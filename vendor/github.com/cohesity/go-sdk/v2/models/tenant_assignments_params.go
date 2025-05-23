// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenantAssignmentsParams TenantAssignmentsParams
//
// Parameters to be specified for assigning properties like storage domain,
// entities, policies to the tenant.
//
// swagger:model TenantAssignmentsParams
type TenantAssignmentsParams struct {

	// List of storage domains on the cluster, to be associated to the Tenant.
	StorageDomainIds []int64 `json:"storageDomainIds"`

	// List of objects on the cluster, to be associated to the Tenant.
	ObjectIds []int64 `json:"objectIds"`

	// List of vlans on the cluster, to be associated to the Tenant.
	VlanIfaceNames []string `json:"vlanIfaceNames"`

	// List of views on the cluster, to be associated to the Tenant.
	ViewIds []int64 `json:"viewIds"`

	// List of policies on the cluster, to be associated to the Tenant.
	PolicyIds []string `json:"policyIds"`
}

// Validate validates this tenant assignments params
func (m *TenantAssignmentsParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tenant assignments params based on context it is used
func (m *TenantAssignmentsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TenantAssignmentsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantAssignmentsParams) UnmarshalBinary(b []byte) error {
	var res TenantAssignmentsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
