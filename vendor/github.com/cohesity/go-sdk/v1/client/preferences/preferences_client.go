// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new preferences API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new preferences API client with basic auth credentials.
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

// New creates a new preferences API client with a bearer token for authentication.
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
Client for preferences API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetUserPreferences(params *GetUserPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserPreferencesOK, error)

	PatchUserPreferences(params *PatchUserPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchUserPreferencesOK, error)

	UpdateUserPreferences(params *UpdateUserPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserPreferencesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetUserPreferences lists the preferences of the session user

```No Privileges Required``` <br><br>
*/
func (a *Client) GetUserPreferences(params *GetUserPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserPreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserPreferencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUserPreferences",
		Method:             "GET",
		PathPattern:        "/public/sessionUser/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserPreferencesReader{formats: a.formats},
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
	success, ok := result.(*GetUserPreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUserPreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchUserPreferences updates specific preferences of the session user

```No Privileges Required``` <br><br>Returns the updated user preferences.
*/
func (a *Client) PatchUserPreferences(params *PatchUserPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchUserPreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchUserPreferencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchUserPreferences",
		Method:             "PATCH",
		PathPattern:        "/public/sessionUser/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchUserPreferencesReader{formats: a.formats},
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
	success, ok := result.(*PatchUserPreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchUserPreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateUserPreferences updates the preferences of the session user

```No Privileges Required``` <br><br>Returns the updated user preferences.
*/
func (a *Client) UpdateUserPreferences(params *UpdateUserPreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateUserPreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserPreferencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateUserPreferences",
		Method:             "PUT",
		PathPattern:        "/public/sessionUser/preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserPreferencesReader{formats: a.formats},
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
	success, ok := result.(*UpdateUserPreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateUserPreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
