// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AntivirusScanConfig Antivirus scan config
//
// Specifies the antivirus scan config settings for this View.
//
// swagger:model AntivirusScanConfig
type AntivirusScanConfig struct {

	// Specifies whether block access to the file when antivirus scan fails.
	BlockAccessOnScanFailure *bool `json:"blockAccessOnScanFailure,omitempty"`

	// Specifies whether the antivirus service is enabled or not.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Specifies maximum file size that will be sent to antivirus server for
	// scanning. if greater than zero, the file size that exceeds this size
	// would be skipped from virus scan.
	MaximumScanFileSize *int64 `json:"maximumScanFileSize,omitempty"`

	// Specifies whether to scan a SMB file or S3 object before it is opened/GET.
	ScanOnAccess *bool `json:"scanOnAccess,omitempty"`

	// Specifies whether to scan a S3 object after it is PUT.
	ScanOnPut *bool `json:"scanOnPut,omitempty"`

	// Specifies whether to scan a SMB file when it is closed after modify.
	ScanOnClose *bool `json:"scanOnClose,omitempty"`

	// Specifies the maximum amount of time that a scan can take before timing
	// out.
	// Required: true
	ScanTimeoutUsecs *int32 `json:"scanTimeoutUsecs"`

	// For S3 usage only. Prefixes that meets these filter criteria will be sent to
	// antivirus server for the scan. This filter is case-sensitive.
	PrefixScanFilter *FileExtensionFilter `json:"prefixScanFilter,omitempty"`

	// For S3 usage only. If set, only objects with these taggings will be
	// scanned or not scanned, depending on the mode.
	S3TaggingFilter *S3TaggingFilter `json:"s3TaggingFilter,omitempty"`

	// Files extension that meets these filter criteria will be sent to
	// antivirus server for the scan. This filter is case-insensitive.
	ScanFilter *FileExtensionFilter `json:"scanFilter,omitempty"`
}

// Validate validates this antivirus scan config
func (m *AntivirusScanConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScanTimeoutUsecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixScanFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3TaggingFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntivirusScanConfig) validateScanTimeoutUsecs(formats strfmt.Registry) error {

	if err := validate.Required("scanTimeoutUsecs", "body", m.ScanTimeoutUsecs); err != nil {
		return err
	}

	return nil
}

func (m *AntivirusScanConfig) validatePrefixScanFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.PrefixScanFilter) { // not required
		return nil
	}

	if m.PrefixScanFilter != nil {
		if err := m.PrefixScanFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prefixScanFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prefixScanFilter")
			}
			return err
		}
	}

	return nil
}

func (m *AntivirusScanConfig) validateS3TaggingFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.S3TaggingFilter) { // not required
		return nil
	}

	if m.S3TaggingFilter != nil {
		if err := m.S3TaggingFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3TaggingFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3TaggingFilter")
			}
			return err
		}
	}

	return nil
}

func (m *AntivirusScanConfig) validateScanFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanFilter) { // not required
		return nil
	}

	if m.ScanFilter != nil {
		if err := m.ScanFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanFilter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this antivirus scan config based on the context it is used
func (m *AntivirusScanConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrefixScanFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3TaggingFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScanFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntivirusScanConfig) contextValidatePrefixScanFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.PrefixScanFilter != nil {

		if swag.IsZero(m.PrefixScanFilter) { // not required
			return nil
		}

		if err := m.PrefixScanFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prefixScanFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prefixScanFilter")
			}
			return err
		}
	}

	return nil
}

func (m *AntivirusScanConfig) contextValidateS3TaggingFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.S3TaggingFilter != nil {

		if swag.IsZero(m.S3TaggingFilter) { // not required
			return nil
		}

		if err := m.S3TaggingFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3TaggingFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3TaggingFilter")
			}
			return err
		}
	}

	return nil
}

func (m *AntivirusScanConfig) contextValidateScanFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.ScanFilter != nil {

		if swag.IsZero(m.ScanFilter) { // not required
			return nil
		}

		if err := m.ScanFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntivirusScanConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntivirusScanConfig) UnmarshalBinary(b []byte) error {
	var res AntivirusScanConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
