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

// GroupParameters Group Request.
//
// Specifies the settings used to add/create a new group or modify an existing
// group.
//
// swagger:model GroupParameters
type GroupParameters struct {

	// Specifies a description of the group.
	Description *string `json:"description,omitempty"`

	// Specifies the domain of the group.
	Domain *string `json:"domain,omitempty"`

	// Specifies the name of the group.
	Name *string `json:"name,omitempty"`

	// Whether the group is a restricted group. Users belonging to a restricted
	// group can only view objects they have permissions to.
	Restricted *bool `json:"restricted,omitempty"`

	// Array of Roles.
	//
	// Specifies the Cohesity role names to associate with the group.
	// role names can be viewed by GET /irisservices/api/v1/public/roles,
	// example: COHESITY_VIEWER
	// The Cohesity roles determine privileges on the Cohesity Cluster
	// for all the users in this group.
	Roles []string `json:"roles"`

	// Specifies the SMB principals. Principals will be added to this group only
	// if IsSmbPrincipalOnly set to true.
	SmbPrincipals []*SmbPrincipal `json:"smbPrincipals"`

	// Specifies the tenants to which the group belongs to. If not specified,
	// session user's tenant id is assumed.
	TenantIds []string `json:"tenantIds"`

	// Specifies the SID of users who are members of the group.
	// This field is used only for local groups.
	Users []string `json:"users"`
}

// Validate validates this group parameters
func (m *GroupParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSmbPrincipals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupParameters) validateSmbPrincipals(formats strfmt.Registry) error {
	if swag.IsZero(m.SmbPrincipals) { // not required
		return nil
	}

	for i := 0; i < len(m.SmbPrincipals); i++ {
		if swag.IsZero(m.SmbPrincipals[i]) { // not required
			continue
		}

		if m.SmbPrincipals[i] != nil {
			if err := m.SmbPrincipals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("smbPrincipals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("smbPrincipals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this group parameters based on the context it is used
func (m *GroupParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSmbPrincipals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupParameters) contextValidateSmbPrincipals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SmbPrincipals); i++ {

		if m.SmbPrincipals[i] != nil {

			if swag.IsZero(m.SmbPrincipals[i]) { // not required
				return nil
			}

			if err := m.SmbPrincipals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("smbPrincipals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("smbPrincipals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupParameters) UnmarshalBinary(b []byte) error {
	var res GroupParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
