// Code generated by go-swagger; DO NOT EDIT.

package remote_restore

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

// NewListRemoteVaultSearchJobByIDParams creates a new ListRemoteVaultSearchJobByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRemoteVaultSearchJobByIDParams() *ListRemoteVaultSearchJobByIDParams {
	return &ListRemoteVaultSearchJobByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRemoteVaultSearchJobByIDParamsWithTimeout creates a new ListRemoteVaultSearchJobByIDParams object
// with the ability to set a timeout on a request.
func NewListRemoteVaultSearchJobByIDParamsWithTimeout(timeout time.Duration) *ListRemoteVaultSearchJobByIDParams {
	return &ListRemoteVaultSearchJobByIDParams{
		timeout: timeout,
	}
}

// NewListRemoteVaultSearchJobByIDParamsWithContext creates a new ListRemoteVaultSearchJobByIDParams object
// with the ability to set a context for a request.
func NewListRemoteVaultSearchJobByIDParamsWithContext(ctx context.Context) *ListRemoteVaultSearchJobByIDParams {
	return &ListRemoteVaultSearchJobByIDParams{
		Context: ctx,
	}
}

// NewListRemoteVaultSearchJobByIDParamsWithHTTPClient creates a new ListRemoteVaultSearchJobByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRemoteVaultSearchJobByIDParamsWithHTTPClient(client *http.Client) *ListRemoteVaultSearchJobByIDParams {
	return &ListRemoteVaultSearchJobByIDParams{
		HTTPClient: client,
	}
}

/*
ListRemoteVaultSearchJobByIDParams contains all the parameters to send to the API endpoint

	for the list remote vault search job by Id operation.

	Typically these are written to a http.Request.
*/
type ListRemoteVaultSearchJobByIDParams struct {

	/* ID.

	   Specifies the id of the remote Vault search Job to return.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list remote vault search job by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRemoteVaultSearchJobByIDParams) WithDefaults() *ListRemoteVaultSearchJobByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list remote vault search job by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRemoteVaultSearchJobByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) WithTimeout(timeout time.Duration) *ListRemoteVaultSearchJobByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) WithContext(ctx context.Context) *ListRemoteVaultSearchJobByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) WithHTTPClient(client *http.Client) *ListRemoteVaultSearchJobByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) WithID(id int64) *ListRemoteVaultSearchJobByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list remote vault search job by Id params
func (o *ListRemoteVaultSearchJobByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListRemoteVaultSearchJobByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
