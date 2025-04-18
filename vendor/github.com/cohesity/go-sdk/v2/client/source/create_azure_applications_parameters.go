// Code generated by go-swagger; DO NOT EDIT.

package source

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

// NewCreateAzureApplicationsParams creates a new CreateAzureApplicationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAzureApplicationsParams() *CreateAzureApplicationsParams {
	return &CreateAzureApplicationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAzureApplicationsParamsWithTimeout creates a new CreateAzureApplicationsParams object
// with the ability to set a timeout on a request.
func NewCreateAzureApplicationsParamsWithTimeout(timeout time.Duration) *CreateAzureApplicationsParams {
	return &CreateAzureApplicationsParams{
		timeout: timeout,
	}
}

// NewCreateAzureApplicationsParamsWithContext creates a new CreateAzureApplicationsParams object
// with the ability to set a context for a request.
func NewCreateAzureApplicationsParamsWithContext(ctx context.Context) *CreateAzureApplicationsParams {
	return &CreateAzureApplicationsParams{
		Context: ctx,
	}
}

// NewCreateAzureApplicationsParamsWithHTTPClient creates a new CreateAzureApplicationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAzureApplicationsParamsWithHTTPClient(client *http.Client) *CreateAzureApplicationsParams {
	return &CreateAzureApplicationsParams{
		HTTPClient: client,
	}
}

/*
CreateAzureApplicationsParams contains all the parameters to send to the API endpoint

	for the create azure applications operation.

	Typically these are written to a http.Request.
*/
type CreateAzureApplicationsParams struct {

	/* Body.

	   Specifies the parameters to create Azure applications within a given Microsoft365 source.
	*/
	Body *models.CreateAzureApplicationRequestParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create azure applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureApplicationsParams) WithDefaults() *CreateAzureApplicationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create azure applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureApplicationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create azure applications params
func (o *CreateAzureApplicationsParams) WithTimeout(timeout time.Duration) *CreateAzureApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create azure applications params
func (o *CreateAzureApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create azure applications params
func (o *CreateAzureApplicationsParams) WithContext(ctx context.Context) *CreateAzureApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create azure applications params
func (o *CreateAzureApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create azure applications params
func (o *CreateAzureApplicationsParams) WithHTTPClient(client *http.Client) *CreateAzureApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create azure applications params
func (o *CreateAzureApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create azure applications params
func (o *CreateAzureApplicationsParams) WithBody(body *models.CreateAzureApplicationRequestParams) *CreateAzureApplicationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create azure applications params
func (o *CreateAzureApplicationsParams) SetBody(body *models.CreateAzureApplicationRequestParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAzureApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
