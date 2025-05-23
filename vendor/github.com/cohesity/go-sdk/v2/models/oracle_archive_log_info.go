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

// OracleArchiveLogInfo OracleArchiveLogInfo
//
// Specifies information related to archive log restore.
//
// swagger:model OracleArchiveLogInfo
type OracleArchiveLogInfo struct {

	// Specifies the type of range.
	// Enum: ["Time","Scn","Sequence"]
	RangeType *string `json:"rangeType,omitempty"`

	// Specifies an array of oracle restore ranges.
	RangeInfoVec []*OracleRangeMetaInfo `json:"rangeInfoVec"`

	// Specifies destination where archive logs are to be restored.
	ArchiveLogRestoreDest *string `json:"archiveLogRestoreDest,omitempty"`
}

// Validate validates this oracle archive log info
func (m *OracleArchiveLogInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRangeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRangeInfoVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleArchiveLogInfoTypeRangeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Time","Scn","Sequence"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleArchiveLogInfoTypeRangeTypePropEnum = append(oracleArchiveLogInfoTypeRangeTypePropEnum, v)
	}
}

const (

	// OracleArchiveLogInfoRangeTypeTime captures enum value "Time"
	OracleArchiveLogInfoRangeTypeTime string = "Time"

	// OracleArchiveLogInfoRangeTypeScn captures enum value "Scn"
	OracleArchiveLogInfoRangeTypeScn string = "Scn"

	// OracleArchiveLogInfoRangeTypeSequence captures enum value "Sequence"
	OracleArchiveLogInfoRangeTypeSequence string = "Sequence"
)

// prop value enum
func (m *OracleArchiveLogInfo) validateRangeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleArchiveLogInfoTypeRangeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleArchiveLogInfo) validateRangeType(formats strfmt.Registry) error {
	if swag.IsZero(m.RangeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRangeTypeEnum("rangeType", "body", *m.RangeType); err != nil {
		return err
	}

	return nil
}

func (m *OracleArchiveLogInfo) validateRangeInfoVec(formats strfmt.Registry) error {
	if swag.IsZero(m.RangeInfoVec) { // not required
		return nil
	}

	for i := 0; i < len(m.RangeInfoVec); i++ {
		if swag.IsZero(m.RangeInfoVec[i]) { // not required
			continue
		}

		if m.RangeInfoVec[i] != nil {
			if err := m.RangeInfoVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rangeInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rangeInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this oracle archive log info based on the context it is used
func (m *OracleArchiveLogInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRangeInfoVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleArchiveLogInfo) contextValidateRangeInfoVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RangeInfoVec); i++ {

		if m.RangeInfoVec[i] != nil {

			if swag.IsZero(m.RangeInfoVec[i]) { // not required
				return nil
			}

			if err := m.RangeInfoVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rangeInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rangeInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OracleArchiveLogInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleArchiveLogInfo) UnmarshalBinary(b []byte) error {
	var res OracleArchiveLogInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
