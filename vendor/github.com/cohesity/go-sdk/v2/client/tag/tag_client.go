// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new tag API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new tag API client with basic auth credentials.
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

// New creates a new tag API client with a bearer token for authentication.
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
Client for tag API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTag(params *CreateTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTagCreated, error)

	DeleteTag(params *DeleteTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTagNoContent, error)

	GetTagByID(params *GetTagByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTagByIDOK, error)

	GetTags(params *GetTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTagsOK, error)

	UpdateTag(params *UpdateTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTagOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateTag creates a tag

**Privileges:** ```TAGS_MODIFY``` <br><br>Creates a Tag.
*/
func (a *Client) CreateTag(params *CreateTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateTagCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTag",
		Method:             "POST",
		PathPattern:        "/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTagReader{formats: a.formats},
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
	success, ok := result.(*CreateTagCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTagDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteTag deletes a tag

**Privileges:** ```TAGS_MODIFY``` <br><br>Deletes a Tag by id.
*/
func (a *Client) DeleteTag(params *DeleteTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTagNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTag",
		Method:             "DELETE",
		PathPattern:        "/tags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTagReader{formats: a.formats},
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
	success, ok := result.(*DeleteTagNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTagDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetTagByID gets tag by id

**Privileges:** ```TAGS_VIEW``` <br><br>Get Tag by id.
*/
func (a *Client) GetTagByID(params *GetTagByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTagByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTagById",
		Method:             "GET",
		PathPattern:        "/tags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTagByIDReader{formats: a.formats},
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
	success, ok := result.(*GetTagByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTagByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetTags gets tags based on filters

	**Privileges:** ```TAGS_VIEW``` <br><br>If no parameters are specified, all tags are returned.

Specifying parameters filters the results that are returned.
*/
func (a *Client) GetTags(params *GetTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTags",
		Method:             "GET",
		PathPattern:        "/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTagsReader{formats: a.formats},
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
	success, ok := result.(*GetTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTagsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateTag updates a tag

**Privileges:** ```TAGS_MODIFY``` <br><br>Updates a Tag by id.
*/
func (a *Client) UpdateTag(params *UpdateTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTag",
		Method:             "PUT",
		PathPattern:        "/tags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTagReader{formats: a.formats},
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
	success, ok := result.(*UpdateTagOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTagDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
