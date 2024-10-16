// Code generated by go-swagger; DO NOT EDIT.

package s_m_b_file_opens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new s m b file opens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new s m b file opens API client with basic auth credentials.
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

// New creates a new s m b file opens API client with a bearer token for authentication.
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
Client for s m b file opens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CloseSmbFileOpen(params *CloseSmbFileOpenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloseSmbFileOpenNoContent, error)

	GetSmbFileOpens(params *GetSmbFileOpensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSmbFileOpensOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CloseSmbFileOpen closes an active s m b file open

**Privileges:** ```STORAGE_MODIFY``` <br><br>Returns nothing upon success.
*/
func (a *Client) CloseSmbFileOpen(params *CloseSmbFileOpenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloseSmbFileOpenNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloseSmbFileOpenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CloseSmbFileOpen",
		Method:             "POST",
		PathPattern:        "/public/smbFileOpens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CloseSmbFileOpenReader{formats: a.formats},
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
	success, ok := result.(*CloseSmbFileOpenNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CloseSmbFileOpenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	GetSmbFileOpens lists the active s m b file opens that match the filter criteria specified using parameters

	**Privileges:** ```STORAGE_VIEW``` <br><br>If no parameters are specified, all active SMB file opens currently on the

Cohesity Cluster are returned. Specifying parameters filters the results that
are returned.
*/
func (a *Client) GetSmbFileOpens(params *GetSmbFileOpensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSmbFileOpensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSmbFileOpensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSmbFileOpens",
		Method:             "GET",
		PathPattern:        "/public/smbFileOpens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSmbFileOpensReader{formats: a.formats},
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
	success, ok := result.(*GetSmbFileOpensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSmbFileOpensDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
