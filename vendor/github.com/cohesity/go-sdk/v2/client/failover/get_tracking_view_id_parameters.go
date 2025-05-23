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
	"github.com/go-openapi/swag"
)

// NewGetTrackingViewIDParams creates a new GetTrackingViewIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTrackingViewIDParams() *GetTrackingViewIDParams {
	return &GetTrackingViewIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrackingViewIDParamsWithTimeout creates a new GetTrackingViewIDParams object
// with the ability to set a timeout on a request.
func NewGetTrackingViewIDParamsWithTimeout(timeout time.Duration) *GetTrackingViewIDParams {
	return &GetTrackingViewIDParams{
		timeout: timeout,
	}
}

// NewGetTrackingViewIDParamsWithContext creates a new GetTrackingViewIDParams object
// with the ability to set a context for a request.
func NewGetTrackingViewIDParamsWithContext(ctx context.Context) *GetTrackingViewIDParams {
	return &GetTrackingViewIDParams{
		Context: ctx,
	}
}

// NewGetTrackingViewIDParamsWithHTTPClient creates a new GetTrackingViewIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTrackingViewIDParamsWithHTTPClient(client *http.Client) *GetTrackingViewIDParams {
	return &GetTrackingViewIDParams{
		HTTPClient: client,
	}
}

/*
GetTrackingViewIDParams contains all the parameters to send to the API endpoint

	for the get tracking view Id operation.

	Typically these are written to a http.Request.
*/
type GetTrackingViewIDParams struct {

	/* ID.

	   Specifies the view_uid of the source view.
	*/
	ID string

	/* IsForwarded.

	   Indicates whether the request is forwarded
	*/
	IsForwarded *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tracking view Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrackingViewIDParams) WithDefaults() *GetTrackingViewIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tracking view Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrackingViewIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tracking view Id params
func (o *GetTrackingViewIDParams) WithTimeout(timeout time.Duration) *GetTrackingViewIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tracking view Id params
func (o *GetTrackingViewIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tracking view Id params
func (o *GetTrackingViewIDParams) WithContext(ctx context.Context) *GetTrackingViewIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tracking view Id params
func (o *GetTrackingViewIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tracking view Id params
func (o *GetTrackingViewIDParams) WithHTTPClient(client *http.Client) *GetTrackingViewIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tracking view Id params
func (o *GetTrackingViewIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get tracking view Id params
func (o *GetTrackingViewIDParams) WithID(id string) *GetTrackingViewIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get tracking view Id params
func (o *GetTrackingViewIDParams) SetID(id string) {
	o.ID = id
}

// WithIsForwarded adds the isForwarded to the get tracking view Id params
func (o *GetTrackingViewIDParams) WithIsForwarded(isForwarded *bool) *GetTrackingViewIDParams {
	o.SetIsForwarded(isForwarded)
	return o
}

// SetIsForwarded adds the isForwarded to the get tracking view Id params
func (o *GetTrackingViewIDParams) SetIsForwarded(isForwarded *bool) {
	o.IsForwarded = isForwarded
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrackingViewIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IsForwarded != nil {

		// query param isForwarded
		var qrIsForwarded bool

		if o.IsForwarded != nil {
			qrIsForwarded = *o.IsForwarded
		}
		qIsForwarded := swag.FormatBool(qrIsForwarded)
		if qIsForwarded != "" {

			if err := r.SetQueryParam("isForwarded", qIsForwarded); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
