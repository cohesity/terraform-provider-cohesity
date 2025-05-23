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

// GenericNasProtectionGroupExtraParams Specifies the extra parameters which are specific to NAS related Protection Groups.
//
// swagger:model GenericNasProtectionGroupExtraParams
type GenericNasProtectionGroupExtraParams struct {

	// Specifies whether or not to store the snapshots in this run directly in an Archive Target instead of on the Cluster. If this is set to true, the associated policy must have exactly one Archive Target associated with it and the policy must be set up to archive after every run. Also, a Storage Domain cannot be specified. Default behavior is 'false'.
	DirectCloudArchive *bool `json:"directCloudArchive,omitempty"`

	// Specifies whether or not to enable native format for direct archive job. This field is set to true if native format should be used for archiving.
	NativeFormat *bool `json:"nativeFormat,omitempty"`

	// Specifies the preferred protocol to use if this device supports multiple protocols.
	// Enum: ["kNoProtocol","kNfs3","kNfs4_1","kCifs1","kCifs2","kCifs3"]
	Protocol *string `json:"protocol,omitempty"`
}

// Validate validates this generic nas protection group extra params
func (m *GenericNasProtectionGroupExtraParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var genericNasProtectionGroupExtraParamsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNoProtocol","kNfs3","kNfs4_1","kCifs1","kCifs2","kCifs3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		genericNasProtectionGroupExtraParamsTypeProtocolPropEnum = append(genericNasProtectionGroupExtraParamsTypeProtocolPropEnum, v)
	}
}

const (

	// GenericNasProtectionGroupExtraParamsProtocolKNoProtocol captures enum value "kNoProtocol"
	GenericNasProtectionGroupExtraParamsProtocolKNoProtocol string = "kNoProtocol"

	// GenericNasProtectionGroupExtraParamsProtocolKNfs3 captures enum value "kNfs3"
	GenericNasProtectionGroupExtraParamsProtocolKNfs3 string = "kNfs3"

	// GenericNasProtectionGroupExtraParamsProtocolKNfs41 captures enum value "kNfs4_1"
	GenericNasProtectionGroupExtraParamsProtocolKNfs41 string = "kNfs4_1"

	// GenericNasProtectionGroupExtraParamsProtocolKCifs1 captures enum value "kCifs1"
	GenericNasProtectionGroupExtraParamsProtocolKCifs1 string = "kCifs1"

	// GenericNasProtectionGroupExtraParamsProtocolKCifs2 captures enum value "kCifs2"
	GenericNasProtectionGroupExtraParamsProtocolKCifs2 string = "kCifs2"

	// GenericNasProtectionGroupExtraParamsProtocolKCifs3 captures enum value "kCifs3"
	GenericNasProtectionGroupExtraParamsProtocolKCifs3 string = "kCifs3"
)

// prop value enum
func (m *GenericNasProtectionGroupExtraParams) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, genericNasProtectionGroupExtraParamsTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GenericNasProtectionGroupExtraParams) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this generic nas protection group extra params based on context it is used
func (m *GenericNasProtectionGroupExtraParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenericNasProtectionGroupExtraParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericNasProtectionGroupExtraParams) UnmarshalBinary(b []byte) error {
	var res GenericNasProtectionGroupExtraParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
