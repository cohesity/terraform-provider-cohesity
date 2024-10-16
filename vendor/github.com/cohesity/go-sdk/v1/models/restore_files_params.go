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

// RestoreFilesParams This message captures all the necessary arguments specified by the user to
// restore files to the source.
//
// swagger:model RestoreFilesParams
type RestoreFilesParams struct {

	// A list of target IP addresses that should not be used.
	BlacklistedIPAddrs []string `json:"blacklistedIpAddrs"`

	// Destination endpoint UUID for source s3 objectstore.
	DestinationEpUUID *string `json:"destinationEpUuid,omitempty"`

	// Directory name security style map contains mapping of the directory name
	// to security style it supports.  This is needed to restore the same
	// permission for the given directory for Qtrees.
	DirectoryNameSecurityStyleMap map[string]string `json:"directoryNameSecurityStyleMap,omitempty"`

	// Glacier restore option chosen by the user.
	GlacierFlrRestoreOption *int32 `json:"glacierFlrRestoreOption,omitempty"`

	// Whether this is a file restore operation from an archive.
	IsArchiveFlr *bool `json:"isArchiveFlr,omitempty"`

	// Whether this is a file based volume restore.
	IsFileVolumeRestore *bool `json:"isFileVolumeRestore,omitempty"`

	// Whether this is a mount based file restore operation
	IsMountBasedFlr *bool `json:"isMountBasedFlr,omitempty"`

	// Whether this is a source initiated restore.
	IsSourceInitiatedRestore *bool `json:"isSourceInitiatedRestore,omitempty"`

	// This is applicable if target entity is of kIsilon type.
	IsilonEnvParams *IsilonEnvParams `json:"isilonEnvParams,omitempty"`

	// Whether this will attach disks or mount disks on the VM side OR use
	// Storage Proxy RPCs to stream data.
	MountDisksOnVM *bool `json:"mountDisksOnVm,omitempty"`

	// Used to determining filtering_policy for NAS Migration uptier operation.
	// Currently this is set only if this is NAS Migration uptier operation.
	NasBackupParams *NasBackupParams `json:"nasBackupParams,omitempty"`

	// The NAS protocol type(s) of this restore task. Currently we only support a
	// single restore type. This field is only populated for NAS restore tasks.
	NasProtocolTypeVec []int32 `json:"nasProtocolTypeVec"`

	// Object store config name for source initiated backup.
	ObjectstoreConfigName *string `json:"objectstoreConfigName,omitempty"`

	// If enabled, magneto physical file restore will be enabled via job
	// framework
	PhysicalFlrParallelRestore *bool `json:"physicalFlrParallelRestore,omitempty"`

	// If the proxy entity is specified, then the virtual disks are attached to
	// the proxy server and the file copy will be initiated through this server.
	ProxyEntity *EntityProto `json:"proxyEntity,omitempty"`

	// If the proxy entity above is specified, then it's parent entity must be
	// specified as well.
	ProxyEntityParentSource *EntityProto `json:"proxyEntityParentSource,omitempty"`

	// Preferences for the restore files operation.
	RestoreFilesPreferences *RestoreFilesPreferences `json:"restoreFilesPreferences,omitempty"`

	// Determines the type of method to be used to perform FLR.
	RestoreMethod *int32 `json:"restoreMethod,omitempty"`

	// Information regarding files and directories.
	RestoredFileInfoVec []*RestoredFileInfo `json:"restoredFileInfoVec"`

	// This message captures all the details of S3 view from where the
	// data is restored.
	S3Viewbackupproperties *S3ViewBackupProperties `json:"s3Viewbackupproperties,omitempty"`

	// Snapshot name need by source to start the restore.
	SourceSnapshotName *string `json:"sourceSnapshotName,omitempty"`

	// The following fields only need to be set if restoring to a target location
	// (i.e., not downloading the files directly).
	//
	// Target entity where the files are being restored to.
	TargetEntity *EntityProto `json:"targetEntity,omitempty"`

	// Credentials to access the target entity such as a VM. In case of physical
	// server, the copy process will be launched as this user.
	// NOTE: If proxy entity is specified, then this credentials will be used for
	// running operations on proxy server as well.
	TargetEntityCredentials *Credentials `json:"targetEntityCredentials,omitempty"`

	// The registered source (i.e vCenter or ESXi host in VMware environment)
	// under which the target entity is present.
	TargetEntityParentSource *EntityProto `json:"targetEntityParentSource,omitempty"`

	// Host entity where the target entity resides. This is only populated for
	// Netapp environments right now if target_entity_parent_source is a
	// cluster. For example, the host of a target Netapp volume will be the
	// vserver it belongs to.
	TargetHostEntity *EntityProto `json:"targetHostEntity,omitempty"`

	// The host environment type. This is set in VMware environment to
	// indicate the OS type of the target entity.
	// NOTE: This is expected to be set since magneto does not know the host type
	// for VMware entities.
	TargetHostType *int32 `json:"targetHostType,omitempty"`

	// target pvc entity
	TargetPvcEntity *EntityProto `json:"targetPvcEntity,omitempty"`

	// Set if this is NAS Migration uptier operation.
	UptierParams *FileUptieringParams `json:"uptierParams,omitempty"`

	// Whether this will use an existing agent on the target VM to do the
	// restore.
	// This field is deprecated. restore_method should be used for populating
	// use of existing agent.
	UseExistingAgent *bool `json:"useExistingAgent,omitempty"`

	// View ID
	ViewID *int64 `json:"viewId,omitempty"`

	// Name of the S3 view
	ViewName *string `json:"viewName,omitempty"`

	// Entity of the VPC connector, required in case of GCP FLR.
	VpcConnectorEntity *EntityProto `json:"vpcConnectorEntity,omitempty"`

	// A list of target IP addresses that should be used exclusively.
	WhitelistedIPAddrs []string `json:"whitelistedIpAddrs"`
}

// Validate validates this restore files params
func (m *RestoreFilesParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsilonEnvParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNasBackupParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyEntityParentSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreFilesPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoredFileInfoVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3Viewbackupproperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEntityCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEntityParentSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetHostEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetPvcEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUptierParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcConnectorEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreFilesParams) validateIsilonEnvParams(formats strfmt.Registry) error {
	if swag.IsZero(m.IsilonEnvParams) { // not required
		return nil
	}

	if m.IsilonEnvParams != nil {
		if err := m.IsilonEnvParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isilonEnvParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isilonEnvParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateNasBackupParams(formats strfmt.Registry) error {
	if swag.IsZero(m.NasBackupParams) { // not required
		return nil
	}

	if m.NasBackupParams != nil {
		if err := m.NasBackupParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nasBackupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nasBackupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateProxyEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyEntity) { // not required
		return nil
	}

	if m.ProxyEntity != nil {
		if err := m.ProxyEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxyEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateProxyEntityParentSource(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyEntityParentSource) { // not required
		return nil
	}

	if m.ProxyEntityParentSource != nil {
		if err := m.ProxyEntityParentSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyEntityParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxyEntityParentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateRestoreFilesPreferences(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreFilesPreferences) { // not required
		return nil
	}

	if m.RestoreFilesPreferences != nil {
		if err := m.RestoreFilesPreferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreFilesPreferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreFilesPreferences")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateRestoredFileInfoVec(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoredFileInfoVec) { // not required
		return nil
	}

	for i := 0; i < len(m.RestoredFileInfoVec); i++ {
		if swag.IsZero(m.RestoredFileInfoVec[i]) { // not required
			continue
		}

		if m.RestoredFileInfoVec[i] != nil {
			if err := m.RestoredFileInfoVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoredFileInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoredFileInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestoreFilesParams) validateS3Viewbackupproperties(formats strfmt.Registry) error {
	if swag.IsZero(m.S3Viewbackupproperties) { // not required
		return nil
	}

	if m.S3Viewbackupproperties != nil {
		if err := m.S3Viewbackupproperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3Viewbackupproperties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3Viewbackupproperties")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateTargetEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetEntity) { // not required
		return nil
	}

	if m.TargetEntity != nil {
		if err := m.TargetEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateTargetEntityCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetEntityCredentials) { // not required
		return nil
	}

	if m.TargetEntityCredentials != nil {
		if err := m.TargetEntityCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntityCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntityCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateTargetEntityParentSource(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetEntityParentSource) { // not required
		return nil
	}

	if m.TargetEntityParentSource != nil {
		if err := m.TargetEntityParentSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntityParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntityParentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateTargetHostEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetHostEntity) { // not required
		return nil
	}

	if m.TargetHostEntity != nil {
		if err := m.TargetHostEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetHostEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetHostEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateTargetPvcEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetPvcEntity) { // not required
		return nil
	}

	if m.TargetPvcEntity != nil {
		if err := m.TargetPvcEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetPvcEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetPvcEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateUptierParams(formats strfmt.Registry) error {
	if swag.IsZero(m.UptierParams) { // not required
		return nil
	}

	if m.UptierParams != nil {
		if err := m.UptierParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uptierParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uptierParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) validateVpcConnectorEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcConnectorEntity) { // not required
		return nil
	}

	if m.VpcConnectorEntity != nil {
		if err := m.VpcConnectorEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpcConnectorEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpcConnectorEntity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this restore files params based on the context it is used
func (m *RestoreFilesParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIsilonEnvParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNasBackupParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxyEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxyEntityParentSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoreFilesPreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoredFileInfoVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3Viewbackupproperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetEntityCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetEntityParentSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetHostEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetPvcEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUptierParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpcConnectorEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreFilesParams) contextValidateIsilonEnvParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IsilonEnvParams != nil {

		if swag.IsZero(m.IsilonEnvParams) { // not required
			return nil
		}

		if err := m.IsilonEnvParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isilonEnvParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isilonEnvParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateNasBackupParams(ctx context.Context, formats strfmt.Registry) error {

	if m.NasBackupParams != nil {

		if swag.IsZero(m.NasBackupParams) { // not required
			return nil
		}

		if err := m.NasBackupParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nasBackupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nasBackupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateProxyEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.ProxyEntity != nil {

		if swag.IsZero(m.ProxyEntity) { // not required
			return nil
		}

		if err := m.ProxyEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxyEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateProxyEntityParentSource(ctx context.Context, formats strfmt.Registry) error {

	if m.ProxyEntityParentSource != nil {

		if swag.IsZero(m.ProxyEntityParentSource) { // not required
			return nil
		}

		if err := m.ProxyEntityParentSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyEntityParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxyEntityParentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateRestoreFilesPreferences(ctx context.Context, formats strfmt.Registry) error {

	if m.RestoreFilesPreferences != nil {

		if swag.IsZero(m.RestoreFilesPreferences) { // not required
			return nil
		}

		if err := m.RestoreFilesPreferences.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreFilesPreferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreFilesPreferences")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateRestoredFileInfoVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RestoredFileInfoVec); i++ {

		if m.RestoredFileInfoVec[i] != nil {

			if swag.IsZero(m.RestoredFileInfoVec[i]) { // not required
				return nil
			}

			if err := m.RestoredFileInfoVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoredFileInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoredFileInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestoreFilesParams) contextValidateS3Viewbackupproperties(ctx context.Context, formats strfmt.Registry) error {

	if m.S3Viewbackupproperties != nil {

		if swag.IsZero(m.S3Viewbackupproperties) { // not required
			return nil
		}

		if err := m.S3Viewbackupproperties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3Viewbackupproperties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3Viewbackupproperties")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateTargetEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetEntity != nil {

		if swag.IsZero(m.TargetEntity) { // not required
			return nil
		}

		if err := m.TargetEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateTargetEntityCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetEntityCredentials != nil {

		if swag.IsZero(m.TargetEntityCredentials) { // not required
			return nil
		}

		if err := m.TargetEntityCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntityCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntityCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateTargetEntityParentSource(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetEntityParentSource != nil {

		if swag.IsZero(m.TargetEntityParentSource) { // not required
			return nil
		}

		if err := m.TargetEntityParentSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntityParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntityParentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateTargetHostEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetHostEntity != nil {

		if swag.IsZero(m.TargetHostEntity) { // not required
			return nil
		}

		if err := m.TargetHostEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetHostEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetHostEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateTargetPvcEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetPvcEntity != nil {

		if swag.IsZero(m.TargetPvcEntity) { // not required
			return nil
		}

		if err := m.TargetPvcEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetPvcEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetPvcEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateUptierParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UptierParams != nil {

		if swag.IsZero(m.UptierParams) { // not required
			return nil
		}

		if err := m.UptierParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uptierParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uptierParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesParams) contextValidateVpcConnectorEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcConnectorEntity != nil {

		if swag.IsZero(m.VpcConnectorEntity) { // not required
			return nil
		}

		if err := m.VpcConnectorEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpcConnectorEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpcConnectorEntity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreFilesParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreFilesParams) UnmarshalBinary(b []byte) error {
	var res RestoreFilesParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
