package cohesity

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCohesityRestoreVMwareVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceCohesityRestoreVMwareVMCreate,
		Read:   resourceCohesityRestoreVMwareVMRead,
		Delete: resourceCohesityRestoreVMwareVMDelete,
		Update: resourceCohesityRestoreVMwareVMUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the name of the restore task",
			},
			"job_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the name of the protection job that backed up the objects to be restored",
			},
			"backup_timestamp": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Specifies the time of the protection job run.
				 Should be in the format YYYY-MM-DD HH:MM Area/Location`,
			},
			"vm_names": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Specifies the names of the virtual machines to restore",
			},
			"operation_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     120,
				Description: "The time to wait in minutes for restore task start/stop operation",
			},
			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "start",
				ForceNew:    true,
				Description: "Specifies whether to start or stop a restore task",
			},
			"vmware_parameters": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Specifies vmware parameters for the restore task",
			},
		},
	}
}

// Used in converting mircoseconds to seconds
const epochTimestampToSeconds = 1000000

func restoreStartStopUtil(resourceData *schema.ResourceData, configMetaData interface{}) (*int64, error) {
	// Authenticate with Cohesity cluster
	log.Printf("[INFO] Authenticate with Cohesity cluster")
	var cohesityConfig = configMetaData.(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return nil, errors.New("Failed to authenticate with Cohesity")
	}

	timeout := resourceData.Get("operation_timeout").(int) * WaitTimeToSeconds
	var taskID int64
	var restoreTaskDetails models.RestoreTask

	// Get the restore tasks and check if a task with given name exists
	log.Printf("[INFO] Get restore tasks")
	result, err := client.RestoreTasks().GetRestoreTasks(nil, nil,
		nil, nil, []string{"kRecoverVMs"}, models.EnvironmentGetRestoreTasks_KVMWARE)
	for _, restoreTask := range result {
		if restoreTask.Name == resourceData.Get("name").(string) {
			log.Printf("[INFO] The restore task %s exists", resourceData.Get("name").(string))
			taskID = *restoreTask.Id
			restoreTaskDetails = *restoreTask
		}
	}

	// Start or stop a restore task based on the state
	if resourceData.Get("state").(string) == "start" {
		log.Printf("[INFO] Creating the restore task %s", resourceData.Get("name").(string))
		if taskID != 0 {
			return nil, fmt.Errorf("Failed to create the restore task. The restore task (%s) already exists",
				resourceData.Get("name").(string))
		}
		var recoverRequest models.RecoverTaskRequest
		var taskName = resourceData.Get("name").(string)

		// Get the protection job id
		log.Printf("[INFO] Get id of protection job %s", resourceData.Get("job_name").(string))
		protectionJobs, err := client.ProtectionJobs().GetProtectionJobs(nil,
			[]string{resourceData.Get("job_name").(string)}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return nil, fmt.Errorf("Failed to get protection job from Cohesity cluster")
		}
		if len(protectionJobs) == 0 {
			return nil, fmt.Errorf("Failed to find the protection job %s", resourceData.Get("job_name").(string))
		}
		protectionJobID := *protectionJobs[0].Id

		var jobRunID int64
		var startTime int64
		objectList := make([]*models.RestoreObjectDetails, 0)
		var foundObject = false

		// Search for vm's in the snapshots and construct the objects list
		if len(resourceData.Get("vm_names").(*schema.Set).List()) != 0 {
			log.Printf("[INFO] Start searching for given vm's snapshots")
			for _, vm := range resourceData.Get("vm_names").(*schema.Set).List() {
				vmName := vm.(string)
				objects, err := client.RestoreTasks().SearchObjects(&vmName,
					[]int64{protectionJobID}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
				if err != nil {
					log.Printf(err.Error())
					return nil, fmt.Errorf("Error in getting the snapshot object for vm %s", vmName)
				}
				for _, snapshotObject := range objects.ObjectSnapshotInfo {
					if *snapshotObject.ObjectName == vmName {
						// If backup timestamp is not given, latest snaphot is considered
						if resourceData.Get("backup_timestamp").(string) == "" {
							log.Printf("[INFO] Found the snapshot for vm %s", vmName)
							var object models.RestoreObjectDetails
							object.JobId = &protectionJobID
							object.JobRunId = snapshotObject.Versions[0].JobRunId
							object.StartedTimeUsecs = snapshotObject.Versions[0].StartedTimeUsecs
							object.ProtectionSourceId = snapshotObject.SnapshottedSource.Id
							objectList = append(objectList, &object)
							foundObject = true
							break

						} else {
							userDateTime := strings.Split(resourceData.Get("backup_timestamp").(string), " ")
							location, _ := time.LoadLocation(userDateTime[2])
							// Search for snapshot with given backup timestamp
							for _, snapshotObjectVersion := range snapshotObject.Versions {
								backupDateTime := strings.Split(time.Unix((*snapshotObjectVersion.
									StartedTimeUsecs)/epochTimestampToSeconds, 0).In(location).String(), " ")
								if backupDateTime[0]+" "+backupDateTime[1][:5] == userDateTime[0]+" "+userDateTime[1] {
									log.Printf("[INFO] Found the snapshot for vm %s", vmName)
									var object models.RestoreObjectDetails
									object.JobId = &protectionJobID
									object.JobRunId = snapshotObjectVersion.JobRunId
									object.StartedTimeUsecs = snapshotObjectVersion.StartedTimeUsecs
									object.ProtectionSourceId = snapshotObject.SnapshottedSource.Id
									objectList = append(objectList, &object)
									foundObject = true
									break
								}
							}
						}
					}
				}
				if !foundObject {
					return nil, fmt.Errorf("Failed to find the snapshot for vm %s", vmName)
				}
				foundObject = false
			}
		} else {
			// If no VM's are given by the user, restore all the vm's in the protection job
			log.Printf("[INFO] No VM's are given. All vm's backed by the job run are restored")
			excludeRunsWithErrors := true
			foundObject = false
			log.Printf("[INFO] Get the protection job runs")
			protectionRuns, err := client.ProtectionRuns().GetProtectionRuns(&protectionJobID,
				nil, nil, nil, nil, nil, nil, nil, nil, &excludeRunsWithErrors, nil)
			if err != nil {
				log.Printf("[ERROR] %s", err.Error())
				return nil, fmt.Errorf("Failed to get the protection job runs")
			} else if len(protectionRuns) == 0 {
				return nil, fmt.Errorf("There are no successful job runs for protection job %s",
					resourceData.Get("job_name").(string))
			}
			for _, protectionRun := range protectionRuns {
				// When backup timestamp is not given, latest successful backup is considered
				if resourceData.Get("backup_timestamp").(string) == "" {
					if protectionRun.BackupRun.Status == models.StatusBackupRun_KSUCCESS {
						log.Printf("[INFO] Latest successful protection job run is used for restore")
						startTime = *protectionRun.BackupRun.Stats.StartTimeUsecs
						jobRunID = *protectionRun.BackupRun.JobRunId
						foundObject = true
						break
					}
				} else {
					userDateTime := strings.Split(resourceData.Get("backup_timestamp").(string), " ")
					location, _ := time.LoadLocation(userDateTime[2])
					backupDateTime := strings.Split(time.Unix((*protectionRun.BackupRun.Stats.
						StartTimeUsecs)/epochTimestampToSeconds, 0).In(location).String(), " ")
					if backupDateTime[0]+" "+backupDateTime[1][:5] == userDateTime[0]+" "+userDateTime[1] &&
						protectionRun.BackupRun.Status == models.StatusBackupRun_KSUCCESS {
						log.Printf("[INFO] Found a protection job run with given timestamp")
						startTime = *protectionRun.BackupRun.Stats.StartTimeUsecs
						jobRunID = *protectionRun.BackupRun.JobRunId
						foundObject = true
						break
					}
				}
			}
			if !foundObject {
				return nil, fmt.Errorf("Failed to find job run for restore")
			}
			var object models.RestoreObjectDetails
			object.JobId = &protectionJobID
			object.JobRunId = &jobRunID
			object.StartedTimeUsecs = &startTime
			objectList = append(objectList, &object)
		}
		recoverRequest.Name = taskName
		recoverRequest.Objects = objectList
		recoverRequest.Type = models.TypeRecoverTaskRequest_KRECOVERVMS
		// Set vmware parameters
		log.Printf("[INFO] Setting the VMware parameters")
		if len(resourceData.Get("vmware_parameters").(map[string]interface{})) != 0 {
			var vmwareParameters models.VmwareRestoreParameters
			vmParams := resourceData.Get("vmware_parameters").(map[string]interface{})
			if prefix, ok := vmParams["prefix"]; ok {
				vmPrefix := prefix.(string)
				vmwareParameters.Prefix = &vmPrefix
			}
			if suffix, ok := vmParams["suffix"]; ok {
				vmSuffix := suffix.(string)
				vmwareParameters.Suffix = &vmSuffix
			}
			recoverRequest.VmwareParameters = &vmwareParameters
		}
		log.Printf("[INFO] Request for restore task creation")
		result, err := client.RestoreTasks().CreateRecoverTask(&recoverRequest)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return nil, fmt.Errorf("Failed to create restore task %s", taskName)
		}
		// Wait for restore task completion
		log.Printf("[INFO] Wait till restore task is complete")
		for timeout > 0 {
			result, err := client.RestoreTasks().GetRestoreTaskById(*result.Id)
			if err == nil && result.Status == models.StatusRestoreTask_KFINISHED &&
				result.Error == nil {
				log.Printf("[INFO] The restore task (%s) is successful", taskName)
				break
			} else if err == nil && result.Status == models.StatusRestoreTask_KFINISHED &&
				result.Error != nil {
				return nil, fmt.Errorf("The restore task is created but the task run errored out")
			} else if err == nil && result.Status == models.StatusRestoreTask_KCANCELLED {
				return nil, fmt.Errorf("The restore task is started, but eventually cancelled")
			}
			time.Sleep(30 * time.Second)
			timeout -= 30
		}

		return result.Id, nil

	} else if resourceData.Get("state") == "stop" {

		log.Printf("[INFO] Cancel the restore task %s", resourceData.Get("name").(string))
		if taskID == 0 {
			return nil, fmt.Errorf("The restore task doesn't exist")
		}
		if restoreTaskDetails.Status == models.StatusRestoreTask_KCANCELLED ||
			restoreTaskDetails.Status == models.StatusRestoreTask_KFINISHED {
			return nil, fmt.Errorf("The restore task is already finished or cancelled")
		}
		err = client.RestoreTasks().UpdateCancelRestoreTask(taskID)
		if err != nil {
			log.Printf("[ERROR] %s", err.Error())
			return nil, fmt.Errorf("Failed to cancel the restore task")
		}
		// Wait till restore task is cancelled
		log.Printf("[INFO] Wait till the restore task is cancelled")
		for timeout > 0 {
			result, err := client.RestoreTasks().GetRestoreTaskById(taskID)
			if err == nil && result.Status == models.StatusRestoreTask_KFINISHED {
				return nil, fmt.Errorf("Failed to cancel the restore task")
			} else if err == nil && result.Status == models.StatusRestoreTask_KCANCELLED {
				log.Printf("[INFO] Successfully cancelled the restore task")
				return &taskID, nil
			}
			time.Sleep(30 * time.Second)
			timeout -= 30
		}

		return &taskID, nil
	}
	return nil, fmt.Errorf("Invalid restore task state. Should be start or stop")
}

func resourceCohesityRestoreVMwareVMCreate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	taskID, err := restoreStartStopUtil(resourceData, configMetaData)
	if err != nil {
		return err
	}
	resourceData.SetId(strconv.FormatInt(*taskID, 10))
	return nil
}

func resourceCohesityRestoreVMwareVMRead(resourceData *schema.ResourceData, configMetaData interface{}) error {
	return nil
}

func resourceCohesityRestoreVMwareVMDelete(resourceData *schema.ResourceData, configMetaData interface{}) error {
	return nil
}

func resourceCohesityRestoreVMwareVMUpdate(resourceData *schema.ResourceData, configMetaData interface{}) error {
	return nil
}
