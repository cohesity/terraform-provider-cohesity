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

// AwsRdsPostgresProtectionGroupParams AWS RDS Snapshot Manager Protection Group Request Params.
//
// Specifies the parameters which are specific to AWS RDS Postgres related Protection Groups.
//
// swagger:model AwsRdsPostgresProtectionGroupParams
type AwsRdsPostgresProtectionGroupParams struct {

	// Specifies the objects to be included in the Protection Group.
	// Unique: true
	Objects []*AwsRdsPostgresProtectionGroupObjectParams `json:"objects"`

	// Specifies the objects to be excluded in the Protection Group.
	// Unique: true
	ExcludeObjectIds []int64 `json:"excludeObjectIds"`

	// Specifies the id of the parent of the objects.
	// Read Only: true
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies the name of the parent of the objects.
	// Read Only: true
	SourceName *string `json:"sourceName,omitempty"`
}

// Validate validates this aws rds postgres protection group params
func (m *AwsRdsPostgresProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludeObjectIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRdsPostgresProtectionGroupParams) validateObjects(formats strfmt.Registry) error {
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

func (m *AwsRdsPostgresProtectionGroupParams) validateExcludeObjectIds(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeObjectIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("excludeObjectIds", "body", m.ExcludeObjectIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws rds postgres protection group params based on the context it is used
func (m *AwsRdsPostgresProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRdsPostgresProtectionGroupParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AwsRdsPostgresProtectionGroupParams) contextValidateSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *AwsRdsPostgresProtectionGroupParams) contextValidateSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceName", "body", m.SourceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsRdsPostgresProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsRdsPostgresProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res AwsRdsPostgresProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
