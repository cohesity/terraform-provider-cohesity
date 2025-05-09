package cohesity

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceCohesityBackupPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesityBackupPolicyCreate,
		ReadContext:   resourceCohesityBackupPolicyRead,
		UpdateContext: resourceCohesityBackupPolicyUpdate,
		DeleteContext: resourceCohesityBackupPolicyDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     backupPolicySchema(),
			},
			"retry_options": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retries": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3,
							Description: "Specifies the number of times to retry capturing Snapshots before the Protection Group Run fails.",
						},
						"retry_interval_mins": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
							Description: "Specifies the number of minutes before retrying a failed Protection Group.",
						},
					},
				},
			},
			"data_lock": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Compliance", "Administrative",
				}, false),
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"is_cbs_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"last_modification_time_usecs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// "rpo_policy_settings": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	MaxItems: 1,
			// 	Elem:     rpoPolicySettingsSchema(),
			// },
			"skip_interval_mins": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enable_smart_local_retention_adjustment": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func backupPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regular": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"incremental": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem:     incrementalScheduleSchema(),
						},
						"full": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem:     fullScheduleSchema(),
						},
						"full_backups": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedule": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem:     fullScheduleSchema(),
									},
									"retention": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										Elem:     retentionSchema(),
									},
								},
							},
						},
						"retention": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     retentionSchema(),
						},
					},
				},
			},
			"log": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     logScheduleSchema(),
						},
						"retention": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     retentionSchema(),
						},
					},
				},
			},
			"bmr": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     fullScheduleSchema(),
						},
						"retention": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     retentionSchema(),
						},
					},
				},
			},
			"cdp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retention": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     cdpRetentionSchema(),
						},
					},
				},
			},
			"storage_array_snapshot": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     incrementalScheduleSchema(),
						},
						"retention": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem:     retentionSchema(),
						},
					},
				},
			},
			"run_timeouts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"backup_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"kRegular", "kFull", "kLog", "kSystem", "kHydrateCDP", "kStorageArraySnapshot"}, false),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func logScheduleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"unit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"Minutes", "Hours"}, false),
				Required:     true,
			},
			"minute_schedule": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     frequencyScheduleSchema(),
			},
			"hour_schedule": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     frequencyScheduleSchema(),
			},
		},
	}
}
func incrementalScheduleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"unit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"Minutes", "Hours", "Days", "Weeks", "Months", "Years"}, false),
				Required:     true,
			},
			"minute_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     frequencyScheduleSchema(),
			},
			"hour_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     frequencyScheduleSchema(),
			},
			"day_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     frequencyScheduleSchema(),
			},
			"week_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     weekScheduleSchema(),
			},
			"month_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     monthScheduleSchema(),
			},
			"year_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     yearScheduleSchema(),
			},
		},
	}
}
func retentionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"unit": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Specifies the Retention Unit of a backup measured in days, months or years. If unit is 'Months', then number specified in duration is multiplied by 30. If 'Years', it is multiplied by 365.",
				ValidateFunc: validation.StringInSlice([]string{"Days", "Weeks", "Months", "Years"}, false),
			},
			"duration": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Specifies the duration for a backup retention.",
				ValidateFunc: validation.IntAtLeast(1),
			},
			"data_lock_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Specifies WORM retention settings.",
				Elem: &schema.Resource{
					Schema: dataLockConfigSchema(),
				},
			},
		},
	}
}
func cdpRetentionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"unit": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Specifies the Retention Unit of a backup measured in days, months or years. If unit is 'Months', then number specified in duration is multiplied by 30. If 'Years', it is multiplied by 365.",
				ValidateFunc: validation.StringInSlice([]string{"Minutes", "Hours"}, false),
			},
			"duration": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Specifies the duration for a backup retention.",
				ValidateFunc: validation.IntAtLeast(1),
			},
			"data_lock_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Specifies WORM retention settings.",
				Elem: &schema.Resource{
					Schema: dataLockConfigSchema(),
				},
			},
		},
	}
}
func dataLockConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mode": {
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"Compliance", "Administrative",
			}, false),
		},
		"unit": {
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"Days", "Weeks", "Months", "Years",
			}, false),
		},
		"duration": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntAtLeast(1),
		},
		"enable_worm_on_external_target": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func frequencyScheduleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"frequency": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Multiplier for frequency unit (e.g., every N hours).",
				ValidateFunc: validation.IntAtLeast(1),
			},
		},
	}
}

func weekScheduleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"day_of_week": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
					}, false),
				},
				Required:    true,
				Description: "Days of the week to run (e.g., Monday, Tuesday).",
			},
		},
	}
}

func monthScheduleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"day_of_week": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
					}, false),
				},
				Description: "List of days of the week for monthly run.",
			},
			"week_of_month": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Week of month (e.g., First, Second, Last).",
				ValidateFunc: validation.StringInSlice([]string{"First", "Second", "Third", "Fourth", "Last"}, false),
			},
			"day_of_month": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Day of month to run (e.g., 18).",
			},
		},
	}
}

func yearScheduleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"day_of_year": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Day of the year to run (e.g., First, Last).",
				ValidateFunc: validation.StringInSlice([]string{"First", "Last"}, false),
			},
		},
	}
}

func fullScheduleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"unit": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "How often to start new runs (e.g., Days, Weeks, Months).",
				ValidateFunc: validation.StringInSlice([]string{"Days", "Weeks", "Months", "Years", "ProtectOnce"}, false),
			},
			"day_schedule": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Run schedule by day frequency.",
				Elem:        frequencyScheduleSchema(),
			},
			"week_schedule": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Weekly run schedule.",
				Elem:        weekScheduleSchema(),
			},
			"month_schedule": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Monthly run schedule.",
				Elem:        monthScheduleSchema(),
			},
			"year_schedule": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Yearly run schedule.",
				Elem:        yearScheduleSchema(),
			},
		},
	}
}

func resourceCohesityBackupPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.FromErr(err)
	}
	ProtectionPolicyRequest := services.BuildProtectionPolicyRequest(d)
	log.Printf("[INFO] Creating backup policy : %+v, %+v, %+v \n done", ProtectionPolicyRequest, ProtectionPolicyRequest.BackupPolicy.Regular.Incremental.Schedule, *ProtectionPolicyRequest.BackupPolicy.Regular.Incremental.Schedule.Unit)
	// Call the Cohesity API to create the backup policy
	id, err := client.CreateProtectionPolicy(ProtectionPolicyRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the resource ID (policy ID)
	d.SetId(id)

	return resourceCohesityBackupPolicyRead(ctx, d, m)
}

func resourceCohesityBackupPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Initialize the Cohesity client
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.FromErr(err)
	}
	// Get the policy ID from the resource data
	policyID := d.Id()

	// Call the Cohesity API to get the backup policy
	policy, err := client.GetProtectionPolicyByID(policyID)
	if err != nil {
		return diag.FromErr(err)
	}
	// log.Printf("[INFO] Getting backup policy : %+v, %+v, %+v \n done", policy, *policy.ProtectionPolicy.Version, *policy.BackupPolicy.Regular)

	// policy := policyResp.ProtectionPolicy
	d.Set("name", policy.Name)
	d.Set("description", policy.Description)
	d.Set("data_lock", policy.DataLock)
	d.Set("version", policy.Version)
	d.Set("is_cbs_enabled", policy.IsCBSEnabled)
	d.Set("last_modification_time_usecs", policy.LastModificationTimeUsecs)
	d.Set("skip_interval_mins", policy.SkipIntervalMins)
	d.Set("enable_smart_local_retention_adjustment", policy.EnableSmartLocalRetentionAdjustment)

	if policy.RetryOptions != nil {
		d.Set("retry_options", []interface{}{
			map[string]interface{}{
				"retries":             policy.RetryOptions.Retries,
				"retry_interval_mins": policy.RetryOptions.RetryIntervalMins,
			},
		})
	} else {
		d.Set("retry_options", nil)
	}

	if policy.BackupPolicy != nil {
		backupPolicy := map[string]interface{}{}

		if policy.BackupPolicy.Regular != nil {
			regular := map[string]interface{}{}

			if policy.BackupPolicy.Regular.Incremental != nil && policy.BackupPolicy.Regular.Incremental.Schedule != nil {
				regular["incremental"] = services.FlattenIncrementalSchedule(policy.BackupPolicy.Regular.Incremental.Schedule)
			}
			if policy.BackupPolicy.Regular.Full != nil && policy.BackupPolicy.Regular.Full.Schedule != nil {
				regular["full"] = services.FlattenFullSchedule(policy.BackupPolicy.Regular.Full.Schedule)
			}
			if len(policy.BackupPolicy.Regular.FullBackups) > 0 {
				fullBackups := make([]interface{}, 0, len(policy.BackupPolicy.Regular.FullBackups))
				for _, fbObj := range policy.BackupPolicy.Regular.FullBackups {
					fb := map[string]interface{}{}
					if fbObj.Schedule != nil {
						fb["schedule"] = services.FlattenFullSchedule(fbObj.Schedule)
					}
					if fbObj.Retention != nil {
						fb["retention"] = services.FlattenRetention(fbObj.Retention)
					}
					fullBackups = append(fullBackups, fbObj)
				}
				backupPolicy["run_timeouts"] = fullBackups
			}
			if policy.BackupPolicy.Regular.Retention != nil {
				regular["retention"] = services.FlattenRetention(policy.BackupPolicy.Regular.Retention)
			}
			backupPolicy["regular"] = []interface{}{regular}
		}

		if policy.BackupPolicy.Log != nil {
			log := map[string]interface{}{}
			if policy.BackupPolicy.Log.Schedule != nil {
				log["schedule"] = services.FlattenLogSchedule(policy.BackupPolicy.Log.Schedule)
			}
			if policy.BackupPolicy.Log.Retention != nil {
				log["retention"] = services.FlattenRetention(policy.BackupPolicy.Log.Retention)
			}
			backupPolicy["log"] = []interface{}{log}
		}

		if policy.BackupPolicy.Bmr != nil {
			bmr := map[string]interface{}{}
			if policy.BackupPolicy.Bmr.Schedule != nil {
				bmr["schedule"] = services.FlattenBmrSchedule(policy.BackupPolicy.Bmr.Schedule)
			}
			if policy.BackupPolicy.Bmr.Retention != nil {
				bmr["retention"] = services.FlattenRetention(policy.BackupPolicy.Bmr.Retention)
			}
			backupPolicy["bmr"] = []interface{}{bmr}
		}

		if policy.BackupPolicy.Cdp != nil {
			cdp := map[string]interface{}{}
			if policy.BackupPolicy.Cdp.Retention != nil {
				cdp["retention"] = services.FlattenCdpRetention(policy.BackupPolicy.Cdp.Retention)
			}
			backupPolicy["cdp"] = []interface{}{cdp}
		}

		if policy.BackupPolicy.StorageArraySnapshot != nil {
			storage := map[string]interface{}{}
			if policy.BackupPolicy.StorageArraySnapshot.Schedule != nil {
				storage["schedule"] = services.FlattenStorageArraySnapshotSchedule(policy.BackupPolicy.StorageArraySnapshot.Schedule)
			}
			if policy.BackupPolicy.StorageArraySnapshot.Retention != nil {
				storage["retention"] = services.FlattenRetention(policy.BackupPolicy.StorageArraySnapshot.Retention)
			}
			backupPolicy["storage_array_snapshot"] = []interface{}{storage}
		}

		if len(policy.BackupPolicy.RunTimeouts) > 0 {
			runTimeouts := make([]interface{}, 0, len(policy.BackupPolicy.RunTimeouts))
			for _, rt := range policy.BackupPolicy.RunTimeouts {
				runTimeouts = append(runTimeouts, map[string]interface{}{
					"timeout_minutes": rt.TimeoutMins,
					"backup_type":     rt.BackupType,
				})
			}
			backupPolicy["run_timeouts"] = runTimeouts
		}

		d.Set("backup_policy", []interface{}{backupPolicy})
	} else {
		d.Set("backup_policy", nil)
	}

	return nil
}

func resourceCohesityBackupPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Initialize the Cohesity client
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.FromErr(err)
	}
	// Get the policy ID from the resource data
	policyID := d.Id()

	policy := services.BuildProtectionPolicyRequest(d)
	log.Printf("[INFO] Updating backup policy : %+v, %+v, %+v \n done", policy, *policy.ProtectionPolicy.Version, *policy.BackupPolicy.Regular)

	// Call the Cohesity API to update the backup policy
	err = client.UpdateProtectionPolicy(policyID, policy)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceCohesityBackupPolicyRead(ctx, d, m)
}

func resourceCohesityBackupPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Initialize the Cohesity client
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.FromErr(err)
	}
	// Get the policy ID from the resource data
	policyID := d.Id()

	// Call the Cohesity API to delete the backup policy
	err = client.DeleteProtectionPolicy(policyID)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
