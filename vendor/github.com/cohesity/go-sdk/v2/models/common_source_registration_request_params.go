// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommonSourceRegistrationRequestParams Specifies the parameters which are common between all Protection Source registrations.
//
// swagger:model CommonSourceRegistrationRequestParams
type CommonSourceRegistrationRequestParams struct {

	// Specifies the environment type of the Protection Source.
	// Required: true
	// Enum: ["kVMware","kHyperV","kAcropolis","kKVM","kAWS","kGCP","kAzure","kPhysical","kPure","kIbmFlashSystem","kNimble","kNetapp","kGenericNas","kIsilon","kFlashBlade","kGPFS","kElastifile","kO365","kHyperFlex","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kSAPHANA","kUDA","kSQL","kOracle","kSfdc","kExperimentalAdapter","kMongoDBPhysical"]
	Environment *string `json:"environment"`

	// A user specified name for this source.
	Name *string `json:"name,omitempty"`

	// Specifies if credentials are encrypted by internal key.
	IsInternalEncrypted *bool `json:"isInternalEncrypted,omitempty"`

	// Specifies the key that user has encrypted the credential with.
	EncryptionKey *string `json:"encryptionKey,omitempty"`

	// Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user.
	ConnectionID *int64 `json:"connectionId,omitempty"`

	// Specfies the list of connections for the source.
	Connections []*ConnectionConfig `json:"connections"`

	// Specifies the connector group id of connector groups.
	ConnectorGroupID *int64 `json:"connectorGroupId,omitempty"`

	// Specifies the advanced configuration for a protection source.
	AdvancedConfigs []*KeyValuePair `json:"advancedConfigs"`

	// Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. Also, this is the 'string' of connectionId. This property was added to accommodate for ID values that exceed 2^53 - 1, which is the max value for which JS maintains precision.
	DataSourceConnectionID *string `json:"dataSourceConnectionId,omitempty"`
}

// Validate validates this common source registration request params
func (m *CommonSourceRegistrationRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvancedConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var commonSourceRegistrationRequestParamsTypeEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kAcropolis","kKVM","kAWS","kGCP","kAzure","kPhysical","kPure","kIbmFlashSystem","kNimble","kNetapp","kGenericNas","kIsilon","kFlashBlade","kGPFS","kElastifile","kO365","kHyperFlex","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kSAPHANA","kUDA","kSQL","kOracle","kSfdc","kExperimentalAdapter","kMongoDBPhysical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonSourceRegistrationRequestParamsTypeEnvironmentPropEnum = append(commonSourceRegistrationRequestParamsTypeEnvironmentPropEnum, v)
	}
}

const (

	// CommonSourceRegistrationRequestParamsEnvironmentKVMware captures enum value "kVMware"
	CommonSourceRegistrationRequestParamsEnvironmentKVMware string = "kVMware"

	// CommonSourceRegistrationRequestParamsEnvironmentKHyperV captures enum value "kHyperV"
	CommonSourceRegistrationRequestParamsEnvironmentKHyperV string = "kHyperV"

	// CommonSourceRegistrationRequestParamsEnvironmentKAcropolis captures enum value "kAcropolis"
	CommonSourceRegistrationRequestParamsEnvironmentKAcropolis string = "kAcropolis"

	// CommonSourceRegistrationRequestParamsEnvironmentKKVM captures enum value "kKVM"
	CommonSourceRegistrationRequestParamsEnvironmentKKVM string = "kKVM"

	// CommonSourceRegistrationRequestParamsEnvironmentKAWS captures enum value "kAWS"
	CommonSourceRegistrationRequestParamsEnvironmentKAWS string = "kAWS"

	// CommonSourceRegistrationRequestParamsEnvironmentKGCP captures enum value "kGCP"
	CommonSourceRegistrationRequestParamsEnvironmentKGCP string = "kGCP"

	// CommonSourceRegistrationRequestParamsEnvironmentKAzure captures enum value "kAzure"
	CommonSourceRegistrationRequestParamsEnvironmentKAzure string = "kAzure"

	// CommonSourceRegistrationRequestParamsEnvironmentKPhysical captures enum value "kPhysical"
	CommonSourceRegistrationRequestParamsEnvironmentKPhysical string = "kPhysical"

	// CommonSourceRegistrationRequestParamsEnvironmentKPure captures enum value "kPure"
	CommonSourceRegistrationRequestParamsEnvironmentKPure string = "kPure"

	// CommonSourceRegistrationRequestParamsEnvironmentKIbmFlashSystem captures enum value "kIbmFlashSystem"
	CommonSourceRegistrationRequestParamsEnvironmentKIbmFlashSystem string = "kIbmFlashSystem"

	// CommonSourceRegistrationRequestParamsEnvironmentKNimble captures enum value "kNimble"
	CommonSourceRegistrationRequestParamsEnvironmentKNimble string = "kNimble"

	// CommonSourceRegistrationRequestParamsEnvironmentKNetapp captures enum value "kNetapp"
	CommonSourceRegistrationRequestParamsEnvironmentKNetapp string = "kNetapp"

	// CommonSourceRegistrationRequestParamsEnvironmentKGenericNas captures enum value "kGenericNas"
	CommonSourceRegistrationRequestParamsEnvironmentKGenericNas string = "kGenericNas"

	// CommonSourceRegistrationRequestParamsEnvironmentKIsilon captures enum value "kIsilon"
	CommonSourceRegistrationRequestParamsEnvironmentKIsilon string = "kIsilon"

	// CommonSourceRegistrationRequestParamsEnvironmentKFlashBlade captures enum value "kFlashBlade"
	CommonSourceRegistrationRequestParamsEnvironmentKFlashBlade string = "kFlashBlade"

	// CommonSourceRegistrationRequestParamsEnvironmentKGPFS captures enum value "kGPFS"
	CommonSourceRegistrationRequestParamsEnvironmentKGPFS string = "kGPFS"

	// CommonSourceRegistrationRequestParamsEnvironmentKElastifile captures enum value "kElastifile"
	CommonSourceRegistrationRequestParamsEnvironmentKElastifile string = "kElastifile"

	// CommonSourceRegistrationRequestParamsEnvironmentKO365 captures enum value "kO365"
	CommonSourceRegistrationRequestParamsEnvironmentKO365 string = "kO365"

	// CommonSourceRegistrationRequestParamsEnvironmentKHyperFlex captures enum value "kHyperFlex"
	CommonSourceRegistrationRequestParamsEnvironmentKHyperFlex string = "kHyperFlex"

	// CommonSourceRegistrationRequestParamsEnvironmentKKubernetes captures enum value "kKubernetes"
	CommonSourceRegistrationRequestParamsEnvironmentKKubernetes string = "kKubernetes"

	// CommonSourceRegistrationRequestParamsEnvironmentKCassandra captures enum value "kCassandra"
	CommonSourceRegistrationRequestParamsEnvironmentKCassandra string = "kCassandra"

	// CommonSourceRegistrationRequestParamsEnvironmentKMongoDB captures enum value "kMongoDB"
	CommonSourceRegistrationRequestParamsEnvironmentKMongoDB string = "kMongoDB"

	// CommonSourceRegistrationRequestParamsEnvironmentKCouchbase captures enum value "kCouchbase"
	CommonSourceRegistrationRequestParamsEnvironmentKCouchbase string = "kCouchbase"

	// CommonSourceRegistrationRequestParamsEnvironmentKHdfs captures enum value "kHdfs"
	CommonSourceRegistrationRequestParamsEnvironmentKHdfs string = "kHdfs"

	// CommonSourceRegistrationRequestParamsEnvironmentKHive captures enum value "kHive"
	CommonSourceRegistrationRequestParamsEnvironmentKHive string = "kHive"

	// CommonSourceRegistrationRequestParamsEnvironmentKHBase captures enum value "kHBase"
	CommonSourceRegistrationRequestParamsEnvironmentKHBase string = "kHBase"

	// CommonSourceRegistrationRequestParamsEnvironmentKSAPHANA captures enum value "kSAPHANA"
	CommonSourceRegistrationRequestParamsEnvironmentKSAPHANA string = "kSAPHANA"

	// CommonSourceRegistrationRequestParamsEnvironmentKUDA captures enum value "kUDA"
	CommonSourceRegistrationRequestParamsEnvironmentKUDA string = "kUDA"

	// CommonSourceRegistrationRequestParamsEnvironmentKSQL captures enum value "kSQL"
	CommonSourceRegistrationRequestParamsEnvironmentKSQL string = "kSQL"

	// CommonSourceRegistrationRequestParamsEnvironmentKOracle captures enum value "kOracle"
	CommonSourceRegistrationRequestParamsEnvironmentKOracle string = "kOracle"

	// CommonSourceRegistrationRequestParamsEnvironmentKSfdc captures enum value "kSfdc"
	CommonSourceRegistrationRequestParamsEnvironmentKSfdc string = "kSfdc"

	// CommonSourceRegistrationRequestParamsEnvironmentKExperimentalAdapter captures enum value "kExperimentalAdapter"
	CommonSourceRegistrationRequestParamsEnvironmentKExperimentalAdapter string = "kExperimentalAdapter"

	// CommonSourceRegistrationRequestParamsEnvironmentKMongoDBPhysical captures enum value "kMongoDBPhysical"
	CommonSourceRegistrationRequestParamsEnvironmentKMongoDBPhysical string = "kMongoDBPhysical"
)

// prop value enum
func (m *CommonSourceRegistrationRequestParams) validateEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonSourceRegistrationRequestParamsTypeEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonSourceRegistrationRequestParams) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	// value enum
	if err := m.validateEnvironmentEnum("environment", "body", *m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *CommonSourceRegistrationRequestParams) validateConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.Connections) { // not required
		return nil
	}

	for i := 0; i < len(m.Connections); i++ {
		if swag.IsZero(m.Connections[i]) { // not required
			continue
		}

		if m.Connections[i] != nil {
			if err := m.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonSourceRegistrationRequestParams) validateAdvancedConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvancedConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.AdvancedConfigs); i++ {
		if swag.IsZero(m.AdvancedConfigs[i]) { // not required
			continue
		}

		if m.AdvancedConfigs[i] != nil {
			if err := m.AdvancedConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advancedConfigs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advancedConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common source registration request params based on the context it is used
func (m *CommonSourceRegistrationRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdvancedConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonSourceRegistrationRequestParams) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connections); i++ {

		if m.Connections[i] != nil {

			if swag.IsZero(m.Connections[i]) { // not required
				return nil
			}

			if err := m.Connections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonSourceRegistrationRequestParams) contextValidateAdvancedConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdvancedConfigs); i++ {

		if m.AdvancedConfigs[i] != nil {

			if swag.IsZero(m.AdvancedConfigs[i]) { // not required
				return nil
			}

			if err := m.AdvancedConfigs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advancedConfigs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advancedConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonSourceRegistrationRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonSourceRegistrationRequestParams) UnmarshalBinary(b []byte) error {
	var res CommonSourceRegistrationRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
