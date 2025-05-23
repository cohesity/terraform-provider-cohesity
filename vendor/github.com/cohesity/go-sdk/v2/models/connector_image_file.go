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

// ConnectorImageFile Specifies the URL to access the connector image file and the platform on which the image can be deployed. The software version of the connector is assumed to be derivable from the name of the image file.
//
// swagger:model ConnectorImageFile
type ConnectorImageFile struct {

	// Specifies the URL to access the file.
	// Required: true
	URL *string `json:"url"`

	// Specifies the platform on which the image can be deployed.
	// Required: true
	// Enum: ["VSI","VMware"]
	ImageType *string `json:"imageType"`
}

// Validate validates this connector image file
func (m *ConnectorImageFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectorImageFile) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

var connectorImageFileTypeImageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VSI","VMware"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectorImageFileTypeImageTypePropEnum = append(connectorImageFileTypeImageTypePropEnum, v)
	}
}

const (

	// ConnectorImageFileImageTypeVSI captures enum value "VSI"
	ConnectorImageFileImageTypeVSI string = "VSI"

	// ConnectorImageFileImageTypeVMware captures enum value "VMware"
	ConnectorImageFileImageTypeVMware string = "VMware"
)

// prop value enum
func (m *ConnectorImageFile) validateImageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectorImageFileTypeImageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConnectorImageFile) validateImageType(formats strfmt.Registry) error {

	if err := validate.Required("imageType", "body", m.ImageType); err != nil {
		return err
	}

	// value enum
	if err := m.validateImageTypeEnum("imageType", "body", *m.ImageType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this connector image file based on context it is used
func (m *ConnectorImageFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectorImageFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectorImageFile) UnmarshalBinary(b []byte) error {
	var res ConnectorImageFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
