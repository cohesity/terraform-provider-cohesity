// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetTrackingViewIDResponse Specifies the response upon requesting a tracking view id
//
// swagger:model GetTrackingViewIdResponse
type GetTrackingViewIDResponse struct {

	// Specifies the local view id corresponding to the view uid.
	TrackingViewID *int64 `json:"trackingViewId,omitempty"`
}

// Validate validates this get tracking view Id response
func (m *GetTrackingViewIDResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get tracking view Id response based on context it is used
func (m *GetTrackingViewIDResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetTrackingViewIDResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTrackingViewIDResponse) UnmarshalBinary(b []byte) error {
	var res GetTrackingViewIDResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
