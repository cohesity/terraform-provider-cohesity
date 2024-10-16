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
	"github.com/go-openapi/validate"
)

// HeliosOnPremConfig Helios OnPrem VM Config Params
//
// Params for Helios OnPrem VM Configuration.
//
// swagger:model HeliosOnPremConfig
type HeliosOnPremConfig struct {

	// Name of the new Helios OnPrem VM.
	// Required: true
	Name *string `json:"name"`

	// Specifies the ID of the Cluster.
	// Read Only: true
	ClusterID *int64 `json:"clusterId,omitempty"`

	// Specifies the Nodes present in this Cluster.
	Nodes []*HeliosOnPremVMNode `json:"nodes"`

	// Subnet to use for setting up the Kubernetes cluster's internal network on which Cohesity Helios will run.
	// Required: true
	KubernetesSubnetCidr *string `json:"kubernetesSubnetCidr"`

	// network config
	NetworkConfig *ClusterCreateNetworkConfig `json:"networkConfig,omitempty"`

	// proxy server config
	ProxyServerConfig *ClusterProxyServerConfig `json:"proxyServerConfig,omitempty"`

	// ssh config
	SSHConfig *HeliosOnPremSSHConfig `json:"sshConfig,omitempty"`
}

// Validate validates this helios on prem config
func (m *HeliosOnPremConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesSubnetCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosOnPremConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HeliosOnPremConfig) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HeliosOnPremConfig) validateKubernetesSubnetCidr(formats strfmt.Registry) error {

	if err := validate.Required("kubernetesSubnetCidr", "body", m.KubernetesSubnetCidr); err != nil {
		return err
	}

	return nil
}

func (m *HeliosOnPremConfig) validateNetworkConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkConfig) { // not required
		return nil
	}

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HeliosOnPremConfig) validateProxyServerConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyServerConfig) { // not required
		return nil
	}

	if m.ProxyServerConfig != nil {
		if err := m.ProxyServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyServerConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxyServerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HeliosOnPremConfig) validateSSHConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHConfig) { // not required
		return nil
	}

	if m.SSHConfig != nil {
		if err := m.SSHConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this helios on prem config based on the context it is used
func (m *HeliosOnPremConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxyServerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSSHConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HeliosOnPremConfig) contextValidateClusterID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *HeliosOnPremConfig) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {

			if swag.IsZero(m.Nodes[i]) { // not required
				return nil
			}

			if err := m.Nodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HeliosOnPremConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkConfig != nil {

		if swag.IsZero(m.NetworkConfig) { // not required
			return nil
		}

		if err := m.NetworkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HeliosOnPremConfig) contextValidateProxyServerConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ProxyServerConfig != nil {

		if swag.IsZero(m.ProxyServerConfig) { // not required
			return nil
		}

		if err := m.ProxyServerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyServerConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxyServerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HeliosOnPremConfig) contextValidateSSHConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SSHConfig != nil {

		if swag.IsZero(m.SSHConfig) { // not required
			return nil
		}

		if err := m.SSHConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HeliosOnPremConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeliosOnPremConfig) UnmarshalBinary(b []byte) error {
	var res HeliosOnPremConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
