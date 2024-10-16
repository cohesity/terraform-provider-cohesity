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

// GetRestoreAppTimeRangesArg API to get the available restore time ranges for an application restore
// operation. In the example of SQL/Oracle restore, Iris will use this RPC
// to get the time ranges to which point-in-time database restore can be
// performed.
//
// swagger:model GetRestoreAppTimeRangesArg
type GetRestoreAppTimeRangesArg struct {

	// Specifies the request attributes.
	APIRequestAttr *APIRequestAttr `json:"apiRequestAttr,omitempty"`

	// Specifies the API version used by this arg.
	APIVersion *APIVersion `json:"apiVersion,omitempty"`

	// The set of owner restore objects for which application restore time
	// ranges are desired. The owner restore object represents the full snapshot
	// of the owner object (such as a VM) from which an application (such as SQL)
	// needs to be restored to a point in time. Note that all the owner objects
	// should represent the same entity (for example, the same VM).
	OwnerObjectVec []*RestoreObject `json:"ownerObjectVec"`

	// The application level objects that needs to be restored. If a single
	// object without its 'app_entity' is specified, the restore time ranges will
	// apply to all application objects in the owner restore object. Otherwise
	// the restore time ranges will only apply to the entities of these objects.
	// If multiple objects are specified, the 'app_entity' field must be
	// specified for all of them.
	RestoreAppObjectVec []*RestoreAppObject `json:"restoreAppObjectVec"`

	// The application environment.
	Type *int32 `json:"type,omitempty"`
}

// Validate validates this get restore app time ranges arg
func (m *GetRestoreAppTimeRangesArg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIRequestAttr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerObjectVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreAppObjectVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetRestoreAppTimeRangesArg) validateAPIRequestAttr(formats strfmt.Registry) error {
	if swag.IsZero(m.APIRequestAttr) { // not required
		return nil
	}

	if m.APIRequestAttr != nil {
		if err := m.APIRequestAttr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiRequestAttr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiRequestAttr")
			}
			return err
		}
	}

	return nil
}

func (m *GetRestoreAppTimeRangesArg) validateAPIVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.APIVersion) { // not required
		return nil
	}

	if m.APIVersion != nil {
		if err := m.APIVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiVersion")
			}
			return err
		}
	}

	return nil
}

func (m *GetRestoreAppTimeRangesArg) validateOwnerObjectVec(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerObjectVec) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnerObjectVec); i++ {
		if swag.IsZero(m.OwnerObjectVec[i]) { // not required
			continue
		}

		if m.OwnerObjectVec[i] != nil {
			if err := m.OwnerObjectVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ownerObjectVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ownerObjectVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetRestoreAppTimeRangesArg) validateRestoreAppObjectVec(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreAppObjectVec) { // not required
		return nil
	}

	for i := 0; i < len(m.RestoreAppObjectVec); i++ {
		if swag.IsZero(m.RestoreAppObjectVec[i]) { // not required
			continue
		}

		if m.RestoreAppObjectVec[i] != nil {
			if err := m.RestoreAppObjectVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoreAppObjectVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoreAppObjectVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get restore app time ranges arg based on the context it is used
func (m *GetRestoreAppTimeRangesArg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIRequestAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwnerObjectVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoreAppObjectVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetRestoreAppTimeRangesArg) contextValidateAPIRequestAttr(ctx context.Context, formats strfmt.Registry) error {

	if m.APIRequestAttr != nil {

		if swag.IsZero(m.APIRequestAttr) { // not required
			return nil
		}

		if err := m.APIRequestAttr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiRequestAttr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiRequestAttr")
			}
			return err
		}
	}

	return nil
}

func (m *GetRestoreAppTimeRangesArg) contextValidateAPIVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.APIVersion != nil {

		if swag.IsZero(m.APIVersion) { // not required
			return nil
		}

		if err := m.APIVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiVersion")
			}
			return err
		}
	}

	return nil
}

func (m *GetRestoreAppTimeRangesArg) contextValidateOwnerObjectVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OwnerObjectVec); i++ {

		if m.OwnerObjectVec[i] != nil {

			if swag.IsZero(m.OwnerObjectVec[i]) { // not required
				return nil
			}

			if err := m.OwnerObjectVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ownerObjectVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ownerObjectVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetRestoreAppTimeRangesArg) contextValidateRestoreAppObjectVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RestoreAppObjectVec); i++ {

		if m.RestoreAppObjectVec[i] != nil {

			if swag.IsZero(m.RestoreAppObjectVec[i]) { // not required
				return nil
			}

			if err := m.RestoreAppObjectVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoreAppObjectVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoreAppObjectVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetRestoreAppTimeRangesArg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRestoreAppTimeRangesArg) UnmarshalBinary(b []byte) error {
	var res GetRestoreAppTimeRangesArg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
