package utils

// Config has cohesity provider configuration
type Config struct {
	Endpoint                string
	APIKey                  string
	CACertificatePEM        string
	BootstrapCertificatePEM string
	ClusterVIP              string
	ClusterUsername         string
	ClusterPassword         string
	SupportPassword         string
	ClusterDomain           string
}
