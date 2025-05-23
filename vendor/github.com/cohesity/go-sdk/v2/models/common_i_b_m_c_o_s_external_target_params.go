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

// CommonIBMCOSExternalTargetParams IBM COS external target request common parameters.
//
// Specifies the common parameters for IBM COS external targets.
//
// swagger:model CommonIBMCOSExternalTargetParams
type CommonIBMCOSExternalTargetParams struct {

	// Specifies bucket name of the external target.
	// Required: true
	BucketName *string `json:"bucketName"`

	// Specifies location of the external target. A location can be a geo(Ex. us-geo), region(Ex. us-east) or a datacenter site (Ex. sjc04) and is used to identify the external target endpoint.
	// Required: true
	Location *string `json:"location"`

	// Specifies the endpoint type to be used to access the external target.
	// Enum: ["Private","Public","Direct"]
	EndpointType *string `json:"endpointType,omitempty"`

	// authentication method
	// Required: true
	AuthenticationMethod *IBMAuthenticationMethodsParams `json:"authenticationMethod"`
}

// Validate validates this common i b m c o s external target params
func (m *CommonIBMCOSExternalTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucketName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonIBMCOSExternalTargetParams) validateBucketName(formats strfmt.Registry) error {

	if err := validate.Required("bucketName", "body", m.BucketName); err != nil {
		return err
	}

	return nil
}

func (m *CommonIBMCOSExternalTargetParams) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	return nil
}

var commonIBMCOSExternalTargetParamsTypeEndpointTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Private","Public","Direct"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonIBMCOSExternalTargetParamsTypeEndpointTypePropEnum = append(commonIBMCOSExternalTargetParamsTypeEndpointTypePropEnum, v)
	}
}

const (

	// CommonIBMCOSExternalTargetParamsEndpointTypePrivate captures enum value "Private"
	CommonIBMCOSExternalTargetParamsEndpointTypePrivate string = "Private"

	// CommonIBMCOSExternalTargetParamsEndpointTypePublic captures enum value "Public"
	CommonIBMCOSExternalTargetParamsEndpointTypePublic string = "Public"

	// CommonIBMCOSExternalTargetParamsEndpointTypeDirect captures enum value "Direct"
	CommonIBMCOSExternalTargetParamsEndpointTypeDirect string = "Direct"
)

// prop value enum
func (m *CommonIBMCOSExternalTargetParams) validateEndpointTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonIBMCOSExternalTargetParamsTypeEndpointTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonIBMCOSExternalTargetParams) validateEndpointType(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEndpointTypeEnum("endpointType", "body", *m.EndpointType); err != nil {
		return err
	}

	return nil
}

func (m *CommonIBMCOSExternalTargetParams) validateAuthenticationMethod(formats strfmt.Registry) error {

	if err := validate.Required("authenticationMethod", "body", m.AuthenticationMethod); err != nil {
		return err
	}

	if m.AuthenticationMethod != nil {
		if err := m.AuthenticationMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticationMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticationMethod")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common i b m c o s external target params based on the context it is used
func (m *CommonIBMCOSExternalTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthenticationMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonIBMCOSExternalTargetParams) contextValidateAuthenticationMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthenticationMethod != nil {

		if err := m.AuthenticationMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticationMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authenticationMethod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonIBMCOSExternalTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonIBMCOSExternalTargetParams) UnmarshalBinary(b []byte) error {
	var res CommonIBMCOSExternalTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
