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

// GetChassisByIDReader is a Reader for the GetChassisByID structure.
type GetChassisByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChassisByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChassisByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetChassisByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetChassisByIDOK creates a GetChassisByIDOK with default headers values
func NewGetChassisByIDOK() *GetChassisByIDOK {
	return &GetChassisByIDOK{}
}

/*
GetChassisByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetChassisByIDOK struct {
	Payload *models.Chassis
}

// IsSuccess returns true when this get chassis by Id o k response has a 2xx status code
func (o *GetChassisByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get chassis by Id o k response has a 3xx status code
func (o *GetChassisByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get chassis by Id o k response has a 4xx status code
func (o *GetChassisByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get chassis by Id o k response has a 5xx status code
func (o *GetChassisByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get chassis by Id o k response a status code equal to that given
func (o *GetChassisByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get chassis by Id o k response
func (o *GetChassisByIDOK) Code() int {
	return 200
}

func (o *GetChassisByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/{id}][%d] getChassisByIdOK %s", 200, payload)
}

func (o *GetChassisByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/{id}][%d] getChassisByIdOK %s", 200, payload)
}

func (o *GetChassisByIDOK) GetPayload() *models.Chassis {
	return o.Payload
}

func (o *GetChassisByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Chassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChassisByIDDefault creates a GetChassisByIDDefault with default headers values
func NewGetChassisByIDDefault(code int) *GetChassisByIDDefault {
	return &GetChassisByIDDefault{
		_statusCode: code,
	}
}

/*
GetChassisByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetChassisByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get chassis by Id default response has a 2xx status code
func (o *GetChassisByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get chassis by Id default response has a 3xx status code
func (o *GetChassisByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get chassis by Id default response has a 4xx status code
func (o *GetChassisByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get chassis by Id default response has a 5xx status code
func (o *GetChassisByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get chassis by Id default response a status code equal to that given
func (o *GetChassisByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get chassis by Id default response
func (o *GetChassisByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetChassisByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/{id}][%d] GetChassisById default %s", o._statusCode, payload)
}

func (o *GetChassisByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /chassis/{id}][%d] GetChassisById default %s", o._statusCode, payload)
}

func (o *GetChassisByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChassisByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
