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

// FileSearchResult File/Folder Search Result.
//
// Specifies details about the found file or folder.
//
// swagger:model FileSearchResult
type FileSearchResult struct {

	// Specifies the metadata about the AD objects.
	AdObjectMetaData *AdObjectMetaData `json:"adObjectMetaData,omitempty"`

	// Specifies the inferred document type.
	DocumentType *string `json:"documentType,omitempty"`

	// Specifies the metadata about the Outlook Emails.
	EmailMetaData *EmailMetaData `json:"emailMetaData,omitempty"`

	// Array of File Versions.
	//
	// Specifies the different snapshot versions of a file or folder
	// that were captured at different times.
	FileVersions []*FileVersion `json:"fileVersions"`

	// Specifies the name of the found file or folder.
	Filename *string `json:"filename,omitempty"`

	// Specifies if the found item is a folder.
	// If true, the found item is a folder.
	IsFolder *bool `json:"isFolder,omitempty"`

	// Specifies the Job id for the Protection Job that is currently
	// associated with object that contains the backed up file or folder.
	// If the file or folder was backed up on current Cohesity Cluster,
	// this field contains the id for the Job that captured the object
	// that contains the file or folder.
	// If the file or folder was backed up on a Primary Cluster and
	// replicated to this Cohesity Cluster, a new Inactive Job is created,
	// the object that contains the file or folder is now associated
	// with new Inactive Job, and this field
	// contains the id of the new Inactive Job.
	JobID *int64 `json:"jobId,omitempty"`

	// Specifies the universal id of the Protection Job that backed up
	// the object that contains the file or folder.
	JobUID struct {
		UniversalID
	} `json:"jobUid,omitempty"`

	// Specifies the metadata about the OneDrive documents.
	OneDriveDocumentMetadata *OneDriveDocumentMetadata `json:"oneDriveDocumentMetadata,omitempty"`

	// Specifies the Protection Source to which the file or folder belongs.
	ProtectionSource *ProtectionSource `json:"protectionSource,omitempty"`

	// Specifies the id of the top-level registered source (such as a
	// vCenter Server) where the source object that contains the
	// the file or folder is stored.
	RegisteredSourceID *int64 `json:"registeredSourceId,omitempty"`

	// Specifies the metadata about the Sharepoint documents.
	SharepointDocumentMetadata *SharepointDocumentMetadata `json:"sharepointDocumentMetadata,omitempty"`

	// Snapshot tags present on this document
	SnapshotTags []string `json:"snapshotTags"`

	// Specifies the source id of the object that contains the file or folder.
	SourceID *int64 `json:"sourceId,omitempty"`

	// Tags present on this document.
	Tags []string `json:"tags"`

	// Mapping from snapshot tags to.
	TagsToSnapshotsMap map[string][]int64 `json:"tagsToSnapshotsMap,omitempty"`

	// Specifies the type of the file document such as KDirectory, kFile, etc.
	// Enum: ["kDirectory","kFile","kEmail","kSymlink"]
	Type *string `json:"type,omitempty"`

	// Specifies the id of the Domain (View Box) where the source object that
	// contains the file or folder is stored.
	ViewBoxID *int64 `json:"viewBoxId,omitempty"`
}

// Validate validates this file search result
func (m *FileSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdObjectMetaData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailMetaData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneDriveDocumentMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharepointDocumentMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileSearchResult) validateAdObjectMetaData(formats strfmt.Registry) error {
	if swag.IsZero(m.AdObjectMetaData) { // not required
		return nil
	}

	if m.AdObjectMetaData != nil {
		if err := m.AdObjectMetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adObjectMetaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adObjectMetaData")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) validateEmailMetaData(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailMetaData) { // not required
		return nil
	}

	if m.EmailMetaData != nil {
		if err := m.EmailMetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailMetaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailMetaData")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) validateFileVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.FileVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.FileVersions); i++ {
		if swag.IsZero(m.FileVersions[i]) { // not required
			continue
		}

		if m.FileVersions[i] != nil {
			if err := m.FileVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileVersions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FileSearchResult) validateJobUID(formats strfmt.Registry) error {
	if swag.IsZero(m.JobUID) { // not required
		return nil
	}

	return nil
}

func (m *FileSearchResult) validateOneDriveDocumentMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.OneDriveDocumentMetadata) { // not required
		return nil
	}

	if m.OneDriveDocumentMetadata != nil {
		if err := m.OneDriveDocumentMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oneDriveDocumentMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oneDriveDocumentMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) validateProtectionSource(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionSource) { // not required
		return nil
	}

	if m.ProtectionSource != nil {
		if err := m.ProtectionSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectionSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectionSource")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) validateSharepointDocumentMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.SharepointDocumentMetadata) { // not required
		return nil
	}

	if m.SharepointDocumentMetadata != nil {
		if err := m.SharepointDocumentMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharepointDocumentMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharepointDocumentMetadata")
			}
			return err
		}
	}

	return nil
}

var fileSearchResultTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kDirectory","kFile","kEmail","kSymlink"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileSearchResultTypeTypePropEnum = append(fileSearchResultTypeTypePropEnum, v)
	}
}

const (

	// FileSearchResultTypeKDirectory captures enum value "kDirectory"
	FileSearchResultTypeKDirectory string = "kDirectory"

	// FileSearchResultTypeKFile captures enum value "kFile"
	FileSearchResultTypeKFile string = "kFile"

	// FileSearchResultTypeKEmail captures enum value "kEmail"
	FileSearchResultTypeKEmail string = "kEmail"

	// FileSearchResultTypeKSymlink captures enum value "kSymlink"
	FileSearchResultTypeKSymlink string = "kSymlink"
)

// prop value enum
func (m *FileSearchResult) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileSearchResultTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileSearchResult) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this file search result based on the context it is used
func (m *FileSearchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdObjectMetaData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailMetaData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJobUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOneDriveDocumentMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSharepointDocumentMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileSearchResult) contextValidateAdObjectMetaData(ctx context.Context, formats strfmt.Registry) error {

	if m.AdObjectMetaData != nil {

		if swag.IsZero(m.AdObjectMetaData) { // not required
			return nil
		}

		if err := m.AdObjectMetaData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adObjectMetaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adObjectMetaData")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) contextValidateEmailMetaData(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailMetaData != nil {

		if swag.IsZero(m.EmailMetaData) { // not required
			return nil
		}

		if err := m.EmailMetaData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailMetaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailMetaData")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) contextValidateFileVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileVersions); i++ {

		if m.FileVersions[i] != nil {

			if swag.IsZero(m.FileVersions[i]) { // not required
				return nil
			}

			if err := m.FileVersions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileVersions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FileSearchResult) contextValidateJobUID(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *FileSearchResult) contextValidateOneDriveDocumentMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.OneDriveDocumentMetadata != nil {

		if swag.IsZero(m.OneDriveDocumentMetadata) { // not required
			return nil
		}

		if err := m.OneDriveDocumentMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oneDriveDocumentMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oneDriveDocumentMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) contextValidateProtectionSource(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionSource != nil {

		if swag.IsZero(m.ProtectionSource) { // not required
			return nil
		}

		if err := m.ProtectionSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectionSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectionSource")
			}
			return err
		}
	}

	return nil
}

func (m *FileSearchResult) contextValidateSharepointDocumentMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.SharepointDocumentMetadata != nil {

		if swag.IsZero(m.SharepointDocumentMetadata) { // not required
			return nil
		}

		if err := m.SharepointDocumentMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharepointDocumentMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharepointDocumentMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileSearchResult) UnmarshalBinary(b []byte) error {
	var res FileSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
