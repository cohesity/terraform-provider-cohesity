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

// EntityMetadataParams Specifies the parameters to associate metadata with entities in the entity hierarchy.
//
// swagger:model EntityMetadataParams
type EntityMetadataParams struct {

	// Specifies the entity id of the entity whose metadata is being updated.
	// Required: true
	EntityID *int64 `json:"entityId"`

	// Sepecifies the parameters if the entity is a aws type entity.
	AwsParams *AwsEntityMetadata `json:"awsParams,omitempty"`

	// Specifies the parameters to be filled for maintenance mode.
	MaintenanceModeConfig *MaintenanceModeConfig `json:"maintenanceModeConfig,omitempty"`

	// Specifies the parameters if the entity is a azure entity.
	AzureParams *AzureEntityMetadata `json:"azureParams,omitempty"`
}

// Validate validates this entity metadata params
func (m *EntityMetadataParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceModeConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityMetadataParams) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entityId", "body", m.EntityID); err != nil {
		return err
	}

	return nil
}

func (m *EntityMetadataParams) validateAwsParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsParams) { // not required
		return nil
	}

	if m.AwsParams != nil {
		if err := m.AwsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsParams")
			}
			return err
		}
	}

	return nil
}

func (m *EntityMetadataParams) validateMaintenanceModeConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MaintenanceModeConfig) { // not required
		return nil
	}

	if m.MaintenanceModeConfig != nil {
		if err := m.MaintenanceModeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenanceModeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenanceModeConfig")
			}
			return err
		}
	}

	return nil
}

func (m *EntityMetadataParams) validateAzureParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureParams) { // not required
		return nil
	}

	if m.AzureParams != nil {
		if err := m.AzureParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this entity metadata params based on the context it is used
func (m *EntityMetadataParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaintenanceModeConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityMetadataParams) contextValidateAwsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsParams != nil {

		if swag.IsZero(m.AwsParams) { // not required
			return nil
		}

		if err := m.AwsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsParams")
			}
			return err
		}
	}

	return nil
}

func (m *EntityMetadataParams) contextValidateMaintenanceModeConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.MaintenanceModeConfig != nil {

		if swag.IsZero(m.MaintenanceModeConfig) { // not required
			return nil
		}

		if err := m.MaintenanceModeConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenanceModeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maintenanceModeConfig")
			}
			return err
		}
	}

	return nil
}

func (m *EntityMetadataParams) contextValidateAzureParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureParams != nil {

		if swag.IsZero(m.AzureParams) { // not required
			return nil
		}

		if err := m.AzureParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityMetadataParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityMetadataParams) UnmarshalBinary(b []byte) error {
	var res EntityMetadataParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
