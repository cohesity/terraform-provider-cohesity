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

// NewGetSharesParams creates a new GetSharesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSharesParams() *GetSharesParams {
	return &GetSharesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSharesParamsWithTimeout creates a new GetSharesParams object
// with the ability to set a timeout on a request.
func NewGetSharesParamsWithTimeout(timeout time.Duration) *GetSharesParams {
	return &GetSharesParams{
		timeout: timeout,
	}
}

// NewGetSharesParamsWithContext creates a new GetSharesParams object
// with the ability to set a context for a request.
func NewGetSharesParamsWithContext(ctx context.Context) *GetSharesParams {
	return &GetSharesParams{
		Context: ctx,
	}
}

// NewGetSharesParamsWithHTTPClient creates a new GetSharesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSharesParamsWithHTTPClient(client *http.Client) *GetSharesParams {
	return &GetSharesParams{
		HTTPClient: client,
	}
}

/*
GetSharesParams contains all the parameters to send to the API endpoint

	for the get shares operation.

	Typically these are written to a http.Request.
*/
type GetSharesParams struct {

	/* Cookie.

	   Specifies the pagination cookie. Expected to be empty in the first call to the API. To get the next set of results, set this value to the pagination cookie value returned in the response of the previous call.
	*/
	Cookie *string

	/* IncludeTenants.

	   IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned.
	*/
	IncludeTenants *bool

	/* MatchPartialName.

	   If true, the share name is matched by any partial rather than exactly matched.
	*/
	MatchPartialName *bool

	/* MaxCount.

	   Specifies a limit on the number of Shares returned. If maxCount is not specified, the first 2000 Shares.

	   Format: int32
	*/
	MaxCount *int32

	/* Name.

	   Specifies the Share name.
	*/
	Name *string

	/* TenantIds.

	   TenantIds contains ids of the tenants for which objects are to be returned.
	*/
	TenantIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get shares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSharesParams) WithDefaults() *GetSharesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get shares params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSharesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get shares params
func (o *GetSharesParams) WithTimeout(timeout time.Duration) *GetSharesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get shares params
func (o *GetSharesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get shares params
func (o *GetSharesParams) WithContext(ctx context.Context) *GetSharesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get shares params
func (o *GetSharesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get shares params
func (o *GetSharesParams) WithHTTPClient(client *http.Client) *GetSharesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get shares params
func (o *GetSharesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCookie adds the cookie to the get shares params
func (o *GetSharesParams) WithCookie(cookie *string) *GetSharesParams {
	o.SetCookie(cookie)
	return o
}

// SetCookie adds the cookie to the get shares params
func (o *GetSharesParams) SetCookie(cookie *string) {
	o.Cookie = cookie
}

// WithIncludeTenants adds the includeTenants to the get shares params
func (o *GetSharesParams) WithIncludeTenants(includeTenants *bool) *GetSharesParams {
	o.SetIncludeTenants(includeTenants)
	return o
}

// SetIncludeTenants adds the includeTenants to the get shares params
func (o *GetSharesParams) SetIncludeTenants(includeTenants *bool) {
	o.IncludeTenants = includeTenants
}

// WithMatchPartialName adds the matchPartialName to the get shares params
func (o *GetSharesParams) WithMatchPartialName(matchPartialName *bool) *GetSharesParams {
	o.SetMatchPartialName(matchPartialName)
	return o
}

// SetMatchPartialName adds the matchPartialName to the get shares params
func (o *GetSharesParams) SetMatchPartialName(matchPartialName *bool) {
	o.MatchPartialName = matchPartialName
}

// WithMaxCount adds the maxCount to the get shares params
func (o *GetSharesParams) WithMaxCount(maxCount *int32) *GetSharesParams {
	o.SetMaxCount(maxCount)
	return o
}

// SetMaxCount adds the maxCount to the get shares params
func (o *GetSharesParams) SetMaxCount(maxCount *int32) {
	o.MaxCount = maxCount
}

// WithName adds the name to the get shares params
func (o *GetSharesParams) WithName(name *string) *GetSharesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get shares params
func (o *GetSharesParams) SetName(name *string) {
	o.Name = name
}

// WithTenantIds adds the tenantIds to the get shares params
func (o *GetSharesParams) WithTenantIds(tenantIds []string) *GetSharesParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get shares params
func (o *GetSharesParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetSharesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cookie != nil {

		// query param cookie
		var qrCookie string

		if o.Cookie != nil {
			qrCookie = *o.Cookie
		}
		qCookie := qrCookie
		if qCookie != "" {

			if err := r.SetQueryParam("cookie", qCookie); err != nil {
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

	if o.MatchPartialName != nil {

		// query param matchPartialName
		var qrMatchPartialName bool

		if o.MatchPartialName != nil {
			qrMatchPartialName = *o.MatchPartialName
		}
		qMatchPartialName := swag.FormatBool(qrMatchPartialName)
		if qMatchPartialName != "" {

			if err := r.SetQueryParam("matchPartialName", qMatchPartialName); err != nil {
				return err
			}
		}
	}

	if o.MaxCount != nil {

		// query param maxCount
		var qrMaxCount int32

		if o.MaxCount != nil {
			qrMaxCount = *o.MaxCount
		}
		qMaxCount := swag.FormatInt32(qrMaxCount)
		if qMaxCount != "" {

			if err := r.SetQueryParam("maxCount", qMaxCount); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

// bindParamGetShares binds the parameter tenantIds
func (o *GetSharesParams) bindParamTenantIds(formats strfmt.Registry) []string {
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
