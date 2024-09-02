package cohesity

// Config has cohesity provider configuration
type Config struct {
	ClusterVIP      string
	ClusterUsername string
	ClusterPassword string
	SupportPassword string
	ClusterDomain   string
}

// WaitTimeToSeconds is used to convert operation time to seconds
const WaitTimeToSeconds int = 60
