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

// RecoverAwsFileAndFolderParams Recover AWS File And Folder Params
//
// Specifies the parameters to recover files and folders.
//
// swagger:model RecoverAwsFileAndFolderParams
type RecoverAwsFileAndFolderParams struct {

	// Specifies the info about the files and folders to be recovered.
	// Required: true
	FilesAndFolders []*CommonRecoverFileAndFolderInfo `json:"filesAndFolders"`

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kAWS"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the parameters to recover to an AWS target.
	AwsTargetParams *AwsTargetParamsForRecoverFileAndFolder `json:"awsTargetParams,omitempty"`
}

// Validate validates this recover aws file and folder params
func (m *RecoverAwsFileAndFolderParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilesAndFolders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAwsFileAndFolderParams) validateFilesAndFolders(formats strfmt.Registry) error {

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

var recoverAwsFileAndFolderParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAWS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverAwsFileAndFolderParamsTypeTargetEnvironmentPropEnum = append(recoverAwsFileAndFolderParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverAwsFileAndFolderParamsTargetEnvironmentKAWS captures enum value "kAWS"
	RecoverAwsFileAndFolderParamsTargetEnvironmentKAWS string = "kAWS"
)

// prop value enum
func (m *RecoverAwsFileAndFolderParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverAwsFileAndFolderParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverAwsFileAndFolderParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverAwsFileAndFolderParams) validateAwsTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsTargetParams) { // not required
		return nil
	}

	if m.AwsTargetParams != nil {
		if err := m.AwsTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover aws file and folder params based on the context it is used
func (m *RecoverAwsFileAndFolderParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilesAndFolders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAwsFileAndFolderParams) contextValidateFilesAndFolders(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverAwsFileAndFolderParams) contextValidateAwsTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsTargetParams != nil {

		if swag.IsZero(m.AwsTargetParams) { // not required
			return nil
		}

		if err := m.AwsTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverAwsFileAndFolderParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAwsFileAndFolderParams) UnmarshalBinary(b []byte) error {
	var res RecoverAwsFileAndFolderParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
