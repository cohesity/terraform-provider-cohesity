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
	"github.com/go-openapi/swag"

	"github.com/cohesity/go-sdk/v2/models"
)

// NewMarkDiskRemovalParams creates a new MarkDiskRemovalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMarkDiskRemovalParams() *MarkDiskRemovalParams {
	return &MarkDiskRemovalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMarkDiskRemovalParamsWithTimeout creates a new MarkDiskRemovalParams object
// with the ability to set a timeout on a request.
func NewMarkDiskRemovalParamsWithTimeout(timeout time.Duration) *MarkDiskRemovalParams {
	return &MarkDiskRemovalParams{
		timeout: timeout,
	}
}

// NewMarkDiskRemovalParamsWithContext creates a new MarkDiskRemovalParams object
// with the ability to set a context for a request.
func NewMarkDiskRemovalParamsWithContext(ctx context.Context) *MarkDiskRemovalParams {
	return &MarkDiskRemovalParams{
		Context: ctx,
	}
}

// NewMarkDiskRemovalParamsWithHTTPClient creates a new MarkDiskRemovalParams object
// with the ability to set a custom HTTPClient for a request.
func NewMarkDiskRemovalParamsWithHTTPClient(client *http.Client) *MarkDiskRemovalParams {
	return &MarkDiskRemovalParams{
		HTTPClient: client,
	}
}

/*
MarkDiskRemovalParams contains all the parameters to send to the API endpoint

	for the mark disk removal operation.

	Typically these are written to a http.Request.
*/
type MarkDiskRemovalParams struct {

	/* Body.

	   Specifies parameters to mark/cancel disk removal.
	*/
	Body *models.DiskRemovalParams

	/* ID.

	   Specifies unique id of the disk to mark for removal.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mark disk removal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MarkDiskRemovalParams) WithDefaults() *MarkDiskRemovalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mark disk removal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MarkDiskRemovalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mark disk removal params
func (o *MarkDiskRemovalParams) WithTimeout(timeout time.Duration) *MarkDiskRemovalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mark disk removal params
func (o *MarkDiskRemovalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mark disk removal params
func (o *MarkDiskRemovalParams) WithContext(ctx context.Context) *MarkDiskRemovalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mark disk removal params
func (o *MarkDiskRemovalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mark disk removal params
func (o *MarkDiskRemovalParams) WithHTTPClient(client *http.Client) *MarkDiskRemovalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mark disk removal params
func (o *MarkDiskRemovalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the mark disk removal params
func (o *MarkDiskRemovalParams) WithBody(body *models.DiskRemovalParams) *MarkDiskRemovalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the mark disk removal params
func (o *MarkDiskRemovalParams) SetBody(body *models.DiskRemovalParams) {
	o.Body = body
}

// WithID adds the id to the mark disk removal params
func (o *MarkDiskRemovalParams) WithID(id int64) *MarkDiskRemovalParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the mark disk removal params
func (o *MarkDiskRemovalParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MarkDiskRemovalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
