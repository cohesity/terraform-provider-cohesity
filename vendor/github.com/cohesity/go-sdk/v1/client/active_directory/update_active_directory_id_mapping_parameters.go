// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

// NewUpdateActiveDirectoryIDMappingParams creates a new UpdateActiveDirectoryIDMappingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateActiveDirectoryIDMappingParams() *UpdateActiveDirectoryIDMappingParams {
	return &UpdateActiveDirectoryIDMappingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActiveDirectoryIDMappingParamsWithTimeout creates a new UpdateActiveDirectoryIDMappingParams object
// with the ability to set a timeout on a request.
func NewUpdateActiveDirectoryIDMappingParamsWithTimeout(timeout time.Duration) *UpdateActiveDirectoryIDMappingParams {
	return &UpdateActiveDirectoryIDMappingParams{
		timeout: timeout,
	}
}

// NewUpdateActiveDirectoryIDMappingParamsWithContext creates a new UpdateActiveDirectoryIDMappingParams object
// with the ability to set a context for a request.
func NewUpdateActiveDirectoryIDMappingParamsWithContext(ctx context.Context) *UpdateActiveDirectoryIDMappingParams {
	return &UpdateActiveDirectoryIDMappingParams{
		Context: ctx,
	}
}

// NewUpdateActiveDirectoryIDMappingParamsWithHTTPClient creates a new UpdateActiveDirectoryIDMappingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateActiveDirectoryIDMappingParamsWithHTTPClient(client *http.Client) *UpdateActiveDirectoryIDMappingParams {
	return &UpdateActiveDirectoryIDMappingParams{
		HTTPClient: client,
	}
}

/*
UpdateActiveDirectoryIDMappingParams contains all the parameters to send to the API endpoint

	for the update active directory Id mapping operation.

	Typically these are written to a http.Request.
*/
type UpdateActiveDirectoryIDMappingParams struct {

	/* Body.

	   Request to update user id mapping of an Active Directory.
	*/
	Body *models.IDMappingInfo

	/* Name.

	   Specifies the Active Directory Domain Name.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update active directory Id mapping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateActiveDirectoryIDMappingParams) WithDefaults() *UpdateActiveDirectoryIDMappingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update active directory Id mapping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateActiveDirectoryIDMappingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) WithTimeout(timeout time.Duration) *UpdateActiveDirectoryIDMappingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) WithContext(ctx context.Context) *UpdateActiveDirectoryIDMappingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) WithHTTPClient(client *http.Client) *UpdateActiveDirectoryIDMappingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) WithBody(body *models.IDMappingInfo) *UpdateActiveDirectoryIDMappingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) SetBody(body *models.IDMappingInfo) {
	o.Body = body
}

// WithName adds the name to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) WithName(name string) *UpdateActiveDirectoryIDMappingParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update active directory Id mapping params
func (o *UpdateActiveDirectoryIDMappingParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActiveDirectoryIDMappingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
