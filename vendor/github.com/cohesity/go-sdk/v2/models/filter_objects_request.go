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

// FilterObjectsRequest Specifies the filter details.
//
// swagger:model FilterObjectsRequest
type FilterObjectsRequest struct {

	// Specifies the type of filtering user wants to perform. Currently, we only support exclude type of filter.
	// Required: true
	// Enum: ["exclude"]
	FilterType *string `json:"filterType"`

	// Specifies the list of filters that need to be applied on given list of discovered objects.
	// Required: true
	// Min Items: 1
	Filters []*Filter `json:"filters"`

	// Specifies a list of non leaf object ids to filter the leaf level objects. Non leaf object such host (physical or vm) or database instance can be specified.
	// Required: true
	// Min Items: 1
	// Unique: true
	ObjectIds []int64 `json:"objectIds"`

	// Specifies the type of application enviornment needed for filtering to be applied on. This is needed because in case of applications like SQL, Oracle, a single source can contain multiple application enviornments.
	// Enum: ["kSQL"]
	ApplicationEnvironment *string `json:"applicationEnvironment,omitempty"`

	// TenantIds contains list of the tenant for which objects are to be returned.
	TenantIds []string `json:"tenantIds"`

	// If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false.
	IncludeTenants *bool `json:"includeTenants,omitempty"`
}

// Validate validates this filter objects request
func (m *FilterObjectsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var filterObjectsRequestTypeFilterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["exclude"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filterObjectsRequestTypeFilterTypePropEnum = append(filterObjectsRequestTypeFilterTypePropEnum, v)
	}
}

const (

	// FilterObjectsRequestFilterTypeExclude captures enum value "exclude"
	FilterObjectsRequestFilterTypeExclude string = "exclude"
)

// prop value enum
func (m *FilterObjectsRequest) validateFilterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, filterObjectsRequestTypeFilterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FilterObjectsRequest) validateFilterType(formats strfmt.Registry) error {

	if err := validate.Required("filterType", "body", m.FilterType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFilterTypeEnum("filterType", "body", *m.FilterType); err != nil {
		return err
	}

	return nil
}

func (m *FilterObjectsRequest) validateFilters(formats strfmt.Registry) error {

	if err := validate.Required("filters", "body", m.Filters); err != nil {
		return err
	}

	iFiltersSize := int64(len(m.Filters))

	if err := validate.MinItems("filters", "body", iFiltersSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FilterObjectsRequest) validateObjectIds(formats strfmt.Registry) error {

	if err := validate.Required("objectIds", "body", m.ObjectIds); err != nil {
		return err
	}

	iObjectIdsSize := int64(len(m.ObjectIds))

	if err := validate.MinItems("objectIds", "body", iObjectIdsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("objectIds", "body", m.ObjectIds); err != nil {
		return err
	}

	return nil
}

var filterObjectsRequestTypeApplicationEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kSQL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filterObjectsRequestTypeApplicationEnvironmentPropEnum = append(filterObjectsRequestTypeApplicationEnvironmentPropEnum, v)
	}
}

const (

	// FilterObjectsRequestApplicationEnvironmentKSQL captures enum value "kSQL"
	FilterObjectsRequestApplicationEnvironmentKSQL string = "kSQL"
)

// prop value enum
func (m *FilterObjectsRequest) validateApplicationEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, filterObjectsRequestTypeApplicationEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FilterObjectsRequest) validateApplicationEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationEnvironment) { // not required
		return nil
	}

	// value enum
	if err := m.validateApplicationEnvironmentEnum("applicationEnvironment", "body", *m.ApplicationEnvironment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this filter objects request based on the context it is used
func (m *FilterObjectsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilterObjectsRequest) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {

			if swag.IsZero(m.Filters[i]) { // not required
				return nil
			}

			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilterObjectsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterObjectsRequest) UnmarshalBinary(b []byte) error {
	var res FilterObjectsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
