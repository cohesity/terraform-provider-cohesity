// Code generated by go-swagger; DO NOT EDIT.

package ldap_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new ldap provider API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new ldap provider API client with basic auth credentials.
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

// New creates a new ldap provider API client with a bearer token for authentication.
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
Client for ldap provider API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLdapProvider(params *CreateLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateLdapProviderCreated, error)

	DeleteLdapProvider(params *DeleteLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteLdapProviderNoContent, error)

	GetLdapProvider(params *GetLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLdapProviderOK, error)

	GetLdapProviderStatus(params *GetLdapProviderStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLdapProviderStatusCreated, error)

	UpdateLdapProvider(params *UpdateLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLdapProviderCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateLdapProvider creates a l d a p provider

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>Returns the created LDAP provider.
*/
func (a *Client) CreateLdapProvider(params *CreateLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateLdapProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLdapProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateLdapProvider",
		Method:             "POST",
		PathPattern:        "/public/ldapProvider",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLdapProviderReader{formats: a.formats},
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
	success, ok := result.(*CreateLdapProviderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateLdapProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteLdapProvider deletes an l d a p provider

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>
*/
func (a *Client) DeleteLdapProvider(params *DeleteLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteLdapProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLdapProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteLdapProvider",
		Method:             "DELETE",
		PathPattern:        "/public/ldapProvider/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLdapProviderReader{formats: a.formats},
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
	success, ok := result.(*DeleteLdapProviderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteLdapProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLdapProvider lists the l d a p providers

**Privileges:** ```AD_LDAP_VIEW``` <br><br>
*/
func (a *Client) GetLdapProvider(params *GetLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLdapProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLdapProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLdapProvider",
		Method:             "GET",
		PathPattern:        "/public/ldapProvider",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLdapProviderReader{formats: a.formats},
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
	success, ok := result.(*GetLdapProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLdapProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLdapProviderStatus gets the connection status of an l d a p provider

**Privileges:** ```AD_LDAP_VIEW``` <br><br>
*/
func (a *Client) GetLdapProviderStatus(params *GetLdapProviderStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLdapProviderStatusCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLdapProviderStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLdapProviderStatus",
		Method:             "GET",
		PathPattern:        "/public/ldapProvider/{id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLdapProviderStatusReader{formats: a.formats},
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
	success, ok := result.(*GetLdapProviderStatusCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLdapProviderStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateLdapProvider updates an l d a p provider

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>Returns the updated LDAP provider.
*/
func (a *Client) UpdateLdapProvider(params *UpdateLdapProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateLdapProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLdapProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateLdapProvider",
		Method:             "PUT",
		PathPattern:        "/public/ldapProvider",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLdapProviderReader{formats: a.formats},
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
	success, ok := result.(*UpdateLdapProviderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateLdapProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
