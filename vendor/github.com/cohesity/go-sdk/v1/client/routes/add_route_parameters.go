// Code generated by go-swagger; DO NOT EDIT.

package routes

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

// NewAddRouteParams creates a new AddRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRouteParams() *AddRouteParams {
	return &AddRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRouteParamsWithTimeout creates a new AddRouteParams object
// with the ability to set a timeout on a request.
func NewAddRouteParamsWithTimeout(timeout time.Duration) *AddRouteParams {
	return &AddRouteParams{
		timeout: timeout,
	}
}

// NewAddRouteParamsWithContext creates a new AddRouteParams object
// with the ability to set a context for a request.
func NewAddRouteParamsWithContext(ctx context.Context) *AddRouteParams {
	return &AddRouteParams{
		Context: ctx,
	}
}

// NewAddRouteParamsWithHTTPClient creates a new AddRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRouteParamsWithHTTPClient(client *http.Client) *AddRouteParams {
	return &AddRouteParams{
		HTTPClient: client,
	}
}

/*
AddRouteParams contains all the parameters to send to the API endpoint

	for the add route operation.

	Typically these are written to a http.Request.
*/
type AddRouteParams struct {

	// Body.
	Body *models.Route

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRouteParams) WithDefaults() *AddRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add route params
func (o *AddRouteParams) WithTimeout(timeout time.Duration) *AddRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add route params
func (o *AddRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add route params
func (o *AddRouteParams) WithContext(ctx context.Context) *AddRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add route params
func (o *AddRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add route params
func (o *AddRouteParams) WithHTTPClient(client *http.Client) *AddRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add route params
func (o *AddRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add route params
func (o *AddRouteParams) WithBody(body *models.Route) *AddRouteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add route params
func (o *AddRouteParams) SetBody(body *models.Route) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
