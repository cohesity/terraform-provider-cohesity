// Code generated by go-swagger; DO NOT EDIT.

package antivirus_service

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

// NewGetIcapURIConnectionStatusParams creates a new GetIcapURIConnectionStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIcapURIConnectionStatusParams() *GetIcapURIConnectionStatusParams {
	return &GetIcapURIConnectionStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIcapURIConnectionStatusParamsWithTimeout creates a new GetIcapURIConnectionStatusParams object
// with the ability to set a timeout on a request.
func NewGetIcapURIConnectionStatusParamsWithTimeout(timeout time.Duration) *GetIcapURIConnectionStatusParams {
	return &GetIcapURIConnectionStatusParams{
		timeout: timeout,
	}
}

// NewGetIcapURIConnectionStatusParamsWithContext creates a new GetIcapURIConnectionStatusParams object
// with the ability to set a context for a request.
func NewGetIcapURIConnectionStatusParamsWithContext(ctx context.Context) *GetIcapURIConnectionStatusParams {
	return &GetIcapURIConnectionStatusParams{
		Context: ctx,
	}
}

// NewGetIcapURIConnectionStatusParamsWithHTTPClient creates a new GetIcapURIConnectionStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIcapURIConnectionStatusParamsWithHTTPClient(client *http.Client) *GetIcapURIConnectionStatusParams {
	return &GetIcapURIConnectionStatusParams{
		HTTPClient: client,
	}
}

/*
GetIcapURIConnectionStatusParams contains all the parameters to send to the API endpoint

	for the get icap Uri connection status operation.

	Typically these are written to a http.Request.
*/
type GetIcapURIConnectionStatusParams struct {

	/* Uris.

	   Specifies a list of URIs to check connection status.
	*/
	Uris []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get icap Uri connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIcapURIConnectionStatusParams) WithDefaults() *GetIcapURIConnectionStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get icap Uri connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIcapURIConnectionStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) WithTimeout(timeout time.Duration) *GetIcapURIConnectionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) WithContext(ctx context.Context) *GetIcapURIConnectionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) WithHTTPClient(client *http.Client) *GetIcapURIConnectionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUris adds the uris to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) WithUris(uris []string) *GetIcapURIConnectionStatusParams {
	o.SetUris(uris)
	return o
}

// SetUris adds the uris to the get icap Uri connection status params
func (o *GetIcapURIConnectionStatusParams) SetUris(uris []string) {
	o.Uris = uris
}

// WriteToRequest writes these params to a swagger request
func (o *GetIcapURIConnectionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Uris != nil {

		// binding items for uris
		joinedUris := o.bindParamUris(reg)

		// query array param uris
		if err := r.SetQueryParam("uris", joinedUris...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetIcapURIConnectionStatus binds the parameter uris
func (o *GetIcapURIConnectionStatusParams) bindParamUris(formats strfmt.Registry) []string {
	urisIR := o.Uris

	var urisIC []string
	for _, urisIIR := range urisIR { // explode []string

		urisIIV := urisIIR // string as string
		urisIC = append(urisIC, urisIIV)
	}

	// items.CollectionFormat: ""
	urisIS := swag.JoinByFormat(urisIC, "")

	return urisIS
}
