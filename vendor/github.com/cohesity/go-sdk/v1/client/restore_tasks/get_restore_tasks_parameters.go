// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRestoreTasksParams creates a new GetRestoreTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRestoreTasksParams() *GetRestoreTasksParams {
	return &GetRestoreTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestoreTasksParamsWithTimeout creates a new GetRestoreTasksParams object
// with the ability to set a timeout on a request.
func NewGetRestoreTasksParamsWithTimeout(timeout time.Duration) *GetRestoreTasksParams {
	return &GetRestoreTasksParams{
		timeout: timeout,
	}
}

// NewGetRestoreTasksParamsWithContext creates a new GetRestoreTasksParams object
// with the ability to set a context for a request.
func NewGetRestoreTasksParamsWithContext(ctx context.Context) *GetRestoreTasksParams {
	return &GetRestoreTasksParams{
		Context: ctx,
	}
}

// NewGetRestoreTasksParamsWithHTTPClient creates a new GetRestoreTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRestoreTasksParamsWithHTTPClient(client *http.Client) *GetRestoreTasksParams {
	return &GetRestoreTasksParams{
		HTTPClient: client,
	}
}

/*
GetRestoreTasksParams contains all the parameters to send to the API endpoint

	for the get restore tasks operation.

	Typically these are written to a http.Request.
*/
type GetRestoreTasksParams struct {

	/* EndTimeUsecs.

	     Filter by an end time specified as a Unix epoch Timestamp
	(in microseconds). All Restore Tasks (both completed and running)
	on the Cohesity Cluster that started after the specified start time
	but before the specified end time are returned.
	If not set, the end time is the current time.

	     Format: int64
	*/
	EndTimeUsecs *int64

	/* Environment.

	     Specifies the environment like VMware, SQL, where the
	Protection Source exists.
	Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
	NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.
	'kVMware' indicates the VMware Protection Source environment.
	'kHyperV' indicates the HyperV Protection Source environment.
	'kSQL' indicates the SQL Protection Source environment.
	'kView' indicates the View Protection Source environment.
	'kPuppeteer' indicates the Cohesity's Remote Adapter.
	'kPhysical' indicates the physical Protection Source environment.
	'kPure' indicates the Pure Storage Protection Source environment.
	'kNimble' indicates the Nimble Storage Protection Source environment.
	'kAzure' indicates the Microsoft's Azure Protection Source environment.
	'kNetapp' indicates the Netapp Protection Source environment.
	'kAgent' indicates the Agent Protection Source environment.
	'kGenericNas' indicates the Generic Network Attached Storage Protection
	Source environment.
	'kAcropolis' indicates the Acropolis Protection Source environment.
	'kPhysicalFiles' indicates the Physical Files Protection Source environment.
	'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment.
	'kIsilon' indicates the Dell EMC's Isilon Protection Source environment.
	'kGPFS' indicates IBM's GPFS Protection Source environment.
	'kKVM' indicates the KVM Protection Source environment.
	'kAWS' indicates the AWS Protection Source environment.
	'kExchange' indicates the Exchange Protection Source environment.
	'kHyperVVSS' indicates the HyperV VSS Protection Source
	environment.
	'kOracle' indicates the Oracle Protection Source environment.
	'kGCP' indicates the Google Cloud Platform Protection Source environment.
	'kFlashBlade' indicates the Flash Blade Protection Source environment.
	'kAWSNative' indicates the AWS Native Protection Source environment.
	'kO365' indicates the Office 365 Protection Source environment.
	'kO365Outlook' indicates Office 365 outlook Protection Source environment.
	'kHyperFlex' indicates the Hyper Flex Protection Source environment.
	'kGCPNative' indicates the GCP Native Protection Source environment.
	'kAzureNative' indicates the Azure Native Protection Source environment.
	'kKubernetes' indicates a Kubernetes Protection Source environment.
	'kElastifile' indicates Elastifile Protection Source environment.
	'kAD' indicates Active Directory Protection Source environment.
	'kRDSSnapshotManager' indicates AWS RDS Protection Source environment.
	'kCassandra' indicates Cassandra Protection Source environment.
	'kMongoDB' indicates MongoDB Protection Source environment.
	'kCouchbase' indicates Couchbase Protection Source environment.
	'kHdfs' indicates Hdfs Protection Source environment.
	'kHive' indicates Hive Protection Source environment.
	'kHBase' indicates HBase Protection Source environment.
	'kUDA' indicates Universal Data Adapter Protection Source environment.
	'kSAPHANA' indicates SAP HANA protection source environment.
	'kO365Teams' indicates the Office365 Teams Protection Source environment.
	'kO365Group' indicates the Office365 Groups Protection Source environment.
	'kO365Exchange' indicates the Office365 Mailbox Protection Source environment.
	'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment.
	'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment.
	'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment.
	kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis,
	kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle,
	kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative,
	kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager,
	kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA,
	kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint,
	kO365PublicFolders
	*/
	Environment *string

	/* PageCount.

	     Specifies the number of completed Restore Tasks to return in the response
	for pagination purposes. Running Restore Tasks are always returned.
	The newest completed Restore Tasks are returned.

	     Format: int64
	*/
	PageCount *int64

	/* StartTimeUsecs.

	     Filter by a start time specified as a Unix epoch Timestamp
	(in microseconds). All Restore Tasks (both completed and running)
	on the Cohesity Cluster that started after the specified start time
	but before the specified end time are returned.
	If not set, the start time is creation time of the Cohesity Cluster.

	     Format: int64
	*/
	StartTimeUsecs *int64

	/* StorageDomainIds.

	     Filter by a list of Storage Domain IDs. This field applies only if
	'TaskTypes' includes 'kCloneView'. Only task writing data to these storage
	domains will be returned.
	*/
	StorageDomainIds []int64

	/* TaskIds.

	   Filter by a list of Restore Task ids.
	*/
	TaskIds []int64

	/* TaskTypes.

	     Filter by the types of Restore Tasks such as 'kRecoverVMs',
	'kCloneVMs', 'kCloneView' or 'kMountVolumes'.
	*/
	TaskTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get restore tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRestoreTasksParams) WithDefaults() *GetRestoreTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get restore tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRestoreTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get restore tasks params
func (o *GetRestoreTasksParams) WithTimeout(timeout time.Duration) *GetRestoreTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restore tasks params
func (o *GetRestoreTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restore tasks params
func (o *GetRestoreTasksParams) WithContext(ctx context.Context) *GetRestoreTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restore tasks params
func (o *GetRestoreTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restore tasks params
func (o *GetRestoreTasksParams) WithHTTPClient(client *http.Client) *GetRestoreTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restore tasks params
func (o *GetRestoreTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTimeUsecs adds the endTimeUsecs to the get restore tasks params
func (o *GetRestoreTasksParams) WithEndTimeUsecs(endTimeUsecs *int64) *GetRestoreTasksParams {
	o.SetEndTimeUsecs(endTimeUsecs)
	return o
}

// SetEndTimeUsecs adds the endTimeUsecs to the get restore tasks params
func (o *GetRestoreTasksParams) SetEndTimeUsecs(endTimeUsecs *int64) {
	o.EndTimeUsecs = endTimeUsecs
}

// WithEnvironment adds the environment to the get restore tasks params
func (o *GetRestoreTasksParams) WithEnvironment(environment *string) *GetRestoreTasksParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the get restore tasks params
func (o *GetRestoreTasksParams) SetEnvironment(environment *string) {
	o.Environment = environment
}

// WithPageCount adds the pageCount to the get restore tasks params
func (o *GetRestoreTasksParams) WithPageCount(pageCount *int64) *GetRestoreTasksParams {
	o.SetPageCount(pageCount)
	return o
}

// SetPageCount adds the pageCount to the get restore tasks params
func (o *GetRestoreTasksParams) SetPageCount(pageCount *int64) {
	o.PageCount = pageCount
}

// WithStartTimeUsecs adds the startTimeUsecs to the get restore tasks params
func (o *GetRestoreTasksParams) WithStartTimeUsecs(startTimeUsecs *int64) *GetRestoreTasksParams {
	o.SetStartTimeUsecs(startTimeUsecs)
	return o
}

// SetStartTimeUsecs adds the startTimeUsecs to the get restore tasks params
func (o *GetRestoreTasksParams) SetStartTimeUsecs(startTimeUsecs *int64) {
	o.StartTimeUsecs = startTimeUsecs
}

// WithStorageDomainIds adds the storageDomainIds to the get restore tasks params
func (o *GetRestoreTasksParams) WithStorageDomainIds(storageDomainIds []int64) *GetRestoreTasksParams {
	o.SetStorageDomainIds(storageDomainIds)
	return o
}

// SetStorageDomainIds adds the storageDomainIds to the get restore tasks params
func (o *GetRestoreTasksParams) SetStorageDomainIds(storageDomainIds []int64) {
	o.StorageDomainIds = storageDomainIds
}

// WithTaskIds adds the taskIds to the get restore tasks params
func (o *GetRestoreTasksParams) WithTaskIds(taskIds []int64) *GetRestoreTasksParams {
	o.SetTaskIds(taskIds)
	return o
}

// SetTaskIds adds the taskIds to the get restore tasks params
func (o *GetRestoreTasksParams) SetTaskIds(taskIds []int64) {
	o.TaskIds = taskIds
}

// WithTaskTypes adds the taskTypes to the get restore tasks params
func (o *GetRestoreTasksParams) WithTaskTypes(taskTypes []string) *GetRestoreTasksParams {
	o.SetTaskTypes(taskTypes)
	return o
}

// SetTaskTypes adds the taskTypes to the get restore tasks params
func (o *GetRestoreTasksParams) SetTaskTypes(taskTypes []string) {
	o.TaskTypes = taskTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestoreTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndTimeUsecs != nil {

		// query param endTimeUsecs
		var qrEndTimeUsecs int64

		if o.EndTimeUsecs != nil {
			qrEndTimeUsecs = *o.EndTimeUsecs
		}
		qEndTimeUsecs := swag.FormatInt64(qrEndTimeUsecs)
		if qEndTimeUsecs != "" {

			if err := r.SetQueryParam("endTimeUsecs", qEndTimeUsecs); err != nil {
				return err
			}
		}
	}

	if o.Environment != nil {

		// query param environment
		var qrEnvironment string

		if o.Environment != nil {
			qrEnvironment = *o.Environment
		}
		qEnvironment := qrEnvironment
		if qEnvironment != "" {

			if err := r.SetQueryParam("environment", qEnvironment); err != nil {
				return err
			}
		}
	}

	if o.PageCount != nil {

		// query param pageCount
		var qrPageCount int64

		if o.PageCount != nil {
			qrPageCount = *o.PageCount
		}
		qPageCount := swag.FormatInt64(qrPageCount)
		if qPageCount != "" {

			if err := r.SetQueryParam("pageCount", qPageCount); err != nil {
				return err
			}
		}
	}

	if o.StartTimeUsecs != nil {

		// query param startTimeUsecs
		var qrStartTimeUsecs int64

		if o.StartTimeUsecs != nil {
			qrStartTimeUsecs = *o.StartTimeUsecs
		}
		qStartTimeUsecs := swag.FormatInt64(qrStartTimeUsecs)
		if qStartTimeUsecs != "" {

			if err := r.SetQueryParam("startTimeUsecs", qStartTimeUsecs); err != nil {
				return err
			}
		}
	}

	if o.StorageDomainIds != nil {

		// binding items for storageDomainIds
		joinedStorageDomainIds := o.bindParamStorageDomainIds(reg)

		// query array param storageDomainIds
		if err := r.SetQueryParam("storageDomainIds", joinedStorageDomainIds...); err != nil {
			return err
		}
	}

	if o.TaskIds != nil {

		// binding items for taskIds
		joinedTaskIds := o.bindParamTaskIds(reg)

		// query array param taskIds
		if err := r.SetQueryParam("taskIds", joinedTaskIds...); err != nil {
			return err
		}
	}

	if o.TaskTypes != nil {

		// binding items for taskTypes
		joinedTaskTypes := o.bindParamTaskTypes(reg)

		// query array param taskTypes
		if err := r.SetQueryParam("taskTypes", joinedTaskTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetRestoreTasks binds the parameter storageDomainIds
func (o *GetRestoreTasksParams) bindParamStorageDomainIds(formats strfmt.Registry) []string {
	storageDomainIdsIR := o.StorageDomainIds

	var storageDomainIdsIC []string
	for _, storageDomainIdsIIR := range storageDomainIdsIR { // explode []int64

		storageDomainIdsIIV := swag.FormatInt64(storageDomainIdsIIR) // int64 as string
		storageDomainIdsIC = append(storageDomainIdsIC, storageDomainIdsIIV)
	}

	// items.CollectionFormat: ""
	storageDomainIdsIS := swag.JoinByFormat(storageDomainIdsIC, "")

	return storageDomainIdsIS
}

// bindParamGetRestoreTasks binds the parameter taskIds
func (o *GetRestoreTasksParams) bindParamTaskIds(formats strfmt.Registry) []string {
	taskIdsIR := o.TaskIds

	var taskIdsIC []string
	for _, taskIdsIIR := range taskIdsIR { // explode []int64

		taskIdsIIV := swag.FormatInt64(taskIdsIIR) // int64 as string
		taskIdsIC = append(taskIdsIC, taskIdsIIV)
	}

	// items.CollectionFormat: ""
	taskIdsIS := swag.JoinByFormat(taskIdsIC, "")

	return taskIdsIS
}

// bindParamGetRestoreTasks binds the parameter taskTypes
func (o *GetRestoreTasksParams) bindParamTaskTypes(formats strfmt.Registry) []string {
	taskTypesIR := o.TaskTypes

	var taskTypesIC []string
	for _, taskTypesIIR := range taskTypesIR { // explode []string

		taskTypesIIV := taskTypesIIR // string as string
		taskTypesIC = append(taskTypesIC, taskTypesIIV)
	}

	// items.CollectionFormat: ""
	taskTypesIS := swag.JoinByFormat(taskTypesIC, "")

	return taskTypesIS
}
