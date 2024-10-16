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

// IbmFlashSystemProtectionGroupParams Specifies the parameters which are specific to IBM Flash System related Protection Groups.
//
// swagger:model IbmFlashSystemProtectionGroupParams
type IbmFlashSystemProtectionGroupParams struct {

	// Specifies the objects to be included in the Protection Group.
	// Required: true
	// Min Items: 1
	// Unique: true
	Objects []*IbmFlashSystemProtectionGroupObjectParams `json:"objects"`

	// Specifies the id of the parent of the objects.
	// Read Only: true
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies the name of the parent of the objects.
	// Read Only: true
	SourceName *string `json:"sourceName,omitempty"`

	// Specifies the pre script and post script to run before and after the protection group.
	PrePostScript *HostBasedBackupScriptParams `json:"prePostScript,omitempty"`

	// Specifies whether the safeguarded copy snapshots are allowed or not
	IsSafeGuardedCopySnapshot *bool `json:"isSafeGuardedCopySnapshot,omitempty"`
}

// Validate validates this ibm flash system protection group params
func (m *IbmFlashSystemProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrePostScript(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbmFlashSystemProtectionGroupParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
		return err
	}

	iObjectsSize := int64(len(m.Objects))

	if err := validate.MinItems("objects", "body", iObjectsSize, 1); err != nil {
		return err
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

func (m *IbmFlashSystemProtectionGroupParams) validatePrePostScript(formats strfmt.Registry) error {
	if swag.IsZero(m.PrePostScript) { // not required
		return nil
	}

	if m.PrePostScript != nil {
		if err := m.PrePostScript.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prePostScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prePostScript")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibm flash system protection group params based on the context it is used
func (m *IbmFlashSystemProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidatePrePostScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbmFlashSystemProtectionGroupParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IbmFlashSystemProtectionGroupParams) contextValidateSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *IbmFlashSystemProtectionGroupParams) contextValidateSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceName", "body", m.SourceName); err != nil {
		return err
	}

	return nil
}

func (m *IbmFlashSystemProtectionGroupParams) contextValidatePrePostScript(ctx context.Context, formats strfmt.Registry) error {

	if m.PrePostScript != nil {

		if swag.IsZero(m.PrePostScript) { // not required
			return nil
		}

		if err := m.PrePostScript.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prePostScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prePostScript")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbmFlashSystemProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbmFlashSystemProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res IbmFlashSystemProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
