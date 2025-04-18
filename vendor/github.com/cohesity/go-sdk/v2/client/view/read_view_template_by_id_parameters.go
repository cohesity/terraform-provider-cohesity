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

// NewReadViewTemplateByIDParams creates a new ReadViewTemplateByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadViewTemplateByIDParams() *ReadViewTemplateByIDParams {
	return &ReadViewTemplateByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadViewTemplateByIDParamsWithTimeout creates a new ReadViewTemplateByIDParams object
// with the ability to set a timeout on a request.
func NewReadViewTemplateByIDParamsWithTimeout(timeout time.Duration) *ReadViewTemplateByIDParams {
	return &ReadViewTemplateByIDParams{
		timeout: timeout,
	}
}

// NewReadViewTemplateByIDParamsWithContext creates a new ReadViewTemplateByIDParams object
// with the ability to set a context for a request.
func NewReadViewTemplateByIDParamsWithContext(ctx context.Context) *ReadViewTemplateByIDParams {
	return &ReadViewTemplateByIDParams{
		Context: ctx,
	}
}

// NewReadViewTemplateByIDParamsWithHTTPClient creates a new ReadViewTemplateByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadViewTemplateByIDParamsWithHTTPClient(client *http.Client) *ReadViewTemplateByIDParams {
	return &ReadViewTemplateByIDParams{
		HTTPClient: client,
	}
}

/*
ReadViewTemplateByIDParams contains all the parameters to send to the API endpoint

	for the read view template by Id operation.

	Typically these are written to a http.Request.
*/
type ReadViewTemplateByIDParams struct {

	/* ID.

	   Specifies a unique id of the view template.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read view template by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadViewTemplateByIDParams) WithDefaults() *ReadViewTemplateByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read view template by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadViewTemplateByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read view template by Id params
func (o *ReadViewTemplateByIDParams) WithTimeout(timeout time.Duration) *ReadViewTemplateByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read view template by Id params
func (o *ReadViewTemplateByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read view template by Id params
func (o *ReadViewTemplateByIDParams) WithContext(ctx context.Context) *ReadViewTemplateByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read view template by Id params
func (o *ReadViewTemplateByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read view template by Id params
func (o *ReadViewTemplateByIDParams) WithHTTPClient(client *http.Client) *ReadViewTemplateByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read view template by Id params
func (o *ReadViewTemplateByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the read view template by Id params
func (o *ReadViewTemplateByIDParams) WithID(id int64) *ReadViewTemplateByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read view template by Id params
func (o *ReadViewTemplateByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadViewTemplateByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
