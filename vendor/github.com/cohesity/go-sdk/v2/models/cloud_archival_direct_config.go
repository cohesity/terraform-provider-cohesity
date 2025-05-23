// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudArchivalDirectConfig Cloud Archival Direct Config
//
// Specifies the properties of vaults used to perform Cloud Archive Direct (CAD)
//
// swagger:model CloudArchivalDirectConfig
type CloudArchivalDirectConfig struct {

	// Specifies a namespace under the bucket used for archival.
	BucketNamespace *string `json:"bucketNamespace,omitempty"`

	// Specifies quota limit in bytes for physical usage associated with this vault during the first CAD job.
	PhysicalQuota *QuotaPolicy `json:"physicalQuota,omitempty"`
}

// Validate validates this cloud archival direct config
func (m *CloudArchivalDirectConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhysicalQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudArchivalDirectConfig) validatePhysicalQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.PhysicalQuota) { // not required
		return nil
	}

	if m.PhysicalQuota != nil {
		if err := m.PhysicalQuota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("physicalQuota")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud archival direct config based on the context it is used
func (m *CloudArchivalDirectConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhysicalQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudArchivalDirectConfig) contextValidatePhysicalQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.PhysicalQuota != nil {

		if swag.IsZero(m.PhysicalQuota) { // not required
			return nil
		}

		if err := m.PhysicalQuota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("physicalQuota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudArchivalDirectConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudArchivalDirectConfig) UnmarshalBinary(b []byte) error {
	var res CloudArchivalDirectConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
