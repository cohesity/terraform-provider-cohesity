// Code generated by go-swagger; DO NOT EDIT.

package recovery

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

// NewTearDownRecoveryByIDParams creates a new TearDownRecoveryByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTearDownRecoveryByIDParams() *TearDownRecoveryByIDParams {
	return &TearDownRecoveryByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTearDownRecoveryByIDParamsWithTimeout creates a new TearDownRecoveryByIDParams object
// with the ability to set a timeout on a request.
func NewTearDownRecoveryByIDParamsWithTimeout(timeout time.Duration) *TearDownRecoveryByIDParams {
	return &TearDownRecoveryByIDParams{
		timeout: timeout,
	}
}

// NewTearDownRecoveryByIDParamsWithContext creates a new TearDownRecoveryByIDParams object
// with the ability to set a context for a request.
func NewTearDownRecoveryByIDParamsWithContext(ctx context.Context) *TearDownRecoveryByIDParams {
	return &TearDownRecoveryByIDParams{
		Context: ctx,
	}
}

// NewTearDownRecoveryByIDParamsWithHTTPClient creates a new TearDownRecoveryByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewTearDownRecoveryByIDParamsWithHTTPClient(client *http.Client) *TearDownRecoveryByIDParams {
	return &TearDownRecoveryByIDParams{
		HTTPClient: client,
	}
}

/*
TearDownRecoveryByIDParams contains all the parameters to send to the API endpoint

	for the tear down recovery by Id operation.

	Typically these are written to a http.Request.
*/
type TearDownRecoveryByIDParams struct {

	/* ID.

	   Specifies the id of a Recovery.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tear down recovery by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TearDownRecoveryByIDParams) WithDefaults() *TearDownRecoveryByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tear down recovery by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TearDownRecoveryByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) WithTimeout(timeout time.Duration) *TearDownRecoveryByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) WithContext(ctx context.Context) *TearDownRecoveryByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) WithHTTPClient(client *http.Client) *TearDownRecoveryByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) WithID(id string) *TearDownRecoveryByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tear down recovery by Id params
func (o *TearDownRecoveryByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TearDownRecoveryByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
