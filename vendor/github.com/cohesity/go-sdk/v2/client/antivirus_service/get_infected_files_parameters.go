// Code generated by go-swagger; DO NOT EDIT.

package antivirus_service

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

// NewGetInfectedFilesParams creates a new GetInfectedFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInfectedFilesParams() *GetInfectedFilesParams {
	return &GetInfectedFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInfectedFilesParamsWithTimeout creates a new GetInfectedFilesParams object
// with the ability to set a timeout on a request.
func NewGetInfectedFilesParamsWithTimeout(timeout time.Duration) *GetInfectedFilesParams {
	return &GetInfectedFilesParams{
		timeout: timeout,
	}
}

// NewGetInfectedFilesParamsWithContext creates a new GetInfectedFilesParams object
// with the ability to set a context for a request.
func NewGetInfectedFilesParamsWithContext(ctx context.Context) *GetInfectedFilesParams {
	return &GetInfectedFilesParams{
		Context: ctx,
	}
}

// NewGetInfectedFilesParamsWithHTTPClient creates a new GetInfectedFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInfectedFilesParamsWithHTTPClient(client *http.Client) *GetInfectedFilesParams {
	return &GetInfectedFilesParams{
		HTTPClient: client,
	}
}

/*
GetInfectedFilesParams contains all the parameters to send to the API endpoint

	for the get infected files operation.

	Typically these are written to a http.Request.
*/
type GetInfectedFilesParams struct {

	/* Cookie.

	   Specifies the pagination cookie.
	*/
	Cookie *string

	/* MaxCount.

	   Specifies the max number of files to be returned.

	   Format: int64
	*/
	MaxCount *int64

	/* Path.

	   Specifies the file path.
	*/
	Path *string

	/* States.

	   Specifies the file states.
	*/
	States []string

	/* ViewIds.

	   Specifies a list of view ids. Only infected entities from these views will be returned.
	*/
	ViewIds []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get infected files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInfectedFilesParams) WithDefaults() *GetInfectedFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get infected files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInfectedFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get infected files params
func (o *GetInfectedFilesParams) WithTimeout(timeout time.Duration) *GetInfectedFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get infected files params
func (o *GetInfectedFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get infected files params
func (o *GetInfectedFilesParams) WithContext(ctx context.Context) *GetInfectedFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get infected files params
func (o *GetInfectedFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get infected files params
func (o *GetInfectedFilesParams) WithHTTPClient(client *http.Client) *GetInfectedFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get infected files params
func (o *GetInfectedFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCookie adds the cookie to the get infected files params
func (o *GetInfectedFilesParams) WithCookie(cookie *string) *GetInfectedFilesParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the get infected files params
func (o *GetInfectedFilesParams) SetCookie(cookie *string) {
	o.Cookie = cookie
}

// WithMaxCount adds the maxCount to the get infected files params
func (o *GetInfectedFilesParams) WithMaxCount(maxCount *int64) *GetInfectedFilesParams {
	o.SetMaxCount(maxCount)
	return o
}

// SetMaxCount adds the maxCount to the get infected files params
func (o *GetInfectedFilesParams) SetMaxCount(maxCount *int64) {
	o.MaxCount = maxCount
}

// WithPath adds the path to the get infected files params
func (o *GetInfectedFilesParams) WithPath(path *string) *GetInfectedFilesParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get infected files params
func (o *GetInfectedFilesParams) SetPath(path *string) {
	o.Path = path
}

// WithStates adds the states to the get infected files params
func (o *GetInfectedFilesParams) WithStates(states []string) *GetInfectedFilesParams {
	o.SetStates(states)
	return o
}

// SetStates adds the states to the get infected files params
func (o *GetInfectedFilesParams) SetStates(states []string) {
	o.States = states
}

// WithViewIds adds the viewIds to the get infected files params
func (o *GetInfectedFilesParams) WithViewIds(viewIds []int64) *GetInfectedFilesParams {
	o.SetViewIds(viewIds)
	return o
}

// SetViewIds adds the viewIds to the get infected files params
func (o *GetInfectedFilesParams) SetViewIds(viewIds []int64) {
	o.ViewIds = viewIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetInfectedFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cookie != nil {

		// query param cookie
		var qrCookie string

		if o.Cookie != nil {
			qrCookie = *o.Cookie
		}
		qCookie := qrCookie
		if qCookie != "" {

			if err := r.SetQueryParam("cookie", qCookie); err != nil {
				return err
			}
		}
	}

	if o.MaxCount != nil {

		// query param maxCount
		var qrMaxCount int64

		if o.MaxCount != nil {
			qrMaxCount = *o.MaxCount
		}
		qMaxCount := swag.FormatInt64(qrMaxCount)
		if qMaxCount != "" {

			if err := r.SetQueryParam("maxCount", qMaxCount); err != nil {
				return err
			}
		}
	}

	if o.Path != nil {

		// query param path
		var qrPath string

		if o.Path != nil {
			qrPath = *o.Path
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	if o.States != nil {

		// binding items for states
		joinedStates := o.bindParamStates(reg)

		// query array param states
		if err := r.SetQueryParam("states", joinedStates...); err != nil {
			return err
		}
	}

	if o.ViewIds != nil {

		// binding items for viewIds
		joinedViewIds := o.bindParamViewIds(reg)

		// query array param viewIds
		if err := r.SetQueryParam("viewIds", joinedViewIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetInfectedFiles binds the parameter states
func (o *GetInfectedFilesParams) bindParamStates(formats strfmt.Registry) []string {
	statesIR := o.States

	var statesIC []string
	for _, statesIIR := range statesIR { // explode []string

		statesIIV := statesIIR // string as string
		statesIC = append(statesIC, statesIIV)
	}

	// items.CollectionFormat: ""
	statesIS := swag.JoinByFormat(statesIC, "")

	return statesIS
}

// bindParamGetInfectedFiles binds the parameter viewIds
func (o *GetInfectedFilesParams) bindParamViewIds(formats strfmt.Registry) []string {
	viewIdsIR := o.ViewIds

	var viewIdsIC []string
	for _, viewIdsIIR := range viewIdsIR { // explode []int64

		viewIdsIIV := swag.FormatInt64(viewIdsIIR) // int64 as string
		viewIdsIC = append(viewIdsIC, viewIdsIIV)
	}

	// items.CollectionFormat: ""
	viewIdsIS := swag.JoinByFormat(viewIdsIC, "")

	return viewIdsIS
}
