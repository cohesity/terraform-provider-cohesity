// Code generated by go-swagger; DO NOT EDIT.

package principals

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

// GetUserPrivilegesReader is a Reader for the GetUserPrivileges structure.
type GetUserPrivilegesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserPrivilegesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserPrivilegesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUserPrivilegesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserPrivilegesOK creates a GetUserPrivilegesOK with default headers values
func NewGetUserPrivilegesOK() *GetUserPrivilegesOK {
	return &GetUserPrivilegesOK{}
}

/*
GetUserPrivilegesOK describes a response with status code 200, with default header values.

Success
*/
type GetUserPrivilegesOK struct {
	Payload []string
}

// IsSuccess returns true when this get user privileges o k response has a 2xx status code
func (o *GetUserPrivilegesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user privileges o k response has a 3xx status code
func (o *GetUserPrivilegesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user privileges o k response has a 4xx status code
func (o *GetUserPrivilegesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user privileges o k response has a 5xx status code
func (o *GetUserPrivilegesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user privileges o k response a status code equal to that given
func (o *GetUserPrivilegesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user privileges o k response
func (o *GetUserPrivilegesOK) Code() int {
	return 200
}

func (o *GetUserPrivilegesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/users/privileges][%d] getUserPrivilegesOK %s", 200, payload)
}

func (o *GetUserPrivilegesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/users/privileges][%d] getUserPrivilegesOK %s", 200, payload)
}

func (o *GetUserPrivilegesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetUserPrivilegesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPrivilegesDefault creates a GetUserPrivilegesDefault with default headers values
func NewGetUserPrivilegesDefault(code int) *GetUserPrivilegesDefault {
	return &GetUserPrivilegesDefault{
		_statusCode: code,
	}
}

/*
GetUserPrivilegesDefault describes a response with status code -1, with default header values.

Error
*/
type GetUserPrivilegesDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get user privileges default response has a 2xx status code
func (o *GetUserPrivilegesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get user privileges default response has a 3xx status code
func (o *GetUserPrivilegesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get user privileges default response has a 4xx status code
func (o *GetUserPrivilegesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get user privileges default response has a 5xx status code
func (o *GetUserPrivilegesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get user privileges default response a status code equal to that given
func (o *GetUserPrivilegesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get user privileges default response
func (o *GetUserPrivilegesDefault) Code() int {
	return o._statusCode
}

func (o *GetUserPrivilegesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/users/privileges][%d] GetUserPrivileges default %s", o._statusCode, payload)
}

func (o *GetUserPrivilegesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/users/privileges][%d] GetUserPrivileges default %s", o._statusCode, payload)
}

func (o *GetUserPrivilegesDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetUserPrivilegesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
