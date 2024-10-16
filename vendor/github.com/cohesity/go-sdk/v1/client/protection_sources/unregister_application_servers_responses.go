// Code generated by go-swagger; DO NOT EDIT.

package protection_sources

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

// UnregisterApplicationServersReader is a Reader for the UnregisterApplicationServers structure.
type UnregisterApplicationServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnregisterApplicationServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnregisterApplicationServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnregisterApplicationServersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnregisterApplicationServersOK creates a UnregisterApplicationServersOK with default headers values
func NewUnregisterApplicationServersOK() *UnregisterApplicationServersOK {
	return &UnregisterApplicationServersOK{}
}

/*
UnregisterApplicationServersOK describes a response with status code 200, with default header values.

Success
*/
type UnregisterApplicationServersOK struct {
	Payload *models.ProtectionSource
}

// IsSuccess returns true when this unregister application servers o k response has a 2xx status code
func (o *UnregisterApplicationServersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unregister application servers o k response has a 3xx status code
func (o *UnregisterApplicationServersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unregister application servers o k response has a 4xx status code
func (o *UnregisterApplicationServersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unregister application servers o k response has a 5xx status code
func (o *UnregisterApplicationServersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unregister application servers o k response a status code equal to that given
func (o *UnregisterApplicationServersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unregister application servers o k response
func (o *UnregisterApplicationServersOK) Code() int {
	return 200
}

func (o *UnregisterApplicationServersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/protectionSources/applicationServers/{id}][%d] unregisterApplicationServersOK %s", 200, payload)
}

func (o *UnregisterApplicationServersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/protectionSources/applicationServers/{id}][%d] unregisterApplicationServersOK %s", 200, payload)
}

func (o *UnregisterApplicationServersOK) GetPayload() *models.ProtectionSource {
	return o.Payload
}

func (o *UnregisterApplicationServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnregisterApplicationServersDefault creates a UnregisterApplicationServersDefault with default headers values
func NewUnregisterApplicationServersDefault(code int) *UnregisterApplicationServersDefault {
	return &UnregisterApplicationServersDefault{
		_statusCode: code,
	}
}

/*
UnregisterApplicationServersDefault describes a response with status code -1, with default header values.

Error
*/
type UnregisterApplicationServersDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this unregister application servers default response has a 2xx status code
func (o *UnregisterApplicationServersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unregister application servers default response has a 3xx status code
func (o *UnregisterApplicationServersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unregister application servers default response has a 4xx status code
func (o *UnregisterApplicationServersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unregister application servers default response has a 5xx status code
func (o *UnregisterApplicationServersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unregister application servers default response a status code equal to that given
func (o *UnregisterApplicationServersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unregister application servers default response
func (o *UnregisterApplicationServersDefault) Code() int {
	return o._statusCode
}

func (o *UnregisterApplicationServersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/protectionSources/applicationServers/{id}][%d] UnregisterApplicationServers default %s", o._statusCode, payload)
}

func (o *UnregisterApplicationServersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/protectionSources/applicationServers/{id}][%d] UnregisterApplicationServers default %s", o._statusCode, payload)
}

func (o *UnregisterApplicationServersDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UnregisterApplicationServersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
