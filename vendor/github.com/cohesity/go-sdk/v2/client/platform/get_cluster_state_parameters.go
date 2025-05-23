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
	"github.com/go-openapi/swag"
)

// NewGetClusterStateParams creates a new GetClusterStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterStateParams() *GetClusterStateParams {
	return &GetClusterStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterStateParamsWithTimeout creates a new GetClusterStateParams object
// with the ability to set a timeout on a request.
func NewGetClusterStateParamsWithTimeout(timeout time.Duration) *GetClusterStateParams {
	return &GetClusterStateParams{
		timeout: timeout,
	}
}

// NewGetClusterStateParamsWithContext creates a new GetClusterStateParams object
// with the ability to set a context for a request.
func NewGetClusterStateParamsWithContext(ctx context.Context) *GetClusterStateParams {
	return &GetClusterStateParams{
		Context: ctx,
	}
}

// NewGetClusterStateParamsWithHTTPClient creates a new GetClusterStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterStateParamsWithHTTPClient(client *http.Client) *GetClusterStateParams {
	return &GetClusterStateParams{
		HTTPClient: client,
	}
}

/*
GetClusterStateParams contains all the parameters to send to the API endpoint

	for the get cluster state operation.

	Typically these are written to a http.Request.
*/
type GetClusterStateParams struct {

	/* SystemApps.

	   The filter whether or not to get the system apps state details.
	*/
	SystemApps *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterStateParams) WithDefaults() *GetClusterStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster state params
func (o *GetClusterStateParams) WithTimeout(timeout time.Duration) *GetClusterStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster state params
func (o *GetClusterStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster state params
func (o *GetClusterStateParams) WithContext(ctx context.Context) *GetClusterStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster state params
func (o *GetClusterStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster state params
func (o *GetClusterStateParams) WithHTTPClient(client *http.Client) *GetClusterStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster state params
func (o *GetClusterStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSystemApps adds the systemApps to the get cluster state params
func (o *GetClusterStateParams) WithSystemApps(systemApps *bool) *GetClusterStateParams {
	o.SetSystemApps(systemApps)
	return o
}

// SetSystemApps adds the systemApps to the get cluster state params
func (o *GetClusterStateParams) SetSystemApps(systemApps *bool) {
	o.SystemApps = systemApps
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SystemApps != nil {

		// query param systemApps
		var qrSystemApps bool

		if o.SystemApps != nil {
			qrSystemApps = *o.SystemApps
		}
		qSystemApps := swag.FormatBool(qrSystemApps)
		if qSystemApps != "" {

			if err := r.SetQueryParam("systemApps", qSystemApps); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
