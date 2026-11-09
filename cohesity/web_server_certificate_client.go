package cohesity

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const certificatePath = "/irisservices/api/v1/public/certificates/webServer"

type client struct {
	roots    *x509.CertPool
	endpoint string
	apiKey   string
	http     *http.Client
}
type certificateConfig struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"privateKey,omitempty"`
}

func certificates(s string) ([]*x509.Certificate, error) {
	var result []*x509.Certificate
	rest := []byte(s)
	for len(bytes.TrimSpace(rest)) > 0 {
		block, next := pem.Decode(rest)
		if block == nil || block.Type != "CERTIFICATE" {
			return nil, errors.New("expected PEM CERTIFICATE blocks only")
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, errors.New("invalid X.509 certificate")
		}
		result = append(result, cert)
		rest = next
	}
	if len(result) == 0 {
		return nil, errors.New("certificate chain is empty")
	}
	return result, nil
}
func fingerprint(cert *x509.Certificate) string {
	sum := sha256.Sum256(cert.Raw)
	return hex.EncodeToString(sum[:])
}
func sameCertificates(a, b string) bool {
	aa, e := certificates(a)
	if e != nil {
		return false
	}
	bb, e := certificates(b)
	if e != nil || len(aa) != len(bb) {
		return false
	}
	for i := range aa {
		if !bytes.Equal(aa[i].Raw, bb[i].Raw) {
			return false
		}
	}
	return true
}

func newClient(endpoint, apiKey, ca, bootstrap string) (*client, error) {
	u, err := url.Parse(endpoint)
	if err != nil || u.Scheme != "https" || u.Host == "" || u.User != nil || (u.Path != "" && u.Path != "/") || u.RawQuery != "" || u.Fragment != "" {
		return nil, errors.New("endpoint must be an HTTPS origin without credentials, path, query or fragment")
	}
	if strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("api_key or COHESITY_API_KEY is required")
	}
	roots, err := x509.SystemCertPool()
	if err != nil {
		roots = x509.NewCertPool()
	}
	if ca != "" && !roots.AppendCertsFromPEM([]byte(ca)) {
		return nil, errors.New("ca_certificate_pem contains no certificates")
	}
	config := &tls.Config{MinVersion: tls.VersionTLS12, RootCAs: roots, ServerName: u.Hostname()}
	if bootstrap != "" {
		pinned, e := certificates(bootstrap)
		if e != nil {
			return nil, errors.New("bootstrap_certificate_pem must contain PEM server certificate pins")
		}
		// Normal PKI verification remains active via VerifyConnection. The only exceptions are the explicitly pinned bootstrap leaves.
		config.InsecureSkipVerify = true
		config.VerifyConnection = func(state tls.ConnectionState) error {
			if len(state.PeerCertificates) == 0 {
				return errors.New("server provided no certificate")
			}
			intermediates := x509.NewCertPool()
			for _, c := range state.PeerCertificates[1:] {
				intermediates.AddCert(c)
			}
			_, e := state.PeerCertificates[0].Verify(x509.VerifyOptions{Roots: roots, Intermediates: intermediates, DNSName: u.Hostname()})
			if e == nil {
				return nil
			}
			for _, pin := range pinned {
				if bytes.Equal(state.PeerCertificates[0].Raw, pin.Raw) {
					return nil
				}
			}
			return errors.New("server certificate is neither trusted nor the pinned bootstrap certificate")
		}
	}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = config
	return &client{roots: roots, endpoint: strings.TrimRight(endpoint, "/"), apiKey: apiKey, http: &http.Client{Transport: transport, Timeout: 30 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}}, nil
}

func (c *client) request(ctx context.Context, method string, body *certificateConfig) (string, error) {
	var data io.Reader
	if body != nil {
		encoded, e := json.Marshal(body)
		if e != nil {
			return "", errors.New("could not encode certificate request")
		}
		data = bytes.NewReader(encoded)
	}
	req, e := http.NewRequestWithContext(ctx, method, c.endpoint+certificatePath, data)
	if e != nil {
		return "", e
	}
	req.Header.Set("apiKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	response, e := c.http.Do(req)
	if e != nil {
		return "", errors.New("Cohesity HTTPS request failed; check connectivity and certificate trust")
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("Cohesity certificate API returned HTTP %d", response.StatusCode)
	}
	// PUT responses may echo private keys. Never decode, log or persist those responses.
	if method != http.MethodGet {
		return "", nil
	}
	var result struct {
		Certificate string `json:"certificate"`
	}
	if e = json.NewDecoder(io.LimitReader(response.Body, 4<<20)).Decode(&result); e != nil {
		return "", errors.New("invalid Cohesity certificate API response")
	}
	if result.Certificate == "" {
		return c.servedCertificate(ctx)
	}
	return result.Certificate, nil
}
func (c *client) servedCertificate(ctx context.Context) (string, error) {
	c.http.CloseIdleConnections()
	req, _ := http.NewRequestWithContext(ctx, http.MethodHead, c.endpoint+"/", nil)
	response, e := c.http.Do(req)
	if e != nil {
		return "", errors.New("could not verify the HTTPS certificate served by Cohesity")
	}
	defer response.Body.Close()
	if response.TLS == nil || len(response.TLS.PeerCertificates) == 0 {
		return "", errors.New("Cohesity returned no TLS peer certificate")
	}
	var result strings.Builder
	for _, cert := range response.TLS.PeerCertificates {
		result.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}))
	}
	return result.String(), nil
}
func (c *client) install(ctx context.Context, certificate, key string, restartIris bool) error {
	chain, e := certificates(certificate)
	if e != nil {
		return e
	}
	if _, e = tls.X509KeyPair([]byte(certificate), []byte(key)); e != nil {
		return errors.New("private key does not match the certificate, or is not an unencrypted PEM key")
	}
	host, _ := url.Parse(c.endpoint)
	if e = chain[0].VerifyHostname(host.Hostname()); e != nil {
		return errors.New("certificate SAN does not cover the Cohesity endpoint")
	}
	intermediates := x509.NewCertPool()
	for _, intermediate := range chain[1:] {
		intermediates.AddCert(intermediate)
	}
	if _, err := chain[0].Verify(x509.VerifyOptions{Roots: c.roots, Intermediates: intermediates, DNSName: host.Hostname()}); err != nil {
		return errors.New("replacement certificate chain is not trusted; provide its issuing CA before installation")
	}
	now := time.Now()
	if now.Before(chain[0].NotBefore) || !now.Before(chain[0].NotAfter) {
		return errors.New("certificate is not currently valid")
	}
	// A retry or a reload-policy-only update must not restart an already-correct listener.
	if remote, err := c.request(ctx, http.MethodGet, nil); err == nil && sameCertificates(remote, certificate) {
		if served, err := c.servedCertificate(ctx); err == nil {
			if actual, err := certificates(served); err == nil && bytes.Equal(actual[0].Raw, chain[0].Raw) {
				return nil
			}
		}
	}
	if _, e = c.request(ctx, http.MethodPut, &certificateConfig{Certificate: certificate, PrivateKey: key}); e != nil {
		return e
	}
	if restartIris {
		if e := c.restartIris(ctx); e != nil {
			return e
		}
	}
	// Reloads can disconnect the API. Wait for the actual HTTPS listener, not just the stored API value.
	deadline := time.NewTimer(2 * time.Minute)
	defer deadline.Stop()
	for {
		served, err := c.servedCertificate(ctx)
		if err == nil {
			actual, err := certificates(served)
			if err == nil && bytes.Equal(actual[0].Raw, chain[0].Raw) {
				return nil
			}
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-deadline.C:
			return errors.New("certificate uploaded but HTTPS did not present it within two minutes; verify the cluster reload procedure before retrying")
		case <-time.After(2 * time.Second):
		}
	}
}

// Restart only the UI/REST API service. Never send an empty service list.
func (c *client) restartIris(ctx context.Context) error {
	body := strings.NewReader(`{"action":"kRestart","services":["kIris"]}`)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint+"/irisservices/api/v1/public/clusters/services/states", body)
	if err != nil {
		return errors.New("could not build Iris restart request")
	}
	req.Header.Set("apiKey", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	response, err := c.http.Do(req)
	if err != nil {
		return errors.New("certificate uploaded but Iris restart request failed; verify service state before retrying")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("certificate uploaded but Iris restart returned HTTP %d", response.StatusCode)
	}
	return nil
}
