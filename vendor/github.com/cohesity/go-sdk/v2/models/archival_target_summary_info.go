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

// ArchivalTargetSummaryInfo Archival target.
//
// Specifies archival target summary information.
//
// swagger:model ArchivalTargetSummaryInfo
type ArchivalTargetSummaryInfo struct {

	// Specifies the archival target ID.
	TargetID *int64 `json:"targetId,omitempty"`

	// Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
	ArchivalTaskID *string `json:"archivalTaskId,omitempty"`

	// Specifies the archival target name.
	TargetName *string `json:"targetName,omitempty"`

	// Specifies the archival target type.
	// Enum: ["Tape","Cloud","Nas"]
	TargetType *string `json:"targetType,omitempty"`

	// Specifies the usage type for the target.
	// Enum: ["Archival","Tiering","Rpaas"]
	UsageType *string `json:"usageType,omitempty"`

	// Specifies the ownership context for the target.
	// Enum: ["Local","FortKnox"]
	OwnershipContext *string `json:"ownershipContext,omitempty"`

	// Specifies the tier level settings configured with archival target. This will be specified if the run is a CAD run.
	TierSettings *ArchivalTargetTierInfo `json:"tierSettings,omitempty"`
}

// Validate validates this archival target summary info
func (m *ArchivalTargetSummaryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnershipContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var archivalTargetSummaryInfoTypeTargetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Tape","Cloud","Nas"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archivalTargetSummaryInfoTypeTargetTypePropEnum = append(archivalTargetSummaryInfoTypeTargetTypePropEnum, v)
	}
}

const (

	// ArchivalTargetSummaryInfoTargetTypeTape captures enum value "Tape"
	ArchivalTargetSummaryInfoTargetTypeTape string = "Tape"

	// ArchivalTargetSummaryInfoTargetTypeCloud captures enum value "Cloud"
	ArchivalTargetSummaryInfoTargetTypeCloud string = "Cloud"

	// ArchivalTargetSummaryInfoTargetTypeNas captures enum value "Nas"
	ArchivalTargetSummaryInfoTargetTypeNas string = "Nas"
)

// prop value enum
func (m *ArchivalTargetSummaryInfo) validateTargetTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archivalTargetSummaryInfoTypeTargetTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchivalTargetSummaryInfo) validateTargetType(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetTypeEnum("targetType", "body", *m.TargetType); err != nil {
		return err
	}

	return nil
}

var archivalTargetSummaryInfoTypeUsageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Archival","Tiering","Rpaas"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archivalTargetSummaryInfoTypeUsageTypePropEnum = append(archivalTargetSummaryInfoTypeUsageTypePropEnum, v)
	}
}

const (

	// ArchivalTargetSummaryInfoUsageTypeArchival captures enum value "Archival"
	ArchivalTargetSummaryInfoUsageTypeArchival string = "Archival"

	// ArchivalTargetSummaryInfoUsageTypeTiering captures enum value "Tiering"
	ArchivalTargetSummaryInfoUsageTypeTiering string = "Tiering"

	// ArchivalTargetSummaryInfoUsageTypeRpaas captures enum value "Rpaas"
	ArchivalTargetSummaryInfoUsageTypeRpaas string = "Rpaas"
)

// prop value enum
func (m *ArchivalTargetSummaryInfo) validateUsageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archivalTargetSummaryInfoTypeUsageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchivalTargetSummaryInfo) validateUsageType(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsageTypeEnum("usageType", "body", *m.UsageType); err != nil {
		return err
	}

	return nil
}

var archivalTargetSummaryInfoTypeOwnershipContextPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Local","FortKnox"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archivalTargetSummaryInfoTypeOwnershipContextPropEnum = append(archivalTargetSummaryInfoTypeOwnershipContextPropEnum, v)
	}
}

const (

	// ArchivalTargetSummaryInfoOwnershipContextLocal captures enum value "Local"
	ArchivalTargetSummaryInfoOwnershipContextLocal string = "Local"

	// ArchivalTargetSummaryInfoOwnershipContextFortKnox captures enum value "FortKnox"
	ArchivalTargetSummaryInfoOwnershipContextFortKnox string = "FortKnox"
)

// prop value enum
func (m *ArchivalTargetSummaryInfo) validateOwnershipContextEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archivalTargetSummaryInfoTypeOwnershipContextPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchivalTargetSummaryInfo) validateOwnershipContext(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnershipContext) { // not required
		return nil
	}

	// value enum
	if err := m.validateOwnershipContextEnum("ownershipContext", "body", *m.OwnershipContext); err != nil {
		return err
	}

	return nil
}

func (m *ArchivalTargetSummaryInfo) validateTierSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.TierSettings) { // not required
		return nil
	}

	if m.TierSettings != nil {
		if err := m.TierSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tierSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tierSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this archival target summary info based on the context it is used
func (m *ArchivalTargetSummaryInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTierSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArchivalTargetSummaryInfo) contextValidateTierSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.TierSettings != nil {

		if swag.IsZero(m.TierSettings) { // not required
			return nil
		}

		if err := m.TierSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tierSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tierSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArchivalTargetSummaryInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchivalTargetSummaryInfo) UnmarshalBinary(b []byte) error {
	var res ArchivalTargetSummaryInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
