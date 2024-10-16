// Code generated by go-swagger; DO NOT EDIT.

package statistics

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

// NewGetTasksParams creates a new GetTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTasksParams() *GetTasksParams {
	return &GetTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTasksParamsWithTimeout creates a new GetTasksParams object
// with the ability to set a timeout on a request.
func NewGetTasksParamsWithTimeout(timeout time.Duration) *GetTasksParams {
	return &GetTasksParams{
		timeout: timeout,
	}
}

// NewGetTasksParamsWithContext creates a new GetTasksParams object
// with the ability to set a context for a request.
func NewGetTasksParamsWithContext(ctx context.Context) *GetTasksParams {
	return &GetTasksParams{
		Context: ctx,
	}
}

// NewGetTasksParamsWithHTTPClient creates a new GetTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTasksParamsWithHTTPClient(client *http.Client) *GetTasksParams {
	return &GetTasksParams{
		HTTPClient: client,
	}
}

/*
GetTasksParams contains all the parameters to send to the API endpoint

	for the get tasks operation.

	Typically these are written to a http.Request.
*/
type GetTasksParams struct {

	/* Attributes.

	     If specified, tasks matching the current query are further filtered by
	these KeyValuePairs. This gives client an ability to search by custom
	attributes that they specified during the task creation. Only the Tasks
	having 'all' of the specified key=value pairs will be returned.
	Attributes passed in should be separated by commas and each must follow
	the pattern "name:type:value". Valid types are "kInt64", "kDouble",
	"kString", and "kBytes".
	*/
	Attributes []string

	/* EndTimeSeconds.

	     Specifies an end time by which to filter tasks. Including a value
	here will result in tasks only being included if the ended before
	this time.

	     Format: int64
	*/
	EndTimeSeconds *int64

	/* ExcludeSubTasks.

	     Specifies whether or not to include subtasks. By default all subtasks
	of any task matching a query will be returned.
	*/
	ExcludeSubTasks *bool

	/* IncludeFinishedTasks.

	     Specifies whether or not to include finished tasks. By default,
	Pulse will only include unfinished tasks.
	*/
	IncludeFinishedTasks *bool

	/* MaxTasks.

	   Specifies the maximum number of tasks to display. Defaults to 100.

	   Format: int32
	*/
	MaxTasks *int32

	/* StartTimeSeconds.

	     Specifies a start time by which to filter tasks. Including a value
	here will result in tasks only being included if they started after
	the time specified.

	     Format: int64
	*/
	StartTimeSeconds *int64

	/* TaskPaths.

	     Specifies a list of tasks path by which to filter the results. Otherwise
	all of the tasks will be returned.
	*/
	TaskPaths []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksParams) WithDefaults() *GetTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) WithTimeout(timeout time.Duration) *GetTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tasks params
func (o *GetTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tasks params
func (o *GetTasksParams) WithContext(ctx context.Context) *GetTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tasks params
func (o *GetTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) WithHTTPClient(client *http.Client) *GetTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tasks params
func (o *GetTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the get tasks params
func (o *GetTasksParams) WithAttributes(attributes []string) *GetTasksParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get tasks params
func (o *GetTasksParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithEndTimeSeconds adds the endTimeSeconds to the get tasks params
func (o *GetTasksParams) WithEndTimeSeconds(endTimeSeconds *int64) *GetTasksParams {
	o.SetEndTimeSeconds(endTimeSeconds)
	return o
}

// SetEndTimeSeconds adds the endTimeSeconds to the get tasks params
func (o *GetTasksParams) SetEndTimeSeconds(endTimeSeconds *int64) {
	o.EndTimeSeconds = endTimeSeconds
}

// WithExcludeSubTasks adds the excludeSubTasks to the get tasks params
func (o *GetTasksParams) WithExcludeSubTasks(excludeSubTasks *bool) *GetTasksParams {
	o.SetExcludeSubTasks(excludeSubTasks)
	return o
}

// SetExcludeSubTasks adds the excludeSubTasks to the get tasks params
func (o *GetTasksParams) SetExcludeSubTasks(excludeSubTasks *bool) {
	o.ExcludeSubTasks = excludeSubTasks
}

// WithIncludeFinishedTasks adds the includeFinishedTasks to the get tasks params
func (o *GetTasksParams) WithIncludeFinishedTasks(includeFinishedTasks *bool) *GetTasksParams {
	o.SetIncludeFinishedTasks(includeFinishedTasks)
	return o
}

// SetIncludeFinishedTasks adds the includeFinishedTasks to the get tasks params
func (o *GetTasksParams) SetIncludeFinishedTasks(includeFinishedTasks *bool) {
	o.IncludeFinishedTasks = includeFinishedTasks
}

// WithMaxTasks adds the maxTasks to the get tasks params
func (o *GetTasksParams) WithMaxTasks(maxTasks *int32) *GetTasksParams {
	o.SetMaxTasks(maxTasks)
	return o
}

// SetMaxTasks adds the maxTasks to the get tasks params
func (o *GetTasksParams) SetMaxTasks(maxTasks *int32) {
	o.MaxTasks = maxTasks
}

// WithStartTimeSeconds adds the startTimeSeconds to the get tasks params
func (o *GetTasksParams) WithStartTimeSeconds(startTimeSeconds *int64) *GetTasksParams {
	o.SetStartTimeSeconds(startTimeSeconds)
	return o
}

// SetStartTimeSeconds adds the startTimeSeconds to the get tasks params
func (o *GetTasksParams) SetStartTimeSeconds(startTimeSeconds *int64) {
	o.StartTimeSeconds = startTimeSeconds
}

// WithTaskPaths adds the taskPaths to the get tasks params
func (o *GetTasksParams) WithTaskPaths(taskPaths []string) *GetTasksParams {
	o.SetTaskPaths(taskPaths)
	return o
}

// SetTaskPaths adds the taskPaths to the get tasks params
func (o *GetTasksParams) SetTaskPaths(taskPaths []string) {
	o.TaskPaths = taskPaths
}

// WriteToRequest writes these params to a swagger request
func (o *GetTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Attributes != nil {

		// binding items for attributes
		joinedAttributes := o.bindParamAttributes(reg)

		// query array param attributes
		if err := r.SetQueryParam("attributes", joinedAttributes...); err != nil {
			return err
		}
	}

	if o.EndTimeSeconds != nil {

		// query param endTimeSeconds
		var qrEndTimeSeconds int64

		if o.EndTimeSeconds != nil {
			qrEndTimeSeconds = *o.EndTimeSeconds
		}
		qEndTimeSeconds := swag.FormatInt64(qrEndTimeSeconds)
		if qEndTimeSeconds != "" {

			if err := r.SetQueryParam("endTimeSeconds", qEndTimeSeconds); err != nil {
				return err
			}
		}
	}

	if o.ExcludeSubTasks != nil {

		// query param excludeSubTasks
		var qrExcludeSubTasks bool

		if o.ExcludeSubTasks != nil {
			qrExcludeSubTasks = *o.ExcludeSubTasks
		}
		qExcludeSubTasks := swag.FormatBool(qrExcludeSubTasks)
		if qExcludeSubTasks != "" {

			if err := r.SetQueryParam("excludeSubTasks", qExcludeSubTasks); err != nil {
				return err
			}
		}
	}

	if o.IncludeFinishedTasks != nil {

		// query param includeFinishedTasks
		var qrIncludeFinishedTasks bool

		if o.IncludeFinishedTasks != nil {
			qrIncludeFinishedTasks = *o.IncludeFinishedTasks
		}
		qIncludeFinishedTasks := swag.FormatBool(qrIncludeFinishedTasks)
		if qIncludeFinishedTasks != "" {

			if err := r.SetQueryParam("includeFinishedTasks", qIncludeFinishedTasks); err != nil {
				return err
			}
		}
	}

	if o.MaxTasks != nil {

		// query param maxTasks
		var qrMaxTasks int32

		if o.MaxTasks != nil {
			qrMaxTasks = *o.MaxTasks
		}
		qMaxTasks := swag.FormatInt32(qrMaxTasks)
		if qMaxTasks != "" {

			if err := r.SetQueryParam("maxTasks", qMaxTasks); err != nil {
				return err
			}
		}
	}

	if o.StartTimeSeconds != nil {

		// query param startTimeSeconds
		var qrStartTimeSeconds int64

		if o.StartTimeSeconds != nil {
			qrStartTimeSeconds = *o.StartTimeSeconds
		}
		qStartTimeSeconds := swag.FormatInt64(qrStartTimeSeconds)
		if qStartTimeSeconds != "" {

			if err := r.SetQueryParam("startTimeSeconds", qStartTimeSeconds); err != nil {
				return err
			}
		}
	}

	if o.TaskPaths != nil {

		// binding items for taskPaths
		joinedTaskPaths := o.bindParamTaskPaths(reg)

		// query array param taskPaths
		if err := r.SetQueryParam("taskPaths", joinedTaskPaths...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetTasks binds the parameter attributes
func (o *GetTasksParams) bindParamAttributes(formats strfmt.Registry) []string {
	attributesIR := o.Attributes

	var attributesIC []string
	for _, attributesIIR := range attributesIR { // explode []string

		attributesIIV := attributesIIR // string as string
		attributesIC = append(attributesIC, attributesIIV)
	}

	// items.CollectionFormat: ""
	attributesIS := swag.JoinByFormat(attributesIC, "")

	return attributesIS
}

// bindParamGetTasks binds the parameter taskPaths
func (o *GetTasksParams) bindParamTaskPaths(formats strfmt.Registry) []string {
	taskPathsIR := o.TaskPaths

	var taskPathsIC []string
	for _, taskPathsIIR := range taskPathsIR { // explode []string

		taskPathsIIV := taskPathsIIR // string as string
		taskPathsIC = append(taskPathsIC, taskPathsIIV)
	}

	// items.CollectionFormat: ""
	taskPathsIS := swag.JoinByFormat(taskPathsIC, "")

	return taskPathsIS
}
