// Code generated by go-swagger; DO NOT EDIT.

package protection_group

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

// UpdateProtectionGroupReader is a Reader for the UpdateProtectionGroup structure.
type UpdateProtectionGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProtectionGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProtectionGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateProtectionGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProtectionGroupOK creates a UpdateProtectionGroupOK with default headers values
func NewUpdateProtectionGroupOK() *UpdateProtectionGroupOK {
	return &UpdateProtectionGroupOK{}
}

/*
UpdateProtectionGroupOK describes a response with status code 200, with default header values.

Success
*/
type UpdateProtectionGroupOK struct {
	Payload *models.ProtectionGroup
}

// IsSuccess returns true when this update protection group o k response has a 2xx status code
func (o *UpdateProtectionGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update protection group o k response has a 3xx status code
func (o *UpdateProtectionGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update protection group o k response has a 4xx status code
func (o *UpdateProtectionGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update protection group o k response has a 5xx status code
func (o *UpdateProtectionGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update protection group o k response a status code equal to that given
func (o *UpdateProtectionGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update protection group o k response
func (o *UpdateProtectionGroupOK) Code() int {
	return 200
}

func (o *UpdateProtectionGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/protection-groups/{id}][%d] updateProtectionGroupOK %s", 200, payload)
}

func (o *UpdateProtectionGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/protection-groups/{id}][%d] updateProtectionGroupOK %s", 200, payload)
}

func (o *UpdateProtectionGroupOK) GetPayload() *models.ProtectionGroup {
	return o.Payload
}

func (o *UpdateProtectionGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProtectionGroupDefault creates a UpdateProtectionGroupDefault with default headers values
func NewUpdateProtectionGroupDefault(code int) *UpdateProtectionGroupDefault {
	return &UpdateProtectionGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateProtectionGroupDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateProtectionGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update protection group default response has a 2xx status code
func (o *UpdateProtectionGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update protection group default response has a 3xx status code
func (o *UpdateProtectionGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update protection group default response has a 4xx status code
func (o *UpdateProtectionGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update protection group default response has a 5xx status code
func (o *UpdateProtectionGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update protection group default response a status code equal to that given
func (o *UpdateProtectionGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update protection group default response
func (o *UpdateProtectionGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProtectionGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/protection-groups/{id}][%d] UpdateProtectionGroup default %s", o._statusCode, payload)
}

func (o *UpdateProtectionGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/protection-groups/{id}][%d] UpdateProtectionGroup default %s", o._statusCode, payload)
}

func (o *UpdateProtectionGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateProtectionGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
