// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdpLogRunParamsVirtualDiskInfo cdp log run params virtual disk info
//
// swagger:model CdpLogRunParams_VirtualDiskInfo
type CdpLogRunParamsVirtualDiskInfo struct {

	// Name of the disk file.
	FileName *string `json:"fileName,omitempty"`

	// Unique disk id within the object.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this cdp log run params virtual disk info
func (m *CdpLogRunParamsVirtualDiskInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cdp log run params virtual disk info based on context it is used
func (m *CdpLogRunParamsVirtualDiskInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdpLogRunParamsVirtualDiskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdpLogRunParamsVirtualDiskInfo) UnmarshalBinary(b []byte) error {
	var res CdpLogRunParamsVirtualDiskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
