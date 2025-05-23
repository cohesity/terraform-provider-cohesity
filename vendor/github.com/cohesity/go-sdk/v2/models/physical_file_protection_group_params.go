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

// PhysicalFileProtectionGroupParams Specifies the parameters which are specific to Physical related Protection Groups.
//
// swagger:model PhysicalFileProtectionGroupParams
type PhysicalFileProtectionGroupParams struct {

	// Specifies the list of objects protected by this Protection Group.
	// Required: true
	// Min Items: 1
	Objects []*PhysicalFileProtectionGroupObjectParams `json:"objects"`

	// Specifies the fields required to enable indexing of the protected objects such as files and directories.
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`

	// Specifies whether or not to perform source side deduplication on this Protection Group.
	PerformSourceSideDeduplication *bool `json:"performSourceSideDeduplication,omitempty"`

	// Specifies whether or not to perform brick based deduplication on this Protection Group.
	PerformBrickBasedDeduplication *bool `json:"performBrickBasedDeduplication,omitempty"`

	// Specifies the timeouts for all the objects inside this Protection Group, for both full and incremental backups.
	TaskTimeouts []*CancellationTimeoutParams `json:"taskTimeouts"`

	// Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup.
	Quiesce *bool `json:"quiesce,omitempty"`

	// Specifies whether to continue backing up on quiesce failure.
	ContinueOnQuiesceFailure *bool `json:"continueOnQuiesceFailure,omitempty"`

	// Specifies whether to take CoBMR backup.
	CobmrBackup *bool `json:"cobmrBackup,omitempty"`

	// Specifies the pre and post script parameters associated with a protection group.
	PrePostScript *PrePostScriptParams `json:"prePostScript,omitempty"`

	// Specifies ids of sources for which deduplication has to be disabled.
	DedupExclusionSourceIds []int64 `json:"dedupExclusionSourceIds"`

	// Specifies global exclude filters which are applied to all sources in a job.
	GlobalExcludePaths []string `json:"globalExcludePaths"`

	// Specifies global exclude filesystems which are applied to all sources in a job.
	GlobalExcludeFS []string `json:"globalExcludeFS"`

	// Specifies the Errors to be ignored in error db.
	IgnorableErrors []string `json:"ignorableErrors"`

	// Specifies whether or not this job can have parallel runs.
	AllowParallelRuns *bool `json:"allowParallelRuns,omitempty"`

	// Specifies writer names which should be excluded from physical file based backups.
	ExcludedVssWriters []string `json:"excludedVssWriters"`
}

// Validate validates this physical file protection group params
func (m *PhysicalFileProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskTimeouts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrePostScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgnorableErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalFileProtectionGroupParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
		return err
	}

	iObjectsSize := int64(len(m.Objects))

	if err := validate.MinItems("objects", "body", iObjectsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PhysicalFileProtectionGroupParams) validateIndexingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexingPolicy) { // not required
		return nil
	}

	if m.IndexingPolicy != nil {
		if err := m.IndexingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("indexingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("indexingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *PhysicalFileProtectionGroupParams) validateTaskTimeouts(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskTimeouts) { // not required
		return nil
	}

	for i := 0; i < len(m.TaskTimeouts); i++ {
		if swag.IsZero(m.TaskTimeouts[i]) { // not required
			continue
		}

		if m.TaskTimeouts[i] != nil {
			if err := m.TaskTimeouts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taskTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taskTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PhysicalFileProtectionGroupParams) validatePrePostScript(formats strfmt.Registry) error {
	if swag.IsZero(m.PrePostScript) { // not required
		return nil
	}

	if m.PrePostScript != nil {
		if err := m.PrePostScript.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prePostScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prePostScript")
			}
			return err
		}
	}

	return nil
}

var physicalFileProtectionGroupParamsIgnorableErrorsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kEOF","kNonExistent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		physicalFileProtectionGroupParamsIgnorableErrorsItemsEnum = append(physicalFileProtectionGroupParamsIgnorableErrorsItemsEnum, v)
	}
}

func (m *PhysicalFileProtectionGroupParams) validateIgnorableErrorsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, physicalFileProtectionGroupParamsIgnorableErrorsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PhysicalFileProtectionGroupParams) validateIgnorableErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.IgnorableErrors) { // not required
		return nil
	}

	for i := 0; i < len(m.IgnorableErrors); i++ {

		// value enum
		if err := m.validateIgnorableErrorsItemsEnum("ignorableErrors"+"."+strconv.Itoa(i), "body", m.IgnorableErrors[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this physical file protection group params based on the context it is used
func (m *PhysicalFileProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndexingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaskTimeouts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrePostScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalFileProtectionGroupParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PhysicalFileProtectionGroupParams) contextValidateIndexingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IndexingPolicy != nil {

		if swag.IsZero(m.IndexingPolicy) { // not required
			return nil
		}

		if err := m.IndexingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("indexingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("indexingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *PhysicalFileProtectionGroupParams) contextValidateTaskTimeouts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaskTimeouts); i++ {

		if m.TaskTimeouts[i] != nil {

			if swag.IsZero(m.TaskTimeouts[i]) { // not required
				return nil
			}

			if err := m.TaskTimeouts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taskTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taskTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PhysicalFileProtectionGroupParams) contextValidatePrePostScript(ctx context.Context, formats strfmt.Registry) error {

	if m.PrePostScript != nil {

		if swag.IsZero(m.PrePostScript) { // not required
			return nil
		}

		if err := m.PrePostScript.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prePostScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prePostScript")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PhysicalFileProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhysicalFileProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res PhysicalFileProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
