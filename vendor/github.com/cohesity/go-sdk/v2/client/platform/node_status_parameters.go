// Code generated by go-swagger; DO NOT EDIT.

package platform

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewNodeStatusParams creates a new NodeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodeStatusParams() *NodeStatusParams {
	return &NodeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodeStatusParamsWithTimeout creates a new NodeStatusParams object
// with the ability to set a timeout on a request.
func NewNodeStatusParamsWithTimeout(timeout time.Duration) *NodeStatusParams {
	return &NodeStatusParams{
		timeout: timeout,
	}
}

// NewNodeStatusParamsWithContext creates a new NodeStatusParams object
// with the ability to set a context for a request.
func NewNodeStatusParamsWithContext(ctx context.Context) *NodeStatusParams {
	return &NodeStatusParams{
		Context: ctx,
	}
}

// NewNodeStatusParamsWithHTTPClient creates a new NodeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodeStatusParamsWithHTTPClient(client *http.Client) *NodeStatusParams {
	return &NodeStatusParams{
		HTTPClient: client,
	}
}

/*
NodeStatusParams contains all the parameters to send to the API endpoint

	for the node status operation.

	Typically these are written to a http.Request.
*/
type NodeStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeStatusParams) WithDefaults() *NodeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the node status params
func (o *NodeStatusParams) WithTimeout(timeout time.Duration) *NodeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node status params
func (o *NodeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node status params
func (o *NodeStatusParams) WithContext(ctx context.Context) *NodeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node status params
func (o *NodeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node status params
func (o *NodeStatusParams) WithHTTPClient(client *http.Client) *NodeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node status params
func (o *NodeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *NodeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
