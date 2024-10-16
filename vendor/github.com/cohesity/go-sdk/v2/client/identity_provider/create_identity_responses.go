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

// CreateIdentityReader is a Reader for the CreateIdentity structure.
type CreateIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateIdentityCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateIdentityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIdentityCreated creates a CreateIdentityCreated with default headers values
func NewCreateIdentityCreated() *CreateIdentityCreated {
	return &CreateIdentityCreated{}
}

/*
CreateIdentityCreated describes a response with status code 201, with default header values.

Success
*/
type CreateIdentityCreated struct {
	Payload *models.IdentityConfig
}

// IsSuccess returns true when this create identity created response has a 2xx status code
func (o *CreateIdentityCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create identity created response has a 3xx status code
func (o *CreateIdentityCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create identity created response has a 4xx status code
func (o *CreateIdentityCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create identity created response has a 5xx status code
func (o *CreateIdentityCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create identity created response a status code equal to that given
func (o *CreateIdentityCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create identity created response
func (o *CreateIdentityCreated) Code() int {
	return 201
}

func (o *CreateIdentityCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identity-providers][%d] createIdentityCreated %s", 201, payload)
}

func (o *CreateIdentityCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identity-providers][%d] createIdentityCreated %s", 201, payload)
}

func (o *CreateIdentityCreated) GetPayload() *models.IdentityConfig {
	return o.Payload
}

func (o *CreateIdentityCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentityConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIdentityDefault creates a CreateIdentityDefault with default headers values
func NewCreateIdentityDefault(code int) *CreateIdentityDefault {
	return &CreateIdentityDefault{
		_statusCode: code,
	}
}

/*
CreateIdentityDefault describes a response with status code -1, with default header values.

Error
*/
type CreateIdentityDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create identity default response has a 2xx status code
func (o *CreateIdentityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create identity default response has a 3xx status code
func (o *CreateIdentityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create identity default response has a 4xx status code
func (o *CreateIdentityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create identity default response has a 5xx status code
func (o *CreateIdentityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create identity default response a status code equal to that given
func (o *CreateIdentityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create identity default response
func (o *CreateIdentityDefault) Code() int {
	return o._statusCode
}

func (o *CreateIdentityDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identity-providers][%d] CreateIdentity default %s", o._statusCode, payload)
}

func (o *CreateIdentityDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /identity-providers][%d] CreateIdentity default %s", o._statusCode, payload)
}

func (o *CreateIdentityDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIdentityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
