// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExchangeDAGProtectionPreference Exchange DAG Backup preference
//
// Specifies the information about the preference order while choosing
// between which database copy of the database which is part of DAG should
// be protected.
//
// swagger:model ExchangeDAGProtectionPreference
type ExchangeDAGProtectionPreference struct {

	// Specifies the preference order of the exchange servers from which passive
	// database copies should be protected.
	// The preference order is descending which indicates that passive database
	// copy in the first server in the list gets the highest preference.
	PassiveCopyPreferenceServerGUIDList []string `json:"passiveCopyPreferenceServerGuidList"`

	// Specifies that only passive database copies should be protected if
	// this is set to true.
	// If this is set to false, both active and passive database copies can be
	// protected.
	PassiveOnly *bool `json:"passiveOnly,omitempty"`

	// Specifies to use the user specified preference order of exchange servers
	// from which the passive database copies should be protected if this is
	// set to true.
	//
	// Every copy of an Exchange database in a DAG is assigned an activation
	// preference number. This number is used by the system as part of the
	// passive database activation process.
	// If this bool flag is set to false, the reverse order of activation
	// is used while choosing between passive copies.
	UseUserSpecifiedPassivePreferenceOrder *bool `json:"useUserSpecifiedPassivePreferenceOrder,omitempty"`
}

// Validate validates this exchange d a g protection preference
func (m *ExchangeDAGProtectionPreference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exchange d a g protection preference based on context it is used
func (m *ExchangeDAGProtectionPreference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeDAGProtectionPreference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeDAGProtectionPreference) UnmarshalBinary(b []byte) error {
	var res ExchangeDAGProtectionPreference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
