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

// NewCreateProtectionGroupRunParams creates a new CreateProtectionGroupRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProtectionGroupRunParams() *CreateProtectionGroupRunParams {
	return &CreateProtectionGroupRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProtectionGroupRunParamsWithTimeout creates a new CreateProtectionGroupRunParams object
// with the ability to set a timeout on a request.
func NewCreateProtectionGroupRunParamsWithTimeout(timeout time.Duration) *CreateProtectionGroupRunParams {
	return &CreateProtectionGroupRunParams{
		timeout: timeout,
	}
}

// NewCreateProtectionGroupRunParamsWithContext creates a new CreateProtectionGroupRunParams object
// with the ability to set a context for a request.
func NewCreateProtectionGroupRunParamsWithContext(ctx context.Context) *CreateProtectionGroupRunParams {
	return &CreateProtectionGroupRunParams{
		Context: ctx,
	}
}

// NewCreateProtectionGroupRunParamsWithHTTPClient creates a new CreateProtectionGroupRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProtectionGroupRunParamsWithHTTPClient(client *http.Client) *CreateProtectionGroupRunParams {
	return &CreateProtectionGroupRunParams{
		HTTPClient: client,
	}
}

/*
CreateProtectionGroupRunParams contains all the parameters to send to the API endpoint

	for the create protection group run operation.

	Typically these are written to a http.Request.
*/
type CreateProtectionGroupRunParams struct {

	/* Body.

	   Specifies the parameters to start a protection run.
	*/
	Body *models.CreateProtectionGroupRunRequest

	/* ID.

	   Specifies a unique id of the Protection Group.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create protection group run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProtectionGroupRunParams) WithDefaults() *CreateProtectionGroupRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create protection group run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProtectionGroupRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create protection group run params
func (o *CreateProtectionGroupRunParams) WithTimeout(timeout time.Duration) *CreateProtectionGroupRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create protection group run params
func (o *CreateProtectionGroupRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create protection group run params
func (o *CreateProtectionGroupRunParams) WithContext(ctx context.Context) *CreateProtectionGroupRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create protection group run params
func (o *CreateProtectionGroupRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create protection group run params
func (o *CreateProtectionGroupRunParams) WithHTTPClient(client *http.Client) *CreateProtectionGroupRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create protection group run params
func (o *CreateProtectionGroupRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create protection group run params
func (o *CreateProtectionGroupRunParams) WithBody(body *models.CreateProtectionGroupRunRequest) *CreateProtectionGroupRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create protection group run params
func (o *CreateProtectionGroupRunParams) SetBody(body *models.CreateProtectionGroupRunRequest) {
	o.Body = body
}

// WithID adds the id to the create protection group run params
func (o *CreateProtectionGroupRunParams) WithID(id string) *CreateProtectionGroupRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create protection group run params
func (o *CreateProtectionGroupRunParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProtectionGroupRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
