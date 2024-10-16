// Code generated by go-swagger; DO NOT EDIT.

package ldap_provider

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

// GetLdapProviderStatusReader is a Reader for the GetLdapProviderStatus structure.
type GetLdapProviderStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLdapProviderStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGetLdapProviderStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLdapProviderStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLdapProviderStatusCreated creates a GetLdapProviderStatusCreated with default headers values
func NewGetLdapProviderStatusCreated() *GetLdapProviderStatusCreated {
	return &GetLdapProviderStatusCreated{}
}

/*
GetLdapProviderStatusCreated describes a response with status code 201, with default header values.

No Content
*/
type GetLdapProviderStatusCreated struct {
}

// IsSuccess returns true when this get ldap provider status created response has a 2xx status code
func (o *GetLdapProviderStatusCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ldap provider status created response has a 3xx status code
func (o *GetLdapProviderStatusCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ldap provider status created response has a 4xx status code
func (o *GetLdapProviderStatusCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ldap provider status created response has a 5xx status code
func (o *GetLdapProviderStatusCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this get ldap provider status created response a status code equal to that given
func (o *GetLdapProviderStatusCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the get ldap provider status created response
func (o *GetLdapProviderStatusCreated) Code() int {
	return 201
}

func (o *GetLdapProviderStatusCreated) Error() string {
	return fmt.Sprintf("[GET /public/ldapProvider/{id}/status][%d] getLdapProviderStatusCreated", 201)
}

func (o *GetLdapProviderStatusCreated) String() string {
	return fmt.Sprintf("[GET /public/ldapProvider/{id}/status][%d] getLdapProviderStatusCreated", 201)
}

func (o *GetLdapProviderStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapProviderStatusDefault creates a GetLdapProviderStatusDefault with default headers values
func NewGetLdapProviderStatusDefault(code int) *GetLdapProviderStatusDefault {
	return &GetLdapProviderStatusDefault{
		_statusCode: code,
	}
}

/*
GetLdapProviderStatusDefault describes a response with status code -1, with default header values.

Error
*/
type GetLdapProviderStatusDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get ldap provider status default response has a 2xx status code
func (o *GetLdapProviderStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get ldap provider status default response has a 3xx status code
func (o *GetLdapProviderStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get ldap provider status default response has a 4xx status code
func (o *GetLdapProviderStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get ldap provider status default response has a 5xx status code
func (o *GetLdapProviderStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get ldap provider status default response a status code equal to that given
func (o *GetLdapProviderStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get ldap provider status default response
func (o *GetLdapProviderStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetLdapProviderStatusDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/ldapProvider/{id}/status][%d] GetLdapProviderStatus default %s", o._statusCode, payload)
}

func (o *GetLdapProviderStatusDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/ldapProvider/{id}/status][%d] GetLdapProviderStatus default %s", o._statusCode, payload)
}

func (o *GetLdapProviderStatusDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetLdapProviderStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
