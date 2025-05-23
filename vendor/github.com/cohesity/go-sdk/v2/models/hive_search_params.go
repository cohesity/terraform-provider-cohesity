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

// HiveSearchParams Specifies the parameters which are specific for searching Hive objects.
//
// swagger:model HiveSearchParams
type HiveSearchParams struct {

	// Specifies the search string to search the Hive Objects
	// Required: true
	SearchString *string `json:"searchString"`

	// Specifies one or more Hive object types be searched.
	// Required: true
	HiveObjectTypes []string `json:"hiveObjectTypes"`
}

// Validate validates this hive search params
func (m *HiveSearchParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchString(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHiveObjectTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HiveSearchParams) validateSearchString(formats strfmt.Registry) error {

	if err := validate.Required("searchString", "body", m.SearchString); err != nil {
		return err
	}

	return nil
}

var hiveSearchParamsHiveObjectTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HiveDatabases","HiveTables","HivePartitions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hiveSearchParamsHiveObjectTypesItemsEnum = append(hiveSearchParamsHiveObjectTypesItemsEnum, v)
	}
}

func (m *HiveSearchParams) validateHiveObjectTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hiveSearchParamsHiveObjectTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HiveSearchParams) validateHiveObjectTypes(formats strfmt.Registry) error {

	if err := validate.Required("hiveObjectTypes", "body", m.HiveObjectTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.HiveObjectTypes); i++ {

		// value enum
		if err := m.validateHiveObjectTypesItemsEnum("hiveObjectTypes"+"."+strconv.Itoa(i), "body", m.HiveObjectTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this hive search params based on context it is used
func (m *HiveSearchParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HiveSearchParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HiveSearchParams) UnmarshalBinary(b []byte) error {
	var res HiveSearchParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
