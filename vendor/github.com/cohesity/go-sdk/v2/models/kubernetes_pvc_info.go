// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubernetesPvcInfo Kubernetes PVC Info
//
// Specifies the parameters which are related to Kubernetes PVC.
//
// swagger:model KubernetesPvcInfo
type KubernetesPvcInfo struct {

	// Specifies the id of the pvc.
	// Read Only: true
	ID *int64 `json:"id,omitempty"`

	// Name of the pvc.
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this kubernetes pvc info
func (m *KubernetesPvcInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this kubernetes pvc info based on the context it is used
func (m *KubernetesPvcInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesPvcInfo) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesPvcInfo) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesPvcInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesPvcInfo) UnmarshalBinary(b []byte) error {
	var res KubernetesPvcInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
