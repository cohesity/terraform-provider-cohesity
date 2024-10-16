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

// SearchIndexedObjectsRequest Search indexed objects request params.
//
// Specifies the request parameters to search for indexed objects.
//
// swagger:model SearchIndexedObjectsRequest
type SearchIndexedObjectsRequest struct {
	CommonSearchIndexedObjectsRequestParams

	// Specifies the parameters which are specific for searching emails and email folders.
	EmailParams *SearchEmailRequestParams `json:"emailParams,omitempty"`

	// Specifies the parameters which are specific for searching files and file folders.
	FileParams *SearchFileRequestParams `json:"fileParams,omitempty"`

	// cassandra params
	CassandraParams *CassandraOnPremSearchParams `json:"cassandraParams,omitempty"`

	// couchbase params
	CouchbaseParams *CouchBaseOnPremSearchParams `json:"couchbaseParams,omitempty"`

	// hbase params
	HbaseParams *HbaseOnPremSearchParams `json:"hbaseParams,omitempty"`

	// hive params
	HiveParams *HiveOnPremSearchParams `json:"hiveParams,omitempty"`

	// mongodb params
	MongodbParams *MongoDbOnPremSearchParams `json:"mongodbParams,omitempty"`

	// hdfs params
	HdfsParams *HDFSOnPremSearchParams `json:"hdfsParams,omitempty"`

	// Specifies the parameters which are specific for searching Exchange mailboxes.
	ExchangeParams *SearchExchangeObjectsRequestParams `json:"exchangeParams,omitempty"`

	// Specifies the parameters which are specific for searching Public folder items.
	PublicFolderParams *SearchPublicFolderRequestParams `json:"publicFolderParams,omitempty"`

	// Specifies the parameters which are specific for searching Microsoft365 Group items.
	MsGroupsParams *SearchMsGroupsRequestParams `json:"msGroupsParams,omitempty"`

	// Specifies the parameters which are specific for searching Microsoft365 Teams items.
	MsTeamsParams *SearchMsTeamsRequestParams `json:"msTeamsParams,omitempty"`

	// Specifies the parameters which are specific for searching files/folders in a document library in a Sharepoint site.
	SharepointParams *SearchDocumentLibraryRequestParams `json:"sharepointParams,omitempty"`

	// Specifies the parameters which are specific for searching files/folders in a document library in a OneDrive.
	OneDriveParams *SearchDocumentLibraryRequestParams `json:"oneDriveParams,omitempty"`

	// Specifies the parameters which are specific for searching Universal Data Adapter objects.
	UdaParams *UdaOnPremSearchParams `json:"udaParams,omitempty"`

	// Specifies the parameters which are specific for searching records related to a salesforce Object.
	SfdcParams *SearchSfdcRecordsRequestParams `json:"sfdcParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SearchIndexedObjectsRequest) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonSearchIndexedObjectsRequestParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonSearchIndexedObjectsRequestParams = aO0

	// AO1
	var dataAO1 struct {
		EmailParams *SearchEmailRequestParams `json:"emailParams,omitempty"`

		FileParams *SearchFileRequestParams `json:"fileParams,omitempty"`

		CassandraParams *CassandraOnPremSearchParams `json:"cassandraParams,omitempty"`

		CouchbaseParams *CouchBaseOnPremSearchParams `json:"couchbaseParams,omitempty"`

		HbaseParams *HbaseOnPremSearchParams `json:"hbaseParams,omitempty"`

		HiveParams *HiveOnPremSearchParams `json:"hiveParams,omitempty"`

		MongodbParams *MongoDbOnPremSearchParams `json:"mongodbParams,omitempty"`

		HdfsParams *HDFSOnPremSearchParams `json:"hdfsParams,omitempty"`

		ExchangeParams *SearchExchangeObjectsRequestParams `json:"exchangeParams,omitempty"`

		PublicFolderParams *SearchPublicFolderRequestParams `json:"publicFolderParams,omitempty"`

		MsGroupsParams *SearchMsGroupsRequestParams `json:"msGroupsParams,omitempty"`

		MsTeamsParams *SearchMsTeamsRequestParams `json:"msTeamsParams,omitempty"`

		SharepointParams *SearchDocumentLibraryRequestParams `json:"sharepointParams,omitempty"`

		OneDriveParams *SearchDocumentLibraryRequestParams `json:"oneDriveParams,omitempty"`

		UdaParams *UdaOnPremSearchParams `json:"udaParams,omitempty"`

		SfdcParams *SearchSfdcRecordsRequestParams `json:"sfdcParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EmailParams = dataAO1.EmailParams

	m.FileParams = dataAO1.FileParams

	m.CassandraParams = dataAO1.CassandraParams

	m.CouchbaseParams = dataAO1.CouchbaseParams

	m.HbaseParams = dataAO1.HbaseParams

	m.HiveParams = dataAO1.HiveParams

	m.MongodbParams = dataAO1.MongodbParams

	m.HdfsParams = dataAO1.HdfsParams

	m.ExchangeParams = dataAO1.ExchangeParams

	m.PublicFolderParams = dataAO1.PublicFolderParams

	m.MsGroupsParams = dataAO1.MsGroupsParams

	m.MsTeamsParams = dataAO1.MsTeamsParams

	m.SharepointParams = dataAO1.SharepointParams

	m.OneDriveParams = dataAO1.OneDriveParams

	m.UdaParams = dataAO1.UdaParams

	m.SfdcParams = dataAO1.SfdcParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SearchIndexedObjectsRequest) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonSearchIndexedObjectsRequestParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		EmailParams *SearchEmailRequestParams `json:"emailParams,omitempty"`

		FileParams *SearchFileRequestParams `json:"fileParams,omitempty"`

		CassandraParams *CassandraOnPremSearchParams `json:"cassandraParams,omitempty"`

		CouchbaseParams *CouchBaseOnPremSearchParams `json:"couchbaseParams,omitempty"`

		HbaseParams *HbaseOnPremSearchParams `json:"hbaseParams,omitempty"`

		HiveParams *HiveOnPremSearchParams `json:"hiveParams,omitempty"`

		MongodbParams *MongoDbOnPremSearchParams `json:"mongodbParams,omitempty"`

		HdfsParams *HDFSOnPremSearchParams `json:"hdfsParams,omitempty"`

		ExchangeParams *SearchExchangeObjectsRequestParams `json:"exchangeParams,omitempty"`

		PublicFolderParams *SearchPublicFolderRequestParams `json:"publicFolderParams,omitempty"`

		MsGroupsParams *SearchMsGroupsRequestParams `json:"msGroupsParams,omitempty"`

		MsTeamsParams *SearchMsTeamsRequestParams `json:"msTeamsParams,omitempty"`

		SharepointParams *SearchDocumentLibraryRequestParams `json:"sharepointParams,omitempty"`

		OneDriveParams *SearchDocumentLibraryRequestParams `json:"oneDriveParams,omitempty"`

		UdaParams *UdaOnPremSearchParams `json:"udaParams,omitempty"`

		SfdcParams *SearchSfdcRecordsRequestParams `json:"sfdcParams,omitempty"`
	}

	dataAO1.EmailParams = m.EmailParams

	dataAO1.FileParams = m.FileParams

	dataAO1.CassandraParams = m.CassandraParams

	dataAO1.CouchbaseParams = m.CouchbaseParams

	dataAO1.HbaseParams = m.HbaseParams

	dataAO1.HiveParams = m.HiveParams

	dataAO1.MongodbParams = m.MongodbParams

	dataAO1.HdfsParams = m.HdfsParams

	dataAO1.ExchangeParams = m.ExchangeParams

	dataAO1.PublicFolderParams = m.PublicFolderParams

	dataAO1.MsGroupsParams = m.MsGroupsParams

	dataAO1.MsTeamsParams = m.MsTeamsParams

	dataAO1.SharepointParams = m.SharepointParams

	dataAO1.OneDriveParams = m.OneDriveParams

	dataAO1.UdaParams = m.UdaParams

	dataAO1.SfdcParams = m.SfdcParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this search indexed objects request
func (m *SearchIndexedObjectsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSearchIndexedObjectsRequestParams
	if err := m.CommonSearchIndexedObjectsRequestParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCassandraParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCouchbaseParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHbaseParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHiveParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongodbParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdfsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicFolderParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsGroupsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsTeamsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharepointParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneDriveParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUdaParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSfdcParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchIndexedObjectsRequest) validateEmailParams(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailParams) { // not required
		return nil
	}

	if m.EmailParams != nil {
		if err := m.EmailParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateFileParams(formats strfmt.Registry) error {

	if swag.IsZero(m.FileParams) { // not required
		return nil
	}

	if m.FileParams != nil {
		if err := m.FileParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateCassandraParams(formats strfmt.Registry) error {

	if swag.IsZero(m.CassandraParams) { // not required
		return nil
	}

	if m.CassandraParams != nil {
		if err := m.CassandraParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateCouchbaseParams(formats strfmt.Registry) error {

	if swag.IsZero(m.CouchbaseParams) { // not required
		return nil
	}

	if m.CouchbaseParams != nil {
		if err := m.CouchbaseParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("couchbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("couchbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateHbaseParams(formats strfmt.Registry) error {

	if swag.IsZero(m.HbaseParams) { // not required
		return nil
	}

	if m.HbaseParams != nil {
		if err := m.HbaseParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateHiveParams(formats strfmt.Registry) error {

	if swag.IsZero(m.HiveParams) { // not required
		return nil
	}

	if m.HiveParams != nil {
		if err := m.HiveParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hiveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hiveParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateMongodbParams(formats strfmt.Registry) error {

	if swag.IsZero(m.MongodbParams) { // not required
		return nil
	}

	if m.MongodbParams != nil {
		if err := m.MongodbParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongodbParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongodbParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateHdfsParams(formats strfmt.Registry) error {

	if swag.IsZero(m.HdfsParams) { // not required
		return nil
	}

	if m.HdfsParams != nil {
		if err := m.HdfsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hdfsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hdfsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateExchangeParams(formats strfmt.Registry) error {

	if swag.IsZero(m.ExchangeParams) { // not required
		return nil
	}

	if m.ExchangeParams != nil {
		if err := m.ExchangeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exchangeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exchangeParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validatePublicFolderParams(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicFolderParams) { // not required
		return nil
	}

	if m.PublicFolderParams != nil {
		if err := m.PublicFolderParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateMsGroupsParams(formats strfmt.Registry) error {

	if swag.IsZero(m.MsGroupsParams) { // not required
		return nil
	}

	if m.MsGroupsParams != nil {
		if err := m.MsGroupsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("msGroupsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("msGroupsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateMsTeamsParams(formats strfmt.Registry) error {

	if swag.IsZero(m.MsTeamsParams) { // not required
		return nil
	}

	if m.MsTeamsParams != nil {
		if err := m.MsTeamsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("msTeamsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("msTeamsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateSharepointParams(formats strfmt.Registry) error {

	if swag.IsZero(m.SharepointParams) { // not required
		return nil
	}

	if m.SharepointParams != nil {
		if err := m.SharepointParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharepointParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharepointParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateOneDriveParams(formats strfmt.Registry) error {

	if swag.IsZero(m.OneDriveParams) { // not required
		return nil
	}

	if m.OneDriveParams != nil {
		if err := m.OneDriveParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oneDriveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oneDriveParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateUdaParams(formats strfmt.Registry) error {

	if swag.IsZero(m.UdaParams) { // not required
		return nil
	}

	if m.UdaParams != nil {
		if err := m.UdaParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("udaParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("udaParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) validateSfdcParams(formats strfmt.Registry) error {

	if swag.IsZero(m.SfdcParams) { // not required
		return nil
	}

	if m.SfdcParams != nil {
		if err := m.SfdcParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfdcParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sfdcParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search indexed objects request based on the context it is used
func (m *SearchIndexedObjectsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSearchIndexedObjectsRequestParams
	if err := m.CommonSearchIndexedObjectsRequestParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCassandraParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCouchbaseParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHbaseParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHiveParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMongodbParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHdfsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExchangeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicFolderParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsGroupsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsTeamsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSharepointParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOneDriveParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUdaParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSfdcParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateEmailParams(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailParams != nil {

		if swag.IsZero(m.EmailParams) { // not required
			return nil
		}

		if err := m.EmailParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateFileParams(ctx context.Context, formats strfmt.Registry) error {

	if m.FileParams != nil {

		if swag.IsZero(m.FileParams) { // not required
			return nil
		}

		if err := m.FileParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateCassandraParams(ctx context.Context, formats strfmt.Registry) error {

	if m.CassandraParams != nil {

		if swag.IsZero(m.CassandraParams) { // not required
			return nil
		}

		if err := m.CassandraParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateCouchbaseParams(ctx context.Context, formats strfmt.Registry) error {

	if m.CouchbaseParams != nil {

		if swag.IsZero(m.CouchbaseParams) { // not required
			return nil
		}

		if err := m.CouchbaseParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("couchbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("couchbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateHbaseParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HbaseParams != nil {

		if swag.IsZero(m.HbaseParams) { // not required
			return nil
		}

		if err := m.HbaseParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateHiveParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HiveParams != nil {

		if swag.IsZero(m.HiveParams) { // not required
			return nil
		}

		if err := m.HiveParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hiveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hiveParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateMongodbParams(ctx context.Context, formats strfmt.Registry) error {

	if m.MongodbParams != nil {

		if swag.IsZero(m.MongodbParams) { // not required
			return nil
		}

		if err := m.MongodbParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongodbParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongodbParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateHdfsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HdfsParams != nil {

		if swag.IsZero(m.HdfsParams) { // not required
			return nil
		}

		if err := m.HdfsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hdfsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hdfsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateExchangeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ExchangeParams != nil {

		if swag.IsZero(m.ExchangeParams) { // not required
			return nil
		}

		if err := m.ExchangeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exchangeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exchangeParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidatePublicFolderParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicFolderParams != nil {

		if swag.IsZero(m.PublicFolderParams) { // not required
			return nil
		}

		if err := m.PublicFolderParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicFolderParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicFolderParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateMsGroupsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.MsGroupsParams != nil {

		if swag.IsZero(m.MsGroupsParams) { // not required
			return nil
		}

		if err := m.MsGroupsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("msGroupsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("msGroupsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateMsTeamsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.MsTeamsParams != nil {

		if swag.IsZero(m.MsTeamsParams) { // not required
			return nil
		}

		if err := m.MsTeamsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("msTeamsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("msTeamsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateSharepointParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SharepointParams != nil {

		if swag.IsZero(m.SharepointParams) { // not required
			return nil
		}

		if err := m.SharepointParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharepointParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharepointParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateOneDriveParams(ctx context.Context, formats strfmt.Registry) error {

	if m.OneDriveParams != nil {

		if swag.IsZero(m.OneDriveParams) { // not required
			return nil
		}

		if err := m.OneDriveParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oneDriveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oneDriveParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateUdaParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UdaParams != nil {

		if swag.IsZero(m.UdaParams) { // not required
			return nil
		}

		if err := m.UdaParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("udaParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("udaParams")
			}
			return err
		}
	}

	return nil
}

func (m *SearchIndexedObjectsRequest) contextValidateSfdcParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SfdcParams != nil {

		if swag.IsZero(m.SfdcParams) { // not required
			return nil
		}

		if err := m.SfdcParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfdcParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sfdcParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchIndexedObjectsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchIndexedObjectsRequest) UnmarshalBinary(b []byte) error {
	var res SearchIndexedObjectsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
