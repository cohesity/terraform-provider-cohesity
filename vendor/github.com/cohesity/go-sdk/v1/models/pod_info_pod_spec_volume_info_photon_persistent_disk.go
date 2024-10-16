// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodInfoPodSpecVolumeInfoPhotonPersistentDisk Represents a Photon Controller persistent disk resource.
//
// swagger:model PodInfo_PodSpec_VolumeInfo_PhotonPersistentDisk
type PodInfoPodSpecVolumeInfoPhotonPersistentDisk struct {

	// fs type
	FsType *string `json:"fsType,omitempty"`

	// pd Id
	PdID *string `json:"pdId,omitempty"`
}

// Validate validates this pod info pod spec volume info photon persistent disk
func (m *PodInfoPodSpecVolumeInfoPhotonPersistentDisk) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod info pod spec volume info photon persistent disk based on context it is used
func (m *PodInfoPodSpecVolumeInfoPhotonPersistentDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodInfoPodSpecVolumeInfoPhotonPersistentDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodInfoPodSpecVolumeInfoPhotonPersistentDisk) UnmarshalBinary(b []byte) error {
	var res PodInfoPodSpecVolumeInfoPhotonPersistentDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
