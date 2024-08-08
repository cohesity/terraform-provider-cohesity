package cohesity

import (
	"errors"
	"log"
	"strconv"
	"time"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCohesityVirtualEditionCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceCohesityVirtualEditionClusterCreate,
		Read:   resourceCohesityVirtualEditionClusterRead,
		Delete: resourceCohesityVirtualEditionClusterDelete,
		Update: resourceCohesityVirtualEditionClusterUpdate,
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the new Virtual edition cluster",
			},
			"license_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("VIRTUAL_COHESITY_CLUSTER_LICENSE_KEY", ""),
				Description: "Cohesity license key to apply after cluster creation",
			},
			"metadata_fault_tolerance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The metadata fault tolerance",
			},
			"enable_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				Description: `Specifies whether or not to enable encryption.
				 If encryption is enabled, all data on the cluster will be encrypted`,
			},
			"enable_fips_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				Description: `Specifies whether or not to enable FIPS mode.
				 This must be set to true in order to enable FIPS`,
			},
			"encryption_keys_rotation_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "The rotation period for encryption keys in days",
			},
			"cluster_gateway": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The default gateway IP address for the cluster network",
			},
			"cluster_subnet_mask": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The subnet mask of the cluster network",
			},
			"domain_names": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The domain names to configure on the cluster",
			},
			"ntp_servers": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The NTP servers to configure on the cluster",
			},
			"dns_servers": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The DNS servers to configure on the cluster",
			},
			"virtual_ips": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The virtauls IPs for the new cluster",
			},
			"operation_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     120,
				Description: "The time to wait in minutes for cluster creation or destruction",
			},
			"virtual_ip_hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The virtual IP hostname",
			},
			"node_ips": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "IP addresses of the nodes in the cluster",
			},
		},
	}
}

func resourceCohesityVirtualEditionClusterCreate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}

	var clusterName = resourceData.Get("cluster_name").(string)
	var metadataFaultTolerance = int64(resourceData.Get("metadata_fault_tolerance").(int))
	var enableEncryption = resourceData.Get("enable_encryption").(bool)
	var enableFipsMode = resourceData.Get("enable_fips_mode").(bool)
	var encryptionKeysRotationPeriod = int64(resourceData.Get("encryption_keys_rotation_period").(int))
	var clusterGateway = resourceData.Get("cluster_gateway").(string)
	var clusterSubnetMask = resourceData.Get("cluster_subnet_mask").(string)
	var vipHostName = resourceData.Get("virtual_ip_hostname").(string)

	log.Printf("[INFO] Create virtual edition cluster: %s", clusterName)

	domainNames := make([]string, resourceData.Get("domain_names").(*schema.Set).Len())
	for i, name := range resourceData.Get("domain_names").(*schema.Set).List() {
		domainNames[i] = name.(string)
	}

	ntpServers := make([]string, resourceData.Get("ntp_servers").(*schema.Set).Len())
	for i, ntp := range resourceData.Get("ntp_servers").(*schema.Set).List() {
		ntpServers[i] = ntp.(string)
	}

	dnsServers := make([]string, resourceData.Get("dns_servers").(*schema.Set).Len())
	for i, dns := range resourceData.Get("dns_servers").(*schema.Set).List() {
		dnsServers[i] = dns.(string)
	}

	vips := make([]string, resourceData.Get("virtual_ips").(*schema.Set).Len())
	for i, vip := range resourceData.Get("virtual_ips").(*schema.Set).List() {
		vips[i] = vip.(string)
	}

	nodeConfigs := make([]*models.VirtualNodeConfiguration, resourceData.Get("node_ips").(*schema.Set).Len())
	for i, ip := range resourceData.Get("node_ips").(*schema.Set).List() {
		nodeIP := ip.(string)
		nodeConfigs[i] = &models.VirtualNodeConfiguration{
			NodeIp: &nodeIP,
		}
	}

	var requestBody models.CreateVirtualClusterParameters
	requestBody.ClusterName = clusterName
	requestBody.MetadataFaultTolerance = &metadataFaultTolerance
	requestBody.NodeConfigs = nodeConfigs
	requestBody.EncryptionConfig = &models.EncryptionConfiguration{
		EnableEncryption: &enableEncryption,
		EnableFipsMode:   &enableFipsMode,
		RotationPeriod:   &encryptionKeysRotationPeriod,
	}
	requestBody.NetworkConfig = models.NetworkConfiguration{
		ClusterGateway:    &clusterGateway,
		ClusterSubnetMask: &clusterSubnetMask,
		DnsServers:        &dnsServers,
		DomainNames:       &domainNames,
		NtpServers:        &ntpServers,
		VipHostname:       &vipHostName,
		Vips:              &vips,
	}
	result, err := client.Clusters().CreateVirtualCluster(&requestBody)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to create virtual edition cluster")
	}
	log.Printf("[INFO] Create request for virtual edition cluster (%s) successful", clusterName)
	var timeout = resourceData.Get("operation_timeout").(int) * WaitTimeToSeconds
	for timeout > 0 {
		result, err2 := client.Cluster().GetBasicClusterInfo()
		if err2 == nil && result.Name != nil {
			log.Printf("[INFO] Successfully created virtual edition cluster %s", clusterName)
			break
		}
		time.Sleep(30 * time.Second)
		timeout -= 30
	}

	client, err = CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)

	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}

	var requestBodyLicence models.LicenceClusterParameters
	signedTime := int64(time.Now().Unix())
	signedVersion := int64(2)
	requestBodyLicence.LicenseKey = resourceData.Get("license_key").(string)
	requestBodyLicence.SignedByUser = cohesityConfig.clusterUsername
	requestBodyLicence.SignedVersion = &signedVersion
	requestBodyLicence.SignedTime = &signedTime

	err = client.Clusters().ApplyClusterLicence(&requestBodyLicence)

	if err != nil {
		log.Printf("[WARNING] Failed to apply license for virtaul edition cluster %s, %s", clusterName, err.Error())
		resourceData.Set("license_key", "")
	}
	log.Printf("[INFO] Successfully created and applied license to  virtual edition cluster %s", clusterName)
	resourceData.SetId(strconv.FormatInt(*result.ClusterId, 10))
	return resourceCohesityVirtualEditionClusterRead(resourceData, configMetaData)
}

func resourceCohesityVirtualEditionClusterRead(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var clusterName = resourceData.Get("cluster_name").(string)
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}

	result, err := client.Cluster().GetBasicClusterInfo()
	if err == nil && result.Name == nil {
		log.Printf("[INFO] %s virtaul edition cluster doesn't exist", clusterName)
		resourceData.SetId("")
	}
	return nil
}

func resourceCohesityVirtualEditionClusterDelete(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var clusterName = resourceData.Get("cluster_name").(string)
	log.Printf("[INFO] Destroy virtual edition cluster: %s", clusterName)
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	destroyErr := client.Clusters().DestroyCluster()
	if destroyErr != nil {
		log.Printf(destroyErr.Error())
		return errors.New("Failed to destroy virtual edition cluster")
	}
	log.Printf("[INFO] Destroy request for virtual edition cluster (%s) successful", clusterName)
	var timeout = resourceData.Get("operation_timeout").(int) * WaitTimeToSeconds
	for timeout > 0 {
		result, infoErr := client.Cluster().GetBasicClusterInfo()
		if infoErr == nil && result.Name == nil {
			log.Printf("[INFO] Destroyed virtual edition cluster: %s", clusterName)
			break
		}
		time.Sleep(30 * time.Second)
		timeout -= 30
	}
	return nil
}

func resourceCohesityVirtualEditionClusterUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var clusterName = resourceData.Get("cluster_name").(string)
	log.Printf("[INFO] Update virtual edition cluster: %s", clusterName)
	resourceData.Partial(true)
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)

	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	oldLicenceValue, _ := resourceData.GetChange("license_key")
	if resourceData.HasChange("license_key") && oldLicenceValue == "" {
		var requestBodyLicence models.LicenceClusterParameters
		signedTime := int64(time.Now().Unix())
		signedVersion := int64(2)
		requestBodyLicence.LicenseKey = resourceData.Get("license_key").(string)
		requestBodyLicence.SignedByUser = cohesityConfig.clusterUsername
		requestBodyLicence.SignedVersion = &signedVersion
		requestBodyLicence.SignedTime = &signedTime
		err := client.Clusters().ApplyClusterLicence(&requestBodyLicence)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to update virtual edition cluster")
		}
		resourceData.SetPartial("license_key")
		log.Printf("[INFO] Applied license to virtual edition cluster: %s", clusterName)
		return resourceCohesityVirtualEditionClusterRead(resourceData, configMetaData)
	}
	return errors.New("Failed to update virtual edition cluster")
}
