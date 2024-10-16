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

// AdObjectRestoreInformation AD Object Restore Information.
//
// Represents the details about the restore of the AD object.
//
// swagger:model AdObjectRestoreInformation
type AdObjectRestoreInformation struct {

	// Specifies the list of attributes of the AD object whose restore failed.
	AttributeRestoreInfo []*AttributeRestoreInformation `json:"attributeRestoreInfo"`

	// Specifies the error message while restoring the AD Object.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Specifies the name of the AD object.
	Name *string `json:"name,omitempty"`

	// Specifies the start time of the restore of the AD object specified as a
	// Unix epoch Timestamp(in microseconds).
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`

	// Specifies the time taken for restore of AD Object and its attributes in
	// milliseconds.
	TimeTakenMsecs *uint64 `json:"timeTakenMsecs,omitempty"`
}

// Validate validates this ad object restore information
func (m *AdObjectRestoreInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeRestoreInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdObjectRestoreInformation) validateAttributeRestoreInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeRestoreInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeRestoreInfo); i++ {
		if swag.IsZero(m.AttributeRestoreInfo[i]) { // not required
			continue
		}

		if m.AttributeRestoreInfo[i] != nil {
			if err := m.AttributeRestoreInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeRestoreInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeRestoreInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ad object restore information based on the context it is used
func (m *AdObjectRestoreInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributeRestoreInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdObjectRestoreInformation) contextValidateAttributeRestoreInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttributeRestoreInfo); i++ {

		if m.AttributeRestoreInfo[i] != nil {

			if swag.IsZero(m.AttributeRestoreInfo[i]) { // not required
				return nil
			}

			if err := m.AttributeRestoreInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeRestoreInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeRestoreInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdObjectRestoreInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdObjectRestoreInformation) UnmarshalBinary(b []byte) error {
	var res AdObjectRestoreInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
