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

// RecoverO365Params Recover Office 365 environment params.
//
// Specifies the recovery options specific to Office 365 environment.
//
// swagger:model RecoverO365Params
type RecoverO365Params struct {

	// Specifies the list of recover Object parameters.
	Objects []*CommonRecoverObjectSnapshotParams `json:"objects"`

	// Specifies the type of recovery action to be performed.
	// Required: true
	// Enum: ["RecoverMailbox","RecoverOneDrive","RecoverSharePoint","RecoverPublicFolders","RecoverMsGroup","RecoverMsTeam","ConvertToPst","DownloadChats","RecoverMailboxCSM","RecoverOneDriveCSM","RecoverSharePointCSM"]
	RecoveryAction *string `json:"recoveryAction"`

	// Specifies the download chats specific parameters for downloading posts for a team/channel or downloading private chats for a user.
	DownloadChatsParams *DownloadChatsParams `json:"downloadChatsParams,omitempty"`

	// Specifies the recovery information to download files and folders. For instance, downloading mailbox items as PST.
	DownloadFileAndFolderParams *CommonDownloadFileAndFolderParams `json:"downloadFileAndFolderParams,omitempty"`

	// Specifies the parameters to recover Office 365 Mailbox.
	RecoverMailboxParams *RecoverMailboxParams `json:"recoverMailboxParams,omitempty"`

	// Specifies the parameters to recover Microsoft 365 Group.
	RecoverMsGroupParams *RecoverMsGroupParams `json:"recoverMsGroupParams,omitempty"`

	// Specifies the parameters to recover Microsoft 365 Teams.
	RecoverMsTeamParams *RecoverMsTeamParams `json:"recoverMsTeamParams,omitempty"`

	// Specifies the parameters to recover Office 365 One Drive.
	RecoverOneDriveParams *RecoverOneDriveParams `json:"recoverOneDriveParams,omitempty"`

	// Specifies the parameters to recover Office 365 Public Folders.
	RecoverPublicFoldersParams *RecoverPublicFoldersParams `json:"recoverPublicFoldersParams,omitempty"`

	// Specifies the parameters to recover Microsoft Office 365 Sharepoint Site.
	RecoverSiteParams *RecoverSiteParams `json:"recoverSiteParams,omitempty"`
}

// Validate validates this recover o365 params
func (m *RecoverO365Params) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadChatsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadFileAndFolderParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverMailboxParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverMsGroupParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverMsTeamParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverOneDriveParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverPublicFoldersParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoverSiteParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverO365Params) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var recoverO365ParamsTypeRecoveryActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RecoverMailbox","RecoverOneDrive","RecoverSharePoint","RecoverPublicFolders","RecoverMsGroup","RecoverMsTeam","ConvertToPst","DownloadChats","RecoverMailboxCSM","RecoverOneDriveCSM","RecoverSharePointCSM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverO365ParamsTypeRecoveryActionPropEnum = append(recoverO365ParamsTypeRecoveryActionPropEnum, v)
	}
}

const (

	// RecoverO365ParamsRecoveryActionRecoverMailbox captures enum value "RecoverMailbox"
	RecoverO365ParamsRecoveryActionRecoverMailbox string = "RecoverMailbox"

	// RecoverO365ParamsRecoveryActionRecoverOneDrive captures enum value "RecoverOneDrive"
	RecoverO365ParamsRecoveryActionRecoverOneDrive string = "RecoverOneDrive"

	// RecoverO365ParamsRecoveryActionRecoverSharePoint captures enum value "RecoverSharePoint"
	RecoverO365ParamsRecoveryActionRecoverSharePoint string = "RecoverSharePoint"

	// RecoverO365ParamsRecoveryActionRecoverPublicFolders captures enum value "RecoverPublicFolders"
	RecoverO365ParamsRecoveryActionRecoverPublicFolders string = "RecoverPublicFolders"

	// RecoverO365ParamsRecoveryActionRecoverMsGroup captures enum value "RecoverMsGroup"
	RecoverO365ParamsRecoveryActionRecoverMsGroup string = "RecoverMsGroup"

	// RecoverO365ParamsRecoveryActionRecoverMsTeam captures enum value "RecoverMsTeam"
	RecoverO365ParamsRecoveryActionRecoverMsTeam string = "RecoverMsTeam"

	// RecoverO365ParamsRecoveryActionConvertToPst captures enum value "ConvertToPst"
	RecoverO365ParamsRecoveryActionConvertToPst string = "ConvertToPst"

	// RecoverO365ParamsRecoveryActionDownloadChats captures enum value "DownloadChats"
	RecoverO365ParamsRecoveryActionDownloadChats string = "DownloadChats"

	// RecoverO365ParamsRecoveryActionRecoverMailboxCSM captures enum value "RecoverMailboxCSM"
	RecoverO365ParamsRecoveryActionRecoverMailboxCSM string = "RecoverMailboxCSM"

	// RecoverO365ParamsRecoveryActionRecoverOneDriveCSM captures enum value "RecoverOneDriveCSM"
	RecoverO365ParamsRecoveryActionRecoverOneDriveCSM string = "RecoverOneDriveCSM"

	// RecoverO365ParamsRecoveryActionRecoverSharePointCSM captures enum value "RecoverSharePointCSM"
	RecoverO365ParamsRecoveryActionRecoverSharePointCSM string = "RecoverSharePointCSM"
)

// prop value enum
func (m *RecoverO365Params) validateRecoveryActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverO365ParamsTypeRecoveryActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverO365Params) validateRecoveryAction(formats strfmt.Registry) error {

	if err := validate.Required("recoveryAction", "body", m.RecoveryAction); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecoveryActionEnum("recoveryAction", "body", *m.RecoveryAction); err != nil {
		return err
	}

	return nil
}

func (m *RecoverO365Params) validateDownloadChatsParams(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadChatsParams) { // not required
		return nil
	}

	if m.DownloadChatsParams != nil {
		if err := m.DownloadChatsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadChatsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadChatsParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) validateDownloadFileAndFolderParams(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadFileAndFolderParams) { // not required
		return nil
	}

	if m.DownloadFileAndFolderParams != nil {
		if err := m.DownloadFileAndFolderParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadFileAndFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadFileAndFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) validateRecoverMailboxParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverMailboxParams) { // not required
		return nil
	}

	if m.RecoverMailboxParams != nil {
		if err := m.RecoverMailboxParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverMailboxParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverMailboxParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) validateRecoverMsGroupParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverMsGroupParams) { // not required
		return nil
	}

	if m.RecoverMsGroupParams != nil {
		if err := m.RecoverMsGroupParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverMsGroupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverMsGroupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) validateRecoverMsTeamParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverMsTeamParams) { // not required
		return nil
	}

	if m.RecoverMsTeamParams != nil {
		if err := m.RecoverMsTeamParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverMsTeamParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverMsTeamParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) validateRecoverOneDriveParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverOneDriveParams) { // not required
		return nil
	}

	if m.RecoverOneDriveParams != nil {
		if err := m.RecoverOneDriveParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverOneDriveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverOneDriveParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) validateRecoverPublicFoldersParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverPublicFoldersParams) { // not required
		return nil
	}

	if m.RecoverPublicFoldersParams != nil {
		if err := m.RecoverPublicFoldersParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverPublicFoldersParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverPublicFoldersParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) validateRecoverSiteParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverSiteParams) { // not required
		return nil
	}

	if m.RecoverSiteParams != nil {
		if err := m.RecoverSiteParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverSiteParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverSiteParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover o365 params based on the context it is used
func (m *RecoverO365Params) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDownloadChatsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDownloadFileAndFolderParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverMailboxParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverMsGroupParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverMsTeamParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverOneDriveParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverPublicFoldersParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoverSiteParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverO365Params) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverO365Params) contextValidateDownloadChatsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.DownloadChatsParams != nil {

		if swag.IsZero(m.DownloadChatsParams) { // not required
			return nil
		}

		if err := m.DownloadChatsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadChatsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadChatsParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) contextValidateDownloadFileAndFolderParams(ctx context.Context, formats strfmt.Registry) error {

	if m.DownloadFileAndFolderParams != nil {

		if swag.IsZero(m.DownloadFileAndFolderParams) { // not required
			return nil
		}

		if err := m.DownloadFileAndFolderParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadFileAndFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadFileAndFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) contextValidateRecoverMailboxParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverMailboxParams != nil {

		if swag.IsZero(m.RecoverMailboxParams) { // not required
			return nil
		}

		if err := m.RecoverMailboxParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverMailboxParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverMailboxParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) contextValidateRecoverMsGroupParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverMsGroupParams != nil {

		if swag.IsZero(m.RecoverMsGroupParams) { // not required
			return nil
		}

		if err := m.RecoverMsGroupParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverMsGroupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverMsGroupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) contextValidateRecoverMsTeamParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverMsTeamParams != nil {

		if swag.IsZero(m.RecoverMsTeamParams) { // not required
			return nil
		}

		if err := m.RecoverMsTeamParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverMsTeamParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverMsTeamParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) contextValidateRecoverOneDriveParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverOneDriveParams != nil {

		if swag.IsZero(m.RecoverOneDriveParams) { // not required
			return nil
		}

		if err := m.RecoverOneDriveParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverOneDriveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverOneDriveParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) contextValidateRecoverPublicFoldersParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverPublicFoldersParams != nil {

		if swag.IsZero(m.RecoverPublicFoldersParams) { // not required
			return nil
		}

		if err := m.RecoverPublicFoldersParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverPublicFoldersParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverPublicFoldersParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverO365Params) contextValidateRecoverSiteParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverSiteParams != nil {

		if swag.IsZero(m.RecoverSiteParams) { // not required
			return nil
		}

		if err := m.RecoverSiteParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoverSiteParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoverSiteParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverO365Params) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverO365Params) UnmarshalBinary(b []byte) error {
	var res RecoverO365Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
