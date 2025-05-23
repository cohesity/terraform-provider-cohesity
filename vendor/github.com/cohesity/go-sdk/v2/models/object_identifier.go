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

// ObjectIdentifier Object Identifier.
//
// Specifies the basic info to identify an object.
//
// swagger:model ObjectIdentifier
type ObjectIdentifier struct {

	// Specifies object id.
	ID *int64 `json:"id,omitempty"`

	// Specifies the name of the object.
	Name *string `json:"name,omitempty"`

	// Specifies registered source id to which object belongs.
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies registered source name to which object belongs.
	SourceName *string `json:"sourceName,omitempty"`

	// Specifies the environment of the object.
	// Enum: ["kVMware","kHyperV","kAzure","kKVM","kAWS","kAcropolis","kGCP","kPhysical","kPhysicalFiles","kIsilon","kNetapp","kGenericNas","kFlashBlade","kElastifile","kGPFS","kPure","kIbmFlashSystem","kNimble","kSQL","kOracle","kExchange","kAD","kView","kO365","kHyperFlex","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kSAPHANA","kUDA","kSfdc","kExperimentalAdapter","kAzureEntraID","kMongoDBPhysical"]
	Environment *string `json:"environment,omitempty"`

	// Specifies the string based Id for an object and also provides the history of ids assigned to the object
	EntityID *ObjectStringIdentifier `json:"entityId,omitempty"`
}

// Validate validates this object identifier
func (m *ObjectIdentifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var objectIdentifierTypeEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kAzure","kKVM","kAWS","kAcropolis","kGCP","kPhysical","kPhysicalFiles","kIsilon","kNetapp","kGenericNas","kFlashBlade","kElastifile","kGPFS","kPure","kIbmFlashSystem","kNimble","kSQL","kOracle","kExchange","kAD","kView","kO365","kHyperFlex","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kSAPHANA","kUDA","kSfdc","kExperimentalAdapter","kAzureEntraID","kMongoDBPhysical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectIdentifierTypeEnvironmentPropEnum = append(objectIdentifierTypeEnvironmentPropEnum, v)
	}
}

const (

	// ObjectIdentifierEnvironmentKVMware captures enum value "kVMware"
	ObjectIdentifierEnvironmentKVMware string = "kVMware"

	// ObjectIdentifierEnvironmentKHyperV captures enum value "kHyperV"
	ObjectIdentifierEnvironmentKHyperV string = "kHyperV"

	// ObjectIdentifierEnvironmentKAzure captures enum value "kAzure"
	ObjectIdentifierEnvironmentKAzure string = "kAzure"

	// ObjectIdentifierEnvironmentKKVM captures enum value "kKVM"
	ObjectIdentifierEnvironmentKKVM string = "kKVM"

	// ObjectIdentifierEnvironmentKAWS captures enum value "kAWS"
	ObjectIdentifierEnvironmentKAWS string = "kAWS"

	// ObjectIdentifierEnvironmentKAcropolis captures enum value "kAcropolis"
	ObjectIdentifierEnvironmentKAcropolis string = "kAcropolis"

	// ObjectIdentifierEnvironmentKGCP captures enum value "kGCP"
	ObjectIdentifierEnvironmentKGCP string = "kGCP"

	// ObjectIdentifierEnvironmentKPhysical captures enum value "kPhysical"
	ObjectIdentifierEnvironmentKPhysical string = "kPhysical"

	// ObjectIdentifierEnvironmentKPhysicalFiles captures enum value "kPhysicalFiles"
	ObjectIdentifierEnvironmentKPhysicalFiles string = "kPhysicalFiles"

	// ObjectIdentifierEnvironmentKIsilon captures enum value "kIsilon"
	ObjectIdentifierEnvironmentKIsilon string = "kIsilon"

	// ObjectIdentifierEnvironmentKNetapp captures enum value "kNetapp"
	ObjectIdentifierEnvironmentKNetapp string = "kNetapp"

	// ObjectIdentifierEnvironmentKGenericNas captures enum value "kGenericNas"
	ObjectIdentifierEnvironmentKGenericNas string = "kGenericNas"

	// ObjectIdentifierEnvironmentKFlashBlade captures enum value "kFlashBlade"
	ObjectIdentifierEnvironmentKFlashBlade string = "kFlashBlade"

	// ObjectIdentifierEnvironmentKElastifile captures enum value "kElastifile"
	ObjectIdentifierEnvironmentKElastifile string = "kElastifile"

	// ObjectIdentifierEnvironmentKGPFS captures enum value "kGPFS"
	ObjectIdentifierEnvironmentKGPFS string = "kGPFS"

	// ObjectIdentifierEnvironmentKPure captures enum value "kPure"
	ObjectIdentifierEnvironmentKPure string = "kPure"

	// ObjectIdentifierEnvironmentKIbmFlashSystem captures enum value "kIbmFlashSystem"
	ObjectIdentifierEnvironmentKIbmFlashSystem string = "kIbmFlashSystem"

	// ObjectIdentifierEnvironmentKNimble captures enum value "kNimble"
	ObjectIdentifierEnvironmentKNimble string = "kNimble"

	// ObjectIdentifierEnvironmentKSQL captures enum value "kSQL"
	ObjectIdentifierEnvironmentKSQL string = "kSQL"

	// ObjectIdentifierEnvironmentKOracle captures enum value "kOracle"
	ObjectIdentifierEnvironmentKOracle string = "kOracle"

	// ObjectIdentifierEnvironmentKExchange captures enum value "kExchange"
	ObjectIdentifierEnvironmentKExchange string = "kExchange"

	// ObjectIdentifierEnvironmentKAD captures enum value "kAD"
	ObjectIdentifierEnvironmentKAD string = "kAD"

	// ObjectIdentifierEnvironmentKView captures enum value "kView"
	ObjectIdentifierEnvironmentKView string = "kView"

	// ObjectIdentifierEnvironmentKO365 captures enum value "kO365"
	ObjectIdentifierEnvironmentKO365 string = "kO365"

	// ObjectIdentifierEnvironmentKHyperFlex captures enum value "kHyperFlex"
	ObjectIdentifierEnvironmentKHyperFlex string = "kHyperFlex"

	// ObjectIdentifierEnvironmentKKubernetes captures enum value "kKubernetes"
	ObjectIdentifierEnvironmentKKubernetes string = "kKubernetes"

	// ObjectIdentifierEnvironmentKCassandra captures enum value "kCassandra"
	ObjectIdentifierEnvironmentKCassandra string = "kCassandra"

	// ObjectIdentifierEnvironmentKMongoDB captures enum value "kMongoDB"
	ObjectIdentifierEnvironmentKMongoDB string = "kMongoDB"

	// ObjectIdentifierEnvironmentKCouchbase captures enum value "kCouchbase"
	ObjectIdentifierEnvironmentKCouchbase string = "kCouchbase"

	// ObjectIdentifierEnvironmentKHdfs captures enum value "kHdfs"
	ObjectIdentifierEnvironmentKHdfs string = "kHdfs"

	// ObjectIdentifierEnvironmentKHive captures enum value "kHive"
	ObjectIdentifierEnvironmentKHive string = "kHive"

	// ObjectIdentifierEnvironmentKHBase captures enum value "kHBase"
	ObjectIdentifierEnvironmentKHBase string = "kHBase"

	// ObjectIdentifierEnvironmentKSAPHANA captures enum value "kSAPHANA"
	ObjectIdentifierEnvironmentKSAPHANA string = "kSAPHANA"

	// ObjectIdentifierEnvironmentKUDA captures enum value "kUDA"
	ObjectIdentifierEnvironmentKUDA string = "kUDA"

	// ObjectIdentifierEnvironmentKSfdc captures enum value "kSfdc"
	ObjectIdentifierEnvironmentKSfdc string = "kSfdc"

	// ObjectIdentifierEnvironmentKExperimentalAdapter captures enum value "kExperimentalAdapter"
	ObjectIdentifierEnvironmentKExperimentalAdapter string = "kExperimentalAdapter"

	// ObjectIdentifierEnvironmentKAzureEntraID captures enum value "kAzureEntraID"
	ObjectIdentifierEnvironmentKAzureEntraID string = "kAzureEntraID"

	// ObjectIdentifierEnvironmentKMongoDBPhysical captures enum value "kMongoDBPhysical"
	ObjectIdentifierEnvironmentKMongoDBPhysical string = "kMongoDBPhysical"
)

// prop value enum
func (m *ObjectIdentifier) validateEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectIdentifierTypeEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectIdentifier) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnvironmentEnum("environment", "body", *m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *ObjectIdentifier) validateEntityID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityID) { // not required
		return nil
	}

	if m.EntityID != nil {
		if err := m.EntityID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object identifier based on the context it is used
func (m *ObjectIdentifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectIdentifier) contextValidateEntityID(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityID != nil {

		if swag.IsZero(m.EntityID) { // not required
			return nil
		}

		if err := m.EntityID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectIdentifier) UnmarshalBinary(b []byte) error {
	var res ObjectIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
