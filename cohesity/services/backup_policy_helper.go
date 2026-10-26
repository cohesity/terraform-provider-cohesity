package services

import (
	"github.com/cohesity/go-sdk/v2/client/policy"
	modelsV2 "github.com/cohesity/go-sdk/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func BuildProtectionPolicyRequest(d *schema.ResourceData) *modelsV2.ProtectionPolicyRequest {
	return &modelsV2.ProtectionPolicyRequest{
		ProtectionPolicy: *BuildProtectionPolicy(d),
	}
}
func BuildProtectionPolicy(d *schema.ResourceData) *modelsV2.ProtectionPolicy {
	pg := &modelsV2.ProtectionPolicy{}

	if v, ok := d.GetOk("name"); ok {
		pg.Name = utils.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("description"); ok {
		pg.Description = utils.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("data_lock"); ok {
		pg.DataLock = utils.StringPtr(v.(string))
	}
	if v, ok := d.GetOk("version"); ok {
		pg.Version = utils.Int32Ptr(int32(v.(int)))
	}
	if v, ok := d.GetOk("is_cbs_enabled"); ok {
		pg.IsCBSEnabled = utils.BoolPtr(v.(bool))
	}
	if v, ok := d.GetOk("last_modification_time_usecs"); ok {
		pg.LastModificationTimeUsecs = utils.Int64Ptr(int64(v.(int)))
	}
	if v, ok := d.GetOk("skip_interval_mins"); ok {
		pg.SkipIntervalMins = utils.Int32Ptr(int32(v.(int)))
	}
	if v, ok := d.GetOk("enable_smart_local_retention_adjustment"); ok {
		pg.EnableSmartLocalRetentionAdjustment = utils.BoolPtr(v.(bool))
	}
	if v, ok := d.GetOk("retry_options"); ok && v != nil {
		pg.RetryOptions = expandRetryOptions(v.([]interface{}))
	}
	if v, ok := d.GetOk("backup_policy"); ok && v != nil {
		pg.BackupPolicy = expandBackupPolicy(v.([]interface{}))
	}

	return pg
}

func expandRetryOptions(l []interface{}) *modelsV2.RetryOptions {
	if len(l) == 0 || l[0] == nil {
		return nil
	}
	data := l[0].(map[string]interface{})
	result := &modelsV2.RetryOptions{}

	if v, ok := data["retries"]; ok {
		result.Retries = utils.Int32Ptr(int32(v.(int)))
	}
	if v, ok := data["retry_interval_mins"]; ok {
		result.RetryIntervalMins = utils.Int32Ptr(int32(v.(int)))
	}

	return result
}

func expandBackupPolicy(l []interface{}) *modelsV2.BackupPolicy {
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	bp, ok := l[0].(map[string]interface{})
	if !ok {
		return nil
	}

	result := &modelsV2.BackupPolicy{}

	if val, ok := bp["regular"]; ok && val != nil {
		if expanded := expandRegularBackup(val); expanded != nil {
			result.Regular = expanded
		}
	}

	if val, ok := bp["log"]; ok && val != nil {
		if expanded := expandLogBackup(val); expanded != nil {
			result.Log = expanded
		}
	}

	if val, ok := bp["bmr"]; ok && val != nil {
		if expanded := expandBMRBackup(val); expanded != nil {
			result.Bmr = expanded
		}
	}

	if val, ok := bp["cdp"]; ok && val != nil {
		if expanded := expandCDPBackup(val); expanded != nil {
			result.Cdp = expanded
		}
	}

	if val, ok := bp["storage_array_snapshot"]; ok && val != nil {
		if expanded := expandStorageArraySnapshotBackup(val); expanded != nil {
			result.StorageArraySnapshot = expanded
		}
	}

	if val, ok := bp["run_timeouts"]; ok && val != nil {
		if expanded := expandRunTimeouts(val); expanded != nil {
			result.RunTimeouts = expanded
		}
	}

	return result
}
func getMapInterface(v interface{}) map[string]interface{} {
	if v == nil {
		return nil
	}

	list, ok := v.([]interface{})
	if !ok || len(list) == 0 || list[0] == nil {
		return nil
	}

	data, ok := list[0].(map[string]interface{})
	if !ok {
		return nil
	}
	return data
}

func expandRegularBackup(v interface{}) *modelsV2.RegularBackupPolicy {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}

	result := &modelsV2.RegularBackupPolicy{}

	// Only populate Incremental if present
	if inc, ok := data["incremental"]; ok && inc != nil {
		if expanded := expandIncrementalSchedule(inc); expanded != nil {
			result.Incremental = &modelsV2.IncrementalBackupPolicy{
				Schedule: expanded,
			}
		}
	}

	// Only populate Full if present
	if full, ok := data["full"]; ok && full != nil {
		if expanded := expandFullSchedule(full); expanded != nil {
			result.Full = &modelsV2.FullBackupPolicy{
				Schedule: expanded,
			}
		}
	}

	// Only populate FullBackups if present
	if fb, ok := data["full_backups"]; ok && fb != nil {
		if expanded := expandFullBackups(fb); expanded != nil {
			result.FullBackups = expanded
		}
	}

	// Only populate Retention if present
	if r, ok := data["retention"]; ok && r != nil {
		if expanded := expandRetention(r); expanded != nil {
			result.Retention = expanded
		}
	}

	return result
}

func expandFullBackups(v interface{}) []*modelsV2.FullScheduleAndRetention {

	if v == nil {
		return nil
	}

	list, ok := v.([]interface{})
	if !ok || len(list) == 0 || list[0] == nil {
		return nil
	}
	result := make([]*modelsV2.FullScheduleAndRetention, 0, len(list))
	for _, item := range list {
		entry := item.(map[string]interface{})
		result = append(result, &modelsV2.FullScheduleAndRetention{
			Schedule:  expandFullSchedule(entry["full_backups"]),
			Retention: expandRetention(entry["full_backups"]),
		})
	}
	return result
}

func expandLogBackup(v interface{}) *modelsV2.LogBackupPolicy {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.LogBackupPolicy{
		Schedule:  expandLogSchedule(data["schedule"]),
		Retention: expandRetention(data["retention"]),
	}
}

func expandBMRBackup(v interface{}) *modelsV2.BmrBackupPolicy {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.BmrBackupPolicy{
		Schedule:  expandBmrSchedule(data["schedule"]),
		Retention: expandRetention(data["retention"]),
	}
}

func expandCDPBackup(v interface{}) *modelsV2.CdpBackupPolicy {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.CdpBackupPolicy{
		Retention: expandCDPRetention(data["retention"]),
	}
}

func expandStorageArraySnapshotBackup(v interface{}) *modelsV2.StorageArraySnapshotBackupPolicy {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.StorageArraySnapshotBackupPolicy{
		Schedule:  expandStorageArraySnapshotSchedule(data["schedule"]),
		Retention: expandRetention(data["retention"]),
	}
}
func expandStorageArraySnapshotSchedule(v interface{}) *modelsV2.StorageArraySnapshotSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}

	// Initialize the StorageArraySnapshotSchedule
	result := &modelsV2.StorageArraySnapshotSchedule{}

	// Handle the 'unit' field (string pointer)
	if unit, ok := data["unit"].(string); ok {
		result.Unit = &unit
	}

	// Handle the minute schedule (if present)
	if minuteSchedule, ok := data["minute_schedule"]; ok {
		result.MinuteSchedule = expandMinuteSchedule(minuteSchedule)
	}

	// Handle the hour schedule (if present)
	if hourSchedule, ok := data["hour_schedule"]; ok {
		result.HourSchedule = expandHourSchedule(hourSchedule)
	}

	// Handle the day schedule (if present)
	if daySchedule, ok := data["day_schedule"]; ok {
		result.DaySchedule = expandDaySchedule(daySchedule)
	}

	// Handle the week schedule (if present)
	if weekSchedule, ok := data["week_schedule"]; ok {
		result.WeekSchedule = expandWeekSchedule(weekSchedule)
	}

	// Handle the month schedule (if present)
	if monthSchedule, ok := data["month_schedule"]; ok {
		result.MonthSchedule = expandMonthSchedule(monthSchedule)
	}

	// Handle the year schedule (if present)
	if yearSchedule, ok := data["year_schedule"]; ok {
		result.YearSchedule = expandYearSchedule(yearSchedule)
	}

	return result
}
func expandIncrementalSchedule(v interface{}) *modelsV2.IncrementalSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}

	// Initialize the StorageArraySnapshotSchedule
	result := &modelsV2.IncrementalSchedule{}

	// Handle the 'unit' field (string pointer)
	if unit, ok := data["unit"].(string); ok {
		result.Unit = &unit
	}

	// Handle the minute schedule (if present)
	if minuteSchedule, ok := data["minute_schedule"]; ok {
		result.MinuteSchedule = expandMinuteSchedule(minuteSchedule)
	}

	// Handle the hour schedule (if present)
	if hourSchedule, ok := data["hour_schedule"]; ok {
		result.HourSchedule = expandHourSchedule(hourSchedule)
	}

	// Handle the day schedule (if present)
	if daySchedule, ok := data["day_schedule"]; ok {
		result.DaySchedule = expandDaySchedule(daySchedule)
	}

	// Handle the week schedule (if present)
	if weekSchedule, ok := data["week_schedule"]; ok {
		result.WeekSchedule = expandWeekSchedule(weekSchedule)
	}

	// Handle the month schedule (if present)
	if monthSchedule, ok := data["month_schedule"]; ok {
		result.MonthSchedule = expandMonthSchedule(monthSchedule)
	}

	// Handle the year schedule (if present)
	if yearSchedule, ok := data["year_schedule"]; ok {
		result.YearSchedule = expandYearSchedule(yearSchedule)
	}

	return result
}
func expandRunTimeouts(v interface{}) []*modelsV2.CancellationTimeoutParams {
	if v == nil {
		return nil
	}
	list := v.([]interface{})
	result := make([]*modelsV2.CancellationTimeoutParams, 0, len(list))
	for _, item := range list {
		entry := item.(map[string]interface{})
		result = append(result, &modelsV2.CancellationTimeoutParams{
			TimeoutMins: utils.Int64Ptr(int64(entry["timeout_minutes"].(int))),
			BackupType:  utils.StringPtr(entry["backup_type"].(string)),
		})
	}
	return result
}
func expandFullSchedule(v interface{}) *modelsV2.FullSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	result := &modelsV2.FullSchedule{}

	// Handle the 'unit' field (string pointer)
	if unit, ok := data["unit"].(string); ok {
		result.Unit = &unit
	}
	// Handle the day schedule (if present)
	if daySchedule, ok := data["day_schedule"]; ok {
		result.DaySchedule = expandDaySchedule(daySchedule)
	}

	// Handle the week schedule (if present)
	if weekSchedule, ok := data["week_schedule"]; ok {
		result.WeekSchedule = expandWeekSchedule(weekSchedule)
	}

	// Handle the month schedule (if present)
	if monthSchedule, ok := data["month_schedule"]; ok {
		result.MonthSchedule = expandMonthSchedule(monthSchedule)
	}

	// Handle the year schedule (if present)
	if yearSchedule, ok := data["year_schedule"]; ok {
		result.YearSchedule = expandYearSchedule(yearSchedule)
	}

	return result

}
func expandBmrSchedule(v interface{}) *modelsV2.BmrSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	result := &modelsV2.BmrSchedule{}

	// Handle the 'unit' field (string pointer)
	if unit, ok := data["unit"].(string); ok {
		result.Unit = &unit
	}
	// Handle the day schedule (if present)
	if daySchedule, ok := data["day_schedule"]; ok {
		result.DaySchedule = expandDaySchedule(daySchedule)
	}

	// Handle the week schedule (if present)
	if weekSchedule, ok := data["week_schedule"]; ok {
		result.WeekSchedule = expandWeekSchedule(weekSchedule)
	}

	// Handle the month schedule (if present)
	if monthSchedule, ok := data["month_schedule"]; ok {
		result.MonthSchedule = expandMonthSchedule(monthSchedule)
	}

	// Handle the year schedule (if present)
	if yearSchedule, ok := data["year_schedule"]; ok {
		result.YearSchedule = expandYearSchedule(yearSchedule)
	}

	return result
}
func expandLogSchedule(v interface{}) *modelsV2.LogSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	// Initialize the StorageArraySnapshotSchedule
	result := &modelsV2.LogSchedule{}

	// Handle the 'unit' field (string pointer)
	if unit, ok := data["unit"].(string); ok {
		result.Unit = &unit
	}

	// Handle the minute schedule (if present)
	if minuteSchedule, ok := data["minute_schedule"]; ok {
		result.MinuteSchedule = expandMinuteSchedule(minuteSchedule)
	}

	// Handle the hour schedule (if present)
	if hourSchedule, ok := data["hour_schedule"]; ok {
		result.HourSchedule = expandHourSchedule(hourSchedule)
	}
	return result
}

func expandMinuteSchedule(v interface{}) *modelsV2.MinuteSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.MinuteSchedule{
		FrequencySchedule: modelsV2.FrequencySchedule{
			Frequency: utils.Int64Ptr(int64(data["frequency"].(int))),
		},
	}
}
func expandDaySchedule(v interface{}) *modelsV2.DaySchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.DaySchedule{
		FrequencySchedule: modelsV2.FrequencySchedule{
			Frequency: utils.Int64Ptr(int64(data["frequency"].(int))),
		},
	}
}
func expandHourSchedule(v interface{}) *modelsV2.HourSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}

	return &modelsV2.HourSchedule{
		FrequencySchedule: modelsV2.FrequencySchedule{
			Frequency: utils.Int64Ptr(int64(data["frequency"].(int))),
		},
	}
}
func expandWeekSchedule(v interface{}) *modelsV2.WeekSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.WeekSchedule{
		DayOfWeek: expandDayOfWeek(data["day_of_week"]),
	}
}

func expandMonthSchedule(v interface{}) *modelsV2.MonthSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.MonthSchedule{
		DayOfWeek:   expandDayOfWeek(data["day_of_week"]),
		WeekOfMonth: utils.StringPtr(data["week_of_month"].(string)),
		DayOfMonth:  utils.Int32Ptr(int32(data["day_of_month"].(int))),
	}
}

func expandYearSchedule(v interface{}) *modelsV2.YearSchedule {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.YearSchedule{
		DayOfYear: utils.StringPtr(data["day_of_year"].(string)),
	}
}

func expandRetention(v interface{}) *modelsV2.Retention {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.Retention{
		Unit:           utils.StringPtr(data["unit"].(string)),
		Duration:       utils.Int64Ptr(int64(data["duration"].(int))),
		DataLockConfig: expandDataLockConfig(data["data_lock_config"]),
	}
}

func expandCDPRetention(v interface{}) *modelsV2.CdpRetention {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.CdpRetention{
		Unit:           utils.StringPtr(data["unit"].(string)),
		Duration:       utils.Int32Ptr(int32(data["duration"].(int))),
		DataLockConfig: expandDataLockConfig(data["data_lock_config"]),
	}
}

func expandDataLockConfig(v interface{}) *modelsV2.DataLockConfig {
	data := getMapInterface(v)
	if data==nil {
		return nil
	}
	return &modelsV2.DataLockConfig{
		Mode:                       utils.StringPtr(data["mode"].(string)),
		Unit:                       utils.StringPtr(data["unit"].(string)),
		Duration:                   utils.Int64Ptr(int64(data["duration"].(int))),
		EnableWormOnExternalTarget: utils.BoolPtr(data["enable_worm_on_external_target"].(bool)),
	}
}

func expandDayOfWeek(v interface{}) []string {
	if v == nil {
		return nil
	}
	days := v.([]interface{})
	result := make([]string, 0, len(days))
	for _, day := range days {
		result = append(result, day.(string))
	}
	return result
}
func FlattenRetention(retention *modelsV2.Retention) []interface{} {
	if retention == nil {
		return nil
	}

	data := make(map[string]interface{})

	if retention.Unit != nil {
		data["unit"] = *retention.Unit
	}
	if retention.Duration != nil {
		data["duration"] = *retention.Duration
	}
	if retention.DataLockConfig != nil {
		dlc := make(map[string]interface{})
		if retention.DataLockConfig.Mode != nil {
			dlc["mode"] = *retention.DataLockConfig.Mode
		}
		if retention.DataLockConfig.Duration != nil {
			dlc["duration"] = *retention.DataLockConfig.Duration
		}
		if retention.DataLockConfig.Unit != nil {
			dlc["unit"] = *retention.DataLockConfig.Unit
		}
		data["data_lock_config"] = []interface{}{dlc}
	}

	return []interface{}{data}
}
func FlattenLogSchedule(schedule *modelsV2.LogSchedule) []interface{} {
	if schedule == nil {
		return nil
	}

	data := make(map[string]interface{})
	if schedule.Unit != nil {
		data["unit"] = schedule.Unit
	}
	if schedule.MinuteSchedule != nil {
		data["minute_schedule"] = flattenMinuteSchedule(schedule.MinuteSchedule)
	}
	if schedule.HourSchedule != nil {
		data["hour_schedule"] = flattenHourSchedule(schedule.HourSchedule)
	}

	return []interface{}{data}
}
func flattenWeekSchedule(w *modelsV2.WeekSchedule) []interface{} {
	if w == nil {
		return nil
	}

	m := map[string]interface{}{
		"day_of_week": w.DayOfWeek,
	}
	return []interface{}{m}
}

func flattenMonthSchedule(mo *modelsV2.MonthSchedule) []interface{} {
	if mo == nil {
		return nil
	}

	data := make(map[string]interface{})
	if mo.DayOfWeek != nil {
		data["day_of_week"] = mo.DayOfWeek
	}
	if mo.WeekOfMonth != nil {
		data["week_of_month"] = *mo.WeekOfMonth
	}
	if mo.DayOfMonth != nil {
		data["day_of_month"] = int(*mo.DayOfMonth)
	}

	return []interface{}{data}
}

func flattenYearSchedule(y *modelsV2.YearSchedule) []interface{} {
	if y == nil {
		return nil
	}

	data := make(map[string]interface{})
	if y.DayOfYear != nil {
		data["day_of_year"] = *y.DayOfYear
	}

	return []interface{}{data}
}
func flattenDaySchedule(d *modelsV2.DaySchedule) []interface{} {
	if d == nil {
		return nil
	}

	data := make(map[string]interface{})
	if d.Frequency != nil {
		data["frequency"] = *d.Frequency
	}

	return []interface{}{data}
}
func flattenHourSchedule(d *modelsV2.HourSchedule) []interface{} {
	if d == nil {
		return nil
	}

	data := make(map[string]interface{})
	if d.Frequency != nil {
		data["frequency"] = *d.Frequency
	}

	return []interface{}{data}
}
func flattenMinuteSchedule(d *modelsV2.MinuteSchedule) []interface{} {
	if d == nil {
		return nil
	}

	data := make(map[string]interface{})
	if d.Frequency != nil {
		data["frequency"] = *d.Frequency
	}

	return []interface{}{data}
}
func FlattenIncrementalSchedule(schedule *modelsV2.IncrementalSchedule) []interface{} {
	if schedule == nil {
		return nil
	}

	data := make(map[string]interface{})
	if schedule.Unit != nil {
		data["unit"] = schedule.Unit
	}
	if schedule.MinuteSchedule != nil {
		data["minute_schedule"] = flattenMinuteSchedule(schedule.MinuteSchedule)
	}
	if schedule.HourSchedule != nil {
		data["hour_schedule"] = flattenHourSchedule(schedule.HourSchedule)
	}
	if schedule.DaySchedule != nil {
		data["day_schedule"] = flattenDaySchedule(schedule.DaySchedule)
	}
	if schedule.WeekSchedule != nil {
		data["week_schedule"] = flattenWeekSchedule(schedule.WeekSchedule)
	}
	if schedule.MonthSchedule != nil {
		data["month_schedule"] = flattenMonthSchedule(schedule.MonthSchedule)
	}
	if schedule.YearSchedule != nil {
		data["year_schedule"] = flattenYearSchedule(schedule.YearSchedule)
	}

	return []interface{}{data}
}
func FlattenStorageArraySnapshotSchedule(schedule *modelsV2.StorageArraySnapshotSchedule) []interface{} {
	if schedule == nil {
		return nil
	}

	data := make(map[string]interface{})
	if schedule.Unit != nil {
		data["unit"] = schedule.Unit
	}
	if schedule.MinuteSchedule != nil {
		data["minute_schedule"] = flattenMinuteSchedule(schedule.MinuteSchedule)
	}
	if schedule.HourSchedule != nil {
		data["hour_schedule"] = flattenHourSchedule(schedule.HourSchedule)
	}
	if schedule.DaySchedule != nil {
		data["day_schedule"] = flattenDaySchedule(schedule.DaySchedule)
	}
	if schedule.WeekSchedule != nil {
		data["week_schedule"] = flattenWeekSchedule(schedule.WeekSchedule)
	}
	if schedule.MonthSchedule != nil {
		data["month_schedule"] = flattenMonthSchedule(schedule.MonthSchedule)
	}
	if schedule.YearSchedule != nil {
		data["year_schedule"] = flattenYearSchedule(schedule.YearSchedule)
	}

	return []interface{}{data}
}
func FlattenFullSchedule(schedule *modelsV2.FullSchedule) []interface{} {
	if schedule == nil {
		return nil
	}

	data := make(map[string]interface{})
	if schedule.Unit != nil {
		data["unit"] = schedule.Unit
	}
	if schedule.DaySchedule != nil {
		data["day_schedule"] = flattenDaySchedule(schedule.DaySchedule)
	}
	if schedule.WeekSchedule != nil {
		data["week_schedule"] = flattenWeekSchedule(schedule.WeekSchedule)
	}
	if schedule.MonthSchedule != nil {
		data["month_schedule"] = flattenMonthSchedule(schedule.MonthSchedule)
	}
	if schedule.YearSchedule != nil {
		data["year_schedule"] = flattenYearSchedule(schedule.YearSchedule)
	}

	return []interface{}{data}
}
func FlattenBmrSchedule(schedule *modelsV2.BmrSchedule) []interface{} {
	if schedule == nil {
		return nil
	}

	data := make(map[string]interface{})
	if schedule.Unit != nil {
		data["unit"] = schedule.Unit
	}
	if schedule.DaySchedule != nil {
		data["day_schedule"] = flattenDaySchedule(schedule.DaySchedule)
	}
	if schedule.WeekSchedule != nil {
		data["week_schedule"] = flattenWeekSchedule(schedule.WeekSchedule)
	}
	if schedule.MonthSchedule != nil {
		data["month_schedule"] = flattenMonthSchedule(schedule.MonthSchedule)
	}
	if schedule.YearSchedule != nil {
		data["year_schedule"] = flattenYearSchedule(schedule.YearSchedule)
	}

	return []interface{}{data}
}

func FlattenCdpRetention(retention *modelsV2.CdpRetention) []interface{} {
	if retention == nil {
		return nil
	}

	data := make(map[string]interface{})

	if retention.Unit != nil {
		data["unit"] = *retention.Unit
	}
	if retention.Duration != nil {
		data["duration"] = *retention.Duration
	}
	if retention.DataLockConfig != nil {
		dlc := make(map[string]interface{})
		if retention.DataLockConfig.Mode != nil {
			dlc["mode"] = *retention.DataLockConfig.Mode
		}
		if retention.DataLockConfig.Duration != nil {
			dlc["duration"] = *retention.DataLockConfig.Duration
		}
		if retention.DataLockConfig.Unit != nil {
			dlc["unit"] = *retention.DataLockConfig.Unit
		}
		data["data_lock_config"] = []interface{}{dlc}
	}

	return []interface{}{data}
}
func (c *CohesityClientV2) CreateProtectionPolicy(createBackupPolicyParams *modelsV2.ProtectionPolicyRequest) (string, error) {

	resp, err := c.client.Policy.CreateProtectionPolicy(policy.NewCreateProtectionPolicyParams().WithBody(createBackupPolicyParams), c.bearerToken)
	if err != nil {
		return "", err
	}
	return *resp.Payload.ID, nil
}

func (c *CohesityClientV2) DeleteProtectionPolicy(policyID string) error {
	_, err := c.client.Policy.DeleteProtectionPolicy(policy.NewDeleteProtectionPolicyParams().WithID(policyID), c.bearerToken)
	return err
}

func (c *CohesityClientV2) GetProtectionPolicyByID(policyID string) (*modelsV2.ProtectionPolicyResponse, error) {
	resp, err := c.client.Policy.GetProtectionPolicyByID(policy.NewGetProtectionPolicyByIDParams().WithID(policyID), c.bearerToken)
	if err != nil {
		return nil, err
	}
	return resp.Payload, err
}

func (c *CohesityClientV2) UpdateProtectionPolicy(policyID string, updatePolicyParams *modelsV2.ProtectionPolicyRequest) error {
	_, err := c.client.Policy.UpdateProtectionPolicy(policy.NewUpdateProtectionPolicyParams().WithID(policyID).WithBody(updatePolicyParams), c.bearerToken)
	if err != nil {
		return err
	}
	return nil
}
