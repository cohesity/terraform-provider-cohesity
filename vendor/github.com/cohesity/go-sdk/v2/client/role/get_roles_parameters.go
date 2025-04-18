// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewGetRolesParams creates a new GetRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRolesParams() *GetRolesParams {
	return &GetRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRolesParamsWithTimeout creates a new GetRolesParams object
// with the ability to set a timeout on a request.
func NewGetRolesParamsWithTimeout(timeout time.Duration) *GetRolesParams {
	return &GetRolesParams{
		timeout: timeout,
	}
}

// NewGetRolesParamsWithContext creates a new GetRolesParams object
// with the ability to set a context for a request.
func NewGetRolesParamsWithContext(ctx context.Context) *GetRolesParams {
	return &GetRolesParams{
		Context: ctx,
	}
}

// NewGetRolesParamsWithHTTPClient creates a new GetRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRolesParamsWithHTTPClient(client *http.Client) *GetRolesParams {
	return &GetRolesParams{
		HTTPClient: client,
	}
}

/*
GetRolesParams contains all the parameters to send to the API endpoint

	for the get roles operation.

	Typically these are written to a http.Request.
*/
type GetRolesParams struct {

	/* IncludeTenants.

	   If true, the response will include Roles which were created by all tenants which the current user has permission to see. If false, then only Roles created by the current user will be returned.
	*/
	IncludeTenants *bool

	/* Names.

	   Specifies a list of Role names.
	*/
	Names []string

	/* TenantIds.

	   TenantIds contains ids of the tenants for which Roles are to be returned.
	*/
	TenantIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesParams) WithDefaults() *GetRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get roles params
func (o *GetRolesParams) WithTimeout(timeout time.Duration) *GetRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get roles params
func (o *GetRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get roles params
func (o *GetRolesParams) WithContext(ctx context.Context) *GetRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get roles params
func (o *GetRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) WithHTTPClient(client *http.Client) *GetRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get roles params
func (o *GetRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeTenants adds the includeTenants to the get roles params
func (o *GetRolesParams) WithIncludeTenants(includeTenants *bool) *GetRolesParams {
	o.SetIncludeTenants(includeTenants)
	return o
}

// SetIncludeTenants adds the includeTenants to the get roles params
func (o *GetRolesParams) SetIncludeTenants(includeTenants *bool) {
	o.IncludeTenants = includeTenants
}

// WithNames adds the names to the get roles params
func (o *GetRolesParams) WithNames(names []string) *GetRolesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get roles params
func (o *GetRolesParams) SetNames(names []string) {
	o.Names = names
}

// WithTenantIds adds the tenantIds to the get roles params
func (o *GetRolesParams) WithTenantIds(tenantIds []string) *GetRolesParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get roles params
func (o *GetRolesParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetRoles binds the parameter names
func (o *GetRolesParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamGetRoles binds the parameter tenantIds
func (o *GetRolesParams) bindParamTenantIds(formats strfmt.Registry) []string {
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
