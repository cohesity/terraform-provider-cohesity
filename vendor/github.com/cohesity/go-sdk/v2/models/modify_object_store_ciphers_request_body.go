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

// ModifyObjectStoreCiphersRequestBody Request to enable/disable a list of object store ciphers.
//
// Specifies object store ciphers to enable/disable on the cluster.
//
// swagger:model ModifyObjectStoreCiphersRequestBody
type ModifyObjectStoreCiphersRequestBody struct {

	// If true, the ciphers passed in will be enabled on the cluster and all other ciphers will be disabled. If false, the ciphers specified will be disabled and all other ciphers on the cluster will be enabled.
	// Required: true
	Enable *bool `json:"enable"`

	// Specifies a list of object store ciphers to enable/disable on the cluster.
	// Required: true
	// Min Items: 1
	// Unique: true
	Ciphers []string `json:"ciphers"`
}

// Validate validates this modify object store ciphers request body
func (m *ModifyObjectStoreCiphersRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCiphers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModifyObjectStoreCiphersRequestBody) validateEnable(formats strfmt.Registry) error {

	if err := validate.Required("enable", "body", m.Enable); err != nil {
		return err
	}

	return nil
}

var modifyObjectStoreCiphersRequestBodyCiphersItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TLS_AES_256_GCM_SHA384","TLS_AES_128_GCM_SHA256","TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384","TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256","TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA","TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA","TLS_RSA_WITH_AES_256_GCM_SHA384","TLS_RSA_WITH_AES_128_GCM_SHA256","TLS_RSA_WITH_AES_256_CBC_SHA","TLS_RSA_WITH_AES_128_CBC_SHA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modifyObjectStoreCiphersRequestBodyCiphersItemsEnum = append(modifyObjectStoreCiphersRequestBodyCiphersItemsEnum, v)
	}
}

func (m *ModifyObjectStoreCiphersRequestBody) validateCiphersItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modifyObjectStoreCiphersRequestBodyCiphersItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModifyObjectStoreCiphersRequestBody) validateCiphers(formats strfmt.Registry) error {

	if err := validate.Required("ciphers", "body", m.Ciphers); err != nil {
		return err
	}

	iCiphersSize := int64(len(m.Ciphers))

	if err := validate.MinItems("ciphers", "body", iCiphersSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("ciphers", "body", m.Ciphers); err != nil {
		return err
	}

	for i := 0; i < len(m.Ciphers); i++ {

		// value enum
		if err := m.validateCiphersItemsEnum("ciphers"+"."+strconv.Itoa(i), "body", m.Ciphers[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this modify object store ciphers request body based on context it is used
func (m *ModifyObjectStoreCiphersRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModifyObjectStoreCiphersRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifyObjectStoreCiphersRequestBody) UnmarshalBinary(b []byte) error {
	var res ModifyObjectStoreCiphersRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
