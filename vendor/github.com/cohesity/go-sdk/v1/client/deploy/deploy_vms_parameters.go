// Code generated by go-swagger; DO NOT EDIT.

package deploy

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

// NewDeployVmsParams creates a new DeployVmsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeployVmsParams() *DeployVmsParams {
	return &DeployVmsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeployVmsParamsWithTimeout creates a new DeployVmsParams object
// with the ability to set a timeout on a request.
func NewDeployVmsParamsWithTimeout(timeout time.Duration) *DeployVmsParams {
	return &DeployVmsParams{
		timeout: timeout,
	}
}

// NewDeployVmsParamsWithContext creates a new DeployVmsParams object
// with the ability to set a context for a request.
func NewDeployVmsParamsWithContext(ctx context.Context) *DeployVmsParams {
	return &DeployVmsParams{
		Context: ctx,
	}
}

// NewDeployVmsParamsWithHTTPClient creates a new DeployVmsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeployVmsParamsWithHTTPClient(client *http.Client) *DeployVmsParams {
	return &DeployVmsParams{
		HTTPClient: client,
	}
}

/*
DeployVmsParams contains all the parameters to send to the API endpoint

	for the deploy vms operation.

	Typically these are written to a http.Request.
*/
type DeployVmsParams struct {

	/* Body.

	   Deploy VM argument
	*/
	Body *models.DeployArg

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deploy vms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeployVmsParams) WithDefaults() *DeployVmsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deploy vms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeployVmsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deploy vms params
func (o *DeployVmsParams) WithTimeout(timeout time.Duration) *DeployVmsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deploy vms params
func (o *DeployVmsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deploy vms params
func (o *DeployVmsParams) WithContext(ctx context.Context) *DeployVmsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deploy vms params
func (o *DeployVmsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deploy vms params
func (o *DeployVmsParams) WithHTTPClient(client *http.Client) *DeployVmsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deploy vms params
func (o *DeployVmsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the deploy vms params
func (o *DeployVmsParams) WithBody(body *models.DeployArg) *DeployVmsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the deploy vms params
func (o *DeployVmsParams) SetBody(body *models.DeployArg) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeployVmsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
