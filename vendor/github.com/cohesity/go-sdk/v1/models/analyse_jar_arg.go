// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnalyseJarArg Analyse Jar File.
//
// API to analyse a JAR file. This JAR may contain multiple mappers/reducers.
// Jar will be analysed and list of all mappers/reducers found in the jar will
// be returned.
//
// swagger:model AnalyseJarArg
type AnalyseJarArg struct {

	// Name of the JAR to be analysed.
	JarName *string `json:"jarName,omitempty"`

	// Path of the jar file.
	JarPath *string `json:"jarPath,omitempty"`

	// jar relative path
	JarRelativePath *string `json:"jarRelativePath,omitempty"`

	// If this flag is true, then also save mapper and reducers in scribe.
	SaveEntities *bool `json:"saveEntities,omitempty"`
}

// Validate validates this analyse jar arg
func (m *AnalyseJarArg) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this analyse jar arg based on context it is used
func (m *AnalyseJarArg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalyseJarArg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyseJarArg) UnmarshalBinary(b []byte) error {
	var res AnalyseJarArg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
