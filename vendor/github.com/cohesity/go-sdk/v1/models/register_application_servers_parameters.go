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

// RegisterApplicationServersParameters Register Application Servers Parameters.
//
// Specifies the parameters required to register Application Servers
// running in a Protection Source.
//
// swagger:model RegisterApplicationServersParameters
type RegisterApplicationServersParameters struct {

	// Specifies the types of applications such as 'kSQL', 'kExchange', 'kAD'
	// running on the Protection Source.
	// overrideDescription: true
	// Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
	// NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.
	// 'kVMware' indicates the VMware Protection Source environment.
	// 'kHyperV' indicates the HyperV Protection Source environment.
	// 'kSQL' indicates the SQL Protection Source environment.
	// 'kView' indicates the View Protection Source environment.
	// 'kPuppeteer' indicates the Cohesity's Remote Adapter.
	// 'kPhysical' indicates the physical Protection Source environment.
	// 'kPure' indicates the Pure Storage Protection Source environment.
	// 'kNimble' indicates the Nimble Storage Protection Source environment.
	// 'kAzure' indicates the Microsoft's Azure Protection Source environment.
	// 'kNetapp' indicates the Netapp Protection Source environment.
	// 'kAgent' indicates the Agent Protection Source environment.
	// 'kGenericNas' indicates the Generic Network Attached Storage Protection
	// Source environment.
	// 'kAcropolis' indicates the Acropolis Protection Source environment.
	// 'kPhysicalFiles' indicates the Physical Files Protection Source environment.
	// 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment.
	// 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment.
	// 'kGPFS' indicates IBM's GPFS Protection Source environment.
	// 'kKVM' indicates the KVM Protection Source environment.
	// 'kAWS' indicates the AWS Protection Source environment.
	// 'kExchange' indicates the Exchange Protection Source environment.
	// 'kHyperVVSS' indicates the HyperV VSS Protection Source
	// environment.
	// 'kOracle' indicates the Oracle Protection Source environment.
	// 'kGCP' indicates the Google Cloud Platform Protection Source environment.
	// 'kFlashBlade' indicates the Flash Blade Protection Source environment.
	// 'kAWSNative' indicates the AWS Native Protection Source environment.
	// 'kO365' indicates the Office 365 Protection Source environment.
	// 'kO365Outlook' indicates Office 365 outlook Protection Source environment.
	// 'kHyperFlex' indicates the Hyper Flex Protection Source environment.
	// 'kGCPNative' indicates the GCP Native Protection Source environment.
	// 'kAzureNative' indicates the Azure Native Protection Source environment.
	// 'kKubernetes' indicates a Kubernetes Protection Source environment.
	// 'kElastifile' indicates Elastifile Protection Source environment.
	// 'kAD' indicates Active Directory Protection Source environment.
	// 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment.
	// 'kCassandra' indicates Cassandra Protection Source environment.
	// 'kMongoDB' indicates MongoDB Protection Source environment.
	// 'kCouchbase' indicates Couchbase Protection Source environment.
	// 'kHdfs' indicates Hdfs Protection Source environment.
	// 'kHive' indicates Hive Protection Source environment.
	// 'kHBase' indicates HBase Protection Source environment.
	// 'kUDA' indicates Universal Data Adapter Protection Source environment.
	// 'kSAPHANA' indicates SAP HANA protection source environment.
	// 'kO365Teams' indicates the Office365 Teams Protection Source environment.
	// 'kO365Group' indicates the Office365 Groups Protection Source environment.
	// 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment.
	// 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment.
	// 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment.
	// 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment.
	// kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis,
	// kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle,
	// kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative,
	// kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager,
	// kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA,
	// kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint,
	// kO365PublicFolders
	Applications []string `json:"applications"`

	// Specifies the cloud credentials used to authenticate with cloud(Aws).
	CloudCredentials *CloudCredentials `json:"cloudCredentials,omitempty"`

	// If set, user has encrypted the credential with 'user_ecryption_key'.
	// It is assumed that credentials are first encrypted using
	// internal magento key and then encrypted using user encryption key.
	EncryptionKey *string `json:"encryptionKey,omitempty"`

	// Set this to true if a persistent agent is running on the host. If this is
	// specified, then credentials would not be used to log into the host
	// environment. This mechanism may be used in environments such as VMware
	// to get around UAC permission issues by running the agent as a service
	// with the right credentials. If this field is not specified, credentials
	// must be specified.
	HasPersistentAgent *bool `json:"hasPersistentAgent,omitempty"`

	// Set to true if credentials are encrypted by internal magneto key.
	IsInternalEncrypted *bool `json:"isInternalEncrypted,omitempty"`

	// Specifies password of the username to access the target source.
	Password *string `json:"password,omitempty"`

	// Specifies the Id of the Protection Source that contains one or more
	// Application Servers running on it.
	ProtectionSourceID *int64 `json:"protectionSourceId,omitempty"`

	// Specifies username to access the target source.
	Username *string `json:"username,omitempty"`
}

// Validate validates this register application servers parameters
func (m *RegisterApplicationServersParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var registerApplicationServersParametersApplicationsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kSQL","kView","kPuppeteer","kPhysical","kPure","kNimble"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registerApplicationServersParametersApplicationsItemsEnum = append(registerApplicationServersParametersApplicationsItemsEnum, v)
	}
}

func (m *RegisterApplicationServersParameters) validateApplicationsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, registerApplicationServersParametersApplicationsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RegisterApplicationServersParameters) validateApplications(formats strfmt.Registry) error {
	if swag.IsZero(m.Applications) { // not required
		return nil
	}

	for i := 0; i < len(m.Applications); i++ {

		// value enum
		if err := m.validateApplicationsItemsEnum("applications"+"."+strconv.Itoa(i), "body", m.Applications[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *RegisterApplicationServersParameters) validateCloudCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudCredentials) { // not required
		return nil
	}

	if m.CloudCredentials != nil {
		if err := m.CloudCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this register application servers parameters based on the context it is used
func (m *RegisterApplicationServersParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterApplicationServersParameters) contextValidateCloudCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudCredentials != nil {

		if swag.IsZero(m.CloudCredentials) { // not required
			return nil
		}

		if err := m.CloudCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisterApplicationServersParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterApplicationServersParameters) UnmarshalBinary(b []byte) error {
	var res RegisterApplicationServersParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
