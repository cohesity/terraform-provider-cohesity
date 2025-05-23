// Code generated by go-swagger; DO NOT EDIT.

package platform

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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewUpgradeNodesParams creates a new UpgradeNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeNodesParams() *UpgradeNodesParams {
	return &UpgradeNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeNodesParamsWithTimeout creates a new UpgradeNodesParams object
// with the ability to set a timeout on a request.
func NewUpgradeNodesParamsWithTimeout(timeout time.Duration) *UpgradeNodesParams {
	return &UpgradeNodesParams{
		timeout: timeout,
	}
}

// NewUpgradeNodesParamsWithContext creates a new UpgradeNodesParams object
// with the ability to set a context for a request.
func NewUpgradeNodesParamsWithContext(ctx context.Context) *UpgradeNodesParams {
	return &UpgradeNodesParams{
		Context: ctx,
	}
}

// NewUpgradeNodesParamsWithHTTPClient creates a new UpgradeNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeNodesParamsWithHTTPClient(client *http.Client) *UpgradeNodesParams {
	return &UpgradeNodesParams{
		HTTPClient: client,
	}
}

/*
UpgradeNodesParams contains all the parameters to send to the API endpoint

	for the upgrade nodes operation.

	Typically these are written to a http.Request.
*/
type UpgradeNodesParams struct {

	/* Body.

	   The parameters to upgrade free node(s).
	*/
	Body *models.NodeUpgradeParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeNodesParams) WithDefaults() *UpgradeNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade nodes params
func (o *UpgradeNodesParams) WithTimeout(timeout time.Duration) *UpgradeNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade nodes params
func (o *UpgradeNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade nodes params
func (o *UpgradeNodesParams) WithContext(ctx context.Context) *UpgradeNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade nodes params
func (o *UpgradeNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade nodes params
func (o *UpgradeNodesParams) WithHTTPClient(client *http.Client) *UpgradeNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade nodes params
func (o *UpgradeNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade nodes params
func (o *UpgradeNodesParams) WithBody(body *models.NodeUpgradeParameters) *UpgradeNodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade nodes params
func (o *UpgradeNodesParams) SetBody(body *models.NodeUpgradeParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
