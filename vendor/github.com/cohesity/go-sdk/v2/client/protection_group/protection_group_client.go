// Code generated by go-swagger; DO NOT EDIT.

package protection_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new protection group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new protection group API client with basic auth credentials.
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

// New creates a new protection group API client with a bearer token for authentication.
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
Client for protection group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProtectionGroup(params *CreateProtectionGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProtectionGroupCreated, error)

	CreateProtectionGroupRun(params *CreateProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProtectionGroupRunAccepted, error)

	DeleteProtectionGroup(params *DeleteProtectionGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProtectionGroupNoContent, error)

	GetProtectionGroupByID(params *GetProtectionGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupByIDOK, error)

	GetProtectionGroupRun(params *GetProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupRunOK, error)

	GetProtectionGroupRuns(params *GetProtectionGroupRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupRunsOK, error)

	GetProtectionGroups(params *GetProtectionGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupsOK, error)

	GetProtectionRunProgress(params *GetProtectionRunProgressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunProgressOK, error)

	GetProtectionRunStats(params *GetProtectionRunStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunStatsOK, error)

	GetProtectionRuns(params *GetProtectionRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunsOK, error)

	GetRunDebugLogs(params *GetRunDebugLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunDebugLogsOK, error)

	GetRunDebugLogsForObject(params *GetRunDebugLogsForObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunDebugLogsForObjectOK, error)

	GetRunMessagesReport(params *GetRunMessagesReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunMessagesReportOK, error)

	GetRunsReport(params *GetRunsReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunsReportOK, error)

	PerformActionOnProtectionGroupRun(params *PerformActionOnProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PerformActionOnProtectionGroupRunAccepted, error)

	UpdateProtectionGroup(params *UpdateProtectionGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProtectionGroupOK, error)

	UpdateProtectionGroupRun(params *UpdateProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProtectionGroupRunMultiStatus, error)

	UpdateProtectionGroupsState(params *UpdateProtectionGroupsStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProtectionGroupsStateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateProtectionGroup creates a protection group

**Privileges:** ```PROTECTION_MODIFY``` <br><br>Create a Protection Group.
*/
func (a *Client) CreateProtectionGroup(params *CreateProtectionGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProtectionGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProtectionGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateProtectionGroup",
		Method:             "POST",
		PathPattern:        "/data-protect/protection-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProtectionGroupReader{formats: a.formats},
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
	success, ok := result.(*CreateProtectionGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateProtectionGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateProtectionGroupRun creates a new protection run

**Privileges:** ```PROTECTION_JOB_OPERATE``` <br><br>Create a new protection run. This can be used to start a run for a Protection Group on demand, ignoring the schedule and retention specified in the protection policy.
*/
func (a *Client) CreateProtectionGroupRun(params *CreateProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateProtectionGroupRunAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProtectionGroupRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateProtectionGroupRun",
		Method:             "POST",
		PathPattern:        "/data-protect/protection-groups/{id}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProtectionGroupRunReader{formats: a.formats},
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
	success, ok := result.(*CreateProtectionGroupRunAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateProtectionGroupRunDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteProtectionGroup deletes a protection group

**Privileges:** ```PROTECTION_MODIFY``` <br><br>Returns Success if the Protection Group is deleted.
*/
func (a *Client) DeleteProtectionGroup(params *DeleteProtectionGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteProtectionGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProtectionGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteProtectionGroup",
		Method:             "DELETE",
		PathPattern:        "/data-protect/protection-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProtectionGroupReader{formats: a.formats},
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
	success, ok := result.(*DeleteProtectionGroupNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteProtectionGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionGroupByID lists details about single protection group

**Privileges:** ```PROTECTION_VIEW``` <br><br>Returns the Protection Group corresponding to the specified Group id.
*/
func (a *Client) GetProtectionGroupByID(params *GetProtectionGroupByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionGroupByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionGroupById",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionGroupByIDReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionGroupByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionGroupByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionGroupRun gets a run for a protection group

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get a run for a particular Protection Group.
*/
func (a *Client) GetProtectionGroupRun(params *GetProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionGroupRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionGroupRun",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups/{id}/runs/{runId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionGroupRunReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionGroupRunOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionGroupRunDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionGroupRuns gets the list of runs for a protection group

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get the runs for a particular Protection Group.
*/
func (a *Client) GetProtectionGroupRuns(params *GetProtectionGroupRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionGroupRunsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionGroupRuns",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups/{id}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionGroupRunsReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionGroupRunsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionGroupRunsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionGroups gets the list of protection groups

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get the list of Protection Groups.
*/
func (a *Client) GetProtectionGroups(params *GetProtectionGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionGroups",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionRunProgress gets the progress of a run

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get the progress of a run.
*/
func (a *Client) GetProtectionRunProgress(params *GetProtectionRunProgressParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunProgressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionRunProgressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionRunProgress",
		Method:             "GET",
		PathPattern:        "/data-protect/runs/{runId}/progress",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionRunProgressReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionRunProgressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionRunProgressDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionRunStats gets the stats for a run

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get the stats for a run.
*/
func (a *Client) GetProtectionRunStats(params *GetProtectionRunStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionRunStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionRunStats",
		Method:             "GET",
		PathPattern:        "/data-protect/runs/{runId}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionRunStatsReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionRunStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionRunStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionRuns gets the list of runs

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get a list of protection runs.
*/
func (a *Client) GetProtectionRuns(params *GetProtectionRunsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionRunsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionRuns",
		Method:             "GET",
		PathPattern:        "/data-protect/runs/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionRunsReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionRunsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionRunsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRunDebugLogs gets the debug logs for a run from a protection group

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get the debug logs for all objects of a run for a particular Protection Group.
*/
func (a *Client) GetRunDebugLogs(params *GetRunDebugLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunDebugLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunDebugLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunDebugLogs",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups/{id}/runs/{runId}/debug-logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRunDebugLogsReader{formats: a.formats},
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
	success, ok := result.(*GetRunDebugLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRunDebugLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRunDebugLogsForObject gets the debug logs for a particular object in a run from a protection group

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get the debug logs for a particular object of a run for a particular Protection Group.
*/
func (a *Client) GetRunDebugLogsForObject(params *GetRunDebugLogsForObjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunDebugLogsForObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunDebugLogsForObjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunDebugLogsForObject",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/debug-logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRunDebugLogsForObjectReader{formats: a.formats},
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
	success, ok := result.(*GetRunDebugLogsForObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRunDebugLogsForObjectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRunMessagesReport gets the c s v of various proto messages for a given run and an object

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get an CSV report for given objectId and run id. Each row in CSV report contains the fields from correspoinding proto message.
*/
func (a *Client) GetRunMessagesReport(params *GetRunMessagesReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunMessagesReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunMessagesReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunMessagesReport",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/download-messages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRunMessagesReportReader{formats: a.formats},
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
	success, ok := result.(*GetRunMessagesReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRunMessagesReportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRunsReport gets the c s v of errors warnings for a given run and an object

**Privileges:** ```PROTECTION_VIEW``` <br><br>Get an CSV report for given objectId and run id. Report will depend on the query parameter fileType, default will be: success_files_list where each row contains the name of file backedup successfully.
*/
func (a *Client) GetRunsReport(params *GetRunsReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunsReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunsReportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunsReport",
		Method:             "GET",
		PathPattern:        "/data-protect/protection-groups/{id}/runs/{runId}/objects/{objectId}/downloadFiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRunsReportReader{formats: a.formats},
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
	success, ok := result.(*GetRunsReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRunsReportDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PerformActionOnProtectionGroupRun actions on protection group run

**Privileges:** ```PROTECTION_MODIFY``` <br><br>Perform various actions on a Protection Group run.
*/
func (a *Client) PerformActionOnProtectionGroupRun(params *PerformActionOnProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PerformActionOnProtectionGroupRunAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformActionOnProtectionGroupRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PerformActionOnProtectionGroupRun",
		Method:             "POST",
		PathPattern:        "/data-protect/protection-groups/{id}/runs/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformActionOnProtectionGroupRunReader{formats: a.formats},
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
	success, ok := result.(*PerformActionOnProtectionGroupRunAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PerformActionOnProtectionGroupRunDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateProtectionGroup updates a protection group

**Privileges:** ```PROTECTION_MODIFY``` <br><br>Update the specified Protection Group.
*/
func (a *Client) UpdateProtectionGroup(params *UpdateProtectionGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProtectionGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProtectionGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateProtectionGroup",
		Method:             "PUT",
		PathPattern:        "/data-protect/protection-groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProtectionGroupReader{formats: a.formats},
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
	success, ok := result.(*UpdateProtectionGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProtectionGroupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateProtectionGroupRun updates runs for a particular protection group

**Privileges:** ```PROTECTION_MODIFY``` <br><br>Update runs for a particular Protection Group. A user can perform the following actions: 1. Extend or reduce retention of a local, replication and archival snapshots. 2. Can perform resync operation on failed copy snapshots attempts in this Run. 3. Add new replication and archival snapshot targets to the Run. 4. Add or remove legal hold on the snapshots. Only a user with DSO role can perform this operation. 5. Delete the snapshots that were created as a part of this Run. 6. Apply datalock on existing snapshots where a user cannot manually delete snapshots before the expiry time.
*/
func (a *Client) UpdateProtectionGroupRun(params *UpdateProtectionGroupRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProtectionGroupRunMultiStatus, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProtectionGroupRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateProtectionGroupRun",
		Method:             "PUT",
		PathPattern:        "/data-protect/protection-groups/{id}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProtectionGroupRunReader{formats: a.formats},
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
	success, ok := result.(*UpdateProtectionGroupRunMultiStatus)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProtectionGroupRunDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateProtectionGroupsState performs an action like pause resume active deactivate on all specified protection groups

**Privileges:** ```PROTECTION_JOB_OPERATE``` <br><br>Perform an action like pause, resume, active, deactivate on all specified Protection Groups. Note that the pause or resume actions will take effect from next Protection Run. Also, user can specify only one type of action on all the Protection Groups. Deactivate and activate actions are independent of pause and resume state. Deactivate and activate actions are useful in case of failover situations. Returns success if the state of all the Protection Groups state is changed successfully.
*/
func (a *Client) UpdateProtectionGroupsState(params *UpdateProtectionGroupsStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateProtectionGroupsStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProtectionGroupsStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateProtectionGroupsState",
		Method:             "POST",
		PathPattern:        "/data-protect/protection-groups/states",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProtectionGroupsStateReader{formats: a.formats},
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
	success, ok := result.(*UpdateProtectionGroupsStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProtectionGroupsStateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
