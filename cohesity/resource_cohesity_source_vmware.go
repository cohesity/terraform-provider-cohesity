package cohesity

import (
	"errors"
	"log"
	"strconv"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCohesitySourceVMware() *schema.Resource {
	return &schema.Resource{
		Create: resourceCohesitySourceVMwareCreate,
		Read:   resourceCohesitySourceVMwareRead,
		Delete: resourceCohesitySourceVMwareDelete,
		Update: resourceCohesitySourceVMwareUpdate,
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
				Description: `Specifies the network endpoint of the Protection Source where it is reachable.
							 It could be an URL or hostname or an IP address of the Protection Source`,
			},
			"vmware_type": {
				Type:        schema.TypeString,
				Default:     "VCenter",
				Optional:    true,
				Description: "Specifies the VMware entity type",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies username to access the target source",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Specifies password of the username to access the target source",
			},
		},
	}
}

func resourceCohesitySourceVMwareCreate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}

	var endpoint = resourceData.Get("endpoint").(string)
	var username = resourceData.Get("username").(string)
	var password = resourceData.Get("password").(string)
	var requestBody models.RegisterProtectionSourceParameters
	requestBody.Endpoint = &endpoint
	requestBody.Username = &username
	requestBody.Password = &password

	switch resourceData.Get("vmware_type").(string) {
	case "VCenter":
		requestBody.VmwareType = models.VmwareType_KVCENTER
	case "VCloudDirector":
		requestBody.VmwareType = models.VmwareType_KVCLOUDDIRECTOR
	case "HostSystem":
		requestBody.VmwareType = models.VmwareType_KHOSTSYSTEM
	}

	requestBody.Environment = models.EnvironmentRegisterProtectionSourceParameters_KVMWARE
	log.Printf("[INFO] Register VMware protection source %s", endpoint)
	result, err := client.ProtectionSources().CreateRegisterProtectionSource(&requestBody)

	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to register Cohesity protection source")
	}
	resourceData.SetId(strconv.FormatInt(*result.Id, 10))
	log.Printf("[INFO] Successfully registered VMware protection source %s", endpoint)
	return resourceCohesitySourceVMwareRead(resourceData, configMetaData)
}

func resourceCohesitySourceVMwareRead(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.EnvironmentListProtectionSources_KVMWARE}
	var trueValue = true
	var endpoint = resourceData.Get("endpoint").(string)
	log.Printf("[INFO] Read Cohesity VMware protection source %s", endpoint)
	result, err := client.ProtectionSources().ListProtectionSources(nil, nil, nil, &trueValue, &trueValue, &trueValue, environmentType, nil, nil, nil, nil, nil)
	for _, protectionSource := range result {
		if *protectionSource.ProtectionSource.Name == endpoint {
			log.Printf("[INFO] Found the VMware protection source %s on cohesity cluster", *protectionSource.ProtectionSource.Name)
			return nil
		}
	}
	log.Printf("[INFO] Couldn't find the VMware protection source %s", endpoint)
	resourceData.SetId("")
	return nil
}

func resourceCohesitySourceVMwareDelete(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	sourceID, _ := strconv.ParseInt(resourceData.Id(), 10, 64)
	log.Printf("[INFO] Unregistering the VMware protection source %s", resourceData.Get("endpoint").(string))
	err = client.ProtectionSources().DeleteUnregisterProtectionSource(sourceID)
	if err != nil {
		log.Printf(err.Error())
		log.Printf("[INFO] Failed to unregister VMware protection source %s", resourceData.Get("endpoint").(string))
	}
	log.Printf("[INFO] Successfully unregistered the VMware protection source %s", resourceData.Get("endpoint").(string))
	return nil
}

func resourceCohesitySourceVMwareUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	resourceData.Partial(true)
	return errors.New("Protection source update is not supported")
}
