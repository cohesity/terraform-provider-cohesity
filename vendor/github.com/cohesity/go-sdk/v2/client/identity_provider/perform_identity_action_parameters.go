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

// NewPerformIdentityActionParams creates a new PerformIdentityActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformIdentityActionParams() *PerformIdentityActionParams {
	return &PerformIdentityActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformIdentityActionParamsWithTimeout creates a new PerformIdentityActionParams object
// with the ability to set a timeout on a request.
func NewPerformIdentityActionParamsWithTimeout(timeout time.Duration) *PerformIdentityActionParams {
	return &PerformIdentityActionParams{
		timeout: timeout,
	}
}

// NewPerformIdentityActionParamsWithContext creates a new PerformIdentityActionParams object
// with the ability to set a context for a request.
func NewPerformIdentityActionParamsWithContext(ctx context.Context) *PerformIdentityActionParams {
	return &PerformIdentityActionParams{
		Context: ctx,
	}
}

// NewPerformIdentityActionParamsWithHTTPClient creates a new PerformIdentityActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformIdentityActionParamsWithHTTPClient(client *http.Client) *PerformIdentityActionParams {
	return &PerformIdentityActionParams{
		HTTPClient: client,
	}
}

/*
PerformIdentityActionParams contains all the parameters to send to the API endpoint

	for the perform identity action operation.

	Typically these are written to a http.Request.
*/
type PerformIdentityActionParams struct {

	/* Body.

	   Specifies parameters perform an identity action.
	*/
	Body *models.IdentityAction

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform identity action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformIdentityActionParams) WithDefaults() *PerformIdentityActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform identity action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformIdentityActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform identity action params
func (o *PerformIdentityActionParams) WithTimeout(timeout time.Duration) *PerformIdentityActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform identity action params
func (o *PerformIdentityActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform identity action params
func (o *PerformIdentityActionParams) WithContext(ctx context.Context) *PerformIdentityActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform identity action params
func (o *PerformIdentityActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform identity action params
func (o *PerformIdentityActionParams) WithHTTPClient(client *http.Client) *PerformIdentityActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform identity action params
func (o *PerformIdentityActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the perform identity action params
func (o *PerformIdentityActionParams) WithBody(body *models.IdentityAction) *PerformIdentityActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform identity action params
func (o *PerformIdentityActionParams) SetBody(body *models.IdentityAction) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PerformIdentityActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
