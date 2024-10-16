// Code generated by go-swagger; DO NOT EDIT.

package stats

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

// NewGetProtectedObjectsSummaryParams creates a new GetProtectedObjectsSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProtectedObjectsSummaryParams() *GetProtectedObjectsSummaryParams {
	return &GetProtectedObjectsSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtectedObjectsSummaryParamsWithTimeout creates a new GetProtectedObjectsSummaryParams object
// with the ability to set a timeout on a request.
func NewGetProtectedObjectsSummaryParamsWithTimeout(timeout time.Duration) *GetProtectedObjectsSummaryParams {
	return &GetProtectedObjectsSummaryParams{
		timeout: timeout,
	}
}

// NewGetProtectedObjectsSummaryParamsWithContext creates a new GetProtectedObjectsSummaryParams object
// with the ability to set a context for a request.
func NewGetProtectedObjectsSummaryParamsWithContext(ctx context.Context) *GetProtectedObjectsSummaryParams {
	return &GetProtectedObjectsSummaryParams{
		Context: ctx,
	}
}

// NewGetProtectedObjectsSummaryParamsWithHTTPClient creates a new GetProtectedObjectsSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProtectedObjectsSummaryParamsWithHTTPClient(client *http.Client) *GetProtectedObjectsSummaryParams {
	return &GetProtectedObjectsSummaryParams{
		HTTPClient: client,
	}
}

/*
GetProtectedObjectsSummaryParams contains all the parameters to send to the API endpoint

	for the get protected objects summary operation.

	Typically these are written to a http.Request.
*/
type GetProtectedObjectsSummaryParams struct {

	/* ExcludeTypes.

	   Specifies the environment types to exclude.
	*/
	ExcludeTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get protected objects summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectedObjectsSummaryParams) WithDefaults() *GetProtectedObjectsSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get protected objects summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectedObjectsSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) WithTimeout(timeout time.Duration) *GetProtectedObjectsSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) WithContext(ctx context.Context) *GetProtectedObjectsSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) WithHTTPClient(client *http.Client) *GetProtectedObjectsSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeTypes adds the excludeTypes to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) WithExcludeTypes(excludeTypes []string) *GetProtectedObjectsSummaryParams {
	o.SetExcludeTypes(excludeTypes)
	return o
}

// SetExcludeTypes adds the excludeTypes to the get protected objects summary params
func (o *GetProtectedObjectsSummaryParams) SetExcludeTypes(excludeTypes []string) {
	o.ExcludeTypes = excludeTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtectedObjectsSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeTypes != nil {

		// binding items for excludeTypes
		joinedExcludeTypes := o.bindParamExcludeTypes(reg)

		// query array param excludeTypes
		if err := r.SetQueryParam("excludeTypes", joinedExcludeTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetProtectedObjectsSummary binds the parameter excludeTypes
func (o *GetProtectedObjectsSummaryParams) bindParamExcludeTypes(formats strfmt.Registry) []string {
	excludeTypesIR := o.ExcludeTypes

	var excludeTypesIC []string
	for _, excludeTypesIIR := range excludeTypesIR { // explode []string

		excludeTypesIIV := excludeTypesIIR // string as string
		excludeTypesIC = append(excludeTypesIC, excludeTypesIIV)
	}

	// items.CollectionFormat: ""
	excludeTypesIS := swag.JoinByFormat(excludeTypesIC, "")

	return excludeTypesIS
}
