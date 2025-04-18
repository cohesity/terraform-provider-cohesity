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

// NewUpdateAirgapConfigParams creates a new UpdateAirgapConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAirgapConfigParams() *UpdateAirgapConfigParams {
	return &UpdateAirgapConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAirgapConfigParamsWithTimeout creates a new UpdateAirgapConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateAirgapConfigParamsWithTimeout(timeout time.Duration) *UpdateAirgapConfigParams {
	return &UpdateAirgapConfigParams{
		timeout: timeout,
	}
}

// NewUpdateAirgapConfigParamsWithContext creates a new UpdateAirgapConfigParams object
// with the ability to set a context for a request.
func NewUpdateAirgapConfigParamsWithContext(ctx context.Context) *UpdateAirgapConfigParams {
	return &UpdateAirgapConfigParams{
		Context: ctx,
	}
}

// NewUpdateAirgapConfigParamsWithHTTPClient creates a new UpdateAirgapConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAirgapConfigParamsWithHTTPClient(client *http.Client) *UpdateAirgapConfigParams {
	return &UpdateAirgapConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateAirgapConfigParams contains all the parameters to send to the API endpoint

	for the update airgap config operation.

	Typically these are written to a http.Request.
*/
type UpdateAirgapConfigParams struct {

	/* Body.

	   Specifies the parameters to update airgap config.
	*/
	Body *models.AirgapConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update airgap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAirgapConfigParams) WithDefaults() *UpdateAirgapConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update airgap config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAirgapConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update airgap config params
func (o *UpdateAirgapConfigParams) WithTimeout(timeout time.Duration) *UpdateAirgapConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update airgap config params
func (o *UpdateAirgapConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update airgap config params
func (o *UpdateAirgapConfigParams) WithContext(ctx context.Context) *UpdateAirgapConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update airgap config params
func (o *UpdateAirgapConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update airgap config params
func (o *UpdateAirgapConfigParams) WithHTTPClient(client *http.Client) *UpdateAirgapConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update airgap config params
func (o *UpdateAirgapConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update airgap config params
func (o *UpdateAirgapConfigParams) WithBody(body *models.AirgapConfig) *UpdateAirgapConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update airgap config params
func (o *UpdateAirgapConfigParams) SetBody(body *models.AirgapConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAirgapConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
