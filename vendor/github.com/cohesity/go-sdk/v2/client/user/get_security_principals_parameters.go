// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetSecurityPrincipalsParams creates a new GetSecurityPrincipalsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSecurityPrincipalsParams() *GetSecurityPrincipalsParams {
	return &GetSecurityPrincipalsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityPrincipalsParamsWithTimeout creates a new GetSecurityPrincipalsParams object
// with the ability to set a timeout on a request.
func NewGetSecurityPrincipalsParamsWithTimeout(timeout time.Duration) *GetSecurityPrincipalsParams {
	return &GetSecurityPrincipalsParams{
		timeout: timeout,
	}
}

// NewGetSecurityPrincipalsParamsWithContext creates a new GetSecurityPrincipalsParams object
// with the ability to set a context for a request.
func NewGetSecurityPrincipalsParamsWithContext(ctx context.Context) *GetSecurityPrincipalsParams {
	return &GetSecurityPrincipalsParams{
		Context: ctx,
	}
}

// NewGetSecurityPrincipalsParamsWithHTTPClient creates a new GetSecurityPrincipalsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSecurityPrincipalsParamsWithHTTPClient(client *http.Client) *GetSecurityPrincipalsParams {
	return &GetSecurityPrincipalsParams{
		HTTPClient: client,
	}
}

/*
GetSecurityPrincipalsParams contains all the parameters to send to the API endpoint

	for the get security principals operation.

	Typically these are written to a http.Request.
*/
type GetSecurityPrincipalsParams struct {

	/* Sids.

	   Specifies a list of SIDs.
	*/
	Sids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get security principals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityPrincipalsParams) WithDefaults() *GetSecurityPrincipalsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get security principals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityPrincipalsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get security principals params
func (o *GetSecurityPrincipalsParams) WithTimeout(timeout time.Duration) *GetSecurityPrincipalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security principals params
func (o *GetSecurityPrincipalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security principals params
func (o *GetSecurityPrincipalsParams) WithContext(ctx context.Context) *GetSecurityPrincipalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security principals params
func (o *GetSecurityPrincipalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security principals params
func (o *GetSecurityPrincipalsParams) WithHTTPClient(client *http.Client) *GetSecurityPrincipalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security principals params
func (o *GetSecurityPrincipalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSids adds the sids to the get security principals params
func (o *GetSecurityPrincipalsParams) WithSids(sids []string) *GetSecurityPrincipalsParams {
	o.SetSids(sids)
	return o
}

// SetSids adds the sids to the get security principals params
func (o *GetSecurityPrincipalsParams) SetSids(sids []string) {
	o.Sids = sids
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityPrincipalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Sids != nil {

		// binding items for sids
		joinedSids := o.bindParamSids(reg)

		// query array param sids
		if err := r.SetQueryParam("sids", joinedSids...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetSecurityPrincipals binds the parameter sids
func (o *GetSecurityPrincipalsParams) bindParamSids(formats strfmt.Registry) []string {
	sidsIR := o.Sids

	var sidsIC []string
	for _, sidsIIR := range sidsIR { // explode []string

		sidsIIV := sidsIIR // string as string
		sidsIC = append(sidsIC, sidsIIV)
	}

	// items.CollectionFormat: ""
	sidsIS := swag.JoinByFormat(sidsIC, "")

	return sidsIS
}
