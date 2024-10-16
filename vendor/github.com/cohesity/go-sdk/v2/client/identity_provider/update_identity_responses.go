// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

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

// UpdateIdentityReader is a Reader for the UpdateIdentity structure.
type UpdateIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateIdentityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIdentityOK creates a UpdateIdentityOK with default headers values
func NewUpdateIdentityOK() *UpdateIdentityOK {
	return &UpdateIdentityOK{}
}

/*
UpdateIdentityOK describes a response with status code 200, with default header values.

Success
*/
type UpdateIdentityOK struct {
	Payload *models.IdentityConfig
}

// IsSuccess returns true when this update identity o k response has a 2xx status code
func (o *UpdateIdentityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update identity o k response has a 3xx status code
func (o *UpdateIdentityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update identity o k response has a 4xx status code
func (o *UpdateIdentityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update identity o k response has a 5xx status code
func (o *UpdateIdentityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update identity o k response a status code equal to that given
func (o *UpdateIdentityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update identity o k response
func (o *UpdateIdentityOK) Code() int {
	return 200
}

func (o *UpdateIdentityOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /identity-providers/{id}][%d] updateIdentityOK %s", 200, payload)
}

func (o *UpdateIdentityOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /identity-providers/{id}][%d] updateIdentityOK %s", 200, payload)
}

func (o *UpdateIdentityOK) GetPayload() *models.IdentityConfig {
	return o.Payload
}

func (o *UpdateIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentityConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIdentityDefault creates a UpdateIdentityDefault with default headers values
func NewUpdateIdentityDefault(code int) *UpdateIdentityDefault {
	return &UpdateIdentityDefault{
		_statusCode: code,
	}
}

/*
UpdateIdentityDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateIdentityDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update identity default response has a 2xx status code
func (o *UpdateIdentityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update identity default response has a 3xx status code
func (o *UpdateIdentityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update identity default response has a 4xx status code
func (o *UpdateIdentityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update identity default response has a 5xx status code
func (o *UpdateIdentityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update identity default response a status code equal to that given
func (o *UpdateIdentityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update identity default response
func (o *UpdateIdentityDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIdentityDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /identity-providers/{id}][%d] UpdateIdentity default %s", o._statusCode, payload)
}

func (o *UpdateIdentityDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /identity-providers/{id}][%d] UpdateIdentity default %s", o._statusCode, payload)
}

func (o *UpdateIdentityDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateIdentityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
