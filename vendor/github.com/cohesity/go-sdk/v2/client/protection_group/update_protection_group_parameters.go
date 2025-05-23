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

// NewUpdateProtectionGroupParams creates a new UpdateProtectionGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProtectionGroupParams() *UpdateProtectionGroupParams {
	return &UpdateProtectionGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProtectionGroupParamsWithTimeout creates a new UpdateProtectionGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateProtectionGroupParamsWithTimeout(timeout time.Duration) *UpdateProtectionGroupParams {
	return &UpdateProtectionGroupParams{
		timeout: timeout,
	}
}

// NewUpdateProtectionGroupParamsWithContext creates a new UpdateProtectionGroupParams object
// with the ability to set a context for a request.
func NewUpdateProtectionGroupParamsWithContext(ctx context.Context) *UpdateProtectionGroupParams {
	return &UpdateProtectionGroupParams{
		Context: ctx,
	}
}

// NewUpdateProtectionGroupParamsWithHTTPClient creates a new UpdateProtectionGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProtectionGroupParamsWithHTTPClient(client *http.Client) *UpdateProtectionGroupParams {
	return &UpdateProtectionGroupParams{
		HTTPClient: client,
	}
}

/*
UpdateProtectionGroupParams contains all the parameters to send to the API endpoint

	for the update protection group operation.

	Typically these are written to a http.Request.
*/
type UpdateProtectionGroupParams struct {

	/* Body.

	   Specifies the parameters to update a Protection Group.
	*/
	Body *models.CreateOrUpdateProtectionGroupRequest

	/* ID.

	   Specifies the id of the Protection Group.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update protection group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProtectionGroupParams) WithDefaults() *UpdateProtectionGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update protection group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProtectionGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update protection group params
func (o *UpdateProtectionGroupParams) WithTimeout(timeout time.Duration) *UpdateProtectionGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update protection group params
func (o *UpdateProtectionGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update protection group params
func (o *UpdateProtectionGroupParams) WithContext(ctx context.Context) *UpdateProtectionGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update protection group params
func (o *UpdateProtectionGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update protection group params
func (o *UpdateProtectionGroupParams) WithHTTPClient(client *http.Client) *UpdateProtectionGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update protection group params
func (o *UpdateProtectionGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update protection group params
func (o *UpdateProtectionGroupParams) WithBody(body *models.CreateOrUpdateProtectionGroupRequest) *UpdateProtectionGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update protection group params
func (o *UpdateProtectionGroupParams) SetBody(body *models.CreateOrUpdateProtectionGroupRequest) {
	o.Body = body
}

// WithID adds the id to the update protection group params
func (o *UpdateProtectionGroupParams) WithID(id string) *UpdateProtectionGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update protection group params
func (o *UpdateProtectionGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProtectionGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
