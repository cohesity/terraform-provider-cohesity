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

// NewDownloadTieringReportsParams creates a new DownloadTieringReportsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadTieringReportsParams() *DownloadTieringReportsParams {
	return &DownloadTieringReportsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadTieringReportsParamsWithTimeout creates a new DownloadTieringReportsParams object
// with the ability to set a timeout on a request.
func NewDownloadTieringReportsParamsWithTimeout(timeout time.Duration) *DownloadTieringReportsParams {
	return &DownloadTieringReportsParams{
		timeout: timeout,
	}
}

// NewDownloadTieringReportsParamsWithContext creates a new DownloadTieringReportsParams object
// with the ability to set a context for a request.
func NewDownloadTieringReportsParamsWithContext(ctx context.Context) *DownloadTieringReportsParams {
	return &DownloadTieringReportsParams{
		Context: ctx,
	}
}

// NewDownloadTieringReportsParamsWithHTTPClient creates a new DownloadTieringReportsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadTieringReportsParamsWithHTTPClient(client *http.Client) *DownloadTieringReportsParams {
	return &DownloadTieringReportsParams{
		HTTPClient: client,
	}
}

/*
DownloadTieringReportsParams contains all the parameters to send to the API endpoint

	for the download tiering reports operation.

	Typically these are written to a http.Request.
*/
type DownloadTieringReportsParams struct {

	/* FilePath.

	   Specifies the file path in the targetView.
	*/
	FilePath string

	/* ID.

	   Specifies a unique id of data tiering task.
	*/
	ID string

	/* RunID.

	   Specifies a unique run id of data tiering task.
	*/
	RunID string

	/* TargetViewName.

	     Specifies the View name from which the tiering job report
	file should be read from.
	*/
	TargetViewName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download tiering reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadTieringReportsParams) WithDefaults() *DownloadTieringReportsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download tiering reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadTieringReportsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download tiering reports params
func (o *DownloadTieringReportsParams) WithTimeout(timeout time.Duration) *DownloadTieringReportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download tiering reports params
func (o *DownloadTieringReportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download tiering reports params
func (o *DownloadTieringReportsParams) WithContext(ctx context.Context) *DownloadTieringReportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download tiering reports params
func (o *DownloadTieringReportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download tiering reports params
func (o *DownloadTieringReportsParams) WithHTTPClient(client *http.Client) *DownloadTieringReportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download tiering reports params
func (o *DownloadTieringReportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilePath adds the filePath to the download tiering reports params
func (o *DownloadTieringReportsParams) WithFilePath(filePath string) *DownloadTieringReportsParams {
	o.SetFilePath(filePath)
	return o
}

// SetFilePath adds the filePath to the download tiering reports params
func (o *DownloadTieringReportsParams) SetFilePath(filePath string) {
	o.FilePath = filePath
}

// WithID adds the id to the download tiering reports params
func (o *DownloadTieringReportsParams) WithID(id string) *DownloadTieringReportsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download tiering reports params
func (o *DownloadTieringReportsParams) SetID(id string) {
	o.ID = id
}

// WithRunID adds the runID to the download tiering reports params
func (o *DownloadTieringReportsParams) WithRunID(runID string) *DownloadTieringReportsParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the download tiering reports params
func (o *DownloadTieringReportsParams) SetRunID(runID string) {
	o.RunID = runID
}

// WithTargetViewName adds the targetViewName to the download tiering reports params
func (o *DownloadTieringReportsParams) WithTargetViewName(targetViewName string) *DownloadTieringReportsParams {
	o.SetTargetViewName(targetViewName)
	return o
}

// SetTargetViewName adds the targetViewName to the download tiering reports params
func (o *DownloadTieringReportsParams) SetTargetViewName(targetViewName string) {
	o.TargetViewName = targetViewName
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadTieringReportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param filePath
	qrFilePath := o.FilePath
	qFilePath := qrFilePath
	if qFilePath != "" {

		if err := r.SetQueryParam("filePath", qFilePath); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param runId
	if err := r.SetPathParam("runId", o.RunID); err != nil {
		return err
	}

	// query param targetViewName
	qrTargetViewName := o.TargetViewName
	qTargetViewName := qrTargetViewName
	if qTargetViewName != "" {

		if err := r.SetQueryParam("targetViewName", qTargetViewName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
