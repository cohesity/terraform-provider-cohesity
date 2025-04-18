// Code generated by go-swagger; DO NOT EDIT.

package view

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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewLockFileParams creates a new LockFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLockFileParams() *LockFileParams {
	return &LockFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLockFileParamsWithTimeout creates a new LockFileParams object
// with the ability to set a timeout on a request.
func NewLockFileParamsWithTimeout(timeout time.Duration) *LockFileParams {
	return &LockFileParams{
		timeout: timeout,
	}
}

// NewLockFileParamsWithContext creates a new LockFileParams object
// with the ability to set a context for a request.
func NewLockFileParamsWithContext(ctx context.Context) *LockFileParams {
	return &LockFileParams{
		Context: ctx,
	}
}

// NewLockFileParamsWithHTTPClient creates a new LockFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewLockFileParamsWithHTTPClient(client *http.Client) *LockFileParams {
	return &LockFileParams{
		HTTPClient: client,
	}
}

/*
LockFileParams contains all the parameters to send to the API endpoint

	for the lock file operation.

	Typically these are written to a http.Request.
*/
type LockFileParams struct {

	/* Body.

	   Specifies the request params to lock a file
	*/
	Body *models.LockFileParams

	/* ID.

	   Specifies the id of a view.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lock file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockFileParams) WithDefaults() *LockFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lock file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lock file params
func (o *LockFileParams) WithTimeout(timeout time.Duration) *LockFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lock file params
func (o *LockFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lock file params
func (o *LockFileParams) WithContext(ctx context.Context) *LockFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lock file params
func (o *LockFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lock file params
func (o *LockFileParams) WithHTTPClient(client *http.Client) *LockFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lock file params
func (o *LockFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lock file params
func (o *LockFileParams) WithBody(body *models.LockFileParams) *LockFileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lock file params
func (o *LockFileParams) SetBody(body *models.LockFileParams) {
	o.Body = body
}

// WithID adds the id to the lock file params
func (o *LockFileParams) WithID(id int64) *LockFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the lock file params
func (o *LockFileParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LockFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
