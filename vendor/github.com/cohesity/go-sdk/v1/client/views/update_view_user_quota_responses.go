// Code generated by go-swagger; DO NOT EDIT.

package views

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

// UpdateViewUserQuotaReader is a Reader for the UpdateViewUserQuota structure.
type UpdateViewUserQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateViewUserQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateViewUserQuotaCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateViewUserQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateViewUserQuotaCreated creates a UpdateViewUserQuotaCreated with default headers values
func NewUpdateViewUserQuotaCreated() *UpdateViewUserQuotaCreated {
	return &UpdateViewUserQuotaCreated{}
}

/*
UpdateViewUserQuotaCreated describes a response with status code 201, with default header values.

Success
*/
type UpdateViewUserQuotaCreated struct {
	Payload *models.UserQuotaAndUsage
}

// IsSuccess returns true when this update view user quota created response has a 2xx status code
func (o *UpdateViewUserQuotaCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update view user quota created response has a 3xx status code
func (o *UpdateViewUserQuotaCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update view user quota created response has a 4xx status code
func (o *UpdateViewUserQuotaCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update view user quota created response has a 5xx status code
func (o *UpdateViewUserQuotaCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update view user quota created response a status code equal to that given
func (o *UpdateViewUserQuotaCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update view user quota created response
func (o *UpdateViewUserQuotaCreated) Code() int {
	return 201
}

func (o *UpdateViewUserQuotaCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/viewUserQuotas][%d] updateViewUserQuotaCreated %s", 201, payload)
}

func (o *UpdateViewUserQuotaCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/viewUserQuotas][%d] updateViewUserQuotaCreated %s", 201, payload)
}

func (o *UpdateViewUserQuotaCreated) GetPayload() *models.UserQuotaAndUsage {
	return o.Payload
}

func (o *UpdateViewUserQuotaCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserQuotaAndUsage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateViewUserQuotaDefault creates a UpdateViewUserQuotaDefault with default headers values
func NewUpdateViewUserQuotaDefault(code int) *UpdateViewUserQuotaDefault {
	return &UpdateViewUserQuotaDefault{
		_statusCode: code,
	}
}

/*
UpdateViewUserQuotaDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateViewUserQuotaDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update view user quota default response has a 2xx status code
func (o *UpdateViewUserQuotaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update view user quota default response has a 3xx status code
func (o *UpdateViewUserQuotaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update view user quota default response has a 4xx status code
func (o *UpdateViewUserQuotaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update view user quota default response has a 5xx status code
func (o *UpdateViewUserQuotaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update view user quota default response a status code equal to that given
func (o *UpdateViewUserQuotaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update view user quota default response
func (o *UpdateViewUserQuotaDefault) Code() int {
	return o._statusCode
}

func (o *UpdateViewUserQuotaDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/viewUserQuotas][%d] UpdateViewUserQuota default %s", o._statusCode, payload)
}

func (o *UpdateViewUserQuotaDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/viewUserQuotas][%d] UpdateViewUserQuota default %s", o._statusCode, payload)
}

func (o *UpdateViewUserQuotaDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateViewUserQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
