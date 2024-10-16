// Code generated by go-swagger; DO NOT EDIT.

package restore_files

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

// RestoreFilesReader is a Reader for the RestoreFiles structure.
type RestoreFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreFilesOK creates a RestoreFilesOK with default headers values
func NewRestoreFilesOK() *RestoreFilesOK {
	return &RestoreFilesOK{}
}

/*
RestoreFilesOK describes a response with status code 200, with default header values.

Success
*/
type RestoreFilesOK struct {
	Payload *models.RestoreTaskWrapper
}

// IsSuccess returns true when this restore files o k response has a 2xx status code
func (o *RestoreFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore files o k response has a 3xx status code
func (o *RestoreFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore files o k response has a 4xx status code
func (o *RestoreFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore files o k response has a 5xx status code
func (o *RestoreFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore files o k response a status code equal to that given
func (o *RestoreFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the restore files o k response
func (o *RestoreFilesOK) Code() int {
	return 200
}

func (o *RestoreFilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreFiles][%d] restoreFilesOK %s", 200, payload)
}

func (o *RestoreFilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreFiles][%d] restoreFilesOK %s", 200, payload)
}

func (o *RestoreFilesOK) GetPayload() *models.RestoreTaskWrapper {
	return o.Payload
}

func (o *RestoreFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreTaskWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreFilesDefault creates a RestoreFilesDefault with default headers values
func NewRestoreFilesDefault(code int) *RestoreFilesDefault {
	return &RestoreFilesDefault{
		_statusCode: code,
	}
}

/*
RestoreFilesDefault describes a response with status code -1, with default header values.

(empty)
*/
type RestoreFilesDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this restore files default response has a 2xx status code
func (o *RestoreFilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore files default response has a 3xx status code
func (o *RestoreFilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore files default response has a 4xx status code
func (o *RestoreFilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore files default response has a 5xx status code
func (o *RestoreFilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore files default response a status code equal to that given
func (o *RestoreFilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the restore files default response
func (o *RestoreFilesDefault) Code() int {
	return o._statusCode
}

func (o *RestoreFilesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreFiles][%d] RestoreFiles default %s", o._statusCode, payload)
}

func (o *RestoreFilesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreFiles][%d] RestoreFiles default %s", o._statusCode, payload)
}

func (o *RestoreFilesDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *RestoreFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
