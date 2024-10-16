// Code generated by go-swagger; DO NOT EDIT.

package remote_restore

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

// NewStopRemoteVaultSearchJobParams creates a new StopRemoteVaultSearchJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopRemoteVaultSearchJobParams() *StopRemoteVaultSearchJobParams {
	return &StopRemoteVaultSearchJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopRemoteVaultSearchJobParamsWithTimeout creates a new StopRemoteVaultSearchJobParams object
// with the ability to set a timeout on a request.
func NewStopRemoteVaultSearchJobParamsWithTimeout(timeout time.Duration) *StopRemoteVaultSearchJobParams {
	return &StopRemoteVaultSearchJobParams{
		timeout: timeout,
	}
}

// NewStopRemoteVaultSearchJobParamsWithContext creates a new StopRemoteVaultSearchJobParams object
// with the ability to set a context for a request.
func NewStopRemoteVaultSearchJobParamsWithContext(ctx context.Context) *StopRemoteVaultSearchJobParams {
	return &StopRemoteVaultSearchJobParams{
		Context: ctx,
	}
}

// NewStopRemoteVaultSearchJobParamsWithHTTPClient creates a new StopRemoteVaultSearchJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopRemoteVaultSearchJobParamsWithHTTPClient(client *http.Client) *StopRemoteVaultSearchJobParams {
	return &StopRemoteVaultSearchJobParams{
		HTTPClient: client,
	}
}

/*
StopRemoteVaultSearchJobParams contains all the parameters to send to the API endpoint

	for the stop remote vault search job operation.

	Typically these are written to a http.Request.
*/
type StopRemoteVaultSearchJobParams struct {

	/* Body.

	   Request to stop a Remote Vault Search Job.
	*/
	Body *models.StopRemoteVaultSearchJobParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop remote vault search job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopRemoteVaultSearchJobParams) WithDefaults() *StopRemoteVaultSearchJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop remote vault search job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopRemoteVaultSearchJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) WithTimeout(timeout time.Duration) *StopRemoteVaultSearchJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) WithContext(ctx context.Context) *StopRemoteVaultSearchJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) WithHTTPClient(client *http.Client) *StopRemoteVaultSearchJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) WithBody(body *models.StopRemoteVaultSearchJobParameters) *StopRemoteVaultSearchJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stop remote vault search job params
func (o *StopRemoteVaultSearchJobParams) SetBody(body *models.StopRemoteVaultSearchJobParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StopRemoteVaultSearchJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
