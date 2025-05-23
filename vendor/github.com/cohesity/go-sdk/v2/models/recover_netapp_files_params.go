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

// RecoverNetappFilesParams Recover Netapp files Params.
//
// Specifies the parameters to recover Netapp files.
//
// swagger:model RecoverNetappFilesParams
type RecoverNetappFilesParams struct {

	// Specifies the list of info about the netapp files and folders to be recovered.
	// Required: true
	FilesAndFolders []*NetappRecoverFileAndFolderInfo `json:"filesAndFolders"`

	// Specifies if the snapshot trying to recover is from a source initiated protection.
	IsFromSourceInitiatedProtection *bool `json:"isFromSourceInitiatedProtection,omitempty"`

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kElastifile","kFlashBlade","kGenericNas","kGPFS","kIsilon","kNetapp"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the params for an Elastifile recovery target.
	ElastifileTargetParams *RecoverOtherNasToElastifileFilesTargetParams `json:"elastifileTargetParams,omitempty"`

	// Specifies the params for a Flashblade recovery target.
	FlashbladeTargetParams *RecoverOtherNasToFlashbladeFilesTargetParams `json:"flashbladeTargetParams,omitempty"`

	// Specifies the params for a generic NAS recovery target.
	GenericNasTargetParams *RecoverOtherNasToGenericNasFilesTargetParams `json:"genericNasTargetParams,omitempty"`

	// Specifies the params for a GPFS recovery target.
	GpfsTargetParams *RecoverOtherNasToGpfsFilesTargetParams `json:"gpfsTargetParams,omitempty"`

	// Specifies the params for an Isilon recovery target.
	IsilonTargetParams *RecoverOtherNasToIsilonFilesTargetParams `json:"isilonTargetParams,omitempty"`

	// Specifies the params for a Netapp recovery target.
	NetappTargetParams *RecoverNetappToNetappFilesTargetParams `json:"netappTargetParams,omitempty"`
}

// Validate validates this recover netapp files params
func (m *RecoverNetappFilesParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilesAndFolders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElastifileTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlashbladeTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenericNasTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpfsTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsilonTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetappTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverNetappFilesParams) validateFilesAndFolders(formats strfmt.Registry) error {

	if err := validate.Required("filesAndFolders", "body", m.FilesAndFolders); err != nil {
		return err
	}

	for i := 0; i < len(m.FilesAndFolders); i++ {
		if swag.IsZero(m.FilesAndFolders[i]) { // not required
			continue
		}

		if m.FilesAndFolders[i] != nil {
			if err := m.FilesAndFolders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesAndFolders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filesAndFolders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var recoverNetappFilesParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kElastifile","kFlashBlade","kGenericNas","kGPFS","kIsilon","kNetapp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverNetappFilesParamsTypeTargetEnvironmentPropEnum = append(recoverNetappFilesParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverNetappFilesParamsTargetEnvironmentKElastifile captures enum value "kElastifile"
	RecoverNetappFilesParamsTargetEnvironmentKElastifile string = "kElastifile"

	// RecoverNetappFilesParamsTargetEnvironmentKFlashBlade captures enum value "kFlashBlade"
	RecoverNetappFilesParamsTargetEnvironmentKFlashBlade string = "kFlashBlade"

	// RecoverNetappFilesParamsTargetEnvironmentKGenericNas captures enum value "kGenericNas"
	RecoverNetappFilesParamsTargetEnvironmentKGenericNas string = "kGenericNas"

	// RecoverNetappFilesParamsTargetEnvironmentKGPFS captures enum value "kGPFS"
	RecoverNetappFilesParamsTargetEnvironmentKGPFS string = "kGPFS"

	// RecoverNetappFilesParamsTargetEnvironmentKIsilon captures enum value "kIsilon"
	RecoverNetappFilesParamsTargetEnvironmentKIsilon string = "kIsilon"

	// RecoverNetappFilesParamsTargetEnvironmentKNetapp captures enum value "kNetapp"
	RecoverNetappFilesParamsTargetEnvironmentKNetapp string = "kNetapp"
)

// prop value enum
func (m *RecoverNetappFilesParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverNetappFilesParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverNetappFilesParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverNetappFilesParams) validateElastifileTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ElastifileTargetParams) { // not required
		return nil
	}

	if m.ElastifileTargetParams != nil {
		if err := m.ElastifileTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elastifileTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elastifileTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) validateFlashbladeTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.FlashbladeTargetParams) { // not required
		return nil
	}

	if m.FlashbladeTargetParams != nil {
		if err := m.FlashbladeTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flashbladeTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flashbladeTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) validateGenericNasTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.GenericNasTargetParams) { // not required
		return nil
	}

	if m.GenericNasTargetParams != nil {
		if err := m.GenericNasTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("genericNasTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("genericNasTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) validateGpfsTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.GpfsTargetParams) { // not required
		return nil
	}

	if m.GpfsTargetParams != nil {
		if err := m.GpfsTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpfsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gpfsTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) validateIsilonTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.IsilonTargetParams) { // not required
		return nil
	}

	if m.IsilonTargetParams != nil {
		if err := m.IsilonTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isilonTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isilonTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) validateNetappTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.NetappTargetParams) { // not required
		return nil
	}

	if m.NetappTargetParams != nil {
		if err := m.NetappTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netappTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netappTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover netapp files params based on the context it is used
func (m *RecoverNetappFilesParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilesAndFolders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElastifileTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlashbladeTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGenericNasTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpfsTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsilonTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetappTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverNetappFilesParams) contextValidateFilesAndFolders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FilesAndFolders); i++ {

		if m.FilesAndFolders[i] != nil {

			if swag.IsZero(m.FilesAndFolders[i]) { // not required
				return nil
			}

			if err := m.FilesAndFolders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesAndFolders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filesAndFolders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverNetappFilesParams) contextValidateElastifileTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ElastifileTargetParams != nil {

		if swag.IsZero(m.ElastifileTargetParams) { // not required
			return nil
		}

		if err := m.ElastifileTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elastifileTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elastifileTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) contextValidateFlashbladeTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.FlashbladeTargetParams != nil {

		if swag.IsZero(m.FlashbladeTargetParams) { // not required
			return nil
		}

		if err := m.FlashbladeTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flashbladeTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flashbladeTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) contextValidateGenericNasTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.GenericNasTargetParams != nil {

		if swag.IsZero(m.GenericNasTargetParams) { // not required
			return nil
		}

		if err := m.GenericNasTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("genericNasTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("genericNasTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) contextValidateGpfsTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.GpfsTargetParams != nil {

		if swag.IsZero(m.GpfsTargetParams) { // not required
			return nil
		}

		if err := m.GpfsTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpfsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gpfsTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) contextValidateIsilonTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IsilonTargetParams != nil {

		if swag.IsZero(m.IsilonTargetParams) { // not required
			return nil
		}

		if err := m.IsilonTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isilonTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isilonTargetParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverNetappFilesParams) contextValidateNetappTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.NetappTargetParams != nil {

		if swag.IsZero(m.NetappTargetParams) { // not required
			return nil
		}

		if err := m.NetappTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netappTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netappTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverNetappFilesParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverNetappFilesParams) UnmarshalBinary(b []byte) error {
	var res RecoverNetappFilesParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
