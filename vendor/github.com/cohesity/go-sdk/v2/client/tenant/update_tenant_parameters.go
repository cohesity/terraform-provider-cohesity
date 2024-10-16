// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewUpdateTenantParams creates a new UpdateTenantParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTenantParams() *UpdateTenantParams {
	return &UpdateTenantParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTenantParamsWithTimeout creates a new UpdateTenantParams object
// with the ability to set a timeout on a request.
func NewUpdateTenantParamsWithTimeout(timeout time.Duration) *UpdateTenantParams {
	return &UpdateTenantParams{
		timeout: timeout,
	}
}

// NewUpdateTenantParamsWithContext creates a new UpdateTenantParams object
// with the ability to set a context for a request.
func NewUpdateTenantParamsWithContext(ctx context.Context) *UpdateTenantParams {
	return &UpdateTenantParams{
		Context: ctx,
	}
}

// NewUpdateTenantParamsWithHTTPClient creates a new UpdateTenantParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTenantParamsWithHTTPClient(client *http.Client) *UpdateTenantParams {
	return &UpdateTenantParams{
		HTTPClient: client,
	}
}

/*
UpdateTenantParams contains all the parameters to send to the API endpoint

	for the update tenant operation.

	Typically these are written to a http.Request.
*/
type UpdateTenantParams struct {

	// Body.
	Body *models.UpdateTenantBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update tenant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantParams) WithDefaults() *UpdateTenantParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update tenant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update tenant params
func (o *UpdateTenantParams) WithTimeout(timeout time.Duration) *UpdateTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tenant params
func (o *UpdateTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tenant params
func (o *UpdateTenantParams) WithContext(ctx context.Context) *UpdateTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tenant params
func (o *UpdateTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tenant params
func (o *UpdateTenantParams) WithHTTPClient(client *http.Client) *UpdateTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tenant params
func (o *UpdateTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tenant params
func (o *UpdateTenantParams) WithBody(body *models.UpdateTenantBody) *UpdateTenantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tenant params
func (o *UpdateTenantParams) SetBody(body *models.UpdateTenantBody) {
	o.Body = body
}

// WithID adds the id to the update tenant params
func (o *UpdateTenantParams) WithID(id string) *UpdateTenantParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update tenant params
func (o *UpdateTenantParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
