// Code generated by go-swagger; DO NOT EDIT.

package keystone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new keystone API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new keystone API client with basic auth credentials.
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

// New creates a new keystone API client with a bearer token for authentication.
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
Client for keystone API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateKeystone(params *CreateKeystoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateKeystoneCreated, error)

	DeleteKeystone(params *DeleteKeystoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKeystoneNoContent, error)

	GetKeystones(params *GetKeystonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKeystonesOK, error)

	GetKeystonesByID(params *GetKeystonesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKeystonesByIDOK, error)

	UpdateKeystone(params *UpdateKeystoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateKeystoneOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateKeystone creates a keystone configuration

**Privileges:** ```KEYSTONE_MODIFY``` <br><br>Create a Keystone configuration.
*/
func (a *Client) CreateKeystone(params *CreateKeystoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateKeystoneCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateKeystoneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateKeystone",
		Method:             "POST",
		PathPattern:        "/keystones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateKeystoneReader{formats: a.formats},
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
	success, ok := result.(*CreateKeystoneCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateKeystoneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteKeystone deletes a keystone configuration

**Privileges:** ```KEYSTONE_MODIFY``` <br><br>Delete a Keystone configuration.
*/
func (a *Client) DeleteKeystone(params *DeleteKeystoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKeystoneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKeystoneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteKeystone",
		Method:             "DELETE",
		PathPattern:        "/keystones/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteKeystoneReader{formats: a.formats},
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
	success, ok := result.(*DeleteKeystoneNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteKeystoneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetKeystones gets keystones

**Privileges:** ```KEYSTONE_VIEW``` <br><br>Get Keystones.
*/
func (a *Client) GetKeystones(params *GetKeystonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKeystonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeystonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetKeystones",
		Method:             "GET",
		PathPattern:        "/keystones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeystonesReader{formats: a.formats},
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
	success, ok := result.(*GetKeystonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetKeystonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetKeystonesByID gets a keystone by its id

**Privileges:** ```KEYSTONE_VIEW``` <br><br>Get a Keystone by its id.
*/
func (a *Client) GetKeystonesByID(params *GetKeystonesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKeystonesByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeystonesByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetKeystonesById",
		Method:             "GET",
		PathPattern:        "/keystones/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeystonesByIDReader{formats: a.formats},
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
	success, ok := result.(*GetKeystonesByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetKeystonesByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateKeystone updates a keystone configuration

**Privileges:** ```KEYSTONE_MODIFY``` <br><br>Update a Keystone configuration.
*/
func (a *Client) UpdateKeystone(params *UpdateKeystoneParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateKeystoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateKeystoneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateKeystone",
		Method:             "PUT",
		PathPattern:        "/keystones/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateKeystoneReader{formats: a.formats},
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
	success, ok := result.(*UpdateKeystoneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateKeystoneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
