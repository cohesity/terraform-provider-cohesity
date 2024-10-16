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

// SQLAagHostAndDatabases SQL AAG Host and Databases.
//
// Specifies AAGs and databases information on an SQL server. If AAGs exist
// on the server, specifies information about the AAG and databases in the
// group for each AAG found on the server.
//
// swagger:model SqlAagHostAndDatabases
type SQLAagHostAndDatabases struct {

	// Specifies a list of AAGs and database members in each AAG.
	AagDatabases []*AagAndDatabases `json:"aagDatabases"`

	// Specifies the subtree used to store application-level Objects.
	// Different environments use the subtree to store application-level
	// information. For example for SQL Server, this subtree stores the
	// SQL Server instances running on a VM or Physical hosts.
	ApplicationNode *ProtectionSourceNode `json:"applicationNode,omitempty"`

	// Specifies all database entities found on the server. Database may or
	// may not be in an AAG.
	Databases []*ProtectionSource `json:"databases"`

	// Specifies an error message when the host is not registered as an SQL
	// host.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Specifies the name of the host that is not registered as an SQL server
	// on Cohesity Cluser.
	UnknownHostName *string `json:"unknownHostName,omitempty"`
}

// Validate validates this Sql aag host and databases
func (m *SQLAagHostAndDatabases) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAagDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLAagHostAndDatabases) validateAagDatabases(formats strfmt.Registry) error {
	if swag.IsZero(m.AagDatabases) { // not required
		return nil
	}

	for i := 0; i < len(m.AagDatabases); i++ {
		if swag.IsZero(m.AagDatabases[i]) { // not required
			continue
		}

		if m.AagDatabases[i] != nil {
			if err := m.AagDatabases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aagDatabases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aagDatabases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SQLAagHostAndDatabases) validateApplicationNode(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationNode) { // not required
		return nil
	}

	if m.ApplicationNode != nil {
		if err := m.ApplicationNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationNode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationNode")
			}
			return err
		}
	}

	return nil
}

func (m *SQLAagHostAndDatabases) validateDatabases(formats strfmt.Registry) error {
	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	for i := 0; i < len(m.Databases); i++ {
		if swag.IsZero(m.Databases[i]) { // not required
			continue
		}

		if m.Databases[i] != nil {
			if err := m.Databases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("databases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("databases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Sql aag host and databases based on the context it is used
func (m *SQLAagHostAndDatabases) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAagDatabases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplicationNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatabases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLAagHostAndDatabases) contextValidateAagDatabases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AagDatabases); i++ {

		if m.AagDatabases[i] != nil {

			if swag.IsZero(m.AagDatabases[i]) { // not required
				return nil
			}

			if err := m.AagDatabases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aagDatabases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aagDatabases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SQLAagHostAndDatabases) contextValidateApplicationNode(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationNode != nil {

		if swag.IsZero(m.ApplicationNode) { // not required
			return nil
		}

		if err := m.ApplicationNode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationNode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationNode")
			}
			return err
		}
	}

	return nil
}

func (m *SQLAagHostAndDatabases) contextValidateDatabases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Databases); i++ {

		if m.Databases[i] != nil {

			if swag.IsZero(m.Databases[i]) { // not required
				return nil
			}

			if err := m.Databases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("databases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("databases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLAagHostAndDatabases) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLAagHostAndDatabases) UnmarshalBinary(b []byte) error {
	var res SQLAagHostAndDatabases
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
