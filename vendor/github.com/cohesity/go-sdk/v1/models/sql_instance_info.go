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

// SQLInstanceInfo Information about a SQL Server instance. The fields correspond to the
// server properties as returned by the SERVERPROPERTY built-in function.
// See https://goo.gl/8amCft for details.
//
// swagger:model SQLInstanceInfo
type SQLInstanceInfo struct {

	// The product build type information of the SQL instance.
	BuildTypeInfo *SQLInstanceInfoProductBuildTypeInfo `json:"buildTypeInfo,omitempty"`

	// Populated if the instance is part of a Windows failover cluster. This
	// contains the virtual endpoints assigned to the instance by the cluster.
	ClusterResourceEndpointVec []*ClusterNetworkingInfoEndpoint `json:"clusterResourceEndpointVec"`

	// Default data files location for the SQL instance
	DefaultDataLocation *string `json:"defaultDataLocation,omitempty"`

	// Default log files location for the SQL instance
	DefaultLogLocation *string `json:"defaultLogLocation,omitempty"`

	// The edition information of the SQL Instance.
	EditionInfo *SQLInstanceInfoEditionInfo `json:"editionInfo,omitempty"`

	// Populated if instance is FCI. List of nodes that are Part of the FCI.
	FciNodeNameVec []string `json:"fciNodeNameVec"`

	// This field denotes the resource GUID of the SQL Server resource of this
	// FCI instance in a windows cluster.
	// The GUID is set by the resource data recieved from GetClusteredSQLInfo,
	// which reads the resource id of this FCI resource from the 'Cluster'
	// registry hive under HKLM.
	// eg: "f9332cdf-84c2-4b7c-961c-2fcbaae690f8"
	// This is only applicable if is_fci_instance is set to true.
	GUIDInWindowsCluster *string `json:"guidInWindowsCluster,omitempty"`

	// The id of the SQL server instance.
	InstanceID []uint8 `json:"instanceId"`

	// The name of the SQL server instance.
	InstanceName []uint8 `json:"instanceName"`

	// Specifies if the instance is part of a SQL cluster (FCI).
	IsFciInstance *bool `json:"isFciInstance,omitempty"`

	// Whether the SQL instance is online.
	IsOnline *bool `json:"isOnline,omitempty"`

	// The product level information of the SQL instance.
	LevelInfo *SQLInstanceInfoProductLevelInfo `json:"levelInfo,omitempty"`

	// The product update level information of the SQL instance.
	UpdateLevelInfo *SQLInstanceInfoProductUpdateLevelInfo `json:"updateLevelInfo,omitempty"`

	// The version of the SQL instance.
	Version *SQLServerInstanceVersion `json:"version,omitempty"`
}

// Validate validates this SQL instance info
func (m *SQLInstanceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildTypeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterResourceEndpointVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEditionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevelInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateLevelInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLInstanceInfo) validateBuildTypeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildTypeInfo) { // not required
		return nil
	}

	if m.BuildTypeInfo != nil {
		if err := m.BuildTypeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildTypeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buildTypeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) validateClusterResourceEndpointVec(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterResourceEndpointVec) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterResourceEndpointVec); i++ {
		if swag.IsZero(m.ClusterResourceEndpointVec[i]) { // not required
			continue
		}

		if m.ClusterResourceEndpointVec[i] != nil {
			if err := m.ClusterResourceEndpointVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterResourceEndpointVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterResourceEndpointVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SQLInstanceInfo) validateEditionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.EditionInfo) { // not required
		return nil
	}

	if m.EditionInfo != nil {
		if err := m.EditionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("editionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("editionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) validateLevelInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.LevelInfo) { // not required
		return nil
	}

	if m.LevelInfo != nil {
		if err := m.LevelInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("levelInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("levelInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) validateUpdateLevelInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateLevelInfo) { // not required
		return nil
	}

	if m.UpdateLevelInfo != nil {
		if err := m.UpdateLevelInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateLevelInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateLevelInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this SQL instance info based on the context it is used
func (m *SQLInstanceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuildTypeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterResourceEndpointVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEditionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLevelInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateLevelInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLInstanceInfo) contextValidateBuildTypeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BuildTypeInfo != nil {

		if swag.IsZero(m.BuildTypeInfo) { // not required
			return nil
		}

		if err := m.BuildTypeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildTypeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buildTypeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) contextValidateClusterResourceEndpointVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterResourceEndpointVec); i++ {

		if m.ClusterResourceEndpointVec[i] != nil {

			if swag.IsZero(m.ClusterResourceEndpointVec[i]) { // not required
				return nil
			}

			if err := m.ClusterResourceEndpointVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterResourceEndpointVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterResourceEndpointVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SQLInstanceInfo) contextValidateEditionInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.EditionInfo != nil {

		if swag.IsZero(m.EditionInfo) { // not required
			return nil
		}

		if err := m.EditionInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("editionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("editionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) contextValidateLevelInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.LevelInfo != nil {

		if swag.IsZero(m.LevelInfo) { // not required
			return nil
		}

		if err := m.LevelInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("levelInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("levelInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) contextValidateUpdateLevelInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdateLevelInfo != nil {

		if swag.IsZero(m.UpdateLevelInfo) { // not required
			return nil
		}

		if err := m.UpdateLevelInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateLevelInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateLevelInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SQLInstanceInfo) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {

		if swag.IsZero(m.Version) { // not required
			return nil
		}

		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SQLInstanceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLInstanceInfo) UnmarshalBinary(b []byte) error {
	var res SQLInstanceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
