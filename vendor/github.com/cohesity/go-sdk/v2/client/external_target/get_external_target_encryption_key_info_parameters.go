// Code generated by go-swagger; DO NOT EDIT.

package external_target

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

// NewGetExternalTargetEncryptionKeyInfoParams creates a new GetExternalTargetEncryptionKeyInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalTargetEncryptionKeyInfoParams() *GetExternalTargetEncryptionKeyInfoParams {
	return &GetExternalTargetEncryptionKeyInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalTargetEncryptionKeyInfoParamsWithTimeout creates a new GetExternalTargetEncryptionKeyInfoParams object
// with the ability to set a timeout on a request.
func NewGetExternalTargetEncryptionKeyInfoParamsWithTimeout(timeout time.Duration) *GetExternalTargetEncryptionKeyInfoParams {
	return &GetExternalTargetEncryptionKeyInfoParams{
		timeout: timeout,
	}
}

// NewGetExternalTargetEncryptionKeyInfoParamsWithContext creates a new GetExternalTargetEncryptionKeyInfoParams object
// with the ability to set a context for a request.
func NewGetExternalTargetEncryptionKeyInfoParamsWithContext(ctx context.Context) *GetExternalTargetEncryptionKeyInfoParams {
	return &GetExternalTargetEncryptionKeyInfoParams{
		Context: ctx,
	}
}

// NewGetExternalTargetEncryptionKeyInfoParamsWithHTTPClient creates a new GetExternalTargetEncryptionKeyInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalTargetEncryptionKeyInfoParamsWithHTTPClient(client *http.Client) *GetExternalTargetEncryptionKeyInfoParams {
	return &GetExternalTargetEncryptionKeyInfoParams{
		HTTPClient: client,
	}
}

/*
GetExternalTargetEncryptionKeyInfoParams contains all the parameters to send to the API endpoint

	for the get external target encryption key info operation.

	Typically these are written to a http.Request.
*/
type GetExternalTargetEncryptionKeyInfoParams struct {

	/* ID.

	   Specifies the id of the External Target.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get external target encryption key info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalTargetEncryptionKeyInfoParams) WithDefaults() *GetExternalTargetEncryptionKeyInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get external target encryption key info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalTargetEncryptionKeyInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) WithTimeout(timeout time.Duration) *GetExternalTargetEncryptionKeyInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) WithContext(ctx context.Context) *GetExternalTargetEncryptionKeyInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) WithHTTPClient(client *http.Client) *GetExternalTargetEncryptionKeyInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) WithID(id int64) *GetExternalTargetEncryptionKeyInfoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get external target encryption key info params
func (o *GetExternalTargetEncryptionKeyInfoParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalTargetEncryptionKeyInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
