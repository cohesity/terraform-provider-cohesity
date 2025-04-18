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

// Office365ObjectProtectionParams Specifies the parameters which are specific to Microsoft 365 Object Protection.
//
// swagger:model Office365ObjectProtectionParams
type Office365ObjectProtectionParams struct {

	// Specifies the Microsoft 365 Object Protection type.
	// Required: true
	// Enum: ["kMailbox","kOneDrive","kSharePoint","kPublicFolders","kGroups","kTeams","kMailboxCSM","kOneDriveCSM","kSharePointCSM"]
	ObjectProtectionType *string `json:"objectProtectionType"`

	// Specifies the User Mailbox Object Protection params.
	UserMailboxObjectProtectionParams *Office365UserMailboxObjectProtectionParams `json:"userMailboxObjectProtectionParams,omitempty"`

	// Specifies the User OneDrive Protection params.
	UserOneDriveObjectProtectionParams *Office365UserOneDriveObjectProtectionParams `json:"userOneDriveObjectProtectionParams,omitempty"`

	// Specifies the Sharepoint site object Protection params.
	SharepointSiteObjectProtectionParams *Office365SharepointSiteObjectProtectionParams `json:"sharepointSiteObjectProtectionParams,omitempty"`

	// Specifies the Teams Object Protection params.
	TeamsObjectProtectionParams *Office365TeamsObjectProtectionParams `json:"teamsObjectProtectionParams,omitempty"`

	// Specifies the Groups Object Protection params.
	GroupsObjectProtectionParams *Office365GroupsObjectProtectionParams `json:"groupsObjectProtectionParams,omitempty"`
}

// Validate validates this office365 object protection params
func (m *Office365ObjectProtectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectProtectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserMailboxObjectProtectionParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserOneDriveObjectProtectionParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharepointSiteObjectProtectionParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamsObjectProtectionParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupsObjectProtectionParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var office365ObjectProtectionParamsTypeObjectProtectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kMailbox","kOneDrive","kSharePoint","kPublicFolders","kGroups","kTeams","kMailboxCSM","kOneDriveCSM","kSharePointCSM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		office365ObjectProtectionParamsTypeObjectProtectionTypePropEnum = append(office365ObjectProtectionParamsTypeObjectProtectionTypePropEnum, v)
	}
}

const (

	// Office365ObjectProtectionParamsObjectProtectionTypeKMailbox captures enum value "kMailbox"
	Office365ObjectProtectionParamsObjectProtectionTypeKMailbox string = "kMailbox"

	// Office365ObjectProtectionParamsObjectProtectionTypeKOneDrive captures enum value "kOneDrive"
	Office365ObjectProtectionParamsObjectProtectionTypeKOneDrive string = "kOneDrive"

	// Office365ObjectProtectionParamsObjectProtectionTypeKSharePoint captures enum value "kSharePoint"
	Office365ObjectProtectionParamsObjectProtectionTypeKSharePoint string = "kSharePoint"

	// Office365ObjectProtectionParamsObjectProtectionTypeKPublicFolders captures enum value "kPublicFolders"
	Office365ObjectProtectionParamsObjectProtectionTypeKPublicFolders string = "kPublicFolders"

	// Office365ObjectProtectionParamsObjectProtectionTypeKGroups captures enum value "kGroups"
	Office365ObjectProtectionParamsObjectProtectionTypeKGroups string = "kGroups"

	// Office365ObjectProtectionParamsObjectProtectionTypeKTeams captures enum value "kTeams"
	Office365ObjectProtectionParamsObjectProtectionTypeKTeams string = "kTeams"

	// Office365ObjectProtectionParamsObjectProtectionTypeKMailboxCSM captures enum value "kMailboxCSM"
	Office365ObjectProtectionParamsObjectProtectionTypeKMailboxCSM string = "kMailboxCSM"

	// Office365ObjectProtectionParamsObjectProtectionTypeKOneDriveCSM captures enum value "kOneDriveCSM"
	Office365ObjectProtectionParamsObjectProtectionTypeKOneDriveCSM string = "kOneDriveCSM"

	// Office365ObjectProtectionParamsObjectProtectionTypeKSharePointCSM captures enum value "kSharePointCSM"
	Office365ObjectProtectionParamsObjectProtectionTypeKSharePointCSM string = "kSharePointCSM"
)

// prop value enum
func (m *Office365ObjectProtectionParams) validateObjectProtectionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, office365ObjectProtectionParamsTypeObjectProtectionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Office365ObjectProtectionParams) validateObjectProtectionType(formats strfmt.Registry) error {

	if err := validate.Required("objectProtectionType", "body", m.ObjectProtectionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateObjectProtectionTypeEnum("objectProtectionType", "body", *m.ObjectProtectionType); err != nil {
		return err
	}

	return nil
}

func (m *Office365ObjectProtectionParams) validateUserMailboxObjectProtectionParams(formats strfmt.Registry) error {
	if swag.IsZero(m.UserMailboxObjectProtectionParams) { // not required
		return nil
	}

	if m.UserMailboxObjectProtectionParams != nil {
		if err := m.UserMailboxObjectProtectionParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userMailboxObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userMailboxObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) validateUserOneDriveObjectProtectionParams(formats strfmt.Registry) error {
	if swag.IsZero(m.UserOneDriveObjectProtectionParams) { // not required
		return nil
	}

	if m.UserOneDriveObjectProtectionParams != nil {
		if err := m.UserOneDriveObjectProtectionParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userOneDriveObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userOneDriveObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) validateSharepointSiteObjectProtectionParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SharepointSiteObjectProtectionParams) { // not required
		return nil
	}

	if m.SharepointSiteObjectProtectionParams != nil {
		if err := m.SharepointSiteObjectProtectionParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharepointSiteObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharepointSiteObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) validateTeamsObjectProtectionParams(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamsObjectProtectionParams) { // not required
		return nil
	}

	if m.TeamsObjectProtectionParams != nil {
		if err := m.TeamsObjectProtectionParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamsObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("teamsObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) validateGroupsObjectProtectionParams(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupsObjectProtectionParams) { // not required
		return nil
	}

	if m.GroupsObjectProtectionParams != nil {
		if err := m.GroupsObjectProtectionParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupsObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupsObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this office365 object protection params based on the context it is used
func (m *Office365ObjectProtectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserMailboxObjectProtectionParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserOneDriveObjectProtectionParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSharepointSiteObjectProtectionParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeamsObjectProtectionParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupsObjectProtectionParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Office365ObjectProtectionParams) contextValidateUserMailboxObjectProtectionParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UserMailboxObjectProtectionParams != nil {

		if swag.IsZero(m.UserMailboxObjectProtectionParams) { // not required
			return nil
		}

		if err := m.UserMailboxObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userMailboxObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userMailboxObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) contextValidateUserOneDriveObjectProtectionParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UserOneDriveObjectProtectionParams != nil {

		if swag.IsZero(m.UserOneDriveObjectProtectionParams) { // not required
			return nil
		}

		if err := m.UserOneDriveObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userOneDriveObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userOneDriveObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) contextValidateSharepointSiteObjectProtectionParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SharepointSiteObjectProtectionParams != nil {

		if swag.IsZero(m.SharepointSiteObjectProtectionParams) { // not required
			return nil
		}

		if err := m.SharepointSiteObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharepointSiteObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharepointSiteObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) contextValidateTeamsObjectProtectionParams(ctx context.Context, formats strfmt.Registry) error {

	if m.TeamsObjectProtectionParams != nil {

		if swag.IsZero(m.TeamsObjectProtectionParams) { // not required
			return nil
		}

		if err := m.TeamsObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamsObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("teamsObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

func (m *Office365ObjectProtectionParams) contextValidateGroupsObjectProtectionParams(ctx context.Context, formats strfmt.Registry) error {

	if m.GroupsObjectProtectionParams != nil {

		if swag.IsZero(m.GroupsObjectProtectionParams) { // not required
			return nil
		}

		if err := m.GroupsObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupsObjectProtectionParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupsObjectProtectionParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Office365ObjectProtectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Office365ObjectProtectionParams) UnmarshalBinary(b []byte) error {
	var res Office365ObjectProtectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
