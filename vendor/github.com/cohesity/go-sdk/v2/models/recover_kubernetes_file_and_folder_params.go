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

// RecoverKubernetesFileAndFolderParams Recover Kubernetes File And Folder Params
//
// Specifies the parameters to recover files and folders.
//
// swagger:model RecoverKubernetesFileAndFolderParams
type RecoverKubernetesFileAndFolderParams struct {

	// Specifies the information about the files and folders to be recovered.
	// Required: true
	FilesAndFolders []*CommonRecoverFileAndFolderInfo `json:"filesAndFolders"`

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kKubernetes"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the parameters to recover to a Kubernetes target.
	KubernetesTargetParams *KubernetesTargetParamsForRecoverFileAndFolder `json:"kubernetesTargetParams,omitempty"`
}

// Validate validates this recover kubernetes file and folder params
func (m *RecoverKubernetesFileAndFolderParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilesAndFolders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverKubernetesFileAndFolderParams) validateFilesAndFolders(formats strfmt.Registry) error {

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

var recoverKubernetesFileAndFolderParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kKubernetes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverKubernetesFileAndFolderParamsTypeTargetEnvironmentPropEnum = append(recoverKubernetesFileAndFolderParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverKubernetesFileAndFolderParamsTargetEnvironmentKKubernetes captures enum value "kKubernetes"
	RecoverKubernetesFileAndFolderParamsTargetEnvironmentKKubernetes string = "kKubernetes"
)

// prop value enum
func (m *RecoverKubernetesFileAndFolderParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverKubernetesFileAndFolderParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverKubernetesFileAndFolderParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverKubernetesFileAndFolderParams) validateKubernetesTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.KubernetesTargetParams) { // not required
		return nil
	}

	if m.KubernetesTargetParams != nil {
		if err := m.KubernetesTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover kubernetes file and folder params based on the context it is used
func (m *RecoverKubernetesFileAndFolderParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilesAndFolders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetesTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverKubernetesFileAndFolderParams) contextValidateFilesAndFolders(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverKubernetesFileAndFolderParams) contextValidateKubernetesTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.KubernetesTargetParams != nil {

		if swag.IsZero(m.KubernetesTargetParams) { // not required
			return nil
		}

		if err := m.KubernetesTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverKubernetesFileAndFolderParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverKubernetesFileAndFolderParams) UnmarshalBinary(b []byte) error {
	var res RecoverKubernetesFileAndFolderParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
