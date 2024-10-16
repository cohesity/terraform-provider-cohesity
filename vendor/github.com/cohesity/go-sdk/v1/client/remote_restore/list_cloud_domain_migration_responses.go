// Code generated by go-swagger; DO NOT EDIT.

package remote_restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cohesity/go-sdk/v1/models"
)

// ListCloudDomainMigrationReader is a Reader for the ListCloudDomainMigration structure.
type ListCloudDomainMigrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCloudDomainMigrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCloudDomainMigrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListCloudDomainMigrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCloudDomainMigrationOK creates a ListCloudDomainMigrationOK with default headers values
func NewListCloudDomainMigrationOK() *ListCloudDomainMigrationOK {
	return &ListCloudDomainMigrationOK{}
}

/*
ListCloudDomainMigrationOK describes a response with status code 200, with default header values.

Success
*/
type ListCloudDomainMigrationOK struct {
	Payload *models.CloudDomainMigrationQueryResult
}

// IsSuccess returns true when this list cloud domain migration o k response has a 2xx status code
func (o *ListCloudDomainMigrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list cloud domain migration o k response has a 3xx status code
func (o *ListCloudDomainMigrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list cloud domain migration o k response has a 4xx status code
func (o *ListCloudDomainMigrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list cloud domain migration o k response has a 5xx status code
func (o *ListCloudDomainMigrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list cloud domain migration o k response a status code equal to that given
func (o *ListCloudDomainMigrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list cloud domain migration o k response
func (o *ListCloudDomainMigrationOK) Code() int {
	return 200
}

func (o *ListCloudDomainMigrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/cloudDomainMigration][%d] listCloudDomainMigrationOK %s", 200, payload)
}

func (o *ListCloudDomainMigrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/cloudDomainMigration][%d] listCloudDomainMigrationOK %s", 200, payload)
}

func (o *ListCloudDomainMigrationOK) GetPayload() *models.CloudDomainMigrationQueryResult {
	return o.Payload
}

func (o *ListCloudDomainMigrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudDomainMigrationQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudDomainMigrationDefault creates a ListCloudDomainMigrationDefault with default headers values
func NewListCloudDomainMigrationDefault(code int) *ListCloudDomainMigrationDefault {
	return &ListCloudDomainMigrationDefault{
		_statusCode: code,
	}
}

/*
ListCloudDomainMigrationDefault describes a response with status code -1, with default header values.

Error
*/
type ListCloudDomainMigrationDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this list cloud domain migration default response has a 2xx status code
func (o *ListCloudDomainMigrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list cloud domain migration default response has a 3xx status code
func (o *ListCloudDomainMigrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list cloud domain migration default response has a 4xx status code
func (o *ListCloudDomainMigrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list cloud domain migration default response has a 5xx status code
func (o *ListCloudDomainMigrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list cloud domain migration default response a status code equal to that given
func (o *ListCloudDomainMigrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list cloud domain migration default response
func (o *ListCloudDomainMigrationDefault) Code() int {
	return o._statusCode
}

func (o *ListCloudDomainMigrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/cloudDomainMigration][%d] ListCloudDomainMigration default %s", o._statusCode, payload)
}

func (o *ListCloudDomainMigrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/cloudDomainMigration][%d] ListCloudDomainMigration default %s", o._statusCode, payload)
}

func (o *ListCloudDomainMigrationDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *ListCloudDomainMigrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
