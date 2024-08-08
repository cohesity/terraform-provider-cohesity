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

func resourceCohesityPhysicalEditionCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceCohesityPhysicalEditionClusterCreate,
		Read:   resourceCohesityPhysicalEditionClusterRead,
		Delete: resourceCohesityPhysicalEditionClusterDelete,
		Update: resourceCohesityPhysicalEditionClusterUpdate,
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the new Physical edition cluste",
			},
			"license_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PHYSICAL_COHESITY_CLUSTER_LICENSE_KEY", ""),
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
			"ipmi_gateway": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The default gateway IP address for the IPMI network",
			},
			"ipmi_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PHYSICAL_COHESITY_CLUSTER_IPMI_PASSWORD", ""),
				Description: "The IPMI password",
			},
			"ipmi_subnet_mask": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The subnet mask for the IPMI network",
			},
			"ipmi_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PHYSICAL_COHESITY_CLUSTER_IPMI_USERNAME", ""),
				Description: "The IPMI username",
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
			"virtual_ip_hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The virtual IP hostname",
			},
			"operation_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     120,
				Description: "The time to wait in minutes for cluster creation or destruction",
			},
			"node_configs": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The configuration for the nodes in the new cluster",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address for this node",
						},
						"node_ipmi_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IPMI IP for this node",
						},
					},
				},
			},
		},
	}
}

func resourceCohesityPhysicalEditionClusterCreate(resourceData *schema.ResourceData, configMetaData interface{}) error {
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
	var ipmiGateway = resourceData.Get("ipmi_gateway").(string)
	var ipmiPassword = resourceData.Get("ipmi_password").(string)
	var ipmiSubnetMask = resourceData.Get("ipmi_subnet_mask").(string)
	var ipmiUsername = resourceData.Get("ipmi_username").(string)

	log.Printf("[INFO] Create physical edition cluster: %s", clusterName)

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

	nodeConfigs := make([]*models.PhysicalNodeConfiguration, resourceData.Get("node_configs").(*schema.Set).Len())
	for i, config := range resourceData.Get("node_configs").(*schema.Set).List() {
		nodeConfig := config.(map[string]interface{})
		nodeIP := nodeConfig["node_ip"].(string)
		NodeIpmiIP := nodeConfig["node_ipmi_ip"].(string)
		nodeConfigs[i] = &models.PhysicalNodeConfiguration{
			NodeIp:     &nodeIP,
			NodeIpmiIp: &NodeIpmiIP,
		}
	}

	var requestBody models.CreatePhysicalClusterParameters
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

	requestBody.IpmiConfig = models.IpmiConfiguration{
		IpmiGateway:    &ipmiGateway,
		IpmiPassword:   &ipmiPassword,
		IpmiSubnetMask: &ipmiSubnetMask,
		IpmiUsername:   &ipmiUsername,
	}
	result, err := client.Clusters().CreatePhysicalCluster(&requestBody)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to create physical edition cluster")
	}
	log.Printf("[INFO] Create request for physical edition cluster (%s) successful", clusterName)
	var timeout = resourceData.Get("operation_timeout").(int) * WaitTimeToSeconds
	for timeout > 0 {
		result, err2 := client.Cluster().GetBasicClusterInfo()
		if err2 == nil && result.Name != nil {
			log.Printf("[INFO] Successfully created physical edition cluster %s", clusterName)
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
		log.Printf("[WARNING] Failed to apply license for physical edition cluster %s, %s", clusterName, err.Error())
		resourceData.Set("license_key", "")
	}
	resourceData.SetId(strconv.FormatInt(*result.ClusterId, 10))
	log.Printf("[INFO] Successfully created and applied license to physical edition cluster %s", clusterName)
	return resourceCohesityPhysicalEditionClusterRead(resourceData, configMetaData)
}

func resourceCohesityPhysicalEditionClusterRead(resourceData *schema.ResourceData, configMetaData interface{}) error {
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
		log.Printf("[INFO] %s physical edition cluster doesn't exist", clusterName)
		resourceData.SetId("")
	}
	return nil
}

func resourceCohesityPhysicalEditionClusterDelete(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var clusterName = resourceData.Get("cluster_name").(string)
	log.Printf("[INFO] Destroy physical edition cluster: %s", clusterName)
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
		return errors.New("Failed to destroy physical edition cluster")
	}
	log.Printf("[INFO] Destroy request for physical edition cluster (%s) successful", clusterName)
	var timeout = resourceData.Get("operation_timeout").(int) * WaitTimeToSeconds
	for timeout > 0 {
		result, infoErr := client.Cluster().GetBasicClusterInfo()
		if infoErr == nil && result.Name == nil {
			log.Printf("[INFO] Destroyed physical edition cluster: %s", clusterName)
			break
		}
		time.Sleep(30 * time.Second)
		timeout -= 30
	}
	return nil
}

func resourceCohesityPhysicalEditionClusterUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var clusterName = resourceData.Get("cluster_name").(string)
	log.Printf("[INFO] Update physical edition cluster: %s", clusterName)
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
			return errors.New("Failed to update physical edition cluster")
		}
		resourceData.SetPartial("license_key")
		log.Printf("[INFO] Applied license to physical edition cluster: %s", clusterName)
		return resourceCohesityPhysicalEditionClusterRead(resourceData, configMetaData)
	}
	return errors.New("Failed to update physical edition cluster")
}
