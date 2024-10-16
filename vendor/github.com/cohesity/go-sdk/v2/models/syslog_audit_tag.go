// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SyslogAuditTag SyslogAuditTag
//
// Cohesity audit tag name.
//
// swagger:model SyslogAuditTag
type SyslogAuditTag struct {

	// Cluster audit tagging name.
	ClusterAudit *string `json:"clusterAudit,omitempty"`

	// Filer audit tagging name.
	FilerAudit *string `json:"filerAudit,omitempty"`

	// Data protection events audit tagging name.
	DataProtectionEventsAudit *string `json:"dataProtectionEventsAudit,omitempty"`

	// Alert audit tagging name.
	AlertAudit *string `json:"alertAudit,omitempty"`
}

// Validate validates this syslog audit tag
func (m *SyslogAuditTag) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this syslog audit tag based on context it is used
func (m *SyslogAuditTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SyslogAuditTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyslogAuditTag) UnmarshalBinary(b []byte) error {
	var res SyslogAuditTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
