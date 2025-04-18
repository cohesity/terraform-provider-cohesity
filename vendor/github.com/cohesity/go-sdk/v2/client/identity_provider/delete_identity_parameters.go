// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

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
	"github.com/go-openapi/swag"
)

// NewDeleteIdentityParams creates a new DeleteIdentityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIdentityParams() *DeleteIdentityParams {
	return &DeleteIdentityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIdentityParamsWithTimeout creates a new DeleteIdentityParams object
// with the ability to set a timeout on a request.
func NewDeleteIdentityParamsWithTimeout(timeout time.Duration) *DeleteIdentityParams {
	return &DeleteIdentityParams{
		timeout: timeout,
	}
}

// NewDeleteIdentityParamsWithContext creates a new DeleteIdentityParams object
// with the ability to set a context for a request.
func NewDeleteIdentityParamsWithContext(ctx context.Context) *DeleteIdentityParams {
	return &DeleteIdentityParams{
		Context: ctx,
	}
}

// NewDeleteIdentityParamsWithHTTPClient creates a new DeleteIdentityParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIdentityParamsWithHTTPClient(client *http.Client) *DeleteIdentityParams {
	return &DeleteIdentityParams{
		HTTPClient: client,
	}
}

/*
DeleteIdentityParams contains all the parameters to send to the API endpoint

	for the delete identity operation.

	Typically these are written to a http.Request.
*/
type DeleteIdentityParams struct {

	/* ID.

	   Specifies id of identity provider configuration

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdentityParams) WithDefaults() *DeleteIdentityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdentityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete identity params
func (o *DeleteIdentityParams) WithTimeout(timeout time.Duration) *DeleteIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete identity params
func (o *DeleteIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete identity params
func (o *DeleteIdentityParams) WithContext(ctx context.Context) *DeleteIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete identity params
func (o *DeleteIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete identity params
func (o *DeleteIdentityParams) WithHTTPClient(client *http.Client) *DeleteIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete identity params
func (o *DeleteIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete identity params
func (o *DeleteIdentityParams) WithID(id int64) *DeleteIdentityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete identity params
func (o *DeleteIdentityParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
