// Code generated by go-swagger; DO NOT EDIT.

package security

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

// CreateCsrReader is a Reader for the CreateCsr structure.
type CreateCsrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCsrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCsrCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateCsrDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCsrCreated creates a CreateCsrCreated with default headers values
func NewCreateCsrCreated() *CreateCsrCreated {
	return &CreateCsrCreated{}
}

/*
CreateCsrCreated describes a response with status code 201, with default header values.

Success
*/
type CreateCsrCreated struct {
	Payload *models.CreateCsrResponseBody
}

// IsSuccess returns true when this create csr created response has a 2xx status code
func (o *CreateCsrCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create csr created response has a 3xx status code
func (o *CreateCsrCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create csr created response has a 4xx status code
func (o *CreateCsrCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create csr created response has a 5xx status code
func (o *CreateCsrCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create csr created response a status code equal to that given
func (o *CreateCsrCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create csr created response
func (o *CreateCsrCreated) Code() int {
	return 201
}

func (o *CreateCsrCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /csr][%d] createCsrCreated %s", 201, payload)
}

func (o *CreateCsrCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /csr][%d] createCsrCreated %s", 201, payload)
}

func (o *CreateCsrCreated) GetPayload() *models.CreateCsrResponseBody {
	return o.Payload
}

func (o *CreateCsrCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCsrResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCsrDefault creates a CreateCsrDefault with default headers values
func NewCreateCsrDefault(code int) *CreateCsrDefault {
	return &CreateCsrDefault{
		_statusCode: code,
	}
}

/*
CreateCsrDefault describes a response with status code -1, with default header values.

Error
*/
type CreateCsrDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create csr default response has a 2xx status code
func (o *CreateCsrDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create csr default response has a 3xx status code
func (o *CreateCsrDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create csr default response has a 4xx status code
func (o *CreateCsrDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create csr default response has a 5xx status code
func (o *CreateCsrDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create csr default response a status code equal to that given
func (o *CreateCsrDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create csr default response
func (o *CreateCsrDefault) Code() int {
	return o._statusCode
}

func (o *CreateCsrDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /csr][%d] CreateCsr default %s", o._statusCode, payload)
}

func (o *CreateCsrDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /csr][%d] CreateCsr default %s", o._statusCode, payload)
}

func (o *CreateCsrDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCsrDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
