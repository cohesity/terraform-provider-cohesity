// Code generated by go-swagger; DO NOT EDIT.

package recovery

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

// NewCreateRecoveryParams creates a new CreateRecoveryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRecoveryParams() *CreateRecoveryParams {
	return &CreateRecoveryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRecoveryParamsWithTimeout creates a new CreateRecoveryParams object
// with the ability to set a timeout on a request.
func NewCreateRecoveryParamsWithTimeout(timeout time.Duration) *CreateRecoveryParams {
	return &CreateRecoveryParams{
		timeout: timeout,
	}
}

// NewCreateRecoveryParamsWithContext creates a new CreateRecoveryParams object
// with the ability to set a context for a request.
func NewCreateRecoveryParamsWithContext(ctx context.Context) *CreateRecoveryParams {
	return &CreateRecoveryParams{
		Context: ctx,
	}
}

// NewCreateRecoveryParamsWithHTTPClient creates a new CreateRecoveryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRecoveryParamsWithHTTPClient(client *http.Client) *CreateRecoveryParams {
	return &CreateRecoveryParams{
		HTTPClient: client,
	}
}

/*
CreateRecoveryParams contains all the parameters to send to the API endpoint

	for the create recovery operation.

	Typically these are written to a http.Request.
*/
type CreateRecoveryParams struct {

	/* Body.

	   Specifies the parameters to create a Recovery.
	*/
	Body *models.CreateRecoveryRequest

	/* RequestInitiatorType.

	   Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
	*/
	RequestInitiatorType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create recovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRecoveryParams) WithDefaults() *CreateRecoveryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create recovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRecoveryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create recovery params
func (o *CreateRecoveryParams) WithTimeout(timeout time.Duration) *CreateRecoveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create recovery params
func (o *CreateRecoveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create recovery params
func (o *CreateRecoveryParams) WithContext(ctx context.Context) *CreateRecoveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create recovery params
func (o *CreateRecoveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create recovery params
func (o *CreateRecoveryParams) WithHTTPClient(client *http.Client) *CreateRecoveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create recovery params
func (o *CreateRecoveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create recovery params
func (o *CreateRecoveryParams) WithBody(body *models.CreateRecoveryRequest) *CreateRecoveryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create recovery params
func (o *CreateRecoveryParams) SetBody(body *models.CreateRecoveryRequest) {
	o.Body = body
}

// WithRequestInitiatorType adds the requestInitiatorType to the create recovery params
func (o *CreateRecoveryParams) WithRequestInitiatorType(requestInitiatorType *string) *CreateRecoveryParams {
	o.SetRequestInitiatorType(requestInitiatorType)
	return o
}

// SetRequestInitiatorType adds the requestInitiatorType to the create recovery params
func (o *CreateRecoveryParams) SetRequestInitiatorType(requestInitiatorType *string) {
	o.RequestInitiatorType = requestInitiatorType
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRecoveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.RequestInitiatorType != nil {

		// header param requestInitiatorType
		if err := r.SetHeaderParam("requestInitiatorType", *o.RequestInitiatorType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
