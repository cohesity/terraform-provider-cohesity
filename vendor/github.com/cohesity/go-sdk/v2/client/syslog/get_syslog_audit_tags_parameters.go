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
)

// NewGetSyslogAuditTagsParams creates a new GetSyslogAuditTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSyslogAuditTagsParams() *GetSyslogAuditTagsParams {
	return &GetSyslogAuditTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyslogAuditTagsParamsWithTimeout creates a new GetSyslogAuditTagsParams object
// with the ability to set a timeout on a request.
func NewGetSyslogAuditTagsParamsWithTimeout(timeout time.Duration) *GetSyslogAuditTagsParams {
	return &GetSyslogAuditTagsParams{
		timeout: timeout,
	}
}

// NewGetSyslogAuditTagsParamsWithContext creates a new GetSyslogAuditTagsParams object
// with the ability to set a context for a request.
func NewGetSyslogAuditTagsParamsWithContext(ctx context.Context) *GetSyslogAuditTagsParams {
	return &GetSyslogAuditTagsParams{
		Context: ctx,
	}
}

// NewGetSyslogAuditTagsParamsWithHTTPClient creates a new GetSyslogAuditTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSyslogAuditTagsParamsWithHTTPClient(client *http.Client) *GetSyslogAuditTagsParams {
	return &GetSyslogAuditTagsParams{
		HTTPClient: client,
	}
}

/*
GetSyslogAuditTagsParams contains all the parameters to send to the API endpoint

	for the get syslog audit tags operation.

	Typically these are written to a http.Request.
*/
type GetSyslogAuditTagsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get syslog audit tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyslogAuditTagsParams) WithDefaults() *GetSyslogAuditTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get syslog audit tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyslogAuditTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get syslog audit tags params
func (o *GetSyslogAuditTagsParams) WithTimeout(timeout time.Duration) *GetSyslogAuditTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get syslog audit tags params
func (o *GetSyslogAuditTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get syslog audit tags params
func (o *GetSyslogAuditTagsParams) WithContext(ctx context.Context) *GetSyslogAuditTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get syslog audit tags params
func (o *GetSyslogAuditTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get syslog audit tags params
func (o *GetSyslogAuditTagsParams) WithHTTPClient(client *http.Client) *GetSyslogAuditTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get syslog audit tags params
func (o *GetSyslogAuditTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyslogAuditTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
