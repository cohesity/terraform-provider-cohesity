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

// NewVerifySupportUserTotpParams creates a new VerifySupportUserTotpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVerifySupportUserTotpParams() *VerifySupportUserTotpParams {
	return &VerifySupportUserTotpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVerifySupportUserTotpParamsWithTimeout creates a new VerifySupportUserTotpParams object
// with the ability to set a timeout on a request.
func NewVerifySupportUserTotpParamsWithTimeout(timeout time.Duration) *VerifySupportUserTotpParams {
	return &VerifySupportUserTotpParams{
		timeout: timeout,
	}
}

// NewVerifySupportUserTotpParamsWithContext creates a new VerifySupportUserTotpParams object
// with the ability to set a context for a request.
func NewVerifySupportUserTotpParamsWithContext(ctx context.Context) *VerifySupportUserTotpParams {
	return &VerifySupportUserTotpParams{
		Context: ctx,
	}
}

// NewVerifySupportUserTotpParamsWithHTTPClient creates a new VerifySupportUserTotpParams object
// with the ability to set a custom HTTPClient for a request.
func NewVerifySupportUserTotpParamsWithHTTPClient(client *http.Client) *VerifySupportUserTotpParams {
	return &VerifySupportUserTotpParams{
		HTTPClient: client,
	}
}

/*
VerifySupportUserTotpParams contains all the parameters to send to the API endpoint

	for the verify support user totp operation.

	Typically these are written to a http.Request.
*/
type VerifySupportUserTotpParams struct {

	/* Body.

	   Totp code to be verified.
	*/
	Body *models.VerifyTotpRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the verify support user totp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifySupportUserTotpParams) WithDefaults() *VerifySupportUserTotpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the verify support user totp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VerifySupportUserTotpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the verify support user totp params
func (o *VerifySupportUserTotpParams) WithTimeout(timeout time.Duration) *VerifySupportUserTotpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify support user totp params
func (o *VerifySupportUserTotpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify support user totp params
func (o *VerifySupportUserTotpParams) WithContext(ctx context.Context) *VerifySupportUserTotpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify support user totp params
func (o *VerifySupportUserTotpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify support user totp params
func (o *VerifySupportUserTotpParams) WithHTTPClient(client *http.Client) *VerifySupportUserTotpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify support user totp params
func (o *VerifySupportUserTotpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the verify support user totp params
func (o *VerifySupportUserTotpParams) WithBody(body *models.VerifyTotpRequest) *VerifySupportUserTotpParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the verify support user totp params
func (o *VerifySupportUserTotpParams) SetBody(body *models.VerifyTotpRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *VerifySupportUserTotpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
