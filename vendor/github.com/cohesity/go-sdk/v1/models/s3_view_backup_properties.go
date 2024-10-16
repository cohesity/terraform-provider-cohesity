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

// S3ViewBackupProperties s3 view backup properties
//
// swagger:model S3ViewBackupProperties
type S3ViewBackupProperties struct {

	// Access key for the buckets which will be created for the source initiated
	// jobs. This needs to be passed to Netapp for it to for doing all s3
	// communications.
	AccessKey *string `json:"accessKey,omitempty"`

	// For source initiated backup the target is s3 view. This captures the
	// configuration needed to create the s3 view.
	S3Config *S3BucketConfigProto `json:"s3Config,omitempty"`

	// Secret key for the buckets will be created for the source initiated jobs.
	// This secret key needed to be sent to Netapp for writing data to our S3
	// views.
	SecretKey *string `json:"secretKey,omitempty"`

	// The snapshot prefix which is needed at the time of incremental for
	// backups. This is only needed for case of DP volume.
	SnapshotPrefixName *string `json:"snapshotPrefixName,omitempty"`

	// The id of the S3 view.
	ViewID *int64 `json:"viewId,omitempty"`
}

// Validate validates this s3 view backup properties
func (m *S3ViewBackupProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateS3Config(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3ViewBackupProperties) validateS3Config(formats strfmt.Registry) error {
	if swag.IsZero(m.S3Config) { // not required
		return nil
	}

	if m.S3Config != nil {
		if err := m.S3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3Config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 view backup properties based on the context it is used
func (m *S3ViewBackupProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateS3Config(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3ViewBackupProperties) contextValidateS3Config(ctx context.Context, formats strfmt.Registry) error {

	if m.S3Config != nil {

		if swag.IsZero(m.S3Config) { // not required
			return nil
		}

		if err := m.S3Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3Config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3ViewBackupProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3ViewBackupProperties) UnmarshalBinary(b []byte) error {
	var res S3ViewBackupProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
