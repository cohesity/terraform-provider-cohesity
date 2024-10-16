// Code generated by go-swagger; DO NOT EDIT.

package vaults

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

	"github.com/cohesity/go-sdk/v1/models"
)

// NewDeleteVaultParams creates a new DeleteVaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVaultParams() *DeleteVaultParams {
	return &DeleteVaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVaultParamsWithTimeout creates a new DeleteVaultParams object
// with the ability to set a timeout on a request.
func NewDeleteVaultParamsWithTimeout(timeout time.Duration) *DeleteVaultParams {
	return &DeleteVaultParams{
		timeout: timeout,
	}
}

// NewDeleteVaultParamsWithContext creates a new DeleteVaultParams object
// with the ability to set a context for a request.
func NewDeleteVaultParamsWithContext(ctx context.Context) *DeleteVaultParams {
	return &DeleteVaultParams{
		Context: ctx,
	}
}

// NewDeleteVaultParamsWithHTTPClient creates a new DeleteVaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVaultParamsWithHTTPClient(client *http.Client) *DeleteVaultParams {
	return &DeleteVaultParams{
		HTTPClient: client,
	}
}

/*
DeleteVaultParams contains all the parameters to send to the API endpoint

	for the delete vault operation.

	Typically these are written to a http.Request.
*/
type DeleteVaultParams struct {

	/* Body.

	   Request to delete vault.
	*/
	Body *models.VaultDeleteParams

	/* ID.

	   Specifies a unique id of the Vault.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete vault params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVaultParams) WithDefaults() *DeleteVaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete vault params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete vault params
func (o *DeleteVaultParams) WithTimeout(timeout time.Duration) *DeleteVaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vault params
func (o *DeleteVaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vault params
func (o *DeleteVaultParams) WithContext(ctx context.Context) *DeleteVaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vault params
func (o *DeleteVaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vault params
func (o *DeleteVaultParams) WithHTTPClient(client *http.Client) *DeleteVaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vault params
func (o *DeleteVaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete vault params
func (o *DeleteVaultParams) WithBody(body *models.VaultDeleteParams) *DeleteVaultParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete vault params
func (o *DeleteVaultParams) SetBody(body *models.VaultDeleteParams) {
	o.Body = body
}

// WithID adds the id to the delete vault params
func (o *DeleteVaultParams) WithID(id int64) *DeleteVaultParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vault params
func (o *DeleteVaultParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
