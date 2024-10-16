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

// OraclePdbRestoreParams Specifies information about the list of pdbs to be restored.
//
// swagger:model OraclePdbRestoreParams
type OraclePdbRestoreParams struct {

	// Specifies if the PDB should be ignored if a PDB already exists with same name.
	DropDuplicatePDB *bool `json:"dropDuplicatePDB,omitempty"`

	// Specifies list of PDB objects to restore.
	PdbObjects []*OraclePdbObjectInfo `json:"pdbObjects"`

	// Specifies if pdbs should be restored to an existing CDB.
	RestoreToExistingCdb *bool `json:"restoreToExistingCdb,omitempty"`

	// Specifies the new PDB name mapping to existing PDBs.
	RenamePdbMap []*KeyValuePair `json:"renamePdbMap"`

	// Specifies whether to restore or skip the provided PDBs list.
	IncludeInRestore *bool `json:"includeInRestore,omitempty"`
}

// Validate validates this oracle pdb restore params
func (m *OraclePdbRestoreParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePdbObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenamePdbMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OraclePdbRestoreParams) validatePdbObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.PdbObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.PdbObjects); i++ {
		if swag.IsZero(m.PdbObjects[i]) { // not required
			continue
		}

		if m.PdbObjects[i] != nil {
			if err := m.PdbObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pdbObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pdbObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OraclePdbRestoreParams) validateRenamePdbMap(formats strfmt.Registry) error {
	if swag.IsZero(m.RenamePdbMap) { // not required
		return nil
	}

	for i := 0; i < len(m.RenamePdbMap); i++ {
		if swag.IsZero(m.RenamePdbMap[i]) { // not required
			continue
		}

		if m.RenamePdbMap[i] != nil {
			if err := m.RenamePdbMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this oracle pdb restore params based on the context it is used
func (m *OraclePdbRestoreParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePdbObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRenamePdbMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OraclePdbRestoreParams) contextValidatePdbObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PdbObjects); i++ {

		if m.PdbObjects[i] != nil {

			if swag.IsZero(m.PdbObjects[i]) { // not required
				return nil
			}

			if err := m.PdbObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pdbObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pdbObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OraclePdbRestoreParams) contextValidateRenamePdbMap(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RenamePdbMap); i++ {

		if m.RenamePdbMap[i] != nil {

			if swag.IsZero(m.RenamePdbMap[i]) { // not required
				return nil
			}

			if err := m.RenamePdbMap[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("renamePdbMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OraclePdbRestoreParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OraclePdbRestoreParams) UnmarshalBinary(b []byte) error {
	var res OraclePdbRestoreParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
