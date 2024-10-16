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

// ClusterStateParams Specifies the current cluster state details.
//
// swagger:model ClusterStateParams
type ClusterStateParams struct {

	// Specifies the id of the cluster.
	ClusterID int64 `json:"clusterId,omitempty"`

	// Specifies the incarnation id of the cluster.
	ClusterIncarnationID int64 `json:"clusterIncarnationId,omitempty"`

	// Specifies the name of the cluster.
	Name string `json:"name,omitempty"`

	// Specifies the software version of the cluster.
	SoftwareVersion string `json:"softwareVersion,omitempty"`

	// Specifies the operations running on the cluster. 'None' indicates that there are no operations currently running on the cluster. 'Destroy' indicates that the cluster is currently being destroyed. 'Upgrade' indicates that the cluster is currently being upgraded. 'Clean' indicates that the cluster is being cleaned. 'NodeRemoval' indicates that a node is being removed from the cluster. 'DiskRemoval' indicates that a disk is being removed from the cluster. 'DiskAddition' indicates that a disk is being added tos the cluster. 'UploadPackageByUrl' indicates that a package is being uploaded using a URL. 'UploadPackageAndUpgrade' indicates package upload by URL and upgrade operation. 'BaseOSUpgrade' indicates that the BaseOSUpgrade operation on the cluster is set. 'ServiceRestart' indicates that the services on the Cluster are currently being restarted. 'SystemServiceRestart' indicates that system services on the Cluster are currently being restarted.
	Operations []string `json:"operations,omitempty"`

	// Specifies the details of each system app state on the cluster.
	SystemApps []*SystemAppStatusParams `json:"systemApps,omitempty"`
}

// Validate validates this cluster state params
func (m *ClusterStateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemApps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterStateParamsOperationsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Destroy","Upgrade","Clean","NodeRemoval","DiskRemoval","DiskAddition","NodeAddition","UploadPackageByUrl","UploadPackageAndUpgrade","BaseOSUpgrade","ServiceRestart","SystemServiceRestart"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterStateParamsOperationsItemsEnum = append(clusterStateParamsOperationsItemsEnum, v)
	}
}

func (m *ClusterStateParams) validateOperationsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterStateParamsOperationsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterStateParams) validateOperations(formats strfmt.Registry) error {
	if swag.IsZero(m.Operations) { // not required
		return nil
	}

	for i := 0; i < len(m.Operations); i++ {

		// value enum
		if err := m.validateOperationsItemsEnum("operations"+"."+strconv.Itoa(i), "body", m.Operations[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ClusterStateParams) validateSystemApps(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemApps) { // not required
		return nil
	}

	for i := 0; i < len(m.SystemApps); i++ {
		if swag.IsZero(m.SystemApps[i]) { // not required
			continue
		}

		if m.SystemApps[i] != nil {
			if err := m.SystemApps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemApps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("systemApps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster state params based on the context it is used
func (m *ClusterStateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSystemApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStateParams) contextValidateSystemApps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SystemApps); i++ {

		if m.SystemApps[i] != nil {

			if swag.IsZero(m.SystemApps[i]) { // not required
				return nil
			}

			if err := m.SystemApps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("systemApps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("systemApps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStateParams) UnmarshalBinary(b []byte) error {
	var res ClusterStateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
