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

// ArchivalNasExternalTargetParams Nas External Target Params for archival purpose type.
//
// Specifies the parameters which are specific to Nas related External Targets of archival purpose type.
//
// swagger:model ArchivalNasExternalTargetParams
type ArchivalNasExternalTargetParams struct {

	// Specifies the host of the NAS external target.
	// Required: true
	Host *string `json:"host"`

	// Specifies the mount path of the NAS external target.
	// Required: true
	MountPath *string `json:"mountPath"`

	// Specifies the share type of the NAS external target.
	// Read Only: true
	// Enum: ["CIFS","NFS"]
	ShareType *string `json:"shareType,omitempty"`

	// Specifies the Source Side Deduplication setting for the Nas external target
	SourceSideDeduplication *bool `json:"sourceSideDeduplication,omitempty"`

	// Specifies if Incremental Archival setting is enabled or not.
	IsIncrementalArchivalEnabled *bool `json:"isIncrementalArchivalEnabled,omitempty"`

	// Specifies if Forever Incremental Archival setting is enabled or not.
	IsForeverIncrementalArchivalEnabled *bool `json:"isForeverIncrementalArchivalEnabled,omitempty"`

	// Specifies the NFS version number of the target.
	// Enum: ["NFSv3","NFSv4","NFSv4_0","NFSv4_1","NFSv4_2"]
	NfsVersionNumber *string `json:"nfsVersionNumber,omitempty"`

	// Specifies the NFS security type of the target.
	// Enum: ["Default","None","System","KRB5","KRB5I","KRB5P"]
	NfsSecurityType *string `json:"nfsSecurityType,omitempty"`

	// Specifies the Kerberos realm name for a Kerberos-secured target.
	KerberosRealmName *string `json:"kerberosRealmName,omitempty"`

	// Specifies whether the garbage collection mode is network optimized or storage optimized. If this field is set to true, it refers to network optimized GC and if set to false, it refers to storage optimized GC.
	IsNetworkOptimizedGC *bool `json:"isNetworkOptimizedGC,omitempty"`
}

// Validate validates this archival nas external target params
func (m *ArchivalNasExternalTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShareType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsVersionNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsSecurityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArchivalNasExternalTargetParams) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *ArchivalNasExternalTargetParams) validateMountPath(formats strfmt.Registry) error {

	if err := validate.Required("mountPath", "body", m.MountPath); err != nil {
		return err
	}

	return nil
}

var archivalNasExternalTargetParamsTypeShareTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CIFS","NFS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archivalNasExternalTargetParamsTypeShareTypePropEnum = append(archivalNasExternalTargetParamsTypeShareTypePropEnum, v)
	}
}

const (

	// ArchivalNasExternalTargetParamsShareTypeCIFS captures enum value "CIFS"
	ArchivalNasExternalTargetParamsShareTypeCIFS string = "CIFS"

	// ArchivalNasExternalTargetParamsShareTypeNFS captures enum value "NFS"
	ArchivalNasExternalTargetParamsShareTypeNFS string = "NFS"
)

// prop value enum
func (m *ArchivalNasExternalTargetParams) validateShareTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archivalNasExternalTargetParamsTypeShareTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchivalNasExternalTargetParams) validateShareType(formats strfmt.Registry) error {
	if swag.IsZero(m.ShareType) { // not required
		return nil
	}

	// value enum
	if err := m.validateShareTypeEnum("shareType", "body", *m.ShareType); err != nil {
		return err
	}

	return nil
}

var archivalNasExternalTargetParamsTypeNfsVersionNumberPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NFSv3","NFSv4","NFSv4_0","NFSv4_1","NFSv4_2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archivalNasExternalTargetParamsTypeNfsVersionNumberPropEnum = append(archivalNasExternalTargetParamsTypeNfsVersionNumberPropEnum, v)
	}
}

const (

	// ArchivalNasExternalTargetParamsNfsVersionNumberNFSv3 captures enum value "NFSv3"
	ArchivalNasExternalTargetParamsNfsVersionNumberNFSv3 string = "NFSv3"

	// ArchivalNasExternalTargetParamsNfsVersionNumberNFSv4 captures enum value "NFSv4"
	ArchivalNasExternalTargetParamsNfsVersionNumberNFSv4 string = "NFSv4"

	// ArchivalNasExternalTargetParamsNfsVersionNumberNFSv40 captures enum value "NFSv4_0"
	ArchivalNasExternalTargetParamsNfsVersionNumberNFSv40 string = "NFSv4_0"

	// ArchivalNasExternalTargetParamsNfsVersionNumberNFSv41 captures enum value "NFSv4_1"
	ArchivalNasExternalTargetParamsNfsVersionNumberNFSv41 string = "NFSv4_1"

	// ArchivalNasExternalTargetParamsNfsVersionNumberNFSv42 captures enum value "NFSv4_2"
	ArchivalNasExternalTargetParamsNfsVersionNumberNFSv42 string = "NFSv4_2"
)

// prop value enum
func (m *ArchivalNasExternalTargetParams) validateNfsVersionNumberEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archivalNasExternalTargetParamsTypeNfsVersionNumberPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchivalNasExternalTargetParams) validateNfsVersionNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsVersionNumber) { // not required
		return nil
	}

	// value enum
	if err := m.validateNfsVersionNumberEnum("nfsVersionNumber", "body", *m.NfsVersionNumber); err != nil {
		return err
	}

	return nil
}

var archivalNasExternalTargetParamsTypeNfsSecurityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Default","None","System","KRB5","KRB5I","KRB5P"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archivalNasExternalTargetParamsTypeNfsSecurityTypePropEnum = append(archivalNasExternalTargetParamsTypeNfsSecurityTypePropEnum, v)
	}
}

const (

	// ArchivalNasExternalTargetParamsNfsSecurityTypeDefault captures enum value "Default"
	ArchivalNasExternalTargetParamsNfsSecurityTypeDefault string = "Default"

	// ArchivalNasExternalTargetParamsNfsSecurityTypeNone captures enum value "None"
	ArchivalNasExternalTargetParamsNfsSecurityTypeNone string = "None"

	// ArchivalNasExternalTargetParamsNfsSecurityTypeSystem captures enum value "System"
	ArchivalNasExternalTargetParamsNfsSecurityTypeSystem string = "System"

	// ArchivalNasExternalTargetParamsNfsSecurityTypeKRB5 captures enum value "KRB5"
	ArchivalNasExternalTargetParamsNfsSecurityTypeKRB5 string = "KRB5"

	// ArchivalNasExternalTargetParamsNfsSecurityTypeKRB5I captures enum value "KRB5I"
	ArchivalNasExternalTargetParamsNfsSecurityTypeKRB5I string = "KRB5I"

	// ArchivalNasExternalTargetParamsNfsSecurityTypeKRB5P captures enum value "KRB5P"
	ArchivalNasExternalTargetParamsNfsSecurityTypeKRB5P string = "KRB5P"
)

// prop value enum
func (m *ArchivalNasExternalTargetParams) validateNfsSecurityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archivalNasExternalTargetParamsTypeNfsSecurityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchivalNasExternalTargetParams) validateNfsSecurityType(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsSecurityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNfsSecurityTypeEnum("nfsSecurityType", "body", *m.NfsSecurityType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this archival nas external target params based on the context it is used
func (m *ArchivalNasExternalTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShareType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArchivalNasExternalTargetParams) contextValidateShareType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "shareType", "body", m.ShareType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArchivalNasExternalTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchivalNasExternalTargetParams) UnmarshalBinary(b []byte) error {
	var res ArchivalNasExternalTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
