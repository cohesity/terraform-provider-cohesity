package cohesity

// Config has cohesity provider configuration
type Config struct {
	Type       string
	clusterVip      string
	clusterUsername string
	clusterPassword string
	clusterDomain   string
}

//WaitTimeToSeconds is used to convert operation time to seconds
const WaitTimeToSeconds int = 60
