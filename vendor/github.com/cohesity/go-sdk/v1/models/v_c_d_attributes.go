// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VCDAttributes Message capturing additional vCD related attributes.
//
// swagger:model VCDAttributes
type VCDAttributes struct {

	// The VCD uuid for catalog.
	CatalogID *string `json:"catalogId,omitempty"`

	// The VCD uuid for catalog item.
	CatalogItemUUID *string `json:"catalogItemUuid,omitempty"`

	// Is this is a standalone VCD VM.
	IsStandaloneVM *bool `json:"isStandaloneVm,omitempty"`

	// If the vCD entity is a kVirtualApp and this vApp is auto-generated (for
	// standalone VM), then this will be set to true and false otherwise.
	IsVappAutoGenerated *bool `json:"isVappAutoGenerated,omitempty"`

	// If this is a vApp template. Only applicable of kvApp type objects.
	IsVappTemplate *bool `json:"isVappTemplate,omitempty"`

	// This contains the entity id of the parent VDC.
	ParentVdcUUID *string `json:"parentVdcUuid,omitempty"`

	// This contains the name or the provider VDC.
	ProviderVdcName *string `json:"providerVdcName,omitempty"`

	// This contains the uuid or the provider VDC.
	ProviderVdcUUID *string `json:"providerVdcUuid,omitempty"`

	// This contains the entity-id of the resource pool, the VM is part of.
	ResgrpEntityID *int64 `json:"resgrpEntityId,omitempty"`

	// This contains the MORef of the resource pool, this VM entity is part of.
	ResourcePoolMoref *string `json:"resourcePoolMoref,omitempty"`

	// The storage profiles associated with the VM. This is only populated
	// for VM type entity.
	StorageProfileUUID *string `json:"storageProfileUuid,omitempty"`

	// This contains the MORef of the VCDs resource pool.
	VcdResourcePoolMoref *MORef `json:"vcdResourcePoolMoref,omitempty"`

	// This is set for certain objects of vCloud director.
	VcdUUID *string `json:"vcdUuid,omitempty"`

	// This contains the vCenter endpoint.
	VcenterEndpoint *string `json:"vcenterEndpoint,omitempty"`

	// This contains the vCenter id.
	VcenterID *string `json:"vcenterId,omitempty"`

	// This contains the vCenter name.
	// TODO(Chinmaya): Replace this with entity id or the UUID of the vCenter.
	VcenterName *string `json:"vcenterName,omitempty"`
}

// Validate validates this v c d attributes
func (m *VCDAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcdResourcePoolMoref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VCDAttributes) validateVcdResourcePoolMoref(formats strfmt.Registry) error {
	if swag.IsZero(m.VcdResourcePoolMoref) { // not required
		return nil
	}

	if m.VcdResourcePoolMoref != nil {
		if err := m.VcdResourcePoolMoref.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcdResourcePoolMoref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcdResourcePoolMoref")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v c d attributes based on the context it is used
func (m *VCDAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcdResourcePoolMoref(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VCDAttributes) contextValidateVcdResourcePoolMoref(ctx context.Context, formats strfmt.Registry) error {

	if m.VcdResourcePoolMoref != nil {

		if swag.IsZero(m.VcdResourcePoolMoref) { // not required
			return nil
		}

		if err := m.VcdResourcePoolMoref.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcdResourcePoolMoref")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcdResourcePoolMoref")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VCDAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VCDAttributes) UnmarshalBinary(b []byte) error {
	var res VCDAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
