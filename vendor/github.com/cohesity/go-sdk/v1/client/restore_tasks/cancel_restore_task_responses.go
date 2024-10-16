// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

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

// CancelRestoreTaskReader is a Reader for the CancelRestoreTask structure.
type CancelRestoreTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelRestoreTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelRestoreTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCancelRestoreTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelRestoreTaskOK creates a CancelRestoreTaskOK with default headers values
func NewCancelRestoreTaskOK() *CancelRestoreTaskOK {
	return &CancelRestoreTaskOK{}
}

/*
CancelRestoreTaskOK describes a response with status code 200, with default header values.

Success Response
*/
type CancelRestoreTaskOK struct {
}

// IsSuccess returns true when this cancel restore task o k response has a 2xx status code
func (o *CancelRestoreTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel restore task o k response has a 3xx status code
func (o *CancelRestoreTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel restore task o k response has a 4xx status code
func (o *CancelRestoreTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel restore task o k response has a 5xx status code
func (o *CancelRestoreTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel restore task o k response a status code equal to that given
func (o *CancelRestoreTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancel restore task o k response
func (o *CancelRestoreTaskOK) Code() int {
	return 200
}

func (o *CancelRestoreTaskOK) Error() string {
	return fmt.Sprintf("[PUT /public/restore/tasks/cancel/{id}][%d] cancelRestoreTaskOK", 200)
}

func (o *CancelRestoreTaskOK) String() string {
	return fmt.Sprintf("[PUT /public/restore/tasks/cancel/{id}][%d] cancelRestoreTaskOK", 200)
}

func (o *CancelRestoreTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelRestoreTaskDefault creates a CancelRestoreTaskDefault with default headers values
func NewCancelRestoreTaskDefault(code int) *CancelRestoreTaskDefault {
	return &CancelRestoreTaskDefault{
		_statusCode: code,
	}
}

/*
CancelRestoreTaskDefault describes a response with status code -1, with default header values.

Error
*/
type CancelRestoreTaskDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this cancel restore task default response has a 2xx status code
func (o *CancelRestoreTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cancel restore task default response has a 3xx status code
func (o *CancelRestoreTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cancel restore task default response has a 4xx status code
func (o *CancelRestoreTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cancel restore task default response has a 5xx status code
func (o *CancelRestoreTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cancel restore task default response a status code equal to that given
func (o *CancelRestoreTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cancel restore task default response
func (o *CancelRestoreTaskDefault) Code() int {
	return o._statusCode
}

func (o *CancelRestoreTaskDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/restore/tasks/cancel/{id}][%d] CancelRestoreTask default %s", o._statusCode, payload)
}

func (o *CancelRestoreTaskDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/restore/tasks/cancel/{id}][%d] CancelRestoreTask default %s", o._statusCode, payload)
}

func (o *CancelRestoreTaskDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CancelRestoreTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
