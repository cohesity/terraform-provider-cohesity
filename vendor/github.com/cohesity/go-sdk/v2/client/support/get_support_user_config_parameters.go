// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewGetSupportUserConfigParams creates a new GetSupportUserConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSupportUserConfigParams() *GetSupportUserConfigParams {
	return &GetSupportUserConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSupportUserConfigParamsWithTimeout creates a new GetSupportUserConfigParams object
// with the ability to set a timeout on a request.
func NewGetSupportUserConfigParamsWithTimeout(timeout time.Duration) *GetSupportUserConfigParams {
	return &GetSupportUserConfigParams{
		timeout: timeout,
	}
}

// NewGetSupportUserConfigParamsWithContext creates a new GetSupportUserConfigParams object
// with the ability to set a context for a request.
func NewGetSupportUserConfigParamsWithContext(ctx context.Context) *GetSupportUserConfigParams {
	return &GetSupportUserConfigParams{
		Context: ctx,
	}
}

// NewGetSupportUserConfigParamsWithHTTPClient creates a new GetSupportUserConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSupportUserConfigParamsWithHTTPClient(client *http.Client) *GetSupportUserConfigParams {
	return &GetSupportUserConfigParams{
		HTTPClient: client,
	}
}

/*
GetSupportUserConfigParams contains all the parameters to send to the API endpoint

	for the get support user config operation.

	Typically these are written to a http.Request.
*/
type GetSupportUserConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get support user config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSupportUserConfigParams) WithDefaults() *GetSupportUserConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get support user config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSupportUserConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get support user config params
func (o *GetSupportUserConfigParams) WithTimeout(timeout time.Duration) *GetSupportUserConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get support user config params
func (o *GetSupportUserConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get support user config params
func (o *GetSupportUserConfigParams) WithContext(ctx context.Context) *GetSupportUserConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get support user config params
func (o *GetSupportUserConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get support user config params
func (o *GetSupportUserConfigParams) WithHTTPClient(client *http.Client) *GetSupportUserConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get support user config params
func (o *GetSupportUserConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSupportUserConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
