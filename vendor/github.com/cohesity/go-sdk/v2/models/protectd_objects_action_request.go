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

// ProtectdObjectsActionRequest Specifies the request for performing various actions on protected objects.
//
// swagger:model ProtectdObjectsActionRequest
type ProtectdObjectsActionRequest struct {

	// Specifies the action type to be performed on object getting protected. Based on selected action, provide the action params.
	// Required: true
	// Enum: ["Pause","Resume","UnProtect","ProtectNow"]
	Action *string `json:"action"`

	// Specifies the object action key for any action on the given object.
	// Enum: ["kVMware","kHyperV","kVCD","kAzure","kGCP","kKVM","kAcropolis","kAWS","kAWSNative","kAwsS3","kAWSSnapshotManager","kRDSSnapshotManager","kAuroraSnapshotManager","kAwsRDSPostgresBackup","kAwsRDSPostgres","kAwsAuroraPostgres","kAzureNative","kAzureSQL","kAzureSnapshotManager","kPhysical","kPhysicalFiles","kGPFS","kElastifile","kNetapp","kGenericNas","kIsilon","kFlashBlade","kPure","kIbmFlashSystem","kSQL","kExchange","kAD","kOracle","kView","kRemoteAdapter","kO365","kO365PublicFolders","kO365Teams","kO365Group","kO365Exchange","kO365OneDrive","kO365Sharepoint","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kSAPHANA","kUDA","kSfdc","kO365ExchangeCSM","kO365OneDriveCSM","kO365SharepointCSM"]
	ObjectActionKey *string `json:"objectActionKey,omitempty"`

	// If provided action is 'Pause' then this object should be provided as input.
	PauseParams *ProtectedObjectPauseActionParams `json:"pauseParams,omitempty"`

	// If provided action is 'Resume' then this object should be provided as input.
	ResumeParams *ProtectedObjectResumeActionParams `json:"resumeParams,omitempty"`

	// If provided action is 'RunNow' then this object should be provided as input.
	RunNowParams *ProtectedObjectRunNowActionParams `json:"runNowParams,omitempty"`

	// If provided action is 'UnProtect' then this object should be provided as input.
	UnProtectParams *ProtectedObjectUnProtectActionParams `json:"unProtectParams,omitempty"`

	// Specifies the protections type on which action to be performed. This is used when an object is protected by multiple protection types. If not specified action will be performed on all protection types.
	// Unique: true
	SnapshotBackendTypes []string `json:"snapshotBackendTypes"`
}

// Validate validates this protectd objects action request
func (m *ProtectdObjectsActionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectActionKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePauseParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResumeParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunNowParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnProtectParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotBackendTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var protectdObjectsActionRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pause","Resume","UnProtect","ProtectNow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectdObjectsActionRequestTypeActionPropEnum = append(protectdObjectsActionRequestTypeActionPropEnum, v)
	}
}

const (

	// ProtectdObjectsActionRequestActionPause captures enum value "Pause"
	ProtectdObjectsActionRequestActionPause string = "Pause"

	// ProtectdObjectsActionRequestActionResume captures enum value "Resume"
	ProtectdObjectsActionRequestActionResume string = "Resume"

	// ProtectdObjectsActionRequestActionUnProtect captures enum value "UnProtect"
	ProtectdObjectsActionRequestActionUnProtect string = "UnProtect"

	// ProtectdObjectsActionRequestActionProtectNow captures enum value "ProtectNow"
	ProtectdObjectsActionRequestActionProtectNow string = "ProtectNow"
)

// prop value enum
func (m *ProtectdObjectsActionRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectdObjectsActionRequestTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectdObjectsActionRequest) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

var protectdObjectsActionRequestTypeObjectActionKeyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kVCD","kAzure","kGCP","kKVM","kAcropolis","kAWS","kAWSNative","kAwsS3","kAWSSnapshotManager","kRDSSnapshotManager","kAuroraSnapshotManager","kAwsRDSPostgresBackup","kAwsRDSPostgres","kAwsAuroraPostgres","kAzureNative","kAzureSQL","kAzureSnapshotManager","kPhysical","kPhysicalFiles","kGPFS","kElastifile","kNetapp","kGenericNas","kIsilon","kFlashBlade","kPure","kIbmFlashSystem","kSQL","kExchange","kAD","kOracle","kView","kRemoteAdapter","kO365","kO365PublicFolders","kO365Teams","kO365Group","kO365Exchange","kO365OneDrive","kO365Sharepoint","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kSAPHANA","kUDA","kSfdc","kO365ExchangeCSM","kO365OneDriveCSM","kO365SharepointCSM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectdObjectsActionRequestTypeObjectActionKeyPropEnum = append(protectdObjectsActionRequestTypeObjectActionKeyPropEnum, v)
	}
}

const (

	// ProtectdObjectsActionRequestObjectActionKeyKVMware captures enum value "kVMware"
	ProtectdObjectsActionRequestObjectActionKeyKVMware string = "kVMware"

	// ProtectdObjectsActionRequestObjectActionKeyKHyperV captures enum value "kHyperV"
	ProtectdObjectsActionRequestObjectActionKeyKHyperV string = "kHyperV"

	// ProtectdObjectsActionRequestObjectActionKeyKVCD captures enum value "kVCD"
	ProtectdObjectsActionRequestObjectActionKeyKVCD string = "kVCD"

	// ProtectdObjectsActionRequestObjectActionKeyKAzure captures enum value "kAzure"
	ProtectdObjectsActionRequestObjectActionKeyKAzure string = "kAzure"

	// ProtectdObjectsActionRequestObjectActionKeyKGCP captures enum value "kGCP"
	ProtectdObjectsActionRequestObjectActionKeyKGCP string = "kGCP"

	// ProtectdObjectsActionRequestObjectActionKeyKKVM captures enum value "kKVM"
	ProtectdObjectsActionRequestObjectActionKeyKKVM string = "kKVM"

	// ProtectdObjectsActionRequestObjectActionKeyKAcropolis captures enum value "kAcropolis"
	ProtectdObjectsActionRequestObjectActionKeyKAcropolis string = "kAcropolis"

	// ProtectdObjectsActionRequestObjectActionKeyKAWS captures enum value "kAWS"
	ProtectdObjectsActionRequestObjectActionKeyKAWS string = "kAWS"

	// ProtectdObjectsActionRequestObjectActionKeyKAWSNative captures enum value "kAWSNative"
	ProtectdObjectsActionRequestObjectActionKeyKAWSNative string = "kAWSNative"

	// ProtectdObjectsActionRequestObjectActionKeyKAwsS3 captures enum value "kAwsS3"
	ProtectdObjectsActionRequestObjectActionKeyKAwsS3 string = "kAwsS3"

	// ProtectdObjectsActionRequestObjectActionKeyKAWSSnapshotManager captures enum value "kAWSSnapshotManager"
	ProtectdObjectsActionRequestObjectActionKeyKAWSSnapshotManager string = "kAWSSnapshotManager"

	// ProtectdObjectsActionRequestObjectActionKeyKRDSSnapshotManager captures enum value "kRDSSnapshotManager"
	ProtectdObjectsActionRequestObjectActionKeyKRDSSnapshotManager string = "kRDSSnapshotManager"

	// ProtectdObjectsActionRequestObjectActionKeyKAuroraSnapshotManager captures enum value "kAuroraSnapshotManager"
	ProtectdObjectsActionRequestObjectActionKeyKAuroraSnapshotManager string = "kAuroraSnapshotManager"

	// ProtectdObjectsActionRequestObjectActionKeyKAwsRDSPostgresBackup captures enum value "kAwsRDSPostgresBackup"
	ProtectdObjectsActionRequestObjectActionKeyKAwsRDSPostgresBackup string = "kAwsRDSPostgresBackup"

	// ProtectdObjectsActionRequestObjectActionKeyKAwsRDSPostgres captures enum value "kAwsRDSPostgres"
	ProtectdObjectsActionRequestObjectActionKeyKAwsRDSPostgres string = "kAwsRDSPostgres"

	// ProtectdObjectsActionRequestObjectActionKeyKAwsAuroraPostgres captures enum value "kAwsAuroraPostgres"
	ProtectdObjectsActionRequestObjectActionKeyKAwsAuroraPostgres string = "kAwsAuroraPostgres"

	// ProtectdObjectsActionRequestObjectActionKeyKAzureNative captures enum value "kAzureNative"
	ProtectdObjectsActionRequestObjectActionKeyKAzureNative string = "kAzureNative"

	// ProtectdObjectsActionRequestObjectActionKeyKAzureSQL captures enum value "kAzureSQL"
	ProtectdObjectsActionRequestObjectActionKeyKAzureSQL string = "kAzureSQL"

	// ProtectdObjectsActionRequestObjectActionKeyKAzureSnapshotManager captures enum value "kAzureSnapshotManager"
	ProtectdObjectsActionRequestObjectActionKeyKAzureSnapshotManager string = "kAzureSnapshotManager"

	// ProtectdObjectsActionRequestObjectActionKeyKPhysical captures enum value "kPhysical"
	ProtectdObjectsActionRequestObjectActionKeyKPhysical string = "kPhysical"

	// ProtectdObjectsActionRequestObjectActionKeyKPhysicalFiles captures enum value "kPhysicalFiles"
	ProtectdObjectsActionRequestObjectActionKeyKPhysicalFiles string = "kPhysicalFiles"

	// ProtectdObjectsActionRequestObjectActionKeyKGPFS captures enum value "kGPFS"
	ProtectdObjectsActionRequestObjectActionKeyKGPFS string = "kGPFS"

	// ProtectdObjectsActionRequestObjectActionKeyKElastifile captures enum value "kElastifile"
	ProtectdObjectsActionRequestObjectActionKeyKElastifile string = "kElastifile"

	// ProtectdObjectsActionRequestObjectActionKeyKNetapp captures enum value "kNetapp"
	ProtectdObjectsActionRequestObjectActionKeyKNetapp string = "kNetapp"

	// ProtectdObjectsActionRequestObjectActionKeyKGenericNas captures enum value "kGenericNas"
	ProtectdObjectsActionRequestObjectActionKeyKGenericNas string = "kGenericNas"

	// ProtectdObjectsActionRequestObjectActionKeyKIsilon captures enum value "kIsilon"
	ProtectdObjectsActionRequestObjectActionKeyKIsilon string = "kIsilon"

	// ProtectdObjectsActionRequestObjectActionKeyKFlashBlade captures enum value "kFlashBlade"
	ProtectdObjectsActionRequestObjectActionKeyKFlashBlade string = "kFlashBlade"

	// ProtectdObjectsActionRequestObjectActionKeyKPure captures enum value "kPure"
	ProtectdObjectsActionRequestObjectActionKeyKPure string = "kPure"

	// ProtectdObjectsActionRequestObjectActionKeyKIbmFlashSystem captures enum value "kIbmFlashSystem"
	ProtectdObjectsActionRequestObjectActionKeyKIbmFlashSystem string = "kIbmFlashSystem"

	// ProtectdObjectsActionRequestObjectActionKeyKSQL captures enum value "kSQL"
	ProtectdObjectsActionRequestObjectActionKeyKSQL string = "kSQL"

	// ProtectdObjectsActionRequestObjectActionKeyKExchange captures enum value "kExchange"
	ProtectdObjectsActionRequestObjectActionKeyKExchange string = "kExchange"

	// ProtectdObjectsActionRequestObjectActionKeyKAD captures enum value "kAD"
	ProtectdObjectsActionRequestObjectActionKeyKAD string = "kAD"

	// ProtectdObjectsActionRequestObjectActionKeyKOracle captures enum value "kOracle"
	ProtectdObjectsActionRequestObjectActionKeyKOracle string = "kOracle"

	// ProtectdObjectsActionRequestObjectActionKeyKView captures enum value "kView"
	ProtectdObjectsActionRequestObjectActionKeyKView string = "kView"

	// ProtectdObjectsActionRequestObjectActionKeyKRemoteAdapter captures enum value "kRemoteAdapter"
	ProtectdObjectsActionRequestObjectActionKeyKRemoteAdapter string = "kRemoteAdapter"

	// ProtectdObjectsActionRequestObjectActionKeyKO365 captures enum value "kO365"
	ProtectdObjectsActionRequestObjectActionKeyKO365 string = "kO365"

	// ProtectdObjectsActionRequestObjectActionKeyKO365PublicFolders captures enum value "kO365PublicFolders"
	ProtectdObjectsActionRequestObjectActionKeyKO365PublicFolders string = "kO365PublicFolders"

	// ProtectdObjectsActionRequestObjectActionKeyKO365Teams captures enum value "kO365Teams"
	ProtectdObjectsActionRequestObjectActionKeyKO365Teams string = "kO365Teams"

	// ProtectdObjectsActionRequestObjectActionKeyKO365Group captures enum value "kO365Group"
	ProtectdObjectsActionRequestObjectActionKeyKO365Group string = "kO365Group"

	// ProtectdObjectsActionRequestObjectActionKeyKO365Exchange captures enum value "kO365Exchange"
	ProtectdObjectsActionRequestObjectActionKeyKO365Exchange string = "kO365Exchange"

	// ProtectdObjectsActionRequestObjectActionKeyKO365OneDrive captures enum value "kO365OneDrive"
	ProtectdObjectsActionRequestObjectActionKeyKO365OneDrive string = "kO365OneDrive"

	// ProtectdObjectsActionRequestObjectActionKeyKO365Sharepoint captures enum value "kO365Sharepoint"
	ProtectdObjectsActionRequestObjectActionKeyKO365Sharepoint string = "kO365Sharepoint"

	// ProtectdObjectsActionRequestObjectActionKeyKKubernetes captures enum value "kKubernetes"
	ProtectdObjectsActionRequestObjectActionKeyKKubernetes string = "kKubernetes"

	// ProtectdObjectsActionRequestObjectActionKeyKCassandra captures enum value "kCassandra"
	ProtectdObjectsActionRequestObjectActionKeyKCassandra string = "kCassandra"

	// ProtectdObjectsActionRequestObjectActionKeyKMongoDB captures enum value "kMongoDB"
	ProtectdObjectsActionRequestObjectActionKeyKMongoDB string = "kMongoDB"

	// ProtectdObjectsActionRequestObjectActionKeyKCouchbase captures enum value "kCouchbase"
	ProtectdObjectsActionRequestObjectActionKeyKCouchbase string = "kCouchbase"

	// ProtectdObjectsActionRequestObjectActionKeyKHdfs captures enum value "kHdfs"
	ProtectdObjectsActionRequestObjectActionKeyKHdfs string = "kHdfs"

	// ProtectdObjectsActionRequestObjectActionKeyKHive captures enum value "kHive"
	ProtectdObjectsActionRequestObjectActionKeyKHive string = "kHive"

	// ProtectdObjectsActionRequestObjectActionKeyKHBase captures enum value "kHBase"
	ProtectdObjectsActionRequestObjectActionKeyKHBase string = "kHBase"

	// ProtectdObjectsActionRequestObjectActionKeyKSAPHANA captures enum value "kSAPHANA"
	ProtectdObjectsActionRequestObjectActionKeyKSAPHANA string = "kSAPHANA"

	// ProtectdObjectsActionRequestObjectActionKeyKUDA captures enum value "kUDA"
	ProtectdObjectsActionRequestObjectActionKeyKUDA string = "kUDA"

	// ProtectdObjectsActionRequestObjectActionKeyKSfdc captures enum value "kSfdc"
	ProtectdObjectsActionRequestObjectActionKeyKSfdc string = "kSfdc"

	// ProtectdObjectsActionRequestObjectActionKeyKO365ExchangeCSM captures enum value "kO365ExchangeCSM"
	ProtectdObjectsActionRequestObjectActionKeyKO365ExchangeCSM string = "kO365ExchangeCSM"

	// ProtectdObjectsActionRequestObjectActionKeyKO365OneDriveCSM captures enum value "kO365OneDriveCSM"
	ProtectdObjectsActionRequestObjectActionKeyKO365OneDriveCSM string = "kO365OneDriveCSM"

	// ProtectdObjectsActionRequestObjectActionKeyKO365SharepointCSM captures enum value "kO365SharepointCSM"
	ProtectdObjectsActionRequestObjectActionKeyKO365SharepointCSM string = "kO365SharepointCSM"
)

// prop value enum
func (m *ProtectdObjectsActionRequest) validateObjectActionKeyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectdObjectsActionRequestTypeObjectActionKeyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectdObjectsActionRequest) validateObjectActionKey(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectActionKey) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectActionKeyEnum("objectActionKey", "body", *m.ObjectActionKey); err != nil {
		return err
	}

	return nil
}

func (m *ProtectdObjectsActionRequest) validatePauseParams(formats strfmt.Registry) error {
	if swag.IsZero(m.PauseParams) { // not required
		return nil
	}

	if m.PauseParams != nil {
		if err := m.PauseParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pauseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pauseParams")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectdObjectsActionRequest) validateResumeParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ResumeParams) { // not required
		return nil
	}

	if m.ResumeParams != nil {
		if err := m.ResumeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resumeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resumeParams")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectdObjectsActionRequest) validateRunNowParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RunNowParams) { // not required
		return nil
	}

	if m.RunNowParams != nil {
		if err := m.RunNowParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runNowParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runNowParams")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectdObjectsActionRequest) validateUnProtectParams(formats strfmt.Registry) error {
	if swag.IsZero(m.UnProtectParams) { // not required
		return nil
	}

	if m.UnProtectParams != nil {
		if err := m.UnProtectParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unProtectParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unProtectParams")
			}
			return err
		}
	}

	return nil
}

var protectdObjectsActionRequestSnapshotBackendTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAWSNative","kAWSSnapshotManager","kPhysical","kSQL","kOracle","kRDSSnapshotManager","kAuroraSnapshotManager","kAwsS3","kAwsRDSPostgresBackup","kAzureNative","kAzureSnapshotManager","kAzureSQL","kAwsAuroraPostgres","kAwsRDSPostgres"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectdObjectsActionRequestSnapshotBackendTypesItemsEnum = append(protectdObjectsActionRequestSnapshotBackendTypesItemsEnum, v)
	}
}

func (m *ProtectdObjectsActionRequest) validateSnapshotBackendTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectdObjectsActionRequestSnapshotBackendTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectdObjectsActionRequest) validateSnapshotBackendTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotBackendTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("snapshotBackendTypes", "body", m.SnapshotBackendTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.SnapshotBackendTypes); i++ {

		// value enum
		if err := m.validateSnapshotBackendTypesItemsEnum("snapshotBackendTypes"+"."+strconv.Itoa(i), "body", m.SnapshotBackendTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this protectd objects action request based on the context it is used
func (m *ProtectdObjectsActionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePauseParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResumeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunNowParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnProtectParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectdObjectsActionRequest) contextValidatePauseParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PauseParams != nil {

		if swag.IsZero(m.PauseParams) { // not required
			return nil
		}

		if err := m.PauseParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pauseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pauseParams")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectdObjectsActionRequest) contextValidateResumeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ResumeParams != nil {

		if swag.IsZero(m.ResumeParams) { // not required
			return nil
		}

		if err := m.ResumeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resumeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resumeParams")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectdObjectsActionRequest) contextValidateRunNowParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RunNowParams != nil {

		if swag.IsZero(m.RunNowParams) { // not required
			return nil
		}

		if err := m.RunNowParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runNowParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runNowParams")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectdObjectsActionRequest) contextValidateUnProtectParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UnProtectParams != nil {

		if swag.IsZero(m.UnProtectParams) { // not required
			return nil
		}

		if err := m.UnProtectParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unProtectParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unProtectParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectdObjectsActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectdObjectsActionRequest) UnmarshalBinary(b []byte) error {
	var res ProtectdObjectsActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
