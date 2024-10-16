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

// NewUpdateIsDMaaSClusterParams creates a new UpdateIsDMaaSClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIsDMaaSClusterParams() *UpdateIsDMaaSClusterParams {
	return &UpdateIsDMaaSClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIsDMaaSClusterParamsWithTimeout creates a new UpdateIsDMaaSClusterParams object
// with the ability to set a timeout on a request.
func NewUpdateIsDMaaSClusterParamsWithTimeout(timeout time.Duration) *UpdateIsDMaaSClusterParams {
	return &UpdateIsDMaaSClusterParams{
		timeout: timeout,
	}
}

// NewUpdateIsDMaaSClusterParamsWithContext creates a new UpdateIsDMaaSClusterParams object
// with the ability to set a context for a request.
func NewUpdateIsDMaaSClusterParamsWithContext(ctx context.Context) *UpdateIsDMaaSClusterParams {
	return &UpdateIsDMaaSClusterParams{
		Context: ctx,
	}
}

// NewUpdateIsDMaaSClusterParamsWithHTTPClient creates a new UpdateIsDMaaSClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIsDMaaSClusterParamsWithHTTPClient(client *http.Client) *UpdateIsDMaaSClusterParams {
	return &UpdateIsDMaaSClusterParams{
		HTTPClient: client,
	}
}

/*
UpdateIsDMaaSClusterParams contains all the parameters to send to the API endpoint

	for the update is d maa s cluster operation.

	Typically these are written to a http.Request.
*/
type UpdateIsDMaaSClusterParams struct {

	/* Body.

	   Param to update whether the cluster is a DMaaS cluster.
	*/
	Body *models.DMaaSInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update is d maa s cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIsDMaaSClusterParams) WithDefaults() *UpdateIsDMaaSClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update is d maa s cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIsDMaaSClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) WithTimeout(timeout time.Duration) *UpdateIsDMaaSClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) WithContext(ctx context.Context) *UpdateIsDMaaSClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) WithHTTPClient(client *http.Client) *UpdateIsDMaaSClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) WithBody(body *models.DMaaSInfo) *UpdateIsDMaaSClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update is d maa s cluster params
func (o *UpdateIsDMaaSClusterParams) SetBody(body *models.DMaaSInfo) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIsDMaaSClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
