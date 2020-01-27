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
				Description: `Specifies the network endpoint of the Protection
				Source where it is reachable. It could be an URL or hostname or
				an IP address of the Protection Source
							  `,
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
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_SOURCE_VMWARE_PASSWORD", ""),
				Description: "Specifies password of the username to access the target source",
			},
			"enable_ssl_verification": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: `Specifies whether SSL verification should be performed.`,
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Default:     "",
				Description: "The contents of CA certificate",
			},
			"cap_streams_per_datastore": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `Specifies whether datastore streams are configured
				for all datastores that are part of the registered entity. If set
				to true, number of streams from Cohesity cluster to the registered
				entity will be limited to the value set for number_of_streams. If
				not set or set to false, there is no max limit for the number of 
				concurrent streams.`,
			},
			"number_of_streams": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				Description: `Specifies the limit on the number of streams
				Cohesity cluster will make concurrently to the datastores
				of the registered entity. This limit is enforced only when the
				cap_streams_per_datastore is set to true`,
			},
			"enable_latency_throttling": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `Indicates whether read operations to the datastores,
				which are part of the registered Protection Source, are throttled.`,
			},
			"new_task_latency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
				Description: `If the latency of a datastore is above this value,
				then new backup tasks using the datastore will not be started.`,
			},
			"active_task_latency": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
				Description: `If the latency of a datastore is above this value,
				existing backup tasks using the datastore are throttled.`,
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
	var sslVerification models.SslVerification
	if resourceData.Get("enable_ssl_verification").(bool) {
		var caCertificate = resourceData.Get("ca_certificate").(string)
		var isEnabled = true
		sslVerification.CaCertificate = &caCertificate
		sslVerification.IsEnabled = &isEnabled
		requestBody.SslVerification = &sslVerification
	}

	var capStreamsPerDatastore = resourceData.Get("cap_streams_per_datastore").(bool)
	var enableLatencyThrottling = resourceData.Get("enable_latency_throttling").(bool)
	var throttlingPolicy models.ThrottlingPolicyParameters

	if capStreamsPerDatastore {
		var numberOfStreams = int64(resourceData.Get("number_of_streams").(int))
		throttlingPolicy.EnforceMaxStreams = &capStreamsPerDatastore
		throttlingPolicy.MaxConcurrentStreams = &numberOfStreams
	}

	if enableLatencyThrottling {
		var newTaskLatency = int64(resourceData.Get("new_task_latency").(int))
		var activeTaskLatency = int64(resourceData.Get("active_task_latency").(int))
		var latencyThresholds models.LatencyThresholds
		latencyThresholds.ActiveTaskMsecs = &activeTaskLatency
		latencyThresholds.NewTaskMsecs = &newTaskLatency
		throttlingPolicy.LatencyThresholds = &latencyThresholds
		throttlingPolicy.IsEnabled = &enableLatencyThrottling
	}

	if capStreamsPerDatastore || enableLatencyThrottling {
		requestBody.ThrottlingPolicy = &throttlingPolicy
	}

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
	var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.
		EnvironmentListProtectionSources_KVMWARE}
	var trueValue = true
	var endpoint = resourceData.Get("endpoint").(string)
	log.Printf("[INFO] Read Cohesity VMware protection source %s", endpoint)
	result, err := client.ProtectionSources().
		ListProtectionSources(nil, nil, nil, &trueValue, &trueValue,
			&trueValue, environmentType, nil, nil, nil, nil, nil)
	for _, protectionSource := range result {
		if *protectionSource.ProtectionSource.Name == endpoint {
			log.Printf("[INFO] Found the VMware protection source %s on cohesity cluster",
				*protectionSource.ProtectionSource.Name)
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

	//parse a decimal string of base 10 and converts into int64
	sourceID, _ := strconv.ParseInt(resourceData.Id(), 10, 64)

	log.Printf("[INFO] Unregistering the VMware protection source %s", resourceData.
		Get("endpoint").(string))

	err = client.ProtectionSources().DeleteUnregisterProtectionSource(sourceID)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to unregister VMware protection source")
	}
	log.Printf("[INFO] Successfully unregistered the VMware protection source %s",
		resourceData.Get("endpoint").(string))
	return nil
}

func resourceCohesitySourceVMwareUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	resourceData.Partial(true)
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}

	if resourceData.HasChange("vmwareType") {
		return errors.New("Can't update vmware type of the protection source")
	}

	var updatedParameters models.UpdateProtectionSourceParameters
	var endpoint = resourceData.Get("endpoint").(string)
	var username = resourceData.Get("username").(string)
	var password = resourceData.Get("password").(string)
	updatedParameters.Endpoint = &endpoint
	updatedParameters.Username = &username
	updatedParameters.Password = &password

	var sslVerification models.SslVerification
	var enableSslVerification = resourceData.Get("enable_ssl_verification").(bool)
	sslVerification.IsEnabled = &enableSslVerification
	if enableSslVerification {
		var caCertificate = resourceData.Get("ca_certificate").(string)
		sslVerification.CaCertificate = &caCertificate
	} else {
		resourceData.Set("ca_certificate", "")
	}
	updatedParameters.SslVerification = &sslVerification

	var capStreamsPerDatastore = resourceData.Get("cap_streams_per_datastore").(bool)
	var enableLatencyThrottling = resourceData.Get("enable_latency_throttling").(bool)
	var throttlingPolicy models.ThrottlingPolicyParameters
	throttlingPolicy.EnforceMaxStreams = &capStreamsPerDatastore
	throttlingPolicy.IsEnabled = &enableLatencyThrottling
	if capStreamsPerDatastore {
		var numberOfStreams = int64(resourceData.Get("number_of_streams").(int))
		throttlingPolicy.MaxConcurrentStreams = &numberOfStreams
	} else {
		//set number of streams to default when cap for datastore steams is disabled
		resourceData.Set("number_of_streams", 1)
	}

	if enableLatencyThrottling {
		var newTaskLatency = int64(resourceData.Get("new_task_latency").(int))
		var activeTaskLatency = int64(resourceData.Get("active_task_latency").(int))
		var latencyThresholds models.LatencyThresholds
		latencyThresholds.ActiveTaskMsecs = &activeTaskLatency
		latencyThresholds.NewTaskMsecs = &newTaskLatency
		throttlingPolicy.LatencyThresholds = &latencyThresholds
	} else {
		//set latency values to default when latency throttling is disabled
		resourceData.Set("new_task_latency", 30)
		resourceData.Set("active_task_latency", 30)
	}
	updatedParameters.ThrottlingPolicy = &throttlingPolicy
	log.Printf("[INFO] Update VMware protection source %s", endpoint)

	//parse a decimal string of base 10 and converts into int64
	sourceID, _ := strconv.ParseInt(resourceData.Id(), 10, 64)

	_, err = client.ProtectionSources().UpdateProtectionSource(sourceID, &updatedParameters)

	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to update VMware protection source")
	}
	resourceData.Partial(false)
	return nil
}
