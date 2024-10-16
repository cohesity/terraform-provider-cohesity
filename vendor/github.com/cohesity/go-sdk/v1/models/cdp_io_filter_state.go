// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdpIoFilterState CdpIoFilterState specifies the current state of the CDP IO Filter for a
// single Protection Source.
//
// swagger:model CdpIoFilterState
type CdpIoFilterState struct {

	// Specifies the message of the error, which was encountered duing the last
	// upgrade. If this is specified, then it means that the last upgrade failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Specifies the current status of the IO Filter.
	// See magneto/base/entities.proto > IOFilterState > FilterStatus
	FilterStatus *string `json:"filterStatus,omitempty"`

	// Specifies the current upgradability status of the IO Filter.
	// See magneto/base/common.proto > Upgradability
	Upgradability *string `json:"upgradability,omitempty"`

	// Specifies the current version of the IO filter.
	Version *string `json:"version,omitempty"`
}

// Validate validates this cdp io filter state
func (m *CdpIoFilterState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cdp io filter state based on context it is used
func (m *CdpIoFilterState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdpIoFilterState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdpIoFilterState) UnmarshalBinary(b []byte) error {
	var res CdpIoFilterState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
