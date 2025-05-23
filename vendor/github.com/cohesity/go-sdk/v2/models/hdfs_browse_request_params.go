// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HdfsBrowseRequestParams hdfs browse request params
//
// swagger:model HdfsBrowseRequestParams
type HdfsBrowseRequestParams struct {

	// Specifies the path whose contents are to be returned. The last token in the path can be a regex. In this case the regex is applied on the contents of the path upto the second-last token and the matching contents are returned.
	ParentPath *string `json:"parentPath,omitempty"`
}

// Validate validates this hdfs browse request params
func (m *HdfsBrowseRequestParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hdfs browse request params based on context it is used
func (m *HdfsBrowseRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HdfsBrowseRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HdfsBrowseRequestParams) UnmarshalBinary(b []byte) error {
	var res HdfsBrowseRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
