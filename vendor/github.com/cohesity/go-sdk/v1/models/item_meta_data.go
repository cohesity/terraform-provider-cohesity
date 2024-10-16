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

// ItemMetaData Message encapsulating the metadata for an outlook item. This has data which
// can be indexed by yoda for searching individual items.
//
// swagger:model ItemMetaData
type ItemMetaData struct {

	// List of attachment ids.
	AttachmentVec []*Attachment `json:"attachmentVec"`

	// List of BccRecipients.
	BccRecipientVec []*PrivateUser `json:"bccRecipientVec"`

	// Birthday of the Contact.
	Birthday *int64 `json:"birthday,omitempty"`

	// The first 255 characters of the message body. It is in text format.
	BodyPreview *string `json:"bodyPreview,omitempty"`

	// Graph specific message fields.
	// The categories associated with the message.
	Categories []string `json:"categories"`

	// List of CcRecipients.
	CcRecipientVec []*PrivateUser `json:"ccRecipientVec"`

	// Version of the item that changed.
	ChangeKey *string `json:"changeKey,omitempty"`

	// Task completion date.
	CompletionDate *int64 `json:"completionDate,omitempty"`

	// The type of the content. Possible values are text and html.
	ContentType *string `json:"contentType,omitempty"`

	// The ID of the conversation the message belongs to.
	ConversationID *string `json:"conversationId,omitempty"`

	// Indicates the position of the message within the conversation.
	ConversationIndex *string `json:"conversationIndex,omitempty"`

	// Additional common fields for all item types.
	// Time at which the item is created - DateTimeCreated.
	CreationTime *int64 `json:"creationTime,omitempty"`

	// Task due date.
	DueDate *int64 `json:"dueDate,omitempty"`

	// Email addresses assigned to the Contact.
	EmailAddressesVec []string `json:"emailAddressesVec"`

	// End date/time for the single occurence events.
	EndTime *int64 `json:"endTime,omitempty"`

	// Contact items specific fields.
	// First name of the Contact.
	FirstName *string `json:"firstName,omitempty"`

	// Date/time for the first occurence of the recurring events.
	FirstOccurrence *ItemMetaDataOccurrence `json:"firstOccurrence,omitempty"`

	// The flag value that indicates the status, start date, due date,
	// or completion date for the message.
	Flag *ItemMetaDataFlag `json:"flag,omitempty"`

	// From whom the mail is received.
	From *PrivateUser `json:"from,omitempty"`

	// The graph/rest id of the item.
	GraphID *string `json:"graphId,omitempty"`

	// Does the item have attachments?
	HasAttachments *bool `json:"hasAttachments,omitempty"`

	// Id of the item.
	// If this is a top level item, then we should always have EWS ID populated
	// in the id field, irrespective of the backup method used.
	// For MSGraph API based backups, populate the items's rest/Graph ID in
	// graph_id field.
	// If this item represents an attachment of the another item, then the id
	// field shouldn't be present and is_attachment_item should be set to true.
	ID *string `json:"id,omitempty"`

	// Item importance.
	Importance *int32 `json:"importance,omitempty"`

	// The classification of the message for the user, based on inferred
	// relevance or importance, or on an explicit override.
	// The possible values are: focused or other.
	InferenceClassification *string `json:"inferenceClassification,omitempty"`

	// A collection of message headers defined by RFC5322.
	InternetMessageHeaders []*ItemMetaDataInternetMessageHeader `json:"internetMessageHeaders"`

	// The message ID in the format specified by RFC2822.
	InternetMessageID *string `json:"internetMessageId,omitempty"`

	// This item is an item-attachment or a top level item.
	IsAttachmentItem *bool `json:"isAttachmentItem,omitempty"`

	// Indicates whether a read receipt is requested for the message.
	IsDeliveryReceiptRequested *bool `json:"isDeliveryReceiptRequested,omitempty"`

	// Indicates if mail in draft or not.
	IsDraft *bool `json:"isDraft,omitempty"`

	// Indicates if mail is marked as read or unread.
	IsRead *bool `json:"isRead,omitempty"`

	// Indicates whether a read receipt is requested for the message.
	IsReadReceiptRequested *bool `json:"isReadReceiptRequested,omitempty"`

	// Is this recurring calendar event or task.
	IsRecurring *bool `json:"isRecurring,omitempty"`

	// Class of the item.
	ItemClass *string `json:"itemClass,omitempty"`

	// The item was last modified by.
	LastModifiedName *string `json:"lastModifiedName,omitempty"`

	// Last modificiation time of the item.
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty"`

	// Last name of the Contact.
	LastName *string `json:"lastName,omitempty"`

	// Date/time for the last occurence of the recurring events.
	LastOccurrence *ItemMetaDataOccurrence `json:"lastOccurrence,omitempty"`

	// Parent folder id returned with GetItem API call of the item.
	MetadataParentFolderID *string `json:"metadataParentFolderId,omitempty"`

	// Middle name of the Contact.
	MiddleName *string `json:"middleName,omitempty"`

	// List of multi value extended properties.
	MultiValueExtendedProperties []*ItemMetaDataMultiValueExtendedProperty `json:"multiValueExtendedProperties"`

	// ETag indicating change in content.
	// This is a odata property that can be potentially used to detect change in
	// the message. Currently we use folder changes to detect which messages have
	// changes that require backing up, but this property can also be used to
	// detect changes.
	OdataEtag *string `json:"odataEtag,omitempty"`

	// Optional Attendee list - OptionalAttendees.
	OptionalAttendeesVec []*PrivateUser `json:"optionalAttendeesVec"`

	// Calendar item spcific fields.
	// Appointment/meeting Organizer.
	Organizer *PrivateUser `json:"organizer,omitempty"`

	// Task items specific fields.
	// Task owner.
	Owner *string `json:"owner,omitempty"`

	// Parent folder id of the item.
	ParentFolderID *string `json:"parentFolderId,omitempty"`

	// Time at which the mail is received.
	ReceivedTime *int64 `json:"receivedTime,omitempty"`

	// Recurrence pattern information for the recurring calendar events and
	// tasks.
	RecurrencePattern *int32 `json:"recurrencePattern,omitempty"`

	// The email addresses to use when replying.
	ReplyToVec []*PrivateUser `json:"replyToVec"`

	// Attendee list - RequiredAttendees.
	RequiredAttendeesVec []*PrivateUser `json:"requiredAttendeesVec"`

	// The account that is actually used to generate the message.
	// In most cases, this value is the same as the from property, except for
	// sharing or delegation scenarios.
	Sender *PrivateUser `json:"sender,omitempty"`

	// Item sensitivity.
	Sensitivity *int32 `json:"sensitivity,omitempty"`

	// Time at which the mail is sent.
	SentTime *int64 `json:"sentTime,omitempty"`

	// Sha-256 for the data.
	Sha256Checksum *string `json:"sha_256Checksum,omitempty"`

	// List of single value extended properties. These properties are used by
	// microsoft internally to set the corresponding properties.
	// Examples include creating message in a non-draft format, to set sent and
	// received date and time for the message and much more.
	SingleValueExtendedProperties []*ItemMetaDataSingleValueExtendedProperty `json:"singleValueExtendedProperties"`

	// Size of the item.
	Size *int64 `json:"size,omitempty"`

	// Start date/time for the single occurence events.
	StartTime *int64 `json:"startTime,omitempty"`

	// Subject of the item.
	Subject *string `json:"subject,omitempty"`

	// Status of the task.
	TaskStatus *int32 `json:"taskStatus,omitempty"`

	// List of ToRecipients.
	ToRecipientVec []*PrivateUser `json:"toRecipientVec"`

	// Type of the item.
	Type *int32 `json:"type,omitempty"`

	// The URL to open the message in Outlook on the web.
	Weblink *string `json:"weblink,omitempty"`
}

// Validate validates this item meta data
func (m *ItemMetaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBccRecipientVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCcRecipientVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstOccurrence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternetMessageHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastOccurrence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultiValueExtendedProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionalAttendeesVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplyToVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredAttendeesVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingleValueExtendedProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToRecipientVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemMetaData) validateAttachmentVec(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachmentVec) { // not required
		return nil
	}

	for i := 0; i < len(m.AttachmentVec); i++ {
		if swag.IsZero(m.AttachmentVec[i]) { // not required
			continue
		}

		if m.AttachmentVec[i] != nil {
			if err := m.AttachmentVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachmentVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachmentVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateBccRecipientVec(formats strfmt.Registry) error {
	if swag.IsZero(m.BccRecipientVec) { // not required
		return nil
	}

	for i := 0; i < len(m.BccRecipientVec); i++ {
		if swag.IsZero(m.BccRecipientVec[i]) { // not required
			continue
		}

		if m.BccRecipientVec[i] != nil {
			if err := m.BccRecipientVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bccRecipientVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bccRecipientVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateCcRecipientVec(formats strfmt.Registry) error {
	if swag.IsZero(m.CcRecipientVec) { // not required
		return nil
	}

	for i := 0; i < len(m.CcRecipientVec); i++ {
		if swag.IsZero(m.CcRecipientVec[i]) { // not required
			continue
		}

		if m.CcRecipientVec[i] != nil {
			if err := m.CcRecipientVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ccRecipientVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ccRecipientVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateFirstOccurrence(formats strfmt.Registry) error {
	if swag.IsZero(m.FirstOccurrence) { // not required
		return nil
	}

	if m.FirstOccurrence != nil {
		if err := m.FirstOccurrence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firstOccurrence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firstOccurrence")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) validateFlag(formats strfmt.Registry) error {
	if swag.IsZero(m.Flag) { // not required
		return nil
	}

	if m.Flag != nil {
		if err := m.Flag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flag")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) validateInternetMessageHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.InternetMessageHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.InternetMessageHeaders); i++ {
		if swag.IsZero(m.InternetMessageHeaders[i]) { // not required
			continue
		}

		if m.InternetMessageHeaders[i] != nil {
			if err := m.InternetMessageHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("internetMessageHeaders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("internetMessageHeaders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateLastOccurrence(formats strfmt.Registry) error {
	if swag.IsZero(m.LastOccurrence) { // not required
		return nil
	}

	if m.LastOccurrence != nil {
		if err := m.LastOccurrence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastOccurrence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastOccurrence")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) validateMultiValueExtendedProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.MultiValueExtendedProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.MultiValueExtendedProperties); i++ {
		if swag.IsZero(m.MultiValueExtendedProperties[i]) { // not required
			continue
		}

		if m.MultiValueExtendedProperties[i] != nil {
			if err := m.MultiValueExtendedProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("multiValueExtendedProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("multiValueExtendedProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateOptionalAttendeesVec(formats strfmt.Registry) error {
	if swag.IsZero(m.OptionalAttendeesVec) { // not required
		return nil
	}

	for i := 0; i < len(m.OptionalAttendeesVec); i++ {
		if swag.IsZero(m.OptionalAttendeesVec[i]) { // not required
			continue
		}

		if m.OptionalAttendeesVec[i] != nil {
			if err := m.OptionalAttendeesVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("optionalAttendeesVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("optionalAttendeesVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateOrganizer(formats strfmt.Registry) error {
	if swag.IsZero(m.Organizer) { // not required
		return nil
	}

	if m.Organizer != nil {
		if err := m.Organizer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organizer")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) validateReplyToVec(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplyToVec) { // not required
		return nil
	}

	for i := 0; i < len(m.ReplyToVec); i++ {
		if swag.IsZero(m.ReplyToVec[i]) { // not required
			continue
		}

		if m.ReplyToVec[i] != nil {
			if err := m.ReplyToVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replyToVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replyToVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateRequiredAttendeesVec(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredAttendeesVec) { // not required
		return nil
	}

	for i := 0; i < len(m.RequiredAttendeesVec); i++ {
		if swag.IsZero(m.RequiredAttendeesVec[i]) { // not required
			continue
		}

		if m.RequiredAttendeesVec[i] != nil {
			if err := m.RequiredAttendeesVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requiredAttendeesVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requiredAttendeesVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateSender(formats strfmt.Registry) error {
	if swag.IsZero(m.Sender) { // not required
		return nil
	}

	if m.Sender != nil {
		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) validateSingleValueExtendedProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.SingleValueExtendedProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.SingleValueExtendedProperties); i++ {
		if swag.IsZero(m.SingleValueExtendedProperties[i]) { // not required
			continue
		}

		if m.SingleValueExtendedProperties[i] != nil {
			if err := m.SingleValueExtendedProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("singleValueExtendedProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("singleValueExtendedProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) validateToRecipientVec(formats strfmt.Registry) error {
	if swag.IsZero(m.ToRecipientVec) { // not required
		return nil
	}

	for i := 0; i < len(m.ToRecipientVec); i++ {
		if swag.IsZero(m.ToRecipientVec[i]) { // not required
			continue
		}

		if m.ToRecipientVec[i] != nil {
			if err := m.ToRecipientVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("toRecipientVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("toRecipientVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this item meta data based on the context it is used
func (m *ItemMetaData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachmentVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBccRecipientVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCcRecipientVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstOccurrence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInternetMessageHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastOccurrence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMultiValueExtendedProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptionalAttendeesVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganizer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplyToVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredAttendeesVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSingleValueExtendedProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToRecipientVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemMetaData) contextValidateAttachmentVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttachmentVec); i++ {

		if m.AttachmentVec[i] != nil {

			if swag.IsZero(m.AttachmentVec[i]) { // not required
				return nil
			}

			if err := m.AttachmentVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachmentVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachmentVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateBccRecipientVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BccRecipientVec); i++ {

		if m.BccRecipientVec[i] != nil {

			if swag.IsZero(m.BccRecipientVec[i]) { // not required
				return nil
			}

			if err := m.BccRecipientVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bccRecipientVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bccRecipientVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateCcRecipientVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CcRecipientVec); i++ {

		if m.CcRecipientVec[i] != nil {

			if swag.IsZero(m.CcRecipientVec[i]) { // not required
				return nil
			}

			if err := m.CcRecipientVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ccRecipientVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ccRecipientVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateFirstOccurrence(ctx context.Context, formats strfmt.Registry) error {

	if m.FirstOccurrence != nil {

		if swag.IsZero(m.FirstOccurrence) { // not required
			return nil
		}

		if err := m.FirstOccurrence.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firstOccurrence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firstOccurrence")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) contextValidateFlag(ctx context.Context, formats strfmt.Registry) error {

	if m.Flag != nil {

		if swag.IsZero(m.Flag) { // not required
			return nil
		}

		if err := m.Flag.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flag")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flag")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.From != nil {

		if swag.IsZero(m.From) { // not required
			return nil
		}

		if err := m.From.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) contextValidateInternetMessageHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InternetMessageHeaders); i++ {

		if m.InternetMessageHeaders[i] != nil {

			if swag.IsZero(m.InternetMessageHeaders[i]) { // not required
				return nil
			}

			if err := m.InternetMessageHeaders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("internetMessageHeaders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("internetMessageHeaders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateLastOccurrence(ctx context.Context, formats strfmt.Registry) error {

	if m.LastOccurrence != nil {

		if swag.IsZero(m.LastOccurrence) { // not required
			return nil
		}

		if err := m.LastOccurrence.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastOccurrence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastOccurrence")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) contextValidateMultiValueExtendedProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MultiValueExtendedProperties); i++ {

		if m.MultiValueExtendedProperties[i] != nil {

			if swag.IsZero(m.MultiValueExtendedProperties[i]) { // not required
				return nil
			}

			if err := m.MultiValueExtendedProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("multiValueExtendedProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("multiValueExtendedProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateOptionalAttendeesVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OptionalAttendeesVec); i++ {

		if m.OptionalAttendeesVec[i] != nil {

			if swag.IsZero(m.OptionalAttendeesVec[i]) { // not required
				return nil
			}

			if err := m.OptionalAttendeesVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("optionalAttendeesVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("optionalAttendeesVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateOrganizer(ctx context.Context, formats strfmt.Registry) error {

	if m.Organizer != nil {

		if swag.IsZero(m.Organizer) { // not required
			return nil
		}

		if err := m.Organizer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organizer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organizer")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) contextValidateReplyToVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReplyToVec); i++ {

		if m.ReplyToVec[i] != nil {

			if swag.IsZero(m.ReplyToVec[i]) { // not required
				return nil
			}

			if err := m.ReplyToVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replyToVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replyToVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateRequiredAttendeesVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequiredAttendeesVec); i++ {

		if m.RequiredAttendeesVec[i] != nil {

			if swag.IsZero(m.RequiredAttendeesVec[i]) { // not required
				return nil
			}

			if err := m.RequiredAttendeesVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requiredAttendeesVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requiredAttendeesVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateSender(ctx context.Context, formats strfmt.Registry) error {

	if m.Sender != nil {

		if swag.IsZero(m.Sender) { // not required
			return nil
		}

		if err := m.Sender.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

func (m *ItemMetaData) contextValidateSingleValueExtendedProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SingleValueExtendedProperties); i++ {

		if m.SingleValueExtendedProperties[i] != nil {

			if swag.IsZero(m.SingleValueExtendedProperties[i]) { // not required
				return nil
			}

			if err := m.SingleValueExtendedProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("singleValueExtendedProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("singleValueExtendedProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ItemMetaData) contextValidateToRecipientVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ToRecipientVec); i++ {

		if m.ToRecipientVec[i] != nil {

			if swag.IsZero(m.ToRecipientVec[i]) { // not required
				return nil
			}

			if err := m.ToRecipientVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("toRecipientVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("toRecipientVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemMetaData) UnmarshalBinary(b []byte) error {
	var res ItemMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
