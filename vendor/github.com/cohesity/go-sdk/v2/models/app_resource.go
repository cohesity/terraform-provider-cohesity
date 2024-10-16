// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppResource Specifies the details about App Resource.
//
// swagger:model AppResource
type AppResource struct {

	// Specifies the unique id of the app resource.
	ID *string `json:"id,omitempty"`

	// Specifies the type of the app resource.
	// Enum: ["standaloneServer","windowsCluster","sqlCluster"]
	Type *string `json:"type,omitempty"`

	// Specifies the information about endpoint associated with this resorce.
	Endpoints []*ResourceEndpoint `json:"endpoints"`
}

// Validate validates this app resource
func (m *AppResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var appResourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["standaloneServer","windowsCluster","sqlCluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appResourceTypeTypePropEnum = append(appResourceTypeTypePropEnum, v)
	}
}

const (

	// AppResourceTypeStandaloneServer captures enum value "standaloneServer"
	AppResourceTypeStandaloneServer string = "standaloneServer"

	// AppResourceTypeWindowsCluster captures enum value "windowsCluster"
	AppResourceTypeWindowsCluster string = "windowsCluster"

	// AppResourceTypeSQLCluster captures enum value "sqlCluster"
	AppResourceTypeSQLCluster string = "sqlCluster"
)

// prop value enum
func (m *AppResource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, appResourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AppResource) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AppResource) validateEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.Endpoints); i++ {
		if swag.IsZero(m.Endpoints[i]) { // not required
			continue
		}

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this app resource based on the context it is used
func (m *AppResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppResource) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Endpoints); i++ {

		if m.Endpoints[i] != nil {

			if swag.IsZero(m.Endpoints[i]) { // not required
				return nil
			}

			if err := m.Endpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppResource) UnmarshalBinary(b []byte) error {
	var res AppResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
