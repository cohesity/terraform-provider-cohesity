// Code generated by go-swagger; DO NOT EDIT.

package storage_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cohesity/go-sdk/v2/models"
)

// GetStorageDomainsReader is a Reader for the GetStorageDomains structure.
type GetStorageDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStorageDomainsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStorageDomainsOK creates a GetStorageDomainsOK with default headers values
func NewGetStorageDomainsOK() *GetStorageDomainsOK {
	return &GetStorageDomainsOK{}
}

/*
GetStorageDomainsOK describes a response with status code 200, with default header values.

Success
*/
type GetStorageDomainsOK struct {
	Payload *models.StorageDomains
}

// IsSuccess returns true when this get storage domains o k response has a 2xx status code
func (o *GetStorageDomainsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get storage domains o k response has a 3xx status code
func (o *GetStorageDomainsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage domains o k response has a 4xx status code
func (o *GetStorageDomainsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get storage domains o k response has a 5xx status code
func (o *GetStorageDomainsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get storage domains o k response a status code equal to that given
func (o *GetStorageDomainsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get storage domains o k response
func (o *GetStorageDomainsOK) Code() int {
	return 200
}

func (o *GetStorageDomainsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage-domains][%d] getStorageDomainsOK %s", 200, payload)
}

func (o *GetStorageDomainsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage-domains][%d] getStorageDomainsOK %s", 200, payload)
}

func (o *GetStorageDomainsOK) GetPayload() *models.StorageDomains {
	return o.Payload
}

func (o *GetStorageDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageDomains)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageDomainsDefault creates a GetStorageDomainsDefault with default headers values
func NewGetStorageDomainsDefault(code int) *GetStorageDomainsDefault {
	return &GetStorageDomainsDefault{
		_statusCode: code,
	}
}

/*
GetStorageDomainsDefault describes a response with status code -1, with default header values.

Error
*/
type GetStorageDomainsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get storage domains default response has a 2xx status code
func (o *GetStorageDomainsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get storage domains default response has a 3xx status code
func (o *GetStorageDomainsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get storage domains default response has a 4xx status code
func (o *GetStorageDomainsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get storage domains default response has a 5xx status code
func (o *GetStorageDomainsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get storage domains default response a status code equal to that given
func (o *GetStorageDomainsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get storage domains default response
func (o *GetStorageDomainsDefault) Code() int {
	return o._statusCode
}

func (o *GetStorageDomainsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage-domains][%d] GetStorageDomains default %s", o._statusCode, payload)
}

func (o *GetStorageDomainsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /storage-domains][%d] GetStorageDomains default %s", o._statusCode, payload)
}

func (o *GetStorageDomainsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageDomainsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
