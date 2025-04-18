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

// ViewProtocol Protocol Option
//
// Specifies the protocol options for view.
//
// swagger:model ViewProtocol
type ViewProtocol struct {

	// Type of protocol.
	// Specifies the supported Protocols for the View.
	//   'NFS' enables protocol access to NFS v3.
	//   'NFS4' enables protocol access to NFS v4.1.
	//   'SMB' enables protocol access to SMB.
	//   'S3' enables protocol access to S3.
	//   'Swift' enables protocol access to Swift.
	// Enum: ["NFS","NFS4","SMB","S3","Swift"]
	Type string `json:"type,omitempty"`

	// Mode of protocol access.
	//   'ReadOnly'
	//   'ReadWrite'
	// Enum: ["ReadOnly","ReadWrite"]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this view protocol
func (m *ViewProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var viewProtocolTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NFS","NFS4","SMB","S3","Swift"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewProtocolTypeTypePropEnum = append(viewProtocolTypeTypePropEnum, v)
	}
}

const (

	// ViewProtocolTypeNFS captures enum value "NFS"
	ViewProtocolTypeNFS string = "NFS"

	// ViewProtocolTypeNFS4 captures enum value "NFS4"
	ViewProtocolTypeNFS4 string = "NFS4"

	// ViewProtocolTypeSMB captures enum value "SMB"
	ViewProtocolTypeSMB string = "SMB"

	// ViewProtocolTypeS3 captures enum value "S3"
	ViewProtocolTypeS3 string = "S3"

	// ViewProtocolTypeSwift captures enum value "Swift"
	ViewProtocolTypeSwift string = "Swift"
)

// prop value enum
func (m *ViewProtocol) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewProtocolTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewProtocol) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var viewProtocolTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ReadOnly","ReadWrite"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewProtocolTypeModePropEnum = append(viewProtocolTypeModePropEnum, v)
	}
}

const (

	// ViewProtocolModeReadOnly captures enum value "ReadOnly"
	ViewProtocolModeReadOnly string = "ReadOnly"

	// ViewProtocolModeReadWrite captures enum value "ReadWrite"
	ViewProtocolModeReadWrite string = "ReadWrite"
)

// prop value enum
func (m *ViewProtocol) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewProtocolTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ViewProtocol) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this view protocol based on context it is used
func (m *ViewProtocol) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViewProtocol) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewProtocol) UnmarshalBinary(b []byte) error {
	var res ViewProtocol
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
