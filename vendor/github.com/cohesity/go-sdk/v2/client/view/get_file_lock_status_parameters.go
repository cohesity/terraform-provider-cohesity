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
)

// NewGetFileLockStatusParams creates a new GetFileLockStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileLockStatusParams() *GetFileLockStatusParams {
	return &GetFileLockStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileLockStatusParamsWithTimeout creates a new GetFileLockStatusParams object
// with the ability to set a timeout on a request.
func NewGetFileLockStatusParamsWithTimeout(timeout time.Duration) *GetFileLockStatusParams {
	return &GetFileLockStatusParams{
		timeout: timeout,
	}
}

// NewGetFileLockStatusParamsWithContext creates a new GetFileLockStatusParams object
// with the ability to set a context for a request.
func NewGetFileLockStatusParamsWithContext(ctx context.Context) *GetFileLockStatusParams {
	return &GetFileLockStatusParams{
		Context: ctx,
	}
}

// NewGetFileLockStatusParamsWithHTTPClient creates a new GetFileLockStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileLockStatusParamsWithHTTPClient(client *http.Client) *GetFileLockStatusParams {
	return &GetFileLockStatusParams{
		HTTPClient: client,
	}
}

/*
GetFileLockStatusParams contains all the parameters to send to the API endpoint

	for the get file lock status operation.

	Typically these are written to a http.Request.
*/
type GetFileLockStatusParams struct {

	/* ID.

	   Specifies the id of a view.

	   Format: int64
	*/
	ID int64

	/* Path.

	   Specifies the file path relative to root of the view.
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file lock status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileLockStatusParams) WithDefaults() *GetFileLockStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file lock status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileLockStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file lock status params
func (o *GetFileLockStatusParams) WithTimeout(timeout time.Duration) *GetFileLockStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file lock status params
func (o *GetFileLockStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file lock status params
func (o *GetFileLockStatusParams) WithContext(ctx context.Context) *GetFileLockStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file lock status params
func (o *GetFileLockStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file lock status params
func (o *GetFileLockStatusParams) WithHTTPClient(client *http.Client) *GetFileLockStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file lock status params
func (o *GetFileLockStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get file lock status params
func (o *GetFileLockStatusParams) WithID(id int64) *GetFileLockStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get file lock status params
func (o *GetFileLockStatusParams) SetID(id int64) {
	o.ID = id
}

// WithPath adds the path to the get file lock status params
func (o *GetFileLockStatusParams) WithPath(path string) *GetFileLockStatusParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get file lock status params
func (o *GetFileLockStatusParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileLockStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
