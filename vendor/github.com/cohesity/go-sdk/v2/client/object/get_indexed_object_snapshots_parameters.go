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

// NewGetIndexedObjectSnapshotsParams creates a new GetIndexedObjectSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIndexedObjectSnapshotsParams() *GetIndexedObjectSnapshotsParams {
	return &GetIndexedObjectSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIndexedObjectSnapshotsParamsWithTimeout creates a new GetIndexedObjectSnapshotsParams object
// with the ability to set a timeout on a request.
func NewGetIndexedObjectSnapshotsParamsWithTimeout(timeout time.Duration) *GetIndexedObjectSnapshotsParams {
	return &GetIndexedObjectSnapshotsParams{
		timeout: timeout,
	}
}

// NewGetIndexedObjectSnapshotsParamsWithContext creates a new GetIndexedObjectSnapshotsParams object
// with the ability to set a context for a request.
func NewGetIndexedObjectSnapshotsParamsWithContext(ctx context.Context) *GetIndexedObjectSnapshotsParams {
	return &GetIndexedObjectSnapshotsParams{
		Context: ctx,
	}
}

// NewGetIndexedObjectSnapshotsParamsWithHTTPClient creates a new GetIndexedObjectSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIndexedObjectSnapshotsParamsWithHTTPClient(client *http.Client) *GetIndexedObjectSnapshotsParams {
	return &GetIndexedObjectSnapshotsParams{
		HTTPClient: client,
	}
}

/*
GetIndexedObjectSnapshotsParams contains all the parameters to send to the API endpoint

	for the get indexed object snapshots operation.

	Typically these are written to a http.Request.
*/
type GetIndexedObjectSnapshotsParams struct {

	/* FromTimeUsecs.

	   Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken after this value.

	   Format: int64
	*/
	FromTimeUsecs *int64

	/* IncludeIndexedSnapshotsOnly.

	   Specifies whether to only return snapshots which are indexed. In an indexed snapshots file are guaranteed to exist, while in a non-indexed snapshots file may not exist.
	*/
	IncludeIndexedSnapshotsOnly *bool

	/* IndexedObjectName.

	   Specifies the indexed object name.
	*/
	IndexedObjectName string

	/* ObjectActionKey.

	   Filter by ObjectActionKey, which uniquely represents backup type for a given version. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey and ObjectId. When specified, only versions of given ObjectActionKey are returned for corresponding object id.
	*/
	ObjectActionKey *string

	/* ObjectID.

	   Specifies the object id.

	   Format: int64
	*/
	ObjectID int64

	/* ProtectionGroupID.

	   Specifies the protection group id.
	*/
	ProtectionGroupID string

	/* RunTypes.

	   Filter by run type. Only protection run matching the specified types will be returned. By default, CDP hydration snapshots are not included, unless explicitly queried using this field.
	*/
	RunTypes []string

	/* ToTimeUsecs.

	   Specifies the timestamp in Unix time epoch in microseconds to filter indexed object's snapshots which are taken before this value.

	   Format: int64
	*/
	ToTimeUsecs *int64

	/* UseCachedData.

	   Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.
	*/
	UseCachedData *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get indexed object snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIndexedObjectSnapshotsParams) WithDefaults() *GetIndexedObjectSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get indexed object snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIndexedObjectSnapshotsParams) SetDefaults() {
	var (
		includeIndexedSnapshotsOnlyDefault = bool(false)
	)

	val := GetIndexedObjectSnapshotsParams{
		IncludeIndexedSnapshotsOnly: &includeIndexedSnapshotsOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithTimeout(timeout time.Duration) *GetIndexedObjectSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithContext(ctx context.Context) *GetIndexedObjectSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithHTTPClient(client *http.Client) *GetIndexedObjectSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFromTimeUsecs adds the fromTimeUsecs to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithFromTimeUsecs(fromTimeUsecs *int64) *GetIndexedObjectSnapshotsParams {
	o.SetFromTimeUsecs(fromTimeUsecs)
	return o
}

// SetFromTimeUsecs adds the fromTimeUsecs to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetFromTimeUsecs(fromTimeUsecs *int64) {
	o.FromTimeUsecs = fromTimeUsecs
}

// WithIncludeIndexedSnapshotsOnly adds the includeIndexedSnapshotsOnly to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithIncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly *bool) *GetIndexedObjectSnapshotsParams {
	o.SetIncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly)
	return o
}

// SetIncludeIndexedSnapshotsOnly adds the includeIndexedSnapshotsOnly to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetIncludeIndexedSnapshotsOnly(includeIndexedSnapshotsOnly *bool) {
	o.IncludeIndexedSnapshotsOnly = includeIndexedSnapshotsOnly
}

// WithIndexedObjectName adds the indexedObjectName to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithIndexedObjectName(indexedObjectName string) *GetIndexedObjectSnapshotsParams {
	o.SetIndexedObjectName(indexedObjectName)
	return o
}

// SetIndexedObjectName adds the indexedObjectName to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetIndexedObjectName(indexedObjectName string) {
	o.IndexedObjectName = indexedObjectName
}

// WithObjectActionKey adds the objectActionKey to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithObjectActionKey(objectActionKey *string) *GetIndexedObjectSnapshotsParams {
	o.SetObjectActionKey(objectActionKey)
	return o
}

// SetObjectActionKey adds the objectActionKey to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetObjectActionKey(objectActionKey *string) {
	o.ObjectActionKey = objectActionKey
}

// WithObjectID adds the objectID to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithObjectID(objectID int64) *GetIndexedObjectSnapshotsParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetObjectID(objectID int64) {
	o.ObjectID = objectID
}

// WithProtectionGroupID adds the protectionGroupID to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithProtectionGroupID(protectionGroupID string) *GetIndexedObjectSnapshotsParams {
	o.SetProtectionGroupID(protectionGroupID)
	return o
}

// SetProtectionGroupID adds the protectionGroupId to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetProtectionGroupID(protectionGroupID string) {
	o.ProtectionGroupID = protectionGroupID
}

// WithRunTypes adds the runTypes to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithRunTypes(runTypes []string) *GetIndexedObjectSnapshotsParams {
	o.SetRunTypes(runTypes)
	return o
}

// SetRunTypes adds the runTypes to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetRunTypes(runTypes []string) {
	o.RunTypes = runTypes
}

// WithToTimeUsecs adds the toTimeUsecs to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithToTimeUsecs(toTimeUsecs *int64) *GetIndexedObjectSnapshotsParams {
	o.SetToTimeUsecs(toTimeUsecs)
	return o
}

// SetToTimeUsecs adds the toTimeUsecs to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetToTimeUsecs(toTimeUsecs *int64) {
	o.ToTimeUsecs = toTimeUsecs
}

// WithUseCachedData adds the useCachedData to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) WithUseCachedData(useCachedData *bool) *GetIndexedObjectSnapshotsParams {
	o.SetUseCachedData(useCachedData)
	return o
}

// SetUseCachedData adds the useCachedData to the get indexed object snapshots params
func (o *GetIndexedObjectSnapshotsParams) SetUseCachedData(useCachedData *bool) {
	o.UseCachedData = useCachedData
}

// WriteToRequest writes these params to a swagger request
func (o *GetIndexedObjectSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FromTimeUsecs != nil {

		// query param fromTimeUsecs
		var qrFromTimeUsecs int64

		if o.FromTimeUsecs != nil {
			qrFromTimeUsecs = *o.FromTimeUsecs
		}
		qFromTimeUsecs := swag.FormatInt64(qrFromTimeUsecs)
		if qFromTimeUsecs != "" {

			if err := r.SetQueryParam("fromTimeUsecs", qFromTimeUsecs); err != nil {
				return err
			}
		}
	}

	if o.IncludeIndexedSnapshotsOnly != nil {

		// query param includeIndexedSnapshotsOnly
		var qrIncludeIndexedSnapshotsOnly bool

		if o.IncludeIndexedSnapshotsOnly != nil {
			qrIncludeIndexedSnapshotsOnly = *o.IncludeIndexedSnapshotsOnly
		}
		qIncludeIndexedSnapshotsOnly := swag.FormatBool(qrIncludeIndexedSnapshotsOnly)
		if qIncludeIndexedSnapshotsOnly != "" {

			if err := r.SetQueryParam("includeIndexedSnapshotsOnly", qIncludeIndexedSnapshotsOnly); err != nil {
				return err
			}
		}
	}

	// query param indexedObjectName
	qrIndexedObjectName := o.IndexedObjectName
	qIndexedObjectName := qrIndexedObjectName
	if qIndexedObjectName != "" {

		if err := r.SetQueryParam("indexedObjectName", qIndexedObjectName); err != nil {
			return err
		}
	}

	if o.ObjectActionKey != nil {

		// query param objectActionKey
		var qrObjectActionKey string

		if o.ObjectActionKey != nil {
			qrObjectActionKey = *o.ObjectActionKey
		}
		qObjectActionKey := qrObjectActionKey
		if qObjectActionKey != "" {

			if err := r.SetQueryParam("objectActionKey", qObjectActionKey); err != nil {
				return err
			}
		}
	}

	// path param objectId
	if err := r.SetPathParam("objectId", swag.FormatInt64(o.ObjectID)); err != nil {
		return err
	}

	// path param protectionGroupId
	if err := r.SetPathParam("protectionGroupId", o.ProtectionGroupID); err != nil {
		return err
	}

	if o.RunTypes != nil {

		// binding items for runTypes
		joinedRunTypes := o.bindParamRunTypes(reg)

		// query array param runTypes
		if err := r.SetQueryParam("runTypes", joinedRunTypes...); err != nil {
			return err
		}
	}

	if o.ToTimeUsecs != nil {

		// query param toTimeUsecs
		var qrToTimeUsecs int64

		if o.ToTimeUsecs != nil {
			qrToTimeUsecs = *o.ToTimeUsecs
		}
		qToTimeUsecs := swag.FormatInt64(qrToTimeUsecs)
		if qToTimeUsecs != "" {

			if err := r.SetQueryParam("toTimeUsecs", qToTimeUsecs); err != nil {
				return err
			}
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

// bindParamGetIndexedObjectSnapshots binds the parameter runTypes
func (o *GetIndexedObjectSnapshotsParams) bindParamRunTypes(formats strfmt.Registry) []string {
	runTypesIR := o.RunTypes

	var runTypesIC []string
	for _, runTypesIIR := range runTypesIR { // explode []string

		runTypesIIV := runTypesIIR // string as string
		runTypesIC = append(runTypesIC, runTypesIIV)
	}

	// items.CollectionFormat: ""
	runTypesIS := swag.JoinByFormat(runTypesIC, "")

	return runTypesIS
}
