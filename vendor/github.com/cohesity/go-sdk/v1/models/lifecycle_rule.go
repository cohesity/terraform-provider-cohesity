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

// LifecycleRule Proto to define Lifecycle configuration Rule.
//
// swagger:model LifecycleRule
type LifecycleRule struct {

	// Specifies the days since the initiation of an incomplete multipart upload
	// before permanently removing all parts of the upload.
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUploadAction `json:"abortIncompleteMultipartUpload,omitempty"`

	// Specifies the expiration for the lifecycle of the object in the form of
	// date, days and, whether the object has a delete marker.
	Expiration *ExpirationAction `json:"expiration,omitempty"`

	// The Filter is used to identify objects that a Lifecycle Rule applies to.
	Filter *LifecycleRuleFilter `json:"filter,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255
	// characters.
	ID *string `json:"id,omitempty"`

	// Specifies when noncurrent object versions expire. Upon expiration,
	// noncurrent object versions are permanently deleted. The action can be
	// specified only in versioning enabled of suspended buckets.
	NoncurrentVersionExpiration *NoncurrentVersionExpirationAction `json:"noncurrentVersionExpiration,omitempty"`

	// The prefix is used to identify objects that a lifecycle rule applies to.
	Prefix *string `json:"prefix,omitempty"`

	// If set, the rule is currently being applied. If 'Disabled', the rule is
	// not currently being applied.
	Status *bool `json:"status,omitempty"`
}

// Validate validates this lifecycle rule
func (m *LifecycleRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbortIncompleteMultipartUpload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoncurrentVersionExpiration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LifecycleRule) validateAbortIncompleteMultipartUpload(formats strfmt.Registry) error {
	if swag.IsZero(m.AbortIncompleteMultipartUpload) { // not required
		return nil
	}

	if m.AbortIncompleteMultipartUpload != nil {
		if err := m.AbortIncompleteMultipartUpload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abortIncompleteMultipartUpload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abortIncompleteMultipartUpload")
			}
			return err
		}
	}

	return nil
}

func (m *LifecycleRule) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if m.Expiration != nil {
		if err := m.Expiration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expiration")
			}
			return err
		}
	}

	return nil
}

func (m *LifecycleRule) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *LifecycleRule) validateNoncurrentVersionExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.NoncurrentVersionExpiration) { // not required
		return nil
	}

	if m.NoncurrentVersionExpiration != nil {
		if err := m.NoncurrentVersionExpiration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("noncurrentVersionExpiration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("noncurrentVersionExpiration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this lifecycle rule based on the context it is used
func (m *LifecycleRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAbortIncompleteMultipartUpload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpiration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNoncurrentVersionExpiration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LifecycleRule) contextValidateAbortIncompleteMultipartUpload(ctx context.Context, formats strfmt.Registry) error {

	if m.AbortIncompleteMultipartUpload != nil {

		if swag.IsZero(m.AbortIncompleteMultipartUpload) { // not required
			return nil
		}

		if err := m.AbortIncompleteMultipartUpload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abortIncompleteMultipartUpload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("abortIncompleteMultipartUpload")
			}
			return err
		}
	}

	return nil
}

func (m *LifecycleRule) contextValidateExpiration(ctx context.Context, formats strfmt.Registry) error {

	if m.Expiration != nil {

		if swag.IsZero(m.Expiration) { // not required
			return nil
		}

		if err := m.Expiration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expiration")
			}
			return err
		}
	}

	return nil
}

func (m *LifecycleRule) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.Filter != nil {

		if swag.IsZero(m.Filter) { // not required
			return nil
		}

		if err := m.Filter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *LifecycleRule) contextValidateNoncurrentVersionExpiration(ctx context.Context, formats strfmt.Registry) error {

	if m.NoncurrentVersionExpiration != nil {

		if swag.IsZero(m.NoncurrentVersionExpiration) { // not required
			return nil
		}

		if err := m.NoncurrentVersionExpiration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("noncurrentVersionExpiration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("noncurrentVersionExpiration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LifecycleRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LifecycleRule) UnmarshalBinary(b []byte) error {
	var res LifecycleRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
