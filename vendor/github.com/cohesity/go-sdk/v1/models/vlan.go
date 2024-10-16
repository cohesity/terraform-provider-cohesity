// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Vlan VLAN.
//
// Specifies the settings of a VLAN.
// Its used by both Request and Response structures.
//
// swagger:model Vlan
type Vlan struct {

	// Specifies whether to add the VLAN IPs to the cluster partition
	// that already has one or more IPs from this VLAN.
	AddToClusterPartition *bool `json:"addToClusterPartition,omitempty"`

	// Specifies if this VLAN can be used by all tenants without explicit
	// assignment to them. This option can only be set true for VLANs that are
	// not assigned to any tenant.
	AllTenantAccess *bool `json:"allTenantAccess,omitempty"`

	// Set to true when ips are in use by Athena Apps.
	// Note: If it is true then vlan interface can't be deleted.
	AppIPVecInUse *bool `json:"appIpVecInUse,omitempty"`

	// Array of Athena Apps IPs.
	//
	// Specifies a list of Athena IPs in the VLAN.
	Appsips []string `json:"appsips"`

	// Specifies a description of the VLAN.
	Description *string `json:"description,omitempty"`

	// Specifies list of dns delegation zones.
	DNSDelegationZones []*DNSDelegationZone `json:"dnsDelegationZones"`

	// EcmpEnabled.
	// Specifies if ECMP is enabled in the VLAN.
	EcmpEnabled *bool `json:"ecmpEnabled,omitempty"`

	// Specifies the Gateway of the VLAN. It can carry V4 or V6 in case of
	// requests, and carrises V4 in case of response.
	Gateway *string `json:"gateway,omitempty"`

	// Specifies the Gateway of the VLAN.
	GatewayV6 *string `json:"gatewayV6,omitempty"`

	// Specifies the hostname of the VLAN.
	Hostname *string `json:"hostname,omitempty"`

	// Specifies the id of the VLAN.
	ID *int32 `json:"id,omitempty"`

	// Specifies the interface group name of the VLAN. It is in the format of
	// <base_interface_group_name>.<vlan_id>.
	IfaceGroupName *string `json:"ifaceGroupName,omitempty"`

	// Specifies the id of the Loopback Interface group. Used only in get,
	// for display.
	InterfaceGroupID *int32 `json:"interfaceGroupId,omitempty"`

	// Specifies the interface name of the VLAN.
	InterfaceName *string `json:"interfaceName,omitempty"`

	// Specifies IP family. Based on this, subnet/gateway field contains
	// V4 or V6 values. Used in Request.
	IPFamily *int32 `json:"ipFamily,omitempty"`

	// IpPoolMap.
	//
	// Pool IPs to program VIP followers.
	IPPoolMap map[string][]string `json:"ipPoolMap,omitempty"`

	// IP Range for vip addition
	IPRange *IPRange `json:"ipRange,omitempty"`

	// Array of range of ips.
	// If specified in PUT request, Ips field will be ignored.
	// Specifies ips in compressed way using list of [start, end] vips.
	IPRanges []*IPRange `json:"ipRanges"`

	// Array of IPs.
	//
	// Specifies a list of IPs in the VLAN.
	Ips []string `json:"ips"`

	// mtu
	Mtu *int32 `json:"mtu,omitempty"`

	// Subnet.
	//
	// Specifies the subnet of the VLAN.
	// The netmask can be specified by setting netmaskBits or netmaskIp4.
	// The netmask can only be set using netmaskIp4 if the IP address is
	// an IPv4 address. It can carry V4 or V6 in case of requests, and carries
	// V4 in case of response.
	Subnet struct {
		Subnet
	} `json:"subnet,omitempty"`

	// Subnet.
	//
	// Specifies the subnet of the VLAN.
	// The netmask can be specified by setting netmaskBits or netmaskIp4.
	// The netmask can only be set using netmaskIp4 if the IP address is
	// an IPv4 address.
	SubnetV6 struct {
		Subnet
	} `json:"subnetV6,omitempty"`

	// Optional tenant id that this vlan belongs to.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies the VLAN name of the vlanId.
	VlanName *string `json:"vlanName,omitempty"`
}

// Validate validates this vlan
func (m *Vlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSDelegationZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetV6(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlan) validateDNSDelegationZones(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSDelegationZones) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSDelegationZones); i++ {
		if swag.IsZero(m.DNSDelegationZones[i]) { // not required
			continue
		}

		if m.DNSDelegationZones[i] != nil {
			if err := m.DNSDelegationZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsDelegationZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsDelegationZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Vlan) validateIPRange(formats strfmt.Registry) error {
	if swag.IsZero(m.IPRange) { // not required
		return nil
	}

	if m.IPRange != nil {
		if err := m.IPRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipRange")
			}
			return err
		}
	}

	return nil
}

func (m *Vlan) validateIPRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.IPRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.IPRanges); i++ {
		if swag.IsZero(m.IPRanges[i]) { // not required
			continue
		}

		if m.IPRanges[i] != nil {
			if err := m.IPRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Vlan) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	return nil
}

func (m *Vlan) validateSubnetV6(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetV6) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this vlan based on the context it is used
func (m *Vlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSDelegationZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlan) contextValidateDNSDelegationZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSDelegationZones); i++ {

		if m.DNSDelegationZones[i] != nil {

			if swag.IsZero(m.DNSDelegationZones[i]) { // not required
				return nil
			}

			if err := m.DNSDelegationZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsDelegationZones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsDelegationZones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Vlan) contextValidateIPRange(ctx context.Context, formats strfmt.Registry) error {

	if m.IPRange != nil {

		if swag.IsZero(m.IPRange) { // not required
			return nil
		}

		if err := m.IPRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipRange")
			}
			return err
		}
	}

	return nil
}

func (m *Vlan) contextValidateIPRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPRanges); i++ {

		if m.IPRanges[i] != nil {

			if swag.IsZero(m.IPRanges[i]) { // not required
				return nil
			}

			if err := m.IPRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Vlan) contextValidateSubnet(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Vlan) contextValidateSubnetV6(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *Vlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vlan) UnmarshalBinary(b []byte) error {
	var res Vlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
