// Code generated by go-swagger; DO NOT EDIT.

package vlan

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

// CreateBifrostVlanReader is a Reader for the CreateBifrostVlan structure.
type CreateBifrostVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBifrostVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBifrostVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateBifrostVlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBifrostVlanOK creates a CreateBifrostVlanOK with default headers values
func NewCreateBifrostVlanOK() *CreateBifrostVlanOK {
	return &CreateBifrostVlanOK{}
}

/*
CreateBifrostVlanOK describes a response with status code 200, with default header values.

Success
*/
type CreateBifrostVlanOK struct {
	Payload *models.BifrostConfig
}

// IsSuccess returns true when this create bifrost vlan o k response has a 2xx status code
func (o *CreateBifrostVlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create bifrost vlan o k response has a 3xx status code
func (o *CreateBifrostVlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bifrost vlan o k response has a 4xx status code
func (o *CreateBifrostVlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create bifrost vlan o k response has a 5xx status code
func (o *CreateBifrostVlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create bifrost vlan o k response a status code equal to that given
func (o *CreateBifrostVlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create bifrost vlan o k response
func (o *CreateBifrostVlanOK) Code() int {
	return 200
}

func (o *CreateBifrostVlanOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/bifrost/vlans][%d] createBifrostVlanOK %s", 200, payload)
}

func (o *CreateBifrostVlanOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/bifrost/vlans][%d] createBifrostVlanOK %s", 200, payload)
}

func (o *CreateBifrostVlanOK) GetPayload() *models.BifrostConfig {
	return o.Payload
}

func (o *CreateBifrostVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BifrostConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBifrostVlanDefault creates a CreateBifrostVlanDefault with default headers values
func NewCreateBifrostVlanDefault(code int) *CreateBifrostVlanDefault {
	return &CreateBifrostVlanDefault{
		_statusCode: code,
	}
}

/*
CreateBifrostVlanDefault describes a response with status code -1, with default header values.

Error
*/
type CreateBifrostVlanDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create bifrost vlan default response has a 2xx status code
func (o *CreateBifrostVlanDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create bifrost vlan default response has a 3xx status code
func (o *CreateBifrostVlanDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create bifrost vlan default response has a 4xx status code
func (o *CreateBifrostVlanDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create bifrost vlan default response has a 5xx status code
func (o *CreateBifrostVlanDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create bifrost vlan default response a status code equal to that given
func (o *CreateBifrostVlanDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create bifrost vlan default response
func (o *CreateBifrostVlanDefault) Code() int {
	return o._statusCode
}

func (o *CreateBifrostVlanDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/bifrost/vlans][%d] CreateBifrostVlan default %s", o._statusCode, payload)
}

func (o *CreateBifrostVlanDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/bifrost/vlans][%d] CreateBifrostVlan default %s", o._statusCode, payload)
}

func (o *CreateBifrostVlanDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateBifrostVlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
