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

// MongodbNewOpsManagerConfig MongoDB Ops Manager Config.
//
// Specifies the new destination Source configuration where the Mongodb cluster will be recovered.
//
// swagger:model MongodbNewOpsManagerConfig
type MongodbNewOpsManagerConfig struct {

	// Specifies the id of the parent source to recover.
	// Required: true
	SourceID *int64 `json:"sourceId"`

	// Specifies the Organization id.
	// Required: true
	OrgID *string `json:"orgId"`

	// Specifies the Project id.
	// Required: true
	ProjectID *string `json:"projectId"`

	// Specifies the Cluster id.
	// Required: true
	ClusterID *string `json:"clusterId"`
}

// Validate validates this mongodb new ops manager config
func (m *MongodbNewOpsManagerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongodbNewOpsManagerConfig) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *MongodbNewOpsManagerConfig) validateOrgID(formats strfmt.Registry) error {

	if err := validate.Required("orgId", "body", m.OrgID); err != nil {
		return err
	}

	return nil
}

func (m *MongodbNewOpsManagerConfig) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *MongodbNewOpsManagerConfig) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mongodb new ops manager config based on context it is used
func (m *MongodbNewOpsManagerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongodbNewOpsManagerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongodbNewOpsManagerConfig) UnmarshalBinary(b []byte) error {
	var res MongodbNewOpsManagerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
