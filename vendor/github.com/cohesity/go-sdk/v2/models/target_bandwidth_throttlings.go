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

// TargetBandwidthThrottlings Target Bandwidth Throttlings
//
// Specifies the bandwidth throttling setting of the External Target.
//
// swagger:model TargetBandwidthThrottlings
type TargetBandwidthThrottlings struct {

	// download
	Download *BandwidthThrottling `json:"download,omitempty"`

	// upload
	Upload *BandwidthThrottling `json:"upload,omitempty"`
}

// Validate validates this target bandwidth throttlings
func (m *TargetBandwidthThrottlings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetBandwidthThrottlings) validateDownload(formats strfmt.Registry) error {
	if swag.IsZero(m.Download) { // not required
		return nil
	}

	if m.Download != nil {
		if err := m.Download.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("download")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("download")
			}
			return err
		}
	}

	return nil
}

func (m *TargetBandwidthThrottlings) validateUpload(formats strfmt.Registry) error {
	if swag.IsZero(m.Upload) { // not required
		return nil
	}

	if m.Upload != nil {
		if err := m.Upload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this target bandwidth throttlings based on the context it is used
func (m *TargetBandwidthThrottlings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDownload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetBandwidthThrottlings) contextValidateDownload(ctx context.Context, formats strfmt.Registry) error {

	if m.Download != nil {

		if swag.IsZero(m.Download) { // not required
			return nil
		}

		if err := m.Download.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("download")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("download")
			}
			return err
		}
	}

	return nil
}

func (m *TargetBandwidthThrottlings) contextValidateUpload(ctx context.Context, formats strfmt.Registry) error {

	if m.Upload != nil {

		if swag.IsZero(m.Upload) { // not required
			return nil
		}

		if err := m.Upload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TargetBandwidthThrottlings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetBandwidthThrottlings) UnmarshalBinary(b []byte) error {
	var res TargetBandwidthThrottlings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
