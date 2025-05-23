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

// UpgradeCheckRunTestsRequest Upgrade checks request
//
// # Specifies upgrade checks request parameters
//
// swagger:model UpgradeCheckRunTestsRequest
type UpgradeCheckRunTestsRequest struct {

	// Type of upgrade checks(pre/post) to run
	// Enum: ["PreUpgrade","PostUpgrade"]
	RequestType string `json:"requestType,omitempty"`
}

// Validate validates this upgrade check run tests request
func (m *UpgradeCheckRunTestsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var upgradeCheckRunTestsRequestTypeRequestTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PreUpgrade","PostUpgrade"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		upgradeCheckRunTestsRequestTypeRequestTypePropEnum = append(upgradeCheckRunTestsRequestTypeRequestTypePropEnum, v)
	}
}

const (

	// UpgradeCheckRunTestsRequestRequestTypePreUpgrade captures enum value "PreUpgrade"
	UpgradeCheckRunTestsRequestRequestTypePreUpgrade string = "PreUpgrade"

	// UpgradeCheckRunTestsRequestRequestTypePostUpgrade captures enum value "PostUpgrade"
	UpgradeCheckRunTestsRequestRequestTypePostUpgrade string = "PostUpgrade"
)

// prop value enum
func (m *UpgradeCheckRunTestsRequest) validateRequestTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, upgradeCheckRunTestsRequestTypeRequestTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpgradeCheckRunTestsRequest) validateRequestType(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRequestTypeEnum("requestType", "body", m.RequestType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this upgrade check run tests request based on context it is used
func (m *UpgradeCheckRunTestsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeCheckRunTestsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeCheckRunTestsRequest) UnmarshalBinary(b []byte) error {
	var res UpgradeCheckRunTestsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
