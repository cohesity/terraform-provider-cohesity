// Code generated by go-swagger; DO NOT EDIT.

package protection_group

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

// NewGetRunDebugLogsForObjectParams creates a new GetRunDebugLogsForObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunDebugLogsForObjectParams() *GetRunDebugLogsForObjectParams {
	return &GetRunDebugLogsForObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunDebugLogsForObjectParamsWithTimeout creates a new GetRunDebugLogsForObjectParams object
// with the ability to set a timeout on a request.
func NewGetRunDebugLogsForObjectParamsWithTimeout(timeout time.Duration) *GetRunDebugLogsForObjectParams {
	return &GetRunDebugLogsForObjectParams{
		timeout: timeout,
	}
}

// NewGetRunDebugLogsForObjectParamsWithContext creates a new GetRunDebugLogsForObjectParams object
// with the ability to set a context for a request.
func NewGetRunDebugLogsForObjectParamsWithContext(ctx context.Context) *GetRunDebugLogsForObjectParams {
	return &GetRunDebugLogsForObjectParams{
		Context: ctx,
	}
}

// NewGetRunDebugLogsForObjectParamsWithHTTPClient creates a new GetRunDebugLogsForObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunDebugLogsForObjectParamsWithHTTPClient(client *http.Client) *GetRunDebugLogsForObjectParams {
	return &GetRunDebugLogsForObjectParams{
		HTTPClient: client,
	}
}

/*
GetRunDebugLogsForObjectParams contains all the parameters to send to the API endpoint

	for the get run debug logs for object operation.

	Typically these are written to a http.Request.
*/
type GetRunDebugLogsForObjectParams struct {

	/* ID.

	   Specifies a unique id of the Protection Group.
	*/
	ID string

	/* ObjectID.

	   Specifies the id of the object for which debug logs are to be returned.
	*/
	ObjectID string

	/* RunID.

	   Specifies a unique run id of the Protection Group run.
	*/
	RunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get run debug logs for object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunDebugLogsForObjectParams) WithDefaults() *GetRunDebugLogsForObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run debug logs for object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunDebugLogsForObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) WithTimeout(timeout time.Duration) *GetRunDebugLogsForObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) WithContext(ctx context.Context) *GetRunDebugLogsForObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) WithHTTPClient(client *http.Client) *GetRunDebugLogsForObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) WithID(id string) *GetRunDebugLogsForObjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) SetID(id string) {
	o.ID = id
}

// WithObjectID adds the objectID to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) WithObjectID(objectID string) *GetRunDebugLogsForObjectParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) SetObjectID(objectID string) {
	o.ObjectID = objectID
}

// WithRunID adds the runID to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) WithRunID(runID string) *GetRunDebugLogsForObjectParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get run debug logs for object params
func (o *GetRunDebugLogsForObjectParams) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunDebugLogsForObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param objectId
	if err := r.SetPathParam("objectId", o.ObjectID); err != nil {
		return err
	}

	// path param runId
	if err := r.SetPathParam("runId", o.RunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
