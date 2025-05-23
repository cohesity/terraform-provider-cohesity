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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewUpdateSMTPConfigurationParams creates a new UpdateSMTPConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSMTPConfigurationParams() *UpdateSMTPConfigurationParams {
	return &UpdateSMTPConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSMTPConfigurationParamsWithTimeout creates a new UpdateSMTPConfigurationParams object
// with the ability to set a timeout on a request.
func NewUpdateSMTPConfigurationParamsWithTimeout(timeout time.Duration) *UpdateSMTPConfigurationParams {
	return &UpdateSMTPConfigurationParams{
		timeout: timeout,
	}
}

// NewUpdateSMTPConfigurationParamsWithContext creates a new UpdateSMTPConfigurationParams object
// with the ability to set a context for a request.
func NewUpdateSMTPConfigurationParamsWithContext(ctx context.Context) *UpdateSMTPConfigurationParams {
	return &UpdateSMTPConfigurationParams{
		Context: ctx,
	}
}

// NewUpdateSMTPConfigurationParamsWithHTTPClient creates a new UpdateSMTPConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSMTPConfigurationParamsWithHTTPClient(client *http.Client) *UpdateSMTPConfigurationParams {
	return &UpdateSMTPConfigurationParams{
		HTTPClient: client,
	}
}

/*
UpdateSMTPConfigurationParams contains all the parameters to send to the API endpoint

	for the update SMTP configuration operation.

	Typically these are written to a http.Request.
*/
type UpdateSMTPConfigurationParams struct {

	/* Body.

	   Specifies the parameters to update cluster SMTP configuration.
	*/
	Body *models.UpdateSMTPParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update SMTP configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSMTPConfigurationParams) WithDefaults() *UpdateSMTPConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update SMTP configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSMTPConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) WithTimeout(timeout time.Duration) *UpdateSMTPConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) WithContext(ctx context.Context) *UpdateSMTPConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) WithHTTPClient(client *http.Client) *UpdateSMTPConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) WithBody(body *models.UpdateSMTPParams) *UpdateSMTPConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update SMTP configuration params
func (o *UpdateSMTPConfigurationParams) SetBody(body *models.UpdateSMTPParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSMTPConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
