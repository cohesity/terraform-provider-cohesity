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
)

// NewDeleteDataTieringAnalysisGroupParams creates a new DeleteDataTieringAnalysisGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDataTieringAnalysisGroupParams() *DeleteDataTieringAnalysisGroupParams {
	return &DeleteDataTieringAnalysisGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDataTieringAnalysisGroupParamsWithTimeout creates a new DeleteDataTieringAnalysisGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteDataTieringAnalysisGroupParamsWithTimeout(timeout time.Duration) *DeleteDataTieringAnalysisGroupParams {
	return &DeleteDataTieringAnalysisGroupParams{
		timeout: timeout,
	}
}

// NewDeleteDataTieringAnalysisGroupParamsWithContext creates a new DeleteDataTieringAnalysisGroupParams object
// with the ability to set a context for a request.
func NewDeleteDataTieringAnalysisGroupParamsWithContext(ctx context.Context) *DeleteDataTieringAnalysisGroupParams {
	return &DeleteDataTieringAnalysisGroupParams{
		Context: ctx,
	}
}

// NewDeleteDataTieringAnalysisGroupParamsWithHTTPClient creates a new DeleteDataTieringAnalysisGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDataTieringAnalysisGroupParamsWithHTTPClient(client *http.Client) *DeleteDataTieringAnalysisGroupParams {
	return &DeleteDataTieringAnalysisGroupParams{
		HTTPClient: client,
	}
}

/*
DeleteDataTieringAnalysisGroupParams contains all the parameters to send to the API endpoint

	for the delete data tiering analysis group operation.

	Typically these are written to a http.Request.
*/
type DeleteDataTieringAnalysisGroupParams struct {

	/* ID.

	   Specifies a unique id of the data tiering analysis group.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete data tiering analysis group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataTieringAnalysisGroupParams) WithDefaults() *DeleteDataTieringAnalysisGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete data tiering analysis group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataTieringAnalysisGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) WithTimeout(timeout time.Duration) *DeleteDataTieringAnalysisGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) WithContext(ctx context.Context) *DeleteDataTieringAnalysisGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) WithHTTPClient(client *http.Client) *DeleteDataTieringAnalysisGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) WithID(id string) *DeleteDataTieringAnalysisGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete data tiering analysis group params
func (o *DeleteDataTieringAnalysisGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDataTieringAnalysisGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
