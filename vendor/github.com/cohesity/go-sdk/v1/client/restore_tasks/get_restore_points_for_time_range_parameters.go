// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

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

// NewGetRestorePointsForTimeRangeParams creates a new GetRestorePointsForTimeRangeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRestorePointsForTimeRangeParams() *GetRestorePointsForTimeRangeParams {
	return &GetRestorePointsForTimeRangeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestorePointsForTimeRangeParamsWithTimeout creates a new GetRestorePointsForTimeRangeParams object
// with the ability to set a timeout on a request.
func NewGetRestorePointsForTimeRangeParamsWithTimeout(timeout time.Duration) *GetRestorePointsForTimeRangeParams {
	return &GetRestorePointsForTimeRangeParams{
		timeout: timeout,
	}
}

// NewGetRestorePointsForTimeRangeParamsWithContext creates a new GetRestorePointsForTimeRangeParams object
// with the ability to set a context for a request.
func NewGetRestorePointsForTimeRangeParamsWithContext(ctx context.Context) *GetRestorePointsForTimeRangeParams {
	return &GetRestorePointsForTimeRangeParams{
		Context: ctx,
	}
}

// NewGetRestorePointsForTimeRangeParamsWithHTTPClient creates a new GetRestorePointsForTimeRangeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRestorePointsForTimeRangeParamsWithHTTPClient(client *http.Client) *GetRestorePointsForTimeRangeParams {
	return &GetRestorePointsForTimeRangeParams{
		HTTPClient: client,
	}
}

/*
GetRestorePointsForTimeRangeParams contains all the parameters to send to the API endpoint

	for the get restore points for time range operation.

	Typically these are written to a http.Request.
*/
type GetRestorePointsForTimeRangeParams struct {

	// Body.
	Body *models.RestorePointsForTimeRangeParam

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get restore points for time range params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRestorePointsForTimeRangeParams) WithDefaults() *GetRestorePointsForTimeRangeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get restore points for time range params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRestorePointsForTimeRangeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) WithTimeout(timeout time.Duration) *GetRestorePointsForTimeRangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) WithContext(ctx context.Context) *GetRestorePointsForTimeRangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) WithHTTPClient(client *http.Client) *GetRestorePointsForTimeRangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) WithBody(body *models.RestorePointsForTimeRangeParam) *GetRestorePointsForTimeRangeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get restore points for time range params
func (o *GetRestorePointsForTimeRangeParams) SetBody(body *models.RestorePointsForTimeRangeParam) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestorePointsForTimeRangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
