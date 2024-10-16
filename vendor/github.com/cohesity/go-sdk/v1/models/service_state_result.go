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

// ServiceStateResult Service State Result.
//
// Specifies the result of querying the state of a specific service on
// the Cluster.
//
// swagger:model ServiceStateResult
type ServiceStateResult struct {

	// Specifies the name of the service.
	// 'kApollo' is a service for reclaiming freed disk sectors on Nodes in the
	// SnapFS distributed file system.
	// 'kBridge' is a service for managing the SnapFS distributed file system.
	// 'kGenie' is a service that is responsible for monitoring hardware health
	// on the Cluster.
	// 'kGenieGofer' is a service that links the Genie service to other services
	// on the Cluster.
	// 'kMagneto' is the data protection service of the Cohesity Data Platform.
	// 'kIris' is the service which serves REST API calls to the UI, CLI, and any
	// scripts written by customers.
	// 'kIrisProxy' is a service that links the Iris service to other services
	// on the Cluster.
	// 'kScribe' is the service responsible for storing filesystem metadata.
	// 'kStats' is the service that is responsible for retrieving and aggregating
	// disk metrics across the Cluster.
	// 'kYoda' is an elastic search indexing service.
	// 'kAlerts' is a publisher and subscribing service for alerts.
	// 'kKeychain' is a service for managing disk encryption keys.
	// 'kLogWatcher' is a service that scans the log directory and reduces
	// the number of logs if required.
	// 'kStatsCollector' is a service that periodically logs system stats.
	// 'kGandalf' is a distributed lock service and coordination manager.
	// 'kNexus' indicates the Nexus service. This is the service that is
	// responsible for creation of Clusters and configuration of Nodes and
	// networking.
	// 'kNexusProxy' is a service that links the Nexus service to other services
	// on the Cluster.
	// 'kStorageProxy' is a service for accessing data on external entities.
	// 'kRtClient' is a reverse tunneling client service.
	// 'kVaultProxy' is a service for managing external targets that Clusters
	// can be backed up to.
	// 'kSmbProxy' is an SMB protocol service.
	// 'kBridgeProxy' is the service that links the Bridge service to other
	// services on the Cluster.
	// 'kLibrarian' is an elastic search indexing service.
	// 'kGroot' is a service for managing replication of SQL databases across
	// multiple nodes in a Cluster.
	// 'kEagleAgent' is a service that is responsible for retrieving information
	// on Cluster health.
	// 'kAthena' is a service for running distributed containerized applications
	// on the Cohesity Data Platform.
	// 'kBifrostBroker' is a service for communicating with the Cohesity proxies
	// for multitenancy.
	// 'kSmb2Proxy' is a new SMB protocol service.
	// 'kOs' can be specified in order to do a full reboot.
	// 'kAtom' is a service for receiving data for the Continuous Data Protection.
	// 'kPatch' is a service for downloading and applying patches.
	// 'kCompass' is a service for serving dns request for external and internal
	// traffic.
	// 'kEtlServer' is a service responsible for ETling data for globalsearch.
	// 'kIcebox' is service that links Icebox service to other services on cluster.
	// kScribe, kStats, kYoda, kAlerts, kKeychain, kLogWatcher, kStatsCollecter,
	// kGandalf, kNexus, kNexusProxy, kStorageProxy, kRtClient, kVaultProxy,
	// kSmbProxy, kBridgeProxy, kLibrarian, kGroot, kEagleAgent, kAthena,
	// kBifrostBroker, kSmb2Proxy, kOs, kAtom, kIcebox
	// Enum: ["kApollo","kBridge","kGenie","kGenieGofer","kMagneto","kIris","kIrisProxy"]
	Service *string `json:"service,omitempty"`

	// Specifies the state of the service.
	// 'kServiceStopped' indicates that the service has been stopped.
	// 'kServiceRunning' indicates that the service is currently running.
	// 'kServiceRestarting' indicates that the service is in the queue to be
	// restarted.
	// Enum: ["kServiceStopped","kServiceRunning","kServiceRestarting"]
	State *string `json:"state,omitempty"`
}

// Validate validates this service state result
func (m *ServiceStateResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceStateResultTypeServicePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kApollo","kBridge","kGenie","kGenieGofer","kMagneto","kIris","kIrisProxy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceStateResultTypeServicePropEnum = append(serviceStateResultTypeServicePropEnum, v)
	}
}

const (

	// ServiceStateResultServiceKApollo captures enum value "kApollo"
	ServiceStateResultServiceKApollo string = "kApollo"

	// ServiceStateResultServiceKBridge captures enum value "kBridge"
	ServiceStateResultServiceKBridge string = "kBridge"

	// ServiceStateResultServiceKGenie captures enum value "kGenie"
	ServiceStateResultServiceKGenie string = "kGenie"

	// ServiceStateResultServiceKGenieGofer captures enum value "kGenieGofer"
	ServiceStateResultServiceKGenieGofer string = "kGenieGofer"

	// ServiceStateResultServiceKMagneto captures enum value "kMagneto"
	ServiceStateResultServiceKMagneto string = "kMagneto"

	// ServiceStateResultServiceKIris captures enum value "kIris"
	ServiceStateResultServiceKIris string = "kIris"

	// ServiceStateResultServiceKIrisProxy captures enum value "kIrisProxy"
	ServiceStateResultServiceKIrisProxy string = "kIrisProxy"
)

// prop value enum
func (m *ServiceStateResult) validateServiceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceStateResultTypeServicePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceStateResult) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceEnum("service", "body", *m.Service); err != nil {
		return err
	}

	return nil
}

var serviceStateResultTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kServiceStopped","kServiceRunning","kServiceRestarting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceStateResultTypeStatePropEnum = append(serviceStateResultTypeStatePropEnum, v)
	}
}

const (

	// ServiceStateResultStateKServiceStopped captures enum value "kServiceStopped"
	ServiceStateResultStateKServiceStopped string = "kServiceStopped"

	// ServiceStateResultStateKServiceRunning captures enum value "kServiceRunning"
	ServiceStateResultStateKServiceRunning string = "kServiceRunning"

	// ServiceStateResultStateKServiceRestarting captures enum value "kServiceRestarting"
	ServiceStateResultStateKServiceRestarting string = "kServiceRestarting"
)

// prop value enum
func (m *ServiceStateResult) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceStateResultTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceStateResult) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service state result based on context it is used
func (m *ServiceStateResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceStateResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceStateResult) UnmarshalBinary(b []byte) error {
	var res ServiceStateResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
