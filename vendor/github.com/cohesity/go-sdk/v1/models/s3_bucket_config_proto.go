// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// S3BucketConfigProto Proto to define the config/metadata of a S3 bucket.
//
// swagger:model S3BucketConfigProto
type S3BucketConfigProto struct {

	// bool representing if the view is ABAC enabled or not
	AbacEnabled *bool `json:"abacEnabled,omitempty"`

	// ACL of the bucket.
	ACL *ACLProto `json:"acl,omitempty"`

	// Policy in effect for the bucket.
	BucketPolicy *BucketPolicy `json:"bucketPolicy,omitempty"`

	// bool representing whether MPUs are organized using S3 MPU 2.0 design.
	// Should only be set while view creation and immutable there after.
	EfficientMpuEnabled *bool `json:"efficientMpuEnabled,omitempty"`

	// If set to false, we disable access to view using S3/Swift protocol. This
	// overrides any 'protocol_access_info' set on view for S3/Swift. This flag
	// is added as the transition for S3 native to non-S3 native is disabled
	// and therefore the access using S3/Swift protocol cannot be disabled by
	// madrox.
	EnableObjStoreAccess *bool `json:"enableObjStoreAccess,omitempty"`

	// The cluster id for the cluster where the view was initially created
	// without replication. For view on Rx, the init_cluster_id will remain same
	// as that of the initial cluster.
	InitClusterID *int64 `json:"initClusterId,omitempty"`

	// The cluster incarnation id for the cluster where the view was initially
	// created without replication. For view on Rx, the
	// init_incarnation_cluster_id will remain same as that of the initial
	// cluster.
	InitClusterIncarnationID *int64 `json:"initClusterIncarnationId,omitempty"`

	// Lifecycle policy of the bucket. If not specified, no lifecycle management
	// is performed for objects in this bucket.
	LifecycleConfig *LifecycleConfigProto `json:"lifecycleConfig,omitempty"`

	// This tunable field defines the number of maximum subfiles a
	// MPU directory can have.
	MaxSubfilesPerMpu *int32 `json:"maxSubfilesPerMpu,omitempty"`

	// For a view under-migration from S3 1.0 to 2.0, we will use newly created
	// object and versions directories with following suffix in obj namespace.
	// Prefix for these root directories is defined in s3_constants.cc
	ObjNsRootDirsSuffix *string `json:"objNsRootDirsSuffix,omitempty"`

	// This field stores all the metadata for migration of bucket from
	// S3 1.0 to 2.0
	ObjectIDMigration *ObjectIDMigrationProto `json:"objectIdMigration,omitempty"`

	// Whether this view has ever had any object with tags. This should be
	// enabled only when first object tag is ever put in this view.
	ObjectTagsAdded *bool `json:"objectTagsAdded,omitempty"`

	// Information about the bucket owner.
	OwnerInfo *OwnerInfo `json:"ownerInfo,omitempty"`

	// Bucket Ownership Controls for the bucket.
	OwnershipControls *BucketOwnershipControls `json:"ownershipControls,omitempty"`

	// Delimiter used in prefix based request routing. An application needs to
	// explicitly set the prefix_delimiter during bucket creation. If the
	// prefix_delimiter is not explicitly set, '/' will be used as the default
	// prefix delimiter for buckets that has prefix-based-routing enabled.
	// SnapDiff backups uses '/' in the object names hence it was chosen as the
	// default prefix to avoid further UI changes in this project. If there are
	// other use cases that require a different prefix_delimiter, it has to be
	// set during bucket creation. Once prefix_delimiter is chosen, it cannot be
	// changed.
	PrefixDelimiter *string `json:"prefixDelimiter,omitempty"`

	// Stores the prefix to child bucket mapping to enable prefix based routing
	// of object requests to child buckets.
	PrefixToChildBucketMap map[string]string `json:"prefixToChildBucketMap,omitempty"`

	// Protocol type of this bucket.
	ProtocolType *int32 `json:"protocolType,omitempty"`

	// Swfit container tagging information.
	SwiftContainerTag *SwiftContainerTaggingProto `json:"swiftContainerTag,omitempty"`

	// Tags (or labels) assigned to the bucket. Tags are set of <key, value>
	// pairs.
	TagMap map[string]string `json:"tagMap,omitempty"`

	// versioning state
	VersioningState *int32 `json:"versioningState,omitempty"`
}

// Validate validates this s3 bucket config proto
func (m *S3BucketConfigProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucketPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycleConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectIDMigration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnershipControls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwiftContainerTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketConfigProto) validateACL(formats strfmt.Registry) error {
	if swag.IsZero(m.ACL) { // not required
		return nil
	}

	if m.ACL != nil {
		if err := m.ACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acl")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) validateBucketPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.BucketPolicy) { // not required
		return nil
	}

	if m.BucketPolicy != nil {
		if err := m.BucketPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucketPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucketPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) validateLifecycleConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.LifecycleConfig) { // not required
		return nil
	}

	if m.LifecycleConfig != nil {
		if err := m.LifecycleConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycleConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lifecycleConfig")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) validateObjectIDMigration(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectIDMigration) { // not required
		return nil
	}

	if m.ObjectIDMigration != nil {
		if err := m.ObjectIDMigration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectIdMigration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectIdMigration")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) validateOwnerInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerInfo) { // not required
		return nil
	}

	if m.OwnerInfo != nil {
		if err := m.OwnerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) validateOwnershipControls(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnershipControls) { // not required
		return nil
	}

	if m.OwnershipControls != nil {
		if err := m.OwnershipControls.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownershipControls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownershipControls")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) validateSwiftContainerTag(formats strfmt.Registry) error {
	if swag.IsZero(m.SwiftContainerTag) { // not required
		return nil
	}

	if m.SwiftContainerTag != nil {
		if err := m.SwiftContainerTag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swiftContainerTag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swiftContainerTag")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s3 bucket config proto based on the context it is used
func (m *S3BucketConfigProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBucketPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLifecycleConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectIDMigration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwnerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwnershipControls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwiftContainerTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3BucketConfigProto) contextValidateACL(ctx context.Context, formats strfmt.Registry) error {

	if m.ACL != nil {

		if swag.IsZero(m.ACL) { // not required
			return nil
		}

		if err := m.ACL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acl")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) contextValidateBucketPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.BucketPolicy != nil {

		if swag.IsZero(m.BucketPolicy) { // not required
			return nil
		}

		if err := m.BucketPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucketPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucketPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) contextValidateLifecycleConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.LifecycleConfig != nil {

		if swag.IsZero(m.LifecycleConfig) { // not required
			return nil
		}

		if err := m.LifecycleConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycleConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lifecycleConfig")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) contextValidateObjectIDMigration(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectIDMigration != nil {

		if swag.IsZero(m.ObjectIDMigration) { // not required
			return nil
		}

		if err := m.ObjectIDMigration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectIdMigration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("objectIdMigration")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) contextValidateOwnerInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OwnerInfo != nil {

		if swag.IsZero(m.OwnerInfo) { // not required
			return nil
		}

		if err := m.OwnerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) contextValidateOwnershipControls(ctx context.Context, formats strfmt.Registry) error {

	if m.OwnershipControls != nil {

		if swag.IsZero(m.OwnershipControls) { // not required
			return nil
		}

		if err := m.OwnershipControls.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ownershipControls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ownershipControls")
			}
			return err
		}
	}

	return nil
}

func (m *S3BucketConfigProto) contextValidateSwiftContainerTag(ctx context.Context, formats strfmt.Registry) error {

	if m.SwiftContainerTag != nil {

		if swag.IsZero(m.SwiftContainerTag) { // not required
			return nil
		}

		if err := m.SwiftContainerTag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swiftContainerTag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swiftContainerTag")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketConfigProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketConfigProto) UnmarshalBinary(b []byte) error {
	var res S3BucketConfigProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
