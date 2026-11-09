package cohesity

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

type fixture struct {
	server         *httptest.Server
	mu             sync.Mutex
	cert, key      string
	active         tls.Certificate
	puts, deletes  int
	restarts       int
	deferredReload bool
	pending        tls.Certificate
}

func issue(t *testing.T) (string, string) {
	t.Helper()
	key, e := rsa.GenerateKey(rand.Reader, 2048)
	if e != nil {
		t.Fatal(e)
	}
	serial, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 120))
	template := &x509.Certificate{SerialNumber: serial, Subject: pkix.Name{CommonName: "127.0.0.1"}, IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(24 * time.Hour), KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}}
	der, e := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if e != nil {
		t.Fatal(e)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})), string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}))
}
func newFixture(t *testing.T) *fixture {
	t.Helper()
	f := &fixture{}
	f.cert, f.key = issue(t)
	f.active, _ = tls.X509KeyPair([]byte(f.cert), []byte(f.key))
	f.server = httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		defer f.mu.Unlock()
		if r.URL.Path == "/" {
			return
		}
		if r.URL.Path == "/irisservices/api/v1/public/clusters/services/states" {
			var action struct {
				Action   string   `json:"action"`
				Services []string `json:"services"`
			}
			if r.Method != http.MethodPost || r.Header.Get("apiKey") != "test-secret" || json.NewDecoder(r.Body).Decode(&action) != nil || action.Action != "kRestart" || len(action.Services) != 1 || action.Services[0] != "kIris" {
				w.WriteHeader(400)
				return
			}
			f.restarts++
			if f.deferredReload {
				f.active = f.pending
			}
			w.WriteHeader(http.StatusAccepted)
			return
		}

		if r.URL.Path != certificatePath {
			http.NotFound(w, r)
			return
		}
		if r.Header.Get("apiKey") != "test-secret" {
			w.WriteHeader(401)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPut:
			var cfg certificateConfig
			if json.NewDecoder(r.Body).Decode(&cfg) != nil {
				w.WriteHeader(400)
				return
			}
			pair, e := tls.X509KeyPair([]byte(cfg.Certificate), []byte(cfg.PrivateKey))
			if e != nil {
				w.WriteHeader(400)
				return
			}
			f.cert, f.key = cfg.Certificate, cfg.PrivateKey
			if f.deferredReload {
				f.pending = pair
			} else {
				f.active = pair
			}
			f.puts++
		case http.MethodDelete:
			f.deletes++
		}
		_ = json.NewEncoder(w).Encode(certificateConfig{Certificate: f.cert, PrivateKey: f.key})
	}))
	f.server.TLS = &tls.Config{GetConfigForClient: func(*tls.ClientHelloInfo) (*tls.Config, error) {
		f.mu.Lock()
		defer f.mu.Unlock()
		return &tls.Config{Certificates: []tls.Certificate{f.active}, MinVersion: tls.VersionTLS12}, nil
	}}
	f.server.StartTLS()
	t.Cleanup(f.server.Close)
	return f
}
func TestInstallAndPinnedBootstrap(t *testing.T) {
	f := newFixture(t)
	original := f.cert
	// The new self-signed fixture leaf stands in for an issued leaf trusted through the configured CA pool.
	cert, key := issue(t)
	c, e := newClient(f.server.URL, "test-secret", cert, original)
	if e != nil {
		t.Fatal(e)
	}
	if e = c.install(context.Background(), cert, key, false); e != nil {
		t.Fatal(e)
	}
	if f.puts != 1 {
		t.Fatalf("expected one PUT, got %d", f.puts)
	}
	got, e := c.request(context.Background(), http.MethodGet, nil)
	if e != nil || !sameCertificates(cert, got) {
		t.Fatalf("readback failed: %v", e)
	}
	// A different self-signed certificate must be rejected once neither CA trust nor pin matches.
	unrelated, unrelatedKey := issue(t)
	f.mu.Lock()
	f.active, _ = tls.X509KeyPair([]byte(unrelated), []byte(unrelatedKey))
	f.mu.Unlock()
	c.http.CloseIdleConnections()
	if _, e = c.servedCertificate(context.Background()); e == nil {
		t.Fatal("accepted an untrusted, unpinned replacement")
	}
}
func TestRejectMismatchedKeyBeforeUpload(t *testing.T) {
	f := newFixture(t)
	c, _ := newClient(f.server.URL, "test-secret", f.cert, "")
	cert, _ := issue(t)
	_, key := issue(t)
	if e := c.install(context.Background(), cert, key, false); e == nil {
		t.Fatal("accepted mismatched key")
	}
	if f.puts != 0 {
		t.Fatal("uploaded invalid certificate")
	}
}
func TestRejectRedirectAndRedactErrors(t *testing.T) {
	f := newFixture(t)
	destination := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { t.Error("followed certificate API redirect") }))
	defer destination.Close()
	f.server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", destination.URL)
		w.WriteHeader(307)
		_, _ = w.Write([]byte("test-secret PRIVATE KEY"))
	})
	c, _ := newClient(f.server.URL, "test-secret", f.cert, "")
	_, e := c.request(context.Background(), http.MethodGet, nil)
	if e == nil || strings.Contains(e.Error(), "test-secret") || strings.Contains(e.Error(), "PRIVATE KEY") {
		t.Fatalf("unsafe error: %v", e)
	}
}
func TestValidateEndpoint(t *testing.T) {
	for _, endpoint := range []string{"http://cluster", "https://user:pass@cluster", "https://cluster/path", "https://cluster?query", "https://cluster/#fragment"} {
		if _, e := newClient(endpoint, "test", "", ""); e == nil {
			t.Errorf("accepted %s", endpoint)
		}
	}
}

func TestRejectUntrustedReplacementBeforeUpload(t *testing.T) {
	f := newFixture(t)
	c, _ := newClient(f.server.URL, "test-secret", f.cert, "")
	cert, key := issue(t)
	if e := c.install(context.Background(), cert, key, false); e == nil {
		t.Fatal("accepted an untrusted replacement certificate")
	}
	if f.puts != 0 {
		t.Fatal("uploaded an untrusted replacement certificate")
	}
}

func TestBootstrapPinBundleAcrossClusterNodes(t *testing.T) {
	f := newFixture(t)
	second, secondKey := issue(t)
	c, err := newClient(f.server.URL, "test-secret", "", f.cert+second)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = c.servedCertificate(context.Background()); err != nil {
		t.Fatal("first cluster node rejected:", err)
	}
	f.mu.Lock()
	f.active, _ = tls.X509KeyPair([]byte(second), []byte(secondKey))
	f.mu.Unlock()
	if _, err = c.servedCertificate(context.Background()); err != nil {
		t.Fatal("second pinned cluster node rejected:", err)
	}
	unrelated, unrelatedKey := issue(t)
	f.mu.Lock()
	f.active, _ = tls.X509KeyPair([]byte(unrelated), []byte(unrelatedKey))
	f.mu.Unlock()
	if _, err = c.servedCertificate(context.Background()); err == nil {
		t.Fatal("accepted a node certificate outside the explicit pin bundle")
	}
}

func TestIrisReloadAndIdempotentRetry(t *testing.T) {
	f := newFixture(t)
	f.deferredReload = true
	cert, key := issue(t)
	c, err := newClient(f.server.URL, "test-secret", cert, f.cert)
	if err != nil {
		t.Fatal(err)
	}
	if err = c.install(context.Background(), cert, key, true); err != nil {
		t.Fatal(err)
	}
	if f.puts != 1 || f.restarts != 1 {
		t.Fatalf("expected one upload and Iris restart; got %d, %d", f.puts, f.restarts)
	}
	if err = c.install(context.Background(), cert, key, true); err != nil {
		t.Fatal(err)
	}
	if f.puts != 1 || f.restarts != 1 {
		t.Fatal("retry re-uploaded or restarted an unchanged certificate")
	}
}
