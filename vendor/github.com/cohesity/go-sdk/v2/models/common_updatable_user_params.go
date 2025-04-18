// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonUpdatableUserParams Specifies user properties which can be updated.
//
// swagger:model CommonUpdatableUserParams
type CommonUpdatableUserParams struct {

	// Specifies the description of the User.
	Description *string `json:"description,omitempty"`

	// Specifies the Cohesity roles to associate with the user. The Cohesity roles determine privileges on the Cohesity Cluster for this user.
	Roles []string `json:"roles"`

	// Specifies whether the User is restricted. A restricted user can only view & manage the objects it has permissions to.
	Restricted *bool `json:"restricted,omitempty"`

	// Specifies the epoch time in milliseconds since when the user can login.
	EffectiveTimeMsecs *int64 `json:"effectiveTimeMsecs,omitempty"`

	// Specifies the epoch time in milliseconds when the user expires. Post expiry the user cannot access Cohesity cluster.
	ExpiryTimeMsecs *int64 `json:"expiryTimeMsecs,omitempty"`

	// Specifies whether the User is locked.
	Locked *bool `json:"locked,omitempty"`
}

// Validate validates this common updatable user params
func (m *CommonUpdatableUserParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common updatable user params based on context it is used
func (m *CommonUpdatableUserParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonUpdatableUserParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonUpdatableUserParams) UnmarshalBinary(b []byte) error {
	var res CommonUpdatableUserParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
