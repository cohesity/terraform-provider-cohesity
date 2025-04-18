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

// VmwareObjectProtectionRequestParams Specifies the parameters which are specific to VMware object protection.
//
// swagger:model VmwareObjectProtectionRequestParams
type VmwareObjectProtectionRequestParams struct {

	// Specifies the objects to include in the backup.
	// Required: true
	Objects []*VmwareObjectProtectionRequest `json:"objects"`

	// Specifies a list of disks to exclude from the backup.
	GlobalExcludeDisks []*DiskInfo `json:"globalExcludeDisks"`

	CommonVmwareProtectionParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VmwareObjectProtectionRequestParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Objects []*VmwareObjectProtectionRequest `json:"objects"`

		GlobalExcludeDisks []*DiskInfo `json:"globalExcludeDisks"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Objects = dataAO0.Objects

	m.GlobalExcludeDisks = dataAO0.GlobalExcludeDisks

	// AO1
	var aO1 CommonVmwareProtectionParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CommonVmwareProtectionParams = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VmwareObjectProtectionRequestParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Objects []*VmwareObjectProtectionRequest `json:"objects"`

		GlobalExcludeDisks []*DiskInfo `json:"globalExcludeDisks"`
	}

	dataAO0.Objects = m.Objects

	dataAO0.GlobalExcludeDisks = m.GlobalExcludeDisks

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.CommonVmwareProtectionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vmware object protection request params
func (m *VmwareObjectProtectionRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalExcludeDisks(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with CommonVmwareProtectionParams
	if err := m.CommonVmwareProtectionParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareObjectProtectionRequestParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
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

func (m *VmwareObjectProtectionRequestParams) validateGlobalExcludeDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.GlobalExcludeDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.GlobalExcludeDisks); i++ {
		if swag.IsZero(m.GlobalExcludeDisks[i]) { // not required
			continue
		}

		if m.GlobalExcludeDisks[i] != nil {
			if err := m.GlobalExcludeDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("globalExcludeDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("globalExcludeDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vmware object protection request params based on the context it is used
func (m *VmwareObjectProtectionRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalExcludeDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with CommonVmwareProtectionParams
	if err := m.CommonVmwareProtectionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareObjectProtectionRequestParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VmwareObjectProtectionRequestParams) contextValidateGlobalExcludeDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GlobalExcludeDisks); i++ {

		if m.GlobalExcludeDisks[i] != nil {

			if swag.IsZero(m.GlobalExcludeDisks[i]) { // not required
				return nil
			}

			if err := m.GlobalExcludeDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("globalExcludeDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("globalExcludeDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareObjectProtectionRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareObjectProtectionRequestParams) UnmarshalBinary(b []byte) error {
	var res VmwareObjectProtectionRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
