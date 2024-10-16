// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

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

	"github.com/cohesity/go-sdk/v1/models"
)

// NewSearchProductionAdObjectsParams creates a new SearchProductionAdObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchProductionAdObjectsParams() *SearchProductionAdObjectsParams {
	return &SearchProductionAdObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchProductionAdObjectsParamsWithTimeout creates a new SearchProductionAdObjectsParams object
// with the ability to set a timeout on a request.
func NewSearchProductionAdObjectsParamsWithTimeout(timeout time.Duration) *SearchProductionAdObjectsParams {
	return &SearchProductionAdObjectsParams{
		timeout: timeout,
	}
}

// NewSearchProductionAdObjectsParamsWithContext creates a new SearchProductionAdObjectsParams object
// with the ability to set a context for a request.
func NewSearchProductionAdObjectsParamsWithContext(ctx context.Context) *SearchProductionAdObjectsParams {
	return &SearchProductionAdObjectsParams{
		Context: ctx,
	}
}

// NewSearchProductionAdObjectsParamsWithHTTPClient creates a new SearchProductionAdObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchProductionAdObjectsParamsWithHTTPClient(client *http.Client) *SearchProductionAdObjectsParams {
	return &SearchProductionAdObjectsParams{
		HTTPClient: client,
	}
}

/*
SearchProductionAdObjectsParams contains all the parameters to send to the API endpoint

	for the search production ad objects operation.

	Typically these are written to a http.Request.
*/
type SearchProductionAdObjectsParams struct {

	/* Body.

	   Specifies the Request to search the AD Objects from Production AD.
	*/
	Body *models.SearchProductionAdObjectsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search production ad objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchProductionAdObjectsParams) WithDefaults() *SearchProductionAdObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search production ad objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchProductionAdObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search production ad objects params
func (o *SearchProductionAdObjectsParams) WithTimeout(timeout time.Duration) *SearchProductionAdObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search production ad objects params
func (o *SearchProductionAdObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search production ad objects params
func (o *SearchProductionAdObjectsParams) WithContext(ctx context.Context) *SearchProductionAdObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search production ad objects params
func (o *SearchProductionAdObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search production ad objects params
func (o *SearchProductionAdObjectsParams) WithHTTPClient(client *http.Client) *SearchProductionAdObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search production ad objects params
func (o *SearchProductionAdObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search production ad objects params
func (o *SearchProductionAdObjectsParams) WithBody(body *models.SearchProductionAdObjectsRequest) *SearchProductionAdObjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search production ad objects params
func (o *SearchProductionAdObjectsParams) SetBody(body *models.SearchProductionAdObjectsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchProductionAdObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
