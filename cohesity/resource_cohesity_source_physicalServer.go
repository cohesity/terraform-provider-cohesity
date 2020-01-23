package cohesity

import (
	"errors"
	"log"
	"strconv"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

//ProtectionSourceID Id of protection source registered.
var ProtectionSourceID int64

func resourceCohesitySourcePhysicalServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCohesitySourcePhysicalServerCreate,
		Read:   resourceCohesitySourcePhysicalServerRead,
		Delete: resourceCohesitySourcePhysicalServerDelete,
		Update: resourceCohesitySourcePhysicalServerUpdate,
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
				Description: `Specifies the network endpoint of the Protection Source where it is reachable. It could be an URL or hostname or 
				an IP address of the Protection Source`,
			},
			"force_register": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Forcefully register physical server to target cluster",
			},
			"host_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "OS type on the physical server.",
			},
			"physical_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host Type for physical server.",
			},
		},
	}
}

func resourceCohesitySourcePhysicalServerCreate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	log.Printf("[INFO] Starting Physical Server Registration")
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	var endpoint = resourceData.Get("endpoint").(string)
	var forceRegister = resourceData.Get("force_register").(bool)

	//Creating Request Body
	var requestBody models.RegisterProtectionSourceParameters
	requestBody.Endpoint = &endpoint
	requestBody.ForceRegister = &forceRegister
	requestBody.Environment = models.EnvironmentRegisterProtectionSourceParameters_KPHYSICAL
	requestBody.HostType = models.HostTypeRegisterProtectionSourceParameters_KLINUX
	requestBody.PhysicalType = models.PhysicalType_KHOST

	log.Printf("[INFO] Register Physical Server as protection source %s", endpoint)
	result, err := client.ProtectionSources().CreateRegisterProtectionSource(&requestBody)

	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to register source Physical Server")
	}
	resourceData.SetId(strconv.FormatInt(*result.Id, 10))
	log.Printf("[INFO] Successfully registered physical server protection source %s", endpoint)
	log.Printf("[INFO] ID %s", resourceData.Id())
	return resourceCohesitySourcePhysicalServerRead(resourceData, configMetaData)
}

func resourceCohesitySourcePhysicalServerRead(resourceData *schema.ResourceData, configMetaData interface{}) error {
	log.Printf("[INFO] Starting Read")
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	var trueValue = true
	var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.EnvironmentListProtectionSources_KPHYSICAL}
	var endpoint = resourceData.Get("endpoint").(string)

	// Converting to int64
	ProtectionSourceID, _ = strconv.ParseInt(resourceData.Id(), 10, 64)
	log.Printf("[INFO] Read Cohesity Physical Server protection source %v, %T", ProtectionSourceID, ProtectionSourceID)

	result, err := client.ProtectionSources().ListProtectionSources(&ProtectionSourceID, nil, nil, &trueValue, &trueValue, &trueValue, environmentType, nil, &trueValue, nil, nil, &trueValue)

	for _, protectionSource := range result {
		log.Printf("[INFO] Protection Name: %s", *protectionSource.ProtectionSource.Name)
		if *protectionSource.ProtectionSource.Name == endpoint {
			log.Printf("[INFO] Found the Physical Server protection source %s on cohesity cluster",
				*protectionSource.ProtectionSource.Name)
			return nil
		}
	}
	log.Printf("[INFO] Couldn't find the Physical Server protection source %s", err)
	resourceData.SetId("")
	return nil
}

func resourceCohesitySourcePhysicalServerDelete(resourceData *schema.ResourceData, configMetaData interface{}) error {
	log.Printf("[INFO] Starting Delete")
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}

	// Converting type to int64
	protectionSourceID, _ := strconv.ParseInt(resourceData.Id(), 10, 64)
	log.Printf("[INFO] Unregistering the Physical Server protection source %s", resourceData.
		Get("endpoint").(string))

	err = client.ProtectionSources().DeleteUnregisterProtectionSource(protectionSourceID)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to unregister Physical Server protection source")
	}
	log.Printf("[INFO] Successfully unregistered the Physical Server protection source %s",
		resourceData.Get("endpoint").(string))
	return nil
}

func resourceCohesitySourcePhysicalServerUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	return errors.New("Update Operation is unavailable for source Physical Server")
}
