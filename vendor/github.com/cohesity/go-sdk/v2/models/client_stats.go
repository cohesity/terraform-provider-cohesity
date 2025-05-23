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

// ClientStats Specifies the Client stats.
//
// swagger:model ClientStats
type ClientStats struct {

	// Specifies the stats metric.
	// Enum: ["kNumBytesRead","kNumBytesWritten","kReadIos","kWriteIos","kReadLatencyUsecs","kWriteLatencyUsecs","kNumReadErrors","kNumWriteErrors","kNFSv3MntLatUsecs","kNFSv3MntOps","kNFSv3MntErr","kNFSv3UmntLatUsecs","kNFSv3UmntOps","kNFSv3UmntErr","kNFSv3CreateLatUsecs","kNFSv3CreateOps","kNFSv3CreateErr","kNFSv3RemoveLatUsecs","kNFSv3RemoveOps","kNFSv3RemoveErr","kNFSv3MkdirLatUsecs","kNFSv3MkdirOps","kNFSv3MkdirErr","kNFSv3RmdirLatUsecs","kNFSv3RmdirOps","kNFSv3RmdirErr","kNFSv3LookupLatUsecs","kNFSv3LookupOps","kNFSv3LookupErr","kNFSv3ReaddirLatUsecs","kNFSv3ReaddirOps","kNFSv3ReaddirErr","kNFSv3ReaddirplusLatUsecs","kNFSv3ReaddirplusOps","kNFSv3ReaddirplusErr","kNFSv3SymlinkLatUsecs","kNFSv3SymlinkOps","kNFSv3SymlinkErr","kNFSv3RenameLatUsecs","kNFSv3RenameOps","kNFSv3RenameErr","kSMBSessionSetupLatUsecs","kSMBSessionSetupOps","kSMBSessionSetupErr","kSMBLogoffLatUsecs","kSMBLogoffOps","kSMBLogoffErr","kSMBTreeConnectLatUsecs","kSMBTreeConnectOps","kSMBTreeConnectErr","kSMBTreeDisconnectLatUsecs","kSMBTreeDisconnectOps","kSMBTreeDisconnectErr","kSMBOpenLatUsecs","kSMBOpenOps","kSMBOpenErr","kSMBCloseLatUsecs","kSMBCloseOps","kSMBCloseErr","kSMBCreateFileLatUsecs","kSMBCreateFileOps","kSMBCreateFileErr","kSMBDeleteFileLatUsecs","kSMBDeleteFileOps","kSMBDeleteFileErr","kSMBMkdirLatUsecs","kSMBMkdirOps","kSMBMkdirErr","kSMBRmdirLatUsecs","kSMBRmdirOps","kSMBRmdirErr","kSMBReaddirLatUsecs","kSMBReaddirOps","kSMBReaddirErr","kSMBGetInfoLatUsecs","kSMBGetInfoOps","kSMBGetInfoErr","kSMBSetInfoLatUsecs","kSMBSetInfoOps","kSMBSetInfoErr","kSMBRenameLatUsecs","kSMBRenameOps","kSMBRenameErr"]
	Metric *string `json:"metric,omitempty"`

	// Specifies the stats value.
	Value *int64 `json:"value,omitempty"`

	// Specifies the stats value in last hours.
	ValueInLastHours []*ClientStatsInLastHours `json:"valueInLastHours"`
}

// Validate validates this client stats
func (m *ClientStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueInLastHours(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clientStatsTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNumBytesRead","kNumBytesWritten","kReadIos","kWriteIos","kReadLatencyUsecs","kWriteLatencyUsecs","kNumReadErrors","kNumWriteErrors","kNFSv3MntLatUsecs","kNFSv3MntOps","kNFSv3MntErr","kNFSv3UmntLatUsecs","kNFSv3UmntOps","kNFSv3UmntErr","kNFSv3CreateLatUsecs","kNFSv3CreateOps","kNFSv3CreateErr","kNFSv3RemoveLatUsecs","kNFSv3RemoveOps","kNFSv3RemoveErr","kNFSv3MkdirLatUsecs","kNFSv3MkdirOps","kNFSv3MkdirErr","kNFSv3RmdirLatUsecs","kNFSv3RmdirOps","kNFSv3RmdirErr","kNFSv3LookupLatUsecs","kNFSv3LookupOps","kNFSv3LookupErr","kNFSv3ReaddirLatUsecs","kNFSv3ReaddirOps","kNFSv3ReaddirErr","kNFSv3ReaddirplusLatUsecs","kNFSv3ReaddirplusOps","kNFSv3ReaddirplusErr","kNFSv3SymlinkLatUsecs","kNFSv3SymlinkOps","kNFSv3SymlinkErr","kNFSv3RenameLatUsecs","kNFSv3RenameOps","kNFSv3RenameErr","kSMBSessionSetupLatUsecs","kSMBSessionSetupOps","kSMBSessionSetupErr","kSMBLogoffLatUsecs","kSMBLogoffOps","kSMBLogoffErr","kSMBTreeConnectLatUsecs","kSMBTreeConnectOps","kSMBTreeConnectErr","kSMBTreeDisconnectLatUsecs","kSMBTreeDisconnectOps","kSMBTreeDisconnectErr","kSMBOpenLatUsecs","kSMBOpenOps","kSMBOpenErr","kSMBCloseLatUsecs","kSMBCloseOps","kSMBCloseErr","kSMBCreateFileLatUsecs","kSMBCreateFileOps","kSMBCreateFileErr","kSMBDeleteFileLatUsecs","kSMBDeleteFileOps","kSMBDeleteFileErr","kSMBMkdirLatUsecs","kSMBMkdirOps","kSMBMkdirErr","kSMBRmdirLatUsecs","kSMBRmdirOps","kSMBRmdirErr","kSMBReaddirLatUsecs","kSMBReaddirOps","kSMBReaddirErr","kSMBGetInfoLatUsecs","kSMBGetInfoOps","kSMBGetInfoErr","kSMBSetInfoLatUsecs","kSMBSetInfoOps","kSMBSetInfoErr","kSMBRenameLatUsecs","kSMBRenameOps","kSMBRenameErr"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientStatsTypeMetricPropEnum = append(clientStatsTypeMetricPropEnum, v)
	}
}

const (

	// ClientStatsMetricKNumBytesRead captures enum value "kNumBytesRead"
	ClientStatsMetricKNumBytesRead string = "kNumBytesRead"

	// ClientStatsMetricKNumBytesWritten captures enum value "kNumBytesWritten"
	ClientStatsMetricKNumBytesWritten string = "kNumBytesWritten"

	// ClientStatsMetricKReadIos captures enum value "kReadIos"
	ClientStatsMetricKReadIos string = "kReadIos"

	// ClientStatsMetricKWriteIos captures enum value "kWriteIos"
	ClientStatsMetricKWriteIos string = "kWriteIos"

	// ClientStatsMetricKReadLatencyUsecs captures enum value "kReadLatencyUsecs"
	ClientStatsMetricKReadLatencyUsecs string = "kReadLatencyUsecs"

	// ClientStatsMetricKWriteLatencyUsecs captures enum value "kWriteLatencyUsecs"
	ClientStatsMetricKWriteLatencyUsecs string = "kWriteLatencyUsecs"

	// ClientStatsMetricKNumReadErrors captures enum value "kNumReadErrors"
	ClientStatsMetricKNumReadErrors string = "kNumReadErrors"

	// ClientStatsMetricKNumWriteErrors captures enum value "kNumWriteErrors"
	ClientStatsMetricKNumWriteErrors string = "kNumWriteErrors"

	// ClientStatsMetricKNFSv3MntLatUsecs captures enum value "kNFSv3MntLatUsecs"
	ClientStatsMetricKNFSv3MntLatUsecs string = "kNFSv3MntLatUsecs"

	// ClientStatsMetricKNFSv3MntOps captures enum value "kNFSv3MntOps"
	ClientStatsMetricKNFSv3MntOps string = "kNFSv3MntOps"

	// ClientStatsMetricKNFSv3MntErr captures enum value "kNFSv3MntErr"
	ClientStatsMetricKNFSv3MntErr string = "kNFSv3MntErr"

	// ClientStatsMetricKNFSv3UmntLatUsecs captures enum value "kNFSv3UmntLatUsecs"
	ClientStatsMetricKNFSv3UmntLatUsecs string = "kNFSv3UmntLatUsecs"

	// ClientStatsMetricKNFSv3UmntOps captures enum value "kNFSv3UmntOps"
	ClientStatsMetricKNFSv3UmntOps string = "kNFSv3UmntOps"

	// ClientStatsMetricKNFSv3UmntErr captures enum value "kNFSv3UmntErr"
	ClientStatsMetricKNFSv3UmntErr string = "kNFSv3UmntErr"

	// ClientStatsMetricKNFSv3CreateLatUsecs captures enum value "kNFSv3CreateLatUsecs"
	ClientStatsMetricKNFSv3CreateLatUsecs string = "kNFSv3CreateLatUsecs"

	// ClientStatsMetricKNFSv3CreateOps captures enum value "kNFSv3CreateOps"
	ClientStatsMetricKNFSv3CreateOps string = "kNFSv3CreateOps"

	// ClientStatsMetricKNFSv3CreateErr captures enum value "kNFSv3CreateErr"
	ClientStatsMetricKNFSv3CreateErr string = "kNFSv3CreateErr"

	// ClientStatsMetricKNFSv3RemoveLatUsecs captures enum value "kNFSv3RemoveLatUsecs"
	ClientStatsMetricKNFSv3RemoveLatUsecs string = "kNFSv3RemoveLatUsecs"

	// ClientStatsMetricKNFSv3RemoveOps captures enum value "kNFSv3RemoveOps"
	ClientStatsMetricKNFSv3RemoveOps string = "kNFSv3RemoveOps"

	// ClientStatsMetricKNFSv3RemoveErr captures enum value "kNFSv3RemoveErr"
	ClientStatsMetricKNFSv3RemoveErr string = "kNFSv3RemoveErr"

	// ClientStatsMetricKNFSv3MkdirLatUsecs captures enum value "kNFSv3MkdirLatUsecs"
	ClientStatsMetricKNFSv3MkdirLatUsecs string = "kNFSv3MkdirLatUsecs"

	// ClientStatsMetricKNFSv3MkdirOps captures enum value "kNFSv3MkdirOps"
	ClientStatsMetricKNFSv3MkdirOps string = "kNFSv3MkdirOps"

	// ClientStatsMetricKNFSv3MkdirErr captures enum value "kNFSv3MkdirErr"
	ClientStatsMetricKNFSv3MkdirErr string = "kNFSv3MkdirErr"

	// ClientStatsMetricKNFSv3RmdirLatUsecs captures enum value "kNFSv3RmdirLatUsecs"
	ClientStatsMetricKNFSv3RmdirLatUsecs string = "kNFSv3RmdirLatUsecs"

	// ClientStatsMetricKNFSv3RmdirOps captures enum value "kNFSv3RmdirOps"
	ClientStatsMetricKNFSv3RmdirOps string = "kNFSv3RmdirOps"

	// ClientStatsMetricKNFSv3RmdirErr captures enum value "kNFSv3RmdirErr"
	ClientStatsMetricKNFSv3RmdirErr string = "kNFSv3RmdirErr"

	// ClientStatsMetricKNFSv3LookupLatUsecs captures enum value "kNFSv3LookupLatUsecs"
	ClientStatsMetricKNFSv3LookupLatUsecs string = "kNFSv3LookupLatUsecs"

	// ClientStatsMetricKNFSv3LookupOps captures enum value "kNFSv3LookupOps"
	ClientStatsMetricKNFSv3LookupOps string = "kNFSv3LookupOps"

	// ClientStatsMetricKNFSv3LookupErr captures enum value "kNFSv3LookupErr"
	ClientStatsMetricKNFSv3LookupErr string = "kNFSv3LookupErr"

	// ClientStatsMetricKNFSv3ReaddirLatUsecs captures enum value "kNFSv3ReaddirLatUsecs"
	ClientStatsMetricKNFSv3ReaddirLatUsecs string = "kNFSv3ReaddirLatUsecs"

	// ClientStatsMetricKNFSv3ReaddirOps captures enum value "kNFSv3ReaddirOps"
	ClientStatsMetricKNFSv3ReaddirOps string = "kNFSv3ReaddirOps"

	// ClientStatsMetricKNFSv3ReaddirErr captures enum value "kNFSv3ReaddirErr"
	ClientStatsMetricKNFSv3ReaddirErr string = "kNFSv3ReaddirErr"

	// ClientStatsMetricKNFSv3ReaddirplusLatUsecs captures enum value "kNFSv3ReaddirplusLatUsecs"
	ClientStatsMetricKNFSv3ReaddirplusLatUsecs string = "kNFSv3ReaddirplusLatUsecs"

	// ClientStatsMetricKNFSv3ReaddirplusOps captures enum value "kNFSv3ReaddirplusOps"
	ClientStatsMetricKNFSv3ReaddirplusOps string = "kNFSv3ReaddirplusOps"

	// ClientStatsMetricKNFSv3ReaddirplusErr captures enum value "kNFSv3ReaddirplusErr"
	ClientStatsMetricKNFSv3ReaddirplusErr string = "kNFSv3ReaddirplusErr"

	// ClientStatsMetricKNFSv3SymlinkLatUsecs captures enum value "kNFSv3SymlinkLatUsecs"
	ClientStatsMetricKNFSv3SymlinkLatUsecs string = "kNFSv3SymlinkLatUsecs"

	// ClientStatsMetricKNFSv3SymlinkOps captures enum value "kNFSv3SymlinkOps"
	ClientStatsMetricKNFSv3SymlinkOps string = "kNFSv3SymlinkOps"

	// ClientStatsMetricKNFSv3SymlinkErr captures enum value "kNFSv3SymlinkErr"
	ClientStatsMetricKNFSv3SymlinkErr string = "kNFSv3SymlinkErr"

	// ClientStatsMetricKNFSv3RenameLatUsecs captures enum value "kNFSv3RenameLatUsecs"
	ClientStatsMetricKNFSv3RenameLatUsecs string = "kNFSv3RenameLatUsecs"

	// ClientStatsMetricKNFSv3RenameOps captures enum value "kNFSv3RenameOps"
	ClientStatsMetricKNFSv3RenameOps string = "kNFSv3RenameOps"

	// ClientStatsMetricKNFSv3RenameErr captures enum value "kNFSv3RenameErr"
	ClientStatsMetricKNFSv3RenameErr string = "kNFSv3RenameErr"

	// ClientStatsMetricKSMBSessionSetupLatUsecs captures enum value "kSMBSessionSetupLatUsecs"
	ClientStatsMetricKSMBSessionSetupLatUsecs string = "kSMBSessionSetupLatUsecs"

	// ClientStatsMetricKSMBSessionSetupOps captures enum value "kSMBSessionSetupOps"
	ClientStatsMetricKSMBSessionSetupOps string = "kSMBSessionSetupOps"

	// ClientStatsMetricKSMBSessionSetupErr captures enum value "kSMBSessionSetupErr"
	ClientStatsMetricKSMBSessionSetupErr string = "kSMBSessionSetupErr"

	// ClientStatsMetricKSMBLogoffLatUsecs captures enum value "kSMBLogoffLatUsecs"
	ClientStatsMetricKSMBLogoffLatUsecs string = "kSMBLogoffLatUsecs"

	// ClientStatsMetricKSMBLogoffOps captures enum value "kSMBLogoffOps"
	ClientStatsMetricKSMBLogoffOps string = "kSMBLogoffOps"

	// ClientStatsMetricKSMBLogoffErr captures enum value "kSMBLogoffErr"
	ClientStatsMetricKSMBLogoffErr string = "kSMBLogoffErr"

	// ClientStatsMetricKSMBTreeConnectLatUsecs captures enum value "kSMBTreeConnectLatUsecs"
	ClientStatsMetricKSMBTreeConnectLatUsecs string = "kSMBTreeConnectLatUsecs"

	// ClientStatsMetricKSMBTreeConnectOps captures enum value "kSMBTreeConnectOps"
	ClientStatsMetricKSMBTreeConnectOps string = "kSMBTreeConnectOps"

	// ClientStatsMetricKSMBTreeConnectErr captures enum value "kSMBTreeConnectErr"
	ClientStatsMetricKSMBTreeConnectErr string = "kSMBTreeConnectErr"

	// ClientStatsMetricKSMBTreeDisconnectLatUsecs captures enum value "kSMBTreeDisconnectLatUsecs"
	ClientStatsMetricKSMBTreeDisconnectLatUsecs string = "kSMBTreeDisconnectLatUsecs"

	// ClientStatsMetricKSMBTreeDisconnectOps captures enum value "kSMBTreeDisconnectOps"
	ClientStatsMetricKSMBTreeDisconnectOps string = "kSMBTreeDisconnectOps"

	// ClientStatsMetricKSMBTreeDisconnectErr captures enum value "kSMBTreeDisconnectErr"
	ClientStatsMetricKSMBTreeDisconnectErr string = "kSMBTreeDisconnectErr"

	// ClientStatsMetricKSMBOpenLatUsecs captures enum value "kSMBOpenLatUsecs"
	ClientStatsMetricKSMBOpenLatUsecs string = "kSMBOpenLatUsecs"

	// ClientStatsMetricKSMBOpenOps captures enum value "kSMBOpenOps"
	ClientStatsMetricKSMBOpenOps string = "kSMBOpenOps"

	// ClientStatsMetricKSMBOpenErr captures enum value "kSMBOpenErr"
	ClientStatsMetricKSMBOpenErr string = "kSMBOpenErr"

	// ClientStatsMetricKSMBCloseLatUsecs captures enum value "kSMBCloseLatUsecs"
	ClientStatsMetricKSMBCloseLatUsecs string = "kSMBCloseLatUsecs"

	// ClientStatsMetricKSMBCloseOps captures enum value "kSMBCloseOps"
	ClientStatsMetricKSMBCloseOps string = "kSMBCloseOps"

	// ClientStatsMetricKSMBCloseErr captures enum value "kSMBCloseErr"
	ClientStatsMetricKSMBCloseErr string = "kSMBCloseErr"

	// ClientStatsMetricKSMBCreateFileLatUsecs captures enum value "kSMBCreateFileLatUsecs"
	ClientStatsMetricKSMBCreateFileLatUsecs string = "kSMBCreateFileLatUsecs"

	// ClientStatsMetricKSMBCreateFileOps captures enum value "kSMBCreateFileOps"
	ClientStatsMetricKSMBCreateFileOps string = "kSMBCreateFileOps"

	// ClientStatsMetricKSMBCreateFileErr captures enum value "kSMBCreateFileErr"
	ClientStatsMetricKSMBCreateFileErr string = "kSMBCreateFileErr"

	// ClientStatsMetricKSMBDeleteFileLatUsecs captures enum value "kSMBDeleteFileLatUsecs"
	ClientStatsMetricKSMBDeleteFileLatUsecs string = "kSMBDeleteFileLatUsecs"

	// ClientStatsMetricKSMBDeleteFileOps captures enum value "kSMBDeleteFileOps"
	ClientStatsMetricKSMBDeleteFileOps string = "kSMBDeleteFileOps"

	// ClientStatsMetricKSMBDeleteFileErr captures enum value "kSMBDeleteFileErr"
	ClientStatsMetricKSMBDeleteFileErr string = "kSMBDeleteFileErr"

	// ClientStatsMetricKSMBMkdirLatUsecs captures enum value "kSMBMkdirLatUsecs"
	ClientStatsMetricKSMBMkdirLatUsecs string = "kSMBMkdirLatUsecs"

	// ClientStatsMetricKSMBMkdirOps captures enum value "kSMBMkdirOps"
	ClientStatsMetricKSMBMkdirOps string = "kSMBMkdirOps"

	// ClientStatsMetricKSMBMkdirErr captures enum value "kSMBMkdirErr"
	ClientStatsMetricKSMBMkdirErr string = "kSMBMkdirErr"

	// ClientStatsMetricKSMBRmdirLatUsecs captures enum value "kSMBRmdirLatUsecs"
	ClientStatsMetricKSMBRmdirLatUsecs string = "kSMBRmdirLatUsecs"

	// ClientStatsMetricKSMBRmdirOps captures enum value "kSMBRmdirOps"
	ClientStatsMetricKSMBRmdirOps string = "kSMBRmdirOps"

	// ClientStatsMetricKSMBRmdirErr captures enum value "kSMBRmdirErr"
	ClientStatsMetricKSMBRmdirErr string = "kSMBRmdirErr"

	// ClientStatsMetricKSMBReaddirLatUsecs captures enum value "kSMBReaddirLatUsecs"
	ClientStatsMetricKSMBReaddirLatUsecs string = "kSMBReaddirLatUsecs"

	// ClientStatsMetricKSMBReaddirOps captures enum value "kSMBReaddirOps"
	ClientStatsMetricKSMBReaddirOps string = "kSMBReaddirOps"

	// ClientStatsMetricKSMBReaddirErr captures enum value "kSMBReaddirErr"
	ClientStatsMetricKSMBReaddirErr string = "kSMBReaddirErr"

	// ClientStatsMetricKSMBGetInfoLatUsecs captures enum value "kSMBGetInfoLatUsecs"
	ClientStatsMetricKSMBGetInfoLatUsecs string = "kSMBGetInfoLatUsecs"

	// ClientStatsMetricKSMBGetInfoOps captures enum value "kSMBGetInfoOps"
	ClientStatsMetricKSMBGetInfoOps string = "kSMBGetInfoOps"

	// ClientStatsMetricKSMBGetInfoErr captures enum value "kSMBGetInfoErr"
	ClientStatsMetricKSMBGetInfoErr string = "kSMBGetInfoErr"

	// ClientStatsMetricKSMBSetInfoLatUsecs captures enum value "kSMBSetInfoLatUsecs"
	ClientStatsMetricKSMBSetInfoLatUsecs string = "kSMBSetInfoLatUsecs"

	// ClientStatsMetricKSMBSetInfoOps captures enum value "kSMBSetInfoOps"
	ClientStatsMetricKSMBSetInfoOps string = "kSMBSetInfoOps"

	// ClientStatsMetricKSMBSetInfoErr captures enum value "kSMBSetInfoErr"
	ClientStatsMetricKSMBSetInfoErr string = "kSMBSetInfoErr"

	// ClientStatsMetricKSMBRenameLatUsecs captures enum value "kSMBRenameLatUsecs"
	ClientStatsMetricKSMBRenameLatUsecs string = "kSMBRenameLatUsecs"

	// ClientStatsMetricKSMBRenameOps captures enum value "kSMBRenameOps"
	ClientStatsMetricKSMBRenameOps string = "kSMBRenameOps"

	// ClientStatsMetricKSMBRenameErr captures enum value "kSMBRenameErr"
	ClientStatsMetricKSMBRenameErr string = "kSMBRenameErr"
)

// prop value enum
func (m *ClientStats) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientStatsTypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientStats) validateMetric(formats strfmt.Registry) error {
	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", *m.Metric); err != nil {
		return err
	}

	return nil
}

func (m *ClientStats) validateValueInLastHours(formats strfmt.Registry) error {
	if swag.IsZero(m.ValueInLastHours) { // not required
		return nil
	}

	for i := 0; i < len(m.ValueInLastHours); i++ {
		if swag.IsZero(m.ValueInLastHours[i]) { // not required
			continue
		}

		if m.ValueInLastHours[i] != nil {
			if err := m.ValueInLastHours[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("valueInLastHours" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("valueInLastHours" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this client stats based on the context it is used
func (m *ClientStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValueInLastHours(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientStats) contextValidateValueInLastHours(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ValueInLastHours); i++ {

		if m.ValueInLastHours[i] != nil {

			if swag.IsZero(m.ValueInLastHours[i]) { // not required
				return nil
			}

			if err := m.ValueInLastHours[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("valueInLastHours" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("valueInLastHours" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientStats) UnmarshalBinary(b []byte) error {
	var res ClientStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
