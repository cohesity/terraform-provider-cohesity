// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// NewUpdateHostsParams creates a new UpdateHostsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHostsParams() *UpdateHostsParams {
	return &UpdateHostsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHostsParamsWithTimeout creates a new UpdateHostsParams object
// with the ability to set a timeout on a request.
func NewUpdateHostsParamsWithTimeout(timeout time.Duration) *UpdateHostsParams {
	return &UpdateHostsParams{
		timeout: timeout,
	}
}

// NewUpdateHostsParamsWithContext creates a new UpdateHostsParams object
// with the ability to set a context for a request.
func NewUpdateHostsParamsWithContext(ctx context.Context) *UpdateHostsParams {
	return &UpdateHostsParams{
		Context: ctx,
	}
}

// NewUpdateHostsParamsWithHTTPClient creates a new UpdateHostsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHostsParamsWithHTTPClient(client *http.Client) *UpdateHostsParams {
	return &UpdateHostsParams{
		HTTPClient: client,
	}
}

/*
UpdateHostsParams contains all the parameters to send to the API endpoint

	for the update hosts operation.

	Typically these are written to a http.Request.
*/
type UpdateHostsParams struct {

	// Body.
	Body models.HostMappingsParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHostsParams) WithDefaults() *UpdateHostsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHostsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update hosts params
func (o *UpdateHostsParams) WithTimeout(timeout time.Duration) *UpdateHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update hosts params
func (o *UpdateHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update hosts params
func (o *UpdateHostsParams) WithContext(ctx context.Context) *UpdateHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update hosts params
func (o *UpdateHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update hosts params
func (o *UpdateHostsParams) WithHTTPClient(client *http.Client) *UpdateHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update hosts params
func (o *UpdateHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update hosts params
func (o *UpdateHostsParams) WithBody(body models.HostMappingsParameters) *UpdateHostsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update hosts params
func (o *UpdateHostsParams) SetBody(body models.HostMappingsParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
