// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new stats API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new stats API client with basic auth credentials.
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

// New creates a new stats API client with a bearer token for authentication.
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
Client for stats API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetActiveAlertsStats(params *GetActiveAlertsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveAlertsStatsOK, error)

	GetConsumerStats(params *GetConsumerStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsumerStatsOK, error)

	GetFileDistributionStats(params *GetFileDistributionStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileDistributionStatsOK, error)

	GetLastProtectionRunStats(params *GetLastProtectionRunStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLastProtectionRunStatsOK, error)

	GetProtectedObjectsSummary(params *GetProtectedObjectsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectedObjectsSummaryOK, error)

	GetProtectionRunsStats(params *GetProtectionRunsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunsStatsOK, error)

	GetRestoreStats(params *GetRestoreStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRestoreStatsOK, error)

	GetStorageStats(params *GetStorageStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageStatsOK, error)

	GetTenantStats(params *GetTenantStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantStatsOK, error)

	GetVaultProviderStats(params *GetVaultProviderStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVaultProviderStatsOK, error)

	GetVaultRunStats(params *GetVaultRunStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVaultRunStatsOK, error)

	GetVaultStats(params *GetVaultStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVaultStatsOK, error)

	GetViewBoxStats(params *GetViewBoxStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetViewBoxStatsOK, error)

	GetViewProtocolStats(params *GetViewProtocolStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetViewProtocolStatsOK, error)

	GetViewStats(params *GetViewStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetViewStatsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetActiveAlertsStats computes the statistics on the active alerts generated on the cluster

**Privileges:** ```ALERT_VIEW``` <br><br>Compute the statistics on the active Alerts generated on the cluster based on the provided time interval.
*/
func (a *Client) GetActiveAlertsStats(params *GetActiveAlertsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveAlertsStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActiveAlertsStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetActiveAlertsStats",
		Method:             "GET",
		PathPattern:        "/public/stats/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActiveAlertsStatsReader{formats: a.formats},
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
	success, ok := result.(*GetActiveAlertsStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActiveAlertsStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetConsumerStats gets the statistics of consumers

**Privileges:** ```STORAGE_VIEW, PROTECTION_VIEW``` <br><br>
*/
func (a *Client) GetConsumerStats(params *GetConsumerStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetConsumerStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConsumerStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetConsumerStats",
		Method:             "GET",
		PathPattern:        "/public/stats/consumers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConsumerStatsReader{formats: a.formats},
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
	success, ok := result.(*GetConsumerStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetConsumerStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFileDistributionStats computes the file distribution statistics on a given entity like cluster or storage domain

**Privileges:** ```CLUSTER_VIEW``` <br><br>Compute the file distribution statistics on a given entity like cluster or storage domain.
*/
func (a *Client) GetFileDistributionStats(params *GetFileDistributionStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileDistributionStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileDistributionStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFileDistributionStats",
		Method:             "GET",
		PathPattern:        "/public/stats/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileDistributionStatsReader{formats: a.formats},
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
	success, ok := result.(*GetFileDistributionStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFileDistributionStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLastProtectionRunStats computes stats on last protection run for every job

**Privileges:** ```PROTECTION_VIEW``` <br><br>Compute stats on last Protection Run for every job.
*/
func (a *Client) GetLastProtectionRunStats(params *GetLastProtectionRunStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLastProtectionRunStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLastProtectionRunStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLastProtectionRunStats",
		Method:             "GET",
		PathPattern:        "/public/stats/protectionRuns/lastRun",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLastProtectionRunStatsReader{formats: a.formats},
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
	success, ok := result.(*GetLastProtectionRunStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLastProtectionRunStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectedObjectsSummary computes the protected objects summary on the cluster

**Privileges:** ```PROTECTION_VIEW``` <br><br>Computes the protected objects summary on the cluster.
*/
func (a *Client) GetProtectedObjectsSummary(params *GetProtectedObjectsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectedObjectsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectedObjectsSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectedObjectsSummary",
		Method:             "GET",
		PathPattern:        "/public/stats/protectionSummary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectedObjectsSummaryReader{formats: a.formats},
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
	success, ok := result.(*GetProtectedObjectsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectedObjectsSummaryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProtectionRunsStats computes the statistics on the protection runs for the cluster

**Privileges:** ```PROTECTION_VIEW``` <br><br>Compute the statistics of the Protection Runs based on the input filters. This endpoint provides a snapshot of count of Protection Runs of a specified status on a specified time interval.
*/
func (a *Client) GetProtectionRunsStats(params *GetProtectionRunsStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProtectionRunsStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProtectionRunsStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProtectionRunsStats",
		Method:             "GET",
		PathPattern:        "/public/stats/protectionRuns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProtectionRunsStatsReader{formats: a.formats},
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
	success, ok := result.(*GetProtectionRunsStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProtectionRunsStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRestoreStats computes the statistics on the restore tasks on the cluster

**Privileges:** ```RESTORE_VIEW``` <br><br>Compute the statistics on the Restore tasks on the cluster based on the provided time interval.
*/
func (a *Client) GetRestoreStats(params *GetRestoreStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRestoreStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRestoreStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRestoreStats",
		Method:             "GET",
		PathPattern:        "/public/stats/restores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRestoreStatsReader{formats: a.formats},
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
	success, ok := result.(*GetRestoreStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRestoreStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStorageStats computes the storage stats on the cluster

**Privileges:** ```CLUSTER_VIEW``` <br><br>Computes the storage stats on the cluster.
*/
func (a *Client) GetStorageStats(params *GetStorageStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStorageStats",
		Method:             "GET",
		PathPattern:        "/public/stats/storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageStatsReader{formats: a.formats},
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
	success, ok := result.(*GetStorageStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStorageStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTenantStats gets the statistics of organizations tenant

**Privileges:** ```STORAGE_VIEW``` <br><br>
*/
func (a *Client) GetTenantStats(params *GetTenantStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTenantStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTenantStats",
		Method:             "GET",
		PathPattern:        "/public/stats/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTenantStatsReader{formats: a.formats},
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
	success, ok := result.(*GetTenantStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTenantStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVaultProviderStats computes the size and count of entities on vaults

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_VIEW``` <br><br>Compute the size and count of entities on vaults.
*/
func (a *Client) GetVaultProviderStats(params *GetVaultProviderStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVaultProviderStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVaultProviderStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVaultProviderStats",
		Method:             "GET",
		PathPattern:        "/public/stats/vaults/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVaultProviderStatsReader{formats: a.formats},
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
	success, ok := result.(*GetVaultProviderStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVaultProviderStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVaultRunStats computes the statistics on protection runs to or from a vault

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_VIEW``` <br><br>Compute the statistics on protection runs to or from a vault and return a trend line of the result.
*/
func (a *Client) GetVaultRunStats(params *GetVaultRunStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVaultRunStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVaultRunStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVaultRunStats",
		Method:             "GET",
		PathPattern:        "/public/stats/vaults/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVaultRunStatsReader{formats: a.formats},
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
	success, ok := result.(*GetVaultRunStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVaultRunStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVaultStats computes the statistics on vaults

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_VIEW``` <br><br>Compute the statistics on vaults.
*/
func (a *Client) GetVaultStats(params *GetVaultStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVaultStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVaultStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVaultStats",
		Method:             "GET",
		PathPattern:        "/public/stats/vaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVaultStatsReader{formats: a.formats},
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
	success, ok := result.(*GetVaultStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVaultStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetViewBoxStats gets the statistics of view boxes storage domain

**Privileges:** ```STORAGE_DOMAIN_VIEW``` <br><br>
*/
func (a *Client) GetViewBoxStats(params *GetViewBoxStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetViewBoxStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetViewBoxStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetViewBoxStats",
		Method:             "GET",
		PathPattern:        "/public/stats/viewBoxes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetViewBoxStatsReader{formats: a.formats},
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
	success, ok := result.(*GetViewBoxStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetViewBoxStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetViewProtocolStats computes the protocol statistics on views

**Privileges:** ```STORAGE_VIEW``` <br><br>Compute the protocol statistics on Views.
*/
func (a *Client) GetViewProtocolStats(params *GetViewProtocolStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetViewProtocolStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetViewProtocolStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetViewProtocolStats",
		Method:             "GET",
		PathPattern:        "/public/stats/views/protocols",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetViewProtocolStatsReader{formats: a.formats},
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
	success, ok := result.(*GetViewProtocolStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetViewProtocolStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetViewStats computes the statistics on views

**Privileges:** ```STORAGE_VIEW``` <br><br>Compute the statistics on Views.
*/
func (a *Client) GetViewStats(params *GetViewStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetViewStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetViewStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetViewStats",
		Method:             "GET",
		PathPattern:        "/public/stats/views",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetViewStatsReader{formats: a.formats},
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
	success, ok := result.(*GetViewStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetViewStatsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
