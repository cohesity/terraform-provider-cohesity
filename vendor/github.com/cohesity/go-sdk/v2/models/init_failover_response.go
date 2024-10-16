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

// InitFailoverResponse Specifies the response after succesfully initiating the failover request.
//
// swagger:model InitFailoverResponse
type InitFailoverResponse struct {

	// Specifies the list of corrosponding source objects mapped with replica objects provided at the time of initiating failover request.
	ReplicaToSourceObjects []*SourceReplicaObject `json:"replicaToSourceObjects"`

	// Specifies the information about source cluster in failover workflow.
	SourceClusterInfo *FailoverSourceCluster `json:"sourceClusterInfo,omitempty"`
}

// Validate validates this init failover response
func (m *InitFailoverResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplicaToSourceObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceClusterInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitFailoverResponse) validateReplicaToSourceObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicaToSourceObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.ReplicaToSourceObjects); i++ {
		if swag.IsZero(m.ReplicaToSourceObjects[i]) { // not required
			continue
		}

		if m.ReplicaToSourceObjects[i] != nil {
			if err := m.ReplicaToSourceObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaToSourceObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaToSourceObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InitFailoverResponse) validateSourceClusterInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceClusterInfo) { // not required
		return nil
	}

	if m.SourceClusterInfo != nil {
		if err := m.SourceClusterInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceClusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceClusterInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this init failover response based on the context it is used
func (m *InitFailoverResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReplicaToSourceObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceClusterInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitFailoverResponse) contextValidateReplicaToSourceObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReplicaToSourceObjects); i++ {

		if m.ReplicaToSourceObjects[i] != nil {

			if swag.IsZero(m.ReplicaToSourceObjects[i]) { // not required
				return nil
			}

			if err := m.ReplicaToSourceObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaToSourceObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaToSourceObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InitFailoverResponse) contextValidateSourceClusterInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SourceClusterInfo != nil {

		if swag.IsZero(m.SourceClusterInfo) { // not required
			return nil
		}

		if err := m.SourceClusterInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceClusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sourceClusterInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitFailoverResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitFailoverResponse) UnmarshalBinary(b []byte) error {
	var res InitFailoverResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
