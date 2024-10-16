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

// SearchIndexedObjectsResponseBody Search Indexed objects response body.
//
// Specifies the search indexed objects response body.
//
// swagger:model SearchIndexedObjectsResponseBody
type SearchIndexedObjectsResponseBody struct {
	CommonSearchIndexedObjectsResponseParams

	// emails
	Emails Emails `json:"emails,omitempty"`

	// files
	Files Files `json:"files,omitempty"`

	// cassandra objects
	CassandraObjects CassandraIndexedObjects `json:"cassandraObjects,omitempty"`

	// couchbase objects
	CouchbaseObjects CouchbaseIndexedObjects `json:"couchbaseObjects,omitempty"`

	// hbase objects
	HbaseObjects HbaseIndexedObjects `json:"hbaseObjects,omitempty"`

	// hive objects
	HiveObjects HiveIndexedObjects `json:"hiveObjects,omitempty"`

	// mongo objects
	MongoObjects MongoIndexedObjects `json:"mongoObjects,omitempty"`

	// hdfs objects
	HdfsObjects HDFSIndexedObjects `json:"hdfsObjects,omitempty"`

	// exchange objects
	ExchangeObjects ExchangeIndexedObjects `json:"exchangeObjects,omitempty"`

	// public folder items
	PublicFolderItems PublicFolderItems `json:"publicFolderItems,omitempty"`

	// sharepoint items
	SharepointItems SharepointItems `json:"sharepointItems,omitempty"`

	// one drive items
	OneDriveItems OneDriveItems `json:"oneDriveItems,omitempty"`

	// uda objects
	UdaObjects UdaIndexedObjects `json:"udaObjects,omitempty"`

	// teams items
	TeamsItems TeamsItems `json:"teamsItems,omitempty"`

	// sfdc records
	SfdcRecords *SfdcRecords `json:"sfdcRecords,omitempty"`

	// ms group items
	MsGroupItems MsGroupItems `json:"msGroupItems,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SearchIndexedObjectsResponseBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonSearchIndexedObjectsResponseParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonSearchIndexedObjectsResponseParams = aO0

	// AO1
	var dataAO1 struct {
		Emails Emails `json:"emails,omitempty"`

		Files Files `json:"files,omitempty"`

		CassandraObjects CassandraIndexedObjects `json:"cassandraObjects,omitempty"`

		CouchbaseObjects CouchbaseIndexedObjects `json:"couchbaseObjects,omitempty"`

		HbaseObjects HbaseIndexedObjects `json:"hbaseObjects,omitempty"`

		HiveObjects HiveIndexedObjects `json:"hiveObjects,omitempty"`

		MongoObjects MongoIndexedObjects `json:"mongoObjects,omitempty"`

		HdfsObjects HDFSIndexedObjects `json:"hdfsObjects,omitempty"`

		ExchangeObjects ExchangeIndexedObjects `json:"exchangeObjects,omitempty"`

		PublicFolderItems PublicFolderItems `json:"publicFolderItems,omitempty"`

		SharepointItems SharepointItems `json:"sharepointItems,omitempty"`

		OneDriveItems OneDriveItems `json:"oneDriveItems,omitempty"`

		UdaObjects UdaIndexedObjects `json:"udaObjects,omitempty"`

		TeamsItems TeamsItems `json:"teamsItems,omitempty"`

		SfdcRecords *SfdcRecords `json:"sfdcRecords,omitempty"`

		MsGroupItems MsGroupItems `json:"msGroupItems,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Emails = dataAO1.Emails

	m.Files = dataAO1.Files

	m.CassandraObjects = dataAO1.CassandraObjects

	m.CouchbaseObjects = dataAO1.CouchbaseObjects

	m.HbaseObjects = dataAO1.HbaseObjects

	m.HiveObjects = dataAO1.HiveObjects

	m.MongoObjects = dataAO1.MongoObjects

	m.HdfsObjects = dataAO1.HdfsObjects

	m.ExchangeObjects = dataAO1.ExchangeObjects

	m.PublicFolderItems = dataAO1.PublicFolderItems

	m.SharepointItems = dataAO1.SharepointItems

	m.OneDriveItems = dataAO1.OneDriveItems

	m.UdaObjects = dataAO1.UdaObjects

	m.TeamsItems = dataAO1.TeamsItems

	m.SfdcRecords = dataAO1.SfdcRecords

	m.MsGroupItems = dataAO1.MsGroupItems

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SearchIndexedObjectsResponseBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonSearchIndexedObjectsResponseParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Emails Emails `json:"emails,omitempty"`

		Files Files `json:"files,omitempty"`

		CassandraObjects CassandraIndexedObjects `json:"cassandraObjects,omitempty"`

		CouchbaseObjects CouchbaseIndexedObjects `json:"couchbaseObjects,omitempty"`

		HbaseObjects HbaseIndexedObjects `json:"hbaseObjects,omitempty"`

		HiveObjects HiveIndexedObjects `json:"hiveObjects,omitempty"`

		MongoObjects MongoIndexedObjects `json:"mongoObjects,omitempty"`

		HdfsObjects HDFSIndexedObjects `json:"hdfsObjects,omitempty"`

		ExchangeObjects ExchangeIndexedObjects `json:"exchangeObjects,omitempty"`

		PublicFolderItems PublicFolderItems `json:"publicFolderItems,omitempty"`

		SharepointItems SharepointItems `json:"sharepointItems,omitempty"`

		OneDriveItems OneDriveItems `json:"oneDriveItems,omitempty"`

		UdaObjects UdaIndexedObjects `json:"udaObjects,omitempty"`

		TeamsItems TeamsItems `json:"teamsItems,omitempty"`

		SfdcRecords *SfdcRecords `json:"sfdcRecords,omitempty"`

		MsGroupItems MsGroupItems `json:"msGroupItems,omitempty"`
	}

	dataAO1.Emails = m.Emails

	dataAO1.Files = m.Files

	dataAO1.CassandraObjects = m.CassandraObjects

	dataAO1.CouchbaseObjects = m.CouchbaseObjects

	dataAO1.HbaseObjects = m.HbaseObjects

	dataAO1.HiveObjects = m.HiveObjects

	dataAO1.MongoObjects = m.MongoObjects

	dataAO1.HdfsObjects = m.HdfsObjects

	dataAO1.ExchangeObjects = m.ExchangeObjects

	dataAO1.PublicFolderItems = m.PublicFolderItems

	dataAO1.SharepointItems = m.SharepointItems

	dataAO1.OneDriveItems = m.OneDriveItems

	dataAO1.UdaObjects = m.UdaObjects

	dataAO1.TeamsItems = m.TeamsItems

	dataAO1.SfdcRecords = m.SfdcRecords

	dataAO1.MsGroupItems = m.MsGroupItems

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this search indexed objects response body
func (m *SearchIndexedObjectsResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSearchIndexedObjectsResponseParams
	if err := m.CommonSearchIndexedObjectsResponseParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCassandraObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCouchbaseObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHbaseObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHiveObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongoObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdfsObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicFolderItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharepointItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneDriveItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUdaObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamsItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSfdcRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsGroupItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateEmails(formats strfmt.Registry) error {

	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	if err := m.Emails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("emails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("emails")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Files) { // not required
		return nil
	}

	if err := m.Files.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("files")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("files")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateCassandraObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.CassandraObjects) { // not required
		return nil
	}

	if err := m.CassandraObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cassandraObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cassandraObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateCouchbaseObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.CouchbaseObjects) { // not required
		return nil
	}

	if err := m.CouchbaseObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("couchbaseObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("couchbaseObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateHbaseObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.HbaseObjects) { // not required
		return nil
	}

	if err := m.HbaseObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hbaseObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hbaseObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateHiveObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.HiveObjects) { // not required
		return nil
	}

	if err := m.HiveObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hiveObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hiveObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateMongoObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.MongoObjects) { // not required
		return nil
	}

	if err := m.MongoObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mongoObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mongoObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateHdfsObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.HdfsObjects) { // not required
		return nil
	}

	if err := m.HdfsObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hdfsObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hdfsObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateExchangeObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.ExchangeObjects) { // not required
		return nil
	}

	if err := m.ExchangeObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("exchangeObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("exchangeObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validatePublicFolderItems(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicFolderItems) { // not required
		return nil
	}

	if err := m.PublicFolderItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("publicFolderItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("publicFolderItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateSharepointItems(formats strfmt.Registry) error {

	if swag.IsZero(m.SharepointItems) { // not required
		return nil
	}

	if err := m.SharepointItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sharepointItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sharepointItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateOneDriveItems(formats strfmt.Registry) error {

	if swag.IsZero(m.OneDriveItems) { // not required
		return nil
	}

	if err := m.OneDriveItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("oneDriveItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("oneDriveItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateUdaObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.UdaObjects) { // not required
		return nil
	}

	if err := m.UdaObjects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("udaObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("udaObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateTeamsItems(formats strfmt.Registry) error {

	if swag.IsZero(m.TeamsItems) { // not required
		return nil
	}

	if err := m.TeamsItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("teamsItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("teamsItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateSfdcRecords(formats strfmt.Registry) error {

	if swag.IsZero(m.SfdcRecords) { // not required
		return nil
	}

	if m.SfdcRecords != nil {
		if err := m.SfdcRecords.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfdcRecords")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sfdcRecords")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) validateMsGroupItems(formats strfmt.Registry) error {

	if swag.IsZero(m.MsGroupItems) { // not required
		return nil
	}

	if err := m.MsGroupItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("msGroupItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("msGroupItems")
		}
		return err
	}

	return nil
}

// ContextValidate validate this search indexed objects response body based on the context it is used
func (m *SearchIndexedObjectsResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSearchIndexedObjectsResponseParams
	if err := m.CommonSearchIndexedObjectsResponseParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCassandraObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCouchbaseObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHbaseObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHiveObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMongoObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHdfsObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExchangeObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicFolderItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSharepointItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOneDriveItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUdaObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeamsItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSfdcRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsGroupItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Emails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("emails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("emails")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Files.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("files")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("files")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateCassandraObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CassandraObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cassandraObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cassandraObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateCouchbaseObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CouchbaseObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("couchbaseObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("couchbaseObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateHbaseObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HbaseObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hbaseObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hbaseObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateHiveObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HiveObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hiveObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hiveObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateMongoObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MongoObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mongoObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("mongoObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateHdfsObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HdfsObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hdfsObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hdfsObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateExchangeObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ExchangeObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("exchangeObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("exchangeObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidatePublicFolderItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PublicFolderItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("publicFolderItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("publicFolderItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateSharepointItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SharepointItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sharepointItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sharepointItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateOneDriveItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OneDriveItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("oneDriveItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("oneDriveItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateUdaObjects(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UdaObjects.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("udaObjects")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("udaObjects")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateTeamsItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TeamsItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("teamsItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("teamsItems")
		}
		return err
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateSfdcRecords(ctx context.Context, formats strfmt.Registry) error {

	if m.SfdcRecords != nil {

		if swag.IsZero(m.SfdcRecords) { // not required
			return nil
		}

		if err := m.SfdcRecords.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfdcRecords")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sfdcRecords")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsResponseBody) contextValidateMsGroupItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MsGroupItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("msGroupItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("msGroupItems")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchIndexedObjectsResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchIndexedObjectsResponseBody) UnmarshalBinary(b []byte) error {
	var res SearchIndexedObjectsResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
