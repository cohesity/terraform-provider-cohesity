// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CancellationTimeoutParams Cancellation Timeout Params.
//
// Specifies timeouts for different backup types (kFull, kRegular etc.)
//
// swagger:model CancellationTimeoutParams
type CancellationTimeoutParams struct {

	// Specifies the timeout in mins
	TimeoutMins *int64 `json:"timeoutMins,omitempty"`

	// The scheduled backup type(kFull, kRegular etc.)
	// Enum: ["kRegular","kFull","kLog","kSystem","kHydrateCDP","kStorageArraySnapshot"]
	BackupType *string `json:"backupType,omitempty"`
}

// Validate validates this cancellation timeout params
func (m *CancellationTimeoutParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cancellationTimeoutParamsTypeBackupTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kRegular","kFull","kLog","kSystem","kHydrateCDP","kStorageArraySnapshot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cancellationTimeoutParamsTypeBackupTypePropEnum = append(cancellationTimeoutParamsTypeBackupTypePropEnum, v)
	}
}

const (

	// CancellationTimeoutParamsBackupTypeKRegular captures enum value "kRegular"
	CancellationTimeoutParamsBackupTypeKRegular string = "kRegular"

	// CancellationTimeoutParamsBackupTypeKFull captures enum value "kFull"
	CancellationTimeoutParamsBackupTypeKFull string = "kFull"

	// CancellationTimeoutParamsBackupTypeKLog captures enum value "kLog"
	CancellationTimeoutParamsBackupTypeKLog string = "kLog"

	// CancellationTimeoutParamsBackupTypeKSystem captures enum value "kSystem"
	CancellationTimeoutParamsBackupTypeKSystem string = "kSystem"

	// CancellationTimeoutParamsBackupTypeKHydrateCDP captures enum value "kHydrateCDP"
	CancellationTimeoutParamsBackupTypeKHydrateCDP string = "kHydrateCDP"

	// CancellationTimeoutParamsBackupTypeKStorageArraySnapshot captures enum value "kStorageArraySnapshot"
	CancellationTimeoutParamsBackupTypeKStorageArraySnapshot string = "kStorageArraySnapshot"
)

// prop value enum
func (m *CancellationTimeoutParams) validateBackupTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cancellationTimeoutParamsTypeBackupTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CancellationTimeoutParams) validateBackupType(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBackupTypeEnum("backupType", "body", *m.BackupType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cancellation timeout params based on context it is used
func (m *CancellationTimeoutParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CancellationTimeoutParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancellationTimeoutParams) UnmarshalBinary(b []byte) error {
	var res CancellationTimeoutParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
