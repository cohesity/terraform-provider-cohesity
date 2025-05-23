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

// RecoverVMwareFileAndFolderParams Recover VMware File And Folder Params
//
// Specifies the parameters to recover files and folders.
//
// swagger:model RecoverVMwareFileAndFolderParams
type RecoverVMwareFileAndFolderParams struct {

	// Specifies the info about the files and folders to be recovered.
	// Required: true
	FilesAndFolders []*CommonRecoverFileAndFolderInfo `json:"filesAndFolders"`

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kVMware"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery.
	// Pattern: ^\d+:\d+:\d+$
	ParentRecoveryID *string `json:"parentRecoveryId,omitempty"`

	// Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot.
	// Enum: ["kStandard","kExpeditedNoPCU","kExpeditedWithPCU"]
	GlacierRetrievalType *string `json:"glacierRetrievalType,omitempty"`

	// Specifies the parameters to recover to a VMware target.
	VmwareTargetParams *VmwareTargetParamsForRecoverFileAndFolder `json:"vmwareTargetParams,omitempty"`
}

// Validate validates this recover v mware file and folder params
func (m *RecoverVMwareFileAndFolderParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilesAndFolders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentRecoveryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlacierRetrievalType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmwareTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVMwareFileAndFolderParams) validateFilesAndFolders(formats strfmt.Registry) error {

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

var recoverVMwareFileAndFolderParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverVMwareFileAndFolderParamsTypeTargetEnvironmentPropEnum = append(recoverVMwareFileAndFolderParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverVMwareFileAndFolderParamsTargetEnvironmentKVMware captures enum value "kVMware"
	RecoverVMwareFileAndFolderParamsTargetEnvironmentKVMware string = "kVMware"
)

// prop value enum
func (m *RecoverVMwareFileAndFolderParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverVMwareFileAndFolderParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverVMwareFileAndFolderParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverVMwareFileAndFolderParams) validateParentRecoveryID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentRecoveryID) { // not required
		return nil
	}

	if err := validate.Pattern("parentRecoveryId", "body", *m.ParentRecoveryID, `^\d+:\d+:\d+$`); err != nil {
		return err
	}

	return nil
}

var recoverVMwareFileAndFolderParamsTypeGlacierRetrievalTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kStandard","kExpeditedNoPCU","kExpeditedWithPCU"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverVMwareFileAndFolderParamsTypeGlacierRetrievalTypePropEnum = append(recoverVMwareFileAndFolderParamsTypeGlacierRetrievalTypePropEnum, v)
	}
}

const (

	// RecoverVMwareFileAndFolderParamsGlacierRetrievalTypeKStandard captures enum value "kStandard"
	RecoverVMwareFileAndFolderParamsGlacierRetrievalTypeKStandard string = "kStandard"

	// RecoverVMwareFileAndFolderParamsGlacierRetrievalTypeKExpeditedNoPCU captures enum value "kExpeditedNoPCU"
	RecoverVMwareFileAndFolderParamsGlacierRetrievalTypeKExpeditedNoPCU string = "kExpeditedNoPCU"

	// RecoverVMwareFileAndFolderParamsGlacierRetrievalTypeKExpeditedWithPCU captures enum value "kExpeditedWithPCU"
	RecoverVMwareFileAndFolderParamsGlacierRetrievalTypeKExpeditedWithPCU string = "kExpeditedWithPCU"
)

// prop value enum
func (m *RecoverVMwareFileAndFolderParams) validateGlacierRetrievalTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverVMwareFileAndFolderParamsTypeGlacierRetrievalTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverVMwareFileAndFolderParams) validateGlacierRetrievalType(formats strfmt.Registry) error {
	if swag.IsZero(m.GlacierRetrievalType) { // not required
		return nil
	}

	// value enum
	if err := m.validateGlacierRetrievalTypeEnum("glacierRetrievalType", "body", *m.GlacierRetrievalType); err != nil {
		return err
	}

	return nil
}

func (m *RecoverVMwareFileAndFolderParams) validateVmwareTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.VmwareTargetParams) { // not required
		return nil
	}

	if m.VmwareTargetParams != nil {
		if err := m.VmwareTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover v mware file and folder params based on the context it is used
func (m *RecoverVMwareFileAndFolderParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilesAndFolders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmwareTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVMwareFileAndFolderParams) contextValidateFilesAndFolders(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverVMwareFileAndFolderParams) contextValidateVmwareTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VmwareTargetParams != nil {

		if swag.IsZero(m.VmwareTargetParams) { // not required
			return nil
		}

		if err := m.VmwareTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverVMwareFileAndFolderParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverVMwareFileAndFolderParams) UnmarshalBinary(b []byte) error {
	var res RecoverVMwareFileAndFolderParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
