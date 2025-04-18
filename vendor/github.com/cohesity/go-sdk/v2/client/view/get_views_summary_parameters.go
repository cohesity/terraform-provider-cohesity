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
)

// NewGetViewsSummaryParams creates a new GetViewsSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetViewsSummaryParams() *GetViewsSummaryParams {
	return &GetViewsSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetViewsSummaryParamsWithTimeout creates a new GetViewsSummaryParams object
// with the ability to set a timeout on a request.
func NewGetViewsSummaryParamsWithTimeout(timeout time.Duration) *GetViewsSummaryParams {
	return &GetViewsSummaryParams{
		timeout: timeout,
	}
}

// NewGetViewsSummaryParamsWithContext creates a new GetViewsSummaryParams object
// with the ability to set a context for a request.
func NewGetViewsSummaryParamsWithContext(ctx context.Context) *GetViewsSummaryParams {
	return &GetViewsSummaryParams{
		Context: ctx,
	}
}

// NewGetViewsSummaryParamsWithHTTPClient creates a new GetViewsSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetViewsSummaryParamsWithHTTPClient(client *http.Client) *GetViewsSummaryParams {
	return &GetViewsSummaryParams{
		HTTPClient: client,
	}
}

/*
GetViewsSummaryParams contains all the parameters to send to the API endpoint

	for the get views summary operation.

	Typically these are written to a http.Request.
*/
type GetViewsSummaryParams struct {

	/* IncludeDeletedProtectionGroups.

	   Specifies if deleted Protection Groups information needs to be returned along with view metadata. By default, deleted Protection Groups are not returned. This is only applied if used along with any view protection related parameter.
	*/
	IncludeDeletedProtectionGroups *bool

	/* IncludeInternalViews.

	   Specifies if internal Views created by the Cohesity Cluster are also returned. In addition, regular Views are returned.
	*/
	IncludeInternalViews *bool

	/* IncludeTenants.

	   IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned.
	*/
	IncludeTenants *bool

	/* MsecsBeforeCurrentTimeToCompare.

	   Specifies the time in msecs before current time to compare with.

	   Format: int64
	*/
	MsecsBeforeCurrentTimeToCompare *int64

	/* TenantIds.

	   TenantIds contains ids of the tenants for which objects are to be returned.
	*/
	TenantIds []string

	/* UseCachedData.

	   Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.
	*/
	UseCachedData *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get views summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewsSummaryParams) WithDefaults() *GetViewsSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get views summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewsSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get views summary params
func (o *GetViewsSummaryParams) WithTimeout(timeout time.Duration) *GetViewsSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get views summary params
func (o *GetViewsSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get views summary params
func (o *GetViewsSummaryParams) WithContext(ctx context.Context) *GetViewsSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get views summary params
func (o *GetViewsSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get views summary params
func (o *GetViewsSummaryParams) WithHTTPClient(client *http.Client) *GetViewsSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get views summary params
func (o *GetViewsSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeDeletedProtectionGroups adds the includeDeletedProtectionGroups to the get views summary params
func (o *GetViewsSummaryParams) WithIncludeDeletedProtectionGroups(includeDeletedProtectionGroups *bool) *GetViewsSummaryParams {
	o.SetIncludeDeletedProtectionGroups(includeDeletedProtectionGroups)
	return o
}

// SetIncludeDeletedProtectionGroups adds the includeDeletedProtectionGroups to the get views summary params
func (o *GetViewsSummaryParams) SetIncludeDeletedProtectionGroups(includeDeletedProtectionGroups *bool) {
	o.IncludeDeletedProtectionGroups = includeDeletedProtectionGroups
}

// WithIncludeInternalViews adds the includeInternalViews to the get views summary params
func (o *GetViewsSummaryParams) WithIncludeInternalViews(includeInternalViews *bool) *GetViewsSummaryParams {
	o.SetIncludeInternalViews(includeInternalViews)
	return o
}

// SetIncludeInternalViews adds the includeInternalViews to the get views summary params
func (o *GetViewsSummaryParams) SetIncludeInternalViews(includeInternalViews *bool) {
	o.IncludeInternalViews = includeInternalViews
}

// WithIncludeTenants adds the includeTenants to the get views summary params
func (o *GetViewsSummaryParams) WithIncludeTenants(includeTenants *bool) *GetViewsSummaryParams {
	o.SetIncludeTenants(includeTenants)
	return o
}

// SetIncludeTenants adds the includeTenants to the get views summary params
func (o *GetViewsSummaryParams) SetIncludeTenants(includeTenants *bool) {
	o.IncludeTenants = includeTenants
}

// WithMsecsBeforeCurrentTimeToCompare adds the msecsBeforeCurrentTimeToCompare to the get views summary params
func (o *GetViewsSummaryParams) WithMsecsBeforeCurrentTimeToCompare(msecsBeforeCurrentTimeToCompare *int64) *GetViewsSummaryParams {
	o.SetMsecsBeforeCurrentTimeToCompare(msecsBeforeCurrentTimeToCompare)
	return o
}

// SetMsecsBeforeCurrentTimeToCompare adds the msecsBeforeCurrentTimeToCompare to the get views summary params
func (o *GetViewsSummaryParams) SetMsecsBeforeCurrentTimeToCompare(msecsBeforeCurrentTimeToCompare *int64) {
	o.MsecsBeforeCurrentTimeToCompare = msecsBeforeCurrentTimeToCompare
}

// WithTenantIds adds the tenantIds to the get views summary params
func (o *GetViewsSummaryParams) WithTenantIds(tenantIds []string) *GetViewsSummaryParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get views summary params
func (o *GetViewsSummaryParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WithUseCachedData adds the useCachedData to the get views summary params
func (o *GetViewsSummaryParams) WithUseCachedData(useCachedData *bool) *GetViewsSummaryParams {
	o.SetUseCachedData(useCachedData)
	return o
}

// SetUseCachedData adds the useCachedData to the get views summary params
func (o *GetViewsSummaryParams) SetUseCachedData(useCachedData *bool) {
	o.UseCachedData = useCachedData
}

// WriteToRequest writes these params to a swagger request
func (o *GetViewsSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeDeletedProtectionGroups != nil {

		// query param includeDeletedProtectionGroups
		var qrIncludeDeletedProtectionGroups bool

		if o.IncludeDeletedProtectionGroups != nil {
			qrIncludeDeletedProtectionGroups = *o.IncludeDeletedProtectionGroups
		}
		qIncludeDeletedProtectionGroups := swag.FormatBool(qrIncludeDeletedProtectionGroups)
		if qIncludeDeletedProtectionGroups != "" {

			if err := r.SetQueryParam("includeDeletedProtectionGroups", qIncludeDeletedProtectionGroups); err != nil {
				return err
			}
		}
	}

	if o.IncludeInternalViews != nil {

		// query param includeInternalViews
		var qrIncludeInternalViews bool

		if o.IncludeInternalViews != nil {
			qrIncludeInternalViews = *o.IncludeInternalViews
		}
		qIncludeInternalViews := swag.FormatBool(qrIncludeInternalViews)
		if qIncludeInternalViews != "" {

			if err := r.SetQueryParam("includeInternalViews", qIncludeInternalViews); err != nil {
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

	if o.MsecsBeforeCurrentTimeToCompare != nil {

		// query param msecsBeforeCurrentTimeToCompare
		var qrMsecsBeforeCurrentTimeToCompare int64

		if o.MsecsBeforeCurrentTimeToCompare != nil {
			qrMsecsBeforeCurrentTimeToCompare = *o.MsecsBeforeCurrentTimeToCompare
		}
		qMsecsBeforeCurrentTimeToCompare := swag.FormatInt64(qrMsecsBeforeCurrentTimeToCompare)
		if qMsecsBeforeCurrentTimeToCompare != "" {

			if err := r.SetQueryParam("msecsBeforeCurrentTimeToCompare", qMsecsBeforeCurrentTimeToCompare); err != nil {
				return err
			}
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

	if o.UseCachedData != nil {

		// query param useCachedData
		var qrUseCachedData bool

		if o.UseCachedData != nil {
			qrUseCachedData = *o.UseCachedData
		}
		qUseCachedData := swag.FormatBool(qrUseCachedData)
		if qUseCachedData != "" {

			if err := r.SetQueryParam("useCachedData", qUseCachedData); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetViewsSummary binds the parameter tenantIds
func (o *GetViewsSummaryParams) bindParamTenantIds(formats strfmt.Registry) []string {
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
