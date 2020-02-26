package cohesity

// variables for vmware job run acceptance test
var (
	jobRunUpdateName  = "terraform_protect_vcenter_updated"
	jobRunUpdateState = "stop"
)

// variables for job vmware acceptance test
var (
	jobVMwareCreateSourceEndpoint = "vc-67.eco.eng.cohesity.com"
	jobVMwareUpdateSourceEndpoint = "vc-67.eco.eng.cohesity.com"
)

// variables for restore vmware vm acceptance test
var (
	restoreVMSourceEndpoint       = "vc-67.eco.eng.cohesity.com"
	restoreVMProtectionJobInclude = "cirros-automation-vm"
	restoreVMVMNames              = "cirros-automation-vm"
)

// variables for source vmware acceptance test
var (
	sourceVMwareCreateSourceEndpoint = "vc-67.eco.eng.cohesity.com"
	sourceVMwareUpdateSourceEndpoint = "vc-67.eco.eng.cohesity.com"
)

// variables for virtual edition cluster acceptance test
var (
	virtualEditionClusterNodeIP     = "10.2.145.49"
	virtualEditionClusterVip        = "10.2.145.49"
	virtualEditionClusterGateway    = "10.2.144.1"
	virtualEditionClusterSubnetMask = "255.255.240.0"
	virtualEditionDomainName        = "eng.cohesity.com"
	virtualEditionDNSServer         = "10.2.145.14"
)

// variables for cloud edition cluster acceptance test
var (
	cloudEditionClusterNodeIP1    = "10.2.45.143"
	cloudEditionClusterNodeIP2    = "10.2.45.144"
	cloudEditionClusterNodeIP3    = "10.2.45.145"
	cloudEditionClusterGateway    = "10.2.144.1"
	cloudEditionClusterSubnetMask = "255.255.240.0"
	cloudEditionDomainName        = "eng.cohesity.com"
	cloudEditionDNSServer         = "10.2.145.14"
)

// variables for physical edition cluster acceptance test
var (
	physicalEditionClusterNodeIP1     = "10.9.33.113"
	physicalEditionClusterNodeIP2     = "10.9.33.114"
	physicalEditionClusterNodeIP3     = "10.9.33.115"
	physicalEditionClusterNodeIpmiIP1 = "10.9.33.113"
	physicalEditionClusterNodeIpmiIP2 = "10.9.33.114"
	physicalEditionClusterNodeIpmiIP3 = "10.9.33.115"
	physicalEditionClusterVip         = "10.2.145.27"
	physicalEditionClusterGateway     = "10.2.144.1"
	physicalEditionClusterSubnetMask  = "255.255.240.0"
	physicalEditionDomainName         = "eng.cohesity.com"
	physicalEditionDNSServer          = "10.2.145.14"
	ipmiGateway                       = "10.2.144.54"
	ipmiSubnetMask                    = "255.255.240.0"
	ipmiUsername                      = "cohesity"
)
