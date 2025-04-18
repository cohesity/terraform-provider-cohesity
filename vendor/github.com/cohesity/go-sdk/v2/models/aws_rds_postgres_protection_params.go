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
	"github.com/go-openapi/validate"
)

// AwsRdsPostgresProtectionParams AWS RDS Snapshot Manager Object protection Request Params.
//
// Specifies the parameters which are specific to AWS RDS Postgres related Object protection.
//
// swagger:model AwsRdsPostgresProtectionParams
type AwsRdsPostgresProtectionParams struct {

	// Specifies the objects to be protected.
	// Unique: true
	Objects []*AwsObjectLevelParams `json:"objects"`

	// Specifies the id of the source of the objects.
	// Read Only: true
	SourceID *int64 `json:"sourceId,omitempty"`
}

// Validate validates this aws rds postgres protection params
func (m *AwsRdsPostgresProtectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRdsPostgresProtectionParams) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	if err := validate.UniqueItems("objects", "body", m.Objects); err != nil {
		return err
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this aws rds postgres protection params based on the context it is used
func (m *AwsRdsPostgresProtectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRdsPostgresProtectionParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsRdsPostgresProtectionParams) contextValidateSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsRdsPostgresProtectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsRdsPostgresProtectionParams) UnmarshalBinary(b []byte) error {
	var res AwsRdsPostgresProtectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
