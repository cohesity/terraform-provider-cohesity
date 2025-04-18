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

// RecoverFlashbladeParams Recover Flashblade Params.
//
// Specifies the recovery options specific to Flashblade environment.
//
// swagger:model RecoverFlashbladeParams
type RecoverFlashbladeParams struct {

	// Specifies the list of recover Object parameters.
	// Required: true
	Objects []*CommonRecoverObjectSnapshotParams `json:"objects"`

	// Specifies the type of recover action to be performed.
	// Required: true
	// Enum: ["RecoverNasVolume","RecoverFiles"]
	RecoveryAction *string `json:"recoveryAction"`

	// Specifies the parameters to download files and folders.
	DownloadFileAndFolderParams *CommonDownloadFileAndFolderParams `json:"downloadFileAndFolderParams,omitempty"`

	// Specifies the parameters to recover files.
	RecoverFileAndFolderParams *RecoverFlashbladeFilesParams `json:"recoverFileAndFolderParams,omitempty"`

	// Specifies the parameters to recover Nas Volumes.
	RecoverNasVolumeParams *RecoverFlashbladeNasVolumeParams `json:"recoverNasVolumeParams,omitempty"`
}

// Validate validates this recover flashblade params
func (m *RecoverFlashbladeParams) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRecoverNasVolumeParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverFlashbladeParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
		return err
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

var recoverFlashbladeParamsTypeRecoveryActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RecoverNasVolume","RecoverFiles"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverFlashbladeParamsTypeRecoveryActionPropEnum = append(recoverFlashbladeParamsTypeRecoveryActionPropEnum, v)
	}
}

const (

	// RecoverFlashbladeParamsRecoveryActionRecoverNasVolume captures enum value "RecoverNasVolume"
	RecoverFlashbladeParamsRecoveryActionRecoverNasVolume string = "RecoverNasVolume"

	// RecoverFlashbladeParamsRecoveryActionRecoverFiles captures enum value "RecoverFiles"
	RecoverFlashbladeParamsRecoveryActionRecoverFiles string = "RecoverFiles"
)

// prop value enum
func (m *RecoverFlashbladeParams) validateRecoveryActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverFlashbladeParamsTypeRecoveryActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverFlashbladeParams) validateRecoveryAction(formats strfmt.Registry) error {

	if err := validate.Required("recoveryAction", "body", m.RecoveryAction); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecoveryActionEnum("recoveryAction", "body", *m.RecoveryAction); err != nil {
		return err
	}

	return nil
}

func (m *RecoverFlashbladeParams) validateDownloadFileAndFolderParams(formats strfmt.Registry) error {
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

func (m *RecoverFlashbladeParams) validateRecoverFileAndFolderParams(formats strfmt.Registry) error {
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

func (m *RecoverFlashbladeParams) validateRecoverNasVolumeParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverNasVolumeParams) { // not required
		return nil
	}

	if m.RecoverNasVolumeParams != nil {
		if err := m.RecoverNasVolumeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverNasVolumeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverNasVolumeParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover flashblade params based on the context it is used
func (m *RecoverFlashbladeParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateRecoverNasVolumeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverFlashbladeParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverFlashbladeParams) contextValidateDownloadFileAndFolderParams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverFlashbladeParams) contextValidateRecoverFileAndFolderParams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverFlashbladeParams) contextValidateRecoverNasVolumeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverNasVolumeParams != nil {

		if swag.IsZero(m.RecoverNasVolumeParams) { // not required
			return nil
		}

		if err := m.RecoverNasVolumeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverNasVolumeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverNasVolumeParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverFlashbladeParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverFlashbladeParams) UnmarshalBinary(b []byte) error {
	var res RecoverFlashbladeParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
