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

// MsGroupItem Indexed M365 Group item.
//
// Specifies the indexed M365 Group item.
//
// swagger:model MsGroupItem
type MsGroupItem struct {
	CommonIndexedObjectParams

	// Specifies the M365 Group item type.
	// Enum: ["Email","EmailFolder","SiteFile","SiteFolder"]
	Type *string `json:"type,omitempty"`

	// Specifies the indexed M365 Group mailbox item.
	MailboxItem *Email `json:"mailboxItem,omitempty"`

	// Specifies the indexed Group site document library item.
	SiteItem *DocumentLibraryItem `json:"siteItem,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MsGroupItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonIndexedObjectParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonIndexedObjectParams = aO0

	// AO1
	var dataAO1 struct {
		Type *string `json:"type,omitempty"`

		MailboxItem *Email `json:"mailboxItem,omitempty"`

		SiteItem *DocumentLibraryItem `json:"siteItem,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Type = dataAO1.Type

	m.MailboxItem = dataAO1.MailboxItem

	m.SiteItem = dataAO1.SiteItem

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MsGroupItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonIndexedObjectParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Type *string `json:"type,omitempty"`

		MailboxItem *Email `json:"mailboxItem,omitempty"`

		SiteItem *DocumentLibraryItem `json:"siteItem,omitempty"`
	}

	dataAO1.Type = m.Type

	dataAO1.MailboxItem = m.MailboxItem

	dataAO1.SiteItem = m.SiteItem

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ms group item
func (m *MsGroupItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIndexedObjectParams
	if err := m.CommonIndexedObjectParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMailboxItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiteItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msGroupItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Email","EmailFolder","SiteFile","SiteFolder"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msGroupItemTypeTypePropEnum = append(msGroupItemTypeTypePropEnum, v)
	}
}

// property enum
func (m *MsGroupItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msGroupItemTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsGroupItem) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *MsGroupItem) validateMailboxItem(formats strfmt.Registry) error {

	if swag.IsZero(m.MailboxItem) { // not required
		return nil
	}

	if m.MailboxItem != nil {
		if err := m.MailboxItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mailboxItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mailboxItem")
			}
			return err
		}
	}

	return nil
}

func (m *MsGroupItem) validateSiteItem(formats strfmt.Registry) error {

	if swag.IsZero(m.SiteItem) { // not required
		return nil
	}

	if m.SiteItem != nil {
		if err := m.SiteItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("siteItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("siteItem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ms group item based on the context it is used
func (m *MsGroupItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIndexedObjectParams
	if err := m.CommonIndexedObjectParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMailboxItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSiteItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsGroupItem) contextValidateMailboxItem(ctx context.Context, formats strfmt.Registry) error {

	if m.MailboxItem != nil {

		if swag.IsZero(m.MailboxItem) { // not required
			return nil
		}

		if err := m.MailboxItem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mailboxItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mailboxItem")
			}
			return err
		}
	}

	return nil
}

func (m *MsGroupItem) contextValidateSiteItem(ctx context.Context, formats strfmt.Registry) error {

	if m.SiteItem != nil {

		if swag.IsZero(m.SiteItem) { // not required
			return nil
		}

		if err := m.SiteItem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("siteItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("siteItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsGroupItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsGroupItem) UnmarshalBinary(b []byte) error {
	var res MsGroupItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
