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

// LicenseState license state
//
// swagger:model LicenseState
type LicenseState struct {

	// Specifies no of failed attempts at claiming the license server
	FailedAttempts *int64 `json:"failedAttempts,omitempty"`

	// Specifies the current state of licensing workflow.
	// Enum: ["kInProgressNewCluster","kInProgressOldCluster","kClaimed","kSkipped","kStarted"]
	State *string `json:"state,omitempty"`
}

// Validate validates this license state
func (m *LicenseState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var licenseStateTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kInProgressNewCluster","kInProgressOldCluster","kClaimed","kSkipped","kStarted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		licenseStateTypeStatePropEnum = append(licenseStateTypeStatePropEnum, v)
	}
}

const (

	// LicenseStateStateKInProgressNewCluster captures enum value "kInProgressNewCluster"
	LicenseStateStateKInProgressNewCluster string = "kInProgressNewCluster"

	// LicenseStateStateKInProgressOldCluster captures enum value "kInProgressOldCluster"
	LicenseStateStateKInProgressOldCluster string = "kInProgressOldCluster"

	// LicenseStateStateKClaimed captures enum value "kClaimed"
	LicenseStateStateKClaimed string = "kClaimed"

	// LicenseStateStateKSkipped captures enum value "kSkipped"
	LicenseStateStateKSkipped string = "kSkipped"

	// LicenseStateStateKStarted captures enum value "kStarted"
	LicenseStateStateKStarted string = "kStarted"
)

// prop value enum
func (m *LicenseState) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, licenseStateTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LicenseState) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this license state based on context it is used
func (m *LicenseState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicenseState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseState) UnmarshalBinary(b []byte) error {
	var res LicenseState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
