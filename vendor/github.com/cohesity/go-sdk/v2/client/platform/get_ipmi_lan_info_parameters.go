// Code generated by go-swagger; DO NOT EDIT.

package platform

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
)

// NewGetIpmiLanInfoParams creates a new GetIpmiLanInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIpmiLanInfoParams() *GetIpmiLanInfoParams {
	return &GetIpmiLanInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIpmiLanInfoParamsWithTimeout creates a new GetIpmiLanInfoParams object
// with the ability to set a timeout on a request.
func NewGetIpmiLanInfoParamsWithTimeout(timeout time.Duration) *GetIpmiLanInfoParams {
	return &GetIpmiLanInfoParams{
		timeout: timeout,
	}
}

// NewGetIpmiLanInfoParamsWithContext creates a new GetIpmiLanInfoParams object
// with the ability to set a context for a request.
func NewGetIpmiLanInfoParamsWithContext(ctx context.Context) *GetIpmiLanInfoParams {
	return &GetIpmiLanInfoParams{
		Context: ctx,
	}
}

// NewGetIpmiLanInfoParamsWithHTTPClient creates a new GetIpmiLanInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIpmiLanInfoParamsWithHTTPClient(client *http.Client) *GetIpmiLanInfoParams {
	return &GetIpmiLanInfoParams{
		HTTPClient: client,
	}
}

/*
GetIpmiLanInfoParams contains all the parameters to send to the API endpoint

	for the get ipmi lan info operation.

	Typically these are written to a http.Request.
*/
type GetIpmiLanInfoParams struct {

	/* NodeID.

	   Specifies the node id of the node for which lan info is requested. This parameter is incompatible with 'nodeIp'.
	*/
	NodeID *string

	/* NodeIP.

	   Specifies the IP Address of the node for which lan info is requested. This parameter is incompatible with 'nodeId'.
	*/
	NodeIP *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ipmi lan info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpmiLanInfoParams) WithDefaults() *GetIpmiLanInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ipmi lan info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpmiLanInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) WithTimeout(timeout time.Duration) *GetIpmiLanInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) WithContext(ctx context.Context) *GetIpmiLanInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) WithHTTPClient(client *http.Client) *GetIpmiLanInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeID adds the nodeID to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) WithNodeID(nodeID *string) *GetIpmiLanInfoParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) SetNodeID(nodeID *string) {
	o.NodeID = nodeID
}

// WithNodeIP adds the nodeIP to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) WithNodeIP(nodeIP *string) *GetIpmiLanInfoParams {
	o.SetNodeIP(nodeIP)
	return o
}

// SetNodeIP adds the nodeIp to the get ipmi lan info params
func (o *GetIpmiLanInfoParams) SetNodeIP(nodeIP *string) {
	o.NodeIP = nodeIP
}

// WriteToRequest writes these params to a swagger request
func (o *GetIpmiLanInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NodeID != nil {

		// query param nodeId
		var qrNodeID string

		if o.NodeID != nil {
			qrNodeID = *o.NodeID
		}
		qNodeID := qrNodeID
		if qNodeID != "" {

			if err := r.SetQueryParam("nodeId", qNodeID); err != nil {
				return err
			}
		}
	}

	if o.NodeIP != nil {

		// query param nodeIp
		var qrNodeIP string

		if o.NodeIP != nil {
			qrNodeIP = *o.NodeIP
		}
		qNodeIP := qrNodeIP
		if qNodeIP != "" {

			if err := r.SetQueryParam("nodeIp", qNodeIP); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
