// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// UpdateTenantViewReader is a Reader for the UpdateTenantView structure.
type UpdateTenantViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTenantViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTenantViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateTenantViewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTenantViewOK creates a UpdateTenantViewOK with default headers values
func NewUpdateTenantViewOK() *UpdateTenantViewOK {
	return &UpdateTenantViewOK{}
}

/*
UpdateTenantViewOK describes a response with status code 200, with default header values.

Tenant View Mapping Response.
*/
type UpdateTenantViewOK struct {
	Payload *models.TenantViewUpdate
}

// IsSuccess returns true when this update tenant view o k response has a 2xx status code
func (o *UpdateTenantViewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update tenant view o k response has a 3xx status code
func (o *UpdateTenantViewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tenant view o k response has a 4xx status code
func (o *UpdateTenantViewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update tenant view o k response has a 5xx status code
func (o *UpdateTenantViewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update tenant view o k response a status code equal to that given
func (o *UpdateTenantViewOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update tenant view o k response
func (o *UpdateTenantViewOK) Code() int {
	return 200
}

func (o *UpdateTenantViewOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/view][%d] updateTenantViewOK %s", 200, payload)
}

func (o *UpdateTenantViewOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/view][%d] updateTenantViewOK %s", 200, payload)
}

func (o *UpdateTenantViewOK) GetPayload() *models.TenantViewUpdate {
	return o.Payload
}

func (o *UpdateTenantViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantViewUpdate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTenantViewDefault creates a UpdateTenantViewDefault with default headers values
func NewUpdateTenantViewDefault(code int) *UpdateTenantViewDefault {
	return &UpdateTenantViewDefault{
		_statusCode: code,
	}
}

/*
UpdateTenantViewDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateTenantViewDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update tenant view default response has a 2xx status code
func (o *UpdateTenantViewDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update tenant view default response has a 3xx status code
func (o *UpdateTenantViewDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update tenant view default response has a 4xx status code
func (o *UpdateTenantViewDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update tenant view default response has a 5xx status code
func (o *UpdateTenantViewDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update tenant view default response a status code equal to that given
func (o *UpdateTenantViewDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update tenant view default response
func (o *UpdateTenantViewDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTenantViewDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/view][%d] UpdateTenantView default %s", o._statusCode, payload)
}

func (o *UpdateTenantViewDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/view][%d] UpdateTenantView default %s", o._statusCode, payload)
}

func (o *UpdateTenantViewDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateTenantViewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
