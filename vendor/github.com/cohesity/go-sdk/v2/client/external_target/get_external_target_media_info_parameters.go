// Code generated by go-swagger; DO NOT EDIT.

package external_target

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

// NewGetExternalTargetMediaInfoParams creates a new GetExternalTargetMediaInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalTargetMediaInfoParams() *GetExternalTargetMediaInfoParams {
	return &GetExternalTargetMediaInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalTargetMediaInfoParamsWithTimeout creates a new GetExternalTargetMediaInfoParams object
// with the ability to set a timeout on a request.
func NewGetExternalTargetMediaInfoParamsWithTimeout(timeout time.Duration) *GetExternalTargetMediaInfoParams {
	return &GetExternalTargetMediaInfoParams{
		timeout: timeout,
	}
}

// NewGetExternalTargetMediaInfoParamsWithContext creates a new GetExternalTargetMediaInfoParams object
// with the ability to set a context for a request.
func NewGetExternalTargetMediaInfoParamsWithContext(ctx context.Context) *GetExternalTargetMediaInfoParams {
	return &GetExternalTargetMediaInfoParams{
		Context: ctx,
	}
}

// NewGetExternalTargetMediaInfoParamsWithHTTPClient creates a new GetExternalTargetMediaInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalTargetMediaInfoParamsWithHTTPClient(client *http.Client) *GetExternalTargetMediaInfoParams {
	return &GetExternalTargetMediaInfoParams{
		HTTPClient: client,
	}
}

/*
GetExternalTargetMediaInfoParams contains all the parameters to send to the API endpoint

	for the get external target media info operation.

	Typically these are written to a http.Request.
*/
type GetExternalTargetMediaInfoParams struct {

	/* ArchivalJobID.

	   Specifies the id of the Job that archived to a QStar media Vault.

	   Format: int64
	*/
	ArchivalJobID int64

	/* ClusterID.

	   Specifies the id of the Cohesity cluster which archived to a QStart media target.

	   Format: int64
	*/
	ClusterID int64

	/* ClusterIncarnationID.

	   Specifies the incarnation Id of the Cohesity cluster which archived to a QStart media target.

	   Format: int64
	*/
	ClusterIncarnationID int64

	/* EntityIds.

	   Specifies an array of entityIds to optionally filter by. An entityId is a unique id for a VM assigned by the Cohesity Cluster.
	*/
	EntityIds []int64

	/* RestoreTaskID.

	   Specifies the id of the restore task to optionally filter by.

	   Format: int64
	*/
	RestoreTaskID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get external target media info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalTargetMediaInfoParams) WithDefaults() *GetExternalTargetMediaInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get external target media info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalTargetMediaInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithTimeout(timeout time.Duration) *GetExternalTargetMediaInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithContext(ctx context.Context) *GetExternalTargetMediaInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithHTTPClient(client *http.Client) *GetExternalTargetMediaInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchivalJobID adds the archivalJobID to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithArchivalJobID(archivalJobID int64) *GetExternalTargetMediaInfoParams {
	o.SetArchivalJobID(archivalJobID)
	return o
}

// SetArchivalJobID adds the archivalJobId to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetArchivalJobID(archivalJobID int64) {
	o.ArchivalJobID = archivalJobID
}

// WithClusterID adds the clusterID to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithClusterID(clusterID int64) *GetExternalTargetMediaInfoParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetClusterID(clusterID int64) {
	o.ClusterID = clusterID
}

// WithClusterIncarnationID adds the clusterIncarnationID to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithClusterIncarnationID(clusterIncarnationID int64) *GetExternalTargetMediaInfoParams {
	o.SetClusterIncarnationID(clusterIncarnationID)
	return o
}

// SetClusterIncarnationID adds the clusterIncarnationId to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetClusterIncarnationID(clusterIncarnationID int64) {
	o.ClusterIncarnationID = clusterIncarnationID
}

// WithEntityIds adds the entityIds to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithEntityIds(entityIds []int64) *GetExternalTargetMediaInfoParams {
	o.SetEntityIds(entityIds)
	return o
}

// SetEntityIds adds the entityIds to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetEntityIds(entityIds []int64) {
	o.EntityIds = entityIds
}

// WithRestoreTaskID adds the restoreTaskID to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) WithRestoreTaskID(restoreTaskID *int64) *GetExternalTargetMediaInfoParams {
	o.SetRestoreTaskID(restoreTaskID)
	return o
}

// SetRestoreTaskID adds the restoreTaskId to the get external target media info params
func (o *GetExternalTargetMediaInfoParams) SetRestoreTaskID(restoreTaskID *int64) {
	o.RestoreTaskID = restoreTaskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalTargetMediaInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param archivalJobId
	qrArchivalJobID := o.ArchivalJobID
	qArchivalJobID := swag.FormatInt64(qrArchivalJobID)
	if qArchivalJobID != "" {

		if err := r.SetQueryParam("archivalJobId", qArchivalJobID); err != nil {
			return err
		}
	}

	// query param clusterId
	qrClusterID := o.ClusterID
	qClusterID := swag.FormatInt64(qrClusterID)
	if qClusterID != "" {

		if err := r.SetQueryParam("clusterId", qClusterID); err != nil {
			return err
		}
	}

	// query param clusterIncarnationId
	qrClusterIncarnationID := o.ClusterIncarnationID
	qClusterIncarnationID := swag.FormatInt64(qrClusterIncarnationID)
	if qClusterIncarnationID != "" {

		if err := r.SetQueryParam("clusterIncarnationId", qClusterIncarnationID); err != nil {
			return err
		}
	}

	if o.EntityIds != nil {

		// binding items for entityIds
		joinedEntityIds := o.bindParamEntityIds(reg)

		// query array param entityIds
		if err := r.SetQueryParam("entityIds", joinedEntityIds...); err != nil {
			return err
		}
	}

	if o.RestoreTaskID != nil {

		// query param restoreTaskId
		var qrRestoreTaskID int64

		if o.RestoreTaskID != nil {
			qrRestoreTaskID = *o.RestoreTaskID
		}
		qRestoreTaskID := swag.FormatInt64(qrRestoreTaskID)
		if qRestoreTaskID != "" {

			if err := r.SetQueryParam("restoreTaskId", qRestoreTaskID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetExternalTargetMediaInfo binds the parameter entityIds
func (o *GetExternalTargetMediaInfoParams) bindParamEntityIds(formats strfmt.Registry) []string {
	entityIdsIR := o.EntityIds

	var entityIdsIC []string
	for _, entityIdsIIR := range entityIdsIR { // explode []int64

		entityIdsIIV := swag.FormatInt64(entityIdsIIR) // int64 as string
		entityIdsIC = append(entityIdsIC, entityIdsIIV)
	}

	// items.CollectionFormat: ""
	entityIdsIS := swag.JoinByFormat(entityIdsIC, "")

	return entityIdsIS
}
