// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new identity provider API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new identity provider API client with basic auth credentials.
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

// New creates a new identity provider API client with a bearer token for authentication.
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
Client for identity provider API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateIdentity(params *CreateIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateIdentityCreated, error)

	CreateIdentityProvider(params *CreateIdentityProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateIdentityProviderCreated, error)

	DeleteIdentity(params *DeleteIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteIdentityNoContent, error)

	DeleteIdentityProvider(params *DeleteIdentityProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteIdentityProviderNoContent, error)

	GetIdentities(params *GetIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIdentitiesOK, error)

	GetIdentityProviders(params *GetIdentityProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIdentityProvidersOK, error)

	IdpsLogin(params *IdpsLoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	PerformIdentityAction(params *PerformIdentityActionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PerformIdentityActionCreated, error)

	UpdateIdentity(params *UpdateIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateIdentityOK, error)

	UpdateIdentityProvider(params *UpdateIdentityProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateIdentityProviderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateIdentity configures identity provider

**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>Configure Identity Provider on the cluster. Currently this API is only for Open ID providers, but will be expanded to include SAML providers in the future.
*/
func (a *Client) CreateIdentity(params *CreateIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateIdentityCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIdentityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateIdentity",
		Method:             "POST",
		PathPattern:        "/identity-providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIdentityReader{formats: a.formats},
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
	success, ok := result.(*CreateIdentityCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateIdentityDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateIdentityProvider configures identity provider

**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>Configure SAML based identity provider on the cluster
*/
func (a *Client) CreateIdentityProvider(params *CreateIdentityProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateIdentityProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIdentityProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateIdentityProvider",
		Method:             "POST",
		PathPattern:        "/idps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIdentityProviderReader{formats: a.formats},
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
	success, ok := result.(*CreateIdentityProviderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateIdentityProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteIdentity deletes identity provider

**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>Delete identity provider configuration on the cluster. Currently this API only supports Open ID based SSO providers, but it will be expanded in the future to support SAML SSO providers.
*/
func (a *Client) DeleteIdentity(params *DeleteIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteIdentityNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIdentityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteIdentity",
		Method:             "DELETE",
		PathPattern:        "/identity-providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIdentityReader{formats: a.formats},
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
	success, ok := result.(*DeleteIdentityNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIdentityDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteIdentityProvider deletes identity provider

**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>Delete SAML based identity provider configuration on the cluster
*/
func (a *Client) DeleteIdentityProvider(params *DeleteIdentityProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteIdentityProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIdentityProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteIdentityProvider",
		Method:             "DELETE",
		PathPattern:        "/idps/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIdentityProviderReader{formats: a.formats},
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
	success, ok := result.(*DeleteIdentityProviderNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteIdentityProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIdentities gets identities

**Privileges:** ```PRINCIPAL_VIEW``` <br><br>Get Identity Providers configured on the cluster. Currently this API only supports Open ID based SSO providers, but it will be expanded in the future to support SAML SSO providers.
*/
func (a *Client) GetIdentities(params *GetIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIdentitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIdentities",
		Method:             "GET",
		PathPattern:        "/identity-providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIdentitiesReader{formats: a.formats},
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
	success, ok := result.(*GetIdentitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIdentitiesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetIdentityProviders gets identity providers

**Privileges:** ```PRINCIPAL_VIEW``` <br><br>Get SAML based identity providers configured on the cluster
*/
func (a *Client) GetIdentityProviders(params *GetIdentityProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIdentityProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityProvidersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetIdentityProviders",
		Method:             "GET",
		PathPattern:        "/idps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIdentityProvidersReader{formats: a.formats},
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
	success, ok := result.(*GetIdentityProvidersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIdentityProvidersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdpsLogin logins to cluster using idp

```No Privileges Required``` <br><br>Redirects the client to the idp site with the URI to login
*/
func (a *Client) IdpsLogin(params *IdpsLoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdpsLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdpsLogin",
		Method:             "GET",
		PathPattern:        "/idps/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdpsLoginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
PerformIdentityAction performs identity action

**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>Perform an action on an Identity Provider. Currently this API only supports Open ID based SSO providers, but it will be expanded in the future to support SAML SSO providers.
*/
func (a *Client) PerformIdentityAction(params *PerformIdentityActionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PerformIdentityActionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformIdentityActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PerformIdentityAction",
		Method:             "POST",
		PathPattern:        "/identity-providers/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformIdentityActionReader{formats: a.formats},
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
	success, ok := result.(*PerformIdentityActionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PerformIdentityActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateIdentity updates identity provider

**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>Update Identity Provider on the cluster. Currently this API only supports Open ID based SSO providers, but it will be expanded in the future to support SAML SSO providers.
*/
func (a *Client) UpdateIdentity(params *UpdateIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIdentityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateIdentity",
		Method:             "PUT",
		PathPattern:        "/identity-providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIdentityReader{formats: a.formats},
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
	success, ok := result.(*UpdateIdentityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateIdentityDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateIdentityProvider updates identity provider

**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>Update SAML based identity provider configurartion on the cluster
*/
func (a *Client) UpdateIdentityProvider(params *UpdateIdentityProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateIdentityProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIdentityProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateIdentityProvider",
		Method:             "PUT",
		PathPattern:        "/idps/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIdentityProviderReader{formats: a.formats},
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
	success, ok := result.(*UpdateIdentityProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateIdentityProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
