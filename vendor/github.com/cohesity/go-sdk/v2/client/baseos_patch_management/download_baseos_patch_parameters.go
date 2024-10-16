// Code generated by go-swagger; DO NOT EDIT.

package baseos_patch_management

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

// NewDownloadBaseosPatchParams creates a new DownloadBaseosPatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadBaseosPatchParams() *DownloadBaseosPatchParams {
	return &DownloadBaseosPatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadBaseosPatchParamsWithTimeout creates a new DownloadBaseosPatchParams object
// with the ability to set a timeout on a request.
func NewDownloadBaseosPatchParamsWithTimeout(timeout time.Duration) *DownloadBaseosPatchParams {
	return &DownloadBaseosPatchParams{
		timeout: timeout,
	}
}

// NewDownloadBaseosPatchParamsWithContext creates a new DownloadBaseosPatchParams object
// with the ability to set a context for a request.
func NewDownloadBaseosPatchParamsWithContext(ctx context.Context) *DownloadBaseosPatchParams {
	return &DownloadBaseosPatchParams{
		Context: ctx,
	}
}

// NewDownloadBaseosPatchParamsWithHTTPClient creates a new DownloadBaseosPatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadBaseosPatchParamsWithHTTPClient(client *http.Client) *DownloadBaseosPatchParams {
	return &DownloadBaseosPatchParams{
		HTTPClient: client,
	}
}

/*
DownloadBaseosPatchParams contains all the parameters to send to the API endpoint

	for the download baseos patch operation.

	Typically these are written to a http.Request.
*/
type DownloadBaseosPatchParams struct {

	/* Body.

	   Request to download a new baseos patch.
	*/
	Body *models.DownloadBaseosPatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download baseos patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadBaseosPatchParams) WithDefaults() *DownloadBaseosPatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download baseos patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadBaseosPatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download baseos patch params
func (o *DownloadBaseosPatchParams) WithTimeout(timeout time.Duration) *DownloadBaseosPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download baseos patch params
func (o *DownloadBaseosPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download baseos patch params
func (o *DownloadBaseosPatchParams) WithContext(ctx context.Context) *DownloadBaseosPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download baseos patch params
func (o *DownloadBaseosPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download baseos patch params
func (o *DownloadBaseosPatchParams) WithHTTPClient(client *http.Client) *DownloadBaseosPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download baseos patch params
func (o *DownloadBaseosPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the download baseos patch params
func (o *DownloadBaseosPatchParams) WithBody(body *models.DownloadBaseosPatchRequest) *DownloadBaseosPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the download baseos patch params
func (o *DownloadBaseosPatchParams) SetBody(body *models.DownloadBaseosPatchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadBaseosPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
