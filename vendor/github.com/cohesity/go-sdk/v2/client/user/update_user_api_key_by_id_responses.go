// Code generated by go-swagger; DO NOT EDIT.

package user

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

// UpdateUserAPIKeyByIDReader is a Reader for the UpdateUserAPIKeyByID structure.
type UpdateUserAPIKeyByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserAPIKeyByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserAPIKeyByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateUserAPIKeyByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserAPIKeyByIDOK creates a UpdateUserAPIKeyByIDOK with default headers values
func NewUpdateUserAPIKeyByIDOK() *UpdateUserAPIKeyByIDOK {
	return &UpdateUserAPIKeyByIDOK{}
}

/*
UpdateUserAPIKeyByIDOK describes a response with status code 200, with default header values.

Success
*/
type UpdateUserAPIKeyByIDOK struct {
	Payload *models.UserAPIKey
}

// IsSuccess returns true when this update user Api key by Id o k response has a 2xx status code
func (o *UpdateUserAPIKeyByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user Api key by Id o k response has a 3xx status code
func (o *UpdateUserAPIKeyByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user Api key by Id o k response has a 4xx status code
func (o *UpdateUserAPIKeyByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user Api key by Id o k response has a 5xx status code
func (o *UpdateUserAPIKeyByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user Api key by Id o k response a status code equal to that given
func (o *UpdateUserAPIKeyByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update user Api key by Id o k response
func (o *UpdateUserAPIKeyByIDOK) Code() int {
	return 200
}

func (o *UpdateUserAPIKeyByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{userSid}/api-keys/{id}][%d] updateUserApiKeyByIdOK %s", 200, payload)
}

func (o *UpdateUserAPIKeyByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{userSid}/api-keys/{id}][%d] updateUserApiKeyByIdOK %s", 200, payload)
}

func (o *UpdateUserAPIKeyByIDOK) GetPayload() *models.UserAPIKey {
	return o.Payload
}

func (o *UpdateUserAPIKeyByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAPIKeyByIDDefault creates a UpdateUserAPIKeyByIDDefault with default headers values
func NewUpdateUserAPIKeyByIDDefault(code int) *UpdateUserAPIKeyByIDDefault {
	return &UpdateUserAPIKeyByIDDefault{
		_statusCode: code,
	}
}

/*
UpdateUserAPIKeyByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateUserAPIKeyByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update user API key by Id default response has a 2xx status code
func (o *UpdateUserAPIKeyByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update user API key by Id default response has a 3xx status code
func (o *UpdateUserAPIKeyByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update user API key by Id default response has a 4xx status code
func (o *UpdateUserAPIKeyByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update user API key by Id default response has a 5xx status code
func (o *UpdateUserAPIKeyByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update user API key by Id default response a status code equal to that given
func (o *UpdateUserAPIKeyByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update user API key by Id default response
func (o *UpdateUserAPIKeyByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserAPIKeyByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{userSid}/api-keys/{id}][%d] UpdateUserAPIKeyById default %s", o._statusCode, payload)
}

func (o *UpdateUserAPIKeyByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{userSid}/api-keys/{id}][%d] UpdateUserAPIKeyById default %s", o._statusCode, payload)
}

func (o *UpdateUserAPIKeyByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserAPIKeyByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
