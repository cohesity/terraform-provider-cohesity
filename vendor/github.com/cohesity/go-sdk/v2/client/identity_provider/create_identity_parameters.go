// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

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

// NewCreateIdentityParams creates a new CreateIdentityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateIdentityParams() *CreateIdentityParams {
	return &CreateIdentityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateIdentityParamsWithTimeout creates a new CreateIdentityParams object
// with the ability to set a timeout on a request.
func NewCreateIdentityParamsWithTimeout(timeout time.Duration) *CreateIdentityParams {
	return &CreateIdentityParams{
		timeout: timeout,
	}
}

// NewCreateIdentityParamsWithContext creates a new CreateIdentityParams object
// with the ability to set a context for a request.
func NewCreateIdentityParamsWithContext(ctx context.Context) *CreateIdentityParams {
	return &CreateIdentityParams{
		Context: ctx,
	}
}

// NewCreateIdentityParamsWithHTTPClient creates a new CreateIdentityParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateIdentityParamsWithHTTPClient(client *http.Client) *CreateIdentityParams {
	return &CreateIdentityParams{
		HTTPClient: client,
	}
}

/*
CreateIdentityParams contains all the parameters to send to the API endpoint

	for the create identity operation.

	Typically these are written to a http.Request.
*/
type CreateIdentityParams struct {

	/* Body.

	   Specifies parameters to configure Identity
	*/
	Body *models.IdentityConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIdentityParams) WithDefaults() *CreateIdentityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateIdentityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create identity params
func (o *CreateIdentityParams) WithTimeout(timeout time.Duration) *CreateIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create identity params
func (o *CreateIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create identity params
func (o *CreateIdentityParams) WithContext(ctx context.Context) *CreateIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create identity params
func (o *CreateIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create identity params
func (o *CreateIdentityParams) WithHTTPClient(client *http.Client) *CreateIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create identity params
func (o *CreateIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create identity params
func (o *CreateIdentityParams) WithBody(body *models.IdentityConfig) *CreateIdentityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create identity params
func (o *CreateIdentityParams) SetBody(body *models.IdentityConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
