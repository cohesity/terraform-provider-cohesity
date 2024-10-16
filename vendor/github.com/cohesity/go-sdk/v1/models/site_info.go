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
)

// SiteInfo Information about the site backed up that is used to restore. For example
// we need to create the extract site type that was backed up if the site does
// not exist. All fields are case insensitive. This is an aggregation of all
// required and optional parameters in New-PnPSite/New-TenantSite/New-PnPWeb
// cmdlets.
//
// swagger:model SiteInfo
type SiteInfo struct {

	// Site classification assigned by administrator.
	Classification *string `json:"classification,omitempty"`

	// Compatibility level . Its the site's compatibility to SPO server version.
	CompatibilityLevel *int32 `json:"compatibilityLevel,omitempty"`

	// Admin entered description of this site.
	Description *string `json:"description,omitempty"`

	// Needed for site collection create. Maps to CommunicationSiteDesign enum
	// (Topic = 0,Showcase = 1,Blank = 2)
	Design *string `json:"design,omitempty"`

	// Design template id.
	DesignID *string `json:"designId,omitempty"`

	// Geo location of the site.
	Geolocation *string `json:"geolocation,omitempty"`

	// Group Id to which this site belongs.
	GroupID *string `json:"groupId,omitempty"`

	// If it is a hub member site, the parent hub site id.
	HubsiteID *string `json:"hubsiteId,omitempty"`

	// If it is a hub member site, the parent hub site url. This can be used to
	// restore a hub site member to same hub site even when the site is deleted
	// and recreated or across tenants.
	HubsiteURL *string `json:"hubsiteUrl,omitempty"`

	// Is it a Hub? This will be false for a hub member.
	IsHubsite *bool `json:"isHubsite,omitempty"`

	// Site is multilingual, needs to get multilingual resource pages.
	IsMultilingual *bool `json:"isMultilingual,omitempty"`

	// Is this a personal site with 'https://*-my.sharepoint.com/*' pattern.
	IsPersonalsite *bool `json:"isPersonalsite,omitempty"`

	// Is this a public or private site?
	IsPublic *bool `json:"isPublic,omitempty"`

	// Is this a subsite or root site? If this is a subsite, it will inherit the
	// template of root site.
	IsSubsite *bool `json:"isSubsite,omitempty"`

	// Site LcId=1033 etc..
	Lcid *int32 `json:"lcid,omitempty"`

	// Locale id of the site. Normally 1033 for US English locale.
	LocaleID *int32 `json:"localeId,omitempty"`

	// Owner. This is an email. Note that across tenants, this email address may
	// be invalid. user@tenant.com. At least one owner must be there.
	OwnerVec []string `json:"ownerVec"`

	// Site is read-only, locked, and unavailable for write access.
	ReadOnly *bool `json:"readOnly,omitempty"`

	// Rootwebid.
	RootWebID *string `json:"rootWebId,omitempty"`

	// Site identifier echo.
	Site *SiteIdentity `json:"site,omitempty"`

	// Vector of sites collection properties (for classic site collections). Got
	// from Get-PnPSite.
	SitepropVec []*SiteProperty `json:"sitepropVec"`

	// Site Template such as GROUP#0, POINTPUBLISHINGTOPIC#0,...
	Template *string `json:"template,omitempty"`

	// Vector of tenant sites collection properties (for modern site
	// collections). Got from Get-PnPTenantSite.
	TenantsitepropVec []*SiteProperty `json:"tenantsitepropVec"`

	// Timezone Id. This is needed to create a site collection(tenant site).
	TimezoneID *int32 `json:"timezoneId,omitempty"`

	// Vector of sites collection properties (for subsites). Got from Get-PnPWeb.
	WebpropVec []*SiteProperty `json:"webpropVec"`
}

// Validate validates this site info
func (m *SiteInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitepropVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantsitepropVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebpropVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteInfo) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *SiteInfo) validateSitepropVec(formats strfmt.Registry) error {
	if swag.IsZero(m.SitepropVec) { // not required
		return nil
	}

	for i := 0; i < len(m.SitepropVec); i++ {
		if swag.IsZero(m.SitepropVec[i]) { // not required
			continue
		}

		if m.SitepropVec[i] != nil {
			if err := m.SitepropVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sitepropVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sitepropVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SiteInfo) validateTenantsitepropVec(formats strfmt.Registry) error {
	if swag.IsZero(m.TenantsitepropVec) { // not required
		return nil
	}

	for i := 0; i < len(m.TenantsitepropVec); i++ {
		if swag.IsZero(m.TenantsitepropVec[i]) { // not required
			continue
		}

		if m.TenantsitepropVec[i] != nil {
			if err := m.TenantsitepropVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenantsitepropVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenantsitepropVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SiteInfo) validateWebpropVec(formats strfmt.Registry) error {
	if swag.IsZero(m.WebpropVec) { // not required
		return nil
	}

	for i := 0; i < len(m.WebpropVec); i++ {
		if swag.IsZero(m.WebpropVec[i]) { // not required
			continue
		}

		if m.WebpropVec[i] != nil {
			if err := m.WebpropVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webpropVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webpropVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this site info based on the context it is used
func (m *SiteInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSitepropVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenantsitepropVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebpropVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteInfo) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {

		if swag.IsZero(m.Site) { // not required
			return nil
		}

		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *SiteInfo) contextValidateSitepropVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SitepropVec); i++ {

		if m.SitepropVec[i] != nil {

			if swag.IsZero(m.SitepropVec[i]) { // not required
				return nil
			}

			if err := m.SitepropVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sitepropVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sitepropVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SiteInfo) contextValidateTenantsitepropVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TenantsitepropVec); i++ {

		if m.TenantsitepropVec[i] != nil {

			if swag.IsZero(m.TenantsitepropVec[i]) { // not required
				return nil
			}

			if err := m.TenantsitepropVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenantsitepropVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenantsitepropVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SiteInfo) contextValidateWebpropVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WebpropVec); i++ {

		if m.WebpropVec[i] != nil {

			if swag.IsZero(m.WebpropVec[i]) { // not required
				return nil
			}

			if err := m.WebpropVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webpropVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webpropVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteInfo) UnmarshalBinary(b []byte) error {
	var res SiteInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
