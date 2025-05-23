// Code generated by go-swagger; DO NOT EDIT.

package protected_object

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
	"github.com/go-openapi/swag"

	"github.com/cohesity/go-sdk/v2/models"
)

// NewUpdateProtectedObjectsOfAnyTypeParams creates a new UpdateProtectedObjectsOfAnyTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProtectedObjectsOfAnyTypeParams() *UpdateProtectedObjectsOfAnyTypeParams {
	return &UpdateProtectedObjectsOfAnyTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProtectedObjectsOfAnyTypeParamsWithTimeout creates a new UpdateProtectedObjectsOfAnyTypeParams object
// with the ability to set a timeout on a request.
func NewUpdateProtectedObjectsOfAnyTypeParamsWithTimeout(timeout time.Duration) *UpdateProtectedObjectsOfAnyTypeParams {
	return &UpdateProtectedObjectsOfAnyTypeParams{
		timeout: timeout,
	}
}

// NewUpdateProtectedObjectsOfAnyTypeParamsWithContext creates a new UpdateProtectedObjectsOfAnyTypeParams object
// with the ability to set a context for a request.
func NewUpdateProtectedObjectsOfAnyTypeParamsWithContext(ctx context.Context) *UpdateProtectedObjectsOfAnyTypeParams {
	return &UpdateProtectedObjectsOfAnyTypeParams{
		Context: ctx,
	}
}

// NewUpdateProtectedObjectsOfAnyTypeParamsWithHTTPClient creates a new UpdateProtectedObjectsOfAnyTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProtectedObjectsOfAnyTypeParamsWithHTTPClient(client *http.Client) *UpdateProtectedObjectsOfAnyTypeParams {
	return &UpdateProtectedObjectsOfAnyTypeParams{
		HTTPClient: client,
	}
}

/*
UpdateProtectedObjectsOfAnyTypeParams contains all the parameters to send to the API endpoint

	for the update protected objects of any type operation.

	Typically these are written to a http.Request.
*/
type UpdateProtectedObjectsOfAnyTypeParams struct {

	/* Body.

	   Specifies the parameters to perform an update on protected objects.
	*/
	Body *models.UpdateProtectedObjectsRequest

	/* ID.

	   Specifies the id of the Protected Object.

	   Format: int64
	*/
	ID int64

	/* RequestInitiatorType.

	   Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
	*/
	RequestInitiatorType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update protected objects of any type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProtectedObjectsOfAnyTypeParams) WithDefaults() *UpdateProtectedObjectsOfAnyTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update protected objects of any type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProtectedObjectsOfAnyTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) WithTimeout(timeout time.Duration) *UpdateProtectedObjectsOfAnyTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) WithContext(ctx context.Context) *UpdateProtectedObjectsOfAnyTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) WithHTTPClient(client *http.Client) *UpdateProtectedObjectsOfAnyTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) WithBody(body *models.UpdateProtectedObjectsRequest) *UpdateProtectedObjectsOfAnyTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) SetBody(body *models.UpdateProtectedObjectsRequest) {
	o.Body = body
}

// WithID adds the id to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) WithID(id int64) *UpdateProtectedObjectsOfAnyTypeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) SetID(id int64) {
	o.ID = id
}

// WithRequestInitiatorType adds the requestInitiatorType to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) WithRequestInitiatorType(requestInitiatorType *string) *UpdateProtectedObjectsOfAnyTypeParams {
	o.SetRequestInitiatorType(requestInitiatorType)
	return o
}

// SetRequestInitiatorType adds the requestInitiatorType to the update protected objects of any type params
func (o *UpdateProtectedObjectsOfAnyTypeParams) SetRequestInitiatorType(requestInitiatorType *string) {
	o.RequestInitiatorType = requestInitiatorType
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProtectedObjectsOfAnyTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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
