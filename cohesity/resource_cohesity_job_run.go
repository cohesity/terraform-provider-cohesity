package cohesity

import (
	"context"
	"log"
	"strconv"
	"time"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceCohesityJobRun() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesityJobRunCreate,
		ReadContext:   resourceCohesityJobRunRead,
		DeleteContext: resourceCohesityJobRunDelete,
		UpdateContext: resourceCohesityJobRunUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the name of the Protection Job",
			},
			"run_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Regular",
				Description: "Specifies the type of backup",
			},
			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "start",
				Description: "Specifies whether to start or stop a protection job run",
			},
			"operation_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  120,
				Description: `Specifies the time to wait in minutes for the protection job run
				 to complete the run or stop the run`,
			},
			"timestamp": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the current timestamp to trigger starting or stopping a job run",
			},
		},
	}
}

func jobStartStopUtil(resourceData *schema.ResourceData, configMetaData interface{}) (*int64, diag.Diagnostics) {
	var cohesityConfig = configMetaData.(utils.Config)
	// authenticate with Cohesity cluster
	log.Printf("[INFO] Authenticate with Cohesity cluster")
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.ClusterVIP,
		cohesityConfig.ClusterUsername, cohesityConfig.ClusterPassword, cohesityConfig.ClusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return nil, diag.Errorf("Failed to authenticate with Cohesity")
	}
	jobName := resourceData.Get("name").(string)
	timeout := resourceData.Get("operation_timeout").(int) * utils.WaitTimeToSeconds
	//get the protection job details
	log.Printf("[INFO] Get the protection job (%s) details", jobName)
	protectionJob, err := client.ProtectionJobs().GetProtectionJobs(nil, []string{jobName}, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return nil, diag.Errorf("Failed to find the protection job")
	}
	jobID := *protectionJob[0].Id
	//get the protection runs to check the the status of recent job run
	log.Printf("[INFO] Get the protection job runs for %s", jobName)
	response, err := client.ProtectionRuns().
		GetProtectionRuns(&jobID, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return nil, diag.Errorf("Failed to get the protection job runs")
	}
	//start or stop the job run and wait for completion or operation timeout
	if resourceData.Get("state").(string) == "start" {
		if len(response) == 0 || (response[0].BackupRun.Status != models.StatusBackupRun_KACCEPTED &&
			response[0].BackupRun.Status != models.StatusBackupRun_KRUNNING) {
			log.Printf("[INFO] Start the protection job %s", jobName)
			var requestParams models.RunProtectionJobParam
			switch resourceData.Get("run_type").(string) {
			case "Regular":
				requestParams.RunType = models.RunTypeRunProtectionJobParam_KREGULAR
			case "Full":
				requestParams.RunType = models.RunTypeRunProtectionJobParam_KFULL
			case "Log":
				requestParams.RunType = models.RunTypeRunProtectionJobParam_KLOG
			case "System":
				requestParams.RunType = models.RunTypeRunProtectionJobParam_KSYSTEM
			}
			err = client.ProtectionJobs().CreateRunProtectionJob(jobID, &requestParams)
			time.Sleep(5 * time.Second)
			if err != nil {
				log.Printf(err.Error())
				return nil, diag.Errorf("Failed to run the protection job")
			}
		}
		for timeout > 0 {
			log.Printf("[INFO] Wait for protection job (%s) run completion", jobName)
			result, err := client.ProtectionRuns().
				GetProtectionRuns(&jobID, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
			if err == nil && result[0].BackupRun.Status == models.StatusBackupRun_KSUCCESS {
				log.Printf("[INFO] The protection job (%s) run is successfull", jobName)
				break
			} else if err == nil && result[0].BackupRun.Status == models.StatusBackupRun_KCANCELED {
				return nil, diag.Errorf("The protection job run is canceled")
			} else if err == nil && result[0].BackupRun.Status == models.StatusBackupRun_KFAILURE {
				return nil, diag.Errorf("The protection job run failed")
			}
			time.Sleep(30 * time.Second)
			timeout -= 30
		}
	} else if resourceData.Get("state").(string) == "stop" && len(response) != 0 &&
		(response[0].BackupRun.Status == models.StatusBackupRun_KACCEPTED ||
			response[0].BackupRun.Status == models.StatusBackupRun_KRUNNING) {
		log.Printf("[INFO] Stop the protection job run (%s)", jobName)
		var cancelJobRunParams models.CancelProtectionJobRunParam
		cancelJobRunParams.JobRunId = response[0].BackupRun.JobRunId
		err = client.ProtectionRuns().CreateCancelProtectionJobRun(jobID, &cancelJobRunParams)
		if err != nil {
			log.Printf(err.Error())
			return nil, diag.Errorf("Failed to stop the protection job run")
		}
		for timeout > 0 {
			log.Printf("[INFO] Wait for the protection job (%s) run to be stopped", jobName)
			result, err := client.ProtectionRuns().
				GetProtectionRuns(&jobID, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
			if err == nil && result[0].BackupRun.Status == models.StatusBackupRun_KCANCELED {
				log.Printf("[INFO] The protection job run (%s) is stopped successfully", jobName)
				break
			}
			time.Sleep(30 * time.Second)
			timeout -= 30
		}
	}
	return &jobID, nil
}

func resourceCohesityJobRunCreate(ctx context.Context, resourceData *schema.ResourceData, configMetaData interface{}) diag.Diagnostics {
	jobID, err := jobStartStopUtil(resourceData, configMetaData)
	if err != nil {
		return err
	}
	resourceData.SetId(strconv.FormatInt(*jobID, 10))
	return nil
}

func resourceCohesityJobRunRead(ctx context.Context, resourceData *schema.ResourceData, configMetaData interface{}) diag.Diagnostics {
	return nil
}

func resourceCohesityJobRunDelete(ctx context.Context, resourceData *schema.ResourceData, configMetaData interface{}) diag.Diagnostics {
	return nil
}

func resourceCohesityJobRunUpdate(ctx context.Context, resourceData *schema.ResourceData, configMetaData interface{}) diag.Diagnostics {
	_, err := jobStartStopUtil(resourceData, configMetaData)
	if err != nil {
		return err
	}
	return nil
}
