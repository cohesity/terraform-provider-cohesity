// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetrieveArchiveTaskStateProtoDownloadFilesInfo Information required for Icebox when downloading files from an archived
// entity. This proto is specifically just for the current temporary solution
// for downloading a single from an archive, where we let icebox download the
// file for us. In the future once the new Yoda APIs for downloading files
// from archive stub views are ready, we will just discard this proto and
// make field 20 reserved.
//
// swagger:model RetrieveArchiveTaskStateProto_DownloadFilesInfo
type RetrieveArchiveTaskStateProtoDownloadFilesInfo struct {

	// The file paths to be retrieved from the archive.
	// For example, if the file paths are /foo/bar and /foo2/bar
	// and target_dir_path is "downloaded_files". The retrieved files will have
	// the full path structure maintained in target_dir_path such as
	// downloaded_files/foo/bar and /downloaded_files/foo2/bar.
	FilePathVec []string `json:"filePathVec"`

	// Instance from which the requested files will be restored.
	MagnetoInstanceID *MagnetoInstanceID `json:"magnetoInstanceId,omitempty"`

	// Object from which the requested files will be restored.
	ObjectID *MagnetoObjectID `json:"objectId,omitempty"`

	// Ask Icebox to only create the stub view and skip restoring files in
	// the stub view.
	SkipRestoreInStubView *bool `json:"skipRestoreInStubView,omitempty"`

	// Path to the directory under which the downloaded files will be placed.
	TargetDirPath *string `json:"targetDirPath,omitempty"`

	// Target view name where the downloaded files will be placed.
	TargetViewName *string `json:"targetViewName,omitempty"`
}

// Validate validates this retrieve archive task state proto download files info
func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMagnetoInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) validateMagnetoInstanceID(formats strfmt.Registry) error {
	if swag.IsZero(m.MagnetoInstanceID) { // not required
		return nil
	}

	if m.MagnetoInstanceID != nil {
		if err := m.MagnetoInstanceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("magnetoInstanceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("magnetoInstanceId")
			}
			return err
		}
	}

	return nil
}

func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) validateObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectID) { // not required
		return nil
	}

	if m.ObjectID != nil {
		if err := m.ObjectID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this retrieve archive task state proto download files info based on the context it is used
func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMagnetoInstanceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) contextValidateMagnetoInstanceID(ctx context.Context, formats strfmt.Registry) error {

	if m.MagnetoInstanceID != nil {

		if swag.IsZero(m.MagnetoInstanceID) { // not required
			return nil
		}

		if err := m.MagnetoInstanceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("magnetoInstanceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("magnetoInstanceId")
			}
			return err
		}
	}

	return nil
}

func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) contextValidateObjectID(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectID != nil {

		if swag.IsZero(m.ObjectID) { // not required
			return nil
		}

		if err := m.ObjectID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrieveArchiveTaskStateProtoDownloadFilesInfo) UnmarshalBinary(b []byte) error {
	var res RetrieveArchiveTaskStateProtoDownloadFilesInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
