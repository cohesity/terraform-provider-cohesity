// Code generated by go-swagger; DO NOT EDIT.

package m_f_a

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

// NewCreateTotpKeyParams creates a new CreateTotpKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTotpKeyParams() *CreateTotpKeyParams {
	return &CreateTotpKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTotpKeyParamsWithTimeout creates a new CreateTotpKeyParams object
// with the ability to set a timeout on a request.
func NewCreateTotpKeyParamsWithTimeout(timeout time.Duration) *CreateTotpKeyParams {
	return &CreateTotpKeyParams{
		timeout: timeout,
	}
}

// NewCreateTotpKeyParamsWithContext creates a new CreateTotpKeyParams object
// with the ability to set a context for a request.
func NewCreateTotpKeyParamsWithContext(ctx context.Context) *CreateTotpKeyParams {
	return &CreateTotpKeyParams{
		Context: ctx,
	}
}

// NewCreateTotpKeyParamsWithHTTPClient creates a new CreateTotpKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTotpKeyParamsWithHTTPClient(client *http.Client) *CreateTotpKeyParams {
	return &CreateTotpKeyParams{
		HTTPClient: client,
	}
}

/*
CreateTotpKeyParams contains all the parameters to send to the API endpoint

	for the create totp key operation.

	Typically these are written to a http.Request.
*/
type CreateTotpKeyParams struct {

	/* Body.

	   Specifies the key id for creating the TOTP key.
	*/
	Body *models.CreateTotpKeyRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create totp key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTotpKeyParams) WithDefaults() *CreateTotpKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create totp key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTotpKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create totp key params
func (o *CreateTotpKeyParams) WithTimeout(timeout time.Duration) *CreateTotpKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create totp key params
func (o *CreateTotpKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create totp key params
func (o *CreateTotpKeyParams) WithContext(ctx context.Context) *CreateTotpKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create totp key params
func (o *CreateTotpKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create totp key params
func (o *CreateTotpKeyParams) WithHTTPClient(client *http.Client) *CreateTotpKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create totp key params
func (o *CreateTotpKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create totp key params
func (o *CreateTotpKeyParams) WithBody(body *models.CreateTotpKeyRequestBody) *CreateTotpKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create totp key params
func (o *CreateTotpKeyParams) SetBody(body *models.CreateTotpKeyRequestBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTotpKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
