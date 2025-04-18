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

// CommonBackupParams Specifies the common parameters for backup. These parameters are common across protection group and object protection.
//
// swagger:model CommonBackupParams
type CommonBackupParams struct {

	// Specifies the unique id of the Protection Policy. The Policy settings will be attached with every object and will be used in backup.
	PolicyID *string `json:"policyId,omitempty"`

	// Specifies the policies for protecting an object with multiple policies. One of policyId and policyConfig is required.
	PolicyConfig *PolicyConfig `json:"policyConfig,omitempty"`

	// Specifies the Storage Domain (View Box) ID where the object backup will be taken. This is not required if Cloud archive direct is benig used.
	StorageDomainID *int64 `json:"storageDomainId,omitempty"`

	// Specifies the start time for the backup to trigger. If no start time is specified, the start time will default to 2am. This field is only used if the policy has a daily or monthly schedule.
	StartTime *TimeOfDay `json:"startTime,omitempty"`

	// Specifies the priority for the objects backup.
	// Enum: ["kLow","kMedium","kHigh"]
	Priority *string `json:"priority,omitempty"`

	// Specifies the SLA parameters for list of objects.
	SLA []*SLARule `json:"sla"`

	// Specifies whether object backup will be written to HDD or SSD.
	// Enum: ["kBackupHDD","kBackupSSD","kTestAndDevHigh","kBackupAll"]
	QosPolicy *string `json:"qosPolicy,omitempty"`

	// Specifies whether currently executing object backup should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false.
	AbortInBlackouts *bool `json:"abortInBlackouts,omitempty"`

	// Specifies whether to skip Rigel for backup or not.
	SkipRigelForBackup *bool `json:"skipRigelForBackup,omitempty"`

	// Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`
}

// Validate validates this common backup params
func (m *CommonBackupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSLA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQosPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonBackupParams) validatePolicyConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyConfig) { // not required
		return nil
	}

	if m.PolicyConfig != nil {
		if err := m.PolicyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policyConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CommonBackupParams) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if m.StartTime != nil {
		if err := m.StartTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startTime")
			}
			return err
		}
	}

	return nil
}

var commonBackupParamsTypePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kLow","kMedium","kHigh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonBackupParamsTypePriorityPropEnum = append(commonBackupParamsTypePriorityPropEnum, v)
	}
}

const (

	// CommonBackupParamsPriorityKLow captures enum value "kLow"
	CommonBackupParamsPriorityKLow string = "kLow"

	// CommonBackupParamsPriorityKMedium captures enum value "kMedium"
	CommonBackupParamsPriorityKMedium string = "kMedium"

	// CommonBackupParamsPriorityKHigh captures enum value "kHigh"
	CommonBackupParamsPriorityKHigh string = "kHigh"
)

// prop value enum
func (m *CommonBackupParams) validatePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonBackupParamsTypePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonBackupParams) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriorityEnum("priority", "body", *m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *CommonBackupParams) validateSLA(formats strfmt.Registry) error {
	if swag.IsZero(m.SLA) { // not required
		return nil
	}

	for i := 0; i < len(m.SLA); i++ {
		if swag.IsZero(m.SLA[i]) { // not required
			continue
		}

		if m.SLA[i] != nil {
			if err := m.SLA[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sla" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sla" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var commonBackupParamsTypeQosPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kBackupHDD","kBackupSSD","kTestAndDevHigh","kBackupAll"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonBackupParamsTypeQosPolicyPropEnum = append(commonBackupParamsTypeQosPolicyPropEnum, v)
	}
}

const (

	// CommonBackupParamsQosPolicyKBackupHDD captures enum value "kBackupHDD"
	CommonBackupParamsQosPolicyKBackupHDD string = "kBackupHDD"

	// CommonBackupParamsQosPolicyKBackupSSD captures enum value "kBackupSSD"
	CommonBackupParamsQosPolicyKBackupSSD string = "kBackupSSD"

	// CommonBackupParamsQosPolicyKTestAndDevHigh captures enum value "kTestAndDevHigh"
	CommonBackupParamsQosPolicyKTestAndDevHigh string = "kTestAndDevHigh"

	// CommonBackupParamsQosPolicyKBackupAll captures enum value "kBackupAll"
	CommonBackupParamsQosPolicyKBackupAll string = "kBackupAll"
)

// prop value enum
func (m *CommonBackupParams) validateQosPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonBackupParamsTypeQosPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonBackupParams) validateQosPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.QosPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateQosPolicyEnum("qosPolicy", "body", *m.QosPolicy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this common backup params based on the context it is used
func (m *CommonBackupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSLA(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonBackupParams) contextValidatePolicyConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.PolicyConfig != nil {

		if swag.IsZero(m.PolicyConfig) { // not required
			return nil
		}

		if err := m.PolicyConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policyConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CommonBackupParams) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if m.StartTime != nil {

		if swag.IsZero(m.StartTime) { // not required
			return nil
		}

		if err := m.StartTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startTime")
			}
			return err
		}
	}

	return nil
}

func (m *CommonBackupParams) contextValidateSLA(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SLA); i++ {

		if m.SLA[i] != nil {

			if swag.IsZero(m.SLA[i]) { // not required
				return nil
			}

			if err := m.SLA[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sla" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sla" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonBackupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonBackupParams) UnmarshalBinary(b []byte) error {
	var res CommonBackupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
