// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImportCertificateByClientcsrResponseBody Import Certificate Response
//
// Specifies the response to import a certificate.
//
// swagger:model ImportCertificateByClientcsrResponseBody
type ImportCertificateByClientcsrResponseBody struct {

	// Specifies the server certificate.
	CertificateServer *string `json:"certificateServer,omitempty"`

	// Specifies the private key of agent.
	PrivateKey *string `json:"privateKey,omitempty"`

	// Specifies the path to the file to be uploaded to server. This file has the server cert, id and encrypted private key
	FileServerCert *string `json:"fileServerCert,omitempty"`
}

// Validate validates this import certificate by clientcsr response body
func (m *ImportCertificateByClientcsrResponseBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this import certificate by clientcsr response body based on context it is used
func (m *ImportCertificateByClientcsrResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImportCertificateByClientcsrResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportCertificateByClientcsrResponseBody) UnmarshalBinary(b []byte) error {
	var res ImportCertificateByClientcsrResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
