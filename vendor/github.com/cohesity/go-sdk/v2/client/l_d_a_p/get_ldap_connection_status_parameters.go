// Code generated by go-swagger; DO NOT EDIT.

package l_d_a_p

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

// NewGetLdapConnectionStatusParams creates a new GetLdapConnectionStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLdapConnectionStatusParams() *GetLdapConnectionStatusParams {
	return &GetLdapConnectionStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapConnectionStatusParamsWithTimeout creates a new GetLdapConnectionStatusParams object
// with the ability to set a timeout on a request.
func NewGetLdapConnectionStatusParamsWithTimeout(timeout time.Duration) *GetLdapConnectionStatusParams {
	return &GetLdapConnectionStatusParams{
		timeout: timeout,
	}
}

// NewGetLdapConnectionStatusParamsWithContext creates a new GetLdapConnectionStatusParams object
// with the ability to set a context for a request.
func NewGetLdapConnectionStatusParamsWithContext(ctx context.Context) *GetLdapConnectionStatusParams {
	return &GetLdapConnectionStatusParams{
		Context: ctx,
	}
}

// NewGetLdapConnectionStatusParamsWithHTTPClient creates a new GetLdapConnectionStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLdapConnectionStatusParamsWithHTTPClient(client *http.Client) *GetLdapConnectionStatusParams {
	return &GetLdapConnectionStatusParams{
		HTTPClient: client,
	}
}

/*
GetLdapConnectionStatusParams contains all the parameters to send to the API endpoint

	for the get ldap connection status operation.

	Typically these are written to a http.Request.
*/
type GetLdapConnectionStatusParams struct {

	/* ID.

	   Specifies the LDAP id.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ldap connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapConnectionStatusParams) WithDefaults() *GetLdapConnectionStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ldap connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapConnectionStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) WithTimeout(timeout time.Duration) *GetLdapConnectionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) WithContext(ctx context.Context) *GetLdapConnectionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) WithHTTPClient(client *http.Client) *GetLdapConnectionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) WithID(id int64) *GetLdapConnectionStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ldap connection status params
func (o *GetLdapConnectionStatusParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapConnectionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
