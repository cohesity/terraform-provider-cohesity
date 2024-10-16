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

// RestoreFilesTaskRequest Restore Task.
//
// Specifies information about a Restore Task that recovers files and
// folders.
//
// swagger:model RestoreFilesTaskRequest
type RestoreFilesTaskRequest struct {

	// Specifies the cloud credentials used to authenticate with cloud(Aws).
	CloudCredentials *CloudCredentials `json:"cloudCredentials,omitempty"`

	// Specifies if the Restore Task should continue even if the copy operation
	// of some files and folders fails. If true, the Cohesity Cluster
	// ignores intermittent errors and recovers as many files and folders
	// as possible.
	// If false, the Restore Task stops recovering when a copy operation fails.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies the type of method to be used to perform file recovery.
	// 'kAutoDeploy' indicates that file restore operation wiil be performed using
	// an ephemeral agent.
	// 'kUseExistingAgent' indicates that file restore operation wiil be performed
	// using an persistent agent.
	// 'kUseHypervisorAPIs' indicates that file restore operation wiil be performed
	// using an hypervisor API's.
	// Enum: ["kAutoDeploy","kUseExistingAgent","kUseHypervisorAPIs"]
	FileRecoveryMethod *string `json:"fileRecoveryMethod,omitempty"`

	// Array of Files or Folders.
	//
	// Specifies the files and folders to recover from the snapshot.
	Filenames []string `json:"filenames"`

	// Specifies the list of IP addresses that are allowed or denied during
	// restore. Allowed IPs and Denied IPs cannot be used together.
	FilterIPConfig *FilterIPConfig `json:"filterIpConfig,omitempty"`

	// Specifies whether this is a file based volume restore.
	IsFileBasedVolumeRestore *bool `json:"isFileBasedVolumeRestore,omitempty"`

	// Sepcifies whether this will attach disks or mount disks on the VM side
	// OR use Storage Proxy RPCs to stream data
	MountDisksOnVM *bool `json:"mountDisksOnVm,omitempty"`

	// Specifies the name of the Restore Task. This field must be set and
	// must be a unique name.
	Name *string `json:"name,omitempty"`

	// Specifies an optional root folder where to recover the selected
	// files and folders.
	// By default, files and folders are restored to their original path.
	NewBaseDirectory *string `json:"newBaseDirectory,omitempty"`

	// If true, any existing files and folders on the operating system
	// are overwritten by the recovered files or folders. This is the default.
	// If false, existing files and folders are not overwritten.
	Overwrite *bool `json:"overwrite,omitempty"`

	// Specifies password of the username to access the target source.
	Password *string `json:"password,omitempty"`

	// If true, the Restore Tasks preserves the original file and
	// folder attributes. This is the default.
	PreserveAttributes *bool `json:"preserveAttributes,omitempty"`

	// Specifies information regarding files and directories.
	RestoredFileInfoList []*RestoredFileInfoList `json:"restoredFileInfoList"`

	// Restore Object.
	//
	// Specifies information about the source object (such as a VM)
	// that contains the files and folders to recover.
	// In addition, it contains information about the Protection Job and Job
	// Run that captured the snapshot to recover from.
	// To specify a particular snapshot, you must specify a
	// jobRunId and a startTimeUsecs.
	// If jobRunId and startTimeUsecs are not specified,
	// the last Job Run of the specified Job is used.
	SourceObjectInfo struct {
		RestoreObjectDetails
	} `json:"sourceObjectInfo,omitempty"`

	// Specifies the target host types to be restored.
	// 'kLinux' indicates the Linux operating system.
	// 'kWindows' indicates the Microsoft Windows operating system.
	// 'kAix' indicates the IBM AIX operating system.
	// 'kSolaris' indicates the Oracle Solaris operating system.
	// 'kSapHana' indicates the Sap Hana database system developed by SAP SE.
	// 'kSapOracle' indicates the Sap Oracle database system developed by SAP SE.
	// 'kCockroachDB' indicates the CockroachDB database system.
	// 'kMySQL' indicates the MySQL database system.
	// 'kSapSybase' indicates the SapSybase database system.
	// 'kSapMaxDB' indicates the SapMaxDB database system.
	// 'kSapSybaseIQ' indicates the SapSybaseIQ database system.
	// 'kDB2' indicates the DB2 database system.
	// 'kSapASE' indicates the SapASE database system.
	// 'kMariaDB' indicates the MariaDB database system.
	// 'kPostgreSQL' indicates the PostgreSQL database system.
	// 'kHPUX' indicates the HPUX database system.
	// 'kVOS' indicates the VOS database system.
	// 'kOther' indicates the other types of operating system.
	// Enum: ["kLinux","kWindows","kAix","kSolaris","kSapHana","kSapOracle","kCockroachDB","kMySQL","kOther","kSapSybase","kSapMaxDB","kSapSybaseIQ","kDB2","kSapASE","kMariaDB","kPostgreSQL","kVOS","kHPUX"]
	TargetHostType *string `json:"targetHostType,omitempty"`

	// Specifies the registered source (such as a vCenter Server)
	// that contains the target protection source (such as a VM)
	// where the files and folders are recovered to.
	// This field is not required for a Physical Server.
	TargetParentSourceID *int64 `json:"targetParentSourceId,omitempty"`

	// Specifies the id of the target protection source (such as a VM)
	// where the files and folders are recovered to.
	TargetSourceID *int64 `json:"targetSourceId,omitempty"`

	// Specifies whether this will use an existing agent on the target vm
	// to do restore. Following field is deprecated and shall not be used.
	// Please refer to the FileRecoveryMethod field for more information.
	UseExistingAgent *bool `json:"useExistingAgent,omitempty"`

	// Specifies username to access the target source.
	Username *string `json:"username,omitempty"`
}

// Validate validates this restore files task request
func (m *RestoreFilesTaskRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileRecoveryMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterIPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoredFileInfoList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceObjectInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetHostType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreFilesTaskRequest) validateCloudCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudCredentials) { // not required
		return nil
	}

	if m.CloudCredentials != nil {
		if err := m.CloudCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudCredentials")
			}
			return err
		}
	}

	return nil
}

var restoreFilesTaskRequestTypeFileRecoveryMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAutoDeploy","kUseExistingAgent","kUseHypervisorAPIs"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restoreFilesTaskRequestTypeFileRecoveryMethodPropEnum = append(restoreFilesTaskRequestTypeFileRecoveryMethodPropEnum, v)
	}
}

const (

	// RestoreFilesTaskRequestFileRecoveryMethodKAutoDeploy captures enum value "kAutoDeploy"
	RestoreFilesTaskRequestFileRecoveryMethodKAutoDeploy string = "kAutoDeploy"

	// RestoreFilesTaskRequestFileRecoveryMethodKUseExistingAgent captures enum value "kUseExistingAgent"
	RestoreFilesTaskRequestFileRecoveryMethodKUseExistingAgent string = "kUseExistingAgent"

	// RestoreFilesTaskRequestFileRecoveryMethodKUseHypervisorAPIs captures enum value "kUseHypervisorAPIs"
	RestoreFilesTaskRequestFileRecoveryMethodKUseHypervisorAPIs string = "kUseHypervisorAPIs"
)

// prop value enum
func (m *RestoreFilesTaskRequest) validateFileRecoveryMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, restoreFilesTaskRequestTypeFileRecoveryMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RestoreFilesTaskRequest) validateFileRecoveryMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.FileRecoveryMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateFileRecoveryMethodEnum("fileRecoveryMethod", "body", *m.FileRecoveryMethod); err != nil {
		return err
	}

	return nil
}

func (m *RestoreFilesTaskRequest) validateFilterIPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterIPConfig) { // not required
		return nil
	}

	if m.FilterIPConfig != nil {
		if err := m.FilterIPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterIpConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filterIpConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesTaskRequest) validateRestoredFileInfoList(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoredFileInfoList) { // not required
		return nil
	}

	for i := 0; i < len(m.RestoredFileInfoList); i++ {
		if swag.IsZero(m.RestoredFileInfoList[i]) { // not required
			continue
		}

		if m.RestoredFileInfoList[i] != nil {
			if err := m.RestoredFileInfoList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoredFileInfoList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoredFileInfoList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestoreFilesTaskRequest) validateSourceObjectInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceObjectInfo) { // not required
		return nil
	}

	return nil
}

var restoreFilesTaskRequestTypeTargetHostTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kLinux","kWindows","kAix","kSolaris","kSapHana","kSapOracle","kCockroachDB","kMySQL","kOther","kSapSybase","kSapMaxDB","kSapSybaseIQ","kDB2","kSapASE","kMariaDB","kPostgreSQL","kVOS","kHPUX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restoreFilesTaskRequestTypeTargetHostTypePropEnum = append(restoreFilesTaskRequestTypeTargetHostTypePropEnum, v)
	}
}

const (

	// RestoreFilesTaskRequestTargetHostTypeKLinux captures enum value "kLinux"
	RestoreFilesTaskRequestTargetHostTypeKLinux string = "kLinux"

	// RestoreFilesTaskRequestTargetHostTypeKWindows captures enum value "kWindows"
	RestoreFilesTaskRequestTargetHostTypeKWindows string = "kWindows"

	// RestoreFilesTaskRequestTargetHostTypeKAix captures enum value "kAix"
	RestoreFilesTaskRequestTargetHostTypeKAix string = "kAix"

	// RestoreFilesTaskRequestTargetHostTypeKSolaris captures enum value "kSolaris"
	RestoreFilesTaskRequestTargetHostTypeKSolaris string = "kSolaris"

	// RestoreFilesTaskRequestTargetHostTypeKSapHana captures enum value "kSapHana"
	RestoreFilesTaskRequestTargetHostTypeKSapHana string = "kSapHana"

	// RestoreFilesTaskRequestTargetHostTypeKSapOracle captures enum value "kSapOracle"
	RestoreFilesTaskRequestTargetHostTypeKSapOracle string = "kSapOracle"

	// RestoreFilesTaskRequestTargetHostTypeKCockroachDB captures enum value "kCockroachDB"
	RestoreFilesTaskRequestTargetHostTypeKCockroachDB string = "kCockroachDB"

	// RestoreFilesTaskRequestTargetHostTypeKMySQL captures enum value "kMySQL"
	RestoreFilesTaskRequestTargetHostTypeKMySQL string = "kMySQL"

	// RestoreFilesTaskRequestTargetHostTypeKOther captures enum value "kOther"
	RestoreFilesTaskRequestTargetHostTypeKOther string = "kOther"

	// RestoreFilesTaskRequestTargetHostTypeKSapSybase captures enum value "kSapSybase"
	RestoreFilesTaskRequestTargetHostTypeKSapSybase string = "kSapSybase"

	// RestoreFilesTaskRequestTargetHostTypeKSapMaxDB captures enum value "kSapMaxDB"
	RestoreFilesTaskRequestTargetHostTypeKSapMaxDB string = "kSapMaxDB"

	// RestoreFilesTaskRequestTargetHostTypeKSapSybaseIQ captures enum value "kSapSybaseIQ"
	RestoreFilesTaskRequestTargetHostTypeKSapSybaseIQ string = "kSapSybaseIQ"

	// RestoreFilesTaskRequestTargetHostTypeKDB2 captures enum value "kDB2"
	RestoreFilesTaskRequestTargetHostTypeKDB2 string = "kDB2"

	// RestoreFilesTaskRequestTargetHostTypeKSapASE captures enum value "kSapASE"
	RestoreFilesTaskRequestTargetHostTypeKSapASE string = "kSapASE"

	// RestoreFilesTaskRequestTargetHostTypeKMariaDB captures enum value "kMariaDB"
	RestoreFilesTaskRequestTargetHostTypeKMariaDB string = "kMariaDB"

	// RestoreFilesTaskRequestTargetHostTypeKPostgreSQL captures enum value "kPostgreSQL"
	RestoreFilesTaskRequestTargetHostTypeKPostgreSQL string = "kPostgreSQL"

	// RestoreFilesTaskRequestTargetHostTypeKVOS captures enum value "kVOS"
	RestoreFilesTaskRequestTargetHostTypeKVOS string = "kVOS"

	// RestoreFilesTaskRequestTargetHostTypeKHPUX captures enum value "kHPUX"
	RestoreFilesTaskRequestTargetHostTypeKHPUX string = "kHPUX"
)

// prop value enum
func (m *RestoreFilesTaskRequest) validateTargetHostTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, restoreFilesTaskRequestTypeTargetHostTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RestoreFilesTaskRequest) validateTargetHostType(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetHostType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetHostTypeEnum("targetHostType", "body", *m.TargetHostType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this restore files task request based on the context it is used
func (m *RestoreFilesTaskRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterIPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoredFileInfoList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceObjectInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreFilesTaskRequest) contextValidateCloudCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudCredentials != nil {

		if swag.IsZero(m.CloudCredentials) { // not required
			return nil
		}

		if err := m.CloudCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesTaskRequest) contextValidateFilterIPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterIPConfig != nil {

		if swag.IsZero(m.FilterIPConfig) { // not required
			return nil
		}

		if err := m.FilterIPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterIpConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filterIpConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreFilesTaskRequest) contextValidateRestoredFileInfoList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RestoredFileInfoList); i++ {

		if m.RestoredFileInfoList[i] != nil {

			if swag.IsZero(m.RestoredFileInfoList[i]) { // not required
				return nil
			}

			if err := m.RestoredFileInfoList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("restoredFileInfoList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("restoredFileInfoList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestoreFilesTaskRequest) contextValidateSourceObjectInfo(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreFilesTaskRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreFilesTaskRequest) UnmarshalBinary(b []byte) error {
	var res RestoreFilesTaskRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
