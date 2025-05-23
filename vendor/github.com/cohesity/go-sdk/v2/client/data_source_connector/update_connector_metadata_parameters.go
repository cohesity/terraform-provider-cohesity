// Code generated by go-swagger; DO NOT EDIT.

package data_source_connector

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

// NewUpdateConnectorMetadataParams creates a new UpdateConnectorMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConnectorMetadataParams() *UpdateConnectorMetadataParams {
	return &UpdateConnectorMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConnectorMetadataParamsWithTimeout creates a new UpdateConnectorMetadataParams object
// with the ability to set a timeout on a request.
func NewUpdateConnectorMetadataParamsWithTimeout(timeout time.Duration) *UpdateConnectorMetadataParams {
	return &UpdateConnectorMetadataParams{
		timeout: timeout,
	}
}

// NewUpdateConnectorMetadataParamsWithContext creates a new UpdateConnectorMetadataParams object
// with the ability to set a context for a request.
func NewUpdateConnectorMetadataParamsWithContext(ctx context.Context) *UpdateConnectorMetadataParams {
	return &UpdateConnectorMetadataParams{
		Context: ctx,
	}
}

// NewUpdateConnectorMetadataParamsWithHTTPClient creates a new UpdateConnectorMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConnectorMetadataParamsWithHTTPClient(client *http.Client) *UpdateConnectorMetadataParams {
	return &UpdateConnectorMetadataParams{
		HTTPClient: client,
	}
}

/*
UpdateConnectorMetadataParams contains all the parameters to send to the API endpoint

	for the update connector metadata operation.

	Typically these are written to a http.Request.
*/
type UpdateConnectorMetadataParams struct {

	/* Body.

	   Specifies information about the connectors.
	*/
	Body *models.CreateOrUpdateConnectorMetadataRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update connector metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConnectorMetadataParams) WithDefaults() *UpdateConnectorMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update connector metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConnectorMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update connector metadata params
func (o *UpdateConnectorMetadataParams) WithTimeout(timeout time.Duration) *UpdateConnectorMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update connector metadata params
func (o *UpdateConnectorMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update connector metadata params
func (o *UpdateConnectorMetadataParams) WithContext(ctx context.Context) *UpdateConnectorMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update connector metadata params
func (o *UpdateConnectorMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update connector metadata params
func (o *UpdateConnectorMetadataParams) WithHTTPClient(client *http.Client) *UpdateConnectorMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update connector metadata params
func (o *UpdateConnectorMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update connector metadata params
func (o *UpdateConnectorMetadataParams) WithBody(body *models.CreateOrUpdateConnectorMetadataRequest) *UpdateConnectorMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update connector metadata params
func (o *UpdateConnectorMetadataParams) SetBody(body *models.CreateOrUpdateConnectorMetadataRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConnectorMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
