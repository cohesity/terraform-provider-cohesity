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

// NetappVolumeInfo NetApp Volume Information.
//
// Specifies information about a volume in a NetApp cluster.
//
// swagger:model NetappVolumeInfo
type NetappVolumeInfo struct {

	// Specifies the containing aggregate name of this volume.
	AggregateName *string `json:"aggregateName,omitempty"`

	// Specifies the total capacity in bytes of this volume.
	CapacityBytes *int64 `json:"capacityBytes,omitempty"`

	// Array of CIFS Shares.
	//
	// Specifies the set of CIFS Shares exported for this volume.
	CifsShares []*CifsShareInfo `json:"cifsShares"`

	// Specifies the creation time of the volume specified in Unix epoch time
	// (in microseconds).
	CreationTimeUsecs *int64 `json:"creationTimeUsecs,omitempty"`

	// Array of Data Protocols.
	//
	// Specifies the set of data protocols supported by this volume.
	// 'kNfs' indicates NFS connections.
	// 'kCifs' indicates SMB (CIFS) connections.
	// 'kIscsi' indicates iSCSI connections.
	// 'kFc' indicates Fiber Channel connections.
	// 'kFcache' indicates Flex Cache connections.
	// 'kHttp' indicates HTTP connections.
	// 'kNdmp' indicates NDMP connections.
	// 'kManagement' indicates non-data connections used for management purposes.
	// 'kNvme' indicates NVMe connections.
	DataProtocols []string `json:"dataProtocols"`

	// Specifies the name of the export policy (which defines the access
	// permissions for the mount client) that has been assigned to this volume.
	ExportPolicyName *string `json:"exportPolicyName,omitempty"`

	// Specifies the Extended style information of a NetApp volume.
	// Specifies the extended style info of a NetApp Volume.
	// 'kFlexGroup' indicates FlexGroup volume. A FlexGroup volume contains
	// several constituents (which themselves are Netapp volumes) that
	// automatically and transparently share the traffic. Cohesity does
	// not need to deal with the individual consituents,
	// just the main FlexGroup volume.
	// 'kFlexVol' indicates FlexVol volume. A typical NAS share.
	// Enum: ["kFlexVol","kFlexGroup"]
	ExtendedStyle *string `json:"extendedStyle,omitempty"`

	// Specifies the junction path of this volume.
	// This path can be used to mount this volume via protocols such as NFS.
	JunctionPath *string `json:"junctionPath,omitempty"`

	// Specifies the name of the NetApp Vserver that this volume belongs to.
	Name *string `json:"name,omitempty"`

	// Specifies the security information of this volume.
	SecurityInfo *VolumeSecurityInfo `json:"securityInfo,omitempty"`

	// Specifies the state of this volume.
	// Specifies the state of a NetApp Volume.
	// 'kOnline' indicates the volume is online. Read and write access to this
	// volume is allowed.
	// 'kRestricted' indicates the volume is restricted. Some operations,
	// such as parity reconstruction, are allowed, but data access is not
	// allowed.
	// 'kOffline' indicates the volume is offline. No access to the volume is
	// allowed.
	// 'kMixed' indicates the volume is in mixed state, which means its
	// aggregates are not all in the same state.
	// Enum: ["kOnline","kRestricted","kOffline","kMixed"]
	State *string `json:"state,omitempty"`

	// Specifies the NetApp type of this volume.
	// Specifies the type of a NetApp Volume.
	// 'kReadWrite' indicates read-write volume.
	// 'kLoadSharing' indicates load-sharing volume.
	// 'kDataProtection' indicates data-protection volume.
	// 'kDataCache' indicates data-cache volume.
	// 'kTmp' indicates temporary purpose.
	// 'kUnknownType' indicates unknown type.
	// Enum: ["kReadWrite","kLoadSharing","kDataProtection","kDataCache","kTmp","kUnknownType"]
	Type *string `json:"type,omitempty"`

	// Specifies the total space (in bytes) used in this volume.
	UsedBytes *int64 `json:"usedBytes,omitempty"`
}

// Validate validates this netapp volume info
func (m *NetappVolumeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCifsShares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataProtocols(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtendedStyle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetappVolumeInfo) validateCifsShares(formats strfmt.Registry) error {
	if swag.IsZero(m.CifsShares) { // not required
		return nil
	}

	for i := 0; i < len(m.CifsShares); i++ {
		if swag.IsZero(m.CifsShares[i]) { // not required
			continue
		}

		if m.CifsShares[i] != nil {
			if err := m.CifsShares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cifsShares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cifsShares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var netappVolumeInfoDataProtocolsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNfs","kCifs","kIscsi","kFc","kFcache","kHttp","kNdmp","kManagement","kNvme"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		netappVolumeInfoDataProtocolsItemsEnum = append(netappVolumeInfoDataProtocolsItemsEnum, v)
	}
}

func (m *NetappVolumeInfo) validateDataProtocolsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, netappVolumeInfoDataProtocolsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetappVolumeInfo) validateDataProtocols(formats strfmt.Registry) error {
	if swag.IsZero(m.DataProtocols) { // not required
		return nil
	}

	for i := 0; i < len(m.DataProtocols); i++ {

		// value enum
		if err := m.validateDataProtocolsItemsEnum("dataProtocols"+"."+strconv.Itoa(i), "body", m.DataProtocols[i]); err != nil {
			return err
		}

	}

	return nil
}

var netappVolumeInfoTypeExtendedStylePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kFlexVol","kFlexGroup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		netappVolumeInfoTypeExtendedStylePropEnum = append(netappVolumeInfoTypeExtendedStylePropEnum, v)
	}
}

const (

	// NetappVolumeInfoExtendedStyleKFlexVol captures enum value "kFlexVol"
	NetappVolumeInfoExtendedStyleKFlexVol string = "kFlexVol"

	// NetappVolumeInfoExtendedStyleKFlexGroup captures enum value "kFlexGroup"
	NetappVolumeInfoExtendedStyleKFlexGroup string = "kFlexGroup"
)

// prop value enum
func (m *NetappVolumeInfo) validateExtendedStyleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, netappVolumeInfoTypeExtendedStylePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetappVolumeInfo) validateExtendedStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtendedStyle) { // not required
		return nil
	}

	// value enum
	if err := m.validateExtendedStyleEnum("extendedStyle", "body", *m.ExtendedStyle); err != nil {
		return err
	}

	return nil
}

func (m *NetappVolumeInfo) validateSecurityInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityInfo) { // not required
		return nil
	}

	if m.SecurityInfo != nil {
		if err := m.SecurityInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityInfo")
			}
			return err
		}
	}

	return nil
}

var netappVolumeInfoTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kOnline","kRestricted","kOffline","kMixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		netappVolumeInfoTypeStatePropEnum = append(netappVolumeInfoTypeStatePropEnum, v)
	}
}

const (

	// NetappVolumeInfoStateKOnline captures enum value "kOnline"
	NetappVolumeInfoStateKOnline string = "kOnline"

	// NetappVolumeInfoStateKRestricted captures enum value "kRestricted"
	NetappVolumeInfoStateKRestricted string = "kRestricted"

	// NetappVolumeInfoStateKOffline captures enum value "kOffline"
	NetappVolumeInfoStateKOffline string = "kOffline"

	// NetappVolumeInfoStateKMixed captures enum value "kMixed"
	NetappVolumeInfoStateKMixed string = "kMixed"
)

// prop value enum
func (m *NetappVolumeInfo) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, netappVolumeInfoTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetappVolumeInfo) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var netappVolumeInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kReadWrite","kLoadSharing","kDataProtection","kDataCache","kTmp","kUnknownType"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		netappVolumeInfoTypeTypePropEnum = append(netappVolumeInfoTypeTypePropEnum, v)
	}
}

const (

	// NetappVolumeInfoTypeKReadWrite captures enum value "kReadWrite"
	NetappVolumeInfoTypeKReadWrite string = "kReadWrite"

	// NetappVolumeInfoTypeKLoadSharing captures enum value "kLoadSharing"
	NetappVolumeInfoTypeKLoadSharing string = "kLoadSharing"

	// NetappVolumeInfoTypeKDataProtection captures enum value "kDataProtection"
	NetappVolumeInfoTypeKDataProtection string = "kDataProtection"

	// NetappVolumeInfoTypeKDataCache captures enum value "kDataCache"
	NetappVolumeInfoTypeKDataCache string = "kDataCache"

	// NetappVolumeInfoTypeKTmp captures enum value "kTmp"
	NetappVolumeInfoTypeKTmp string = "kTmp"

	// NetappVolumeInfoTypeKUnknownType captures enum value "kUnknownType"
	NetappVolumeInfoTypeKUnknownType string = "kUnknownType"
)

// prop value enum
func (m *NetappVolumeInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, netappVolumeInfoTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetappVolumeInfo) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this netapp volume info based on the context it is used
func (m *NetappVolumeInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCifsShares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetappVolumeInfo) contextValidateCifsShares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CifsShares); i++ {

		if m.CifsShares[i] != nil {

			if swag.IsZero(m.CifsShares[i]) { // not required
				return nil
			}

			if err := m.CifsShares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cifsShares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cifsShares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetappVolumeInfo) contextValidateSecurityInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityInfo != nil {

		if swag.IsZero(m.SecurityInfo) { // not required
			return nil
		}

		if err := m.SecurityInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetappVolumeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetappVolumeInfo) UnmarshalBinary(b []byte) error {
	var res NetappVolumeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
