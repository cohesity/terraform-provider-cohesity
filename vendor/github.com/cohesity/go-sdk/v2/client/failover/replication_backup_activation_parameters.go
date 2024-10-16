// Code generated by go-swagger; DO NOT EDIT.

package failover

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

// NewReplicationBackupActivationParams creates a new ReplicationBackupActivationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplicationBackupActivationParams() *ReplicationBackupActivationParams {
	return &ReplicationBackupActivationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplicationBackupActivationParamsWithTimeout creates a new ReplicationBackupActivationParams object
// with the ability to set a timeout on a request.
func NewReplicationBackupActivationParamsWithTimeout(timeout time.Duration) *ReplicationBackupActivationParams {
	return &ReplicationBackupActivationParams{
		timeout: timeout,
	}
}

// NewReplicationBackupActivationParamsWithContext creates a new ReplicationBackupActivationParams object
// with the ability to set a context for a request.
func NewReplicationBackupActivationParamsWithContext(ctx context.Context) *ReplicationBackupActivationParams {
	return &ReplicationBackupActivationParams{
		Context: ctx,
	}
}

// NewReplicationBackupActivationParamsWithHTTPClient creates a new ReplicationBackupActivationParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplicationBackupActivationParamsWithHTTPClient(client *http.Client) *ReplicationBackupActivationParams {
	return &ReplicationBackupActivationParams{
		HTTPClient: client,
	}
}

/*
ReplicationBackupActivationParams contains all the parameters to send to the API endpoint

	for the replication backup activation operation.

	Typically these are written to a http.Request.
*/
type ReplicationBackupActivationParams struct {

	/* Body.

	   Specifies the paramteres to activate the backup of failover entities.
	*/
	Body *models.ReplicationBackupActivation

	/* ID.

	   Specifies the id of the failover workflow.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replication backup activation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplicationBackupActivationParams) WithDefaults() *ReplicationBackupActivationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replication backup activation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplicationBackupActivationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replication backup activation params
func (o *ReplicationBackupActivationParams) WithTimeout(timeout time.Duration) *ReplicationBackupActivationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replication backup activation params
func (o *ReplicationBackupActivationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replication backup activation params
func (o *ReplicationBackupActivationParams) WithContext(ctx context.Context) *ReplicationBackupActivationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replication backup activation params
func (o *ReplicationBackupActivationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replication backup activation params
func (o *ReplicationBackupActivationParams) WithHTTPClient(client *http.Client) *ReplicationBackupActivationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replication backup activation params
func (o *ReplicationBackupActivationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replication backup activation params
func (o *ReplicationBackupActivationParams) WithBody(body *models.ReplicationBackupActivation) *ReplicationBackupActivationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replication backup activation params
func (o *ReplicationBackupActivationParams) SetBody(body *models.ReplicationBackupActivation) {
	o.Body = body
}

// WithID adds the id to the replication backup activation params
func (o *ReplicationBackupActivationParams) WithID(id string) *ReplicationBackupActivationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replication backup activation params
func (o *ReplicationBackupActivationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReplicationBackupActivationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
