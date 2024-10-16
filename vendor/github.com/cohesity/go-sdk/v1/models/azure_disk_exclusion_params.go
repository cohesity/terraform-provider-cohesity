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
)

// AzureDiskExclusionParams Message defining the different criteria to exclude Azure disks from
// backup. This is used to specify both object-level (BackupSourceParams) and
// job-level (EnvBackupParams) exclusion criteria.
// If a criterion is specified at both object-level and job-level, then
// job-level setting will be ignored.
//
// swagger:model AzureDiskExclusionParams
type AzureDiskExclusionParams struct {

	// List of disk resource IDs to exclude.
	// This field is only for object-level exclusions.
	DiskIDVec []string `json:"diskIdVec"`

	// Raw boolean query given as input by the user to exclude volume based
	// on tags.
	// In the current version, the query contains only tags.
	// Example query 1: "K1" = "V1" AND "K2" IN ("V2", "V3") AND "K4" != "V4"
	// Example query 2: "K1" != "V1" OR "K2" NOT IN ("V2", "V3") OR "K4" != "V4"
	// All Keys and Values must be wrapped inside double quotes.
	// Comparision Operators supported : =, !=, IN, NOT IN.
	// Logical Operators supported : AND, OR.
	// We cannot have AND, OR together in the query. Only one of them is allowed.
	// The processed form for this query is stored in the above tag_params_vec.
	RawQuery *string `json:"rawQuery,omitempty"`

	// List of Tag Params to exclude Azure disks.
	TagParamsVec []*AzureDiskExclusionParamsTagParams `json:"tagParamsVec"`
}

// Validate validates this azure disk exclusion params
func (m *AzureDiskExclusionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagParamsVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureDiskExclusionParams) validateTagParamsVec(formats strfmt.Registry) error {
	if swag.IsZero(m.TagParamsVec) { // not required
		return nil
	}

	for i := 0; i < len(m.TagParamsVec); i++ {
		if swag.IsZero(m.TagParamsVec[i]) { // not required
			continue
		}

		if m.TagParamsVec[i] != nil {
			if err := m.TagParamsVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagParamsVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagParamsVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this azure disk exclusion params based on the context it is used
func (m *AzureDiskExclusionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTagParamsVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureDiskExclusionParams) contextValidateTagParamsVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagParamsVec); i++ {

		if m.TagParamsVec[i] != nil {

			if swag.IsZero(m.TagParamsVec[i]) { // not required
				return nil
			}

			if err := m.TagParamsVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagParamsVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagParamsVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureDiskExclusionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureDiskExclusionParams) UnmarshalBinary(b []byte) error {
	var res AzureDiskExclusionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
