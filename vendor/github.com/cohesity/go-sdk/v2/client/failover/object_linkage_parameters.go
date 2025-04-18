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

// NewObjectLinkageParams creates a new ObjectLinkageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewObjectLinkageParams() *ObjectLinkageParams {
	return &ObjectLinkageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewObjectLinkageParamsWithTimeout creates a new ObjectLinkageParams object
// with the ability to set a timeout on a request.
func NewObjectLinkageParamsWithTimeout(timeout time.Duration) *ObjectLinkageParams {
	return &ObjectLinkageParams{
		timeout: timeout,
	}
}

// NewObjectLinkageParamsWithContext creates a new ObjectLinkageParams object
// with the ability to set a context for a request.
func NewObjectLinkageParamsWithContext(ctx context.Context) *ObjectLinkageParams {
	return &ObjectLinkageParams{
		Context: ctx,
	}
}

// NewObjectLinkageParamsWithHTTPClient creates a new ObjectLinkageParams object
// with the ability to set a custom HTTPClient for a request.
func NewObjectLinkageParamsWithHTTPClient(client *http.Client) *ObjectLinkageParams {
	return &ObjectLinkageParams{
		HTTPClient: client,
	}
}

/*
ObjectLinkageParams contains all the parameters to send to the API endpoint

	for the object linkage operation.

	Typically these are written to a http.Request.
*/
type ObjectLinkageParams struct {

	/* Body.

	   Specifies the paramteres to create links between replicated objects and failover objects.
	*/
	Body *models.ObjectLinkingRequest

	/* ID.

	   Specifies the id of the failover workflow.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the object linkage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectLinkageParams) WithDefaults() *ObjectLinkageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the object linkage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectLinkageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the object linkage params
func (o *ObjectLinkageParams) WithTimeout(timeout time.Duration) *ObjectLinkageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the object linkage params
func (o *ObjectLinkageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the object linkage params
func (o *ObjectLinkageParams) WithContext(ctx context.Context) *ObjectLinkageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the object linkage params
func (o *ObjectLinkageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the object linkage params
func (o *ObjectLinkageParams) WithHTTPClient(client *http.Client) *ObjectLinkageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the object linkage params
func (o *ObjectLinkageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the object linkage params
func (o *ObjectLinkageParams) WithBody(body *models.ObjectLinkingRequest) *ObjectLinkageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the object linkage params
func (o *ObjectLinkageParams) SetBody(body *models.ObjectLinkingRequest) {
	o.Body = body
}

// WithID adds the id to the object linkage params
func (o *ObjectLinkageParams) WithID(id string) *ObjectLinkageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the object linkage params
func (o *ObjectLinkageParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectLinkageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
