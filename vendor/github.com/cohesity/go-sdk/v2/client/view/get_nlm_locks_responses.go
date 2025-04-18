// Code generated by go-swagger; DO NOT EDIT.

package view

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

// GetNlmLocksReader is a Reader for the GetNlmLocks structure.
type GetNlmLocksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNlmLocksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNlmLocksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNlmLocksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNlmLocksOK creates a GetNlmLocksOK with default headers values
func NewGetNlmLocksOK() *GetNlmLocksOK {
	return &GetNlmLocksOK{}
}

/*
GetNlmLocksOK describes a response with status code 200, with default header values.

Success
*/
type GetNlmLocksOK struct {
	Payload *models.GetNlmLocksResult
}

// IsSuccess returns true when this get nlm locks o k response has a 2xx status code
func (o *GetNlmLocksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nlm locks o k response has a 3xx status code
func (o *GetNlmLocksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nlm locks o k response has a 4xx status code
func (o *GetNlmLocksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nlm locks o k response has a 5xx status code
func (o *GetNlmLocksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nlm locks o k response a status code equal to that given
func (o *GetNlmLocksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get nlm locks o k response
func (o *GetNlmLocksOK) Code() int {
	return 200
}

func (o *GetNlmLocksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/nlm-locks][%d] getNlmLocksOK %s", 200, payload)
}

func (o *GetNlmLocksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/nlm-locks][%d] getNlmLocksOK %s", 200, payload)
}

func (o *GetNlmLocksOK) GetPayload() *models.GetNlmLocksResult {
	return o.Payload
}

func (o *GetNlmLocksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNlmLocksResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNlmLocksDefault creates a GetNlmLocksDefault with default headers values
func NewGetNlmLocksDefault(code int) *GetNlmLocksDefault {
	return &GetNlmLocksDefault{
		_statusCode: code,
	}
}

/*
GetNlmLocksDefault describes a response with status code -1, with default header values.

Error
*/
type GetNlmLocksDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get nlm locks default response has a 2xx status code
func (o *GetNlmLocksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get nlm locks default response has a 3xx status code
func (o *GetNlmLocksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get nlm locks default response has a 4xx status code
func (o *GetNlmLocksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get nlm locks default response has a 5xx status code
func (o *GetNlmLocksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get nlm locks default response a status code equal to that given
func (o *GetNlmLocksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get nlm locks default response
func (o *GetNlmLocksDefault) Code() int {
	return o._statusCode
}

func (o *GetNlmLocksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/nlm-locks][%d] GetNlmLocks default %s", o._statusCode, payload)
}

func (o *GetNlmLocksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/nlm-locks][%d] GetNlmLocks default %s", o._statusCode, payload)
}

func (o *GetNlmLocksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNlmLocksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
