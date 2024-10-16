// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

// UpdateActiveDirectoryIDMappingReader is a Reader for the UpdateActiveDirectoryIDMapping structure.
type UpdateActiveDirectoryIDMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateActiveDirectoryIDMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateActiveDirectoryIDMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateActiveDirectoryIDMappingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateActiveDirectoryIDMappingOK creates a UpdateActiveDirectoryIDMappingOK with default headers values
func NewUpdateActiveDirectoryIDMappingOK() *UpdateActiveDirectoryIDMappingOK {
	return &UpdateActiveDirectoryIDMappingOK{}
}

/*
UpdateActiveDirectoryIDMappingOK describes a response with status code 200, with default header values.

Success
*/
type UpdateActiveDirectoryIDMappingOK struct {
	Payload *models.ActiveDirectoryEntry
}

// IsSuccess returns true when this update active directory Id mapping o k response has a 2xx status code
func (o *UpdateActiveDirectoryIDMappingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update active directory Id mapping o k response has a 3xx status code
func (o *UpdateActiveDirectoryIDMappingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update active directory Id mapping o k response has a 4xx status code
func (o *UpdateActiveDirectoryIDMappingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update active directory Id mapping o k response has a 5xx status code
func (o *UpdateActiveDirectoryIDMappingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update active directory Id mapping o k response a status code equal to that given
func (o *UpdateActiveDirectoryIDMappingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update active directory Id mapping o k response
func (o *UpdateActiveDirectoryIDMappingOK) Code() int {
	return 200
}

func (o *UpdateActiveDirectoryIDMappingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/activeDirectory/{name}/idMappingInfo][%d] updateActiveDirectoryIdMappingOK %s", 200, payload)
}

func (o *UpdateActiveDirectoryIDMappingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/activeDirectory/{name}/idMappingInfo][%d] updateActiveDirectoryIdMappingOK %s", 200, payload)
}

func (o *UpdateActiveDirectoryIDMappingOK) GetPayload() *models.ActiveDirectoryEntry {
	return o.Payload
}

func (o *UpdateActiveDirectoryIDMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveDirectoryEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActiveDirectoryIDMappingDefault creates a UpdateActiveDirectoryIDMappingDefault with default headers values
func NewUpdateActiveDirectoryIDMappingDefault(code int) *UpdateActiveDirectoryIDMappingDefault {
	return &UpdateActiveDirectoryIDMappingDefault{
		_statusCode: code,
	}
}

/*
UpdateActiveDirectoryIDMappingDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateActiveDirectoryIDMappingDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update active directory Id mapping default response has a 2xx status code
func (o *UpdateActiveDirectoryIDMappingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update active directory Id mapping default response has a 3xx status code
func (o *UpdateActiveDirectoryIDMappingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update active directory Id mapping default response has a 4xx status code
func (o *UpdateActiveDirectoryIDMappingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update active directory Id mapping default response has a 5xx status code
func (o *UpdateActiveDirectoryIDMappingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update active directory Id mapping default response a status code equal to that given
func (o *UpdateActiveDirectoryIDMappingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update active directory Id mapping default response
func (o *UpdateActiveDirectoryIDMappingDefault) Code() int {
	return o._statusCode
}

func (o *UpdateActiveDirectoryIDMappingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/activeDirectory/{name}/idMappingInfo][%d] UpdateActiveDirectoryIdMapping default %s", o._statusCode, payload)
}

func (o *UpdateActiveDirectoryIDMappingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/activeDirectory/{name}/idMappingInfo][%d] UpdateActiveDirectoryIdMapping default %s", o._statusCode, payload)
}

func (o *UpdateActiveDirectoryIDMappingDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateActiveDirectoryIDMappingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
