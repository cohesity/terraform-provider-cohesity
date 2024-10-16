// Code generated by go-swagger; DO NOT EDIT.

package source

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

// NewGetProtectionSourcesParams creates a new GetProtectionSourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProtectionSourcesParams() *GetProtectionSourcesParams {
	return &GetProtectionSourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtectionSourcesParamsWithTimeout creates a new GetProtectionSourcesParams object
// with the ability to set a timeout on a request.
func NewGetProtectionSourcesParamsWithTimeout(timeout time.Duration) *GetProtectionSourcesParams {
	return &GetProtectionSourcesParams{
		timeout: timeout,
	}
}

// NewGetProtectionSourcesParamsWithContext creates a new GetProtectionSourcesParams object
// with the ability to set a context for a request.
func NewGetProtectionSourcesParamsWithContext(ctx context.Context) *GetProtectionSourcesParams {
	return &GetProtectionSourcesParams{
		Context: ctx,
	}
}

// NewGetProtectionSourcesParamsWithHTTPClient creates a new GetProtectionSourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProtectionSourcesParamsWithHTTPClient(client *http.Client) *GetProtectionSourcesParams {
	return &GetProtectionSourcesParams{
		HTTPClient: client,
	}
}

/*
GetProtectionSourcesParams contains all the parameters to send to the API endpoint

	for the get protection sources operation.

	Typically these are written to a http.Request.
*/
type GetProtectionSourcesParams struct {

	/* EncryptionKey.

	   Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
	*/
	EncryptionKey *string

	/* IncludeSourceCredentials.

	   If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key.
	*/
	IncludeSourceCredentials *bool

	/* IncludeTenants.

	   If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned.
	*/
	IncludeTenants *bool

	/* RequestInitiatorType.

	   Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
	*/
	RequestInitiatorType *string

	/* TenantIds.

	   TenantIds contains ids of the tenants for which Sources are to be returned.
	*/
	TenantIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get protection sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionSourcesParams) WithDefaults() *GetProtectionSourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get protection sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionSourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get protection sources params
func (o *GetProtectionSourcesParams) WithTimeout(timeout time.Duration) *GetProtectionSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protection sources params
func (o *GetProtectionSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protection sources params
func (o *GetProtectionSourcesParams) WithContext(ctx context.Context) *GetProtectionSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protection sources params
func (o *GetProtectionSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protection sources params
func (o *GetProtectionSourcesParams) WithHTTPClient(client *http.Client) *GetProtectionSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protection sources params
func (o *GetProtectionSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncryptionKey adds the encryptionKey to the get protection sources params
func (o *GetProtectionSourcesParams) WithEncryptionKey(encryptionKey *string) *GetProtectionSourcesParams {
	o.SetEncryptionKey(encryptionKey)
	return o
}

// SetEncryptionKey adds the encryptionKey to the get protection sources params
func (o *GetProtectionSourcesParams) SetEncryptionKey(encryptionKey *string) {
	o.EncryptionKey = encryptionKey
}

// WithIncludeSourceCredentials adds the includeSourceCredentials to the get protection sources params
func (o *GetProtectionSourcesParams) WithIncludeSourceCredentials(includeSourceCredentials *bool) *GetProtectionSourcesParams {
	o.SetIncludeSourceCredentials(includeSourceCredentials)
	return o
}

// SetIncludeSourceCredentials adds the includeSourceCredentials to the get protection sources params
func (o *GetProtectionSourcesParams) SetIncludeSourceCredentials(includeSourceCredentials *bool) {
	o.IncludeSourceCredentials = includeSourceCredentials
}

// WithIncludeTenants adds the includeTenants to the get protection sources params
func (o *GetProtectionSourcesParams) WithIncludeTenants(includeTenants *bool) *GetProtectionSourcesParams {
	o.SetIncludeTenants(includeTenants)
	return o
}

// SetIncludeTenants adds the includeTenants to the get protection sources params
func (o *GetProtectionSourcesParams) SetIncludeTenants(includeTenants *bool) {
	o.IncludeTenants = includeTenants
}

// WithRequestInitiatorType adds the requestInitiatorType to the get protection sources params
func (o *GetProtectionSourcesParams) WithRequestInitiatorType(requestInitiatorType *string) *GetProtectionSourcesParams {
	o.SetRequestInitiatorType(requestInitiatorType)
	return o
}

// SetRequestInitiatorType adds the requestInitiatorType to the get protection sources params
func (o *GetProtectionSourcesParams) SetRequestInitiatorType(requestInitiatorType *string) {
	o.RequestInitiatorType = requestInitiatorType
}

// WithTenantIds adds the tenantIds to the get protection sources params
func (o *GetProtectionSourcesParams) WithTenantIds(tenantIds []string) *GetProtectionSourcesParams {
	o.SetTenantIds(tenantIds)
	return o
}

// SetTenantIds adds the tenantIds to the get protection sources params
func (o *GetProtectionSourcesParams) SetTenantIds(tenantIds []string) {
	o.TenantIds = tenantIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtectionSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EncryptionKey != nil {

		// query param encryptionKey
		var qrEncryptionKey string

		if o.EncryptionKey != nil {
			qrEncryptionKey = *o.EncryptionKey
		}
		qEncryptionKey := qrEncryptionKey
		if qEncryptionKey != "" {

			if err := r.SetQueryParam("encryptionKey", qEncryptionKey); err != nil {
				return err
			}
		}
	}

	if o.IncludeSourceCredentials != nil {

		// query param includeSourceCredentials
		var qrIncludeSourceCredentials bool

		if o.IncludeSourceCredentials != nil {
			qrIncludeSourceCredentials = *o.IncludeSourceCredentials
		}
		qIncludeSourceCredentials := swag.FormatBool(qrIncludeSourceCredentials)
		if qIncludeSourceCredentials != "" {

			if err := r.SetQueryParam("includeSourceCredentials", qIncludeSourceCredentials); err != nil {
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

	if o.RequestInitiatorType != nil {

		// header param requestInitiatorType
		if err := r.SetHeaderParam("requestInitiatorType", *o.RequestInitiatorType); err != nil {
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

// bindParamGetProtectionSources binds the parameter tenantIds
func (o *GetProtectionSourcesParams) bindParamTenantIds(formats strfmt.Registry) []string {
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
