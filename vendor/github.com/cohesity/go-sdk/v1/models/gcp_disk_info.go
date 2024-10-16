// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GcpDiskInfo GCP disk Info.
//
// Specified info about the GCP disk.
//
// swagger:model GcpDiskInfo
type GcpDiskInfo struct {

	// Specifies the name of the device. Eg - /dev/sdb.
	DeviceName *string `json:"deviceName,omitempty"`

	// Specified ID of the disk.
	ID *int64 `json:"id,omitempty"`

	// Specifies if the disk is attached as root device.
	IsRootDevice *bool `json:"isRootDevice,omitempty"`

	// Specifies the name of the disk.
	Name *string `json:"name,omitempty"`

	// Specifies the size of the device.
	SizeGb *int64 `json:"sizeGb,omitempty"`

	// Specifies the type of the disk.
	Type *string `json:"type,omitempty"`
}

// Validate validates this gcp disk info
func (m *GcpDiskInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gcp disk info based on context it is used
func (m *GcpDiskInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpDiskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpDiskInfo) UnmarshalBinary(b []byte) error {
	var res GcpDiskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
