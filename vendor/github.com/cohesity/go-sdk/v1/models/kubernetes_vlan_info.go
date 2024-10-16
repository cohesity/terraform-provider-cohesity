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

// KubernetesVlanInfo Vlan information for Kubernetes protection source.
//
// Specifies a Protection Source in Kubernetes environment.
//
// swagger:model KubernetesVlanInfo
type KubernetesVlanInfo struct {

	// Specifies annotations to be put on services for IP allocation. Applicable
	// only when service is of type LoadBalancer.
	ServiceAnnotations []*VlanInfoServiceAnnotationsEntry `json:"serviceAnnotations"`

	// Specifies selected VLAN parameters to be used
	VlanParams *VlanParameters `json:"vlanParams,omitempty"`
}

// Validate validates this kubernetes vlan info
func (m *KubernetesVlanInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesVlanInfo) validateServiceAnnotations(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAnnotations) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceAnnotations); i++ {
		if swag.IsZero(m.ServiceAnnotations[i]) { // not required
			continue
		}

		if m.ServiceAnnotations[i] != nil {
			if err := m.ServiceAnnotations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceAnnotations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceAnnotations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesVlanInfo) validateVlanParams(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanParams) { // not required
		return nil
	}

	if m.VlanParams != nil {
		if err := m.VlanParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kubernetes vlan info based on the context it is used
func (m *KubernetesVlanInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceAnnotations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesVlanInfo) contextValidateServiceAnnotations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceAnnotations); i++ {

		if m.ServiceAnnotations[i] != nil {

			if swag.IsZero(m.ServiceAnnotations[i]) { // not required
				return nil
			}

			if err := m.ServiceAnnotations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceAnnotations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceAnnotations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubernetesVlanInfo) contextValidateVlanParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanParams != nil {

		if swag.IsZero(m.VlanParams) { // not required
			return nil
		}

		if err := m.VlanParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesVlanInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesVlanInfo) UnmarshalBinary(b []byte) error {
	var res KubernetesVlanInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
