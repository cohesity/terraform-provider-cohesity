// Code generated by go-swagger; DO NOT EDIT.

package storage_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new storage domain API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new storage domain API client with basic auth credentials.
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

// New creates a new storage domain API client with a bearer token for authentication.
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
Client for storage domain API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateStorageDomain(params *CreateStorageDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateStorageDomainCreated, error)

	DeleteStorageDomain(params *DeleteStorageDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteStorageDomainNoContent, error)

	GetStorageDomainByID(params *GetStorageDomainByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageDomainByIDOK, error)

	GetStorageDomains(params *GetStorageDomainsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageDomainsOK, error)

	UpdateStorageDomain(params *UpdateStorageDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStorageDomainOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateStorageDomain creates a storage domain

**Privileges:** ```STORAGE_DOMAIN_MODIFY``` <br><br>Create a Storage Domain.
*/
func (a *Client) CreateStorageDomain(params *CreateStorageDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateStorageDomainCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStorageDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateStorageDomain",
		Method:             "POST",
		PathPattern:        "/storage-domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateStorageDomainReader{formats: a.formats},
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
	success, ok := result.(*CreateStorageDomainCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateStorageDomainDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteStorageDomain deletes a storage domain

**Privileges:** ```STORAGE_DOMAIN_MODIFY``` <br><br>Delete a Storage Domain.
*/
func (a *Client) DeleteStorageDomain(params *DeleteStorageDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteStorageDomainNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStorageDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteStorageDomain",
		Method:             "DELETE",
		PathPattern:        "/storage-domains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStorageDomainReader{formats: a.formats},
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
	success, ok := result.(*DeleteStorageDomainNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteStorageDomainDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStorageDomainByID gets a storage domain by id

**Privileges:** ```STORAGE_DOMAIN_VIEW``` <br><br>Get a Storage Domain by id.
*/
func (a *Client) GetStorageDomainByID(params *GetStorageDomainByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageDomainByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageDomainByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStorageDomainById",
		Method:             "GET",
		PathPattern:        "/storage-domains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageDomainByIDReader{formats: a.formats},
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
	success, ok := result.(*GetStorageDomainByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStorageDomainByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetStorageDomains gets storage domains

**Privileges:** ```STORAGE_DOMAIN_VIEW``` <br><br>Get Storage Domains.
*/
func (a *Client) GetStorageDomains(params *GetStorageDomainsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStorageDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageDomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStorageDomains",
		Method:             "GET",
		PathPattern:        "/storage-domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageDomainsReader{formats: a.formats},
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
	success, ok := result.(*GetStorageDomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStorageDomainsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateStorageDomain updates a storage domain

**Privileges:** ```STORAGE_DOMAIN_MODIFY``` <br><br>Update a Storage Domain.
*/
func (a *Client) UpdateStorageDomain(params *UpdateStorageDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStorageDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStorageDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateStorageDomain",
		Method:             "PUT",
		PathPattern:        "/storage-domains/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateStorageDomainReader{formats: a.formats},
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
	success, ok := result.(*UpdateStorageDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateStorageDomainDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
