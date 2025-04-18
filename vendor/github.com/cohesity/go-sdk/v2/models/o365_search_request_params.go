// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// O365SearchRequestParams O365 Search Request Params
//
// Specifies O365 specific params search request params to search for indexed items.
//
// swagger:model O365SearchRequestParams
type O365SearchRequestParams struct {

	// Specifies the domain Ids in which indexed items are searched.
	DomainIds []int64 `json:"domainIds"`

	// Specifies the user ids across which the indexed items needs to be searched.
	UserIds []int64 `json:"userIds"`

	// Specifies the Sharepoint site ids across which the indexed items needs to be searched.
	SiteIds []int64 `json:"siteIds"`

	// Specifies the Group ids across which the indexed items needs to be searched.
	GroupIds []int64 `json:"groupIds"`

	// Specifies the Teams ids across which the indexed items needs to be searched.
	TeamsIds []int64 `json:"teamsIds"`
}

// Validate validates this o365 search request params
func (m *O365SearchRequestParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o365 search request params based on context it is used
func (m *O365SearchRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *O365SearchRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *O365SearchRequestParams) UnmarshalBinary(b []byte) error {
	var res O365SearchRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
