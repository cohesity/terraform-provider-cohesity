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

// RecoverAcropolisParams Recover VM params.
//
// Specifies Acropolis related recovery options.
//
// swagger:model RecoverAcropolisParams
type RecoverAcropolisParams struct {

	// Specifies the list of recover Object parameters. This property is mandatory for all recovery action types except recover vms. While recovering VMs, a user can specify snapshots of VM's or a Protection Group Run details to recover all the VM's that are backed up by that Run. For recovering files, specifies the object contains the file to recover.
	Objects []*RecoverAcropolisSnapshotParams `json:"objects"`

	// Specifies the type of recovery action to be performed.
	// Required: true
	// Enum: ["RecoverVMs","RecoverFiles"]
	RecoveryAction *string `json:"recoveryAction"`

	// Specifies the parameters to download files and folders.
	DownloadFileAndFolderParams *CommonDownloadFileAndFolderParams `json:"downloadFileAndFolderParams,omitempty"`

	// Specifies the parameters to recover Acropolis files and folders.
	RecoverFileAndFolderParams *RecoverAcropolisFileAndFolderParams `json:"recoverFileAndFolderParams,omitempty"`

	// Specifies the parameters to recover Acropolis VMs.
	RecoverVMParams *RecoverAcropolisVMParams `json:"recoverVmParams,omitempty"`
}

// Validate validates this recover acropolis params
func (m *RecoverAcropolisParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadFileAndFolderParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverFileAndFolderParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverVMParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAcropolisParams) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var recoverAcropolisParamsTypeRecoveryActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RecoverVMs","RecoverFiles"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverAcropolisParamsTypeRecoveryActionPropEnum = append(recoverAcropolisParamsTypeRecoveryActionPropEnum, v)
	}
}

const (

	// RecoverAcropolisParamsRecoveryActionRecoverVMs captures enum value "RecoverVMs"
	RecoverAcropolisParamsRecoveryActionRecoverVMs string = "RecoverVMs"

	// RecoverAcropolisParamsRecoveryActionRecoverFiles captures enum value "RecoverFiles"
	RecoverAcropolisParamsRecoveryActionRecoverFiles string = "RecoverFiles"
)

// prop value enum
func (m *RecoverAcropolisParams) validateRecoveryActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverAcropolisParamsTypeRecoveryActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverAcropolisParams) validateRecoveryAction(formats strfmt.Registry) error {

	if err := validate.Required("recoveryAction", "body", m.RecoveryAction); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecoveryActionEnum("recoveryAction", "body", *m.RecoveryAction); err != nil {
		return err
	}

	return nil
}

func (m *RecoverAcropolisParams) validateDownloadFileAndFolderParams(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadFileAndFolderParams) { // not required
		return nil
	}

	if m.DownloadFileAndFolderParams != nil {
		if err := m.DownloadFileAndFolderParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadFileAndFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadFileAndFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAcropolisParams) validateRecoverFileAndFolderParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverFileAndFolderParams) { // not required
		return nil
	}

	if m.RecoverFileAndFolderParams != nil {
		if err := m.RecoverFileAndFolderParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverFileAndFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverFileAndFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAcropolisParams) validateRecoverVMParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverVMParams) { // not required
		return nil
	}

	if m.RecoverVMParams != nil {
		if err := m.RecoverVMParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverVmParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverVmParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover acropolis params based on the context it is used
func (m *RecoverAcropolisParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDownloadFileAndFolderParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverFileAndFolderParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverVMParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAcropolisParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverAcropolisParams) contextValidateDownloadFileAndFolderParams(ctx context.Context, formats strfmt.Registry) error {

	if m.DownloadFileAndFolderParams != nil {

		if swag.IsZero(m.DownloadFileAndFolderParams) { // not required
			return nil
		}

		if err := m.DownloadFileAndFolderParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadFileAndFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadFileAndFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAcropolisParams) contextValidateRecoverFileAndFolderParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverFileAndFolderParams != nil {

		if swag.IsZero(m.RecoverFileAndFolderParams) { // not required
			return nil
		}

		if err := m.RecoverFileAndFolderParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverFileAndFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverFileAndFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAcropolisParams) contextValidateRecoverVMParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverVMParams != nil {

		if swag.IsZero(m.RecoverVMParams) { // not required
			return nil
		}

		if err := m.RecoverVMParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverVmParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverVmParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverAcropolisParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAcropolisParams) UnmarshalBinary(b []byte) error {
	var res RecoverAcropolisParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
