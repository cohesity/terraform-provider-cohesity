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

// AzureAgentProtectionGroupParams Agent based Azure Protection Group Request Params.
//
// Specifies the parameters which are specific to Azure related Protection Groups using cohesity protection-service installed on the instance. Objects must be specified.
//
// swagger:model AzureAgentProtectionGroupParams
type AzureAgentProtectionGroupParams struct {

	// Specifies the objects to be included in the Protection Group.
	// Required: true
	// Min Items: 1
	// Unique: true
	Objects []*AzureAgentProtectionGroupObjectParams `json:"objects"`

	// Specifies the objects to be excluded in the Protection Group.
	// Unique: true
	ExcludeObjectIds []int64 `json:"excludeObjectIds"`

	// Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots.
	AppConsistentSnapshot *bool `json:"appConsistentSnapshot,omitempty"`

	// Specifies the indexing policy for VMs in this Group.
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`

	// Specifies whether or not to move the workload to the cloud.
	CloudMigration *bool `json:"cloudMigration,omitempty"`

	// Specifies the id of the parent of the objects.
	// Read Only: true
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies the name of the parent of the objects.
	// Read Only: true
	SourceName *string `json:"sourceName,omitempty"`
}

// Validate validates this azure agent protection group params
func (m *AzureAgentProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludeObjectIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureAgentProtectionGroupParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
		return err
	}

	iObjectsSize := int64(len(m.Objects))

	if err := validate.MinItems("objects", "body", iObjectsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("objects", "body", m.Objects); err != nil {
		return err
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AzureAgentProtectionGroupParams) validateExcludeObjectIds(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeObjectIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("excludeObjectIds", "body", m.ExcludeObjectIds); err != nil {
		return err
	}

	return nil
}

func (m *AzureAgentProtectionGroupParams) validateIndexingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexingPolicy) { // not required
		return nil
	}

	if m.IndexingPolicy != nil {
		if err := m.IndexingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("indexingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("indexingPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure agent protection group params based on the context it is used
func (m *AzureAgentProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndexingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureAgentProtectionGroupParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AzureAgentProtectionGroupParams) contextValidateIndexingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IndexingPolicy != nil {

		if swag.IsZero(m.IndexingPolicy) { // not required
			return nil
		}

		if err := m.IndexingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("indexingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("indexingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *AzureAgentProtectionGroupParams) contextValidateSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *AzureAgentProtectionGroupParams) contextValidateSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceName", "body", m.SourceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureAgentProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureAgentProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res AzureAgentProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
