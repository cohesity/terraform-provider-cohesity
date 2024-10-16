// Code generated by go-swagger; DO NOT EDIT.

package remote_restore

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

	"github.com/cohesity/go-sdk/v1/models"
)

// NewCreateCloudDomainMigrationRequestParams creates a new CreateCloudDomainMigrationRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCloudDomainMigrationRequestParams() *CreateCloudDomainMigrationRequestParams {
	return &CreateCloudDomainMigrationRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudDomainMigrationRequestParamsWithTimeout creates a new CreateCloudDomainMigrationRequestParams object
// with the ability to set a timeout on a request.
func NewCreateCloudDomainMigrationRequestParamsWithTimeout(timeout time.Duration) *CreateCloudDomainMigrationRequestParams {
	return &CreateCloudDomainMigrationRequestParams{
		timeout: timeout,
	}
}

// NewCreateCloudDomainMigrationRequestParamsWithContext creates a new CreateCloudDomainMigrationRequestParams object
// with the ability to set a context for a request.
func NewCreateCloudDomainMigrationRequestParamsWithContext(ctx context.Context) *CreateCloudDomainMigrationRequestParams {
	return &CreateCloudDomainMigrationRequestParams{
		Context: ctx,
	}
}

// NewCreateCloudDomainMigrationRequestParamsWithHTTPClient creates a new CreateCloudDomainMigrationRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCloudDomainMigrationRequestParamsWithHTTPClient(client *http.Client) *CreateCloudDomainMigrationRequestParams {
	return &CreateCloudDomainMigrationRequestParams{
		HTTPClient: client,
	}
}

/*
CreateCloudDomainMigrationRequestParams contains all the parameters to send to the API endpoint

	for the create cloud domain migration request operation.

	Typically these are written to a http.Request.
*/
type CreateCloudDomainMigrationRequestParams struct {

	/* Body.

	   Request to schedule a cloud domain migration task.
	*/
	Body *models.CreateCloudDomainMigrationParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cloud domain migration request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudDomainMigrationRequestParams) WithDefaults() *CreateCloudDomainMigrationRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cloud domain migration request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudDomainMigrationRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) WithTimeout(timeout time.Duration) *CreateCloudDomainMigrationRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) WithContext(ctx context.Context) *CreateCloudDomainMigrationRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) WithHTTPClient(client *http.Client) *CreateCloudDomainMigrationRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) WithBody(body *models.CreateCloudDomainMigrationParameters) *CreateCloudDomainMigrationRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cloud domain migration request params
func (o *CreateCloudDomainMigrationRequestParams) SetBody(body *models.CreateCloudDomainMigrationParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudDomainMigrationRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
