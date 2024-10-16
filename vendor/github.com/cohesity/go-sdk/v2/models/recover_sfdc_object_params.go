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

// RecoverSfdcObjectParams Recover Salesforce object params.
//
// Specifies the parameters to recover Salesforce objects.
//
// swagger:model RecoverSfdcObjectParams
type RecoverSfdcObjectParams struct {

	// Specifies the name of the object to be restored.
	ObjectName *string `json:"objectName,omitempty"`

	// Specifies a list of records IDs to be recovered for the specified Object.
	Records []string `json:"records"`

	// Specifies a list of parent object IDs to include in recovery. Specified parent objects will also be recovered as part of this recovery.
	ParentObjectIds []string `json:"parentObjectIds"`

	// Specifies a list of child object IDs to include in the recovery. Specified object IDs will also be recovered as part of this recovery.
	ChildObjectIds []string `json:"childObjectIds"`

	// Specifies a Query to filter the records. This filtered list of records will be used for recovery.
	FilterQuery *string `json:"filterQuery,omitempty"`

	// Specifies whether to include deleted Objects in the recovery.
	// Required: true
	IncludeDeletedObjects *bool `json:"includeDeletedObjects"`

	// Specifies a list of mutuation types for an object. Mutation type is required in conjunction with 'filterQuery'.
	// Unique: true
	MutationTypes []string `json:"mutationTypes"`
}

// Validate validates this recover sfdc object params
func (m *RecoverSfdcObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncludeDeletedObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMutationTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverSfdcObjectParams) validateIncludeDeletedObjects(formats strfmt.Registry) error {

	if err := validate.Required("includeDeletedObjects", "body", m.IncludeDeletedObjects); err != nil {
		return err
	}

	return nil
}

var recoverSfdcObjectParamsMutationTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["All","Added","Removed","Changed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverSfdcObjectParamsMutationTypesItemsEnum = append(recoverSfdcObjectParamsMutationTypesItemsEnum, v)
	}
}

func (m *RecoverSfdcObjectParams) validateMutationTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverSfdcObjectParamsMutationTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverSfdcObjectParams) validateMutationTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.MutationTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("mutationTypes", "body", m.MutationTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.MutationTypes); i++ {

		// value enum
		if err := m.validateMutationTypesItemsEnum("mutationTypes"+"."+strconv.Itoa(i), "body", m.MutationTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this recover sfdc object params based on context it is used
func (m *RecoverSfdcObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoverSfdcObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverSfdcObjectParams) UnmarshalBinary(b []byte) error {
	var res RecoverSfdcObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
