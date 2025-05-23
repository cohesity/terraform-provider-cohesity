// Code generated by go-swagger; DO NOT EDIT.

package key_management_system

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

// NewDeleteKmsConfigParams creates a new DeleteKmsConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKmsConfigParams() *DeleteKmsConfigParams {
	return &DeleteKmsConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKmsConfigParamsWithTimeout creates a new DeleteKmsConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteKmsConfigParamsWithTimeout(timeout time.Duration) *DeleteKmsConfigParams {
	return &DeleteKmsConfigParams{
		timeout: timeout,
	}
}

// NewDeleteKmsConfigParamsWithContext creates a new DeleteKmsConfigParams object
// with the ability to set a context for a request.
func NewDeleteKmsConfigParamsWithContext(ctx context.Context) *DeleteKmsConfigParams {
	return &DeleteKmsConfigParams{
		Context: ctx,
	}
}

// NewDeleteKmsConfigParamsWithHTTPClient creates a new DeleteKmsConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKmsConfigParamsWithHTTPClient(client *http.Client) *DeleteKmsConfigParams {
	return &DeleteKmsConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteKmsConfigParams contains all the parameters to send to the API endpoint

	for the delete kms config operation.

	Typically these are written to a http.Request.
*/
type DeleteKmsConfigParams struct {

	/* ID.

	   ID of the KMS configured on the cluster.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete kms config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKmsConfigParams) WithDefaults() *DeleteKmsConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete kms config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKmsConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete kms config params
func (o *DeleteKmsConfigParams) WithTimeout(timeout time.Duration) *DeleteKmsConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete kms config params
func (o *DeleteKmsConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete kms config params
func (o *DeleteKmsConfigParams) WithContext(ctx context.Context) *DeleteKmsConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete kms config params
func (o *DeleteKmsConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete kms config params
func (o *DeleteKmsConfigParams) WithHTTPClient(client *http.Client) *DeleteKmsConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete kms config params
func (o *DeleteKmsConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete kms config params
func (o *DeleteKmsConfigParams) WithID(id int64) *DeleteKmsConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete kms config params
func (o *DeleteKmsConfigParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKmsConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
