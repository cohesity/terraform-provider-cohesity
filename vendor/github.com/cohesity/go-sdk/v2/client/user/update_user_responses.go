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

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*
UpdateUserOK describes a response with status code 200, with default header values.

Success
*/
type UpdateUserOK struct {
	Payload *models.UserParams
}

// IsSuccess returns true when this update user o k response has a 2xx status code
func (o *UpdateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user o k response has a 3xx status code
func (o *UpdateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user o k response has a 4xx status code
func (o *UpdateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user o k response has a 5xx status code
func (o *UpdateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user o k response a status code equal to that given
func (o *UpdateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update user o k response
func (o *UpdateUserOK) Code() int {
	return 200
}

func (o *UpdateUserOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{sid}][%d] updateUserOK %s", 200, payload)
}

func (o *UpdateUserOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{sid}][%d] updateUserOK %s", 200, payload)
}

func (o *UpdateUserOK) GetPayload() *models.UserParams {
	return o.Payload
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserDefault creates a UpdateUserDefault with default headers values
func NewUpdateUserDefault(code int) *UpdateUserDefault {
	return &UpdateUserDefault{
		_statusCode: code,
	}
}

/*
UpdateUserDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update user default response has a 2xx status code
func (o *UpdateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update user default response has a 3xx status code
func (o *UpdateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update user default response has a 4xx status code
func (o *UpdateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update user default response has a 5xx status code
func (o *UpdateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update user default response a status code equal to that given
func (o *UpdateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update user default response
func (o *UpdateUserDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{sid}][%d] UpdateUser default %s", o._statusCode, payload)
}

func (o *UpdateUserDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /users/{sid}][%d] UpdateUser default %s", o._statusCode, payload)
}

func (o *UpdateUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
