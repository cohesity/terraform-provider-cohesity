// Code generated by go-swagger; DO NOT EDIT.

package object

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

// NewFilterObjectsParams creates a new FilterObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFilterObjectsParams() *FilterObjectsParams {
	return &FilterObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFilterObjectsParamsWithTimeout creates a new FilterObjectsParams object
// with the ability to set a timeout on a request.
func NewFilterObjectsParamsWithTimeout(timeout time.Duration) *FilterObjectsParams {
	return &FilterObjectsParams{
		timeout: timeout,
	}
}

// NewFilterObjectsParamsWithContext creates a new FilterObjectsParams object
// with the ability to set a context for a request.
func NewFilterObjectsParamsWithContext(ctx context.Context) *FilterObjectsParams {
	return &FilterObjectsParams{
		Context: ctx,
	}
}

// NewFilterObjectsParamsWithHTTPClient creates a new FilterObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFilterObjectsParamsWithHTTPClient(client *http.Client) *FilterObjectsParams {
	return &FilterObjectsParams{
		HTTPClient: client,
	}
}

/*
FilterObjectsParams contains all the parameters to send to the API endpoint

	for the filter objects operation.

	Typically these are written to a http.Request.
*/
type FilterObjectsParams struct {

	/* Body.

	   Specifies the parameters to filter objects.
	*/
	Body *models.FilterObjectsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the filter objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FilterObjectsParams) WithDefaults() *FilterObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the filter objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FilterObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the filter objects params
func (o *FilterObjectsParams) WithTimeout(timeout time.Duration) *FilterObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the filter objects params
func (o *FilterObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the filter objects params
func (o *FilterObjectsParams) WithContext(ctx context.Context) *FilterObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the filter objects params
func (o *FilterObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the filter objects params
func (o *FilterObjectsParams) WithHTTPClient(client *http.Client) *FilterObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the filter objects params
func (o *FilterObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the filter objects params
func (o *FilterObjectsParams) WithBody(body *models.FilterObjectsRequest) *FilterObjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the filter objects params
func (o *FilterObjectsParams) SetBody(body *models.FilterObjectsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FilterObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
