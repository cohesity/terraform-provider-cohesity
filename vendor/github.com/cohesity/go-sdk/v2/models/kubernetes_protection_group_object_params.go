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

// KubernetesProtectionGroupObjectParams Kubernetes ProtectionGroup Object Params
//
// Specifies the object parameters to create Kubernetes Protection Group.
//
// swagger:model KubernetesProtectionGroupObjectParams
type KubernetesProtectionGroupObjectParams struct {

	// Specifies the id of the object.
	// Required: true
	ID *int64 `json:"id"`

	// Specifies the name of the object.
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Specifies a list of pvcs to exclude from being protected. This is only applicable to kubernetes.
	ExcludePvcs []*KubernetesPvcInfo `json:"excludePvcs"`

	// Specifies a list of Pvcs to include in the protection. This is only applicable to kubernetes.
	IncludePvcs []*KubernetesPvcInfo `json:"includePvcs"`

	// Specifies the quiescing rules are which specified by the user for doing backup.
	QuiesceGroups []*QuiesceGroup `json:"quiesceGroups"`

	// Specifies the resources to include during backup
	IncludedResources []string `json:"includedResources"`

	// Specifies the resources to exclude during backup
	ExcludedResources []string `json:"excludedResources"`

	// Specifies whether to backup pvc and related resources only
	BackupOnlyPvc *bool `json:"backupOnlyPvc,omitempty"`
}

// Validate validates this kubernetes protection group object params
func (m *KubernetesProtectionGroupObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludePvcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludePvcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuiesceGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesProtectionGroupObjectParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesProtectionGroupObjectParams) validateExcludePvcs(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludePvcs) { // not required
		return nil
	}

	for i := 0; i < len(m.ExcludePvcs); i++ {
		if swag.IsZero(m.ExcludePvcs[i]) { // not required
			continue
		}

		if m.ExcludePvcs[i] != nil {
			if err := m.ExcludePvcs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludePvcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludePvcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesProtectionGroupObjectParams) validateIncludePvcs(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludePvcs) { // not required
		return nil
	}

	for i := 0; i < len(m.IncludePvcs); i++ {
		if swag.IsZero(m.IncludePvcs[i]) { // not required
			continue
		}

		if m.IncludePvcs[i] != nil {
			if err := m.IncludePvcs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includePvcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("includePvcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesProtectionGroupObjectParams) validateQuiesceGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.QuiesceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.QuiesceGroups); i++ {
		if swag.IsZero(m.QuiesceGroups[i]) { // not required
			continue
		}

		if m.QuiesceGroups[i] != nil {
			if err := m.QuiesceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quiesceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quiesceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this kubernetes protection group object params based on the context it is used
func (m *KubernetesProtectionGroupObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcludePvcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncludePvcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuiesceGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesProtectionGroupObjectParams) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesProtectionGroupObjectParams) contextValidateExcludePvcs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExcludePvcs); i++ {

		if m.ExcludePvcs[i] != nil {

			if swag.IsZero(m.ExcludePvcs[i]) { // not required
				return nil
			}

			if err := m.ExcludePvcs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludePvcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludePvcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesProtectionGroupObjectParams) contextValidateIncludePvcs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncludePvcs); i++ {

		if m.IncludePvcs[i] != nil {

			if swag.IsZero(m.IncludePvcs[i]) { // not required
				return nil
			}

			if err := m.IncludePvcs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includePvcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("includePvcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesProtectionGroupObjectParams) contextValidateQuiesceGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QuiesceGroups); i++ {

		if m.QuiesceGroups[i] != nil {

			if swag.IsZero(m.QuiesceGroups[i]) { // not required
				return nil
			}

			if err := m.QuiesceGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quiesceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quiesceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesProtectionGroupObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesProtectionGroupObjectParams) UnmarshalBinary(b []byte) error {
	var res KubernetesProtectionGroupObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
