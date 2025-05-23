// Code generated by go-swagger; DO NOT EDIT.

package snmp_config

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

// UpdateSnmpConfigReader is a Reader for the UpdateSnmpConfig structure.
type UpdateSnmpConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnmpConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnmpConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSnmpConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSnmpConfigOK creates a UpdateSnmpConfigOK with default headers values
func NewUpdateSnmpConfigOK() *UpdateSnmpConfigOK {
	return &UpdateSnmpConfigOK{}
}

/*
UpdateSnmpConfigOK describes a response with status code 200, with default header values.

Success
*/
type UpdateSnmpConfigOK struct {
	Payload *models.SnmpConfig
}

// IsSuccess returns true when this update snmp config o k response has a 2xx status code
func (o *UpdateSnmpConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update snmp config o k response has a 3xx status code
func (o *UpdateSnmpConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snmp config o k response has a 4xx status code
func (o *UpdateSnmpConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update snmp config o k response has a 5xx status code
func (o *UpdateSnmpConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update snmp config o k response a status code equal to that given
func (o *UpdateSnmpConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update snmp config o k response
func (o *UpdateSnmpConfigOK) Code() int {
	return 200
}

func (o *UpdateSnmpConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /snmp/config][%d] updateSnmpConfigOK %s", 200, payload)
}

func (o *UpdateSnmpConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /snmp/config][%d] updateSnmpConfigOK %s", 200, payload)
}

func (o *UpdateSnmpConfigOK) GetPayload() *models.SnmpConfig {
	return o.Payload
}

func (o *UpdateSnmpConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnmpConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSnmpConfigDefault creates a UpdateSnmpConfigDefault with default headers values
func NewUpdateSnmpConfigDefault(code int) *UpdateSnmpConfigDefault {
	return &UpdateSnmpConfigDefault{
		_statusCode: code,
	}
}

/*
UpdateSnmpConfigDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateSnmpConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update snmp config default response has a 2xx status code
func (o *UpdateSnmpConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update snmp config default response has a 3xx status code
func (o *UpdateSnmpConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update snmp config default response has a 4xx status code
func (o *UpdateSnmpConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update snmp config default response has a 5xx status code
func (o *UpdateSnmpConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update snmp config default response a status code equal to that given
func (o *UpdateSnmpConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update snmp config default response
func (o *UpdateSnmpConfigDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSnmpConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /snmp/config][%d] UpdateSnmpConfig default %s", o._statusCode, payload)
}

func (o *UpdateSnmpConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /snmp/config][%d] UpdateSnmpConfig default %s", o._statusCode, payload)
}

func (o *UpdateSnmpConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSnmpConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
