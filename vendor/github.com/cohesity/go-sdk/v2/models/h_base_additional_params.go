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

// HBaseAdditionalParams HBAse Additional Params.
//
// Additional params for HBase protection source.
//
// swagger:model HBaseAdditionalParams
type HBaseAdditionalParams struct {

	// The 'Zookeeper Quorum' for this HBase.
	// Read Only: true
	ZookeeperQuorum []string `json:"zookeeperQuorum"`

	// The 'Data root directory' for this HBase.
	// Read Only: true
	DataRootDirectory *string `json:"dataRootDirectory,omitempty"`

	// Authentication type.
	// Read Only: true
	// Enum: ["KERBEROS","NONE"]
	AuthType *string `json:"authType,omitempty"`
}

// Validate validates this h base additional params
func (m *HBaseAdditionalParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hBaseAdditionalParamsTypeAuthTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["KERBEROS","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hBaseAdditionalParamsTypeAuthTypePropEnum = append(hBaseAdditionalParamsTypeAuthTypePropEnum, v)
	}
}

const (

	// HBaseAdditionalParamsAuthTypeKERBEROS captures enum value "KERBEROS"
	HBaseAdditionalParamsAuthTypeKERBEROS string = "KERBEROS"

	// HBaseAdditionalParamsAuthTypeNONE captures enum value "NONE"
	HBaseAdditionalParamsAuthTypeNONE string = "NONE"
)

// prop value enum
func (m *HBaseAdditionalParams) validateAuthTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hBaseAdditionalParamsTypeAuthTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HBaseAdditionalParams) validateAuthType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthTypeEnum("authType", "body", *m.AuthType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this h base additional params based on the context it is used
func (m *HBaseAdditionalParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateZookeeperQuorum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataRootDirectory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HBaseAdditionalParams) contextValidateZookeeperQuorum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "zookeeperQuorum", "body", []string(m.ZookeeperQuorum)); err != nil {
		return err
	}

	return nil
}

func (m *HBaseAdditionalParams) contextValidateDataRootDirectory(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dataRootDirectory", "body", m.DataRootDirectory); err != nil {
		return err
	}

	return nil
}

func (m *HBaseAdditionalParams) contextValidateAuthType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "authType", "body", m.AuthType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HBaseAdditionalParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HBaseAdditionalParams) UnmarshalBinary(b []byte) error {
	var res HBaseAdditionalParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
