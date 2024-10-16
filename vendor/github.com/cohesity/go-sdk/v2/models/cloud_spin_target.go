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

// CloudSpinTarget Specifies the details about Cloud Spin target where backup snapshots may be converted and stored.
//
// swagger:model CloudSpinTarget
type CloudSpinTarget struct {

	// Specifies the unique id of the cloud spin entity.
	ID *int64 `json:"id,omitempty"`

	// Contains information needed to identify various resources when deploying VMs to Cloud.
	//
	// Specifies various resources when converting and deploying a VM to AWS.
	AwsParams *AwsCloudSpinParams `json:"awsParams,omitempty"`

	// Specifies various resources when converting and deploying a VM to Azure.
	AzureParams *AzureCloudSpinParams `json:"azureParams,omitempty"`

	// Specifies the name of the already added cloud spin target.
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this cloud spin target
func (m *CloudSpinTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsParams(formats); err != nil {
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

func (m *CloudSpinTarget) validateAwsParams(formats strfmt.Registry) error {
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

func (m *CloudSpinTarget) validateAzureParams(formats strfmt.Registry) error {
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

// ContextValidate validate this cloud spin target based on the context it is used
func (m *CloudSpinTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureParams(ctx, formats); err != nil {
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

func (m *CloudSpinTarget) contextValidateAwsParams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CloudSpinTarget) contextValidateAzureParams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CloudSpinTarget) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudSpinTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudSpinTarget) UnmarshalBinary(b []byte) error {
	var res CloudSpinTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
