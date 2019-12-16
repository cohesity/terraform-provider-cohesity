package cohesity

import (
	"errors"
	"log"
	"strconv"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/golang-collections/collections/queue"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCohesityJobVMware() *schema.Resource {
	return &schema.Resource{
		Create: resourceCohesityJobVMwareCreate,
		Read:   resourceCohesityJobVMwareRead,
		Delete: resourceCohesityJobVMwareDelete,
		Update: resourceCohesityJobVMwareUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the name of the Protection Job",
			},
			"protection_source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the Protection source",
			},
			"timezone": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "America/Los_Angeles",
				Description: "Specify the timezone in the format: Area/Location",
			},
			"include": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Description: `Specifies the list of vm's from the
				protection source to be protected by this protection job`,
			},
			"exclude": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Description: `Specifies the list of vm's from the protection source
				that should not be protected and are excluded from being
				backed up by the protection job`,
			},
			"policy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the protection policy",
			},
			"storage_domain": {
				Type:     schema.TypeString,
				Required: true,
				Description: `Specifies the storage domain where this job
				writes data`,
			},
			"qos_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "BackupHDD",
				Description: `Specifies the QoS policy type to use for this
				Protection Job. 'BackupHDD' indicates the Cohesity Cluster
				writes data directly to the HDD tier for this Protection Job.
				This is the recommended setting. 'BackupSSD' indicates the
				Cohesity Cluster writes data directly to the SSD tier for this
				Protection Job. Only specify this policy if you need fast ingest
				speed for a small number of Protection Jobs`,
			},
			"full_sla": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  120,
				Description: `Specifies the number of minutes that a Job Run
				of a Full (no CBT) backup schedule is expected to complete,
				which is known as a Service-Level Agreement (SLA).A SLA
				violation is reported when the run time of a Job Run exceeds
				the SLA time period specified for this backup schedule.`,
			},
			"incremental_sla": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
				Description: `Specifies the number of minutes that a Job Run
				of a CBT-based backup schedule is expected to complete, which
				is known as a Service-Level Agreement (SLA). A SLA violation
				is reported when the run time of a Job Run exceeds the SLA
				time period specified for this backup schedule.`,
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Medium",
				Description: `Specifies the priority of execution for a
				Protection Job`,
			},
			"delete_snapshots": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `Specifies if snapshots generated by the Protection
				Job should also be deleted when the Job is deleted.
				`,
			},
		},
	}
}

type empty struct{}

//parse the protection source node structure for vm ids
func parseGetSourceIDs(sources map[string]empty, sourceNode []interface{}) []int64 {
	var nodesQueue *queue.Queue
	nodesQueue = queue.New()
	var sourceIDs = []int64{}
	for _, node := range sourceNode {
		nodesQueue.Enqueue(node)
	}

	for nodesQueue.Len() != 0 && len(sources) != 0 {
		node := nodesQueue.Dequeue()
		var name = (((node.(map[string]interface {
		}))["protectionSource"]).(map[string]interface{}))["name"].(string)
		if _, ok := sources[name]; ok {
			id := int64((((node.(map[string]interface {
			}))["protectionSource"]).(map[string]interface{}))["id"].(float64))
			sourceIDs = append(sourceIDs, id)
			delete(sources, name)
			continue
		}
		nodes, ok := node.(map[string]interface{})["nodes"]
		if ok == true && len(nodes.([]interface{})) != 0 {
			for _, node := range nodes.([]interface{}) {
				nodesQueue.Enqueue(node)
			}
		}
	}
	return sourceIDs
}

func resourceCohesityJobVMwareCreate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var cohesityConfig = configMetaData.(Config)
	// authenticate with Cohesity cluster
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	var requestBody models.ProtectionJobRequestBody
	var name = resourceData.Get("name").(string)
	var protectionSource = resourceData.Get("protection_source").(string)
	var fullSLA = int64(resourceData.Get("full_sla").(int))
	var incrementalSLA = int64(resourceData.Get("incremental_sla").(int))
	var timezone = resourceData.Get("timezone").(string)
	var protectionSourceID int64

	//get the policy id
	log.Printf("[INFO] Get the protection policy id")
	policy, err := client.ProtectionPolicies().GetProtectionPolicies(nil,
		[]string{resourceData.Get("policy").(string)}, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Error in getting the protection policy")
	} else if err == nil && len(policy) == 0 {
		return errors.New("Failed to find the protection policy on Cohesity cluster")
	}
	policyID := *policy[0].Id

	//get the storage domain id
	log.Printf("[INFO] Get the storage domain id")
	storageDomain, err := client.ViewBoxes().GetViewBoxes(nil, nil, nil,
		[]string{resourceData.Get("storage_domain").(string)}, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Error in getting the storage domain")
	} else if err == nil && len(storageDomain) == 0 {
		return errors.New("Failed to find the storage domain on Cohesity cluster")
	}
	storageDomainID := *storageDomain[0].Id

	//get the protection source id
	log.Printf("[INFO] Get the protection source id")
	sources, err := client.ProtectionSources().ListProtectionSourcesRootNodes(nil,
		[]models.EnvironmentListProtectionSourcesRootNodesEnum{models.
			EnvironmentListProtectionSourcesRootNodes_KVMWARE}, nil)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Error in getting the protection source")
	}
	for _, source := range sources {
		if *source.ProtectionSource.Name == protectionSource {
			protectionSourceID = *source.ProtectionSource.Id
		}
	}
	if protectionSourceID == 0 {
		return errors.New("Failed to find the protection source on Cohesity cluster")
	}

	//backup specified vm's or enter vCenter
	protectionSourceWithClildObjects, err := client.ProtectionSources().
		ListProtectionSources(&protectionSourceID, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to get the protection source")
	}
	log.Printf("[INFO] Set the source id's for the protection job")
	if len(resourceData.Get("include").(*schema.Set).List()) == 0 {
		var sourceIds = []int64{protectionSourceID}
		requestBody.SourceIds = &sourceIds
	} else {
		var emptyValue empty
		includeSourceSet := make(map[string]empty)
		for _, source := range resourceData.Get("include").(*schema.Set).List() {
			includeSourceSet[source.(string)] = emptyValue
		}
		n := (*protectionSourceWithClildObjects[0].
			Nodes)[0].(map[string]interface{})["nodes"].([]interface{})
		includeSourceIDs := parseGetSourceIDs(includeSourceSet, n)
		if len(resourceData.Get("include").(*schema.Set).List()) != len(includeSourceIDs) {
			return errors.New("Failed to find all the included protection source objects")
		}
		requestBody.SourceIds = &includeSourceIDs
	}

	//exclude the specified vm's
	if len(resourceData.Get("exclude").(*schema.Set).List()) != 0 {
		log.Printf("[INFO] Set the exclude source id's for the protection job")
		var emptyValue empty
		excludeSourceSet := make(map[string]empty)
		for _, source := range resourceData.Get("exclude").(*schema.Set).List() {
			excludeSourceSet[source.(string)] = emptyValue
		}
		n := (*protectionSourceWithClildObjects[0].
			Nodes)[0].(map[string]interface{})["nodes"].([]interface{})
		excludeSourceIDs := parseGetSourceIDs(excludeSourceSet, n)
		if len(resourceData.Get("exclude").(*schema.Set).List()) != len(excludeSourceIDs) {
			return errors.New("Failed to find all the excluded protection source objects")
		}
		requestBody.ExcludeSourceIds = &excludeSourceIDs
	}

	requestBody.Name = name
	requestBody.ParentSourceId = &protectionSourceID
	requestBody.PolicyId = policyID
	requestBody.ViewBoxId = storageDomainID
	requestBody.Timezone = &timezone

	//set the qos type
	log.Printf("[INFO] Set the qos type for the protection job")
	switch resourceData.Get("qos_type").(string) {
	case "BackupHDD":
		requestBody.QosType = models.QosType_KBACKUPHDD
	case "BackupSSD":
		requestBody.QosType = models.QosType_KBACKUPSSD
	}

	//set the priority
	log.Printf("[INFO] Get the priority for the protection job")
	switch resourceData.Get("priority").(string) {
	case "Medium":
		requestBody.Priority = models.Priority_KMEDIUM
	case "Low":
		requestBody.Priority = models.Priority_KLOW
	case "High":
		requestBody.Priority = models.Priority_KHIGH
	}

	//set the SLA
	log.Printf("[INFO] Get the SLA for the protection job")
	requestBody.FullProtectionSlaTimeMins = &fullSLA
	requestBody.IncrementalProtectionSlaTimeMins = &incrementalSLA

	result, err := client.ProtectionJobs().CreateProtectionJob(&requestBody)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to create the protection job")
	}
	resourceData.SetId(strconv.FormatInt(*result.Id, 10))
	log.Printf("[INFO] Successfully created the protection job %s", name)
	return resourceCohesityJobVMwareRead(resourceData, configMetaData)
}

func resourceCohesityJobVMwareRead(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	//parse a decimal string of base 10 and converts into int64
	sourceID, _ := strconv.ParseInt(resourceData.Id(), 10, 64)

	log.Printf("[INFO] Read Cohesity VMware protection job")
	result, err := client.ProtectionJobs().GetProtectionJobs([]int64{sourceID},
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil || len(result) == 0 {
		resourceData.SetId("")
	}
	return nil
}

func resourceCohesityJobVMwareDelete(resourceData *schema.ResourceData, configMetaData interface{}) error {
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}

	//parse a decimal string of base 10 and converts into int64
	jobID, _ := strconv.ParseInt(resourceData.Id(), 10, 64)
	log.Printf("[INFO] Delete VMware protection job %s", resourceData.
		Get("name").(string))
	deleteSnapshots := resourceData.Get("delete_snapshots").(bool)
	deleteProtectionJobParam := models.
		DeleteProtectionJobParam{DeleteSnapshots: &deleteSnapshots}
	err = client.ProtectionJobs().DeleteProtectionJob(jobID, &deleteProtectionJobParam)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to delete the protection job")
	}
	return nil
}

func resourceCohesityJobVMwareUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	resourceData.Partial(true)
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to authenticate with Cohesity")
	}
	log.Printf("[INFO] Update VMware protection job")
	var requestBody models.ProtectionJobRequestBody

	//parse a decimal string of base 10 and converts into int64
	jobID, _ := strconv.ParseInt(resourceData.Id(), 10, 64)
	result, err := client.ProtectionJobs().GetProtectionJobs([]int64{jobID},
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	var name = resourceData.Get("name").(string)
	var timezone = resourceData.Get("timezone").(string)
	var fullSLA = int64(resourceData.Get("full_sla").(int))
	var incrementalSLA = int64(resourceData.Get("incremental_sla").(int))

	requestBody.Name = name
	requestBody.Timezone = &timezone
	requestBody.FullProtectionSlaTimeMins = &fullSLA
	requestBody.IncrementalProtectionSlaTimeMins = &incrementalSLA

	if resourceData.HasChange("protection_source") {
		return errors.New("Can't update the protection source in VMware protection job")
	}
	requestBody.ParentSourceId = result[0].ParentSourceId

	switch resourceData.Get("qos_type").(string) {
	case "BackupHDD":
		requestBody.QosType = models.QosType_KBACKUPHDD
	case "BackupSSD":
		requestBody.QosType = models.QosType_KBACKUPSSD
	}

	switch resourceData.Get("priority").(string) {
	case "Medium":
		requestBody.Priority = models.Priority_KMEDIUM
	case "Low":
		requestBody.Priority = models.Priority_KLOW
	case "High":
		requestBody.Priority = models.Priority_KHIGH
	}

	if resourceData.HasChange("policy") {
		log.Printf("[INFO] Update VMware protection job policy")
		policy, err := client.ProtectionPolicies().GetProtectionPolicies(nil,
			[]string{resourceData.Get("policy").(string)}, nil, nil, nil, nil)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Error in getting the protection policy")
		} else if err == nil && len(policy) == 0 {
			return errors.New("Failed to find the protection policy on Cohesity cluster")
		}
		policyID := *policy[0].Id
		requestBody.PolicyId = policyID

	} else {
		requestBody.PolicyId = result[0].PolicyId
	}

	if resourceData.HasChange("storage_domain") {
		log.Printf("[INFO] Update VMware protection job storage domain")
		storageDomain, err := client.ViewBoxes().GetViewBoxes(nil, nil, nil,
			[]string{resourceData.Get("storage_domain").(string)}, nil, nil, nil)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Error in getting the storage domain")
		} else if err == nil && len(storageDomain) == 0 {
			return errors.New("Failed to find the storage domain on Cohesity cluster")
		}
		storageDomainID := *storageDomain[0].Id
		requestBody.ViewBoxId = storageDomainID

	} else {
		requestBody.ViewBoxId = result[0].ViewBoxId
	}
	protectionSourceWithClildObjects, err := client.ProtectionSources().ListProtectionSources(result[0].
		ParentSourceId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to get the protection source")
	}

	if resourceData.HasChange("include") {
		log.Printf("[INFO] Update VMware protection job sources id's")
		if len(resourceData.Get("include").(*schema.Set).List()) == 0 {
			var sourceIds = []int64{*result[0].ParentSourceId}
			requestBody.SourceIds = &sourceIds
		} else {
			var emptyValue empty
			includeSourceSet := make(map[string]empty)
			for _, source := range resourceData.Get("include").(*schema.Set).List() {
				includeSourceSet[source.(string)] = emptyValue
			}
			n := (*protectionSourceWithClildObjects[0].
				Nodes)[0].(map[string]interface{})["nodes"].([]interface{})
			includeSourceIDs := parseGetSourceIDs(includeSourceSet, n)
			if len(resourceData.Get("include").(*schema.Set).List()) != len(includeSourceIDs) {
				return errors.New("Failed to find all the included protection source objects")
			}
			requestBody.SourceIds = &includeSourceIDs
		}
	} else {
		requestBody.SourceIds = result[0].SourceIds
	}

	if resourceData.HasChange("exclude") {
		log.Printf("[INFO] Update VMware protection job exclude source id's")
		if len(resourceData.Get("exclude").(*schema.Set).List()) != 0 {
			var emptyValue empty
			excludeSourceSet := make(map[string]empty)
			for _, source := range resourceData.Get("exclude").(*schema.Set).List() {
				excludeSourceSet[source.(string)] = emptyValue
			}
			n := (*protectionSourceWithClildObjects[0].
				Nodes)[0].(map[string]interface{})["nodes"].([]interface{})
			excludeSourceIDs := parseGetSourceIDs(excludeSourceSet, n)
			if len(resourceData.Get("exclude").(*schema.Set).List()) != len(excludeSourceIDs) {
				return errors.New("Failed to find all the excluded protection source objects")
			}
			requestBody.ExcludeSourceIds = &excludeSourceIDs
		} else {
			requestBody.ExcludeSourceIds = nil
		}
	} else {
		requestBody.ExcludeSourceIds = result[0].ExcludeSourceIds
	}
	_, err = client.ProtectionJobs().UpdateProtectionJob(&requestBody, *result[0].Id)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to update the VMware protection job")
	}
	resourceData.Partial(false)
	return nil
}
