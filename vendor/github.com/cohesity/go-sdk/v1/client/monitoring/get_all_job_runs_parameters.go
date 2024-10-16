// Code generated by go-swagger; DO NOT EDIT.

package monitoring

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

// NewGetAllJobRunsParams creates a new GetAllJobRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllJobRunsParams() *GetAllJobRunsParams {
	return &GetAllJobRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllJobRunsParamsWithTimeout creates a new GetAllJobRunsParams object
// with the ability to set a timeout on a request.
func NewGetAllJobRunsParamsWithTimeout(timeout time.Duration) *GetAllJobRunsParams {
	return &GetAllJobRunsParams{
		timeout: timeout,
	}
}

// NewGetAllJobRunsParamsWithContext creates a new GetAllJobRunsParams object
// with the ability to set a context for a request.
func NewGetAllJobRunsParamsWithContext(ctx context.Context) *GetAllJobRunsParams {
	return &GetAllJobRunsParams{
		Context: ctx,
	}
}

// NewGetAllJobRunsParamsWithHTTPClient creates a new GetAllJobRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllJobRunsParamsWithHTTPClient(client *http.Client) *GetAllJobRunsParams {
	return &GetAllJobRunsParams{
		HTTPClient: client,
	}
}

/*
GetAllJobRunsParams contains all the parameters to send to the API endpoint

	for the get all job runs operation.

	Typically these are written to a http.Request.
*/
type GetAllJobRunsParams struct {

	/* EndTimeMsecs.

	   Specifies the end time of the time range of interest.

	   Format: int64
	*/
	EndTimeMsecs int64

	/* EnvTypes.

	     Specifies the environment types of the job.
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
	EnvTypes []string

	/* JobTypes.

	   Specifies the job types for which runs are needed.
	*/
	JobTypes []string

	/* Page.

	   Specifies the page number in case of pagination of response.

	   Format: int32
	*/
	Page *int32

	/* PageSize.

	   Specifies the size of the page in case of pagination of response.

	   Format: int32
	*/
	PageSize *int32

	/* StartTimeMsecs.

	   Specifies the start time of the time range of interest.

	   Format: int64
	*/
	StartTimeMsecs int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all job runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllJobRunsParams) WithDefaults() *GetAllJobRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all job runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllJobRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all job runs params
func (o *GetAllJobRunsParams) WithTimeout(timeout time.Duration) *GetAllJobRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all job runs params
func (o *GetAllJobRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all job runs params
func (o *GetAllJobRunsParams) WithContext(ctx context.Context) *GetAllJobRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all job runs params
func (o *GetAllJobRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all job runs params
func (o *GetAllJobRunsParams) WithHTTPClient(client *http.Client) *GetAllJobRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all job runs params
func (o *GetAllJobRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTimeMsecs adds the endTimeMsecs to the get all job runs params
func (o *GetAllJobRunsParams) WithEndTimeMsecs(endTimeMsecs int64) *GetAllJobRunsParams {
	o.SetEndTimeMsecs(endTimeMsecs)
	return o
}

// SetEndTimeMsecs adds the endTimeMsecs to the get all job runs params
func (o *GetAllJobRunsParams) SetEndTimeMsecs(endTimeMsecs int64) {
	o.EndTimeMsecs = endTimeMsecs
}

// WithEnvTypes adds the envTypes to the get all job runs params
func (o *GetAllJobRunsParams) WithEnvTypes(envTypes []string) *GetAllJobRunsParams {
	o.SetEnvTypes(envTypes)
	return o
}

// SetEnvTypes adds the envTypes to the get all job runs params
func (o *GetAllJobRunsParams) SetEnvTypes(envTypes []string) {
	o.EnvTypes = envTypes
}

// WithJobTypes adds the jobTypes to the get all job runs params
func (o *GetAllJobRunsParams) WithJobTypes(jobTypes []string) *GetAllJobRunsParams {
	o.SetJobTypes(jobTypes)
	return o
}

// SetJobTypes adds the jobTypes to the get all job runs params
func (o *GetAllJobRunsParams) SetJobTypes(jobTypes []string) {
	o.JobTypes = jobTypes
}

// WithPage adds the page to the get all job runs params
func (o *GetAllJobRunsParams) WithPage(page *int32) *GetAllJobRunsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get all job runs params
func (o *GetAllJobRunsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get all job runs params
func (o *GetAllJobRunsParams) WithPageSize(pageSize *int32) *GetAllJobRunsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get all job runs params
func (o *GetAllJobRunsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithStartTimeMsecs adds the startTimeMsecs to the get all job runs params
func (o *GetAllJobRunsParams) WithStartTimeMsecs(startTimeMsecs int64) *GetAllJobRunsParams {
	o.SetStartTimeMsecs(startTimeMsecs)
	return o
}

// SetStartTimeMsecs adds the startTimeMsecs to the get all job runs params
func (o *GetAllJobRunsParams) SetStartTimeMsecs(startTimeMsecs int64) {
	o.StartTimeMsecs = startTimeMsecs
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllJobRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endTimeMsecs
	qrEndTimeMsecs := o.EndTimeMsecs
	qEndTimeMsecs := swag.FormatInt64(qrEndTimeMsecs)
	if qEndTimeMsecs != "" {

		if err := r.SetQueryParam("endTimeMsecs", qEndTimeMsecs); err != nil {
			return err
		}
	}

	if o.EnvTypes != nil {

		// binding items for envTypes
		joinedEnvTypes := o.bindParamEnvTypes(reg)

		// query array param envTypes
		if err := r.SetQueryParam("envTypes", joinedEnvTypes...); err != nil {
			return err
		}
	}

	if o.JobTypes != nil {

		// binding items for jobTypes
		joinedJobTypes := o.bindParamJobTypes(reg)

		// query array param jobTypes
		if err := r.SetQueryParam("jobTypes", joinedJobTypes...); err != nil {
			return err
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	// query param startTimeMsecs
	qrStartTimeMsecs := o.StartTimeMsecs
	qStartTimeMsecs := swag.FormatInt64(qrStartTimeMsecs)
	if qStartTimeMsecs != "" {

		if err := r.SetQueryParam("startTimeMsecs", qStartTimeMsecs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAllJobRuns binds the parameter envTypes
func (o *GetAllJobRunsParams) bindParamEnvTypes(formats strfmt.Registry) []string {
	envTypesIR := o.EnvTypes

	var envTypesIC []string
	for _, envTypesIIR := range envTypesIR { // explode []string

		envTypesIIV := envTypesIIR // string as string
		envTypesIC = append(envTypesIC, envTypesIIV)
	}

	// items.CollectionFormat: ""
	envTypesIS := swag.JoinByFormat(envTypesIC, "")

	return envTypesIS
}

// bindParamGetAllJobRuns binds the parameter jobTypes
func (o *GetAllJobRunsParams) bindParamJobTypes(formats strfmt.Registry) []string {
	jobTypesIR := o.JobTypes

	var jobTypesIC []string
	for _, jobTypesIIR := range jobTypesIR { // explode []string

		jobTypesIIV := jobTypesIIR // string as string
		jobTypesIC = append(jobTypesIC, jobTypesIIV)
	}

	// items.CollectionFormat: ""
	jobTypesIS := swag.JoinByFormat(jobTypesIC, "")

	return jobTypesIS
}
