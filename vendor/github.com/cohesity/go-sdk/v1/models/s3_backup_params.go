// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// S3BackupParams s3 backup params
//
// swagger:model S3BackupParams
type S3BackupParams struct {

	// Type of backup for this aws S3 backup task.
	BackupType *int32 `json:"backupType,omitempty"`

	// Do check for SQS event older than inventory_report_create_secs.
	CheckInventoryReportSqsCoverage *bool `json:"checkInventoryReportSqsCoverage,omitempty"`

	// fusion fs node Id
	FusionFsNodeID *int64 `json:"fusionFsNodeId,omitempty"`

	// Skip prev marker check in SQS fetcher op.
	IgnorePrevSqsMarker *bool `json:"ignorePrevSqsMarker,omitempty"`

	// If run is using an inventory report, name of manifest file used.
	InventoryReportManifestFile *string `json:"inventoryReportManifestFile,omitempty"`
}

// Validate validates this s3 backup params
func (m *S3BackupParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 backup params based on context it is used
func (m *S3BackupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *S3BackupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BackupParams) UnmarshalBinary(b []byte) error {
	var res S3BackupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
