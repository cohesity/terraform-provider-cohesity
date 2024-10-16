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

// ListTenantsReader is a Reader for the ListTenants structure.
type ListTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListTenantsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTenantsOK creates a ListTenantsOK with default headers values
func NewListTenantsOK() *ListTenantsOK {
	return &ListTenantsOK{}
}

/*
ListTenantsOK describes a response with status code 200, with default header values.

Success
*/
type ListTenantsOK struct {
	Payload models.TenantsInfo
}

// IsSuccess returns true when this list tenants o k response has a 2xx status code
func (o *ListTenantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list tenants o k response has a 3xx status code
func (o *ListTenantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list tenants o k response has a 4xx status code
func (o *ListTenantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list tenants o k response has a 5xx status code
func (o *ListTenantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list tenants o k response a status code equal to that given
func (o *ListTenantsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list tenants o k response
func (o *ListTenantsOK) Code() int {
	return 200
}

func (o *ListTenantsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants][%d] listTenantsOK %s", 200, payload)
}

func (o *ListTenantsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants][%d] listTenantsOK %s", 200, payload)
}

func (o *ListTenantsOK) GetPayload() models.TenantsInfo {
	return o.Payload
}

func (o *ListTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTenantsDefault creates a ListTenantsDefault with default headers values
func NewListTenantsDefault(code int) *ListTenantsDefault {
	return &ListTenantsDefault{
		_statusCode: code,
	}
}

/*
ListTenantsDefault describes a response with status code -1, with default header values.

Error
*/
type ListTenantsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list tenants default response has a 2xx status code
func (o *ListTenantsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list tenants default response has a 3xx status code
func (o *ListTenantsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list tenants default response has a 4xx status code
func (o *ListTenantsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list tenants default response has a 5xx status code
func (o *ListTenantsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list tenants default response a status code equal to that given
func (o *ListTenantsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list tenants default response
func (o *ListTenantsDefault) Code() int {
	return o._statusCode
}

func (o *ListTenantsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants][%d] ListTenants default %s", o._statusCode, payload)
}

func (o *ListTenantsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants][%d] ListTenants default %s", o._statusCode, payload)
}

func (o *ListTenantsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListTenantsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
