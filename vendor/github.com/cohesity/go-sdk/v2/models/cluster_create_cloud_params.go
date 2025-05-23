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

// ClusterCreateCloudParams Cloud Cluster Params.
//
// # Params for Cloud Edition Cluster Creation
//
// swagger:model ClusterCreateCloudParams
type ClusterCreateCloudParams struct {

	// node ips
	// Required: true
	// Min Items: 1
	// Unique: true
	NodeIps []string `json:"nodeIps"`

	// Specifies the metadata fault tolerance.
	MetadataFaultTolerance *int32 `json:"metadataFaultTolerance,omitempty"`

	// Specifies the encryption configuration parameters
	EncryptionConfig *EncryptionConfigurationParams `json:"encryptionConfig,omitempty"`

	// Specifies IP preference
	IPPreference *int32 `json:"ipPreference,omitempty"`

	// Specifies Trust Domain used for Service Identity
	TrustDomain *string `json:"trustDomain,omitempty"`

	// Specifies whether or not to enable software encryption
	EnableCloudRf1 *bool `json:"enableCloudRf1,omitempty"`

	// Specifies the size of the cloud platforms.
	// Enum: ["Small","Medium","Large","XLarge","NextGen"]
	ClusterSize *string `json:"clusterSize,omitempty"`

	// Serial number of the disks to designate properties.
	DiskSerials []string `json:"diskSerials"`

	// Component exclusive property of the disks to designate.
	DiskComponentExclusive []string `json:"diskComponentExclusive"`

	// Self fault tolerant property of the disks to designate.
	DiskSelfFaultTolerant []bool `json:"diskSelfFaultTolerant"`

	// All nodes reachable property of the disks to designate.
	DiskAllNodesReachable []bool `json:"diskAllNodesReachable"`

	// Hostname of the cluster partition.
	ClusterPartitionHostname *string `json:"clusterPartitionHostname,omitempty"`
}

// Validate validates this cluster create cloud params
func (m *ClusterCreateCloudParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreateCloudParams) validateNodeIps(formats strfmt.Registry) error {

	if err := validate.Required("nodeIps", "body", m.NodeIps); err != nil {
		return err
	}

	iNodeIpsSize := int64(len(m.NodeIps))

	if err := validate.MinItems("nodeIps", "body", iNodeIpsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("nodeIps", "body", m.NodeIps); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreateCloudParams) validateEncryptionConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionConfig) { // not required
		return nil
	}

	if m.EncryptionConfig != nil {
		if err := m.EncryptionConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionConfig")
			}
			return err
		}
	}

	return nil
}

var clusterCreateCloudParamsTypeClusterSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Small","Medium","Large","XLarge","NextGen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterCreateCloudParamsTypeClusterSizePropEnum = append(clusterCreateCloudParamsTypeClusterSizePropEnum, v)
	}
}

const (

	// ClusterCreateCloudParamsClusterSizeSmall captures enum value "Small"
	ClusterCreateCloudParamsClusterSizeSmall string = "Small"

	// ClusterCreateCloudParamsClusterSizeMedium captures enum value "Medium"
	ClusterCreateCloudParamsClusterSizeMedium string = "Medium"

	// ClusterCreateCloudParamsClusterSizeLarge captures enum value "Large"
	ClusterCreateCloudParamsClusterSizeLarge string = "Large"

	// ClusterCreateCloudParamsClusterSizeXLarge captures enum value "XLarge"
	ClusterCreateCloudParamsClusterSizeXLarge string = "XLarge"

	// ClusterCreateCloudParamsClusterSizeNextGen captures enum value "NextGen"
	ClusterCreateCloudParamsClusterSizeNextGen string = "NextGen"
)

// prop value enum
func (m *ClusterCreateCloudParams) validateClusterSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterCreateCloudParamsTypeClusterSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterCreateCloudParams) validateClusterSize(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSize) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterSizeEnum("clusterSize", "body", *m.ClusterSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster create cloud params based on the context it is used
func (m *ClusterCreateCloudParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncryptionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreateCloudParams) contextValidateEncryptionConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionConfig != nil {

		if swag.IsZero(m.EncryptionConfig) { // not required
			return nil
		}

		if err := m.EncryptionConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreateCloudParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreateCloudParams) UnmarshalBinary(b []byte) error {
	var res ClusterCreateCloudParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
