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

// ExternallyTriggeredEnvJobParameters Externally Triggered Backup Environment Job Parameters.
//
// swagger:model ExternallyTriggeredEnvJobParameters
type ExternallyTriggeredEnvJobParameters struct {

	// Specifies the client type of the externally triggered job
	// kGeneric implies generic externally triggered backups.
	// kSbt implies externally triggered backups for SBT.
	// Enum: ["kGeneric","kSbt"]
	ClientType *string `json:"clientType,omitempty"`

	// Specifies the catalog view associated with this job.
	SbtParams *ExternallyTriggeredSbtParameters `json:"sbtParams,omitempty"`
}

// Validate validates this externally triggered env job parameters
func (m *ExternallyTriggeredEnvJobParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSbtParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var externallyTriggeredEnvJobParametersTypeClientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kGeneric","kSbt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externallyTriggeredEnvJobParametersTypeClientTypePropEnum = append(externallyTriggeredEnvJobParametersTypeClientTypePropEnum, v)
	}
}

const (

	// ExternallyTriggeredEnvJobParametersClientTypeKGeneric captures enum value "kGeneric"
	ExternallyTriggeredEnvJobParametersClientTypeKGeneric string = "kGeneric"

	// ExternallyTriggeredEnvJobParametersClientTypeKSbt captures enum value "kSbt"
	ExternallyTriggeredEnvJobParametersClientTypeKSbt string = "kSbt"
)

// prop value enum
func (m *ExternallyTriggeredEnvJobParameters) validateClientTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externallyTriggeredEnvJobParametersTypeClientTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternallyTriggeredEnvJobParameters) validateClientType(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientType) { // not required
		return nil
	}

	// value enum
	if err := m.validateClientTypeEnum("clientType", "body", *m.ClientType); err != nil {
		return err
	}

	return nil
}

func (m *ExternallyTriggeredEnvJobParameters) validateSbtParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SbtParams) { // not required
		return nil
	}

	if m.SbtParams != nil {
		if err := m.SbtParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbtParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbtParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this externally triggered env job parameters based on the context it is used
func (m *ExternallyTriggeredEnvJobParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSbtParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternallyTriggeredEnvJobParameters) contextValidateSbtParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SbtParams != nil {

		if swag.IsZero(m.SbtParams) { // not required
			return nil
		}

		if err := m.SbtParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbtParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbtParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternallyTriggeredEnvJobParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternallyTriggeredEnvJobParameters) UnmarshalBinary(b []byte) error {
	var res ExternallyTriggeredEnvJobParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
