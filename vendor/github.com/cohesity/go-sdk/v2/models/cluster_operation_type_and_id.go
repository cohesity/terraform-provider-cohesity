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

// ClusterOperationTypeAndID Operation type and id of the cluster operation.
//
// swagger:model ClusterOperationTypeAndId
type ClusterOperationTypeAndID struct {

	// Operation type of cluster operation created for the request.
	//
	// Required: true
	// Enum: ["Destroy","Create","NodeAddition","NodeRemoval","DownloadUpgradePackage","DownloadPatchPackage","DownloadUpgradeAndPatchPackages","DownloadAndUpgrade","DownloadAndApplyPatch","DownloadAndUpgradeWithPatch","Upgrade","ApplyPatch","RevertPatch","UpgradeAndPatch","AssessSoftwareUpdate","AbortApplyPatch","AbortUpgrade"]
	OperationType *string `json:"operationType"`

	// Operation Id of cluster operation.
	//
	// Required: true
	OperationID *string `json:"operationId"`
}

// Validate validates this cluster operation type and Id
func (m *ClusterOperationTypeAndID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterOperationTypeAndIdTypeOperationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Destroy","Create","NodeAddition","NodeRemoval","DownloadUpgradePackage","DownloadPatchPackage","DownloadUpgradeAndPatchPackages","DownloadAndUpgrade","DownloadAndApplyPatch","DownloadAndUpgradeWithPatch","Upgrade","ApplyPatch","RevertPatch","UpgradeAndPatch","AssessSoftwareUpdate","AbortApplyPatch","AbortUpgrade"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterOperationTypeAndIdTypeOperationTypePropEnum = append(clusterOperationTypeAndIdTypeOperationTypePropEnum, v)
	}
}

const (

	// ClusterOperationTypeAndIDOperationTypeDestroy captures enum value "Destroy"
	ClusterOperationTypeAndIDOperationTypeDestroy string = "Destroy"

	// ClusterOperationTypeAndIDOperationTypeCreate captures enum value "Create"
	ClusterOperationTypeAndIDOperationTypeCreate string = "Create"

	// ClusterOperationTypeAndIDOperationTypeNodeAddition captures enum value "NodeAddition"
	ClusterOperationTypeAndIDOperationTypeNodeAddition string = "NodeAddition"

	// ClusterOperationTypeAndIDOperationTypeNodeRemoval captures enum value "NodeRemoval"
	ClusterOperationTypeAndIDOperationTypeNodeRemoval string = "NodeRemoval"

	// ClusterOperationTypeAndIDOperationTypeDownloadUpgradePackage captures enum value "DownloadUpgradePackage"
	ClusterOperationTypeAndIDOperationTypeDownloadUpgradePackage string = "DownloadUpgradePackage"

	// ClusterOperationTypeAndIDOperationTypeDownloadPatchPackage captures enum value "DownloadPatchPackage"
	ClusterOperationTypeAndIDOperationTypeDownloadPatchPackage string = "DownloadPatchPackage"

	// ClusterOperationTypeAndIDOperationTypeDownloadUpgradeAndPatchPackages captures enum value "DownloadUpgradeAndPatchPackages"
	ClusterOperationTypeAndIDOperationTypeDownloadUpgradeAndPatchPackages string = "DownloadUpgradeAndPatchPackages"

	// ClusterOperationTypeAndIDOperationTypeDownloadAndUpgrade captures enum value "DownloadAndUpgrade"
	ClusterOperationTypeAndIDOperationTypeDownloadAndUpgrade string = "DownloadAndUpgrade"

	// ClusterOperationTypeAndIDOperationTypeDownloadAndApplyPatch captures enum value "DownloadAndApplyPatch"
	ClusterOperationTypeAndIDOperationTypeDownloadAndApplyPatch string = "DownloadAndApplyPatch"

	// ClusterOperationTypeAndIDOperationTypeDownloadAndUpgradeWithPatch captures enum value "DownloadAndUpgradeWithPatch"
	ClusterOperationTypeAndIDOperationTypeDownloadAndUpgradeWithPatch string = "DownloadAndUpgradeWithPatch"

	// ClusterOperationTypeAndIDOperationTypeUpgrade captures enum value "Upgrade"
	ClusterOperationTypeAndIDOperationTypeUpgrade string = "Upgrade"

	// ClusterOperationTypeAndIDOperationTypeApplyPatch captures enum value "ApplyPatch"
	ClusterOperationTypeAndIDOperationTypeApplyPatch string = "ApplyPatch"

	// ClusterOperationTypeAndIDOperationTypeRevertPatch captures enum value "RevertPatch"
	ClusterOperationTypeAndIDOperationTypeRevertPatch string = "RevertPatch"

	// ClusterOperationTypeAndIDOperationTypeUpgradeAndPatch captures enum value "UpgradeAndPatch"
	ClusterOperationTypeAndIDOperationTypeUpgradeAndPatch string = "UpgradeAndPatch"

	// ClusterOperationTypeAndIDOperationTypeAssessSoftwareUpdate captures enum value "AssessSoftwareUpdate"
	ClusterOperationTypeAndIDOperationTypeAssessSoftwareUpdate string = "AssessSoftwareUpdate"

	// ClusterOperationTypeAndIDOperationTypeAbortApplyPatch captures enum value "AbortApplyPatch"
	ClusterOperationTypeAndIDOperationTypeAbortApplyPatch string = "AbortApplyPatch"

	// ClusterOperationTypeAndIDOperationTypeAbortUpgrade captures enum value "AbortUpgrade"
	ClusterOperationTypeAndIDOperationTypeAbortUpgrade string = "AbortUpgrade"
)

// prop value enum
func (m *ClusterOperationTypeAndID) validateOperationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterOperationTypeAndIdTypeOperationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterOperationTypeAndID) validateOperationType(formats strfmt.Registry) error {

	if err := validate.Required("operationType", "body", m.OperationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperationTypeEnum("operationType", "body", *m.OperationType); err != nil {
		return err
	}

	return nil
}

func (m *ClusterOperationTypeAndID) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster operation type and Id based on context it is used
func (m *ClusterOperationTypeAndID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterOperationTypeAndID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterOperationTypeAndID) UnmarshalBinary(b []byte) error {
	var res ClusterOperationTypeAndID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
