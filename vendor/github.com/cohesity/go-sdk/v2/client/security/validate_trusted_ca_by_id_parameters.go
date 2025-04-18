// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewValidateTrustedCaByIDParams creates a new ValidateTrustedCaByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateTrustedCaByIDParams() *ValidateTrustedCaByIDParams {
	return &ValidateTrustedCaByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateTrustedCaByIDParamsWithTimeout creates a new ValidateTrustedCaByIDParams object
// with the ability to set a timeout on a request.
func NewValidateTrustedCaByIDParamsWithTimeout(timeout time.Duration) *ValidateTrustedCaByIDParams {
	return &ValidateTrustedCaByIDParams{
		timeout: timeout,
	}
}

// NewValidateTrustedCaByIDParamsWithContext creates a new ValidateTrustedCaByIDParams object
// with the ability to set a context for a request.
func NewValidateTrustedCaByIDParamsWithContext(ctx context.Context) *ValidateTrustedCaByIDParams {
	return &ValidateTrustedCaByIDParams{
		Context: ctx,
	}
}

// NewValidateTrustedCaByIDParamsWithHTTPClient creates a new ValidateTrustedCaByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateTrustedCaByIDParamsWithHTTPClient(client *http.Client) *ValidateTrustedCaByIDParams {
	return &ValidateTrustedCaByIDParams{
		HTTPClient: client,
	}
}

/*
ValidateTrustedCaByIDParams contains all the parameters to send to the API endpoint

	for the validate trusted ca by Id operation.

	Typically these are written to a http.Request.
*/
type ValidateTrustedCaByIDParams struct {

	/* ID.

	   Specifies the id of the certificate to be validated.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate trusted ca by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateTrustedCaByIDParams) WithDefaults() *ValidateTrustedCaByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate trusted ca by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateTrustedCaByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) WithTimeout(timeout time.Duration) *ValidateTrustedCaByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) WithContext(ctx context.Context) *ValidateTrustedCaByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) WithHTTPClient(client *http.Client) *ValidateTrustedCaByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) WithID(id string) *ValidateTrustedCaByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate trusted ca by Id params
func (o *ValidateTrustedCaByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateTrustedCaByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
