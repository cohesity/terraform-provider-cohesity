// Code generated by go-swagger; DO NOT EDIT.

package storage_domain

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

// NewGetStorageDomainsParams creates a new GetStorageDomainsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStorageDomainsParams() *GetStorageDomainsParams {
	return &GetStorageDomainsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageDomainsParamsWithTimeout creates a new GetStorageDomainsParams object
// with the ability to set a timeout on a request.
func NewGetStorageDomainsParamsWithTimeout(timeout time.Duration) *GetStorageDomainsParams {
	return &GetStorageDomainsParams{
		timeout: timeout,
	}
}

// NewGetStorageDomainsParamsWithContext creates a new GetStorageDomainsParams object
// with the ability to set a context for a request.
func NewGetStorageDomainsParamsWithContext(ctx context.Context) *GetStorageDomainsParams {
	return &GetStorageDomainsParams{
		Context: ctx,
	}
}

// NewGetStorageDomainsParamsWithHTTPClient creates a new GetStorageDomainsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStorageDomainsParamsWithHTTPClient(client *http.Client) *GetStorageDomainsParams {
	return &GetStorageDomainsParams{
		HTTPClient: client,
	}
}

/*
GetStorageDomainsParams contains all the parameters to send to the API endpoint

	for the get storage domains operation.

	Typically these are written to a http.Request.
*/
type GetStorageDomainsParams struct {

	/* ClusterPartitionIds.

	   Filter by a list of cluster partition ids.
	*/
	ClusterPartitionIds []int64

	/* Ids.

	   Filter by a list of Storage Domain ids.
	*/
	Ids []int64

	/* IncludeFileCountBySize.

	   Whether to include Storage Domain file count by size.
	*/
	IncludeFileCountBySize *bool

	/* IncludeStats.

	   Whether to include Storage Domain stats in response.
	*/
	IncludeStats *bool

	/* IncludeTenants.

	   IncludeTenants specifies if Storage Domains of all the tenants under the hierarchy of the logged in user's organization should be returned.
	*/
	IncludeTenants *bool

	/* IncludeTimeSeriesSchema.

	   Whether to include Storage Domain time series schema in response.
	*/
	IncludeTimeSeriesSchema *bool

	/* MatchPartialNames.

	   If set to true, names in the 'names' parameter will be matched partially instead of exactly.
	*/
	MatchPartialNames *bool

	/* Names.

	   Filter by a list of Storage Domain names.
	*/
	Names []string

	/* TenantIds.

	   TenantIds contains ids of the tenants for which Storage Domains are to be returned.
	*/
	TenantIds []string

	/* ViewTemplateID.

	   Specifies a view template id for Storage Domain. Storage Domains with same deduplication and compression settings will be recommended.

	   Format: int64
	*/
	ViewTemplateID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get storage domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageDomainsParams) WithDefaults() *GetStorageDomainsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get storage domains params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStorageDomainsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get storage domains params
func (o *GetStorageDomainsParams) WithTimeout(timeout time.Duration) *GetStorageDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage domains params
func (o *GetStorageDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage domains params
func (o *GetStorageDomainsParams) WithContext(ctx context.Context) *GetStorageDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage domains params
func (o *GetStorageDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage domains params
func (o *GetStorageDomainsParams) WithHTTPClient(client *http.Client) *GetStorageDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage domains params
func (o *GetStorageDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterPartitionIds adds the clusterPartitionIds to the get storage domains params
func (o *GetStorageDomainsParams) WithClusterPartitionIds(clusterPartitionIds []int64) *GetStorageDomainsParams {
	o.SetClusterPartitionIds(clusterPartitionIds)
	return o
}

// SetClusterPartitionIds adds the clusterPartitionIds to the get storage domains params
func (o *GetStorageDomainsParams) SetClusterPartitionIds(clusterPartitionIds []int64) {
	o.ClusterPartitionIds = clusterPartitionIds
}

// WithIds adds the ids to the get storage domains params
func (o *GetStorageDomainsParams) WithIds(ids []int64) *GetStorageDomainsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get storage domains params
func (o *GetStorageDomainsParams) SetIds(ids []int64) {
	o.Ids = ids
}

// WithIncludeFileCountBySize adds the includeFileCountBySize to the get storage domains params
func (o *GetStorageDomainsParams) WithIncludeFileCountBySize(includeFileCountBySize *bool) *GetStorageDomainsParams {
	o.SetIncludeFileCountBySize(includeFileCountBySize)
	return o
}

// SetIncludeFileCountBySize adds the includeFileCountBySize to the get storage domains params
func (o *GetStorageDomainsParams) SetIncludeFileCountBySize(includeFileCountBySize *bool) {
	o.IncludeFileCountBySize = includeFileCountBySize
}

// WithIncludeStats adds the includeStats to the get storage domains params
func (o *GetStorageDomainsParams) WithIncludeStats(includeStats *bool) *GetStorageDomainsParams {
	o.SetIncludeStats(includeStats)
	return o
}

// SetIncludeStats adds the includeStats to the get storage domains params
func (o *GetStorageDomainsParams) SetIncludeStats(includeStats *bool) {
	o.IncludeStats = includeStats
}

// WithIncludeTenants adds the includeTenants to the get storage domains params
func (o *GetStorageDomainsParams) WithIncludeTenants(includeTenants *bool) *GetStorageDomainsParams {
	o.SetIncludeTenants(includeTenants)
	return o
}

// SetIncludeTenants adds the includeTenants to the get storage domains params
func (o *GetStorageDomainsParams) SetIncludeTenants(includeTenants *bool) {
	o.IncludeTenants = includeTenants
}

// WithIncludeTimeSeriesSchema adds the includeTimeSeriesSchema to the get storage domains params
func (o *GetStorageDomainsParams) WithIncludeTimeSeriesSchema(includeTimeSeriesSchema *bool) *GetStorageDomainsParams {
	o.SetIncludeTimeSeriesSchema(includeTimeSeriesSchema)
	return o
}

// SetIncludeTimeSeriesSchema adds the includeTimeSeriesSchema to the get storage domains params
func (o *GetStorageDomainsParams) SetIncludeTimeSeriesSchema(includeTimeSeriesSchema *bool) {
	o.IncludeTimeSeriesSchema = includeTimeSeriesSchema
}

// WithMatchPartialNames adds the matchPartialNames to the get storage domains params
func (o *GetStorageDomainsParams) WithMatchPartialNames(matchPartialNames *bool) *GetStorageDomainsParams {
	o.SetMatchPartialNames(matchPartialNames)
	return o
}

// SetMatchPartialNames adds the matchPartialNames to the get storage domains params
func (o *GetStorageDomainsParams) SetMatchPartialNames(matchPartialNames *bool) {
	o.MatchPartialNames = matchPartialNames
}

// WithNames adds the names to the get storage domains params
func (o *GetStorageDomainsParams) WithNames(names []string) *GetStorageDomainsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get storage domains params
func (o *GetStorageDomainsParams) SetNames(names []string) {
	o.Names = names
}

// WithTenantIds adds the tenantIds to the get storage domains params
func (o *GetStorageDomainsParams) WithTenantIds(tenantIds []string) *GetStorageDomainsParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get storage domains params
func (o *GetStorageDomainsParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WithViewTemplateID adds the viewTemplateID to the get storage domains params
func (o *GetStorageDomainsParams) WithViewTemplateID(viewTemplateID *int64) *GetStorageDomainsParams {
	o.SetViewTemplateID(viewTemplateID)
	return o
}

// SetViewTemplateID adds the viewTemplateId to the get storage domains params
func (o *GetStorageDomainsParams) SetViewTemplateID(viewTemplateID *int64) {
	o.ViewTemplateID = viewTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterPartitionIds != nil {

		// binding items for clusterPartitionIds
		joinedClusterPartitionIds := o.bindParamClusterPartitionIds(reg)

		// query array param clusterPartitionIds
		if err := r.SetQueryParam("clusterPartitionIds", joinedClusterPartitionIds...); err != nil {
			return err
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if o.IncludeFileCountBySize != nil {

		// query param includeFileCountBySize
		var qrIncludeFileCountBySize bool

		if o.IncludeFileCountBySize != nil {
			qrIncludeFileCountBySize = *o.IncludeFileCountBySize
		}
		qIncludeFileCountBySize := swag.FormatBool(qrIncludeFileCountBySize)
		if qIncludeFileCountBySize != "" {

			if err := r.SetQueryParam("includeFileCountBySize", qIncludeFileCountBySize); err != nil {
				return err
			}
		}
	}

	if o.IncludeStats != nil {

		// query param includeStats
		var qrIncludeStats bool

		if o.IncludeStats != nil {
			qrIncludeStats = *o.IncludeStats
		}
		qIncludeStats := swag.FormatBool(qrIncludeStats)
		if qIncludeStats != "" {

			if err := r.SetQueryParam("includeStats", qIncludeStats); err != nil {
				return err
			}
		}
	}

	if o.IncludeTenants != nil {

		// query param includeTenants
		var qrIncludeTenants bool

		if o.IncludeTenants != nil {
			qrIncludeTenants = *o.IncludeTenants
		}
		qIncludeTenants := swag.FormatBool(qrIncludeTenants)
		if qIncludeTenants != "" {

			if err := r.SetQueryParam("includeTenants", qIncludeTenants); err != nil {
				return err
			}
		}
	}

	if o.IncludeTimeSeriesSchema != nil {

		// query param includeTimeSeriesSchema
		var qrIncludeTimeSeriesSchema bool

		if o.IncludeTimeSeriesSchema != nil {
			qrIncludeTimeSeriesSchema = *o.IncludeTimeSeriesSchema
		}
		qIncludeTimeSeriesSchema := swag.FormatBool(qrIncludeTimeSeriesSchema)
		if qIncludeTimeSeriesSchema != "" {

			if err := r.SetQueryParam("includeTimeSeriesSchema", qIncludeTimeSeriesSchema); err != nil {
				return err
			}
		}
	}

	if o.MatchPartialNames != nil {

		// query param matchPartialNames
		var qrMatchPartialNames bool

		if o.MatchPartialNames != nil {
			qrMatchPartialNames = *o.MatchPartialNames
		}
		qMatchPartialNames := swag.FormatBool(qrMatchPartialNames)
		if qMatchPartialNames != "" {

			if err := r.SetQueryParam("matchPartialNames", qMatchPartialNames); err != nil {
				return err
			}
		}
	}

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if o.TenantIds != nil {

		// binding items for tenantIds
		joinedTenantIds := o.bindParamTenantIds(reg)

		// query array param tenantIds
		if err := r.SetQueryParam("tenantIds", joinedTenantIds...); err != nil {
			return err
		}
	}

	if o.ViewTemplateID != nil {

		// query param viewTemplateId
		var qrViewTemplateID int64

		if o.ViewTemplateID != nil {
			qrViewTemplateID = *o.ViewTemplateID
		}
		qViewTemplateID := swag.FormatInt64(qrViewTemplateID)
		if qViewTemplateID != "" {

			if err := r.SetQueryParam("viewTemplateId", qViewTemplateID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetStorageDomains binds the parameter clusterPartitionIds
func (o *GetStorageDomainsParams) bindParamClusterPartitionIds(formats strfmt.Registry) []string {
	clusterPartitionIdsIR := o.ClusterPartitionIds

	var clusterPartitionIdsIC []string
	for _, clusterPartitionIdsIIR := range clusterPartitionIdsIR { // explode []int64

		clusterPartitionIdsIIV := swag.FormatInt64(clusterPartitionIdsIIR) // int64 as string
		clusterPartitionIdsIC = append(clusterPartitionIdsIC, clusterPartitionIdsIIV)
	}

	// items.CollectionFormat: ""
	clusterPartitionIdsIS := swag.JoinByFormat(clusterPartitionIdsIC, "")

	return clusterPartitionIdsIS
}

// bindParamGetStorageDomains binds the parameter ids
func (o *GetStorageDomainsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []int64

		idsIIV := swag.FormatInt64(idsIIR) // int64 as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: ""
	idsIS := swag.JoinByFormat(idsIC, "")

	return idsIS
}

// bindParamGetStorageDomains binds the parameter names
func (o *GetStorageDomainsParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: ""
	namesIS := swag.JoinByFormat(namesIC, "")

	return namesIS
}

// bindParamGetStorageDomains binds the parameter tenantIds
func (o *GetStorageDomainsParams) bindParamTenantIds(formats strfmt.Registry) []string {
	tenantIdsIR := o.TenantIds

	var tenantIdsIC []string
	for _, tenantIdsIIR := range tenantIdsIR { // explode []string

		tenantIdsIIV := tenantIdsIIR // string as string
		tenantIdsIC = append(tenantIdsIC, tenantIdsIIV)
	}

	// items.CollectionFormat: ""
	tenantIdsIS := swag.JoinByFormat(tenantIdsIC, "")

	return tenantIdsIS
}
