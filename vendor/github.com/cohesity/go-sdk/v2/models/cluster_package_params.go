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

// ClusterPackageParams Cluster software package parameters.
//
// swagger:model ClusterPackageParams
type ClusterPackageParams struct {

	// Name of the package version.
	// Example: '6.6.0d_u6_release-20210714_0fad884e',
	//   '7.0.1_release-20230623_ddbb8c79' for upgrade packages,
	// '6.8.1-p1s1-2023Jun26-221b8a5c' for patch packages
	//
	VersionName *string `json:"versionName,omitempty"`

	// Release date of the package.
	// Format: date-time
	ReleaseDate *strfmt.DateTime `json:"releaseDate,omitempty"`

	// Release version of the package.
	// Examples:
	// For upgrade package: '6.6.0d_u6', '7.0.'
	// For patch package - '6.8.1-p1s1'
	//
	ReleaseVersion string `json:"releaseVersion,omitempty"`

	// Type of the package - Upgrade or Patch
	// Enum: ["Upgrade","Patch"]
	PackageType string `json:"packageType,omitempty"`

	// Sub-type of package - Security Patch or Product Patch
	// Enum: ["SecurityPatch","ProductPatch"]
	PackageSubType string `json:"packageSubType,omitempty"`

	// Node IDs where package is available
	NodeIds []int64 `json:"nodeIds"`

	// Status of package
	Status *ClusterPackageStatus `json:"status,omitempty"`

	// Indicates whether package need downtime during installation
	IsDowntimeRequired *bool `json:"isDowntimeRequired,omitempty"`

	// Size of file in bytes
	FileSizeBytes int64 `json:"fileSizeBytes,omitempty"`

	// MD5 Checksum
	Md5Checksum string `json:"md5Checksum,omitempty"`

	// SHA256 Checksum
	Sha256Checksum string `json:"sha256Checksum,omitempty"`

	// List of issues fixed in the package
	FixedIssues ClusterPackageFixedIssues `json:"fixedIssues,omitempty"`

	// Array of versionName values, representing compatible packages
	// that are available on system.
	//
	CompatiblePackages []string `json:"compatiblePackages"`

	// List of package componenets. Aplicable for one helios package
	//
	Components []*PackageComponent `json:"components"`
}

// Validate validates this cluster package params
func (m *ClusterPackageParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageSubType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPackageParams) validateReleaseDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseDate", "body", "date-time", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var clusterPackageParamsTypePackageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Upgrade","Patch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterPackageParamsTypePackageTypePropEnum = append(clusterPackageParamsTypePackageTypePropEnum, v)
	}
}

const (

	// ClusterPackageParamsPackageTypeUpgrade captures enum value "Upgrade"
	ClusterPackageParamsPackageTypeUpgrade string = "Upgrade"

	// ClusterPackageParamsPackageTypePatch captures enum value "Patch"
	ClusterPackageParamsPackageTypePatch string = "Patch"
)

// prop value enum
func (m *ClusterPackageParams) validatePackageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterPackageParamsTypePackageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterPackageParams) validatePackageType(formats strfmt.Registry) error {
	if swag.IsZero(m.PackageType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePackageTypeEnum("packageType", "body", m.PackageType); err != nil {
		return err
	}

	return nil
}

var clusterPackageParamsTypePackageSubTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SecurityPatch","ProductPatch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterPackageParamsTypePackageSubTypePropEnum = append(clusterPackageParamsTypePackageSubTypePropEnum, v)
	}
}

const (

	// ClusterPackageParamsPackageSubTypeSecurityPatch captures enum value "SecurityPatch"
	ClusterPackageParamsPackageSubTypeSecurityPatch string = "SecurityPatch"

	// ClusterPackageParamsPackageSubTypeProductPatch captures enum value "ProductPatch"
	ClusterPackageParamsPackageSubTypeProductPatch string = "ProductPatch"
)

// prop value enum
func (m *ClusterPackageParams) validatePackageSubTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterPackageParamsTypePackageSubTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterPackageParams) validatePackageSubType(formats strfmt.Registry) error {
	if swag.IsZero(m.PackageSubType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePackageSubTypeEnum("packageSubType", "body", m.PackageSubType); err != nil {
		return err
	}

	return nil
}

func (m *ClusterPackageParams) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterPackageParams) validateFixedIssues(formats strfmt.Registry) error {
	if swag.IsZero(m.FixedIssues) { // not required
		return nil
	}

	if err := m.FixedIssues.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fixedIssues")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fixedIssues")
		}
		return err
	}

	return nil
}

func (m *ClusterPackageParams) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster package params based on the context it is used
func (m *ClusterPackageParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFixedIssues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPackageParams) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterPackageParams) contextValidateFixedIssues(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FixedIssues.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fixedIssues")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fixedIssues")
		}
		return err
	}

	return nil
}

func (m *ClusterPackageParams) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {

			if swag.IsZero(m.Components[i]) { // not required
				return nil
			}

			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterPackageParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterPackageParams) UnmarshalBinary(b []byte) error {
	var res ClusterPackageParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
