// Code generated by go-swagger; DO NOT EDIT.

package data_source_connector_local

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

// NewGetDataSourceConnectorLogsParams creates a new GetDataSourceConnectorLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataSourceConnectorLogsParams() *GetDataSourceConnectorLogsParams {
	return &GetDataSourceConnectorLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataSourceConnectorLogsParamsWithTimeout creates a new GetDataSourceConnectorLogsParams object
// with the ability to set a timeout on a request.
func NewGetDataSourceConnectorLogsParamsWithTimeout(timeout time.Duration) *GetDataSourceConnectorLogsParams {
	return &GetDataSourceConnectorLogsParams{
		timeout: timeout,
	}
}

// NewGetDataSourceConnectorLogsParamsWithContext creates a new GetDataSourceConnectorLogsParams object
// with the ability to set a context for a request.
func NewGetDataSourceConnectorLogsParamsWithContext(ctx context.Context) *GetDataSourceConnectorLogsParams {
	return &GetDataSourceConnectorLogsParams{
		Context: ctx,
	}
}

// NewGetDataSourceConnectorLogsParamsWithHTTPClient creates a new GetDataSourceConnectorLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataSourceConnectorLogsParamsWithHTTPClient(client *http.Client) *GetDataSourceConnectorLogsParams {
	return &GetDataSourceConnectorLogsParams{
		HTTPClient: client,
	}
}

/*
GetDataSourceConnectorLogsParams contains all the parameters to send to the API endpoint

	for the get data source connector logs operation.

	Typically these are written to a http.Request.
*/
type GetDataSourceConnectorLogsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get data source connector logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataSourceConnectorLogsParams) WithDefaults() *GetDataSourceConnectorLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get data source connector logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataSourceConnectorLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get data source connector logs params
func (o *GetDataSourceConnectorLogsParams) WithTimeout(timeout time.Duration) *GetDataSourceConnectorLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data source connector logs params
func (o *GetDataSourceConnectorLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data source connector logs params
func (o *GetDataSourceConnectorLogsParams) WithContext(ctx context.Context) *GetDataSourceConnectorLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data source connector logs params
func (o *GetDataSourceConnectorLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data source connector logs params
func (o *GetDataSourceConnectorLogsParams) WithHTTPClient(client *http.Client) *GetDataSourceConnectorLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data source connector logs params
func (o *GetDataSourceConnectorLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataSourceConnectorLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
