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

// FilesStatsForEntity Specifies the files stats for an entity.
//
// swagger:model FilesStatsForEntity
type FilesStatsForEntity struct {

	// Specifies the cluster id of the entity.
	ClusterID *int64 `json:"clusterId,omitempty"`

	// Specifies the cluster incarnation id of the entity.
	ClusterIncarnationID *int64 `json:"clusterIncarnationId,omitempty"`

	// Specifies the entity id.
	EntityID *int64 `json:"entityId,omitempty"`

	// Specifies the entity name.
	EntityName *string `json:"entityName,omitempty"`

	// Specifies the entity type.
	// Enum: ["kCluster","kStorageDomain"]
	EntityType *string `json:"entityType,omitempty"`

	// Specifies a list of files stats for the entity.
	FilesStats []*FileStats `json:"filesStats"`
}

// Validate validates this files stats for entity
func (m *FilesStatsForEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var filesStatsForEntityTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kCluster","kStorageDomain"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filesStatsForEntityTypeEntityTypePropEnum = append(filesStatsForEntityTypeEntityTypePropEnum, v)
	}
}

const (

	// FilesStatsForEntityEntityTypeKCluster captures enum value "kCluster"
	FilesStatsForEntityEntityTypeKCluster string = "kCluster"

	// FilesStatsForEntityEntityTypeKStorageDomain captures enum value "kStorageDomain"
	FilesStatsForEntityEntityTypeKStorageDomain string = "kStorageDomain"
)

// prop value enum
func (m *FilesStatsForEntity) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, filesStatsForEntityTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FilesStatsForEntity) validateEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", *m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *FilesStatsForEntity) validateFilesStats(formats strfmt.Registry) error {
	if swag.IsZero(m.FilesStats) { // not required
		return nil
	}

	for i := 0; i < len(m.FilesStats); i++ {
		if swag.IsZero(m.FilesStats[i]) { // not required
			continue
		}

		if m.FilesStats[i] != nil {
			if err := m.FilesStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filesStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this files stats for entity based on the context it is used
func (m *FilesStatsForEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilesStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesStatsForEntity) contextValidateFilesStats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FilesStats); i++ {

		if m.FilesStats[i] != nil {

			if swag.IsZero(m.FilesStats[i]) { // not required
				return nil
			}

			if err := m.FilesStats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filesStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilesStatsForEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesStatsForEntity) UnmarshalBinary(b []byte) error {
	var res FilesStatsForEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
