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

// GetGraphNodeRelationsDiffParams Specify the query params to determine difference of node relation between two snapshots for a given node id.
//
// swagger:model GetGraphNodeRelationsDiffParams
type GetGraphNodeRelationsDiffParams struct {

	// If set to false only the diff of node info will be returned else the diff of relations matching the below edge filters will also be returned. Defaults to false.
	DiffRelation *bool `json:"diffRelation,omitempty"`

	// Specifies the edges filter params for a graph node.
	RelationFilter struct {
		GraphNodeRelationFilterParams
	} `json:"relationFilter,omitempty"`

	// Specifies an optional mask to filter only certain kinds of diffs. Supported diff types - Added/Modified/Deleted/Unmodified
	// Unique: true
	DiffTypes []string `json:"diffTypes"`
}

// Validate validates this get graph node relations diff params
func (m *GetGraphNodeRelationsDiffParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelationFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiffTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetGraphNodeRelationsDiffParams) validateRelationFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.RelationFilter) { // not required
		return nil
	}

	return nil
}

var getGraphNodeRelationsDiffParamsDiffTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Added","Modified","Deleted","Unmodified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getGraphNodeRelationsDiffParamsDiffTypesItemsEnum = append(getGraphNodeRelationsDiffParamsDiffTypesItemsEnum, v)
	}
}

func (m *GetGraphNodeRelationsDiffParams) validateDiffTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getGraphNodeRelationsDiffParamsDiffTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GetGraphNodeRelationsDiffParams) validateDiffTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.DiffTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("diffTypes", "body", m.DiffTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.DiffTypes); i++ {

		// value enum
		if err := m.validateDiffTypesItemsEnum("diffTypes"+"."+strconv.Itoa(i), "body", m.DiffTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this get graph node relations diff params based on the context it is used
func (m *GetGraphNodeRelationsDiffParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRelationFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetGraphNodeRelationsDiffParams) contextValidateRelationFilter(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetGraphNodeRelationsDiffParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetGraphNodeRelationsDiffParams) UnmarshalBinary(b []byte) error {
	var res GetGraphNodeRelationsDiffParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
