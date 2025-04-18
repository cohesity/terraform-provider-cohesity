// Code generated by go-swagger; DO NOT EDIT.

package remote_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new remote storage API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new remote storage API client with basic auth credentials.
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

// New creates a new remote storage API client with a bearer token for authentication.
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
Client for remote storage API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteRemoteStorageRegistration(params *DeleteRemoteStorageRegistrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemoteStorageRegistrationNoContent, error)

	GetRegisteredRemoteStorageList(params *GetRegisteredRemoteStorageListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRegisteredRemoteStorageListOK, error)

	GetRemoteStorageDetails(params *GetRemoteStorageDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRemoteStorageDetailsOK, error)

	RegisterNewRemoteStorage(params *RegisterNewRemoteStorageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterNewRemoteStorageCreated, error)

	UpdateRemoteStorageRegistration(params *UpdateRemoteStorageRegistrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemoteStorageRegistrationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteRemoteStorageRegistration deletes remote storage registration

**Privileges:** ```CLUSTER_MODIFY``` <br><br>Delete remote storage registration.
*/
func (a *Client) DeleteRemoteStorageRegistration(params *DeleteRemoteStorageRegistrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteRemoteStorageRegistrationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRemoteStorageRegistrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRemoteStorageRegistration",
		Method:             "DELETE",
		PathPattern:        "/remote-storage/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRemoteStorageRegistrationReader{formats: a.formats},
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
	success, ok := result.(*DeleteRemoteStorageRegistrationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRemoteStorageRegistrationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRegisteredRemoteStorageList gets registered remote storage servers list

**Privileges:** ```CLUSTER_VIEW``` <br><br>Get summary about list of registered remote storage servers.
*/
func (a *Client) GetRegisteredRemoteStorageList(params *GetRegisteredRemoteStorageListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRegisteredRemoteStorageListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegisteredRemoteStorageListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRegisteredRemoteStorageList",
		Method:             "GET",
		PathPattern:        "/remote-storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRegisteredRemoteStorageListReader{formats: a.formats},
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
	success, ok := result.(*GetRegisteredRemoteStorageListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRegisteredRemoteStorageListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetRemoteStorageDetails gets remote storage details

**Privileges:** ```CLUSTER_VIEW``` <br><br>Get details of remote storage given by id.
*/
func (a *Client) GetRemoteStorageDetails(params *GetRemoteStorageDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRemoteStorageDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRemoteStorageDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRemoteStorageDetails",
		Method:             "GET",
		PathPattern:        "/remote-storage/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRemoteStorageDetailsReader{formats: a.formats},
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
	success, ok := result.(*GetRemoteStorageDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRemoteStorageDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RegisterNewRemoteStorage registers remote storage

**Privileges:** ```CLUSTER_MODIFY``` <br><br>Register a remote storage to be used for disaggregated storage.
*/
func (a *Client) RegisterNewRemoteStorage(params *RegisterNewRemoteStorageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegisterNewRemoteStorageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterNewRemoteStorageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegisterNewRemoteStorage",
		Method:             "POST",
		PathPattern:        "/remote-storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegisterNewRemoteStorageReader{formats: a.formats},
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
	success, ok := result.(*RegisterNewRemoteStorageCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterNewRemoteStorageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateRemoteStorageRegistration updates remote storage config

**Privileges:** ```CLUSTER_MODIFY``` <br><br>Update Registered Remote Storage Config.
*/
func (a *Client) UpdateRemoteStorageRegistration(params *UpdateRemoteStorageRegistrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRemoteStorageRegistrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRemoteStorageRegistrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRemoteStorageRegistration",
		Method:             "PATCH",
		PathPattern:        "/remote-storage/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRemoteStorageRegistrationReader{formats: a.formats},
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
	success, ok := result.(*UpdateRemoteStorageRegistrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRemoteStorageRegistrationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
