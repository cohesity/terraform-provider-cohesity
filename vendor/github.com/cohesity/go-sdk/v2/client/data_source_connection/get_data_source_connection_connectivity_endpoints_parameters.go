// Code generated by go-swagger; DO NOT EDIT.

package data_source_connection

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

// NewGetDataSourceConnectionConnectivityEndpointsParams creates a new GetDataSourceConnectionConnectivityEndpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataSourceConnectionConnectivityEndpointsParams() *GetDataSourceConnectionConnectivityEndpointsParams {
	return &GetDataSourceConnectionConnectivityEndpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataSourceConnectionConnectivityEndpointsParamsWithTimeout creates a new GetDataSourceConnectionConnectivityEndpointsParams object
// with the ability to set a timeout on a request.
func NewGetDataSourceConnectionConnectivityEndpointsParamsWithTimeout(timeout time.Duration) *GetDataSourceConnectionConnectivityEndpointsParams {
	return &GetDataSourceConnectionConnectivityEndpointsParams{
		timeout: timeout,
	}
}

// NewGetDataSourceConnectionConnectivityEndpointsParamsWithContext creates a new GetDataSourceConnectionConnectivityEndpointsParams object
// with the ability to set a context for a request.
func NewGetDataSourceConnectionConnectivityEndpointsParamsWithContext(ctx context.Context) *GetDataSourceConnectionConnectivityEndpointsParams {
	return &GetDataSourceConnectionConnectivityEndpointsParams{
		Context: ctx,
	}
}

// NewGetDataSourceConnectionConnectivityEndpointsParamsWithHTTPClient creates a new GetDataSourceConnectionConnectivityEndpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataSourceConnectionConnectivityEndpointsParamsWithHTTPClient(client *http.Client) *GetDataSourceConnectionConnectivityEndpointsParams {
	return &GetDataSourceConnectionConnectivityEndpointsParams{
		HTTPClient: client,
	}
}

/*
GetDataSourceConnectionConnectivityEndpointsParams contains all the parameters to send to the API endpoint

	for the get data source connection connectivity endpoints operation.

	Typically these are written to a http.Request.
*/
type GetDataSourceConnectionConnectivityEndpointsParams struct {

	/* ConnectionID.

	   Specifies the unique ID of the connection for which connectivity endpoints are to be fetched.
	*/
	ConnectionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get data source connection connectivity endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataSourceConnectionConnectivityEndpointsParams) WithDefaults() *GetDataSourceConnectionConnectivityEndpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get data source connection connectivity endpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataSourceConnectionConnectivityEndpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) WithTimeout(timeout time.Duration) *GetDataSourceConnectionConnectivityEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) WithContext(ctx context.Context) *GetDataSourceConnectionConnectivityEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) WithHTTPClient(client *http.Client) *GetDataSourceConnectionConnectivityEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) WithConnectionID(connectionID *string) *GetDataSourceConnectionConnectivityEndpointsParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get data source connection connectivity endpoints params
func (o *GetDataSourceConnectionConnectivityEndpointsParams) SetConnectionID(connectionID *string) {
	o.ConnectionID = connectionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataSourceConnectionConnectivityEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConnectionID != nil {

		// query param connectionId
		var qrConnectionID string

		if o.ConnectionID != nil {
			qrConnectionID = *o.ConnectionID
		}
		qConnectionID := qrConnectionID
		if qConnectionID != "" {

			if err := r.SetQueryParam("connectionId", qConnectionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
