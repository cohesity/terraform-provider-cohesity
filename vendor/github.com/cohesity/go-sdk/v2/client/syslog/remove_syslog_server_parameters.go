// Code generated by go-swagger; DO NOT EDIT.

package syslog

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

// NewRemoveSyslogServerParams creates a new RemoveSyslogServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveSyslogServerParams() *RemoveSyslogServerParams {
	return &RemoveSyslogServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveSyslogServerParamsWithTimeout creates a new RemoveSyslogServerParams object
// with the ability to set a timeout on a request.
func NewRemoveSyslogServerParamsWithTimeout(timeout time.Duration) *RemoveSyslogServerParams {
	return &RemoveSyslogServerParams{
		timeout: timeout,
	}
}

// NewRemoveSyslogServerParamsWithContext creates a new RemoveSyslogServerParams object
// with the ability to set a context for a request.
func NewRemoveSyslogServerParamsWithContext(ctx context.Context) *RemoveSyslogServerParams {
	return &RemoveSyslogServerParams{
		Context: ctx,
	}
}

// NewRemoveSyslogServerParamsWithHTTPClient creates a new RemoveSyslogServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveSyslogServerParamsWithHTTPClient(client *http.Client) *RemoveSyslogServerParams {
	return &RemoveSyslogServerParams{
		HTTPClient: client,
	}
}

/*
RemoveSyslogServerParams contains all the parameters to send to the API endpoint

	for the remove syslog server operation.

	Typically these are written to a http.Request.
*/
type RemoveSyslogServerParams struct {

	/* ID.

	   Specifies a unique id of the syslog server.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove syslog server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveSyslogServerParams) WithDefaults() *RemoveSyslogServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove syslog server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveSyslogServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove syslog server params
func (o *RemoveSyslogServerParams) WithTimeout(timeout time.Duration) *RemoveSyslogServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove syslog server params
func (o *RemoveSyslogServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove syslog server params
func (o *RemoveSyslogServerParams) WithContext(ctx context.Context) *RemoveSyslogServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove syslog server params
func (o *RemoveSyslogServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove syslog server params
func (o *RemoveSyslogServerParams) WithHTTPClient(client *http.Client) *RemoveSyslogServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove syslog server params
func (o *RemoveSyslogServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove syslog server params
func (o *RemoveSyslogServerParams) WithID(id int32) *RemoveSyslogServerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove syslog server params
func (o *RemoveSyslogServerParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveSyslogServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
