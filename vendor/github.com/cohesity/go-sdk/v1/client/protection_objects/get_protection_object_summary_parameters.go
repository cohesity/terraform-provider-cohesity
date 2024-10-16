// Code generated by go-swagger; DO NOT EDIT.

package protection_objects

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

// NewGetProtectionObjectSummaryParams creates a new GetProtectionObjectSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProtectionObjectSummaryParams() *GetProtectionObjectSummaryParams {
	return &GetProtectionObjectSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtectionObjectSummaryParamsWithTimeout creates a new GetProtectionObjectSummaryParams object
// with the ability to set a timeout on a request.
func NewGetProtectionObjectSummaryParamsWithTimeout(timeout time.Duration) *GetProtectionObjectSummaryParams {
	return &GetProtectionObjectSummaryParams{
		timeout: timeout,
	}
}

// NewGetProtectionObjectSummaryParamsWithContext creates a new GetProtectionObjectSummaryParams object
// with the ability to set a context for a request.
func NewGetProtectionObjectSummaryParamsWithContext(ctx context.Context) *GetProtectionObjectSummaryParams {
	return &GetProtectionObjectSummaryParams{
		Context: ctx,
	}
}

// NewGetProtectionObjectSummaryParamsWithHTTPClient creates a new GetProtectionObjectSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProtectionObjectSummaryParamsWithHTTPClient(client *http.Client) *GetProtectionObjectSummaryParams {
	return &GetProtectionObjectSummaryParams{
		HTTPClient: client,
	}
}

/*
GetProtectionObjectSummaryParams contains all the parameters to send to the API endpoint

	for the get protection object summary operation.

	Typically these are written to a http.Request.
*/
type GetProtectionObjectSummaryParams struct {

	/* ProtectionSourceID.

	   Specifies the id of the Protection Source.

	   Format: int64
	*/
	ProtectionSourceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get protection object summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionObjectSummaryParams) WithDefaults() *GetProtectionObjectSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get protection object summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionObjectSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) WithTimeout(timeout time.Duration) *GetProtectionObjectSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) WithContext(ctx context.Context) *GetProtectionObjectSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) WithHTTPClient(client *http.Client) *GetProtectionObjectSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProtectionSourceID adds the protectionSourceID to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) WithProtectionSourceID(protectionSourceID int64) *GetProtectionObjectSummaryParams {
	o.SetProtectionSourceID(protectionSourceID)
	return o
}

// SetProtectionSourceID adds the protectionSourceId to the get protection object summary params
func (o *GetProtectionObjectSummaryParams) SetProtectionSourceID(protectionSourceID int64) {
	o.ProtectionSourceID = protectionSourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtectionObjectSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param protectionSourceId
	qrProtectionSourceID := o.ProtectionSourceID
	qProtectionSourceID := swag.FormatInt64(qrProtectionSourceID)
	if qProtectionSourceID != "" {

		if err := r.SetQueryParam("protectionSourceId", qProtectionSourceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
