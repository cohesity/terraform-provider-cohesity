// Code generated by go-swagger; DO NOT EDIT.

package antivirus_service_group

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

// AntivirusServiceGroupStateReader is a Reader for the AntivirusServiceGroupState structure.
type AntivirusServiceGroupStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AntivirusServiceGroupStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAntivirusServiceGroupStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAntivirusServiceGroupStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAntivirusServiceGroupStateOK creates a AntivirusServiceGroupStateOK with default headers values
func NewAntivirusServiceGroupStateOK() *AntivirusServiceGroupStateOK {
	return &AntivirusServiceGroupStateOK{}
}

/*
AntivirusServiceGroupStateOK describes a response with status code 200, with default header values.

Success
*/
type AntivirusServiceGroupStateOK struct {
	Payload *models.AntivirusServiceGroupStateParams
}

// IsSuccess returns true when this antivirus service group state o k response has a 2xx status code
func (o *AntivirusServiceGroupStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this antivirus service group state o k response has a 3xx status code
func (o *AntivirusServiceGroupStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this antivirus service group state o k response has a 4xx status code
func (o *AntivirusServiceGroupStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this antivirus service group state o k response has a 5xx status code
func (o *AntivirusServiceGroupStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this antivirus service group state o k response a status code equal to that given
func (o *AntivirusServiceGroupStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the antivirus service group state o k response
func (o *AntivirusServiceGroupStateOK) Code() int {
	return 200
}

func (o *AntivirusServiceGroupStateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups/states][%d] antivirusServiceGroupStateOK %s", 200, payload)
}

func (o *AntivirusServiceGroupStateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups/states][%d] antivirusServiceGroupStateOK %s", 200, payload)
}

func (o *AntivirusServiceGroupStateOK) GetPayload() *models.AntivirusServiceGroupStateParams {
	return o.Payload
}

func (o *AntivirusServiceGroupStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntivirusServiceGroupStateParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAntivirusServiceGroupStateDefault creates a AntivirusServiceGroupStateDefault with default headers values
func NewAntivirusServiceGroupStateDefault(code int) *AntivirusServiceGroupStateDefault {
	return &AntivirusServiceGroupStateDefault{
		_statusCode: code,
	}
}

/*
AntivirusServiceGroupStateDefault describes a response with status code -1, with default header values.

Error
*/
type AntivirusServiceGroupStateDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this antivirus service group state default response has a 2xx status code
func (o *AntivirusServiceGroupStateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this antivirus service group state default response has a 3xx status code
func (o *AntivirusServiceGroupStateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this antivirus service group state default response has a 4xx status code
func (o *AntivirusServiceGroupStateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this antivirus service group state default response has a 5xx status code
func (o *AntivirusServiceGroupStateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this antivirus service group state default response a status code equal to that given
func (o *AntivirusServiceGroupStateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the antivirus service group state default response
func (o *AntivirusServiceGroupStateDefault) Code() int {
	return o._statusCode
}

func (o *AntivirusServiceGroupStateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups/states][%d] AntivirusServiceGroupState default %s", o._statusCode, payload)
}

func (o *AntivirusServiceGroupStateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups/states][%d] AntivirusServiceGroupState default %s", o._statusCode, payload)
}

func (o *AntivirusServiceGroupStateDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *AntivirusServiceGroupStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
