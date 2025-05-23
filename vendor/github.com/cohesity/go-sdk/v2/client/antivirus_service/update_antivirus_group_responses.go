// Code generated by go-swagger; DO NOT EDIT.

package antivirus_service

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

// UpdateAntivirusGroupReader is a Reader for the UpdateAntivirusGroup structure.
type UpdateAntivirusGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAntivirusGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAntivirusGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAntivirusGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAntivirusGroupOK creates a UpdateAntivirusGroupOK with default headers values
func NewUpdateAntivirusGroupOK() *UpdateAntivirusGroupOK {
	return &UpdateAntivirusGroupOK{}
}

/*
UpdateAntivirusGroupOK describes a response with status code 200, with default header values.

Success
*/
type UpdateAntivirusGroupOK struct {
	Payload *models.AntivirusServiceGroup
}

// IsSuccess returns true when this update antivirus group o k response has a 2xx status code
func (o *UpdateAntivirusGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update antivirus group o k response has a 3xx status code
func (o *UpdateAntivirusGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update antivirus group o k response has a 4xx status code
func (o *UpdateAntivirusGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update antivirus group o k response has a 5xx status code
func (o *UpdateAntivirusGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update antivirus group o k response a status code equal to that given
func (o *UpdateAntivirusGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update antivirus group o k response
func (o *UpdateAntivirusGroupOK) Code() int {
	return 200
}

func (o *UpdateAntivirusGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /antivirus-service/groups/{id}][%d] updateAntivirusGroupOK %s", 200, payload)
}

func (o *UpdateAntivirusGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /antivirus-service/groups/{id}][%d] updateAntivirusGroupOK %s", 200, payload)
}

func (o *UpdateAntivirusGroupOK) GetPayload() *models.AntivirusServiceGroup {
	return o.Payload
}

func (o *UpdateAntivirusGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntivirusServiceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAntivirusGroupDefault creates a UpdateAntivirusGroupDefault with default headers values
func NewUpdateAntivirusGroupDefault(code int) *UpdateAntivirusGroupDefault {
	return &UpdateAntivirusGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateAntivirusGroupDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateAntivirusGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update antivirus group default response has a 2xx status code
func (o *UpdateAntivirusGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update antivirus group default response has a 3xx status code
func (o *UpdateAntivirusGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update antivirus group default response has a 4xx status code
func (o *UpdateAntivirusGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update antivirus group default response has a 5xx status code
func (o *UpdateAntivirusGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update antivirus group default response a status code equal to that given
func (o *UpdateAntivirusGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update antivirus group default response
func (o *UpdateAntivirusGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAntivirusGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /antivirus-service/groups/{id}][%d] UpdateAntivirusGroup default %s", o._statusCode, payload)
}

func (o *UpdateAntivirusGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /antivirus-service/groups/{id}][%d] UpdateAntivirusGroup default %s", o._statusCode, payload)
}

func (o *UpdateAntivirusGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAntivirusGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
