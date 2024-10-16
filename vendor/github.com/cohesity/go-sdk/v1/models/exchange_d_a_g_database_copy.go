// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExchangeDAGDatabaseCopy Information about a DAG Database copy status on a member Exchange server in
// DAG. Note that this information may become stale as it contains status
// information. If a DB is created in DAG node, it becomes active copy on that
// node with single copy. This information is assembled from
// Get-MailboxDatabase and Get-MailboxDatabaseCopyStatus PowerShell cmdlets.
//
// swagger:model ExchangeDAGDatabaseCopy
type ExchangeDAGDatabaseCopy struct {

	// Activation preference number assigned for this database on a node.
	// This is 1 based number assigned in Exchange Admin Console.
	ActivationPreference *uint32 `json:"activationPreference,omitempty"`

	// Active copy of the database. Look for active_copy=true and
	// status=kHealthy to find active mounted copy of a database.
	ActiveCopy *bool `json:"activeCopy,omitempty"`

	// Database copy error details.
	CopyError *DatabaseCopyError `json:"copyError,omitempty"`

	// Guid of the DAG db copy. This is unique to a copy of the database on a
	// specific node. Use the ExchangeDatabase::guid to identify the parent
	// database for which this copy exists.
	CopyGUID *string `json:"copyGuid,omitempty"`

	// Database copy status. If the status is in error state, the error info
	// will be present in error object.
	CopyStatus *int32 `json:"copyStatus,omitempty"`

	// Parent database information.
	Db *ExchangeDatabase `json:"db,omitempty"`

	// Outstanding logs to copy for the database copy on this node.
	// 1 = Unknown. This is the 'CopyQueueLength' cmdlet param.
	MaxLogToCopy *int64 `json:"maxLogToCopy,omitempty"`

	// Outstanding logs to replay for the database copy on this node.
	// 1 = Unknown. This is the 'ReplayQueueLength'.
	MaxLogToReplay *int64 `json:"maxLogToReplay,omitempty"`
}

// Validate validates this exchange d a g database copy
func (m *ExchangeDAGDatabaseCopy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCopyError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeDAGDatabaseCopy) validateCopyError(formats strfmt.Registry) error {
	if swag.IsZero(m.CopyError) { // not required
		return nil
	}

	if m.CopyError != nil {
		if err := m.CopyError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("copyError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("copyError")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeDAGDatabaseCopy) validateDb(formats strfmt.Registry) error {
	if swag.IsZero(m.Db) { // not required
		return nil
	}

	if m.Db != nil {
		if err := m.Db.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this exchange d a g database copy based on the context it is used
func (m *ExchangeDAGDatabaseCopy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCopyError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeDAGDatabaseCopy) contextValidateCopyError(ctx context.Context, formats strfmt.Registry) error {

	if m.CopyError != nil {

		if swag.IsZero(m.CopyError) { // not required
			return nil
		}

		if err := m.CopyError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("copyError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("copyError")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeDAGDatabaseCopy) contextValidateDb(ctx context.Context, formats strfmt.Registry) error {

	if m.Db != nil {

		if swag.IsZero(m.Db) { // not required
			return nil
		}

		if err := m.Db.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeDAGDatabaseCopy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeDAGDatabaseCopy) UnmarshalBinary(b []byte) error {
	var res ExchangeDAGDatabaseCopy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
