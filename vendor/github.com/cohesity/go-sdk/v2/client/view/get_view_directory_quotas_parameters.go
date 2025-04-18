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

// NewGetViewDirectoryQuotasParams creates a new GetViewDirectoryQuotasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetViewDirectoryQuotasParams() *GetViewDirectoryQuotasParams {
	return &GetViewDirectoryQuotasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetViewDirectoryQuotasParamsWithTimeout creates a new GetViewDirectoryQuotasParams object
// with the ability to set a timeout on a request.
func NewGetViewDirectoryQuotasParamsWithTimeout(timeout time.Duration) *GetViewDirectoryQuotasParams {
	return &GetViewDirectoryQuotasParams{
		timeout: timeout,
	}
}

// NewGetViewDirectoryQuotasParamsWithContext creates a new GetViewDirectoryQuotasParams object
// with the ability to set a context for a request.
func NewGetViewDirectoryQuotasParamsWithContext(ctx context.Context) *GetViewDirectoryQuotasParams {
	return &GetViewDirectoryQuotasParams{
		Context: ctx,
	}
}

// NewGetViewDirectoryQuotasParamsWithHTTPClient creates a new GetViewDirectoryQuotasParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetViewDirectoryQuotasParamsWithHTTPClient(client *http.Client) *GetViewDirectoryQuotasParams {
	return &GetViewDirectoryQuotasParams{
		HTTPClient: client,
	}
}

/*
GetViewDirectoryQuotasParams contains all the parameters to send to the API endpoint

	for the get view directory quotas operation.

	Typically these are written to a http.Request.
*/
type GetViewDirectoryQuotasParams struct {

	/* Cookie.

	   Specifies the cookie.

	   Format: int64
	*/
	Cookie *int64

	/* ID.

	   Specifies the View id.

	   Format: int64
	*/
	ID int64

	/* MaxCount.

	   Specifies a limit on the number of quotas returned.

	   Format: int64
	*/
	MaxCount *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get view directory quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewDirectoryQuotasParams) WithDefaults() *GetViewDirectoryQuotasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get view directory quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewDirectoryQuotasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) WithTimeout(timeout time.Duration) *GetViewDirectoryQuotasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) WithContext(ctx context.Context) *GetViewDirectoryQuotasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) WithHTTPClient(client *http.Client) *GetViewDirectoryQuotasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCookie adds the cookie to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) WithCookie(cookie *int64) *GetViewDirectoryQuotasParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) SetCookie(cookie *int64) {
	o.Cookie = cookie
}

// WithID adds the id to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) WithID(id int64) *GetViewDirectoryQuotasParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) SetID(id int64) {
	o.ID = id
}

// WithMaxCount adds the maxCount to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) WithMaxCount(maxCount *int64) *GetViewDirectoryQuotasParams {
	o.SetMaxCount(maxCount)
	return o
}

// SetMaxCount adds the maxCount to the get view directory quotas params
func (o *GetViewDirectoryQuotasParams) SetMaxCount(maxCount *int64) {
	o.MaxCount = maxCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetViewDirectoryQuotasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cookie != nil {

		// query param cookie
		var qrCookie int64

		if o.Cookie != nil {
			qrCookie = *o.Cookie
		}
		qCookie := swag.FormatInt64(qrCookie)
		if qCookie != "" {

			if err := r.SetQueryParam("cookie", qCookie); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.MaxCount != nil {

		// query param maxCount
		var qrMaxCount int64

		if o.MaxCount != nil {
			qrMaxCount = *o.MaxCount
		}
		qMaxCount := swag.FormatInt64(qrMaxCount)
		if qMaxCount != "" {

			if err := r.SetQueryParam("maxCount", qMaxCount); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
