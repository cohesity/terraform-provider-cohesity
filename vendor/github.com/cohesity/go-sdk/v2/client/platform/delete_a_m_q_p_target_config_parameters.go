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

// NewDeleteAMQPTargetConfigParams creates a new DeleteAMQPTargetConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAMQPTargetConfigParams() *DeleteAMQPTargetConfigParams {
	return &DeleteAMQPTargetConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAMQPTargetConfigParamsWithTimeout creates a new DeleteAMQPTargetConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteAMQPTargetConfigParamsWithTimeout(timeout time.Duration) *DeleteAMQPTargetConfigParams {
	return &DeleteAMQPTargetConfigParams{
		timeout: timeout,
	}
}

// NewDeleteAMQPTargetConfigParamsWithContext creates a new DeleteAMQPTargetConfigParams object
// with the ability to set a context for a request.
func NewDeleteAMQPTargetConfigParamsWithContext(ctx context.Context) *DeleteAMQPTargetConfigParams {
	return &DeleteAMQPTargetConfigParams{
		Context: ctx,
	}
}

// NewDeleteAMQPTargetConfigParamsWithHTTPClient creates a new DeleteAMQPTargetConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAMQPTargetConfigParamsWithHTTPClient(client *http.Client) *DeleteAMQPTargetConfigParams {
	return &DeleteAMQPTargetConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteAMQPTargetConfigParams contains all the parameters to send to the API endpoint

	for the delete a m q p target config operation.

	Typically these are written to a http.Request.
*/
type DeleteAMQPTargetConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete a m q p target config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAMQPTargetConfigParams) WithDefaults() *DeleteAMQPTargetConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete a m q p target config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAMQPTargetConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete a m q p target config params
func (o *DeleteAMQPTargetConfigParams) WithTimeout(timeout time.Duration) *DeleteAMQPTargetConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete a m q p target config params
func (o *DeleteAMQPTargetConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete a m q p target config params
func (o *DeleteAMQPTargetConfigParams) WithContext(ctx context.Context) *DeleteAMQPTargetConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete a m q p target config params
func (o *DeleteAMQPTargetConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete a m q p target config params
func (o *DeleteAMQPTargetConfigParams) WithHTTPClient(client *http.Client) *DeleteAMQPTargetConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete a m q p target config params
func (o *DeleteAMQPTargetConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAMQPTargetConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
