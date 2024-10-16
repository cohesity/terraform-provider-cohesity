// Code generated by go-swagger; DO NOT EDIT.

package protection_sources

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

// NewListProtectedObjectsParams creates a new ListProtectedObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProtectedObjectsParams() *ListProtectedObjectsParams {
	return &ListProtectedObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProtectedObjectsParamsWithTimeout creates a new ListProtectedObjectsParams object
// with the ability to set a timeout on a request.
func NewListProtectedObjectsParamsWithTimeout(timeout time.Duration) *ListProtectedObjectsParams {
	return &ListProtectedObjectsParams{
		timeout: timeout,
	}
}

// NewListProtectedObjectsParamsWithContext creates a new ListProtectedObjectsParams object
// with the ability to set a context for a request.
func NewListProtectedObjectsParamsWithContext(ctx context.Context) *ListProtectedObjectsParams {
	return &ListProtectedObjectsParams{
		Context: ctx,
	}
}

// NewListProtectedObjectsParamsWithHTTPClient creates a new ListProtectedObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProtectedObjectsParamsWithHTTPClient(client *http.Client) *ListProtectedObjectsParams {
	return &ListProtectedObjectsParams{
		HTTPClient: client,
	}
}

/*
ListProtectedObjectsParams contains all the parameters to send to the API endpoint

	for the list protected objects operation.

	Typically these are written to a http.Request.
*/
type ListProtectedObjectsParams struct {

	/* AllUnderHierarchy.

	     AllUnderHierarchy specifies if objects of all the tenants under the
	hierarchy of the logged in user's organization should be returned.
	*/
	AllUnderHierarchy *bool

	/* Environment.

	     Specifies the environment type of the registered Protection Source
	such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure',
	'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'.
	For example, set this parameter to 'kVMware' if the registered
	Protection Source is of 'kVMware' environment type.
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
	Environment string

	/* ID.

	     Specifies the Id of a registered Protection Source of the type
	given in environment.

	     Format: int64
	*/
	ID int64

	/* IncludeRpoSnapshots.

	     If true, then the Protected Objects protected by RPO policies will also
	be returned.
	*/
	IncludeRpoSnapshots *bool

	/* PruneProtectionJobMetadata.

	     Specifies whether the metadata about the protection job is to be pruned.
	If set to true, only ID, name & policy info will be returned.
	Info about indexing, entities within the job and other additional settings
	will be omitted.
	*/
	PruneProtectionJobMetadata *bool

	/* RequestInitiatorType.

	     Specifies the type of the request. Possible values are UIUser and UIAuto,
	which means the request is triggered by user or is an auto refresh
	request. Services like magneto will use this to determine the priority
	of the requests, so that it can more intelligently handle overload
	situations by prioritizing higher priority requests.
	*/
	RequestInitiatorType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list protected objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProtectedObjectsParams) WithDefaults() *ListProtectedObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list protected objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProtectedObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list protected objects params
func (o *ListProtectedObjectsParams) WithTimeout(timeout time.Duration) *ListProtectedObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list protected objects params
func (o *ListProtectedObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list protected objects params
func (o *ListProtectedObjectsParams) WithContext(ctx context.Context) *ListProtectedObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list protected objects params
func (o *ListProtectedObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list protected objects params
func (o *ListProtectedObjectsParams) WithHTTPClient(client *http.Client) *ListProtectedObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list protected objects params
func (o *ListProtectedObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllUnderHierarchy adds the allUnderHierarchy to the list protected objects params
func (o *ListProtectedObjectsParams) WithAllUnderHierarchy(allUnderHierarchy *bool) *ListProtectedObjectsParams {
	o.SetAllUnderHierarchy(allUnderHierarchy)
	return o
}

// SetAllUnderHierarchy adds the allUnderHierarchy to the list protected objects params
func (o *ListProtectedObjectsParams) SetAllUnderHierarchy(allUnderHierarchy *bool) {
	o.AllUnderHierarchy = allUnderHierarchy
}

// WithEnvironment adds the environment to the list protected objects params
func (o *ListProtectedObjectsParams) WithEnvironment(environment string) *ListProtectedObjectsParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the list protected objects params
func (o *ListProtectedObjectsParams) SetEnvironment(environment string) {
	o.Environment = environment
}

// WithID adds the id to the list protected objects params
func (o *ListProtectedObjectsParams) WithID(id int64) *ListProtectedObjectsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list protected objects params
func (o *ListProtectedObjectsParams) SetID(id int64) {
	o.ID = id
}

// WithIncludeRpoSnapshots adds the includeRpoSnapshots to the list protected objects params
func (o *ListProtectedObjectsParams) WithIncludeRpoSnapshots(includeRpoSnapshots *bool) *ListProtectedObjectsParams {
	o.SetIncludeRpoSnapshots(includeRpoSnapshots)
	return o
}

// SetIncludeRpoSnapshots adds the includeRpoSnapshots to the list protected objects params
func (o *ListProtectedObjectsParams) SetIncludeRpoSnapshots(includeRpoSnapshots *bool) {
	o.IncludeRpoSnapshots = includeRpoSnapshots
}

// WithPruneProtectionJobMetadata adds the pruneProtectionJobMetadata to the list protected objects params
func (o *ListProtectedObjectsParams) WithPruneProtectionJobMetadata(pruneProtectionJobMetadata *bool) *ListProtectedObjectsParams {
	o.SetPruneProtectionJobMetadata(pruneProtectionJobMetadata)
	return o
}

// SetPruneProtectionJobMetadata adds the pruneProtectionJobMetadata to the list protected objects params
func (o *ListProtectedObjectsParams) SetPruneProtectionJobMetadata(pruneProtectionJobMetadata *bool) {
	o.PruneProtectionJobMetadata = pruneProtectionJobMetadata
}

// WithRequestInitiatorType adds the requestInitiatorType to the list protected objects params
func (o *ListProtectedObjectsParams) WithRequestInitiatorType(requestInitiatorType *string) *ListProtectedObjectsParams {
	o.SetRequestInitiatorType(requestInitiatorType)
	return o
}

// SetRequestInitiatorType adds the requestInitiatorType to the list protected objects params
func (o *ListProtectedObjectsParams) SetRequestInitiatorType(requestInitiatorType *string) {
	o.RequestInitiatorType = requestInitiatorType
}

// WriteToRequest writes these params to a swagger request
func (o *ListProtectedObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllUnderHierarchy != nil {

		// query param allUnderHierarchy
		var qrAllUnderHierarchy bool

		if o.AllUnderHierarchy != nil {
			qrAllUnderHierarchy = *o.AllUnderHierarchy
		}
		qAllUnderHierarchy := swag.FormatBool(qrAllUnderHierarchy)
		if qAllUnderHierarchy != "" {

			if err := r.SetQueryParam("allUnderHierarchy", qAllUnderHierarchy); err != nil {
				return err
			}
		}
	}

	// query param environment
	qrEnvironment := o.Environment
	qEnvironment := qrEnvironment
	if qEnvironment != "" {

		if err := r.SetQueryParam("environment", qEnvironment); err != nil {
			return err
		}
	}

	// query param id
	qrID := o.ID
	qID := swag.FormatInt64(qrID)
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if o.IncludeRpoSnapshots != nil {

		// query param includeRpoSnapshots
		var qrIncludeRpoSnapshots bool

		if o.IncludeRpoSnapshots != nil {
			qrIncludeRpoSnapshots = *o.IncludeRpoSnapshots
		}
		qIncludeRpoSnapshots := swag.FormatBool(qrIncludeRpoSnapshots)
		if qIncludeRpoSnapshots != "" {

			if err := r.SetQueryParam("includeRpoSnapshots", qIncludeRpoSnapshots); err != nil {
				return err
			}
		}
	}

	if o.PruneProtectionJobMetadata != nil {

		// query param pruneProtectionJobMetadata
		var qrPruneProtectionJobMetadata bool

		if o.PruneProtectionJobMetadata != nil {
			qrPruneProtectionJobMetadata = *o.PruneProtectionJobMetadata
		}
		qPruneProtectionJobMetadata := swag.FormatBool(qrPruneProtectionJobMetadata)
		if qPruneProtectionJobMetadata != "" {

			if err := r.SetQueryParam("pruneProtectionJobMetadata", qPruneProtectionJobMetadata); err != nil {
				return err
			}
		}
	}

	if o.RequestInitiatorType != nil {

		// header param requestInitiatorType
		if err := r.SetHeaderParam("requestInitiatorType", *o.RequestInitiatorType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
