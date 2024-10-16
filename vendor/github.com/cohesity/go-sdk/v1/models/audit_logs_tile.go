// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditLogsTile Audit Logs Tile.
//
// Audit logs for Dashboard.
//
// swagger:model AuditLogsTile
type AuditLogsTile struct {

	// Array of Cluster Audit Logs.
	//
	// Specifies a list of Cluster audit logs that match the specified
	// filter criteria up to the limit specified in pageCount.
	ClusterAuditLogs []*ClusterAuditLog `json:"clusterAuditLogs"`

	// Specifies the total number of logs that match the specified
	// filter criteria. (This number might be larger than the size of the
	// Cluster Audit Logs array.)
	// This count is provided to indicate if additional requests must be
	// made to get the full result.
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// Validate validates this audit logs tile
func (m *AuditLogsTile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLogsTile) validateClusterAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterAuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterAuditLogs); i++ {
		if swag.IsZero(m.ClusterAuditLogs[i]) { // not required
			continue
		}

		if m.ClusterAuditLogs[i] != nil {
			if err := m.ClusterAuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterAuditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterAuditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this audit logs tile based on the context it is used
func (m *AuditLogsTile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLogsTile) contextValidateClusterAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterAuditLogs); i++ {

		if m.ClusterAuditLogs[i] != nil {

			if swag.IsZero(m.ClusterAuditLogs[i]) { // not required
				return nil
			}

			if err := m.ClusterAuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterAuditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterAuditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLogsTile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLogsTile) UnmarshalBinary(b []byte) error {
	var res AuditLogsTile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
