// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

// GetActiveDirectoryPrincipalsReader is a Reader for the GetActiveDirectoryPrincipals structure.
type GetActiveDirectoryPrincipalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActiveDirectoryPrincipalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActiveDirectoryPrincipalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetActiveDirectoryPrincipalsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetActiveDirectoryPrincipalsOK creates a GetActiveDirectoryPrincipalsOK with default headers values
func NewGetActiveDirectoryPrincipalsOK() *GetActiveDirectoryPrincipalsOK {
	return &GetActiveDirectoryPrincipalsOK{}
}

/*
GetActiveDirectoryPrincipalsOK describes a response with status code 200, with default header values.

Success
*/
type GetActiveDirectoryPrincipalsOK struct {
	Payload *models.ActiveDirectoryPrincipals
}

// IsSuccess returns true when this get active directory principals o k response has a 2xx status code
func (o *GetActiveDirectoryPrincipalsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get active directory principals o k response has a 3xx status code
func (o *GetActiveDirectoryPrincipalsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get active directory principals o k response has a 4xx status code
func (o *GetActiveDirectoryPrincipalsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get active directory principals o k response has a 5xx status code
func (o *GetActiveDirectoryPrincipalsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get active directory principals o k response a status code equal to that given
func (o *GetActiveDirectoryPrincipalsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get active directory principals o k response
func (o *GetActiveDirectoryPrincipalsOK) Code() int {
	return 200
}

func (o *GetActiveDirectoryPrincipalsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /active-directory-principals][%d] getActiveDirectoryPrincipalsOK %s", 200, payload)
}

func (o *GetActiveDirectoryPrincipalsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /active-directory-principals][%d] getActiveDirectoryPrincipalsOK %s", 200, payload)
}

func (o *GetActiveDirectoryPrincipalsOK) GetPayload() *models.ActiveDirectoryPrincipals {
	return o.Payload
}

func (o *GetActiveDirectoryPrincipalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveDirectoryPrincipals)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActiveDirectoryPrincipalsDefault creates a GetActiveDirectoryPrincipalsDefault with default headers values
func NewGetActiveDirectoryPrincipalsDefault(code int) *GetActiveDirectoryPrincipalsDefault {
	return &GetActiveDirectoryPrincipalsDefault{
		_statusCode: code,
	}
}

/*
GetActiveDirectoryPrincipalsDefault describes a response with status code -1, with default header values.

Error
*/
type GetActiveDirectoryPrincipalsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get active directory principals default response has a 2xx status code
func (o *GetActiveDirectoryPrincipalsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get active directory principals default response has a 3xx status code
func (o *GetActiveDirectoryPrincipalsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get active directory principals default response has a 4xx status code
func (o *GetActiveDirectoryPrincipalsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get active directory principals default response has a 5xx status code
func (o *GetActiveDirectoryPrincipalsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get active directory principals default response a status code equal to that given
func (o *GetActiveDirectoryPrincipalsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get active directory principals default response
func (o *GetActiveDirectoryPrincipalsDefault) Code() int {
	return o._statusCode
}

func (o *GetActiveDirectoryPrincipalsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /active-directory-principals][%d] GetActiveDirectoryPrincipals default %s", o._statusCode, payload)
}

func (o *GetActiveDirectoryPrincipalsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /active-directory-principals][%d] GetActiveDirectoryPrincipals default %s", o._statusCode, payload)
}

func (o *GetActiveDirectoryPrincipalsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetActiveDirectoryPrincipalsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
