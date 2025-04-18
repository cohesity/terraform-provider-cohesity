// Code generated by go-swagger; DO NOT EDIT.

package data_tiering

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

// NewUpdateDataTieringAnalysisGroupTagsConfigParams creates a new UpdateDataTieringAnalysisGroupTagsConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDataTieringAnalysisGroupTagsConfigParams() *UpdateDataTieringAnalysisGroupTagsConfigParams {
	return &UpdateDataTieringAnalysisGroupTagsConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDataTieringAnalysisGroupTagsConfigParamsWithTimeout creates a new UpdateDataTieringAnalysisGroupTagsConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateDataTieringAnalysisGroupTagsConfigParamsWithTimeout(timeout time.Duration) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	return &UpdateDataTieringAnalysisGroupTagsConfigParams{
		timeout: timeout,
	}
}

// NewUpdateDataTieringAnalysisGroupTagsConfigParamsWithContext creates a new UpdateDataTieringAnalysisGroupTagsConfigParams object
// with the ability to set a context for a request.
func NewUpdateDataTieringAnalysisGroupTagsConfigParamsWithContext(ctx context.Context) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	return &UpdateDataTieringAnalysisGroupTagsConfigParams{
		Context: ctx,
	}
}

// NewUpdateDataTieringAnalysisGroupTagsConfigParamsWithHTTPClient creates a new UpdateDataTieringAnalysisGroupTagsConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDataTieringAnalysisGroupTagsConfigParamsWithHTTPClient(client *http.Client) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	return &UpdateDataTieringAnalysisGroupTagsConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateDataTieringAnalysisGroupTagsConfigParams contains all the parameters to send to the API endpoint

	for the update data tiering analysis group tags config operation.

	Typically these are written to a http.Request.
*/
type UpdateDataTieringAnalysisGroupTagsConfigParams struct {

	/* Body.

	   Specifies the data tiering analysis Tags Config.
	*/
	Body *models.DataTieringTagConfig

	/* ID.

	   Specifies a unique id of the data tiering analysis group.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update data tiering analysis group tags config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) WithDefaults() *UpdateDataTieringAnalysisGroupTagsConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update data tiering analysis group tags config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) WithTimeout(timeout time.Duration) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) WithContext(ctx context.Context) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) WithHTTPClient(client *http.Client) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) WithBody(body *models.DataTieringTagConfig) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) SetBody(body *models.DataTieringTagConfig) {
	o.Body = body
}

// WithID adds the id to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) WithID(id string) *UpdateDataTieringAnalysisGroupTagsConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update data tiering analysis group tags config params
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDataTieringAnalysisGroupTagsConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
