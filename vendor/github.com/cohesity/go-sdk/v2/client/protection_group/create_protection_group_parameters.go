// Code generated by go-swagger; DO NOT EDIT.

package protection_group

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

// NewCreateProtectionGroupParams creates a new CreateProtectionGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProtectionGroupParams() *CreateProtectionGroupParams {
	return &CreateProtectionGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProtectionGroupParamsWithTimeout creates a new CreateProtectionGroupParams object
// with the ability to set a timeout on a request.
func NewCreateProtectionGroupParamsWithTimeout(timeout time.Duration) *CreateProtectionGroupParams {
	return &CreateProtectionGroupParams{
		timeout: timeout,
	}
}

// NewCreateProtectionGroupParamsWithContext creates a new CreateProtectionGroupParams object
// with the ability to set a context for a request.
func NewCreateProtectionGroupParamsWithContext(ctx context.Context) *CreateProtectionGroupParams {
	return &CreateProtectionGroupParams{
		Context: ctx,
	}
}

// NewCreateProtectionGroupParamsWithHTTPClient creates a new CreateProtectionGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProtectionGroupParamsWithHTTPClient(client *http.Client) *CreateProtectionGroupParams {
	return &CreateProtectionGroupParams{
		HTTPClient: client,
	}
}

/*
CreateProtectionGroupParams contains all the parameters to send to the API endpoint

	for the create protection group operation.

	Typically these are written to a http.Request.
*/
type CreateProtectionGroupParams struct {

	/* Body.

	   Specifies the parameters to create a Protection Group.
	*/
	Body *models.CreateOrUpdateProtectionGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create protection group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProtectionGroupParams) WithDefaults() *CreateProtectionGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create protection group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProtectionGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create protection group params
func (o *CreateProtectionGroupParams) WithTimeout(timeout time.Duration) *CreateProtectionGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create protection group params
func (o *CreateProtectionGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create protection group params
func (o *CreateProtectionGroupParams) WithContext(ctx context.Context) *CreateProtectionGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create protection group params
func (o *CreateProtectionGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create protection group params
func (o *CreateProtectionGroupParams) WithHTTPClient(client *http.Client) *CreateProtectionGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create protection group params
func (o *CreateProtectionGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create protection group params
func (o *CreateProtectionGroupParams) WithBody(body *models.CreateOrUpdateProtectionGroupRequest) *CreateProtectionGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create protection group params
func (o *CreateProtectionGroupParams) SetBody(body *models.CreateOrUpdateProtectionGroupRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProtectionGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
