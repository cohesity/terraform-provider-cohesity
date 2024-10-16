// Code generated by go-swagger; DO NOT EDIT.

package audit_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new audit log API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new audit log API client with basic auth credentials.
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

// New creates a new audit log API client with a bearer token for authentication.
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
Client for audit log API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAuditLogs(params *GetAuditLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogsOK, error)

	GetAuditLogsActions(params *GetAuditLogsActionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogsActionsOK, error)

	GetAuditLogsEntityTypes(params *GetAuditLogsEntityTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogsEntityTypesOK, error)

	GetFilerAuditLogConfigs(params *GetFilerAuditLogConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilerAuditLogConfigsOK, error)

	UpdateFilerAuditLogConfigs(params *UpdateFilerAuditLogConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFilerAuditLogConfigsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAuditLogs gets cluster audit logs

**Privileges:** ```CLUSTER_AUDIT``` <br><br>Get a cluster audit logs.
*/
func (a *Client) GetAuditLogs(params *GetAuditLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuditLogs",
		Method:             "GET",
		PathPattern:        "/audit-logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuditLogsReader{formats: a.formats},
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
	success, ok := result.(*GetAuditLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAuditLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAuditLogsActions gets cluster audit logs actions

**Privileges:** ```CLUSTER_AUDIT``` <br><br>Get all actions of cluster audit logs.
*/
func (a *Client) GetAuditLogsActions(params *GetAuditLogsActionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogsActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditLogsActionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuditLogsActions",
		Method:             "GET",
		PathPattern:        "/audit-logs/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuditLogsActionsReader{formats: a.formats},
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
	success, ok := result.(*GetAuditLogsActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAuditLogsActionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAuditLogsEntityTypes gets cluster audit logs entity types

**Privileges:** ```CLUSTER_AUDIT``` <br><br>Get all entity types of cluster audit logs.
*/
func (a *Client) GetAuditLogsEntityTypes(params *GetAuditLogsEntityTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogsEntityTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditLogsEntityTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuditLogsEntityTypes",
		Method:             "GET",
		PathPattern:        "/audit-logs/entity-types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuditLogsEntityTypesReader{formats: a.formats},
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
	success, ok := result.(*GetAuditLogsEntityTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAuditLogsEntityTypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetFilerAuditLogConfigs gets filer audit log configs

**Privileges:** ```CLUSTER_AUDIT``` <br><br>Get filer audit log configs.
*/
func (a *Client) GetFilerAuditLogConfigs(params *GetFilerAuditLogConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilerAuditLogConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilerAuditLogConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFilerAuditLogConfigs",
		Method:             "GET",
		PathPattern:        "/audit-logs/filer-configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilerAuditLogConfigsReader{formats: a.formats},
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
	success, ok := result.(*GetFilerAuditLogConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFilerAuditLogConfigsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateFilerAuditLogConfigs updates filer audit log configs

**Privileges:** ```CLUSTER_AUDIT``` <br><br>Update filer audit log configs.
*/
func (a *Client) UpdateFilerAuditLogConfigs(params *UpdateFilerAuditLogConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFilerAuditLogConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFilerAuditLogConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateFilerAuditLogConfigs",
		Method:             "PUT",
		PathPattern:        "/audit-logs/filer-configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFilerAuditLogConfigsReader{formats: a.formats},
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
	success, ok := result.(*UpdateFilerAuditLogConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateFilerAuditLogConfigsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
