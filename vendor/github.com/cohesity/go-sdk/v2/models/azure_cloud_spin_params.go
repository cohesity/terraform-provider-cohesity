// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureCloudSpinParams Azure Parameters.
//
// Specifies various resources when converting and deploying a VM to Azure.
//
// swagger:model AzureCloudSpinParams
type AzureCloudSpinParams struct {

	// Specifies id of the resource group for the selected virtual network.
	NetworkResourceGroupID *int64 `json:"networkResourceGroupId,omitempty"`

	// Specifies id of the Azure resource group. Its value is globally unique within Azure.
	ResourceGroupID *int64 `json:"resourceGroupId,omitempty"`

	// Specifies id of the storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
	StorageAccountID *int64 `json:"storageAccountId,omitempty"`

	// Specifies id of the storage container within the above storage account.
	StorageContainerID *int64 `json:"storageContainerId,omitempty"`

	// Specifies id of the resource group for the selected storage account.
	StorageResourceGroupID *int64 `json:"storageResourceGroupId,omitempty"`

	// Specifies the availability set.
	AvailabilitySetID *int64 `json:"availabilitySetId,omitempty"`

	// Specifies id of the temporary Azure resource group.
	TempVMResourceGroupID *int64 `json:"tempVmResourceGroupId,omitempty"`

	// Specifies id of the temporary VM storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
	TempVMStorageAccountID *int64 `json:"tempVmStorageAccountId,omitempty"`

	// Specifies id of the temporary VM storage container within the above storage account.
	TempVMStorageContainerID *int64 `json:"tempVmStorageContainerId,omitempty"`

	// Specifies Id of the temporary VM subnet within the above virtual network.
	TempVMSubnetID *int64 `json:"tempVmSubnetId,omitempty"`

	// Specifies Id of the temporary VM Virtual Network.
	TempVMVirtualNetworkID *int64 `json:"tempVmVirtualNetworkId,omitempty"`
}

// Validate validates this azure cloud spin params
func (m *AzureCloudSpinParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure cloud spin params based on context it is used
func (m *AzureCloudSpinParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureCloudSpinParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCloudSpinParams) UnmarshalBinary(b []byte) error {
	var res AzureCloudSpinParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
