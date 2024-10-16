// Code generated by go-swagger; DO NOT EDIT.

package failover

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

// NewCreatePlannedRunParams creates a new CreatePlannedRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePlannedRunParams() *CreatePlannedRunParams {
	return &CreatePlannedRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePlannedRunParamsWithTimeout creates a new CreatePlannedRunParams object
// with the ability to set a timeout on a request.
func NewCreatePlannedRunParamsWithTimeout(timeout time.Duration) *CreatePlannedRunParams {
	return &CreatePlannedRunParams{
		timeout: timeout,
	}
}

// NewCreatePlannedRunParamsWithContext creates a new CreatePlannedRunParams object
// with the ability to set a context for a request.
func NewCreatePlannedRunParamsWithContext(ctx context.Context) *CreatePlannedRunParams {
	return &CreatePlannedRunParams{
		Context: ctx,
	}
}

// NewCreatePlannedRunParamsWithHTTPClient creates a new CreatePlannedRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePlannedRunParamsWithHTTPClient(client *http.Client) *CreatePlannedRunParams {
	return &CreatePlannedRunParams{
		HTTPClient: client,
	}
}

/*
CreatePlannedRunParams contains all the parameters to send to the API endpoint

	for the create planned run operation.

	Typically these are written to a http.Request.
*/
type CreatePlannedRunParams struct {

	/* Body.

	   Specifies the paramteres to create a planned run while failover workflow.
	*/
	Body *models.FailoverRunConfiguration

	/* ID.

	   Specifies the id of the failover workflow.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create planned run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePlannedRunParams) WithDefaults() *CreatePlannedRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create planned run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePlannedRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create planned run params
func (o *CreatePlannedRunParams) WithTimeout(timeout time.Duration) *CreatePlannedRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create planned run params
func (o *CreatePlannedRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create planned run params
func (o *CreatePlannedRunParams) WithContext(ctx context.Context) *CreatePlannedRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create planned run params
func (o *CreatePlannedRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create planned run params
func (o *CreatePlannedRunParams) WithHTTPClient(client *http.Client) *CreatePlannedRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create planned run params
func (o *CreatePlannedRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create planned run params
func (o *CreatePlannedRunParams) WithBody(body *models.FailoverRunConfiguration) *CreatePlannedRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create planned run params
func (o *CreatePlannedRunParams) SetBody(body *models.FailoverRunConfiguration) {
	o.Body = body
}

// WithID adds the id to the create planned run params
func (o *CreatePlannedRunParams) WithID(id string) *CreatePlannedRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create planned run params
func (o *CreatePlannedRunParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePlannedRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
