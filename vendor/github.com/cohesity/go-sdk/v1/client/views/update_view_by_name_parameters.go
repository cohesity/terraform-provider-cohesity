// Code generated by go-swagger; DO NOT EDIT.

package views

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

// NewUpdateViewByNameParams creates a new UpdateViewByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateViewByNameParams() *UpdateViewByNameParams {
	return &UpdateViewByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateViewByNameParamsWithTimeout creates a new UpdateViewByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateViewByNameParamsWithTimeout(timeout time.Duration) *UpdateViewByNameParams {
	return &UpdateViewByNameParams{
		timeout: timeout,
	}
}

// NewUpdateViewByNameParamsWithContext creates a new UpdateViewByNameParams object
// with the ability to set a context for a request.
func NewUpdateViewByNameParamsWithContext(ctx context.Context) *UpdateViewByNameParams {
	return &UpdateViewByNameParams{
		Context: ctx,
	}
}

// NewUpdateViewByNameParamsWithHTTPClient creates a new UpdateViewByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateViewByNameParamsWithHTTPClient(client *http.Client) *UpdateViewByNameParams {
	return &UpdateViewByNameParams{
		HTTPClient: client,
	}
}

/*
UpdateViewByNameParams contains all the parameters to send to the API endpoint

	for the update view by name operation.

	Typically these are written to a http.Request.
*/
type UpdateViewByNameParams struct {

	/* Body.

	   Request to update a view.
	*/
	Body *models.UpdateViewParam

	/* Name.

	   Specifies the View name.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update view by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewByNameParams) WithDefaults() *UpdateViewByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update view by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update view by name params
func (o *UpdateViewByNameParams) WithTimeout(timeout time.Duration) *UpdateViewByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update view by name params
func (o *UpdateViewByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update view by name params
func (o *UpdateViewByNameParams) WithContext(ctx context.Context) *UpdateViewByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update view by name params
func (o *UpdateViewByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update view by name params
func (o *UpdateViewByNameParams) WithHTTPClient(client *http.Client) *UpdateViewByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update view by name params
func (o *UpdateViewByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update view by name params
func (o *UpdateViewByNameParams) WithBody(body *models.UpdateViewParam) *UpdateViewByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update view by name params
func (o *UpdateViewByNameParams) SetBody(body *models.UpdateViewParam) {
	o.Body = body
}

// WithName adds the name to the update view by name params
func (o *UpdateViewByNameParams) WithName(name string) *UpdateViewByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update view by name params
func (o *UpdateViewByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateViewByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
