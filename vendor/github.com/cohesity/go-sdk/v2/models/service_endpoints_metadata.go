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

// ServiceEndpointsMetadata ServiceEndpointsMetadata
//
// Specifies the service endpoints that can be configured based on the environnment.
//
// swagger:model ServiceEndpointsMetadata
type ServiceEndpointsMetadata struct {

	// Specifies the name of the IBM cloud service for which the endpoints need to be configured. Based on the provided name here, API callers must set the appropriate service metadata.
	// Required: true
	// Enum: ["VPC","IAM"]
	CloudServiceName *string `json:"cloudServiceName"`

	// Specifies the service configuration for VPC service in IBM cloud.
	VpcParams *IbmVPCServiceMetadata `json:"vpcParams,omitempty"`

	// Specifies the service configuration for IAM service in IBM cloud.
	IamParams *IbmIAMCServiceMetadata `json:"iamParams,omitempty"`
}

// Validate validates this service endpoints metadata
func (m *ServiceEndpointsMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudServiceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIamParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceEndpointsMetadataTypeCloudServiceNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VPC","IAM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceEndpointsMetadataTypeCloudServiceNamePropEnum = append(serviceEndpointsMetadataTypeCloudServiceNamePropEnum, v)
	}
}

const (

	// ServiceEndpointsMetadataCloudServiceNameVPC captures enum value "VPC"
	ServiceEndpointsMetadataCloudServiceNameVPC string = "VPC"

	// ServiceEndpointsMetadataCloudServiceNameIAM captures enum value "IAM"
	ServiceEndpointsMetadataCloudServiceNameIAM string = "IAM"
)

// prop value enum
func (m *ServiceEndpointsMetadata) validateCloudServiceNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceEndpointsMetadataTypeCloudServiceNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceEndpointsMetadata) validateCloudServiceName(formats strfmt.Registry) error {

	if err := validate.Required("cloudServiceName", "body", m.CloudServiceName); err != nil {
		return err
	}

	// value enum
	if err := m.validateCloudServiceNameEnum("cloudServiceName", "body", *m.CloudServiceName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceEndpointsMetadata) validateVpcParams(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcParams) { // not required
		return nil
	}

	if m.VpcParams != nil {
		if err := m.VpcParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpcParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpcParams")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEndpointsMetadata) validateIamParams(formats strfmt.Registry) error {
	if swag.IsZero(m.IamParams) { // not required
		return nil
	}

	if m.IamParams != nil {
		if err := m.IamParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iamParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iamParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service endpoints metadata based on the context it is used
func (m *ServiceEndpointsMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVpcParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIamParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEndpointsMetadata) contextValidateVpcParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcParams != nil {

		if swag.IsZero(m.VpcParams) { // not required
			return nil
		}

		if err := m.VpcParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpcParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpcParams")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEndpointsMetadata) contextValidateIamParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IamParams != nil {

		if swag.IsZero(m.IamParams) { // not required
			return nil
		}

		if err := m.IamParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iamParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iamParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceEndpointsMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceEndpointsMetadata) UnmarshalBinary(b []byte) error {
	var res ServiceEndpointsMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
