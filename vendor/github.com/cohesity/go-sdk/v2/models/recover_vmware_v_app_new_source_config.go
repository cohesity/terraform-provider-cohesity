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

// RecoverVmwareVAppNewSourceConfig Recover VMware vApps New Source Config.
//
// Specifies the new destination Source configuration where the vApps will be recovered.
//
// swagger:model RecoverVmwareVAppNewSourceConfig
type RecoverVmwareVAppNewSourceConfig struct {

	// Specifies the type of VMware source to which the VMs are being restored.
	// Required: true
	// Enum: ["kvCloudDirector"]
	SourceType *string `json:"sourceType"`

	// v cloud director params
	VCloudDirectorParams *RecoverVmwareVAppVCDSourceConfig `json:"vCloudDirectorParams,omitempty"`
}

// Validate validates this recover vmware v app new source config
func (m *RecoverVmwareVAppNewSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVCloudDirectorParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recoverVmwareVAppNewSourceConfigTypeSourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kvCloudDirector"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverVmwareVAppNewSourceConfigTypeSourceTypePropEnum = append(recoverVmwareVAppNewSourceConfigTypeSourceTypePropEnum, v)
	}
}

const (

	// RecoverVmwareVAppNewSourceConfigSourceTypeKvCloudDirector captures enum value "kvCloudDirector"
	RecoverVmwareVAppNewSourceConfigSourceTypeKvCloudDirector string = "kvCloudDirector"
)

// prop value enum
func (m *RecoverVmwareVAppNewSourceConfig) validateSourceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverVmwareVAppNewSourceConfigTypeSourceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverVmwareVAppNewSourceConfig) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("sourceType", "body", m.SourceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateSourceTypeEnum("sourceType", "body", *m.SourceType); err != nil {
		return err
	}

	return nil
}

func (m *RecoverVmwareVAppNewSourceConfig) validateVCloudDirectorParams(formats strfmt.Registry) error {
	if swag.IsZero(m.VCloudDirectorParams) { // not required
		return nil
	}

	if m.VCloudDirectorParams != nil {
		if err := m.VCloudDirectorParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vCloudDirectorParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vCloudDirectorParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover vmware v app new source config based on the context it is used
func (m *RecoverVmwareVAppNewSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVCloudDirectorParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVmwareVAppNewSourceConfig) contextValidateVCloudDirectorParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VCloudDirectorParams != nil {

		if swag.IsZero(m.VCloudDirectorParams) { // not required
			return nil
		}

		if err := m.VCloudDirectorParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vCloudDirectorParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vCloudDirectorParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverVmwareVAppNewSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverVmwareVAppNewSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverVmwareVAppNewSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
