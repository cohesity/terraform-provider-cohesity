// Code generated by go-swagger; DO NOT EDIT.

package app_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new app instance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new app instance API client with basic auth credentials.
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

// New creates a new app instance API client with a bearer token for authentication.
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
Client for app instance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAppInstances(params *GetAppInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAppInstancesOK, error)

	LaunchAppInstance(params *LaunchAppInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LaunchAppInstanceAccepted, error)

	UpdateAppInstanceSettings(params *UpdateAppInstanceSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAppInstanceSettingsAccepted, error)

	UpdateAppInstanceState(params *UpdateAppInstanceStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAppInstanceStateAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetAppInstances lists the app instances

	**Privileges:** ```APP_LAUNCH``` <br><br>Api provides the list of the app instances. Instances can be in different

states including stopped.
*/
func (a *Client) GetAppInstances(params *GetAppInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAppInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAppInstances",
		Method:             "GET",
		PathPattern:        "/public/appInstances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAppInstancesReader{formats: a.formats},
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
	success, ok := result.(*GetAppInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAppInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
LaunchAppInstance starts the application instance launch on the cluster

**Privileges:** ```APP_LAUNCH``` <br><br>Only installed apps can be launched.
*/
func (a *Client) LaunchAppInstance(params *LaunchAppInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LaunchAppInstanceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLaunchAppInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LaunchAppInstance",
		Method:             "POST",
		PathPattern:        "/public/appInstances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LaunchAppInstanceReader{formats: a.formats},
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
	success, ok := result.(*LaunchAppInstanceAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LaunchAppInstanceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateAppInstanceSettings updates app instance settings

**Privileges:** ```APP_LAUNCH``` <br><br>Changes the settings of the app instance.
*/
func (a *Client) UpdateAppInstanceSettings(params *UpdateAppInstanceSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAppInstanceSettingsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAppInstanceSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateAppInstanceSettings",
		Method:             "PUT",
		PathPattern:        "/public/appInstanceSettings/{appInstanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAppInstanceSettingsReader{formats: a.formats},
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
	success, ok := result.(*UpdateAppInstanceSettingsAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateAppInstanceSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateAppInstanceState updates app instance state

**Privileges:** ```APP_LAUNCH``` <br><br>Changes the state of the app instances.
*/
func (a *Client) UpdateAppInstanceState(params *UpdateAppInstanceStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateAppInstanceStateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAppInstanceStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateAppInstanceState",
		Method:             "PUT",
		PathPattern:        "/public/appInstances/{appInstanceId}/states",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAppInstanceStateReader{formats: a.formats},
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
	success, ok := result.(*UpdateAppInstanceStateAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateAppInstanceStateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
