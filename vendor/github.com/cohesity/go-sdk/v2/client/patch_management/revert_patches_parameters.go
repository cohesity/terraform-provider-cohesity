// Code generated by go-swagger; DO NOT EDIT.

package patch_management

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

// NewRevertPatchesParams creates a new RevertPatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevertPatchesParams() *RevertPatchesParams {
	return &RevertPatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevertPatchesParamsWithTimeout creates a new RevertPatchesParams object
// with the ability to set a timeout on a request.
func NewRevertPatchesParamsWithTimeout(timeout time.Duration) *RevertPatchesParams {
	return &RevertPatchesParams{
		timeout: timeout,
	}
}

// NewRevertPatchesParamsWithContext creates a new RevertPatchesParams object
// with the ability to set a context for a request.
func NewRevertPatchesParamsWithContext(ctx context.Context) *RevertPatchesParams {
	return &RevertPatchesParams{
		Context: ctx,
	}
}

// NewRevertPatchesParamsWithHTTPClient creates a new RevertPatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevertPatchesParamsWithHTTPClient(client *http.Client) *RevertPatchesParams {
	return &RevertPatchesParams{
		HTTPClient: client,
	}
}

/*
RevertPatchesParams contains all the parameters to send to the API endpoint

	for the revert patches operation.

	Typically these are written to a http.Request.
*/
type RevertPatchesParams struct {

	/* Body.

	   Request to revert patches.
	*/
	Body *models.RevertPatchesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revert patches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevertPatchesParams) WithDefaults() *RevertPatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revert patches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevertPatchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revert patches params
func (o *RevertPatchesParams) WithTimeout(timeout time.Duration) *RevertPatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revert patches params
func (o *RevertPatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revert patches params
func (o *RevertPatchesParams) WithContext(ctx context.Context) *RevertPatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revert patches params
func (o *RevertPatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revert patches params
func (o *RevertPatchesParams) WithHTTPClient(client *http.Client) *RevertPatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revert patches params
func (o *RevertPatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the revert patches params
func (o *RevertPatchesParams) WithBody(body *models.RevertPatchesRequest) *RevertPatchesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the revert patches params
func (o *RevertPatchesParams) SetBody(body *models.RevertPatchesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RevertPatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
