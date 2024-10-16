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

// BackupTaskStateVappInfoProto This message contains information on the VApp being being backed up.
//
// swagger:model BackupTaskStateVappInfoProto
type BackupTaskStateVappInfoProto struct {

	// The catalog to which the template belongs to. This is only applicable if
	// is_vapp_template_backup is true.
	CatalogUUID *string `json:"catalogUuid,omitempty"`

	// Whether the vApp is being backed up during a vApp template backup.
	IsVappTemplateBackup *bool `json:"isVappTemplateBackup,omitempty"`

	// Wehther the VApp is deployed.
	VappDeployed *bool `json:"vappDeployed,omitempty"`

	// Contains the vApp entity id.
	VappEntity *int64 `json:"vappEntity,omitempty"`

	// Contains the VApp href.
	VappHref *string `json:"vappHref,omitempty"`

	// Contains the VApp Id.
	VappID *string `json:"vappId,omitempty"`

	// Wehther the VApp is in maintenance mode.
	VappInMaintenanceMode *bool `json:"vappInMaintenanceMode,omitempty"`

	// Contains the entire VApp info fetched from VCD in xml format.
	VappInfoXML *string `json:"vappInfoXml,omitempty"`

	// Contains the metadata associated with the vApp.
	VappMetadata *MetadataInfo `json:"vappMetadata,omitempty"`

	// Contains the VApp name.
	VappName *string `json:"vappName,omitempty"`

	// Contains the vApp's network config XML string.
	VappNetworkXML *string `json:"vappNetworkXml,omitempty"`

	// Contains the VApp's vDC Id.
	VdcID *string `json:"vdcId,omitempty"`
}

// Validate validates this backup task state vapp info proto
func (m *BackupTaskStateVappInfoProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVappMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateVappInfoProto) validateVappMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.VappMetadata) { // not required
		return nil
	}

	if m.VappMetadata != nil {
		if err := m.VappMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vappMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vappMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup task state vapp info proto based on the context it is used
func (m *BackupTaskStateVappInfoProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVappMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateVappInfoProto) contextValidateVappMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.VappMetadata != nil {

		if swag.IsZero(m.VappMetadata) { // not required
			return nil
		}

		if err := m.VappMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vappMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vappMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupTaskStateVappInfoProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupTaskStateVappInfoProto) UnmarshalBinary(b []byte) error {
	var res BackupTaskStateVappInfoProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
