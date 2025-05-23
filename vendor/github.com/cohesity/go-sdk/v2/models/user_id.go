// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserID Specifies the User Id type. Either sid or unixUid should be set.
//
// swagger:model UserId
type UserID struct {

	// Specifies the user sid.
	Sid *string `json:"sid,omitempty"`

	// Specifies the unix Uid.
	UnixUID *uint32 `json:"unixUid,omitempty"`

	// Specifies the full name of the user
	UserName *string `json:"userName,omitempty"`

	// Specifies the domain name of the user, where the principal' account is maintained.
	Domain *string `json:"domain,omitempty"`
}

// Validate validates this user Id
func (m *UserID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user Id based on context it is used
func (m *UserID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserID) UnmarshalBinary(b []byte) error {
	var res UserID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
