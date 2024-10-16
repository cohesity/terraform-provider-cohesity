// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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
)

// NewListCentrifyZonesParams creates a new ListCentrifyZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCentrifyZonesParams() *ListCentrifyZonesParams {
	return &ListCentrifyZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCentrifyZonesParamsWithTimeout creates a new ListCentrifyZonesParams object
// with the ability to set a timeout on a request.
func NewListCentrifyZonesParamsWithTimeout(timeout time.Duration) *ListCentrifyZonesParams {
	return &ListCentrifyZonesParams{
		timeout: timeout,
	}
}

// NewListCentrifyZonesParamsWithContext creates a new ListCentrifyZonesParams object
// with the ability to set a context for a request.
func NewListCentrifyZonesParamsWithContext(ctx context.Context) *ListCentrifyZonesParams {
	return &ListCentrifyZonesParams{
		Context: ctx,
	}
}

// NewListCentrifyZonesParamsWithHTTPClient creates a new ListCentrifyZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCentrifyZonesParamsWithHTTPClient(client *http.Client) *ListCentrifyZonesParams {
	return &ListCentrifyZonesParams{
		HTTPClient: client,
	}
}

/*
ListCentrifyZonesParams contains all the parameters to send to the API endpoint

	for the list centrify zones operation.

	Typically these are written to a http.Request.
*/
type ListCentrifyZonesParams struct {

	/* DomainName.

	   Specifies the fully qualified domain name (FQDN) of an Active Directory.
	*/
	DomainName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list centrify zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCentrifyZonesParams) WithDefaults() *ListCentrifyZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list centrify zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCentrifyZonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list centrify zones params
func (o *ListCentrifyZonesParams) WithTimeout(timeout time.Duration) *ListCentrifyZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list centrify zones params
func (o *ListCentrifyZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list centrify zones params
func (o *ListCentrifyZonesParams) WithContext(ctx context.Context) *ListCentrifyZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list centrify zones params
func (o *ListCentrifyZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list centrify zones params
func (o *ListCentrifyZonesParams) WithHTTPClient(client *http.Client) *ListCentrifyZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list centrify zones params
func (o *ListCentrifyZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the list centrify zones params
func (o *ListCentrifyZonesParams) WithDomainName(domainName *string) *ListCentrifyZonesParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the list centrify zones params
func (o *ListCentrifyZonesParams) SetDomainName(domainName *string) {
	o.DomainName = domainName
}

// WriteToRequest writes these params to a swagger request
func (o *ListCentrifyZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainName != nil {

		// query param domainName
		var qrDomainName string

		if o.DomainName != nil {
			qrDomainName = *o.DomainName
		}
		qDomainName := qrDomainName
		if qDomainName != "" {

			if err := r.SetQueryParam("domainName", qDomainName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
