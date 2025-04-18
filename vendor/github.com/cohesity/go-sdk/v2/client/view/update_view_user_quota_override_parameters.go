// Code generated by go-swagger; DO NOT EDIT.

package view

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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewUpdateViewUserQuotaOverrideParams creates a new UpdateViewUserQuotaOverrideParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateViewUserQuotaOverrideParams() *UpdateViewUserQuotaOverrideParams {
	return &UpdateViewUserQuotaOverrideParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateViewUserQuotaOverrideParamsWithTimeout creates a new UpdateViewUserQuotaOverrideParams object
// with the ability to set a timeout on a request.
func NewUpdateViewUserQuotaOverrideParamsWithTimeout(timeout time.Duration) *UpdateViewUserQuotaOverrideParams {
	return &UpdateViewUserQuotaOverrideParams{
		timeout: timeout,
	}
}

// NewUpdateViewUserQuotaOverrideParamsWithContext creates a new UpdateViewUserQuotaOverrideParams object
// with the ability to set a context for a request.
func NewUpdateViewUserQuotaOverrideParamsWithContext(ctx context.Context) *UpdateViewUserQuotaOverrideParams {
	return &UpdateViewUserQuotaOverrideParams{
		Context: ctx,
	}
}

// NewUpdateViewUserQuotaOverrideParamsWithHTTPClient creates a new UpdateViewUserQuotaOverrideParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateViewUserQuotaOverrideParamsWithHTTPClient(client *http.Client) *UpdateViewUserQuotaOverrideParams {
	return &UpdateViewUserQuotaOverrideParams{
		HTTPClient: client,
	}
}

/*
UpdateViewUserQuotaOverrideParams contains all the parameters to send to the API endpoint

	for the update view user quota override operation.

	Typically these are written to a http.Request.
*/
type UpdateViewUserQuotaOverrideParams struct {

	/* Body.

	   Specifies the user quota policy of the user.
	*/
	Body *models.QuotaPolicy

	/* UserID.

	   Specifies the unixUid or sid or an user.
	*/
	UserID string

	/* ViewID.

	   Specifies the View id.

	   Format: int64
	*/
	ViewID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update view user quota override params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewUserQuotaOverrideParams) WithDefaults() *UpdateViewUserQuotaOverrideParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update view user quota override params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewUserQuotaOverrideParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) WithTimeout(timeout time.Duration) *UpdateViewUserQuotaOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) WithContext(ctx context.Context) *UpdateViewUserQuotaOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) WithHTTPClient(client *http.Client) *UpdateViewUserQuotaOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) WithBody(body *models.QuotaPolicy) *UpdateViewUserQuotaOverrideParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) SetBody(body *models.QuotaPolicy) {
	o.Body = body
}

// WithUserID adds the userID to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) WithUserID(userID string) *UpdateViewUserQuotaOverrideParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithViewID adds the viewID to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) WithViewID(viewID int64) *UpdateViewUserQuotaOverrideParams {
	o.SetViewID(viewID)
	return o
}

// SetViewID adds the viewId to the update view user quota override params
func (o *UpdateViewUserQuotaOverrideParams) SetViewID(viewID int64) {
	o.ViewID = viewID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateViewUserQuotaOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	// path param viewId
	if err := r.SetPathParam("viewId", swag.FormatInt64(o.ViewID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
