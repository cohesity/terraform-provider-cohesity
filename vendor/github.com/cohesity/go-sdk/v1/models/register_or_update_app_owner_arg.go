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

// RegisterOrUpdateAppOwnerArg Message to register or update the owner entity on which applications are
// to be backed up.
//
// swagger:model RegisterOrUpdateAppOwnerArg
type RegisterOrUpdateAppOwnerArg struct {

	// Specifies the request attributes.
	APIRequestAttr *APIRequestAttr `json:"apiRequestAttr,omitempty"`

	// Specifies the API version used by this arg.
	APIVersion *APIVersion `json:"apiVersion,omitempty"`

	// Application specific credentials vector. This may be needed in cases where
	// after authorizing with the environment, separate authorization is needed
	// to access an application in the environment (for example, kOracle where
	// after authorizing with the database host, to access a database separate
	// credentials are required).
	AppCredentialsVec []*PrivateAppCredentials `json:"appCredentialsVec"`

	// The types of application environments (for example, kSQL, kExchange,
	// kOracle).
	AppEnvVec []int32 `json:"appEnvVec"`

	// Credentials that will be used to log into the application environment.
	Credentials *PrivateCredentials `json:"credentials,omitempty"`

	// This is populated with last known modification time of the Entity
	// (last_modification_time should be updated (to current time) when a public
	// API call updates an Entity) to be updated. Magneto updates the Entity iff
	// this matches the last_modification_time stored in EntityHierarchyProto,
	// to avoid any race conditions.
	ExpectedEntityMtimeUsecs *int64 `json:"expectedEntityMtimeUsecs,omitempty"`

	// Set to true if credentials are encrypted by internal magneto key.
	IsInternalEncrypted *bool `json:"isInternalEncrypted,omitempty"`

	// This is set to the id of the network-realm from where this source is
	// reachable. This should only be set for a source being registered by a
	// tenant user.
	NetworkRealmID *int64 `json:"networkRealmId,omitempty"`

	// The owner entity on which the application environment needs to be
	// backed up. For example, this could be a VM on which we want to protect
	// all the SQL databases. The environment of the owner_entity could be
	// different than the 'type' field specified above.
	OwnerEntity *PrivateEntityProto `json:"ownerEntity,omitempty"`

	// If set to true, the owner will be unregistered for all the application
	// environments for which it was previously registered. It is caller's
	// responsibility to make sure that the 'owner_entity' is already a
	// registered source entity. If there are backup jobs that still protect an
	// appication environment on this owner, this RPC will return an error.
	UnregisterOwner *bool `json:"unregisterOwner,omitempty"`

	// If both 'is_internal_encrypted' and 'user_encryption_key' is set, it is
	// assumed that credentials are first encrypted using
	// 'internal_encryption_key' and then encrypted using 'user_encryption_key'.
	UserEncryptionKey *string `json:"userEncryptionKey,omitempty"`

	// Specifies information about the user who made the request.
	UserInfo *UserInformation `json:"userInfo,omitempty"`

	// Set to true if a persistent agent is running on the host. If this is
	// specified, then credentials would not be used to log into the host
	// environment. This mechanism may be used in environments such as VMware
	// to get around UAC permission issues by running the agent as a service
	// with the right credentials.
	UsesPersistentAgent *bool `json:"usesPersistentAgent,omitempty"`
}

// Validate validates this register or update app owner arg
func (m *RegisterOrUpdateAppOwnerArg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIRequestAttr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppCredentialsVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) validateAPIRequestAttr(formats strfmt.Registry) error {
	if swag.IsZero(m.APIRequestAttr) { // not required
		return nil
	}

	if m.APIRequestAttr != nil {
		if err := m.APIRequestAttr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiRequestAttr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiRequestAttr")
			}
			return err
		}
	}

	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) validateAPIVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.APIVersion) { // not required
		return nil
	}

	if m.APIVersion != nil {
		if err := m.APIVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiVersion")
			}
			return err
		}
	}

	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) validateAppCredentialsVec(formats strfmt.Registry) error {
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

func (m *RegisterOrUpdateAppOwnerArg) validateCredentials(formats strfmt.Registry) error {
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

func (m *RegisterOrUpdateAppOwnerArg) validateOwnerEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerEntity) { // not required
		return nil
	}

	if m.OwnerEntity != nil {
		if err := m.OwnerEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownerEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownerEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) validateUserInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UserInfo) { // not required
		return nil
	}

	if m.UserInfo != nil {
		if err := m.UserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this register or update app owner arg based on the context it is used
func (m *RegisterOrUpdateAppOwnerArg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIRequestAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppCredentialsVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwnerEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) contextValidateAPIRequestAttr(ctx context.Context, formats strfmt.Registry) error {

	if m.APIRequestAttr != nil {

		if swag.IsZero(m.APIRequestAttr) { // not required
			return nil
		}

		if err := m.APIRequestAttr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiRequestAttr")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiRequestAttr")
			}
			return err
		}
	}

	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) contextValidateAPIVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.APIVersion != nil {

		if swag.IsZero(m.APIVersion) { // not required
			return nil
		}

		if err := m.APIVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apiVersion")
			}
			return err
		}
	}

	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) contextValidateAppCredentialsVec(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RegisterOrUpdateAppOwnerArg) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RegisterOrUpdateAppOwnerArg) contextValidateOwnerEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.OwnerEntity != nil {

		if swag.IsZero(m.OwnerEntity) { // not required
			return nil
		}

		if err := m.OwnerEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownerEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownerEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RegisterOrUpdateAppOwnerArg) contextValidateUserInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UserInfo != nil {

		if swag.IsZero(m.UserInfo) { // not required
			return nil
		}

		if err := m.UserInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisterOrUpdateAppOwnerArg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterOrUpdateAppOwnerArg) UnmarshalBinary(b []byte) error {
	var res RegisterOrUpdateAppOwnerArg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
