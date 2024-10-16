// Code generated by go-swagger; DO NOT EDIT.

package protection_sources

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

// NewRegisterProtectionSourceParams creates a new RegisterProtectionSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterProtectionSourceParams() *RegisterProtectionSourceParams {
	return &RegisterProtectionSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterProtectionSourceParamsWithTimeout creates a new RegisterProtectionSourceParams object
// with the ability to set a timeout on a request.
func NewRegisterProtectionSourceParamsWithTimeout(timeout time.Duration) *RegisterProtectionSourceParams {
	return &RegisterProtectionSourceParams{
		timeout: timeout,
	}
}

// NewRegisterProtectionSourceParamsWithContext creates a new RegisterProtectionSourceParams object
// with the ability to set a context for a request.
func NewRegisterProtectionSourceParamsWithContext(ctx context.Context) *RegisterProtectionSourceParams {
	return &RegisterProtectionSourceParams{
		Context: ctx,
	}
}

// NewRegisterProtectionSourceParamsWithHTTPClient creates a new RegisterProtectionSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterProtectionSourceParamsWithHTTPClient(client *http.Client) *RegisterProtectionSourceParams {
	return &RegisterProtectionSourceParams{
		HTTPClient: client,
	}
}

/*
RegisterProtectionSourceParams contains all the parameters to send to the API endpoint

	for the register protection source operation.

	Typically these are written to a http.Request.
*/
type RegisterProtectionSourceParams struct {

	/* Body.

	   Request to register a protection source.
	*/
	Body *models.RegisterProtectionSourceParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register protection source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterProtectionSourceParams) WithDefaults() *RegisterProtectionSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register protection source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterProtectionSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register protection source params
func (o *RegisterProtectionSourceParams) WithTimeout(timeout time.Duration) *RegisterProtectionSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register protection source params
func (o *RegisterProtectionSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register protection source params
func (o *RegisterProtectionSourceParams) WithContext(ctx context.Context) *RegisterProtectionSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register protection source params
func (o *RegisterProtectionSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register protection source params
func (o *RegisterProtectionSourceParams) WithHTTPClient(client *http.Client) *RegisterProtectionSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register protection source params
func (o *RegisterProtectionSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register protection source params
func (o *RegisterProtectionSourceParams) WithBody(body *models.RegisterProtectionSourceParameters) *RegisterProtectionSourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register protection source params
func (o *RegisterProtectionSourceParams) SetBody(body *models.RegisterProtectionSourceParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterProtectionSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
