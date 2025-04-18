// Code generated by go-swagger; DO NOT EDIT.

package external_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new external target API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new external target API client with basic auth credentials.
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

// New creates a new external target API client with a bearer token for authentication.
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
Client for external target API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationOctetStream sets the Accept header to "application/octet-stream".
func WithAcceptApplicationOctetStream(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/octet-stream"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateExternalTarget(params *CreateExternalTargetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExternalTargetCreated, error)

	DeleteExternalTarget(params *DeleteExternalTargetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExternalTargetNoContent, error)

	GetExternalTargetByID(params *GetExternalTargetByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetByIDOK, error)

	GetExternalTargetEncryptionKeyInfo(params *GetExternalTargetEncryptionKeyInfoParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*GetExternalTargetEncryptionKeyInfoOK, error)

	GetExternalTargetMediaInfo(params *GetExternalTargetMediaInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetMediaInfoOK, error)

	GetExternalTargetSettings(params *GetExternalTargetSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetSettingsOK, error)

	GetExternalTargets(params *GetExternalTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetsOK, error)

	UpdateExternalTarget(params *UpdateExternalTargetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExternalTargetOK, error)

	UpdateExternalTargetSettings(params *UpdateExternalTargetSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExternalTargetSettingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateExternalTarget creates a external target

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_MODIFY``` <br><br>Create a External Target.
*/
func (a *Client) CreateExternalTarget(params *CreateExternalTargetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExternalTargetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExternalTargetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateExternalTarget",
		Method:             "POST",
		PathPattern:        "/data-protect/external-targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateExternalTargetReader{formats: a.formats},
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
	success, ok := result.(*CreateExternalTargetCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateExternalTargetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteExternalTarget deletes a external target

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_MODIFY``` <br><br>Returns Success if the External Target is deleted.
*/
func (a *Client) DeleteExternalTarget(params *DeleteExternalTargetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExternalTargetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExternalTargetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteExternalTarget",
		Method:             "DELETE",
		PathPattern:        "/data-protect/external-targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteExternalTargetReader{formats: a.formats},
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
	success, ok := result.(*DeleteExternalTargetNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteExternalTargetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetExternalTargetByID lists details about single external target

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_VIEW``` <br><br>Returns the External Target corresponding to the specified Group id.
*/
func (a *Client) GetExternalTargetByID(params *GetExternalTargetByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalTargetByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetExternalTargetById",
		Method:             "GET",
		PathPattern:        "/data-protect/external-targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalTargetByIDReader{formats: a.formats},
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
	success, ok := result.(*GetExternalTargetByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetExternalTargetByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetExternalTargetEncryptionKeyInfo gets the encryption key info for an external target

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_VIEW``` <br><br>Get the encryption key info for an external target
*/
func (a *Client) GetExternalTargetEncryptionKeyInfo(params *GetExternalTargetEncryptionKeyInfoParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*GetExternalTargetEncryptionKeyInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalTargetEncryptionKeyInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetExternalTargetEncryptionKeyInfo",
		Method:             "GET",
		PathPattern:        "/data-protect/external-targets/{id}/encryption-key",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalTargetEncryptionKeyInfoReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*GetExternalTargetEncryptionKeyInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetExternalTargetEncryptionKeyInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetExternalTargetMediaInfo lists archive media information

```Unknown Privileges``` <br><br>Returns the media information about the specified archive service uid (such as a QStar tape archive service).
*/
func (a *Client) GetExternalTargetMediaInfo(params *GetExternalTargetMediaInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetMediaInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalTargetMediaInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetExternalTargetMediaInfo",
		Method:             "GET",
		PathPattern:        "/data-protect/external-targets/media-info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalTargetMediaInfoReader{formats: a.formats},
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
	success, ok := result.(*GetExternalTargetMediaInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetExternalTargetMediaInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetExternalTargetSettings gets the list of external target settings

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_VIEW``` <br><br>Get the list of External Target Settings
*/
func (a *Client) GetExternalTargetSettings(params *GetExternalTargetSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalTargetSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetExternalTargetSettings",
		Method:             "GET",
		PathPattern:        "/data-protect/external-targets/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalTargetSettingsReader{formats: a.formats},
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
	success, ok := result.(*GetExternalTargetSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetExternalTargetSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetExternalTargets gets the list of external targets

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_VIEW``` <br><br>Get the list of External Targets.
*/
func (a *Client) GetExternalTargets(params *GetExternalTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExternalTargetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalTargetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetExternalTargets",
		Method:             "GET",
		PathPattern:        "/data-protect/external-targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalTargetsReader{formats: a.formats},
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
	success, ok := result.(*GetExternalTargetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetExternalTargetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateExternalTarget updates a external target

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_MODIFY``` <br><br>Update the specified External Target.
*/
func (a *Client) UpdateExternalTarget(params *UpdateExternalTargetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExternalTargetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExternalTargetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateExternalTarget",
		Method:             "PUT",
		PathPattern:        "/data-protect/external-targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExternalTargetReader{formats: a.formats},
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
	success, ok := result.(*UpdateExternalTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateExternalTargetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateExternalTargetSettings updates external target settings

**Privileges:** ```CLUSTER_EXTERNAL_TARGET_MODIFY``` <br><br>Update External Target Settings
*/
func (a *Client) UpdateExternalTargetSettings(params *UpdateExternalTargetSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExternalTargetSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExternalTargetSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateExternalTargetSettings",
		Method:             "PUT",
		PathPattern:        "/data-protect/external-targets/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExternalTargetSettingsReader{formats: a.formats},
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
	success, ok := result.(*UpdateExternalTargetSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateExternalTargetSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
