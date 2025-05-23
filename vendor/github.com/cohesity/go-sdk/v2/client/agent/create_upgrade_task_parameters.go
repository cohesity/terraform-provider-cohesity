// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// NewCreateUpgradeTaskParams creates a new CreateUpgradeTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUpgradeTaskParams() *CreateUpgradeTaskParams {
	return &CreateUpgradeTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUpgradeTaskParamsWithTimeout creates a new CreateUpgradeTaskParams object
// with the ability to set a timeout on a request.
func NewCreateUpgradeTaskParamsWithTimeout(timeout time.Duration) *CreateUpgradeTaskParams {
	return &CreateUpgradeTaskParams{
		timeout: timeout,
	}
}

// NewCreateUpgradeTaskParamsWithContext creates a new CreateUpgradeTaskParams object
// with the ability to set a context for a request.
func NewCreateUpgradeTaskParamsWithContext(ctx context.Context) *CreateUpgradeTaskParams {
	return &CreateUpgradeTaskParams{
		Context: ctx,
	}
}

// NewCreateUpgradeTaskParamsWithHTTPClient creates a new CreateUpgradeTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUpgradeTaskParamsWithHTTPClient(client *http.Client) *CreateUpgradeTaskParams {
	return &CreateUpgradeTaskParams{
		HTTPClient: client,
	}
}

/*
CreateUpgradeTaskParams contains all the parameters to send to the API endpoint

	for the create upgrade task operation.

	Typically these are written to a http.Request.
*/
type CreateUpgradeTaskParams struct {

	/* Body.

	   Specifies parameters to create a schedule based agent upgrade task.
	*/
	Body *models.CreateUpgradeTaskRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create upgrade task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUpgradeTaskParams) WithDefaults() *CreateUpgradeTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create upgrade task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUpgradeTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create upgrade task params
func (o *CreateUpgradeTaskParams) WithTimeout(timeout time.Duration) *CreateUpgradeTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create upgrade task params
func (o *CreateUpgradeTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create upgrade task params
func (o *CreateUpgradeTaskParams) WithContext(ctx context.Context) *CreateUpgradeTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create upgrade task params
func (o *CreateUpgradeTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create upgrade task params
func (o *CreateUpgradeTaskParams) WithHTTPClient(client *http.Client) *CreateUpgradeTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create upgrade task params
func (o *CreateUpgradeTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create upgrade task params
func (o *CreateUpgradeTaskParams) WithBody(body *models.CreateUpgradeTaskRequest) *CreateUpgradeTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create upgrade task params
func (o *CreateUpgradeTaskParams) SetBody(body *models.CreateUpgradeTaskRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUpgradeTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
