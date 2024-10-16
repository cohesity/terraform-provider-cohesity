// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// NewGetDataTransferToVaultsReportRequestParams creates a new GetDataTransferToVaultsReportRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataTransferToVaultsReportRequestParams() *GetDataTransferToVaultsReportRequestParams {
	return &GetDataTransferToVaultsReportRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataTransferToVaultsReportRequestParamsWithTimeout creates a new GetDataTransferToVaultsReportRequestParams object
// with the ability to set a timeout on a request.
func NewGetDataTransferToVaultsReportRequestParamsWithTimeout(timeout time.Duration) *GetDataTransferToVaultsReportRequestParams {
	return &GetDataTransferToVaultsReportRequestParams{
		timeout: timeout,
	}
}

// NewGetDataTransferToVaultsReportRequestParamsWithContext creates a new GetDataTransferToVaultsReportRequestParams object
// with the ability to set a context for a request.
func NewGetDataTransferToVaultsReportRequestParamsWithContext(ctx context.Context) *GetDataTransferToVaultsReportRequestParams {
	return &GetDataTransferToVaultsReportRequestParams{
		Context: ctx,
	}
}

// NewGetDataTransferToVaultsReportRequestParamsWithHTTPClient creates a new GetDataTransferToVaultsReportRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataTransferToVaultsReportRequestParamsWithHTTPClient(client *http.Client) *GetDataTransferToVaultsReportRequestParams {
	return &GetDataTransferToVaultsReportRequestParams{
		HTTPClient: client,
	}
}

/*
GetDataTransferToVaultsReportRequestParams contains all the parameters to send to the API endpoint

	for the get data transfer to vaults report request operation.

	Typically these are written to a http.Request.
*/
type GetDataTransferToVaultsReportRequestParams struct {

	/* EndTimeMsecs.

	     Filter by end time. Specify the end time as a Unix epoch Timestamp
	(in milliseconds).
	If startTimeMsecs and endTimeMsecs are not specified,
	the time period is the last 7 days.

	     Format: int64
	*/
	EndTimeMsecs *int64

	/* GroupBy.

	     Specifies wheather the report should be grouped by target when scheduled
	or downloaded. If not set or set to false, report is grouped by protection
	jobs. It is ignored if outformat is not "csv" and response contains whole
	report.

	     Format: int32
	*/
	GroupBy *int32

	/* OutputFormat.

	     Specifies the format for the output such as 'csv' or 'json'.
	If not specified, the json format is returned.
	If 'csv' is specified, a comma-separated list with a heading
	row is returned.
	*/
	OutputFormat *string

	/* StartTimeMsecs.

	     Filter by a start time. Specify the start time as a Unix epoch Timestamp
	(in milliseconds).
	If startTimeMsecs and endTimeMsecs are not specified,
	the time period is the last 7 days.

	     Format: int64
	*/
	StartTimeMsecs *int64

	/* VaultIds.

	   Filter by a list of Vault ids.
	*/
	VaultIds []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get data transfer to vaults report request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataTransferToVaultsReportRequestParams) WithDefaults() *GetDataTransferToVaultsReportRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get data transfer to vaults report request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataTransferToVaultsReportRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithTimeout(timeout time.Duration) *GetDataTransferToVaultsReportRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithContext(ctx context.Context) *GetDataTransferToVaultsReportRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithHTTPClient(client *http.Client) *GetDataTransferToVaultsReportRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTimeMsecs adds the endTimeMsecs to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithEndTimeMsecs(endTimeMsecs *int64) *GetDataTransferToVaultsReportRequestParams {
	o.SetEndTimeMsecs(endTimeMsecs)
	return o
}

// SetEndTimeMsecs adds the endTimeMsecs to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetEndTimeMsecs(endTimeMsecs *int64) {
	o.EndTimeMsecs = endTimeMsecs
}

// WithGroupBy adds the groupBy to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithGroupBy(groupBy *int32) *GetDataTransferToVaultsReportRequestParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetGroupBy(groupBy *int32) {
	o.GroupBy = groupBy
}

// WithOutputFormat adds the outputFormat to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithOutputFormat(outputFormat *string) *GetDataTransferToVaultsReportRequestParams {
	o.SetOutputFormat(outputFormat)
	return o
}

// SetOutputFormat adds the outputFormat to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetOutputFormat(outputFormat *string) {
	o.OutputFormat = outputFormat
}

// WithStartTimeMsecs adds the startTimeMsecs to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithStartTimeMsecs(startTimeMsecs *int64) *GetDataTransferToVaultsReportRequestParams {
	o.SetStartTimeMsecs(startTimeMsecs)
	return o
}

// SetStartTimeMsecs adds the startTimeMsecs to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetStartTimeMsecs(startTimeMsecs *int64) {
	o.StartTimeMsecs = startTimeMsecs
}

// WithVaultIds adds the vaultIds to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) WithVaultIds(vaultIds []int64) *GetDataTransferToVaultsReportRequestParams {
	o.SetVaultIds(vaultIds)
	return o
}

// SetVaultIds adds the vaultIds to the get data transfer to vaults report request params
func (o *GetDataTransferToVaultsReportRequestParams) SetVaultIds(vaultIds []int64) {
	o.VaultIds = vaultIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataTransferToVaultsReportRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndTimeMsecs != nil {

		// query param endTimeMsecs
		var qrEndTimeMsecs int64

		if o.EndTimeMsecs != nil {
			qrEndTimeMsecs = *o.EndTimeMsecs
		}
		qEndTimeMsecs := swag.FormatInt64(qrEndTimeMsecs)
		if qEndTimeMsecs != "" {

			if err := r.SetQueryParam("endTimeMsecs", qEndTimeMsecs); err != nil {
				return err
			}
		}
	}

	if o.GroupBy != nil {

		// query param groupBy
		var qrGroupBy int32

		if o.GroupBy != nil {
			qrGroupBy = *o.GroupBy
		}
		qGroupBy := swag.FormatInt32(qrGroupBy)
		if qGroupBy != "" {

			if err := r.SetQueryParam("groupBy", qGroupBy); err != nil {
				return err
			}
		}
	}

	if o.OutputFormat != nil {

		// query param outputFormat
		var qrOutputFormat string

		if o.OutputFormat != nil {
			qrOutputFormat = *o.OutputFormat
		}
		qOutputFormat := qrOutputFormat
		if qOutputFormat != "" {

			if err := r.SetQueryParam("outputFormat", qOutputFormat); err != nil {
				return err
			}
		}
	}

	if o.StartTimeMsecs != nil {

		// query param startTimeMsecs
		var qrStartTimeMsecs int64

		if o.StartTimeMsecs != nil {
			qrStartTimeMsecs = *o.StartTimeMsecs
		}
		qStartTimeMsecs := swag.FormatInt64(qrStartTimeMsecs)
		if qStartTimeMsecs != "" {

			if err := r.SetQueryParam("startTimeMsecs", qStartTimeMsecs); err != nil {
				return err
			}
		}
	}

	if o.VaultIds != nil {

		// binding items for vaultIds
		joinedVaultIds := o.bindParamVaultIds(reg)

		// query array param vaultIds
		if err := r.SetQueryParam("vaultIds", joinedVaultIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDataTransferToVaultsReportRequest binds the parameter vaultIds
func (o *GetDataTransferToVaultsReportRequestParams) bindParamVaultIds(formats strfmt.Registry) []string {
	vaultIdsIR := o.VaultIds

	var vaultIdsIC []string
	for _, vaultIdsIIR := range vaultIdsIR { // explode []int64

		vaultIdsIIV := swag.FormatInt64(vaultIdsIIR) // int64 as string
		vaultIdsIC = append(vaultIdsIC, vaultIdsIIV)
	}

	// items.CollectionFormat: ""
	vaultIdsIS := swag.JoinByFormat(vaultIdsIC, "")

	return vaultIdsIS
}
