// Code generated by go-swagger; DO NOT EDIT.

package views

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

// LockFileByIDReader is a Reader for the LockFileByID structure.
type LockFileByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LockFileByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLockFileByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLockFileByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLockFileByIDOK creates a LockFileByIDOK with default headers values
func NewLockFileByIDOK() *LockFileByIDOK {
	return &LockFileByIDOK{}
}

/*
LockFileByIDOK describes a response with status code 200, with default header values.

Get lock file status response
*/
type LockFileByIDOK struct {
	Payload *models.FileLockStatus
}

// IsSuccess returns true when this lock file by Id o k response has a 2xx status code
func (o *LockFileByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lock file by Id o k response has a 3xx status code
func (o *LockFileByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lock file by Id o k response has a 4xx status code
func (o *LockFileByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lock file by Id o k response has a 5xx status code
func (o *LockFileByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lock file by Id o k response a status code equal to that given
func (o *LockFileByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lock file by Id o k response
func (o *LockFileByIDOK) Code() int {
	return 200
}

func (o *LockFileByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/id/{id}/fileLocks][%d] lockFileByIdOK %s", 200, payload)
}

func (o *LockFileByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/id/{id}/fileLocks][%d] lockFileByIdOK %s", 200, payload)
}

func (o *LockFileByIDOK) GetPayload() *models.FileLockStatus {
	return o.Payload
}

func (o *LockFileByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileLockStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLockFileByIDDefault creates a LockFileByIDDefault with default headers values
func NewLockFileByIDDefault(code int) *LockFileByIDDefault {
	return &LockFileByIDDefault{
		_statusCode: code,
	}
}

/*
LockFileByIDDefault describes a response with status code -1, with default header values.

Error
*/
type LockFileByIDDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this lock file by Id default response has a 2xx status code
func (o *LockFileByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lock file by Id default response has a 3xx status code
func (o *LockFileByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lock file by Id default response has a 4xx status code
func (o *LockFileByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lock file by Id default response has a 5xx status code
func (o *LockFileByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lock file by Id default response a status code equal to that given
func (o *LockFileByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lock file by Id default response
func (o *LockFileByIDDefault) Code() int {
	return o._statusCode
}

func (o *LockFileByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/id/{id}/fileLocks][%d] LockFileById default %s", o._statusCode, payload)
}

func (o *LockFileByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/id/{id}/fileLocks][%d] LockFileById default %s", o._statusCode, payload)
}

func (o *LockFileByIDDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *LockFileByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
