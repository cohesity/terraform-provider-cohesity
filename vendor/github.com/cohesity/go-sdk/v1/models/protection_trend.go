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

// ProtectionTrend Protection Trend.
//
// Specifies details of a protected object with it's protection trends.
//
// swagger:model ProtectionTrend
type ProtectionTrend struct {

	// Specifies number of cancelled runs across trends.
	Cancelled *int64 `json:"cancelled,omitempty"`

	// Specifies environment.
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
	// Enum: ["kVMware","kHyperV","kSQL","kView","kPuppeteer","kPhysical","kPure","kNimble"]
	Environment *string `json:"environment,omitempty"`

	// Specifies number of failed runs across trends.
	Failed *int64 `json:"failed,omitempty"`

	// Specifies protected object's Id.
	ID *int64 `json:"id,omitempty"`

	// Specifies protected object's name.
	Name *string `json:"name,omitempty"`

	// Specifies protected object's parent id.
	ParentSourceID *int64 `json:"parentSourceId,omitempty"`

	// Specifies protected object's parent name.
	ParentSourceName *string `json:"parentSourceName,omitempty"`

	// Specifies number of in-progress runs across trends.
	Running *int64 `json:"running,omitempty"`

	// Specifies number of successful runs across trends.
	Successful *int64 `json:"successful,omitempty"`

	// Specifies total number of runs across trends.
	Total *int64 `json:"total,omitempty"`

	// Aggregated protection runs information by days/weeks.
	Trends []*TrendingData `json:"trends"`
}

// Validate validates this protection trend
func (m *ProtectionTrend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrends(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var protectionTrendTypeEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kSQL","kView","kPuppeteer","kPhysical","kPure","kNimble"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectionTrendTypeEnvironmentPropEnum = append(protectionTrendTypeEnvironmentPropEnum, v)
	}
}

const (

	// ProtectionTrendEnvironmentKVMware captures enum value "kVMware"
	ProtectionTrendEnvironmentKVMware string = "kVMware"

	// ProtectionTrendEnvironmentKHyperV captures enum value "kHyperV"
	ProtectionTrendEnvironmentKHyperV string = "kHyperV"

	// ProtectionTrendEnvironmentKSQL captures enum value "kSQL"
	ProtectionTrendEnvironmentKSQL string = "kSQL"

	// ProtectionTrendEnvironmentKView captures enum value "kView"
	ProtectionTrendEnvironmentKView string = "kView"

	// ProtectionTrendEnvironmentKPuppeteer captures enum value "kPuppeteer"
	ProtectionTrendEnvironmentKPuppeteer string = "kPuppeteer"

	// ProtectionTrendEnvironmentKPhysical captures enum value "kPhysical"
	ProtectionTrendEnvironmentKPhysical string = "kPhysical"

	// ProtectionTrendEnvironmentKPure captures enum value "kPure"
	ProtectionTrendEnvironmentKPure string = "kPure"

	// ProtectionTrendEnvironmentKNimble captures enum value "kNimble"
	ProtectionTrendEnvironmentKNimble string = "kNimble"
)

// prop value enum
func (m *ProtectionTrend) validateEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectionTrendTypeEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectionTrend) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnvironmentEnum("environment", "body", *m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *ProtectionTrend) validateTrends(formats strfmt.Registry) error {
	if swag.IsZero(m.Trends) { // not required
		return nil
	}

	for i := 0; i < len(m.Trends); i++ {
		if swag.IsZero(m.Trends[i]) { // not required
			continue
		}

		if m.Trends[i] != nil {
			if err := m.Trends[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trends" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this protection trend based on the context it is used
func (m *ProtectionTrend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrends(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionTrend) contextValidateTrends(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Trends); i++ {

		if m.Trends[i] != nil {

			if swag.IsZero(m.Trends[i]) { // not required
				return nil
			}

			if err := m.Trends[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trends" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionTrend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionTrend) UnmarshalBinary(b []byte) error {
	var res ProtectionTrend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
