package v2

type Cluster struct {
	Id               int64         `json:"id"`
	IncarnationId    int64         `json:"incarnationId"`
	Name             string        `json:"name"`
	Type             string        `json:"type"`
	ClusterSize      string        `json:"clusterSize"`
	SwVersion        string        `json:"swVersion"`
	NetworkConfig    NetworkConfig `json:"networkConfig"`
	EnableEncryption bool          `json:"enableEncryption"`
}

type NetworkConfig struct {
	NtpServers          []string            `json:"ntpServers"`
	DomainNames         []string            `json:"domainNames"`
	VipHostName         string              `json:"vipHostName"`
	UseDhcp             bool                `json:"useDhcp"`
	ManualNetworkConfig ManualNetworkConfig `json:"manualNetworkConfig"`
}

type ManualNetworkConfig struct {
	Gateway    string   `json:"gateway"`
	SubnetIp   string   `json:"subnetIp"`
	SubnetMask string   `json:"subnetMask"`
	DnsServers []string `json:"dnsServers"`
}
