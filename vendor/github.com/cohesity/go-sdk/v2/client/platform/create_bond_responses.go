// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// CreateBondReader is a Reader for the CreateBond structure.
type CreateBondReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBondReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateBondCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateBondDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBondCreated creates a CreateBondCreated with default headers values
func NewCreateBondCreated() *CreateBondCreated {
	return &CreateBondCreated{}
}

/*
CreateBondCreated describes a response with status code 201, with default header values.

Success
*/
type CreateBondCreated struct {
	Payload *models.CreateBondParams
}

// IsSuccess returns true when this create bond created response has a 2xx status code
func (o *CreateBondCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create bond created response has a 3xx status code
func (o *CreateBondCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bond created response has a 4xx status code
func (o *CreateBondCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create bond created response has a 5xx status code
func (o *CreateBondCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create bond created response a status code equal to that given
func (o *CreateBondCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create bond created response
func (o *CreateBondCreated) Code() int {
	return 201
}

func (o *CreateBondCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/bonds][%d] createBondCreated %s", 201, payload)
}

func (o *CreateBondCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/bonds][%d] createBondCreated %s", 201, payload)
}

func (o *CreateBondCreated) GetPayload() *models.CreateBondParams {
	return o.Payload
}

func (o *CreateBondCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateBondParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBondDefault creates a CreateBondDefault with default headers values
func NewCreateBondDefault(code int) *CreateBondDefault {
	return &CreateBondDefault{
		_statusCode: code,
	}
}

/*
CreateBondDefault describes a response with status code -1, with default header values.

Error
*/
type CreateBondDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create bond default response has a 2xx status code
func (o *CreateBondDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create bond default response has a 3xx status code
func (o *CreateBondDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create bond default response has a 4xx status code
func (o *CreateBondDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create bond default response has a 5xx status code
func (o *CreateBondDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create bond default response a status code equal to that given
func (o *CreateBondDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create bond default response
func (o *CreateBondDefault) Code() int {
	return o._statusCode
}

func (o *CreateBondDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/bonds][%d] CreateBond default %s", o._statusCode, payload)
}

func (o *CreateBondDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/bonds][%d] CreateBond default %s", o._statusCode, payload)
}

func (o *CreateBondDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBondDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
