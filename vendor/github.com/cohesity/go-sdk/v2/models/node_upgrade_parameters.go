// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeUpgradeParameters Node Upgrade Parameters
//
// Parameters for Node Upgrade.
//
// swagger:model NodeUpgradeParameters
type NodeUpgradeParameters struct {

	// Specifies the target software version. The node that
	// the request is sent to will search itself for the specified software
	// package and if that package is found, it will be used for the
	// upgrade.
	//
	// Required: true
	TargetSwVersion *string `json:"targetSwVersion"`

	// Specifies that the node that the request is being sent
	// to should be upgraded. By default, this is set to true.
	//
	UpgradeSelf *bool `json:"upgradeSelf,omitempty"`

	// Specifies a list of IDs of additional nodes to be
	// upgraded. These must be free Nodes present on the same local network
	// as the Node that the request was sent to. The ID of the Node the
	// request was sent to should not be included in this list. This
	// parameter can only be specified if upgradeAllFreeNodes is not specified.
	//
	NodeIds []int64 `json:"nodeIds"`

	// Specifies whether or not to attempt to upgrade all free nodes which
	// are currently connected to the same local network as the node that the
	// request was sent to. This parameter can only be specified if nodeIds
	// is not specified.
	//
	UpgradeAllFreeNodes *bool `json:"upgradeAllFreeNodes,omitempty"`
}

// Validate validates this node upgrade parameters
func (m *NodeUpgradeParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetSwVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeUpgradeParameters) validateTargetSwVersion(formats strfmt.Registry) error {

	if err := validate.Required("targetSwVersion", "body", m.TargetSwVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this node upgrade parameters based on context it is used
func (m *NodeUpgradeParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeUpgradeParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeUpgradeParameters) UnmarshalBinary(b []byte) error {
	var res NodeUpgradeParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
