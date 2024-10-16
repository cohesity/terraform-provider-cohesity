// Code generated by go-swagger; DO NOT EDIT.

package principals

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

// NewListSourcesForPrincipalsParams creates a new ListSourcesForPrincipalsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSourcesForPrincipalsParams() *ListSourcesForPrincipalsParams {
	return &ListSourcesForPrincipalsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSourcesForPrincipalsParamsWithTimeout creates a new ListSourcesForPrincipalsParams object
// with the ability to set a timeout on a request.
func NewListSourcesForPrincipalsParamsWithTimeout(timeout time.Duration) *ListSourcesForPrincipalsParams {
	return &ListSourcesForPrincipalsParams{
		timeout: timeout,
	}
}

// NewListSourcesForPrincipalsParamsWithContext creates a new ListSourcesForPrincipalsParams object
// with the ability to set a context for a request.
func NewListSourcesForPrincipalsParamsWithContext(ctx context.Context) *ListSourcesForPrincipalsParams {
	return &ListSourcesForPrincipalsParams{
		Context: ctx,
	}
}

// NewListSourcesForPrincipalsParamsWithHTTPClient creates a new ListSourcesForPrincipalsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSourcesForPrincipalsParamsWithHTTPClient(client *http.Client) *ListSourcesForPrincipalsParams {
	return &ListSourcesForPrincipalsParams{
		HTTPClient: client,
	}
}

/*
ListSourcesForPrincipalsParams contains all the parameters to send to the API endpoint

	for the list sources for principals operation.

	Typically these are written to a http.Request.
*/
type ListSourcesForPrincipalsParams struct {

	/* Sids.

	     Specifies a list of security identifiers (SIDs) that specify user or
	group principals.
	*/
	Sids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list sources for principals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSourcesForPrincipalsParams) WithDefaults() *ListSourcesForPrincipalsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list sources for principals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSourcesForPrincipalsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) WithTimeout(timeout time.Duration) *ListSourcesForPrincipalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) WithContext(ctx context.Context) *ListSourcesForPrincipalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) WithHTTPClient(client *http.Client) *ListSourcesForPrincipalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSids adds the sids to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) WithSids(sids []string) *ListSourcesForPrincipalsParams {
	o.SetSids(sids)
	return o
}

// SetSids adds the sids to the list sources for principals params
func (o *ListSourcesForPrincipalsParams) SetSids(sids []string) {
	o.Sids = sids
}

// WriteToRequest writes these params to a swagger request
func (o *ListSourcesForPrincipalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamListSourcesForPrincipals binds the parameter sids
func (o *ListSourcesForPrincipalsParams) bindParamSids(formats strfmt.Registry) []string {
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
