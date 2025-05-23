// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContinuousSnapshotParams Continuous Snapshot Params
//
// Specifies the source snapshots to be taken even if there is a pending run in a protection group.
//
// swagger:model ContinuousSnapshotParams
type ContinuousSnapshotParams struct {

	// Specifies whether source snapshots should be taken even if there is a pending run.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`

	// Specifies the maximum number of source snapshots allowed for a given object in a protection group. This is only applicable if isContinuousSnapshottingEnabled is set to true.
	MaxAllowedSnapshots *int32 `json:"maxAllowedSnapshots,omitempty"`
}

// Validate validates this continuous snapshot params
func (m *ContinuousSnapshotParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContinuousSnapshotParams) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this continuous snapshot params based on context it is used
func (m *ContinuousSnapshotParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContinuousSnapshotParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContinuousSnapshotParams) UnmarshalBinary(b []byte) error {
	var res ContinuousSnapshotParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
