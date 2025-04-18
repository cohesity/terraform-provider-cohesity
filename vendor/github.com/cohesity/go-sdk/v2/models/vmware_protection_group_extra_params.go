// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VmwareProtectionGroupExtraParams Specifies the extra parameters which are specific to VMware object protection Group.
//
// swagger:model VmwareProtectionGroupExtraParams
type VmwareProtectionGroupExtraParams struct {

	// Specifies the id of the parent of the objects.
	// Read Only: true
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies the name of the parent of the objects.
	// Read Only: true
	SourceName *string `json:"sourceName,omitempty"`

	// Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection.
	ExcludeObjectIds []*int64 `json:"excludeObjectIds"`

	// Array of Array of VM Tag Ids that Specify VMs to Protect. Optionally specify a list of VMs to protect by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to protect which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. To protect only 'Eng' VMs in the East and all the VMs in the West, specify the following tag id array: [ [1101, 2221], [3031] ], where 1101 is the 'Eng' VM Tag id, 2221 is the 'East' VM Tag id and 3031 is the 'West' VM Tag id. The inner array [1101, 2221] produces a list of VMs that are both tagged with 'Eng' and 'East' (an intersection). The outer array combines the list from the inner array with list of VMs tagged with 'West' (a union). The list of resulting VMs are protected by this Protection Group.
	VMTagIds [][]int64 `json:"vmTagIds"`

	// Array of Arrays of VM Tag Ids that Specify VMs to Exclude. Optionally specify a list of VMs to exclude from protecting by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. For example a Datacenter is selected to be protected but you want to exclude all the 'Former Employees' VMs in the East and West but keep all the VMs for 'Former Employees' in the South which are also stored in this Datacenter, by specifying the following tag id array: [ [1000, 2221], [1000, 3031] ], where 1000 is the 'Former Employee' VM Tag id, 2221 is the 'East' VM Tag id and 3031 is the 'West' VM Tag id. The first inner array [1000, 2221] produces a list of VMs that are both tagged with 'Former Employees' and 'East' (an intersection). The second inner array [1000, 3031] produces a list of VMs that are both tagged with 'Former Employees' and 'West' (an intersection). The outer array combines the list of VMs from the two inner arrays. The list of resulting VMs are excluded from being protected this Job.
	ExcludeVMTagIds [][]int64 `json:"excludeVmTagIds"`

	// Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying these filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters.
	ExcludeFilters []*VMFilter `json:"excludeFilters"`

	// Whether to leverage the storage array based snapshots for this backup. To leverage storage snapshots, the storage array has to be registered as a source. If storage based snapshots can not be taken, backup will fallback to the default backup method.
	LeverageStorageSnapshots *bool `json:"leverageStorageSnapshots,omitempty"`

	// Whether to leverage the hyperflex based snapshots for this backup. To leverage hyperflex snapshots, it has to first be registered. If hyperflex based snapshots cannot be taken, backup will fallback to the default backup method.
	LeverageHyperflexSnapshots *bool `json:"leverageHyperflexSnapshots,omitempty"`

	// Whether to leverage the nutanix based snapshots for this backup. To leverage nutanix snapshots, it has to first be registered. If nutanix based snapshots cannot be taken, backup will fallback to the default backup method.
	LeverageNutanixSnapshots *bool `json:"leverageNutanixSnapshots,omitempty"`

	// Specifies whether or not to move the workload to the cloud.
	CloudMigration *bool `json:"cloudMigration,omitempty"`

	// Specifies whether or not this job can have parallel runs.
	AllowParallelRuns *bool `json:"allowParallelRuns,omitempty"`
}

// Validate validates this vmware protection group extra params
func (m *VmwareProtectionGroupExtraParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExcludeFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareProtectionGroupExtraParams) validateExcludeFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.ExcludeFilters); i++ {
		if swag.IsZero(m.ExcludeFilters[i]) { // not required
			continue
		}

		if m.ExcludeFilters[i] != nil {
			if err := m.ExcludeFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludeFilters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludeFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vmware protection group extra params based on the context it is used
func (m *VmwareProtectionGroupExtraParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcludeFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareProtectionGroupExtraParams) contextValidateSourceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

func (m *VmwareProtectionGroupExtraParams) contextValidateSourceName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sourceName", "body", m.SourceName); err != nil {
		return err
	}

	return nil
}

func (m *VmwareProtectionGroupExtraParams) contextValidateExcludeFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExcludeFilters); i++ {

		if m.ExcludeFilters[i] != nil {

			if swag.IsZero(m.ExcludeFilters[i]) { // not required
				return nil
			}

			if err := m.ExcludeFilters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludeFilters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludeFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareProtectionGroupExtraParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareProtectionGroupExtraParams) UnmarshalBinary(b []byte) error {
	var res VmwareProtectionGroupExtraParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
