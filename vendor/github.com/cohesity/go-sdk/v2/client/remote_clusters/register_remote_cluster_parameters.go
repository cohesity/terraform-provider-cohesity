// Code generated by go-swagger; DO NOT EDIT.

package remote_clusters

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

// NewRegisterRemoteClusterParams creates a new RegisterRemoteClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterRemoteClusterParams() *RegisterRemoteClusterParams {
	return &RegisterRemoteClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterRemoteClusterParamsWithTimeout creates a new RegisterRemoteClusterParams object
// with the ability to set a timeout on a request.
func NewRegisterRemoteClusterParamsWithTimeout(timeout time.Duration) *RegisterRemoteClusterParams {
	return &RegisterRemoteClusterParams{
		timeout: timeout,
	}
}

// NewRegisterRemoteClusterParamsWithContext creates a new RegisterRemoteClusterParams object
// with the ability to set a context for a request.
func NewRegisterRemoteClusterParamsWithContext(ctx context.Context) *RegisterRemoteClusterParams {
	return &RegisterRemoteClusterParams{
		Context: ctx,
	}
}

// NewRegisterRemoteClusterParamsWithHTTPClient creates a new RegisterRemoteClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterRemoteClusterParamsWithHTTPClient(client *http.Client) *RegisterRemoteClusterParams {
	return &RegisterRemoteClusterParams{
		HTTPClient: client,
	}
}

/*
RegisterRemoteClusterParams contains all the parameters to send to the API endpoint

	for the register remote cluster operation.

	Typically these are written to a http.Request.
*/
type RegisterRemoteClusterParams struct {

	/* Body.

	   Specifies the request to register Remote Cluster.
	*/
	Body *models.RegisterRemoteClusterParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register remote cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterRemoteClusterParams) WithDefaults() *RegisterRemoteClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register remote cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterRemoteClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register remote cluster params
func (o *RegisterRemoteClusterParams) WithTimeout(timeout time.Duration) *RegisterRemoteClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register remote cluster params
func (o *RegisterRemoteClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register remote cluster params
func (o *RegisterRemoteClusterParams) WithContext(ctx context.Context) *RegisterRemoteClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register remote cluster params
func (o *RegisterRemoteClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register remote cluster params
func (o *RegisterRemoteClusterParams) WithHTTPClient(client *http.Client) *RegisterRemoteClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register remote cluster params
func (o *RegisterRemoteClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register remote cluster params
func (o *RegisterRemoteClusterParams) WithBody(body *models.RegisterRemoteClusterParameters) *RegisterRemoteClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register remote cluster params
func (o *RegisterRemoteClusterParams) SetBody(body *models.RegisterRemoteClusterParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterRemoteClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
