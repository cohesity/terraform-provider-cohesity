// Code generated by go-swagger; DO NOT EDIT.

package active_directory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new active directory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new active directory API client with basic auth credentials.
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

// New creates a new active directory API client with a bearer token for authentication.
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
Client for active directory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddActiveDirectoryPrincipals(params *AddActiveDirectoryPrincipalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddActiveDirectoryPrincipalsOK, error)

	CreateActiveDirectory(params *CreateActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateActiveDirectoryCreated, error)

	DeleteActiveDirectory(params *DeleteActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteActiveDirectoryNoContent, error)

	GetActiveDirectory(params *GetActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveDirectoryOK, error)

	GetActiveDirectoryByID(params *GetActiveDirectoryByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveDirectoryByIDOK, error)

	GetActiveDirectoryPrincipals(params *GetActiveDirectoryPrincipalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveDirectoryPrincipalsOK, error)

	GetCentrifyZones(params *GetCentrifyZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCentrifyZonesOK, error)

	GetDomainControllers(params *GetDomainControllersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomainControllersOK, error)

	GetTrustedDomains(params *GetTrustedDomainsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTrustedDomainsOK, error)

	TriggerTrustedDomainsDiscovery(params *TriggerTrustedDomainsDiscoveryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerTrustedDomainsDiscoveryAccepted, error)

	UpdateActiveDirectory(params *UpdateActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateActiveDirectoryOK, error)

	UpdateTrustedDomains(params *UpdateTrustedDomainsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTrustedDomainsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	AddActiveDirectoryPrincipals adds multiple groups or users on the cohesity cluster for the specified active directory principals in addition assign cohesity roles to the users or groups to define their cohesity privileges

	**Privileges:** ```PRINCIPAL_MODIFY``` <br><br>After a group or user has been added to a Cohesity Cluster, the referenced Active Directory principal can be used by the Cohesity Cluster. In addition, this operation maps Cohesity roles with a group or user and this mapping defines the privileges allowed on the Cohesity Cluster for the group or user. For example if an 'management' group is created on the Cohesity Cluster for the Active Directory 'management' principal group and is associated with the Cohesity 'View' role, all users in the referenced Active Directory 'management' principal group can log in to the Cohesity Dashboard but will only have view-only privileges. These users cannot create new Protection Jobs, Policies, Views, etc.

NOTE: Local Cohesity users and groups cannot be created by this operation. Local Cohesity users or groups do not have an associated Active Directory principals and are created directly in the default LOCAL domain.
*/
func (a *Client) AddActiveDirectoryPrincipals(params *AddActiveDirectoryPrincipalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddActiveDirectoryPrincipalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddActiveDirectoryPrincipalsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddActiveDirectoryPrincipals",
		Method:             "POST",
		PathPattern:        "/active-directory-principals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddActiveDirectoryPrincipalsReader{formats: a.formats},
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
	success, ok := result.(*AddActiveDirectoryPrincipalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddActiveDirectoryPrincipalsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateActiveDirectory creates an active directory

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>Create an Active Directory.
*/
func (a *Client) CreateActiveDirectory(params *CreateActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateActiveDirectoryCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateActiveDirectoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateActiveDirectory",
		Method:             "POST",
		PathPattern:        "/active-directories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateActiveDirectoryReader{formats: a.formats},
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
	success, ok := result.(*CreateActiveDirectoryCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateActiveDirectoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteActiveDirectory deletes an active directory

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>Delete an Active Directory.
*/
func (a *Client) DeleteActiveDirectory(params *DeleteActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteActiveDirectoryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteActiveDirectoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteActiveDirectory",
		Method:             "DELETE",
		PathPattern:        "/active-directories/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteActiveDirectoryReader{formats: a.formats},
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
	success, ok := result.(*DeleteActiveDirectoryNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteActiveDirectoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetActiveDirectory gets the list of active directories

**Privileges:** ```AD_LDAP_VIEW``` <br><br>Get the list of Active Directories.
*/
func (a *Client) GetActiveDirectory(params *GetActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveDirectoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActiveDirectoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetActiveDirectory",
		Method:             "GET",
		PathPattern:        "/active-directories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActiveDirectoryReader{formats: a.formats},
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
	success, ok := result.(*GetActiveDirectoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActiveDirectoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetActiveDirectoryByID gets an active directory by id

**Privileges:** ```AD_LDAP_VIEW``` <br><br>Get an Active Directory by id.
*/
func (a *Client) GetActiveDirectoryByID(params *GetActiveDirectoryByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveDirectoryByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActiveDirectoryByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetActiveDirectoryById",
		Method:             "GET",
		PathPattern:        "/active-directories/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActiveDirectoryByIDReader{formats: a.formats},
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
	success, ok := result.(*GetActiveDirectoryByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActiveDirectoryByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetActiveDirectoryPrincipals gets the list of user and group principals from the active directory that match the specified filter criteria

**Privileges:** ```PRINCIPAL_VIEW, AD_LDAP_VIEW``` <br><br>Get the list of user and group principals from the Active Directory that match the specified filter criteria.
*/
func (a *Client) GetActiveDirectoryPrincipals(params *GetActiveDirectoryPrincipalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActiveDirectoryPrincipalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActiveDirectoryPrincipalsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetActiveDirectoryPrincipals",
		Method:             "GET",
		PathPattern:        "/active-directory-principals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActiveDirectoryPrincipalsReader{formats: a.formats},
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
	success, ok := result.(*GetActiveDirectoryPrincipalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActiveDirectoryPrincipalsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCentrifyZones gets centrify zones

**Privileges:** ```AD_LDAP_VIEW``` <br><br>Get Centrify zones for a specified domain.
*/
func (a *Client) GetCentrifyZones(params *GetCentrifyZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCentrifyZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentrifyZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCentrifyZones",
		Method:             "GET",
		PathPattern:        "/centrify-zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentrifyZonesReader{formats: a.formats},
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
	success, ok := result.(*GetCentrifyZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCentrifyZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDomainControllers gets domain controllers of specified domains

**Privileges:** ```AD_LDAP_VIEW``` <br><br>Get Domain Controllers of specified domains.
*/
func (a *Client) GetDomainControllers(params *GetDomainControllersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDomainControllersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainControllersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDomainControllers",
		Method:             "GET",
		PathPattern:        "/domain-controllers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainControllersReader{formats: a.formats},
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
	success, ok := result.(*GetDomainControllersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDomainControllersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTrustedDomains gets trusted domains

**Privileges:** ```AD_LDAP_VIEW``` <br><br>Get Trusted Domains for a specified domain.
*/
func (a *Client) GetTrustedDomains(params *GetTrustedDomainsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTrustedDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrustedDomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTrustedDomains",
		Method:             "GET",
		PathPattern:        "/trusted-domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrustedDomainsReader{formats: a.formats},
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
	success, ok := result.(*GetTrustedDomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTrustedDomainsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
TriggerTrustedDomainsDiscovery rediscovers trusted domains

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>Re-trigger the trusted domains of an Active Directory.
*/
func (a *Client) TriggerTrustedDomainsDiscovery(params *TriggerTrustedDomainsDiscoveryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerTrustedDomainsDiscoveryAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerTrustedDomainsDiscoveryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TriggerTrustedDomainsDiscovery",
		Method:             "PUT",
		PathPattern:        "/trusted-domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TriggerTrustedDomainsDiscoveryReader{formats: a.formats},
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
	success, ok := result.(*TriggerTrustedDomainsDiscoveryAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TriggerTrustedDomainsDiscoveryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateActiveDirectory updates an active directory

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>Update an Active Directory.
*/
func (a *Client) UpdateActiveDirectory(params *UpdateActiveDirectoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateActiveDirectoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateActiveDirectoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateActiveDirectory",
		Method:             "PUT",
		PathPattern:        "/active-directories/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateActiveDirectoryReader{formats: a.formats},
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
	success, ok := result.(*UpdateActiveDirectoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateActiveDirectoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateTrustedDomains updates trusted domains

**Privileges:** ```AD_LDAP_MODIFY``` <br><br>To update trusted domains of an Active Directory.
*/
func (a *Client) UpdateTrustedDomains(params *UpdateTrustedDomainsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTrustedDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTrustedDomainsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTrustedDomains",
		Method:             "POST",
		PathPattern:        "/trusted-domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTrustedDomainsReader{formats: a.formats},
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
	success, ok := result.(*UpdateTrustedDomainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTrustedDomainsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
