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

// HdfsProtectionGroupParams Specifies the parameters for HDFS Protection Group.
//
// swagger:model HdfsProtectionGroupParams
type HdfsProtectionGroupParams struct {

	// Specifies the paths to be included in the Protection Group.
	// Min Items: 1
	IncludePaths []string `json:"includePaths"`

	// Specifies the paths to be excluded in the Protection Group. excludePaths will ovrride includePaths.
	ExcludePaths []string `json:"excludePaths"`

	// Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster.
	Concurrency *int32 `json:"concurrency,omitempty"`

	// Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster.
	BandwidthMBPS *int64 `json:"bandwidthMBPS,omitempty"`

	// The object ID of the HDFS source for this protection group.
	// Required: true
	HdfsSourceID *int64 `json:"hdfsSourceId"`

	// Object ID of the Source on which this protection was run .
	// Read Only: true
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies settings for indexing files found in an Object so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`

	// Specifies the name of the Source on which this protection was run.
	// Read Only: true
	SourceName *string `json:"sourceName,omitempty"`
}

// Validate validates this hdfs protection group params
func (m *HdfsProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncludePaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdfsSourceID(formats); err != nil {
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

func (m *HdfsProtectionGroupParams) validateIncludePaths(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludePaths) { // not required
		return nil
	}

	iIncludePathsSize := int64(len(m.IncludePaths))

	if err := validate.MinItems("includePaths", "body", iIncludePathsSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *HdfsProtectionGroupParams) validateHdfsSourceID(formats strfmt.Registry) error {

	if err := validate.Required("hdfsSourceId", "body", m.HdfsSourceID); err != nil {
		return err
	}

	return nil
}

func (m *HdfsProtectionGroupParams) validateIndexingPolicy(formats strfmt.Registry) error {
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

// ContextValidate validate this hdfs protection group params based on the context it is used
func (m *HdfsProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndexingPolicy(ctx, formats); err != nil {
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

func (m *HdfsProtectionGroupParams) contextValidateSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *HdfsProtectionGroupParams) contextValidateIndexingPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HdfsProtectionGroupParams) contextValidateSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceName", "body", m.SourceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HdfsProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HdfsProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res HdfsProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
