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

// SapHanaProtectionGroupParams Specifies parameters related to the SAP HANA Protection job.
//
// swagger:model SapHanaProtectionGroupParams
type SapHanaProtectionGroupParams struct {

	// Specifies the source Id of the objects to be protected.
	// Required: true
	SourceID *int64 `json:"sourceId"`

	// Specifies a list of fully qualified names of the objects to be protected.
	// Required: true
	// Min Items: 1
	Objects []*UdaProtectionGroupObjectParams `json:"objects"`

	// Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. If not specified, the default value is taken as 1.
	Concurrency *int32 `json:"concurrency,omitempty"`

	// Specifies the objects to be excluded in the Protection Group.
	// Unique: true
	ExcludeObjectIds []int64 `json:"excludeObjectIds"`

	// Specifies the incremental backup delta (incremental/differential)
	// Required: true
	Delta *string `json:"delta"`
}

// Validate validates this sap hana protection group params
func (m *SapHanaProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludeObjectIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SapHanaProtectionGroupParams) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *SapHanaProtectionGroupParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
		return err
	}

	iObjectsSize := int64(len(m.Objects))

	if err := validate.MinItems("objects", "body", iObjectsSize, 1); err != nil {
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

func (m *SapHanaProtectionGroupParams) validateExcludeObjectIds(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeObjectIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("excludeObjectIds", "body", m.ExcludeObjectIds); err != nil {
		return err
	}

	return nil
}

func (m *SapHanaProtectionGroupParams) validateDelta(formats strfmt.Registry) error {

	if err := validate.Required("delta", "body", m.Delta); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sap hana protection group params based on the context it is used
func (m *SapHanaProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SapHanaProtectionGroupParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SapHanaProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SapHanaProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res SapHanaProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
