// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewListServiceStatesParams creates a new ListServiceStatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServiceStatesParams() *ListServiceStatesParams {
	return &ListServiceStatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceStatesParamsWithTimeout creates a new ListServiceStatesParams object
// with the ability to set a timeout on a request.
func NewListServiceStatesParamsWithTimeout(timeout time.Duration) *ListServiceStatesParams {
	return &ListServiceStatesParams{
		timeout: timeout,
	}
}

// NewListServiceStatesParamsWithContext creates a new ListServiceStatesParams object
// with the ability to set a context for a request.
func NewListServiceStatesParamsWithContext(ctx context.Context) *ListServiceStatesParams {
	return &ListServiceStatesParams{
		Context: ctx,
	}
}

// NewListServiceStatesParamsWithHTTPClient creates a new ListServiceStatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServiceStatesParamsWithHTTPClient(client *http.Client) *ListServiceStatesParams {
	return &ListServiceStatesParams{
		HTTPClient: client,
	}
}

/*
ListServiceStatesParams contains all the parameters to send to the API endpoint

	for the list service states operation.

	Typically these are written to a http.Request.
*/
type ListServiceStatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list service states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceStatesParams) WithDefaults() *ListServiceStatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list service states params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceStatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list service states params
func (o *ListServiceStatesParams) WithTimeout(timeout time.Duration) *ListServiceStatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service states params
func (o *ListServiceStatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service states params
func (o *ListServiceStatesParams) WithContext(ctx context.Context) *ListServiceStatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service states params
func (o *ListServiceStatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service states params
func (o *ListServiceStatesParams) WithHTTPClient(client *http.Client) *ListServiceStatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service states params
func (o *ListServiceStatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceStatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
