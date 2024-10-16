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

// AwsCredentials AWS source Credentials.
//
// Specifies the credentials to authenticate with AWS Cloud Platform.
//
// swagger:model AwsCredentials
type AwsCredentials struct {

	// Specifies Access key of the AWS account.
	AccessKey *string `json:"accessKey,omitempty"`

	// Specifies Amazon Resource Name (owner ID) of the IAM user, act as an
	// unique identifier of as AWS entity.
	AmazonResourceName *string `json:"amazonResourceName,omitempty"`

	// Specifies the authentication method to be used for API calls.
	// Specifies the authentication method to be used for API calls.
	// 'kUseIAMUser' indicates a user based authentication.
	// 'kUseIAMRole' indicates a role based authentication, used only for AWS CE.
	// 'kUseHelios' indicates a Helios based authentication.
	// Enum: ["kUseIAMUser","kUseIAMRole","kUseHelios"]
	AuthMethod *string `json:"authMethod,omitempty"`

	// Specifies the entity type such as 'kIAMUser' if the environment is kAWS.
	// Specifies the type of an AWS source entity.
	// 'kIAMUser' indicates a unique user within an AWS account.
	// 'kRegion' indicates a geographical region in the global infrastructure.
	// 'kAvailabilityZone' indicates an availability zone within a region.
	// 'kEC2Instance' indicates a Virtual Machine running in AWS environment.
	// 'kVPC' indicates a virtual private cloud (VPC) network within AWS.
	// 'kSubnet' indicates a subnet inside the VPC.
	// 'kNetworkSecurityGroup' represents a network security group.
	// 'kInstanceType' represents various machine types.
	// 'kKeyPair' represents a pair of public and private key used to login into a Virtual Machine.
	// 'kTag' represents a tag attached to EC2 instance.
	// 'kRDSOptionGroup' represents a RDS option group for configuring database features.
	// 'kRDSParameterGroup' represents a RDS parameter group.
	// 'kRDSInstance' represents a RDS DB instance.
	// 'kRDSSubnet' represents a RDS subnet.
	// 'kRDSTag' represents a tag attached to RDS instance.
	// 'kAuroraTag' represents a tag attached to an Aurora cluster.
	// 'kAccount' represents an AWS account.
	// 'kAuroraCluster' represents an Aurora cluster.
	// 'kSubTaskPermit' entity type will be used by slave sub-tasks to take permit for native backups,
	// so that we can control the number of concurrent sub-tasks independent of the number of VMs being backed up concurrently.
	// This does not represent any entity type in AWS entity hierarchy.
	// 'kS3Bucket' represents an S3 bucket.
	// 'kS3Tag' represents an S3 tag attached to S3 Bucket.
	// 'kKmsKey' represents a KMS key.
	// Enum: ["kIAMUser","kRegion","kAvailabilityZone","kEC2Instance","kVPC","kSubnet","kNetworkSecurityGroup","kInstanceType","kKeyPair","kTag","kRDSOptionGroup","kRDSParameterGroup","kRDSInstance","kRDSSubnet","kRDSTag","kAuroraTag","kAccount","kAuroraCluster","kSubTaskPermit","kS3Bucket","kS3Tag","kKmsKey"]
	AwsType *string `json:"awsType,omitempty"`

	// Specifies the C2S Access Portal (CAP) server info.
	C2sServerInfo *C2SServerInfo `json:"c2sServerInfo,omitempty"`

	// Specifies the IAM role which will be used to access the security
	// credentials required for API calls.
	IamRoleArn *string `json:"iamRoleArn,omitempty"`

	// Specifies Secret Access key of the AWS account.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`

	// Specifies the subscription type of AWS such as 'kAWSCommercial' or
	// 'kAWSGovCloud'.
	// Specifies the subscription type of an AWS source entity.
	// 'kAWSCommercial' indicates a standard AWS subscription.
	// 'kAWSGovCloud' indicates a govt AWS subscription.
	// Enum: ["kAWSCommercial","kAWSGovCloud"]
	SubscriptionType *string `json:"subscriptionType,omitempty"`
}

// Validate validates this aws credentials
func (m *AwsCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateC2sServerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var awsCredentialsTypeAuthMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kUseIAMUser","kUseIAMRole","kUseHelios"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsCredentialsTypeAuthMethodPropEnum = append(awsCredentialsTypeAuthMethodPropEnum, v)
	}
}

const (

	// AwsCredentialsAuthMethodKUseIAMUser captures enum value "kUseIAMUser"
	AwsCredentialsAuthMethodKUseIAMUser string = "kUseIAMUser"

	// AwsCredentialsAuthMethodKUseIAMRole captures enum value "kUseIAMRole"
	AwsCredentialsAuthMethodKUseIAMRole string = "kUseIAMRole"

	// AwsCredentialsAuthMethodKUseHelios captures enum value "kUseHelios"
	AwsCredentialsAuthMethodKUseHelios string = "kUseHelios"
)

// prop value enum
func (m *AwsCredentials) validateAuthMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsCredentialsTypeAuthMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsCredentials) validateAuthMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthMethodEnum("authMethod", "body", *m.AuthMethod); err != nil {
		return err
	}

	return nil
}

var awsCredentialsTypeAwsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kIAMUser","kRegion","kAvailabilityZone","kEC2Instance","kVPC","kSubnet","kNetworkSecurityGroup","kInstanceType","kKeyPair","kTag","kRDSOptionGroup","kRDSParameterGroup","kRDSInstance","kRDSSubnet","kRDSTag","kAuroraTag","kAccount","kAuroraCluster","kSubTaskPermit","kS3Bucket","kS3Tag","kKmsKey"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsCredentialsTypeAwsTypePropEnum = append(awsCredentialsTypeAwsTypePropEnum, v)
	}
}

const (

	// AwsCredentialsAwsTypeKIAMUser captures enum value "kIAMUser"
	AwsCredentialsAwsTypeKIAMUser string = "kIAMUser"

	// AwsCredentialsAwsTypeKRegion captures enum value "kRegion"
	AwsCredentialsAwsTypeKRegion string = "kRegion"

	// AwsCredentialsAwsTypeKAvailabilityZone captures enum value "kAvailabilityZone"
	AwsCredentialsAwsTypeKAvailabilityZone string = "kAvailabilityZone"

	// AwsCredentialsAwsTypeKEC2Instance captures enum value "kEC2Instance"
	AwsCredentialsAwsTypeKEC2Instance string = "kEC2Instance"

	// AwsCredentialsAwsTypeKVPC captures enum value "kVPC"
	AwsCredentialsAwsTypeKVPC string = "kVPC"

	// AwsCredentialsAwsTypeKSubnet captures enum value "kSubnet"
	AwsCredentialsAwsTypeKSubnet string = "kSubnet"

	// AwsCredentialsAwsTypeKNetworkSecurityGroup captures enum value "kNetworkSecurityGroup"
	AwsCredentialsAwsTypeKNetworkSecurityGroup string = "kNetworkSecurityGroup"

	// AwsCredentialsAwsTypeKInstanceType captures enum value "kInstanceType"
	AwsCredentialsAwsTypeKInstanceType string = "kInstanceType"

	// AwsCredentialsAwsTypeKKeyPair captures enum value "kKeyPair"
	AwsCredentialsAwsTypeKKeyPair string = "kKeyPair"

	// AwsCredentialsAwsTypeKTag captures enum value "kTag"
	AwsCredentialsAwsTypeKTag string = "kTag"

	// AwsCredentialsAwsTypeKRDSOptionGroup captures enum value "kRDSOptionGroup"
	AwsCredentialsAwsTypeKRDSOptionGroup string = "kRDSOptionGroup"

	// AwsCredentialsAwsTypeKRDSParameterGroup captures enum value "kRDSParameterGroup"
	AwsCredentialsAwsTypeKRDSParameterGroup string = "kRDSParameterGroup"

	// AwsCredentialsAwsTypeKRDSInstance captures enum value "kRDSInstance"
	AwsCredentialsAwsTypeKRDSInstance string = "kRDSInstance"

	// AwsCredentialsAwsTypeKRDSSubnet captures enum value "kRDSSubnet"
	AwsCredentialsAwsTypeKRDSSubnet string = "kRDSSubnet"

	// AwsCredentialsAwsTypeKRDSTag captures enum value "kRDSTag"
	AwsCredentialsAwsTypeKRDSTag string = "kRDSTag"

	// AwsCredentialsAwsTypeKAuroraTag captures enum value "kAuroraTag"
	AwsCredentialsAwsTypeKAuroraTag string = "kAuroraTag"

	// AwsCredentialsAwsTypeKAccount captures enum value "kAccount"
	AwsCredentialsAwsTypeKAccount string = "kAccount"

	// AwsCredentialsAwsTypeKAuroraCluster captures enum value "kAuroraCluster"
	AwsCredentialsAwsTypeKAuroraCluster string = "kAuroraCluster"

	// AwsCredentialsAwsTypeKSubTaskPermit captures enum value "kSubTaskPermit"
	AwsCredentialsAwsTypeKSubTaskPermit string = "kSubTaskPermit"

	// AwsCredentialsAwsTypeKS3Bucket captures enum value "kS3Bucket"
	AwsCredentialsAwsTypeKS3Bucket string = "kS3Bucket"

	// AwsCredentialsAwsTypeKS3Tag captures enum value "kS3Tag"
	AwsCredentialsAwsTypeKS3Tag string = "kS3Tag"

	// AwsCredentialsAwsTypeKKmsKey captures enum value "kKmsKey"
	AwsCredentialsAwsTypeKKmsKey string = "kKmsKey"
)

// prop value enum
func (m *AwsCredentials) validateAwsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsCredentialsTypeAwsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsCredentials) validateAwsType(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAwsTypeEnum("awsType", "body", *m.AwsType); err != nil {
		return err
	}

	return nil
}

func (m *AwsCredentials) validateC2sServerInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.C2sServerInfo) { // not required
		return nil
	}

	if m.C2sServerInfo != nil {
		if err := m.C2sServerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("c2sServerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("c2sServerInfo")
			}
			return err
		}
	}

	return nil
}

var awsCredentialsTypeSubscriptionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAWSCommercial","kAWSGovCloud"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		awsCredentialsTypeSubscriptionTypePropEnum = append(awsCredentialsTypeSubscriptionTypePropEnum, v)
	}
}

const (

	// AwsCredentialsSubscriptionTypeKAWSCommercial captures enum value "kAWSCommercial"
	AwsCredentialsSubscriptionTypeKAWSCommercial string = "kAWSCommercial"

	// AwsCredentialsSubscriptionTypeKAWSGovCloud captures enum value "kAWSGovCloud"
	AwsCredentialsSubscriptionTypeKAWSGovCloud string = "kAWSGovCloud"
)

// prop value enum
func (m *AwsCredentials) validateSubscriptionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, awsCredentialsTypeSubscriptionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AwsCredentials) validateSubscriptionType(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscriptionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubscriptionTypeEnum("subscriptionType", "body", *m.SubscriptionType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws credentials based on the context it is used
func (m *AwsCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateC2sServerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCredentials) contextValidateC2sServerInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.C2sServerInfo != nil {

		if swag.IsZero(m.C2sServerInfo) { // not required
			return nil
		}

		if err := m.C2sServerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("c2sServerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("c2sServerInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCredentials) UnmarshalBinary(b []byte) error {
	var res AwsCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
