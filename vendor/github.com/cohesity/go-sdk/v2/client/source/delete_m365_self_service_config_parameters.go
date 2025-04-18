// Code generated by go-swagger; DO NOT EDIT.

package source

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

// NewDeleteM365SelfServiceConfigParams creates a new DeleteM365SelfServiceConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteM365SelfServiceConfigParams() *DeleteM365SelfServiceConfigParams {
	return &DeleteM365SelfServiceConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteM365SelfServiceConfigParamsWithTimeout creates a new DeleteM365SelfServiceConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteM365SelfServiceConfigParamsWithTimeout(timeout time.Duration) *DeleteM365SelfServiceConfigParams {
	return &DeleteM365SelfServiceConfigParams{
		timeout: timeout,
	}
}

// NewDeleteM365SelfServiceConfigParamsWithContext creates a new DeleteM365SelfServiceConfigParams object
// with the ability to set a context for a request.
func NewDeleteM365SelfServiceConfigParamsWithContext(ctx context.Context) *DeleteM365SelfServiceConfigParams {
	return &DeleteM365SelfServiceConfigParams{
		Context: ctx,
	}
}

// NewDeleteM365SelfServiceConfigParamsWithHTTPClient creates a new DeleteM365SelfServiceConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteM365SelfServiceConfigParamsWithHTTPClient(client *http.Client) *DeleteM365SelfServiceConfigParams {
	return &DeleteM365SelfServiceConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteM365SelfServiceConfigParams contains all the parameters to send to the API endpoint

	for the delete m365 self service config operation.

	Typically these are written to a http.Request.
*/
type DeleteM365SelfServiceConfigParams struct {

	/* UUID.

	   Specifies the UUID of the Microsoft365 Source.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete m365 self service config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteM365SelfServiceConfigParams) WithDefaults() *DeleteM365SelfServiceConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete m365 self service config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteM365SelfServiceConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) WithTimeout(timeout time.Duration) *DeleteM365SelfServiceConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) WithContext(ctx context.Context) *DeleteM365SelfServiceConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) WithHTTPClient(client *http.Client) *DeleteM365SelfServiceConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) WithUUID(uuid string) *DeleteM365SelfServiceConfigParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete m365 self service config params
func (o *DeleteM365SelfServiceConfigParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteM365SelfServiceConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
