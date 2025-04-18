// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewDeleteTagParams creates a new DeleteTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTagParams() *DeleteTagParams {
	return &DeleteTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTagParamsWithTimeout creates a new DeleteTagParams object
// with the ability to set a timeout on a request.
func NewDeleteTagParamsWithTimeout(timeout time.Duration) *DeleteTagParams {
	return &DeleteTagParams{
		timeout: timeout,
	}
}

// NewDeleteTagParamsWithContext creates a new DeleteTagParams object
// with the ability to set a context for a request.
func NewDeleteTagParamsWithContext(ctx context.Context) *DeleteTagParams {
	return &DeleteTagParams{
		Context: ctx,
	}
}

// NewDeleteTagParamsWithHTTPClient creates a new DeleteTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTagParamsWithHTTPClient(client *http.Client) *DeleteTagParams {
	return &DeleteTagParams{
		HTTPClient: client,
	}
}

/*
DeleteTagParams contains all the parameters to send to the API endpoint

	for the delete tag operation.

	Typically these are written to a http.Request.
*/
type DeleteTagParams struct {

	/* ID.

	   Specifies the Id of the tag.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTagParams) WithDefaults() *DeleteTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete tag params
func (o *DeleteTagParams) WithTimeout(timeout time.Duration) *DeleteTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete tag params
func (o *DeleteTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete tag params
func (o *DeleteTagParams) WithContext(ctx context.Context) *DeleteTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete tag params
func (o *DeleteTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete tag params
func (o *DeleteTagParams) WithHTTPClient(client *http.Client) *DeleteTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete tag params
func (o *DeleteTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete tag params
func (o *DeleteTagParams) WithID(id string) *DeleteTagParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete tag params
func (o *DeleteTagParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
