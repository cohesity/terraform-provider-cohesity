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

// UpdateGroupReader is a Reader for the UpdateGroup structure.
type UpdateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateGroupOK creates a UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {
	return &UpdateGroupOK{}
}

/*
UpdateGroupOK describes a response with status code 200, with default header values.

Success
*/
type UpdateGroupOK struct {
	Payload *models.GroupParams
}

// IsSuccess returns true when this update group o k response has a 2xx status code
func (o *UpdateGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update group o k response has a 3xx status code
func (o *UpdateGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update group o k response has a 4xx status code
func (o *UpdateGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update group o k response has a 5xx status code
func (o *UpdateGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update group o k response a status code equal to that given
func (o *UpdateGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update group o k response
func (o *UpdateGroupOK) Code() int {
	return 200
}

func (o *UpdateGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groups/{sid}][%d] updateGroupOK %s", 200, payload)
}

func (o *UpdateGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groups/{sid}][%d] updateGroupOK %s", 200, payload)
}

func (o *UpdateGroupOK) GetPayload() *models.GroupParams {
	return o.Payload
}

func (o *UpdateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupDefault creates a UpdateGroupDefault with default headers values
func NewUpdateGroupDefault(code int) *UpdateGroupDefault {
	return &UpdateGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateGroupDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update group default response has a 2xx status code
func (o *UpdateGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update group default response has a 3xx status code
func (o *UpdateGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update group default response has a 4xx status code
func (o *UpdateGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update group default response has a 5xx status code
func (o *UpdateGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update group default response a status code equal to that given
func (o *UpdateGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update group default response
func (o *UpdateGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groups/{sid}][%d] UpdateGroup default %s", o._statusCode, payload)
}

func (o *UpdateGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /groups/{sid}][%d] UpdateGroup default %s", o._statusCode, payload)
}

func (o *UpdateGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
