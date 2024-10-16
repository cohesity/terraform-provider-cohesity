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

	"github.com/cohesity/go-sdk/v2/models"
)

// AssignPropertiesToTenantReader is a Reader for the AssignPropertiesToTenant structure.
type AssignPropertiesToTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignPropertiesToTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignPropertiesToTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignPropertiesToTenantDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignPropertiesToTenantOK creates a AssignPropertiesToTenantOK with default headers values
func NewAssignPropertiesToTenantOK() *AssignPropertiesToTenantOK {
	return &AssignPropertiesToTenantOK{}
}

/*
AssignPropertiesToTenantOK describes a response with status code 200, with default header values.

Success
*/
type AssignPropertiesToTenantOK struct {
	Payload *models.TenantAssignments
}

// IsSuccess returns true when this assign properties to tenant o k response has a 2xx status code
func (o *AssignPropertiesToTenantOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign properties to tenant o k response has a 3xx status code
func (o *AssignPropertiesToTenantOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign properties to tenant o k response has a 4xx status code
func (o *AssignPropertiesToTenantOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign properties to tenant o k response has a 5xx status code
func (o *AssignPropertiesToTenantOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign properties to tenant o k response a status code equal to that given
func (o *AssignPropertiesToTenantOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign properties to tenant o k response
func (o *AssignPropertiesToTenantOK) Code() int {
	return 200
}

func (o *AssignPropertiesToTenantOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenants/{id}/assignments][%d] assignPropertiesToTenantOK %s", 200, payload)
}

func (o *AssignPropertiesToTenantOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenants/{id}/assignments][%d] assignPropertiesToTenantOK %s", 200, payload)
}

func (o *AssignPropertiesToTenantOK) GetPayload() *models.TenantAssignments {
	return o.Payload
}

func (o *AssignPropertiesToTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantAssignments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignPropertiesToTenantDefault creates a AssignPropertiesToTenantDefault with default headers values
func NewAssignPropertiesToTenantDefault(code int) *AssignPropertiesToTenantDefault {
	return &AssignPropertiesToTenantDefault{
		_statusCode: code,
	}
}

/*
AssignPropertiesToTenantDefault describes a response with status code -1, with default header values.

Error
*/
type AssignPropertiesToTenantDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this assign properties to tenant default response has a 2xx status code
func (o *AssignPropertiesToTenantDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign properties to tenant default response has a 3xx status code
func (o *AssignPropertiesToTenantDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign properties to tenant default response has a 4xx status code
func (o *AssignPropertiesToTenantDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign properties to tenant default response has a 5xx status code
func (o *AssignPropertiesToTenantDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign properties to tenant default response a status code equal to that given
func (o *AssignPropertiesToTenantDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the assign properties to tenant default response
func (o *AssignPropertiesToTenantDefault) Code() int {
	return o._statusCode
}

func (o *AssignPropertiesToTenantDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenants/{id}/assignments][%d] AssignPropertiesToTenant default %s", o._statusCode, payload)
}

func (o *AssignPropertiesToTenantDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenants/{id}/assignments][%d] AssignPropertiesToTenant default %s", o._statusCode, payload)
}

func (o *AssignPropertiesToTenantDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignPropertiesToTenantDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
