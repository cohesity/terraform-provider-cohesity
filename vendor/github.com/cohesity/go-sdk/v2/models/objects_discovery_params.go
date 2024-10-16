// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObjectsDiscoveryParams Office365 Objects Discovery parameters.
//
// Specifies the parameters used for discovering the office 365 objects selectively during source registration or refresh.
//
// swagger:model ObjectsDiscoveryParams
type ObjectsDiscoveryParams struct {

	// Specifies the list of object types that will be discovered as part of source registration or refresh.
	DiscoverableObjectTypeList []string `json:"discoverableObjectTypeList"`

	// Specifies the discovery params for User(mailbox/onedrive) entities.
	UsersDiscoveryParams *UsersDiscoveryParams `json:"usersDiscoveryParams,omitempty"`

	// Specifies the discovery params for SharePoint site entities.
	SitesDiscoveryParams *SitesDiscoveryParams `json:"sitesDiscoveryParams,omitempty"`

	// Specifies additional params for Teams entities.
	TeamsAdditionalParams *TeamsAdditionalParams `json:"teamsAdditionalParams,omitempty"`
}

// Validate validates this objects discovery params
func (m *ObjectsDiscoveryParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscoverableObjectTypeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsersDiscoveryParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitesDiscoveryParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamsAdditionalParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var objectsDiscoveryParamsDiscoverableObjectTypeListItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kDomain","kOutlook","kMailbox","kUsers","kGroups","kSites","kUser","kGroup","kSite","kApplication","kGraphUser","kPublicFolders","kPublicFolder","kTeams","kTeam","kRootPublicFolder","kO365Exchange","kO365OneDrive","kO365Sharepoint"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectsDiscoveryParamsDiscoverableObjectTypeListItemsEnum = append(objectsDiscoveryParamsDiscoverableObjectTypeListItemsEnum, v)
	}
}

func (m *ObjectsDiscoveryParams) validateDiscoverableObjectTypeListItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectsDiscoveryParamsDiscoverableObjectTypeListItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectsDiscoveryParams) validateDiscoverableObjectTypeList(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscoverableObjectTypeList) { // not required
		return nil
	}

	for i := 0; i < len(m.DiscoverableObjectTypeList); i++ {

		// value enum
		if err := m.validateDiscoverableObjectTypeListItemsEnum("discoverableObjectTypeList"+"."+strconv.Itoa(i), "body", m.DiscoverableObjectTypeList[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectsDiscoveryParams) validateUsersDiscoveryParams(formats strfmt.Registry) error {
	if swag.IsZero(m.UsersDiscoveryParams) { // not required
		return nil
	}

	if m.UsersDiscoveryParams != nil {
		if err := m.UsersDiscoveryParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usersDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectsDiscoveryParams) validateSitesDiscoveryParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SitesDiscoveryParams) { // not required
		return nil
	}

	if m.SitesDiscoveryParams != nil {
		if err := m.SitesDiscoveryParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sitesDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sitesDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectsDiscoveryParams) validateTeamsAdditionalParams(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamsAdditionalParams) { // not required
		return nil
	}

	if m.TeamsAdditionalParams != nil {
		if err := m.TeamsAdditionalParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamsAdditionalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("teamsAdditionalParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this objects discovery params based on the context it is used
func (m *ObjectsDiscoveryParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsersDiscoveryParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSitesDiscoveryParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeamsAdditionalParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectsDiscoveryParams) contextValidateUsersDiscoveryParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UsersDiscoveryParams != nil {

		if swag.IsZero(m.UsersDiscoveryParams) { // not required
			return nil
		}

		if err := m.UsersDiscoveryParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usersDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectsDiscoveryParams) contextValidateSitesDiscoveryParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SitesDiscoveryParams != nil {

		if swag.IsZero(m.SitesDiscoveryParams) { // not required
			return nil
		}

		if err := m.SitesDiscoveryParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sitesDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sitesDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectsDiscoveryParams) contextValidateTeamsAdditionalParams(ctx context.Context, formats strfmt.Registry) error {

	if m.TeamsAdditionalParams != nil {

		if swag.IsZero(m.TeamsAdditionalParams) { // not required
			return nil
		}

		if err := m.TeamsAdditionalParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamsAdditionalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("teamsAdditionalParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectsDiscoveryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectsDiscoveryParams) UnmarshalBinary(b []byte) error {
	var res ObjectsDiscoveryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
