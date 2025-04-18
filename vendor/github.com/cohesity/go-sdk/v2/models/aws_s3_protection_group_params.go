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

// AwsS3ProtectionGroupParams Create AWS S3 Protection Request Body
//
// Specifies the parameters which are specific to AWS S3 Protection.
//
// swagger:model AwsS3ProtectionGroupParams
type AwsS3ProtectionGroupParams struct {

	// Specifies the objects to be protected.
	Objects []*AwsS3ProtectionGroupObjectParams `json:"objects"`

	// Specifies whether to skip files on error or not. Default value is false.
	SkipOnError *bool `json:"skipOnError,omitempty"`

	// Specifies the AWS S3 Storage classes to backup.
	StorageClass []string `json:"storageClass"`

	// Specifies whether to backup object level acls. Default value is false.
	BackupObjectLevelACLs *bool `json:"backupObjectLevelACLs,omitempty"`

	// Specifies the id of the parent of the objects.
	// Read Only: true
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies the name of the parent of the objects.
	// Read Only: true
	SourceName *string `json:"sourceName,omitempty"`

	// ARN of the inventory report destination bucket for S3 backups.
	InventoryReportDestination *string `json:"inventoryReportDestination,omitempty"`

	// The prefix in the S3 destination bucket where inventory reports will be stored.
	InventoryReportDestinationPrefix *string `json:"inventoryReportDestinationPrefix,omitempty"`

	// Specifies the frequency to generate inventory reports.
	// Enum: ["Weekly","Monthly"]
	InventoryReportFrequency *string `json:"inventoryReportFrequency,omitempty"`

	// Specifies the baseline incremental frequency.
	// Enum: ["Daily","Weekly","Monthly"]
	BaselineIncrementalFrequency *string `json:"baselineIncrementalFrequency,omitempty"`
}

// Validate validates this aws s3 protection group params
func (m *AwsS3ProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInventoryReportFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaselineIncrementalFrequency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsS3ProtectionGroupParams) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var awsS3ProtectionGroupParamsStorageClassItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AmazonS3Standard","AmazonS3IntelligentTiering","AmazonS3StandardIA","AmazonS3OneZoneIA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsS3ProtectionGroupParamsStorageClassItemsEnum = append(awsS3ProtectionGroupParamsStorageClassItemsEnum, v)
	}
}

func (m *AwsS3ProtectionGroupParams) validateStorageClassItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsS3ProtectionGroupParamsStorageClassItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsS3ProtectionGroupParams) validateStorageClass(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageClass) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageClass); i++ {

		// value enum
		if err := m.validateStorageClassItemsEnum("storageClass"+"."+strconv.Itoa(i), "body", m.StorageClass[i]); err != nil {
			return err
		}

	}

	return nil
}

var awsS3ProtectionGroupParamsTypeInventoryReportFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Weekly","Monthly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsS3ProtectionGroupParamsTypeInventoryReportFrequencyPropEnum = append(awsS3ProtectionGroupParamsTypeInventoryReportFrequencyPropEnum, v)
	}
}

const (

	// AwsS3ProtectionGroupParamsInventoryReportFrequencyWeekly captures enum value "Weekly"
	AwsS3ProtectionGroupParamsInventoryReportFrequencyWeekly string = "Weekly"

	// AwsS3ProtectionGroupParamsInventoryReportFrequencyMonthly captures enum value "Monthly"
	AwsS3ProtectionGroupParamsInventoryReportFrequencyMonthly string = "Monthly"
)

// prop value enum
func (m *AwsS3ProtectionGroupParams) validateInventoryReportFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsS3ProtectionGroupParamsTypeInventoryReportFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsS3ProtectionGroupParams) validateInventoryReportFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.InventoryReportFrequency) { // not required
		return nil
	}

	// value enum
	if err := m.validateInventoryReportFrequencyEnum("inventoryReportFrequency", "body", *m.InventoryReportFrequency); err != nil {
		return err
	}

	return nil
}

var awsS3ProtectionGroupParamsTypeBaselineIncrementalFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Daily","Weekly","Monthly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsS3ProtectionGroupParamsTypeBaselineIncrementalFrequencyPropEnum = append(awsS3ProtectionGroupParamsTypeBaselineIncrementalFrequencyPropEnum, v)
	}
}

const (

	// AwsS3ProtectionGroupParamsBaselineIncrementalFrequencyDaily captures enum value "Daily"
	AwsS3ProtectionGroupParamsBaselineIncrementalFrequencyDaily string = "Daily"

	// AwsS3ProtectionGroupParamsBaselineIncrementalFrequencyWeekly captures enum value "Weekly"
	AwsS3ProtectionGroupParamsBaselineIncrementalFrequencyWeekly string = "Weekly"

	// AwsS3ProtectionGroupParamsBaselineIncrementalFrequencyMonthly captures enum value "Monthly"
	AwsS3ProtectionGroupParamsBaselineIncrementalFrequencyMonthly string = "Monthly"
)

// prop value enum
func (m *AwsS3ProtectionGroupParams) validateBaselineIncrementalFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsS3ProtectionGroupParamsTypeBaselineIncrementalFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsS3ProtectionGroupParams) validateBaselineIncrementalFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.BaselineIncrementalFrequency) { // not required
		return nil
	}

	// value enum
	if err := m.validateBaselineIncrementalFrequencyEnum("baselineIncrementalFrequency", "body", *m.BaselineIncrementalFrequency); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws s3 protection group params based on the context it is used
func (m *AwsS3ProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsS3ProtectionGroupParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsS3ProtectionGroupParams) contextValidateSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *AwsS3ProtectionGroupParams) contextValidateSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceName", "body", m.SourceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsS3ProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsS3ProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res AwsS3ProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
