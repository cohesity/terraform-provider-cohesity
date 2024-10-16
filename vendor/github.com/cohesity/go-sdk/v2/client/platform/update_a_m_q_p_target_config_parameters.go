// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// NewUpdateAMQPTargetConfigParams creates a new UpdateAMQPTargetConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAMQPTargetConfigParams() *UpdateAMQPTargetConfigParams {
	return &UpdateAMQPTargetConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAMQPTargetConfigParamsWithTimeout creates a new UpdateAMQPTargetConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateAMQPTargetConfigParamsWithTimeout(timeout time.Duration) *UpdateAMQPTargetConfigParams {
	return &UpdateAMQPTargetConfigParams{
		timeout: timeout,
	}
}

// NewUpdateAMQPTargetConfigParamsWithContext creates a new UpdateAMQPTargetConfigParams object
// with the ability to set a context for a request.
func NewUpdateAMQPTargetConfigParamsWithContext(ctx context.Context) *UpdateAMQPTargetConfigParams {
	return &UpdateAMQPTargetConfigParams{
		Context: ctx,
	}
}

// NewUpdateAMQPTargetConfigParamsWithHTTPClient creates a new UpdateAMQPTargetConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAMQPTargetConfigParamsWithHTTPClient(client *http.Client) *UpdateAMQPTargetConfigParams {
	return &UpdateAMQPTargetConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateAMQPTargetConfigParams contains all the parameters to send to the API endpoint

	for the update a m q p target config operation.

	Typically these are written to a http.Request.
*/
type UpdateAMQPTargetConfigParams struct {

	/* Body.

	   Specifies the parameters to update cluster AMQP target config.
	*/
	Body *models.ClusterAMQPTargetConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update a m q p target config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAMQPTargetConfigParams) WithDefaults() *UpdateAMQPTargetConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update a m q p target config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAMQPTargetConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) WithTimeout(timeout time.Duration) *UpdateAMQPTargetConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) WithContext(ctx context.Context) *UpdateAMQPTargetConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) WithHTTPClient(client *http.Client) *UpdateAMQPTargetConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) WithBody(body *models.ClusterAMQPTargetConfig) *UpdateAMQPTargetConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update a m q p target config params
func (o *UpdateAMQPTargetConfigParams) SetBody(body *models.ClusterAMQPTargetConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAMQPTargetConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
