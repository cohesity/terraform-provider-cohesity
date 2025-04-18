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

// QoS QoS.
//
// Specifies the Quality of Service (QoS) Policy for the View.
//
// swagger:model QoS
type QoS struct {

	// Specifies the id of the QoS Policy used for the View. (Deprecated)
	// This parameter is deprecated and shall not be supported in future releases.
	// Use name instead.
	PrincipalID *int64 `json:"principalId,omitempty"`

	// Specifies the name of the QoS Policy. If not specified,
	// the default is 'Backup Target Low'.
	// (To be deprecated in future release, use name instead)
	// Enum: ["Backup Target High","Backup Target Low","TestAndDev High","TestAndDev Low","Backup Target SSD","Backup Target Commvault","Journaled Sequential Dump","Backup Target Auto"]
	PrincipalName *string `json:"principalName,omitempty"`

	// Specifies the name of the QoS Policy. If not specified,
	// the default is 'BackupTargetLow'.
	//
	// BackupTargetAuto: (Applicable only for C6K Platform) Use this policy
	// for workloads such as backups, which keep many I/Os outstanding.
	// This policy splits I/Os across SSDs and HDDs to achieve maximum
	// performance based on the current usage.
	// The priority for processing workload with this policy is the same as
	// Backup Target High and Backup Target SSD.
	//
	// JournaledSequentialDump: Use this policy for workloads that write
	// large amounts of data sequentially to a very small number of files
	// and do not keep many I/Os outstanding. By default data is written to
	// the SSD and has the highest priority and low latency.
	//
	// TestAndDevHigh: Use this policy for workloads that require lower I/O
	// latency or do not keep many I/Os outstanding, as the I/Os are given
	// higher priority compared to other QoS policies.
	// Data is written to the SSD.
	//
	// TestAndDevLow: The same as TestAndDev High, except that the I/Os with
	// this QoS policy are given lower priority
	// compared to I/Os with TestAndDev High when there is contention.
	//
	// BackupTargetCommvault: Use this policy to intelligently detect and
	// exclude application-specific markers to achieve better deduplication
	// when CommVault backup application is writing to a Cohesity View.
	// Data is written to SSD and has the same priority and latency as
	// TestAndDev High.
	//
	// BackupTargetSSD: Use this policy for workloads such as backups,
	// which keep many I/Os outstanding, but in this case,
	// DataPlatform sends both sequential and random I/Os to SSD.
	// The latency is lower than other Backup Target policies.
	// The priority for processing workload with this policy
	// is the same as Backup Target Auto.
	//
	// BackupTargetHigh: Use this policy for non-latency sensitive workloads
	// such as backups, which keep many I/Os outstanding.
	// Data is written to HDD and has higher latency compared to other QoS
	// policies writing to a SSD The priority for processing workload with
	// this policy is the same as Backup Target Auto.
	//
	// BackupTargetLow: The same as Backup Target High, except that the
	// priority for processing workloads with this policy is lower than
	// workloads with Backup Target Auto / High /SSD
	// when there is contention.
	// Enum: ["BackupTargetHigh","BackupTargetLow","TestAndDevHigh","TestAndDevLow","BackupTargetSSD","BackupTargetCommvault","JournaledSequentialDump","BackupTargetAuto"]
	Name *string `json:"name,omitempty"`
}

// Validate validates this qo s
func (m *QoS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrincipalName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var qoSTypePrincipalNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Backup Target High","Backup Target Low","TestAndDev High","TestAndDev Low","Backup Target SSD","Backup Target Commvault","Journaled Sequential Dump","Backup Target Auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		qoSTypePrincipalNamePropEnum = append(qoSTypePrincipalNamePropEnum, v)
	}
}

const (

	// QoSPrincipalNameBackupTargetHigh captures enum value "Backup Target High"
	QoSPrincipalNameBackupTargetHigh string = "Backup Target High"

	// QoSPrincipalNameBackupTargetLow captures enum value "Backup Target Low"
	QoSPrincipalNameBackupTargetLow string = "Backup Target Low"

	// QoSPrincipalNameTestAndDevHigh captures enum value "TestAndDev High"
	QoSPrincipalNameTestAndDevHigh string = "TestAndDev High"

	// QoSPrincipalNameTestAndDevLow captures enum value "TestAndDev Low"
	QoSPrincipalNameTestAndDevLow string = "TestAndDev Low"

	// QoSPrincipalNameBackupTargetSSD captures enum value "Backup Target SSD"
	QoSPrincipalNameBackupTargetSSD string = "Backup Target SSD"

	// QoSPrincipalNameBackupTargetCommvault captures enum value "Backup Target Commvault"
	QoSPrincipalNameBackupTargetCommvault string = "Backup Target Commvault"

	// QoSPrincipalNameJournaledSequentialDump captures enum value "Journaled Sequential Dump"
	QoSPrincipalNameJournaledSequentialDump string = "Journaled Sequential Dump"

	// QoSPrincipalNameBackupTargetAuto captures enum value "Backup Target Auto"
	QoSPrincipalNameBackupTargetAuto string = "Backup Target Auto"
)

// prop value enum
func (m *QoS) validatePrincipalNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, qoSTypePrincipalNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QoS) validatePrincipalName(formats strfmt.Registry) error {
	if swag.IsZero(m.PrincipalName) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrincipalNameEnum("principalName", "body", *m.PrincipalName); err != nil {
		return err
	}

	return nil
}

var qoSTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BackupTargetHigh","BackupTargetLow","TestAndDevHigh","TestAndDevLow","BackupTargetSSD","BackupTargetCommvault","JournaledSequentialDump","BackupTargetAuto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		qoSTypeNamePropEnum = append(qoSTypeNamePropEnum, v)
	}
}

const (

	// QoSNameBackupTargetHigh captures enum value "BackupTargetHigh"
	QoSNameBackupTargetHigh string = "BackupTargetHigh"

	// QoSNameBackupTargetLow captures enum value "BackupTargetLow"
	QoSNameBackupTargetLow string = "BackupTargetLow"

	// QoSNameTestAndDevHigh captures enum value "TestAndDevHigh"
	QoSNameTestAndDevHigh string = "TestAndDevHigh"

	// QoSNameTestAndDevLow captures enum value "TestAndDevLow"
	QoSNameTestAndDevLow string = "TestAndDevLow"

	// QoSNameBackupTargetSSD captures enum value "BackupTargetSSD"
	QoSNameBackupTargetSSD string = "BackupTargetSSD"

	// QoSNameBackupTargetCommvault captures enum value "BackupTargetCommvault"
	QoSNameBackupTargetCommvault string = "BackupTargetCommvault"

	// QoSNameJournaledSequentialDump captures enum value "JournaledSequentialDump"
	QoSNameJournaledSequentialDump string = "JournaledSequentialDump"

	// QoSNameBackupTargetAuto captures enum value "BackupTargetAuto"
	QoSNameBackupTargetAuto string = "BackupTargetAuto"
)

// prop value enum
func (m *QoS) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, qoSTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QoS) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this qo s based on context it is used
func (m *QoS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QoS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QoS) UnmarshalBinary(b []byte) error {
	var res QoS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
