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

// KmsCreateRequestParameters Request to create a KMS with specified configuration.
//
// swagger:model KmsCreateRequestParameters
type KmsCreateRequestParameters struct {

	// AWS KMS conifg.
	AwsKms *AwsKmsConfiguration `json:"awsKms,omitempty"`

	// Azure KMS config.
	AzureKms *AzureKmsConfiguration `json:"azureKms,omitempty"`

	// Cryptsoft KMS config.
	CryptsoftKms *CryptsoftKmsConfiguration `json:"cryptsoftKms,omitempty"`

	// The Id of a KMS server.
	ID *int64 `json:"id,omitempty"`

	// Specifies name of the key.
	KeyName *string `json:"keyName,omitempty"`

	// Specifies the consumption model for the KMS Key.
	// 'Local' indicates an internal KMS object.
	// 'FortKnox' indicates an FortKnox KMS object.
	// Enum: ["Local","FortKnox"]
	OwnershipContext *string `json:"ownershipContext,omitempty"`

	// Specifies the name given to the KMS Server.
	ServerName *string `json:"serverName,omitempty"`

	// Specifies the type of key mangement system.
	// 'kInternalKMS' indicates an internal KMS object.
	// 'kAwsKMS' indicates an Aws KMS object.
	// 'kCryptsoftKMS' indicates a Cryptsoft KMS object.
	// 'kAzureKMS' indicates a Azure KMS object.
	// Enum: ["kInternalKMS","kAwsKMS","kCryptsoftKMS","kAzureKMS"]
	ServerType *string `json:"serverType,omitempty"`

	// Specifies the usage type of the kms config. kArchival indicates
	// this is used for regular archival. kRpaasArchival indicates this
	// is used for RPaaS only.
	// 'kArchival' indicates an internal KMS object.
	// 'kRpaasArchival' indicates an Aws KMS object.
	// Enum: ["kArchival","kRpaasArchival"]
	UsageType *string `json:"usageType,omitempty"`

	// Specifies the list of Vault Ids.
	VaultIDList []int64 `json:"vaultIdList"`

	// Specifies the list of View Box Ids.
	ViewBoxIDList []int64 `json:"viewBoxIdList"`
}

// Validate validates this kms create request parameters
func (m *KmsCreateRequestParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsKms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureKms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCryptsoftKms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnershipContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KmsCreateRequestParameters) validateAwsKms(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsKms) { // not required
		return nil
	}

	if m.AwsKms != nil {
		if err := m.AwsKms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsKms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsKms")
			}
			return err
		}
	}

	return nil
}

func (m *KmsCreateRequestParameters) validateAzureKms(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureKms) { // not required
		return nil
	}

	if m.AzureKms != nil {
		if err := m.AzureKms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureKms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureKms")
			}
			return err
		}
	}

	return nil
}

func (m *KmsCreateRequestParameters) validateCryptsoftKms(formats strfmt.Registry) error {
	if swag.IsZero(m.CryptsoftKms) { // not required
		return nil
	}

	if m.CryptsoftKms != nil {
		if err := m.CryptsoftKms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cryptsoftKms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cryptsoftKms")
			}
			return err
		}
	}

	return nil
}

var kmsCreateRequestParametersTypeOwnershipContextPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Local","FortKnox"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kmsCreateRequestParametersTypeOwnershipContextPropEnum = append(kmsCreateRequestParametersTypeOwnershipContextPropEnum, v)
	}
}

const (

	// KmsCreateRequestParametersOwnershipContextLocal captures enum value "Local"
	KmsCreateRequestParametersOwnershipContextLocal string = "Local"

	// KmsCreateRequestParametersOwnershipContextFortKnox captures enum value "FortKnox"
	KmsCreateRequestParametersOwnershipContextFortKnox string = "FortKnox"
)

// prop value enum
func (m *KmsCreateRequestParameters) validateOwnershipContextEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kmsCreateRequestParametersTypeOwnershipContextPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KmsCreateRequestParameters) validateOwnershipContext(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnershipContext) { // not required
		return nil
	}

	// value enum
	if err := m.validateOwnershipContextEnum("ownershipContext", "body", *m.OwnershipContext); err != nil {
		return err
	}

	return nil
}

var kmsCreateRequestParametersTypeServerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kInternalKMS","kAwsKMS","kCryptsoftKMS","kAzureKMS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kmsCreateRequestParametersTypeServerTypePropEnum = append(kmsCreateRequestParametersTypeServerTypePropEnum, v)
	}
}

const (

	// KmsCreateRequestParametersServerTypeKInternalKMS captures enum value "kInternalKMS"
	KmsCreateRequestParametersServerTypeKInternalKMS string = "kInternalKMS"

	// KmsCreateRequestParametersServerTypeKAwsKMS captures enum value "kAwsKMS"
	KmsCreateRequestParametersServerTypeKAwsKMS string = "kAwsKMS"

	// KmsCreateRequestParametersServerTypeKCryptsoftKMS captures enum value "kCryptsoftKMS"
	KmsCreateRequestParametersServerTypeKCryptsoftKMS string = "kCryptsoftKMS"

	// KmsCreateRequestParametersServerTypeKAzureKMS captures enum value "kAzureKMS"
	KmsCreateRequestParametersServerTypeKAzureKMS string = "kAzureKMS"
)

// prop value enum
func (m *KmsCreateRequestParameters) validateServerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kmsCreateRequestParametersTypeServerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KmsCreateRequestParameters) validateServerType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerTypeEnum("serverType", "body", *m.ServerType); err != nil {
		return err
	}

	return nil
}

var kmsCreateRequestParametersTypeUsageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kArchival","kRpaasArchival"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kmsCreateRequestParametersTypeUsageTypePropEnum = append(kmsCreateRequestParametersTypeUsageTypePropEnum, v)
	}
}

const (

	// KmsCreateRequestParametersUsageTypeKArchival captures enum value "kArchival"
	KmsCreateRequestParametersUsageTypeKArchival string = "kArchival"

	// KmsCreateRequestParametersUsageTypeKRpaasArchival captures enum value "kRpaasArchival"
	KmsCreateRequestParametersUsageTypeKRpaasArchival string = "kRpaasArchival"
)

// prop value enum
func (m *KmsCreateRequestParameters) validateUsageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kmsCreateRequestParametersTypeUsageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KmsCreateRequestParameters) validateUsageType(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsageTypeEnum("usageType", "body", *m.UsageType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this kms create request parameters based on the context it is used
func (m *KmsCreateRequestParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsKms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureKms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCryptsoftKms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KmsCreateRequestParameters) contextValidateAwsKms(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsKms != nil {

		if swag.IsZero(m.AwsKms) { // not required
			return nil
		}

		if err := m.AwsKms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsKms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsKms")
			}
			return err
		}
	}

	return nil
}

func (m *KmsCreateRequestParameters) contextValidateAzureKms(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureKms != nil {

		if swag.IsZero(m.AzureKms) { // not required
			return nil
		}

		if err := m.AzureKms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureKms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureKms")
			}
			return err
		}
	}

	return nil
}

func (m *KmsCreateRequestParameters) contextValidateCryptsoftKms(ctx context.Context, formats strfmt.Registry) error {

	if m.CryptsoftKms != nil {

		if swag.IsZero(m.CryptsoftKms) { // not required
			return nil
		}

		if err := m.CryptsoftKms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cryptsoftKms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cryptsoftKms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KmsCreateRequestParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KmsCreateRequestParameters) UnmarshalBinary(b []byte) error {
	var res KmsCreateRequestParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
