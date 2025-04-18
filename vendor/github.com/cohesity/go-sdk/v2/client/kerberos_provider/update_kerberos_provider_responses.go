// Code generated by go-swagger; DO NOT EDIT.

package kerberos_provider

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

// UpdateKerberosProviderReader is a Reader for the UpdateKerberosProvider structure.
type UpdateKerberosProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateKerberosProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateKerberosProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateKerberosProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateKerberosProviderOK creates a UpdateKerberosProviderOK with default headers values
func NewUpdateKerberosProviderOK() *UpdateKerberosProviderOK {
	return &UpdateKerberosProviderOK{}
}

/*
UpdateKerberosProviderOK describes a response with status code 200, with default header values.

Success
*/
type UpdateKerberosProviderOK struct {
	Payload *models.KerberosProvider
}

// IsSuccess returns true when this update kerberos provider o k response has a 2xx status code
func (o *UpdateKerberosProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update kerberos provider o k response has a 3xx status code
func (o *UpdateKerberosProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kerberos provider o k response has a 4xx status code
func (o *UpdateKerberosProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update kerberos provider o k response has a 5xx status code
func (o *UpdateKerberosProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update kerberos provider o k response a status code equal to that given
func (o *UpdateKerberosProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update kerberos provider o k response
func (o *UpdateKerberosProviderOK) Code() int {
	return 200
}

func (o *UpdateKerberosProviderOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /kerberos-providers/{id}][%d] updateKerberosProviderOK %s", 200, payload)
}

func (o *UpdateKerberosProviderOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /kerberos-providers/{id}][%d] updateKerberosProviderOK %s", 200, payload)
}

func (o *UpdateKerberosProviderOK) GetPayload() *models.KerberosProvider {
	return o.Payload
}

func (o *UpdateKerberosProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KerberosProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKerberosProviderDefault creates a UpdateKerberosProviderDefault with default headers values
func NewUpdateKerberosProviderDefault(code int) *UpdateKerberosProviderDefault {
	return &UpdateKerberosProviderDefault{
		_statusCode: code,
	}
}

/*
UpdateKerberosProviderDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateKerberosProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update kerberos provider default response has a 2xx status code
func (o *UpdateKerberosProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update kerberos provider default response has a 3xx status code
func (o *UpdateKerberosProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update kerberos provider default response has a 4xx status code
func (o *UpdateKerberosProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update kerberos provider default response has a 5xx status code
func (o *UpdateKerberosProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update kerberos provider default response a status code equal to that given
func (o *UpdateKerberosProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update kerberos provider default response
func (o *UpdateKerberosProviderDefault) Code() int {
	return o._statusCode
}

func (o *UpdateKerberosProviderDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /kerberos-providers/{id}][%d] UpdateKerberosProvider default %s", o._statusCode, payload)
}

func (o *UpdateKerberosProviderDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /kerberos-providers/{id}][%d] UpdateKerberosProvider default %s", o._statusCode, payload)
}

func (o *UpdateKerberosProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateKerberosProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
