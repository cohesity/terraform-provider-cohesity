// Code generated by go-swagger; DO NOT EDIT.

package view

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

// NewDeleteViewParams creates a new DeleteViewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteViewParams() *DeleteViewParams {
	return &DeleteViewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteViewParamsWithTimeout creates a new DeleteViewParams object
// with the ability to set a timeout on a request.
func NewDeleteViewParamsWithTimeout(timeout time.Duration) *DeleteViewParams {
	return &DeleteViewParams{
		timeout: timeout,
	}
}

// NewDeleteViewParamsWithContext creates a new DeleteViewParams object
// with the ability to set a context for a request.
func NewDeleteViewParamsWithContext(ctx context.Context) *DeleteViewParams {
	return &DeleteViewParams{
		Context: ctx,
	}
}

// NewDeleteViewParamsWithHTTPClient creates a new DeleteViewParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteViewParamsWithHTTPClient(client *http.Client) *DeleteViewParams {
	return &DeleteViewParams{
		HTTPClient: client,
	}
}

/*
DeleteViewParams contains all the parameters to send to the API endpoint

	for the delete view operation.

	Typically these are written to a http.Request.
*/
type DeleteViewParams struct {

	/* ID.

	   Specifies a unique id of the View to delete.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteViewParams) WithDefaults() *DeleteViewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteViewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete view params
func (o *DeleteViewParams) WithTimeout(timeout time.Duration) *DeleteViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete view params
func (o *DeleteViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete view params
func (o *DeleteViewParams) WithContext(ctx context.Context) *DeleteViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete view params
func (o *DeleteViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete view params
func (o *DeleteViewParams) WithHTTPClient(client *http.Client) *DeleteViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete view params
func (o *DeleteViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete view params
func (o *DeleteViewParams) WithID(id int64) *DeleteViewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete view params
func (o *DeleteViewParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
