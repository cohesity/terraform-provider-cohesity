// Code generated by go-swagger; DO NOT EDIT.

package command

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

// NewLinuxCommandExecuteParams creates a new LinuxCommandExecuteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLinuxCommandExecuteParams() *LinuxCommandExecuteParams {
	return &LinuxCommandExecuteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLinuxCommandExecuteParamsWithTimeout creates a new LinuxCommandExecuteParams object
// with the ability to set a timeout on a request.
func NewLinuxCommandExecuteParamsWithTimeout(timeout time.Duration) *LinuxCommandExecuteParams {
	return &LinuxCommandExecuteParams{
		timeout: timeout,
	}
}

// NewLinuxCommandExecuteParamsWithContext creates a new LinuxCommandExecuteParams object
// with the ability to set a context for a request.
func NewLinuxCommandExecuteParamsWithContext(ctx context.Context) *LinuxCommandExecuteParams {
	return &LinuxCommandExecuteParams{
		Context: ctx,
	}
}

// NewLinuxCommandExecuteParamsWithHTTPClient creates a new LinuxCommandExecuteParams object
// with the ability to set a custom HTTPClient for a request.
func NewLinuxCommandExecuteParamsWithHTTPClient(client *http.Client) *LinuxCommandExecuteParams {
	return &LinuxCommandExecuteParams{
		HTTPClient: client,
	}
}

/*
LinuxCommandExecuteParams contains all the parameters to send to the API endpoint

	for the linux command execute operation.

	Typically these are written to a http.Request.
*/
type LinuxCommandExecuteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the linux command execute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LinuxCommandExecuteParams) WithDefaults() *LinuxCommandExecuteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the linux command execute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LinuxCommandExecuteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the linux command execute params
func (o *LinuxCommandExecuteParams) WithTimeout(timeout time.Duration) *LinuxCommandExecuteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the linux command execute params
func (o *LinuxCommandExecuteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the linux command execute params
func (o *LinuxCommandExecuteParams) WithContext(ctx context.Context) *LinuxCommandExecuteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the linux command execute params
func (o *LinuxCommandExecuteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the linux command execute params
func (o *LinuxCommandExecuteParams) WithHTTPClient(client *http.Client) *LinuxCommandExecuteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the linux command execute params
func (o *LinuxCommandExecuteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LinuxCommandExecuteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
