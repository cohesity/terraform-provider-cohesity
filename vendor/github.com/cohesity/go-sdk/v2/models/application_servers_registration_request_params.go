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

// ApplicationServersRegistrationRequestParams Application Servers Registration Request Params
//
// Specifies the application server registration request params.
//
// swagger:model ApplicationServersRegistrationRequestParams
type ApplicationServersRegistrationRequestParams struct {
	CommonApplicationServersRegistrationParams

	// Specifies the credentials that will be used to log into the application environment.
	Credentials *Credentials `json:"credentials,omitempty"`

	// Specifies application specific credentials vec.
	AppCredentialsVec []*ApplicationCredentials `json:"appCredentialsVec"`

	// Set to true if credentials are encrypted by internal magneto key
	IsInternalEncrypted *bool `json:"isInternalEncrypted,omitempty"`

	// If set the user has encrypted the credential with userEncryptionKey. If both isInternalEncrypted and userEncryptionKey is set, it is assumed that credentials are first encrypted using 'internalEncryptionKey' and then encrypted using 'userEncryptionKey'.
	UserEncryptionKey *string `json:"userEncryptionKey,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplicationServersRegistrationRequestParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonApplicationServersRegistrationParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonApplicationServersRegistrationParams = aO0

	// AO1
	var dataAO1 struct {
		Credentials *Credentials `json:"credentials,omitempty"`

		AppCredentialsVec []*ApplicationCredentials `json:"appCredentialsVec"`

		IsInternalEncrypted *bool `json:"isInternalEncrypted,omitempty"`

		UserEncryptionKey *string `json:"userEncryptionKey,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Credentials = dataAO1.Credentials

	m.AppCredentialsVec = dataAO1.AppCredentialsVec

	m.IsInternalEncrypted = dataAO1.IsInternalEncrypted

	m.UserEncryptionKey = dataAO1.UserEncryptionKey

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplicationServersRegistrationRequestParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonApplicationServersRegistrationParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Credentials *Credentials `json:"credentials,omitempty"`

		AppCredentialsVec []*ApplicationCredentials `json:"appCredentialsVec"`

		IsInternalEncrypted *bool `json:"isInternalEncrypted,omitempty"`

		UserEncryptionKey *string `json:"userEncryptionKey,omitempty"`
	}

	dataAO1.Credentials = m.Credentials

	dataAO1.AppCredentialsVec = m.AppCredentialsVec

	dataAO1.IsInternalEncrypted = m.IsInternalEncrypted

	dataAO1.UserEncryptionKey = m.UserEncryptionKey

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this application servers registration request params
func (m *ApplicationServersRegistrationRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonApplicationServersRegistrationParams
	if err := m.CommonApplicationServersRegistrationParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppCredentialsVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationServersRegistrationRequestParams) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationServersRegistrationRequestParams) validateAppCredentialsVec(formats strfmt.Registry) error {

	if swag.IsZero(m.AppCredentialsVec) { // not required
		return nil
	}

	for i := 0; i < len(m.AppCredentialsVec); i++ {
		if swag.IsZero(m.AppCredentialsVec[i]) { // not required
			continue
		}

		if m.AppCredentialsVec[i] != nil {
			if err := m.AppCredentialsVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appCredentialsVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appCredentialsVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this application servers registration request params based on the context it is used
func (m *ApplicationServersRegistrationRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonApplicationServersRegistrationParams
	if err := m.CommonApplicationServersRegistrationParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppCredentialsVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationServersRegistrationRequestParams) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {

		if swag.IsZero(m.Credentials) { // not required
			return nil
		}

		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationServersRegistrationRequestParams) contextValidateAppCredentialsVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppCredentialsVec); i++ {

		if m.AppCredentialsVec[i] != nil {

			if swag.IsZero(m.AppCredentialsVec[i]) { // not required
				return nil
			}

			if err := m.AppCredentialsVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appCredentialsVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appCredentialsVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationServersRegistrationRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationServersRegistrationRequestParams) UnmarshalBinary(b []byte) error {
	var res ApplicationServersRegistrationRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
