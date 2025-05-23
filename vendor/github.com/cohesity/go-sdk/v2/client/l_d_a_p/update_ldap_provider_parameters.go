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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewUpdateLdapProviderParams creates a new UpdateLdapProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLdapProviderParams() *UpdateLdapProviderParams {
	return &UpdateLdapProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLdapProviderParamsWithTimeout creates a new UpdateLdapProviderParams object
// with the ability to set a timeout on a request.
func NewUpdateLdapProviderParamsWithTimeout(timeout time.Duration) *UpdateLdapProviderParams {
	return &UpdateLdapProviderParams{
		timeout: timeout,
	}
}

// NewUpdateLdapProviderParamsWithContext creates a new UpdateLdapProviderParams object
// with the ability to set a context for a request.
func NewUpdateLdapProviderParamsWithContext(ctx context.Context) *UpdateLdapProviderParams {
	return &UpdateLdapProviderParams{
		Context: ctx,
	}
}

// NewUpdateLdapProviderParamsWithHTTPClient creates a new UpdateLdapProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLdapProviderParamsWithHTTPClient(client *http.Client) *UpdateLdapProviderParams {
	return &UpdateLdapProviderParams{
		HTTPClient: client,
	}
}

/*
UpdateLdapProviderParams contains all the parameters to send to the API endpoint

	for the update ldap provider operation.

	Typically these are written to a http.Request.
*/
type UpdateLdapProviderParams struct {

	/* Body.

	   Specifies the parameters to update Ldap provider.
	*/
	Body *models.Ldap

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update ldap provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLdapProviderParams) WithDefaults() *UpdateLdapProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update ldap provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLdapProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update ldap provider params
func (o *UpdateLdapProviderParams) WithTimeout(timeout time.Duration) *UpdateLdapProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ldap provider params
func (o *UpdateLdapProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ldap provider params
func (o *UpdateLdapProviderParams) WithContext(ctx context.Context) *UpdateLdapProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ldap provider params
func (o *UpdateLdapProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ldap provider params
func (o *UpdateLdapProviderParams) WithHTTPClient(client *http.Client) *UpdateLdapProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ldap provider params
func (o *UpdateLdapProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update ldap provider params
func (o *UpdateLdapProviderParams) WithBody(body *models.Ldap) *UpdateLdapProviderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update ldap provider params
func (o *UpdateLdapProviderParams) SetBody(body *models.Ldap) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLdapProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
