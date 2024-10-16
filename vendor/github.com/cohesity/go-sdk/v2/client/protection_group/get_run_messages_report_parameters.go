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

// NewGetRunMessagesReportParams creates a new GetRunMessagesReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunMessagesReportParams() *GetRunMessagesReportParams {
	return &GetRunMessagesReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunMessagesReportParamsWithTimeout creates a new GetRunMessagesReportParams object
// with the ability to set a timeout on a request.
func NewGetRunMessagesReportParamsWithTimeout(timeout time.Duration) *GetRunMessagesReportParams {
	return &GetRunMessagesReportParams{
		timeout: timeout,
	}
}

// NewGetRunMessagesReportParamsWithContext creates a new GetRunMessagesReportParams object
// with the ability to set a context for a request.
func NewGetRunMessagesReportParamsWithContext(ctx context.Context) *GetRunMessagesReportParams {
	return &GetRunMessagesReportParams{
		Context: ctx,
	}
}

// NewGetRunMessagesReportParamsWithHTTPClient creates a new GetRunMessagesReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunMessagesReportParamsWithHTTPClient(client *http.Client) *GetRunMessagesReportParams {
	return &GetRunMessagesReportParams{
		HTTPClient: client,
	}
}

/*
GetRunMessagesReportParams contains all the parameters to send to the API endpoint

	for the get run messages report operation.

	Typically these are written to a http.Request.
*/
type GetRunMessagesReportParams struct {

	/* FileType.

	   Specifies the downloaded type, i.e: inclusion_exclusion_reports, error_files_list. default: error_files_list
	*/
	FileType *string

	/* ID.

	   Specifies a unique id of the Protection Group.
	*/
	ID string

	/* Name.

	   Specifies the name of the source being backed up
	*/
	Name *string

	/* ObjectID.

	   Specifies the id of the object for which errors/warnings are to be returned.
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

// WithDefaults hydrates default values in the get run messages report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunMessagesReportParams) WithDefaults() *GetRunMessagesReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get run messages report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunMessagesReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get run messages report params
func (o *GetRunMessagesReportParams) WithTimeout(timeout time.Duration) *GetRunMessagesReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run messages report params
func (o *GetRunMessagesReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run messages report params
func (o *GetRunMessagesReportParams) WithContext(ctx context.Context) *GetRunMessagesReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run messages report params
func (o *GetRunMessagesReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run messages report params
func (o *GetRunMessagesReportParams) WithHTTPClient(client *http.Client) *GetRunMessagesReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run messages report params
func (o *GetRunMessagesReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileType adds the fileType to the get run messages report params
func (o *GetRunMessagesReportParams) WithFileType(fileType *string) *GetRunMessagesReportParams {
	o.SetFileType(fileType)
	return o
}

// SetFileType adds the fileType to the get run messages report params
func (o *GetRunMessagesReportParams) SetFileType(fileType *string) {
	o.FileType = fileType
}

// WithID adds the id to the get run messages report params
func (o *GetRunMessagesReportParams) WithID(id string) *GetRunMessagesReportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get run messages report params
func (o *GetRunMessagesReportParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the get run messages report params
func (o *GetRunMessagesReportParams) WithName(name *string) *GetRunMessagesReportParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get run messages report params
func (o *GetRunMessagesReportParams) SetName(name *string) {
	o.Name = name
}

// WithObjectID adds the objectID to the get run messages report params
func (o *GetRunMessagesReportParams) WithObjectID(objectID string) *GetRunMessagesReportParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the get run messages report params
func (o *GetRunMessagesReportParams) SetObjectID(objectID string) {
	o.ObjectID = objectID
}

// WithRunID adds the runID to the get run messages report params
func (o *GetRunMessagesReportParams) WithRunID(runID string) *GetRunMessagesReportParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get run messages report params
func (o *GetRunMessagesReportParams) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunMessagesReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FileType != nil {

		// query param fileType
		var qrFileType string

		if o.FileType != nil {
			qrFileType = *o.FileType
		}
		qFileType := qrFileType
		if qFileType != "" {

			if err := r.SetQueryParam("fileType", qFileType); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
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
