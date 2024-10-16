// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataTransferToVaultPerProtectionJob Data Transfer to Vault Per Protection Job.
//
// Specifies statistics about the transfer of data from this Cohesity
// Cluster to this Vault for a Protection Job.
//
// swagger:model DataTransferToVaultPerProtectionJob
type DataTransferToVaultPerProtectionJob struct {

	// Specifies the total number of logical bytes that are transferred
	// from this Cohesity Cluster to this Vault for this Protection Job.
	// The logical size is when the data is fully hydrated or expanded.
	NumLogicalBytesTransferred *int64 `json:"numLogicalBytesTransferred,omitempty"`

	// Specifies the total number of physical bytes that are transferred
	// from this Cohesity Cluster to this Vault for this Protection Job.
	NumPhysicalBytesTransferred *int64 `json:"numPhysicalBytesTransferred,omitempty"`

	// Specifies the name of the Protection Job.
	ProtectionJobName *string `json:"protectionJobName,omitempty"`

	// Specifies the total number of storage bytes consumed that are transferred
	// from this Cohesity Cluster to this vault for this Protection Job.
	StorageConsumed *int64 `json:"storageConsumed,omitempty"`
}

// Validate validates this data transfer to vault per protection job
func (m *DataTransferToVaultPerProtectionJob) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data transfer to vault per protection job based on context it is used
func (m *DataTransferToVaultPerProtectionJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTransferToVaultPerProtectionJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTransferToVaultPerProtectionJob) UnmarshalBinary(b []byte) error {
	var res DataTransferToVaultPerProtectionJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
