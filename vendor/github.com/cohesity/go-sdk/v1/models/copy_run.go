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

// CopyRun Copy Run Task.
//
// Specifies details about the Copy Run for a backup run of a Job Run.
// A Copy task copies snapshots resulted from a backup run to a snapshot
// target which could be 'kLocal', 'kArchival', or 'kRemote'.
//
// swagger:model CopyRun
type CopyRun struct {

	// Specifies the status information of each task that copies the snapshot
	// taken for a Protection Source.
	CopySnapshotTasks []*CopySnapshotTaskStatus `json:"copySnapshotTasks"`

	// Specifies the datalock constraints for a copy run task.
	DataLockConstraints *DataLockConstraints `json:"dataLockConstraints,omitempty"`

	// Specifies if an error occurred (if any) while running this task.
	// This field is populated when the status is equal to 'kFailure'.
	Error *string `json:"error,omitempty"`

	// Specifies expiry time of the copies of the snapshots in this Protection
	// Run.
	ExpiryTimeUsecs *int64 `json:"expiryTimeUsecs,omitempty"`

	// Specifies whether legal hold is enabled on this run. It is true if the
	// run is put on legal hold. Independent of this flag, some of the entities
	// may be on legal hold.
	HoldForLegalPurpose *bool `json:"holdForLegalPurpose,omitempty"`

	// Specifies the list of Protection Source Ids and the legal hold status.
	LegalHoldings []*LegalHoldings `json:"legalHoldings"`

	// Specifies start time of the copy run.
	RunStartTimeUsecs *int64 `json:"runStartTimeUsecs,omitempty"`

	// Specifies the aggregated information of all the copy tasks.
	Stats *CopyRunStats `json:"stats,omitempty"`

	// Specifies the aggregated status of copy tasks such as 'kRunning',
	// 'kSuccess', 'kFailure' etc.
	// 'kAccepted' indicates the task is queued to run but not yet running.
	// 'kRunning' indicates the task is running.
	// 'kCanceling' indicates a request to cancel the task has occurred but
	//
	// the task is not yet canceled.
	//
	// 'kCanceled' indicates the task has been canceled.
	// 'kSuccess' indicates the task was successful.
	// 'kFailure' indicates the task failed.
	// 'kWarning' indicates the task has finished with warning.
	// 'kOnHold' indicates the task is kept onHold.
	// 'kMissed' indicates the task is missed.
	// 'Finalizing' indicates the task is finalizing.
	// Enum: ["kAccepted","kRunning","kCanceling","kCanceled","kSuccess","kFailure","kWarning","kOnHold","kMissed","kFinalizing"]
	Status *string `json:"status,omitempty"`

	// Specifies the target of the copy task such as an external target or a
	// Remote Cohesity Cluster.
	Target *SnapshotTargetSettings `json:"target,omitempty"`

	// Specifies a globally unique id of the copy task.
	TaskUID struct {
		UniversalID
	} `json:"taskUid,omitempty"`

	// Specifies a message to the user if any manual intervention is needed to
	// make forward progress for the archival task. This message is mainly
	// relevant for tape based archival tasks where a backup admin might be
	// asked to load a new media when the tape library does not have any more
	// media to use.
	UserActionMessage *string `json:"userActionMessage,omitempty"`
}

// Validate validates this copy run
func (m *CopyRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCopySnapshotTasks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataLockConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLegalHoldings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyRun) validateCopySnapshotTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.CopySnapshotTasks) { // not required
		return nil
	}

	for i := 0; i < len(m.CopySnapshotTasks); i++ {
		if swag.IsZero(m.CopySnapshotTasks[i]) { // not required
			continue
		}

		if m.CopySnapshotTasks[i] != nil {
			if err := m.CopySnapshotTasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copySnapshotTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copySnapshotTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyRun) validateDataLockConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.DataLockConstraints) { // not required
		return nil
	}

	if m.DataLockConstraints != nil {
		if err := m.DataLockConstraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataLockConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataLockConstraints")
			}
			return err
		}
	}

	return nil
}

func (m *CopyRun) validateLegalHoldings(formats strfmt.Registry) error {
	if swag.IsZero(m.LegalHoldings) { // not required
		return nil
	}

	for i := 0; i < len(m.LegalHoldings); i++ {
		if swag.IsZero(m.LegalHoldings[i]) { // not required
			continue
		}

		if m.LegalHoldings[i] != nil {
			if err := m.LegalHoldings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("legalHoldings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("legalHoldings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyRun) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

var copyRunTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAccepted","kRunning","kCanceling","kCanceled","kSuccess","kFailure","kWarning","kOnHold","kMissed","kFinalizing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		copyRunTypeStatusPropEnum = append(copyRunTypeStatusPropEnum, v)
	}
}

const (

	// CopyRunStatusKAccepted captures enum value "kAccepted"
	CopyRunStatusKAccepted string = "kAccepted"

	// CopyRunStatusKRunning captures enum value "kRunning"
	CopyRunStatusKRunning string = "kRunning"

	// CopyRunStatusKCanceling captures enum value "kCanceling"
	CopyRunStatusKCanceling string = "kCanceling"

	// CopyRunStatusKCanceled captures enum value "kCanceled"
	CopyRunStatusKCanceled string = "kCanceled"

	// CopyRunStatusKSuccess captures enum value "kSuccess"
	CopyRunStatusKSuccess string = "kSuccess"

	// CopyRunStatusKFailure captures enum value "kFailure"
	CopyRunStatusKFailure string = "kFailure"

	// CopyRunStatusKWarning captures enum value "kWarning"
	CopyRunStatusKWarning string = "kWarning"

	// CopyRunStatusKOnHold captures enum value "kOnHold"
	CopyRunStatusKOnHold string = "kOnHold"

	// CopyRunStatusKMissed captures enum value "kMissed"
	CopyRunStatusKMissed string = "kMissed"

	// CopyRunStatusKFinalizing captures enum value "kFinalizing"
	CopyRunStatusKFinalizing string = "kFinalizing"
)

// prop value enum
func (m *CopyRun) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, copyRunTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CopyRun) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CopyRun) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CopyRun) validateTaskUID(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskUID) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this copy run based on the context it is used
func (m *CopyRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCopySnapshotTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataLockConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLegalHoldings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaskUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyRun) contextValidateCopySnapshotTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CopySnapshotTasks); i++ {

		if m.CopySnapshotTasks[i] != nil {

			if swag.IsZero(m.CopySnapshotTasks[i]) { // not required
				return nil
			}

			if err := m.CopySnapshotTasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copySnapshotTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copySnapshotTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyRun) contextValidateDataLockConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.DataLockConstraints != nil {

		if swag.IsZero(m.DataLockConstraints) { // not required
			return nil
		}

		if err := m.DataLockConstraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataLockConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataLockConstraints")
			}
			return err
		}
	}

	return nil
}

func (m *CopyRun) contextValidateLegalHoldings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LegalHoldings); i++ {

		if m.LegalHoldings[i] != nil {

			if swag.IsZero(m.LegalHoldings[i]) { // not required
				return nil
			}

			if err := m.LegalHoldings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("legalHoldings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("legalHoldings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyRun) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if swag.IsZero(m.Stats) { // not required
			return nil
		}

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *CopyRun) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CopyRun) contextValidateTaskUID(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *CopyRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyRun) UnmarshalBinary(b []byte) error {
	var res CopyRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
