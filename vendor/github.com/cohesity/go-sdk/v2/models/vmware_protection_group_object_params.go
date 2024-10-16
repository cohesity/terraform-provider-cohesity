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

// VmwareProtectionGroupObjectParams Specifies the input for a protection object in the VMware environment.
//
// swagger:model VmwareProtectionGroupObjectParams
type VmwareProtectionGroupObjectParams struct {

	// Specifies the id of the object being protected. This can be a leaf level or non leaf level object.
	// Required: true
	ID *int64 `json:"id"`

	// Specifies the type of the VMware object.
	// Enum: ["kVCenter","kStandaloneHost","kvCloudDirector","kFolder","kDatacenter","kComputeResource","kClusterComputeResource","kResourcePool","kDatastore","kHostSystem","kVirtualMachine","kVirtualApp","kStoragePod","kNetwork","kDistributedVirtualPortgroup","kTagCategory","kTag","kOpaqueNetwork","kOrganization","kVirtualDatacenter","kCatalog","kOrgMetadata","kStoragePolicy","kVirtualAppTemplate"]
	Type *string `json:"type,omitempty"`

	// Specifies the name of the virtual machine.
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Specifies whether the vm is part of an Autoprotection and there is at least one object-specific setting applied to this vm. True implies that the vm or its parent directory is autoprotected and will remain part of the autoprotection with additional settings specified here. False implies the object is not part of an Autoprotection and will remain protected and its individual settings here even if a parent directory's Autoprotection setting is altered. Default is false.
	IsAutoprotected *bool `json:"isAutoprotected,omitempty"`

	// Specifies the CDP related information for a given object. This field will only be populated when protection group is configured with policy having CDP retnetion settings.
	// Read Only: true
	CdpInfo *VmwareCdpObject `json:"cdpInfo,omitempty"`

	// Specifies the standby related information for a given object. This field will only be populated when standby is configured in backup job settings.
	// Read Only: true
	StandbyInfo *VmwareStandbyObject `json:"standbyInfo,omitempty"`

	CommonVmwareObjectParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VmwareProtectionGroupObjectParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *int64 `json:"id"`

		Type *string `json:"type,omitempty"`

		Name *string `json:"name,omitempty"`

		IsAutoprotected *bool `json:"isAutoprotected,omitempty"`

		CdpInfo *VmwareCdpObject `json:"cdpInfo,omitempty"`

		StandbyInfo *VmwareStandbyObject `json:"standbyInfo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ID = dataAO0.ID

	m.Type = dataAO0.Type

	m.Name = dataAO0.Name

	m.IsAutoprotected = dataAO0.IsAutoprotected

	m.CdpInfo = dataAO0.CdpInfo

	m.StandbyInfo = dataAO0.StandbyInfo

	// AO1
	var aO1 CommonVmwareObjectParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CommonVmwareObjectParams = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VmwareProtectionGroupObjectParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *int64 `json:"id"`

		Type *string `json:"type,omitempty"`

		Name *string `json:"name,omitempty"`

		IsAutoprotected *bool `json:"isAutoprotected,omitempty"`

		CdpInfo *VmwareCdpObject `json:"cdpInfo,omitempty"`

		StandbyInfo *VmwareStandbyObject `json:"standbyInfo,omitempty"`
	}

	dataAO0.ID = m.ID

	dataAO0.Type = m.Type

	dataAO0.Name = m.Name

	dataAO0.IsAutoprotected = m.IsAutoprotected

	dataAO0.CdpInfo = m.CdpInfo

	dataAO0.StandbyInfo = m.StandbyInfo

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.CommonVmwareObjectParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vmware protection group object params
func (m *VmwareProtectionGroupObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCdpInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandbyInfo(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with CommonVmwareObjectParams
	if err := m.CommonVmwareObjectParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareProtectionGroupObjectParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var vmwareProtectionGroupObjectParamsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVCenter","kStandaloneHost","kvCloudDirector","kFolder","kDatacenter","kComputeResource","kClusterComputeResource","kResourcePool","kDatastore","kHostSystem","kVirtualMachine","kVirtualApp","kStoragePod","kNetwork","kDistributedVirtualPortgroup","kTagCategory","kTag","kOpaqueNetwork","kOrganization","kVirtualDatacenter","kCatalog","kOrgMetadata","kStoragePolicy","kVirtualAppTemplate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmwareProtectionGroupObjectParamsTypeTypePropEnum = append(vmwareProtectionGroupObjectParamsTypeTypePropEnum, v)
	}
}

// property enum
func (m *VmwareProtectionGroupObjectParams) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmwareProtectionGroupObjectParamsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VmwareProtectionGroupObjectParams) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *VmwareProtectionGroupObjectParams) validateCdpInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.CdpInfo) { // not required
		return nil
	}

	if m.CdpInfo != nil {
		if err := m.CdpInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdpInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdpInfo")
			}
			return err
		}
	}

	return nil
}

func (m *VmwareProtectionGroupObjectParams) validateStandbyInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.StandbyInfo) { // not required
		return nil
	}

	if m.StandbyInfo != nil {
		if err := m.StandbyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standbyInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standbyInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vmware protection group object params based on the context it is used
func (m *VmwareProtectionGroupObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCdpInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandbyInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with CommonVmwareObjectParams
	if err := m.CommonVmwareObjectParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareProtectionGroupObjectParams) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VmwareProtectionGroupObjectParams) contextValidateCdpInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CdpInfo != nil {

		if swag.IsZero(m.CdpInfo) { // not required
			return nil
		}

		if err := m.CdpInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdpInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdpInfo")
			}
			return err
		}
	}

	return nil
}

func (m *VmwareProtectionGroupObjectParams) contextValidateStandbyInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.StandbyInfo != nil {

		if swag.IsZero(m.StandbyInfo) { // not required
			return nil
		}

		if err := m.StandbyInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standbyInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standbyInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareProtectionGroupObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareProtectionGroupObjectParams) UnmarshalBinary(b []byte) error {
	var res VmwareProtectionGroupObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
