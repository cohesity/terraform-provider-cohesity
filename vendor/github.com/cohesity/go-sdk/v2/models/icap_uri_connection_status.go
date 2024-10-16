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

// IcapURIConnectionStatus Specifies the ICAP Uri connection status.
//
// swagger:model IcapUriConnectionStatus
type IcapURIConnectionStatus struct {

	// Specifies the ICAP Uri.
	IcapURI *string `json:"icapUri,omitempty"`

	// Specifies the connection status.
	// Enum: ["Succeeded","Failed"]
	ConnectionStatus *string `json:"connectionStatus,omitempty"`
}

// Validate validates this icap Uri connection status
func (m *IcapURIConnectionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var icapUriConnectionStatusTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		icapUriConnectionStatusTypeConnectionStatusPropEnum = append(icapUriConnectionStatusTypeConnectionStatusPropEnum, v)
	}
}

const (

	// IcapURIConnectionStatusConnectionStatusSucceeded captures enum value "Succeeded"
	IcapURIConnectionStatusConnectionStatusSucceeded string = "Succeeded"

	// IcapURIConnectionStatusConnectionStatusFailed captures enum value "Failed"
	IcapURIConnectionStatusConnectionStatusFailed string = "Failed"
)

// prop value enum
func (m *IcapURIConnectionStatus) validateConnectionStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, icapUriConnectionStatusTypeConnectionStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IcapURIConnectionStatus) validateConnectionStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connectionStatus", "body", *m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this icap Uri connection status based on context it is used
func (m *IcapURIConnectionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IcapURIConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IcapURIConnectionStatus) UnmarshalBinary(b []byte) error {
	var res IcapURIConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
