// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KmsConfiguration  Key management system(KMS) configurations.
//
// swagger:model KmsConfiguration
type KmsConfiguration struct {
	KmsConfigurationResponseParams

	// Id of KMS.
	// Read Only: true
	ID *int64 `json:"id,omitempty"`

	// Specifies the state of KMS. 'Active' indicates that KMS is reachable from cluster. 'InActive' indicates that KMS is not reachable from cluster. 'MarkedForRemoval' indicates that KMS is marked for removal and the removal process is in progress.
	// Read Only: true
	// Enum: ["Active","InActive","MarkedForRemoval"]
	State *string `json:"state,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *KmsConfiguration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 KmsConfigurationResponseParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.KmsConfigurationResponseParams = aO0

	// AO1
	var dataAO1 struct {
		ID *int64 `json:"id,omitempty"`

		State *string `json:"state,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.State = dataAO1.State

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m KmsConfiguration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.KmsConfigurationResponseParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ID *int64 `json:"id,omitempty"`

		State *string `json:"state,omitempty"`
	}

	dataAO1.ID = m.ID

	dataAO1.State = m.State

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this kms configuration
func (m *KmsConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with KmsConfigurationResponseParams
	if err := m.KmsConfigurationResponseParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kmsConfigurationTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","InActive","MarkedForRemoval"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kmsConfigurationTypeStatePropEnum = append(kmsConfigurationTypeStatePropEnum, v)
	}
}

// property enum
func (m *KmsConfiguration) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kmsConfigurationTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KmsConfiguration) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this kms configuration based on the context it is used
func (m *KmsConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with KmsConfigurationResponseParams
	if err := m.KmsConfigurationResponseParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KmsConfiguration) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *KmsConfiguration) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KmsConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KmsConfiguration) UnmarshalBinary(b []byte) error {
	var res KmsConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
