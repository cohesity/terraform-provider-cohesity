// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// O365SearchEmailsRequestParams O365 Search Emails Request Params
//
// Specifies email search request params specific to O365 environment.
//
// swagger:model O365SearchEmailsRequestParams
type O365SearchEmailsRequestParams struct {

	// Specifies the domain Ids in which mailboxes are registered.
	DomainIds []int64 `json:"domainIds"`

	// Specifies the mailbox Ids which contains the emails/folders.
	MailboxIds []int64 `json:"mailboxIds"`
}

// Validate validates this o365 search emails request params
func (m *O365SearchEmailsRequestParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o365 search emails request params based on context it is used
func (m *O365SearchEmailsRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *O365SearchEmailsRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *O365SearchEmailsRequestParams) UnmarshalBinary(b []byte) error {
	var res O365SearchEmailsRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
