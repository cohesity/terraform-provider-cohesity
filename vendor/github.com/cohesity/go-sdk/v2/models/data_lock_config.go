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

// DataLockConfig Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
//
// swagger:model DataLockConfig
type DataLockConfig struct {

	// Specifies the type of WORM retention type.
	// 'Compliance' implies WORM retention is set for compliance reason.
	// 'Administrative' implies WORM retention is set for administrative purposes.
	// Required: true
	// Enum: ["Compliance","Administrative"]
	Mode *string `json:"mode"`

	// Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
	// Required: true
	// Enum: ["Days","Weeks","Months","Years"]
	Unit *string `json:"unit"`

	// Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
	// Required: true
	// Minimum: 1
	Duration *int64 `json:"duration"`

	// Specifies whether objects in the external target associated with this policy need to be made immutable.
	EnableWormOnExternalTarget *bool `json:"enableWormOnExternalTarget,omitempty"`
}

// Validate validates this data lock config
func (m *DataLockConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dataLockConfigTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Compliance","Administrative"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataLockConfigTypeModePropEnum = append(dataLockConfigTypeModePropEnum, v)
	}
}

const (

	// DataLockConfigModeCompliance captures enum value "Compliance"
	DataLockConfigModeCompliance string = "Compliance"

	// DataLockConfigModeAdministrative captures enum value "Administrative"
	DataLockConfigModeAdministrative string = "Administrative"
)

// prop value enum
func (m *DataLockConfig) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataLockConfigTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataLockConfig) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

var dataLockConfigTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Days","Weeks","Months","Years"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataLockConfigTypeUnitPropEnum = append(dataLockConfigTypeUnitPropEnum, v)
	}
}

const (

	// DataLockConfigUnitDays captures enum value "Days"
	DataLockConfigUnitDays string = "Days"

	// DataLockConfigUnitWeeks captures enum value "Weeks"
	DataLockConfigUnitWeeks string = "Weeks"

	// DataLockConfigUnitMonths captures enum value "Months"
	DataLockConfigUnitMonths string = "Months"

	// DataLockConfigUnitYears captures enum value "Years"
	DataLockConfigUnitYears string = "Years"
)

// prop value enum
func (m *DataLockConfig) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataLockConfigTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataLockConfig) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *DataLockConfig) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	if err := validate.MinimumInt("duration", "body", *m.Duration, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data lock config based on context it is used
func (m *DataLockConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataLockConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataLockConfig) UnmarshalBinary(b []byte) error {
	var res DataLockConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
