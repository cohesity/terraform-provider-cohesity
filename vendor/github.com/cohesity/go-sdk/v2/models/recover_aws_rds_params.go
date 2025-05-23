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

// RecoverAwsRdsParams Recover AWS RDS params.
//
// Specifies the parameters to recover AWS RDS.
//
// swagger:model RecoverAwsRdsParams
type RecoverAwsRdsParams struct {

	// Specifies the Protection Group Runs params to recover. All the RDS instances that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of RDS instances. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object's snapshot.
	RecoverProtectionGroupRunsParams []*RecoverProtectionGroupRunParams `json:"recoverProtectionGroupRunsParams"`

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kAWS"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the params for recovering to an AWS target.
	AwsTargetParams *AwsTargetParamsForRecoverRds `json:"awsTargetParams,omitempty"`
}

// Validate validates this recover aws rds params
func (m *RecoverAwsRdsParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverProtectionGroupRunsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAwsRdsParams) validateRecoverProtectionGroupRunsParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverProtectionGroupRunsParams) { // not required
		return nil
	}

	for i := 0; i < len(m.RecoverProtectionGroupRunsParams); i++ {
		if swag.IsZero(m.RecoverProtectionGroupRunsParams[i]) { // not required
			continue
		}

		if m.RecoverProtectionGroupRunsParams[i] != nil {
			if err := m.RecoverProtectionGroupRunsParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recoverProtectionGroupRunsParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recoverProtectionGroupRunsParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var recoverAwsRdsParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAWS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverAwsRdsParamsTypeTargetEnvironmentPropEnum = append(recoverAwsRdsParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverAwsRdsParamsTargetEnvironmentKAWS captures enum value "kAWS"
	RecoverAwsRdsParamsTargetEnvironmentKAWS string = "kAWS"
)

// prop value enum
func (m *RecoverAwsRdsParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverAwsRdsParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverAwsRdsParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverAwsRdsParams) validateAwsTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsTargetParams) { // not required
		return nil
	}

	if m.AwsTargetParams != nil {
		if err := m.AwsTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover aws rds params based on the context it is used
func (m *RecoverAwsRdsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecoverProtectionGroupRunsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAwsRdsParams) contextValidateRecoverProtectionGroupRunsParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecoverProtectionGroupRunsParams); i++ {

		if m.RecoverProtectionGroupRunsParams[i] != nil {

			if swag.IsZero(m.RecoverProtectionGroupRunsParams[i]) { // not required
				return nil
			}

			if err := m.RecoverProtectionGroupRunsParams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recoverProtectionGroupRunsParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recoverProtectionGroupRunsParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverAwsRdsParams) contextValidateAwsTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsTargetParams != nil {

		if swag.IsZero(m.AwsTargetParams) { // not required
			return nil
		}

		if err := m.AwsTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverAwsRdsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAwsRdsParams) UnmarshalBinary(b []byte) error {
	var res RecoverAwsRdsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
