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

// NewListFreeNodesParams creates a new ListFreeNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListFreeNodesParams() *ListFreeNodesParams {
	return &ListFreeNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListFreeNodesParamsWithTimeout creates a new ListFreeNodesParams object
// with the ability to set a timeout on a request.
func NewListFreeNodesParamsWithTimeout(timeout time.Duration) *ListFreeNodesParams {
	return &ListFreeNodesParams{
		timeout: timeout,
	}
}

// NewListFreeNodesParamsWithContext creates a new ListFreeNodesParams object
// with the ability to set a context for a request.
func NewListFreeNodesParamsWithContext(ctx context.Context) *ListFreeNodesParams {
	return &ListFreeNodesParams{
		Context: ctx,
	}
}

// NewListFreeNodesParamsWithHTTPClient creates a new ListFreeNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListFreeNodesParamsWithHTTPClient(client *http.Client) *ListFreeNodesParams {
	return &ListFreeNodesParams{
		HTTPClient: client,
	}
}

/*
ListFreeNodesParams contains all the parameters to send to the API endpoint

	for the list free nodes operation.

	Typically these are written to a http.Request.
*/
type ListFreeNodesParams struct {

	/* Ips.

	   "Specifies a list of ips of nodes among which free and compatible nodes to be returned"
	*/
	Ips []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list free nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFreeNodesParams) WithDefaults() *ListFreeNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list free nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListFreeNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list free nodes params
func (o *ListFreeNodesParams) WithTimeout(timeout time.Duration) *ListFreeNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list free nodes params
func (o *ListFreeNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list free nodes params
func (o *ListFreeNodesParams) WithContext(ctx context.Context) *ListFreeNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list free nodes params
func (o *ListFreeNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list free nodes params
func (o *ListFreeNodesParams) WithHTTPClient(client *http.Client) *ListFreeNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list free nodes params
func (o *ListFreeNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIps adds the ips to the list free nodes params
func (o *ListFreeNodesParams) WithIps(ips []string) *ListFreeNodesParams {
	o.SetIps(ips)
	return o
}

// SetIps adds the ips to the list free nodes params
func (o *ListFreeNodesParams) SetIps(ips []string) {
	o.Ips = ips
}

// WriteToRequest writes these params to a swagger request
func (o *ListFreeNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ips != nil {

		// binding items for ips
		joinedIps := o.bindParamIps(reg)

		// query array param ips
		if err := r.SetQueryParam("ips", joinedIps...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListFreeNodes binds the parameter ips
func (o *ListFreeNodesParams) bindParamIps(formats strfmt.Registry) []string {
	ipsIR := o.Ips

	var ipsIC []string
	for _, ipsIIR := range ipsIR { // explode []string

		ipsIIV := ipsIIR // string as string
		ipsIC = append(ipsIC, ipsIIV)
	}

	// items.CollectionFormat: ""
	ipsIS := swag.JoinByFormat(ipsIC, "")

	return ipsIS
}
