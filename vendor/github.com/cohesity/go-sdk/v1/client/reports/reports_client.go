// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new reports API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new reports API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAgentDeploymentReport(params *GetAgentDeploymentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAgentDeploymentReportOK, error)

	GetCloudArchiveReportRequest(params *GetCloudArchiveReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCloudArchiveReportRequestOK, error)

	GetDataTransferFromVaultsReportRequest(params *GetDataTransferFromVaultsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataTransferFromVaultsReportRequestOK, error)

	GetDataTransferToVaultsReportRequest(params *GetDataTransferToVaultsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataTransferToVaultsReportRequestOK, error)

	GetGdprReport(params *GetGdprReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGdprReportOK, error)

	GetProtectedObjectsTrendsReportRequest(params *GetProtectedObjectsTrendsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectedObjectsTrendsReportRequestOK, error)

	GetProtectionSourcesJobRunsReportRequest(params *GetProtectionSourcesJobRunsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionSourcesJobRunsReportRequestOK, error)

	GetProtectionSourcesJobsSummaryReportRequest(params *GetProtectionSourcesJobsSummaryReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionSourcesJobsSummaryReportRequestOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetAgentDeploymentReport **Privileges:** ```REPORTS_VIEW``` <br><br>Get the list of all the installed agents which includes the health status and

upgradability of the agent.
*/
func (a *Client) GetAgentDeploymentReport(params *GetAgentDeploymentReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAgentDeploymentReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentDeploymentReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAgentDeploymentReport",
		Method:             "GET",
		PathPattern:        "/public/reports/agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAgentDeploymentReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAgentDeploymentReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAgentDeploymentReportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetCloudArchiveReportRequest gets summary statistics about succesful failed job runs queued archival jobs

	**Privileges:** ```REPORTS_VIEW``` <br><br>last run status by each protection job.

Cohesity Cluster to Vaults (External Targets).
*/
func (a *Client) GetCloudArchiveReportRequest(params *GetCloudArchiveReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCloudArchiveReportRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCloudArchiveReportRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCloudArchiveReportRequest",
		Method:             "GET",
		PathPattern:        "/public/reports/cloudArchiveReport",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCloudArchiveReportRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCloudArchiveReportRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCloudArchiveReportRequestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetDataTransferFromVaultsReportRequest gets summary statistics about transferring data from vaults external targets to this cohesity cluster

	**Privileges:** ```REPORTS_VIEW``` <br><br>A Vault can provide an additional Cloud Tier where cold data of the

Cohesity Cluster is stored in the Cloud.
A Vault can also provide archive storage for backup data.
This archive data is stored on Tapes and in Cloud Vaults.
*/
func (a *Client) GetDataTransferFromVaultsReportRequest(params *GetDataTransferFromVaultsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataTransferFromVaultsReportRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataTransferFromVaultsReportRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDataTransferFromVaultsReportRequest",
		Method:             "GET",
		PathPattern:        "/public/reports/dataTransferFromVaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDataTransferFromVaultsReportRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataTransferFromVaultsReportRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDataTransferFromVaultsReportRequestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetDataTransferToVaultsReportRequest gets summary statistics about transferring data from the cohesity cluster to vaults external targets

	**Privileges:** ```REPORTS_VIEW``` <br><br>A Vault can provide an additional Cloud Tier where cold data of the

Cohesity Cluster can be stored in the Cloud.
A Vault can also provide archive storage for backup data.
This archive data is stored on Tapes and in Cloud Vaults.
*/
func (a *Client) GetDataTransferToVaultsReportRequest(params *GetDataTransferToVaultsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataTransferToVaultsReportRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataTransferToVaultsReportRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDataTransferToVaultsReportRequest",
		Method:             "GET",
		PathPattern:        "/public/reports/dataTransferToVaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDataTransferToVaultsReportRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataTransferToVaultsReportRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDataTransferToVaultsReportRequestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetGdprReport gets gdpr report related information

**Privileges:** ```GDPR_VIEW``` <br><br>Returns the GDPR report information.
*/
func (a *Client) GetGdprReport(params *GetGdprReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGdprReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGdprReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGdprReport",
		Method:             "GET",
		PathPattern:        "/public/reports/gdpr",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGdprReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGdprReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGdprReportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetProtectedObjectsTrendsReportRequest gets protected objects trends grouped by days weeks months

	**Privileges:** ```REPORTS_VIEW, TENANT_VIEW``` <br><br>This gives a summary of protection trend for protected resources during the

given time range.
If no roll up is specified, then the trends will be grouped by days.
*/
func (a *Client) GetProtectedObjectsTrendsReportRequest(params *GetProtectedObjectsTrendsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectedObjectsTrendsReportRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectedObjectsTrendsReportRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectedObjectsTrendsReportRequest",
		Method:             "GET",
		PathPattern:        "/public/reports/protectedObjectsTrends",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectedObjectsTrendsReportRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProtectedObjectsTrendsReportRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectedObjectsTrendsReportRequestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetProtectionSourcesJobRunsReportRequest gets protection details about the specified list of leaf protection source objects such as a v ms

	**Privileges:** ```REPORTS_VIEW, TENANT_VIEW``` <br><br>Returns the Snapshots that contain backups of the specified

Protection Source Objects and match the specified filter criteria.
*/
func (a *Client) GetProtectionSourcesJobRunsReportRequest(params *GetProtectionSourcesJobRunsReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionSourcesJobRunsReportRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionSourcesJobRunsReportRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionSourcesJobRunsReportRequest",
		Method:             "GET",
		PathPattern:        "/public/reports/protectionSourcesJobRuns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionSourcesJobRunsReportRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProtectionSourcesJobRunsReportRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionSourcesJobRunsReportRequestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetProtectionSourcesJobsSummaryReportRequest gets job run snapshot summary statistics about the leaf protection sources objects that match the specified filter criteria

	**Privileges:** ```REPORTS_VIEW, TENANT_VIEW``` <br><br>For example, if two Job ids are passed in, Snapshot summary statistics about

all the leaf Objects that have been protected by the two specified
Jobs are reported.
For example, if a top level registered Source id is passed in, summary
statistics about all the Snapshots backing up leaf Objects in
the specified Source are reported.
*/
func (a *Client) GetProtectionSourcesJobsSummaryReportRequest(params *GetProtectionSourcesJobsSummaryReportRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionSourcesJobsSummaryReportRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionSourcesJobsSummaryReportRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionSourcesJobsSummaryReportRequest",
		Method:             "GET",
		PathPattern:        "/public/reports/protectionSourcesJobsSummary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionSourcesJobsSummaryReportRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProtectionSourcesJobsSummaryReportRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionSourcesJobsSummaryReportRequestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
