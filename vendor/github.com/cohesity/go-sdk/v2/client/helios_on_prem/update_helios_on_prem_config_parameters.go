// Code generated by go-swagger; DO NOT EDIT.

package helios_on_prem

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

// NewUpdateHeliosOnPremConfigParams creates a new UpdateHeliosOnPremConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHeliosOnPremConfigParams() *UpdateHeliosOnPremConfigParams {
	return &UpdateHeliosOnPremConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHeliosOnPremConfigParamsWithTimeout creates a new UpdateHeliosOnPremConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateHeliosOnPremConfigParamsWithTimeout(timeout time.Duration) *UpdateHeliosOnPremConfigParams {
	return &UpdateHeliosOnPremConfigParams{
		timeout: timeout,
	}
}

// NewUpdateHeliosOnPremConfigParamsWithContext creates a new UpdateHeliosOnPremConfigParams object
// with the ability to set a context for a request.
func NewUpdateHeliosOnPremConfigParamsWithContext(ctx context.Context) *UpdateHeliosOnPremConfigParams {
	return &UpdateHeliosOnPremConfigParams{
		Context: ctx,
	}
}

// NewUpdateHeliosOnPremConfigParamsWithHTTPClient creates a new UpdateHeliosOnPremConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHeliosOnPremConfigParamsWithHTTPClient(client *http.Client) *UpdateHeliosOnPremConfigParams {
	return &UpdateHeliosOnPremConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateHeliosOnPremConfigParams contains all the parameters to send to the API endpoint

	for the update helios on prem config operation.

	Typically these are written to a http.Request.
*/
type UpdateHeliosOnPremConfigParams struct {

	/* Body.

	   Specifies the parameters for config update.
	*/
	Body *models.HeliosOnPremConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update helios on prem config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHeliosOnPremConfigParams) WithDefaults() *UpdateHeliosOnPremConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update helios on prem config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHeliosOnPremConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) WithTimeout(timeout time.Duration) *UpdateHeliosOnPremConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) WithContext(ctx context.Context) *UpdateHeliosOnPremConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) WithHTTPClient(client *http.Client) *UpdateHeliosOnPremConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) WithBody(body *models.HeliosOnPremConfig) *UpdateHeliosOnPremConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update helios on prem config params
func (o *UpdateHeliosOnPremConfigParams) SetBody(body *models.HeliosOnPremConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHeliosOnPremConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
