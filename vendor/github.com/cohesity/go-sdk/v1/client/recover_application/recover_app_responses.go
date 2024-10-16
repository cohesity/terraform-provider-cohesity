// Code generated by go-swagger; DO NOT EDIT.

package recover_application

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

// RecoverAppReader is a Reader for the RecoverApp structure.
type RecoverAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecoverAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRecoverAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRecoverAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRecoverAppOK creates a RecoverAppOK with default headers values
func NewRecoverAppOK() *RecoverAppOK {
	return &RecoverAppOK{}
}

/*
RecoverAppOK describes a response with status code 200, with default header values.

Success
*/
type RecoverAppOK struct {
	Payload *models.RestoreTaskWrapper
}

// IsSuccess returns true when this recover app o k response has a 2xx status code
func (o *RecoverAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this recover app o k response has a 3xx status code
func (o *RecoverAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this recover app o k response has a 4xx status code
func (o *RecoverAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this recover app o k response has a 5xx status code
func (o *RecoverAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this recover app o k response a status code equal to that given
func (o *RecoverAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the recover app o k response
func (o *RecoverAppOK) Code() int {
	return 200
}

func (o *RecoverAppOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /recoverApplication][%d] recoverAppOK %s", 200, payload)
}

func (o *RecoverAppOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /recoverApplication][%d] recoverAppOK %s", 200, payload)
}

func (o *RecoverAppOK) GetPayload() *models.RestoreTaskWrapper {
	return o.Payload
}

func (o *RecoverAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreTaskWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecoverAppDefault creates a RecoverAppDefault with default headers values
func NewRecoverAppDefault(code int) *RecoverAppDefault {
	return &RecoverAppDefault{
		_statusCode: code,
	}
}

/*
RecoverAppDefault describes a response with status code -1, with default header values.

(empty)
*/
type RecoverAppDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this recover app default response has a 2xx status code
func (o *RecoverAppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this recover app default response has a 3xx status code
func (o *RecoverAppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this recover app default response has a 4xx status code
func (o *RecoverAppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this recover app default response has a 5xx status code
func (o *RecoverAppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this recover app default response a status code equal to that given
func (o *RecoverAppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the recover app default response
func (o *RecoverAppDefault) Code() int {
	return o._statusCode
}

func (o *RecoverAppDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /recoverApplication][%d] RecoverApp default %s", o._statusCode, payload)
}

func (o *RecoverAppDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /recoverApplication][%d] RecoverApp default %s", o._statusCode, payload)
}

func (o *RecoverAppDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *RecoverAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
