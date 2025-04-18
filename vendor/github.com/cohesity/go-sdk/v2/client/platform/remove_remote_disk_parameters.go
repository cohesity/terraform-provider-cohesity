// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// NewRemoveRemoteDiskParams creates a new RemoveRemoteDiskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveRemoteDiskParams() *RemoveRemoteDiskParams {
	return &RemoveRemoteDiskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveRemoteDiskParamsWithTimeout creates a new RemoveRemoteDiskParams object
// with the ability to set a timeout on a request.
func NewRemoveRemoteDiskParamsWithTimeout(timeout time.Duration) *RemoveRemoteDiskParams {
	return &RemoveRemoteDiskParams{
		timeout: timeout,
	}
}

// NewRemoveRemoteDiskParamsWithContext creates a new RemoveRemoteDiskParams object
// with the ability to set a context for a request.
func NewRemoveRemoteDiskParamsWithContext(ctx context.Context) *RemoveRemoteDiskParams {
	return &RemoveRemoteDiskParams{
		Context: ctx,
	}
}

// NewRemoveRemoteDiskParamsWithHTTPClient creates a new RemoveRemoteDiskParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveRemoteDiskParamsWithHTTPClient(client *http.Client) *RemoveRemoteDiskParams {
	return &RemoveRemoteDiskParams{
		HTTPClient: client,
	}
}

/*
RemoveRemoteDiskParams contains all the parameters to send to the API endpoint

	for the remove remote disk operation.

	Typically these are written to a http.Request.
*/
type RemoveRemoteDiskParams struct {

	/* ID.

	   Specifies the id of the remote disk to remove.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove remote disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveRemoteDiskParams) WithDefaults() *RemoveRemoteDiskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove remote disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveRemoteDiskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove remote disk params
func (o *RemoveRemoteDiskParams) WithTimeout(timeout time.Duration) *RemoveRemoteDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove remote disk params
func (o *RemoveRemoteDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove remote disk params
func (o *RemoveRemoteDiskParams) WithContext(ctx context.Context) *RemoveRemoteDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove remote disk params
func (o *RemoveRemoteDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove remote disk params
func (o *RemoveRemoteDiskParams) WithHTTPClient(client *http.Client) *RemoveRemoteDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove remote disk params
func (o *RemoveRemoteDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove remote disk params
func (o *RemoveRemoteDiskParams) WithID(id int64) *RemoveRemoteDiskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove remote disk params
func (o *RemoveRemoteDiskParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveRemoteDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
