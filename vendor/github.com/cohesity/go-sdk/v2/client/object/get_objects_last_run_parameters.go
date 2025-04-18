// Code generated by go-swagger; DO NOT EDIT.

package object

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

// NewGetObjectsLastRunParams creates a new GetObjectsLastRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetObjectsLastRunParams() *GetObjectsLastRunParams {
	return &GetObjectsLastRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetObjectsLastRunParamsWithTimeout creates a new GetObjectsLastRunParams object
// with the ability to set a timeout on a request.
func NewGetObjectsLastRunParamsWithTimeout(timeout time.Duration) *GetObjectsLastRunParams {
	return &GetObjectsLastRunParams{
		timeout: timeout,
	}
}

// NewGetObjectsLastRunParamsWithContext creates a new GetObjectsLastRunParams object
// with the ability to set a context for a request.
func NewGetObjectsLastRunParamsWithContext(ctx context.Context) *GetObjectsLastRunParams {
	return &GetObjectsLastRunParams{
		Context: ctx,
	}
}

// NewGetObjectsLastRunParamsWithHTTPClient creates a new GetObjectsLastRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetObjectsLastRunParamsWithHTTPClient(client *http.Client) *GetObjectsLastRunParams {
	return &GetObjectsLastRunParams{
		HTTPClient: client,
	}
}

/*
GetObjectsLastRunParams contains all the parameters to send to the API endpoint

	for the get objects last run operation.

	Typically these are written to a http.Request.
*/
type GetObjectsLastRunParams struct {

	/* Count.

	   Specifies the number of objects to be fetched for the specified pagination cookie.

	   Format: int32
	*/
	Count *int32

	/* Ids.

	   Specifies a list of object ids, only last runs for these objects will be returned.
	*/
	Ids []int64

	/* IncludeTenants.

	   If true, the response will include Objects which belongs to all tenants which the current user has permission to see.
	*/
	IncludeTenants *bool

	/* PaginationCookie.

	   Specifies the pagination cookie with which subsequent parts of the response can be fetched.
	*/
	PaginationCookie *string

	/* TenantIds.

	   TenantIds contains ids of the tenants for which objects are to be returned.
	*/
	TenantIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get objects last run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetObjectsLastRunParams) WithDefaults() *GetObjectsLastRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get objects last run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetObjectsLastRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get objects last run params
func (o *GetObjectsLastRunParams) WithTimeout(timeout time.Duration) *GetObjectsLastRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get objects last run params
func (o *GetObjectsLastRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get objects last run params
func (o *GetObjectsLastRunParams) WithContext(ctx context.Context) *GetObjectsLastRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get objects last run params
func (o *GetObjectsLastRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get objects last run params
func (o *GetObjectsLastRunParams) WithHTTPClient(client *http.Client) *GetObjectsLastRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get objects last run params
func (o *GetObjectsLastRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get objects last run params
func (o *GetObjectsLastRunParams) WithCount(count *int32) *GetObjectsLastRunParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get objects last run params
func (o *GetObjectsLastRunParams) SetCount(count *int32) {
	o.Count = count
}

// WithIds adds the ids to the get objects last run params
func (o *GetObjectsLastRunParams) WithIds(ids []int64) *GetObjectsLastRunParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get objects last run params
func (o *GetObjectsLastRunParams) SetIds(ids []int64) {
	o.Ids = ids
}

// WithIncludeTenants adds the includeTenants to the get objects last run params
func (o *GetObjectsLastRunParams) WithIncludeTenants(includeTenants *bool) *GetObjectsLastRunParams {
	o.SetIncludeTenants(includeTenants)
	return o
}

// SetIncludeTenants adds the includeTenants to the get objects last run params
func (o *GetObjectsLastRunParams) SetIncludeTenants(includeTenants *bool) {
	o.IncludeTenants = includeTenants
}

// WithPaginationCookie adds the paginationCookie to the get objects last run params
func (o *GetObjectsLastRunParams) WithPaginationCookie(paginationCookie *string) *GetObjectsLastRunParams {
	o.SetPaginationCookie(paginationCookie)
	return o
}

// SetPaginationCookie adds the paginationCookie to the get objects last run params
func (o *GetObjectsLastRunParams) SetPaginationCookie(paginationCookie *string) {
	o.PaginationCookie = paginationCookie
}

// WithTenantIds adds the tenantIds to the get objects last run params
func (o *GetObjectsLastRunParams) WithTenantIds(tenantIds []string) *GetObjectsLastRunParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get objects last run params
func (o *GetObjectsLastRunParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetObjectsLastRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
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

	if o.PaginationCookie != nil {

		// query param paginationCookie
		var qrPaginationCookie string

		if o.PaginationCookie != nil {
			qrPaginationCookie = *o.PaginationCookie
		}
		qPaginationCookie := qrPaginationCookie
		if qPaginationCookie != "" {

			if err := r.SetQueryParam("paginationCookie", qPaginationCookie); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetObjectsLastRun binds the parameter ids
func (o *GetObjectsLastRunParams) bindParamIds(formats strfmt.Registry) []string {
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

// bindParamGetObjectsLastRun binds the parameter tenantIds
func (o *GetObjectsLastRunParams) bindParamTenantIds(formats strfmt.Registry) []string {
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
