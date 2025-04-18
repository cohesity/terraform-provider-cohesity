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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewModifyCiphersParams creates a new ModifyCiphersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyCiphersParams() *ModifyCiphersParams {
	return &ModifyCiphersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyCiphersParamsWithTimeout creates a new ModifyCiphersParams object
// with the ability to set a timeout on a request.
func NewModifyCiphersParamsWithTimeout(timeout time.Duration) *ModifyCiphersParams {
	return &ModifyCiphersParams{
		timeout: timeout,
	}
}

// NewModifyCiphersParamsWithContext creates a new ModifyCiphersParams object
// with the ability to set a context for a request.
func NewModifyCiphersParamsWithContext(ctx context.Context) *ModifyCiphersParams {
	return &ModifyCiphersParams{
		Context: ctx,
	}
}

// NewModifyCiphersParamsWithHTTPClient creates a new ModifyCiphersParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyCiphersParamsWithHTTPClient(client *http.Client) *ModifyCiphersParams {
	return &ModifyCiphersParams{
		HTTPClient: client,
	}
}

/*
ModifyCiphersParams contains all the parameters to send to the API endpoint

	for the modify ciphers operation.

	Typically these are written to a http.Request.
*/
type ModifyCiphersParams struct {

	/* Body.

	   Enable/Disable ciphers.
	*/
	Body *models.ModifyCiphersRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify ciphers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyCiphersParams) WithDefaults() *ModifyCiphersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify ciphers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyCiphersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify ciphers params
func (o *ModifyCiphersParams) WithTimeout(timeout time.Duration) *ModifyCiphersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify ciphers params
func (o *ModifyCiphersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify ciphers params
func (o *ModifyCiphersParams) WithContext(ctx context.Context) *ModifyCiphersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify ciphers params
func (o *ModifyCiphersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify ciphers params
func (o *ModifyCiphersParams) WithHTTPClient(client *http.Client) *ModifyCiphersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify ciphers params
func (o *ModifyCiphersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify ciphers params
func (o *ModifyCiphersParams) WithBody(body *models.ModifyCiphersRequestBody) *ModifyCiphersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify ciphers params
func (o *ModifyCiphersParams) SetBody(body *models.ModifyCiphersRequestBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyCiphersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
