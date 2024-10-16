// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdRootTopologyObject AD Root Topology Object.
//
// Represents a node in AD Topology tree.
//
// swagger:model AdRootTopologyObject
type AdRootTopologyObject struct {

	// Specifies the array of children of this object.
	ChildObjects []interface{} `json:"childObjects"`

	// Specifies the 'description' of an object.
	Description *string `json:"description,omitempty"`

	// Specifies the guid of matching 'source_guid' from production AD.
	// This is looked up  based on source_guid or distinguishedName
	// attribute value.
	DestGUID *string `json:"destGuid,omitempty"`

	// Specifies the display name of the object in AD Topology tree.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies the distinguished name of the object in AD Topology tree.
	// Eg: CN=Jone Doe,OU=Users,DC=corp,DC=cohesity,DC=com
	DistinguishedName *string `json:"distinguishedName,omitempty"`

	// Specifies the AD error while fetching the ADRootTopologyObject.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Specifies the LDAP class name such as 'user','computer',
	// 'organizationalUnit'.
	ObjectClass *string `json:"objectClass,omitempty"`

	// Specifies the guid string of the object in AD snapshot database.
	SourceGUID *string `json:"sourceGuid,omitempty"`
}

// Validate validates this ad root topology object
func (m *AdRootTopologyObject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ad root topology object based on context it is used
func (m *AdRootTopologyObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdRootTopologyObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdRootTopologyObject) UnmarshalBinary(b []byte) error {
	var res AdRootTopologyObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
