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

// RecoverAwsS3BucketParams Recover AWS S3 Bucket params.
//
// Specifies the parameters to recover AWS S3 Buckets.
//
// swagger:model RecoverAwsS3BucketParams
type RecoverAwsS3BucketParams struct {

	// Specifies the Protection Group Runs params to recover. All the VM's that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object's snapshot.
	RecoverProtectionGroupRunsParams []*RecoverProtectionGroupRunParams `json:"recoverProtectionGroupRunsParams"`

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kAWS"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the filtering policy for S3 Bucket restore.
	AwsS3BucketRestoreFilterPolicy *AwsS3BucketRestoreFilterPolicy `json:"awsS3BucketRestoreFilterPolicy,omitempty"`

	// Specifies the params for recovering to an AWS target.
	AwsTargetParams *AwsTargetParamsForRecoverS3 `json:"awsTargetParams,omitempty"`
}

// Validate validates this recover aws s3 bucket params
func (m *RecoverAwsS3BucketParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverProtectionGroupRunsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsS3BucketRestoreFilterPolicy(formats); err != nil {
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

func (m *RecoverAwsS3BucketParams) validateRecoverProtectionGroupRunsParams(formats strfmt.Registry) error {
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

var recoverAwsS3BucketParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAWS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverAwsS3BucketParamsTypeTargetEnvironmentPropEnum = append(recoverAwsS3BucketParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverAwsS3BucketParamsTargetEnvironmentKAWS captures enum value "kAWS"
	RecoverAwsS3BucketParamsTargetEnvironmentKAWS string = "kAWS"
)

// prop value enum
func (m *RecoverAwsS3BucketParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverAwsS3BucketParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverAwsS3BucketParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverAwsS3BucketParams) validateAwsS3BucketRestoreFilterPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsS3BucketRestoreFilterPolicy) { // not required
		return nil
	}

	if m.AwsS3BucketRestoreFilterPolicy != nil {
		if err := m.AwsS3BucketRestoreFilterPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsS3BucketRestoreFilterPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsS3BucketRestoreFilterPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsS3BucketParams) validateAwsTargetParams(formats strfmt.Registry) error {
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

// ContextValidate validate this recover aws s3 bucket params based on the context it is used
func (m *RecoverAwsS3BucketParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecoverProtectionGroupRunsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsS3BucketRestoreFilterPolicy(ctx, formats); err != nil {
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

func (m *RecoverAwsS3BucketParams) contextValidateRecoverProtectionGroupRunsParams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverAwsS3BucketParams) contextValidateAwsS3BucketRestoreFilterPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsS3BucketRestoreFilterPolicy != nil {

		if swag.IsZero(m.AwsS3BucketRestoreFilterPolicy) { // not required
			return nil
		}

		if err := m.AwsS3BucketRestoreFilterPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsS3BucketRestoreFilterPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsS3BucketRestoreFilterPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsS3BucketParams) contextValidateAwsTargetParams(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RecoverAwsS3BucketParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAwsS3BucketParams) UnmarshalBinary(b []byte) error {
	var res RecoverAwsS3BucketParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
