// Code generated by go-swagger; DO NOT EDIT.

package views

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

	"github.com/cohesity/go-sdk/v1/models"
)

// NewRenameViewParams creates a new RenameViewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenameViewParams() *RenameViewParams {
	return &RenameViewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenameViewParamsWithTimeout creates a new RenameViewParams object
// with the ability to set a timeout on a request.
func NewRenameViewParamsWithTimeout(timeout time.Duration) *RenameViewParams {
	return &RenameViewParams{
		timeout: timeout,
	}
}

// NewRenameViewParamsWithContext creates a new RenameViewParams object
// with the ability to set a context for a request.
func NewRenameViewParamsWithContext(ctx context.Context) *RenameViewParams {
	return &RenameViewParams{
		Context: ctx,
	}
}

// NewRenameViewParamsWithHTTPClient creates a new RenameViewParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenameViewParamsWithHTTPClient(client *http.Client) *RenameViewParams {
	return &RenameViewParams{
		HTTPClient: client,
	}
}

/*
RenameViewParams contains all the parameters to send to the API endpoint

	for the rename view operation.

	Typically these are written to a http.Request.
*/
type RenameViewParams struct {

	/* Body.

	   Request to rename a View.
	*/
	Body *models.RenameViewParam

	/* Name.

	   Specifies the View name.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rename view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenameViewParams) WithDefaults() *RenameViewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rename view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenameViewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rename view params
func (o *RenameViewParams) WithTimeout(timeout time.Duration) *RenameViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rename view params
func (o *RenameViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rename view params
func (o *RenameViewParams) WithContext(ctx context.Context) *RenameViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rename view params
func (o *RenameViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rename view params
func (o *RenameViewParams) WithHTTPClient(client *http.Client) *RenameViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rename view params
func (o *RenameViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rename view params
func (o *RenameViewParams) WithBody(body *models.RenameViewParam) *RenameViewParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rename view params
func (o *RenameViewParams) SetBody(body *models.RenameViewParam) {
	o.Body = body
}

// WithName adds the name to the rename view params
func (o *RenameViewParams) WithName(name string) *RenameViewParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the rename view params
func (o *RenameViewParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RenameViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
