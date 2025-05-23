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

// IpmiUsers Specifies the list of IPMI users for the given node.
//
// swagger:model IpmiUsers
type IpmiUsers struct {

	// Specifies the channel through which the IPMI interface communicates on the network.
	Channel *int32 `json:"channel,omitempty"`

	// Specifies the maximum value of user id among all the users.
	MaxUserID *int32 `json:"maxUserId,omitempty"`

	// Specifies the list of ipmi users with their permissions.
	UserList []*IpmiUserInfo `json:"userList"`
}

// Validate validates this ipmi users
func (m *IpmiUsers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpmiUsers) validateUserList(formats strfmt.Registry) error {
	if swag.IsZero(m.UserList) { // not required
		return nil
	}

	for i := 0; i < len(m.UserList); i++ {
		if swag.IsZero(m.UserList[i]) { // not required
			continue
		}

		if m.UserList[i] != nil {
			if err := m.UserList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipmi users based on the context it is used
func (m *IpmiUsers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpmiUsers) contextValidateUserList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserList); i++ {

		if m.UserList[i] != nil {

			if swag.IsZero(m.UserList[i]) { // not required
				return nil
			}

			if err := m.UserList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpmiUsers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpmiUsers) UnmarshalBinary(b []byte) error {
	var res IpmiUsers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
