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

// AcropolisDiskInfo Acropolis Disk Information
//
// Specifies information about a disk to be filtered.
//
// swagger:model AcropolisDiskInfo
type AcropolisDiskInfo struct {

	// Specifies the disk controller type.
	// Required: true
	// Enum: ["scsi","ide","pci","sata","spapr","nvme"]
	ControllerType *string `json:"controllerType"`

	// Specifies the disk index number.
	// Required: true
	UnitNumber *int32 `json:"unitNumber"`
}

// Validate validates this acropolis disk info
func (m *AcropolisDiskInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControllerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var acropolisDiskInfoTypeControllerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["scsi","ide","pci","sata","spapr","nvme"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		acropolisDiskInfoTypeControllerTypePropEnum = append(acropolisDiskInfoTypeControllerTypePropEnum, v)
	}
}

const (

	// AcropolisDiskInfoControllerTypeScsi captures enum value "scsi"
	AcropolisDiskInfoControllerTypeScsi string = "scsi"

	// AcropolisDiskInfoControllerTypeIde captures enum value "ide"
	AcropolisDiskInfoControllerTypeIde string = "ide"

	// AcropolisDiskInfoControllerTypePci captures enum value "pci"
	AcropolisDiskInfoControllerTypePci string = "pci"

	// AcropolisDiskInfoControllerTypeSata captures enum value "sata"
	AcropolisDiskInfoControllerTypeSata string = "sata"

	// AcropolisDiskInfoControllerTypeSpapr captures enum value "spapr"
	AcropolisDiskInfoControllerTypeSpapr string = "spapr"

	// AcropolisDiskInfoControllerTypeNvme captures enum value "nvme"
	AcropolisDiskInfoControllerTypeNvme string = "nvme"
)

// prop value enum
func (m *AcropolisDiskInfo) validateControllerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, acropolisDiskInfoTypeControllerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AcropolisDiskInfo) validateControllerType(formats strfmt.Registry) error {

	if err := validate.Required("controllerType", "body", m.ControllerType); err != nil {
		return err
	}

	// value enum
	if err := m.validateControllerTypeEnum("controllerType", "body", *m.ControllerType); err != nil {
		return err
	}

	return nil
}

func (m *AcropolisDiskInfo) validateUnitNumber(formats strfmt.Registry) error {

	if err := validate.Required("unitNumber", "body", m.UnitNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this acropolis disk info based on context it is used
func (m *AcropolisDiskInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AcropolisDiskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcropolisDiskInfo) UnmarshalBinary(b []byte) error {
	var res AcropolisDiskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
