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

// AadGraphNodeFilterParams Determines filter params that can be applied to query aad node.
//
// swagger:model AadGraphNodeFilterParams
type AadGraphNodeFilterParams struct {

	// Filter the nodes which matches with specified aad node type provided. Supported AAD node types - Users/Groups/Applications/ AdministrativeUnits/ServicePrincipals/DirectoryRoles/Contacts/ Devices
	// Enum: ["Unknown","AdminUnit","Application","Contact","Device","DirRole","Group","ServicePrincipal","AppRoleAssignment","User"]
	NodeType *string `json:"nodeType,omitempty"`
}

// Validate validates this aad graph node filter params
func (m *AadGraphNodeFilterParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var aadGraphNodeFilterParamsTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","AdminUnit","Application","Contact","Device","DirRole","Group","ServicePrincipal","AppRoleAssignment","User"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aadGraphNodeFilterParamsTypeNodeTypePropEnum = append(aadGraphNodeFilterParamsTypeNodeTypePropEnum, v)
	}
}

const (

	// AadGraphNodeFilterParamsNodeTypeUnknown captures enum value "Unknown"
	AadGraphNodeFilterParamsNodeTypeUnknown string = "Unknown"

	// AadGraphNodeFilterParamsNodeTypeAdminUnit captures enum value "AdminUnit"
	AadGraphNodeFilterParamsNodeTypeAdminUnit string = "AdminUnit"

	// AadGraphNodeFilterParamsNodeTypeApplication captures enum value "Application"
	AadGraphNodeFilterParamsNodeTypeApplication string = "Application"

	// AadGraphNodeFilterParamsNodeTypeContact captures enum value "Contact"
	AadGraphNodeFilterParamsNodeTypeContact string = "Contact"

	// AadGraphNodeFilterParamsNodeTypeDevice captures enum value "Device"
	AadGraphNodeFilterParamsNodeTypeDevice string = "Device"

	// AadGraphNodeFilterParamsNodeTypeDirRole captures enum value "DirRole"
	AadGraphNodeFilterParamsNodeTypeDirRole string = "DirRole"

	// AadGraphNodeFilterParamsNodeTypeGroup captures enum value "Group"
	AadGraphNodeFilterParamsNodeTypeGroup string = "Group"

	// AadGraphNodeFilterParamsNodeTypeServicePrincipal captures enum value "ServicePrincipal"
	AadGraphNodeFilterParamsNodeTypeServicePrincipal string = "ServicePrincipal"

	// AadGraphNodeFilterParamsNodeTypeAppRoleAssignment captures enum value "AppRoleAssignment"
	AadGraphNodeFilterParamsNodeTypeAppRoleAssignment string = "AppRoleAssignment"

	// AadGraphNodeFilterParamsNodeTypeUser captures enum value "User"
	AadGraphNodeFilterParamsNodeTypeUser string = "User"
)

// prop value enum
func (m *AadGraphNodeFilterParams) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aadGraphNodeFilterParamsTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AadGraphNodeFilterParams) validateNodeType(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNodeTypeEnum("nodeType", "body", *m.NodeType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aad graph node filter params based on context it is used
func (m *AadGraphNodeFilterParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AadGraphNodeFilterParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AadGraphNodeFilterParams) UnmarshalBinary(b []byte) error {
	var res AadGraphNodeFilterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
