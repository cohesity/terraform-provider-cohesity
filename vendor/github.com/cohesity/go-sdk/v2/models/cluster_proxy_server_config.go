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

// ClusterProxyServerConfig Cluster Proxy Server Config.
//
// Specifies the parameters of the proxy server to be used for external traffic.
//
// swagger:model ClusterProxyServerConfig
type ClusterProxyServerConfig struct {

	// Specifies the unique name of the proxy server.
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Specifies the IP address of the proxy server.
	// Required: true
	IP *string `json:"ip"`

	// Specifies the port on which the server is listening.
	// Required: true
	Port *int32 `json:"port"`

	// Specifies the username for the proxy.
	Username *string `json:"username,omitempty"`

	// Specifies the password for the proxy.
	Password *string `json:"password,omitempty"`

	// Disable proxy is used to turn the proxy on and off.
	IsDisabled *bool `json:"isDisabled,omitempty"`
}

// Validate validates this cluster proxy server config
func (m *ClusterProxyServerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterProxyServerConfig) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *ClusterProxyServerConfig) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster proxy server config based on the context it is used
func (m *ClusterProxyServerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterProxyServerConfig) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterProxyServerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterProxyServerConfig) UnmarshalBinary(b []byte) error {
	var res ClusterProxyServerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
