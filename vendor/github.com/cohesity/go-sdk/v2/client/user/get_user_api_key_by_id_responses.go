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

// GetUserAPIKeyByIDReader is a Reader for the GetUserAPIKeyByID structure.
type GetUserAPIKeyByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserAPIKeyByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserAPIKeyByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUserAPIKeyByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserAPIKeyByIDOK creates a GetUserAPIKeyByIDOK with default headers values
func NewGetUserAPIKeyByIDOK() *GetUserAPIKeyByIDOK {
	return &GetUserAPIKeyByIDOK{}
}

/*
GetUserAPIKeyByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetUserAPIKeyByIDOK struct {
	Payload *models.UserAPIKey
}

// IsSuccess returns true when this get user Api key by Id o k response has a 2xx status code
func (o *GetUserAPIKeyByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user Api key by Id o k response has a 3xx status code
func (o *GetUserAPIKeyByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user Api key by Id o k response has a 4xx status code
func (o *GetUserAPIKeyByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user Api key by Id o k response has a 5xx status code
func (o *GetUserAPIKeyByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user Api key by Id o k response a status code equal to that given
func (o *GetUserAPIKeyByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user Api key by Id o k response
func (o *GetUserAPIKeyByIDOK) Code() int {
	return 200
}

func (o *GetUserAPIKeyByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{userSid}/api-keys/{id}][%d] getUserApiKeyByIdOK %s", 200, payload)
}

func (o *GetUserAPIKeyByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{userSid}/api-keys/{id}][%d] getUserApiKeyByIdOK %s", 200, payload)
}

func (o *GetUserAPIKeyByIDOK) GetPayload() *models.UserAPIKey {
	return o.Payload
}

func (o *GetUserAPIKeyByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserAPIKeyByIDDefault creates a GetUserAPIKeyByIDDefault with default headers values
func NewGetUserAPIKeyByIDDefault(code int) *GetUserAPIKeyByIDDefault {
	return &GetUserAPIKeyByIDDefault{
		_statusCode: code,
	}
}

/*
GetUserAPIKeyByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetUserAPIKeyByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get user API key by Id default response has a 2xx status code
func (o *GetUserAPIKeyByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get user API key by Id default response has a 3xx status code
func (o *GetUserAPIKeyByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get user API key by Id default response has a 4xx status code
func (o *GetUserAPIKeyByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get user API key by Id default response has a 5xx status code
func (o *GetUserAPIKeyByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get user API key by Id default response a status code equal to that given
func (o *GetUserAPIKeyByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get user API key by Id default response
func (o *GetUserAPIKeyByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetUserAPIKeyByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{userSid}/api-keys/{id}][%d] GetUserAPIKeyById default %s", o._statusCode, payload)
}

func (o *GetUserAPIKeyByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{userSid}/api-keys/{id}][%d] GetUserAPIKeyById default %s", o._statusCode, payload)
}

func (o *GetUserAPIKeyByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUserAPIKeyByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
