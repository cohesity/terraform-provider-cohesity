// Code generated by go-swagger; DO NOT EDIT.

package external_target

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

// GetExternalTargetsReader is a Reader for the GetExternalTargets structure.
type GetExternalTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalTargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetExternalTargetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalTargetsOK creates a GetExternalTargetsOK with default headers values
func NewGetExternalTargetsOK() *GetExternalTargetsOK {
	return &GetExternalTargetsOK{}
}

/*
GetExternalTargetsOK describes a response with status code 200, with default header values.

Success
*/
type GetExternalTargetsOK struct {
	Payload *models.ExternalTargets
}

// IsSuccess returns true when this get external targets o k response has a 2xx status code
func (o *GetExternalTargetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get external targets o k response has a 3xx status code
func (o *GetExternalTargetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external targets o k response has a 4xx status code
func (o *GetExternalTargetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get external targets o k response has a 5xx status code
func (o *GetExternalTargetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get external targets o k response a status code equal to that given
func (o *GetExternalTargetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get external targets o k response
func (o *GetExternalTargetsOK) Code() int {
	return 200
}

func (o *GetExternalTargetsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets][%d] getExternalTargetsOK %s", 200, payload)
}

func (o *GetExternalTargetsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets][%d] getExternalTargetsOK %s", 200, payload)
}

func (o *GetExternalTargetsOK) GetPayload() *models.ExternalTargets {
	return o.Payload
}

func (o *GetExternalTargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalTargets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalTargetsDefault creates a GetExternalTargetsDefault with default headers values
func NewGetExternalTargetsDefault(code int) *GetExternalTargetsDefault {
	return &GetExternalTargetsDefault{
		_statusCode: code,
	}
}

/*
GetExternalTargetsDefault describes a response with status code -1, with default header values.

Error
*/
type GetExternalTargetsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get external targets default response has a 2xx status code
func (o *GetExternalTargetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get external targets default response has a 3xx status code
func (o *GetExternalTargetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get external targets default response has a 4xx status code
func (o *GetExternalTargetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get external targets default response has a 5xx status code
func (o *GetExternalTargetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get external targets default response a status code equal to that given
func (o *GetExternalTargetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get external targets default response
func (o *GetExternalTargetsDefault) Code() int {
	return o._statusCode
}

func (o *GetExternalTargetsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets][%d] GetExternalTargets default %s", o._statusCode, payload)
}

func (o *GetExternalTargetsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets][%d] GetExternalTargets default %s", o._statusCode, payload)
}

func (o *GetExternalTargetsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExternalTargetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
