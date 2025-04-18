// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupPolicy Backup Schedule and Retention.
//
// Specifies the backup schedule and retentions of a Protection Policy.
//
// swagger:model BackupPolicy
type BackupPolicy struct {

	// Specifies the Incremental and Full policy settings and also the common Retention policy settings.
	// Required: true
	Regular *RegularBackupPolicy `json:"regular"`

	// Specifies the Log backup schedule of a Protection Policy.
	Log *LogBackupPolicy `json:"log,omitempty"`

	// Specifies the BMR backup schedule of a Protection Policy.
	Bmr *BmrBackupPolicy `json:"bmr,omitempty"`

	// Specifies the settings for CDP (Continious Data Protection) Protection policy.
	Cdp *CdpBackupPolicy `json:"cdp,omitempty"`

	// Specifies the settings for Storage Array Snapshot Protection policy.
	StorageArraySnapshot *StorageArraySnapshotBackupPolicy `json:"storageArraySnapshot,omitempty"`

	// Specifies the backup timeouts for different type of runs(kFull, kRegular etc.).
	RunTimeouts []*CancellationTimeoutParams `json:"runTimeouts"`
}

// Validate validates this backup policy
func (m *BackupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegular(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBmr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCdp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArraySnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunTimeouts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPolicy) validateRegular(formats strfmt.Registry) error {

	if err := validate.Required("regular", "body", m.Regular); err != nil {
		return err
	}

	if m.Regular != nil {
		if err := m.Regular.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regular")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("regular")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) validateLog(formats strfmt.Registry) error {
	if swag.IsZero(m.Log) { // not required
		return nil
	}

	if m.Log != nil {
		if err := m.Log.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) validateBmr(formats strfmt.Registry) error {
	if swag.IsZero(m.Bmr) { // not required
		return nil
	}

	if m.Bmr != nil {
		if err := m.Bmr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bmr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bmr")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) validateCdp(formats strfmt.Registry) error {
	if swag.IsZero(m.Cdp) { // not required
		return nil
	}

	if m.Cdp != nil {
		if err := m.Cdp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdp")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) validateStorageArraySnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageArraySnapshot) { // not required
		return nil
	}

	if m.StorageArraySnapshot != nil {
		if err := m.StorageArraySnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) validateRunTimeouts(formats strfmt.Registry) error {
	if swag.IsZero(m.RunTimeouts) { // not required
		return nil
	}

	for i := 0; i < len(m.RunTimeouts); i++ {
		if swag.IsZero(m.RunTimeouts[i]) { // not required
			continue
		}

		if m.RunTimeouts[i] != nil {
			if err := m.RunTimeouts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this backup policy based on the context it is used
func (m *BackupPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegular(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBmr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCdp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageArraySnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunTimeouts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPolicy) contextValidateRegular(ctx context.Context, formats strfmt.Registry) error {

	if m.Regular != nil {

		if err := m.Regular.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regular")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("regular")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

	if m.Log != nil {

		if swag.IsZero(m.Log) { // not required
			return nil
		}

		if err := m.Log.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) contextValidateBmr(ctx context.Context, formats strfmt.Registry) error {

	if m.Bmr != nil {

		if swag.IsZero(m.Bmr) { // not required
			return nil
		}

		if err := m.Bmr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bmr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bmr")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) contextValidateCdp(ctx context.Context, formats strfmt.Registry) error {

	if m.Cdp != nil {

		if swag.IsZero(m.Cdp) { // not required
			return nil
		}

		if err := m.Cdp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdp")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) contextValidateStorageArraySnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageArraySnapshot != nil {

		if swag.IsZero(m.StorageArraySnapshot) { // not required
			return nil
		}

		if err := m.StorageArraySnapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPolicy) contextValidateRunTimeouts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RunTimeouts); i++ {

		if m.RunTimeouts[i] != nil {

			if swag.IsZero(m.RunTimeouts[i]) { // not required
				return nil
			}

			if err := m.RunTimeouts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPolicy) UnmarshalBinary(b []byte) error {
	var res BackupPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
