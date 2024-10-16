// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ADAttributeRestoreParam AD attribute recovery parameters for one or more AD objects.
//
// swagger:model ADAttributeRestoreParam
type ADAttributeRestoreParam struct {

	// Array of LDAP property names to excluded from 'property_vec'. Excluded
	// property name cannot contain wildcard character '*'.  Property names are
	// case insensitive.
	ExcludedPropertyVec []string `json:"excludedPropertyVec"`

	// Array of source and destination object guid pairs to restore attributes.
	// Pair source guid refers to guid in AD snapshot in source_server
	// endpoint. Destination guid refers to guid in production AD. If destination
	// guid is empty, then source guid in AD snapshot should exist in production
	// AD.
	GuidpairVec []*ADGUIDPair `json:"guidpairVec"`

	// Attribute restore option flags of type ADAttributeOptionFlags.
	OptionFlags *uint32 `json:"optionFlags,omitempty"`

	// Array of LDAP property(attribute) names. The name can be standard or
	// custom property defined in AD schema partition. The property can contain
	// wildcard character '*'. If this array is empty, then '*' is assigned,
	// means restore all properties except default system excluded properties.
	// Wildcards will be expanded. If 'memberOf' property is included, group
	// membership of the objects specified in 'guid_vec' will be restored.
	// Property that does not exist for an object is ignored and no error
	// info is returned for that property. Property names are case insensitive.
	// Caller may check the ADAttributeFlags.kSystem obtained during object
	// compare to exclude system properties.
	PropertyVec []string `json:"propertyVec"`

	// When restoring a GPO, need to know the absolute path for SYSVOL folder.
	SrcSysvolFolder *string `json:"srcSysvolFolder,omitempty"`
}

// Validate validates this a d attribute restore param
func (m *ADAttributeRestoreParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGuidpairVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADAttributeRestoreParam) validateGuidpairVec(formats strfmt.Registry) error {
	if swag.IsZero(m.GuidpairVec) { // not required
		return nil
	}

	for i := 0; i < len(m.GuidpairVec); i++ {
		if swag.IsZero(m.GuidpairVec[i]) { // not required
			continue
		}

		if m.GuidpairVec[i] != nil {
			if err := m.GuidpairVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("guidpairVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("guidpairVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this a d attribute restore param based on the context it is used
func (m *ADAttributeRestoreParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGuidpairVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADAttributeRestoreParam) contextValidateGuidpairVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GuidpairVec); i++ {

		if m.GuidpairVec[i] != nil {

			if swag.IsZero(m.GuidpairVec[i]) { // not required
				return nil
			}

			if err := m.GuidpairVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("guidpairVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("guidpairVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ADAttributeRestoreParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ADAttributeRestoreParam) UnmarshalBinary(b []byte) error {
	var res ADAttributeRestoreParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
