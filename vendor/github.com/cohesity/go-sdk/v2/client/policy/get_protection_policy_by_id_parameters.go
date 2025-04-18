// Code generated by go-swagger; DO NOT EDIT.

package policy

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
)

// NewGetProtectionPolicyByIDParams creates a new GetProtectionPolicyByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProtectionPolicyByIDParams() *GetProtectionPolicyByIDParams {
	return &GetProtectionPolicyByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtectionPolicyByIDParamsWithTimeout creates a new GetProtectionPolicyByIDParams object
// with the ability to set a timeout on a request.
func NewGetProtectionPolicyByIDParamsWithTimeout(timeout time.Duration) *GetProtectionPolicyByIDParams {
	return &GetProtectionPolicyByIDParams{
		timeout: timeout,
	}
}

// NewGetProtectionPolicyByIDParamsWithContext creates a new GetProtectionPolicyByIDParams object
// with the ability to set a context for a request.
func NewGetProtectionPolicyByIDParamsWithContext(ctx context.Context) *GetProtectionPolicyByIDParams {
	return &GetProtectionPolicyByIDParams{
		Context: ctx,
	}
}

// NewGetProtectionPolicyByIDParamsWithHTTPClient creates a new GetProtectionPolicyByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProtectionPolicyByIDParamsWithHTTPClient(client *http.Client) *GetProtectionPolicyByIDParams {
	return &GetProtectionPolicyByIDParams{
		HTTPClient: client,
	}
}

/*
GetProtectionPolicyByIDParams contains all the parameters to send to the API endpoint

	for the get protection policy by Id operation.

	Typically these are written to a http.Request.
*/
type GetProtectionPolicyByIDParams struct {

	/* ID.

	   Specifies a unique id of the Protection Policy to return.
	*/
	ID string

	/* RequestInitiatorType.

	   Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
	*/
	RequestInitiatorType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get protection policy by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionPolicyByIDParams) WithDefaults() *GetProtectionPolicyByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get protection policy by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionPolicyByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) WithTimeout(timeout time.Duration) *GetProtectionPolicyByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) WithContext(ctx context.Context) *GetProtectionPolicyByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) WithHTTPClient(client *http.Client) *GetProtectionPolicyByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) WithID(id string) *GetProtectionPolicyByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) SetID(id string) {
	o.ID = id
}

// WithRequestInitiatorType adds the requestInitiatorType to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) WithRequestInitiatorType(requestInitiatorType *string) *GetProtectionPolicyByIDParams {
	o.SetRequestInitiatorType(requestInitiatorType)
	return o
}

// SetRequestInitiatorType adds the requestInitiatorType to the get protection policy by Id params
func (o *GetProtectionPolicyByIDParams) SetRequestInitiatorType(requestInitiatorType *string) {
	o.RequestInitiatorType = requestInitiatorType
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtectionPolicyByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
