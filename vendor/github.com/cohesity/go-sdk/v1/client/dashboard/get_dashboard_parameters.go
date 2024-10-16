// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewGetDashboardParams creates a new GetDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardParams() *GetDashboardParams {
	return &GetDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardParamsWithTimeout creates a new GetDashboardParams object
// with the ability to set a timeout on a request.
func NewGetDashboardParamsWithTimeout(timeout time.Duration) *GetDashboardParams {
	return &GetDashboardParams{
		timeout: timeout,
	}
}

// NewGetDashboardParamsWithContext creates a new GetDashboardParams object
// with the ability to set a context for a request.
func NewGetDashboardParamsWithContext(ctx context.Context) *GetDashboardParams {
	return &GetDashboardParams{
		Context: ctx,
	}
}

// NewGetDashboardParamsWithHTTPClient creates a new GetDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardParamsWithHTTPClient(client *http.Client) *GetDashboardParams {
	return &GetDashboardParams{
		HTTPClient: client,
	}
}

/*
GetDashboardParams contains all the parameters to send to the API endpoint

	for the get dashboard operation.

	Typically these are written to a http.Request.
*/
type GetDashboardParams struct {

	/* AllClusters.

	     Summary data for all clusters. If this is set to true, all other
	parameters will be ignored.
	*/
	AllClusters *bool

	/* ClusterID.

	     Id of the remote cluster for which to fetch the data. If value is not
	specified, it is assumed to be local cluster.

	     Format: int64
	*/
	ClusterID *int64

	/* Refresh.

	   Specifies whether to refresh the tiles selected.
	*/
	Refresh *bool

	/* TileTypes.

	     Specifies the types of the tiles to be returned. If this is not
	specified, all the tiles are returned. This is ignored when allClusters
	is set to true.
	'kHealthTile' is the tile that shows health of the cluster and the alerts in
	the past 24 hours.
	'kJobRunsTile' is the tile that shows job runs in the past 24 hours.
	'kRecoveriesTile' is the tile that shows recoveries done in the past 30 days.
	'kProtectedObjectsTile' is the tile that shows the protected objects details.
	'kProtectionTile' is the tile that shows the protection information in the
	past 24 hours.
	'kAuditLogsTile' is the tile that shows the recent audit logs.
	'kIopsTile' is the tile that shows IP performance in the past 24 hours.
	'kThroughputTile' is the tile that shows job runs in the past 24 hours.
	'kStorageEfficiencyTile' is the tile that shows job runs in the past 7 days.
	*/
	TileTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardParams) WithDefaults() *GetDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dashboard params
func (o *GetDashboardParams) WithTimeout(timeout time.Duration) *GetDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard params
func (o *GetDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard params
func (o *GetDashboardParams) WithContext(ctx context.Context) *GetDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard params
func (o *GetDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard params
func (o *GetDashboardParams) WithHTTPClient(client *http.Client) *GetDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard params
func (o *GetDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllClusters adds the allClusters to the get dashboard params
func (o *GetDashboardParams) WithAllClusters(allClusters *bool) *GetDashboardParams {
	o.SetAllClusters(allClusters)
	return o
}

// SetAllClusters adds the allClusters to the get dashboard params
func (o *GetDashboardParams) SetAllClusters(allClusters *bool) {
	o.AllClusters = allClusters
}

// WithClusterID adds the clusterID to the get dashboard params
func (o *GetDashboardParams) WithClusterID(clusterID *int64) *GetDashboardParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get dashboard params
func (o *GetDashboardParams) SetClusterID(clusterID *int64) {
	o.ClusterID = clusterID
}

// WithRefresh adds the refresh to the get dashboard params
func (o *GetDashboardParams) WithRefresh(refresh *bool) *GetDashboardParams {
	o.SetRefresh(refresh)
	return o
}

// SetRefresh adds the refresh to the get dashboard params
func (o *GetDashboardParams) SetRefresh(refresh *bool) {
	o.Refresh = refresh
}

// WithTileTypes adds the tileTypes to the get dashboard params
func (o *GetDashboardParams) WithTileTypes(tileTypes []string) *GetDashboardParams {
	o.SetTileTypes(tileTypes)
	return o
}

// SetTileTypes adds the tileTypes to the get dashboard params
func (o *GetDashboardParams) SetTileTypes(tileTypes []string) {
	o.TileTypes = tileTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllClusters != nil {

		// query param allClusters
		var qrAllClusters bool

		if o.AllClusters != nil {
			qrAllClusters = *o.AllClusters
		}
		qAllClusters := swag.FormatBool(qrAllClusters)
		if qAllClusters != "" {

			if err := r.SetQueryParam("allClusters", qAllClusters); err != nil {
				return err
			}
		}
	}

	if o.ClusterID != nil {

		// query param clusterId
		var qrClusterID int64

		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := swag.FormatInt64(qrClusterID)
		if qClusterID != "" {

			if err := r.SetQueryParam("clusterId", qClusterID); err != nil {
				return err
			}
		}
	}

	if o.Refresh != nil {

		// query param refresh
		var qrRefresh bool

		if o.Refresh != nil {
			qrRefresh = *o.Refresh
		}
		qRefresh := swag.FormatBool(qrRefresh)
		if qRefresh != "" {

			if err := r.SetQueryParam("refresh", qRefresh); err != nil {
				return err
			}
		}
	}

	if o.TileTypes != nil {

		// binding items for tileTypes
		joinedTileTypes := o.bindParamTileTypes(reg)

		// query array param tileTypes
		if err := r.SetQueryParam("tileTypes", joinedTileTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDashboard binds the parameter tileTypes
func (o *GetDashboardParams) bindParamTileTypes(formats strfmt.Registry) []string {
	tileTypesIR := o.TileTypes

	var tileTypesIC []string
	for _, tileTypesIIR := range tileTypesIR { // explode []string

		tileTypesIIV := tileTypesIIR // string as string
		tileTypesIC = append(tileTypesIC, tileTypesIIV)
	}

	// items.CollectionFormat: ""
	tileTypesIS := swag.JoinByFormat(tileTypesIC, "")

	return tileTypesIS
}
