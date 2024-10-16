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

// GetBifrostTenantVlansReader is a Reader for the GetBifrostTenantVlans structure.
type GetBifrostTenantVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBifrostTenantVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBifrostTenantVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBifrostTenantVlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBifrostTenantVlansOK creates a GetBifrostTenantVlansOK with default headers values
func NewGetBifrostTenantVlansOK() *GetBifrostTenantVlansOK {
	return &GetBifrostTenantVlansOK{}
}

/*
GetBifrostTenantVlansOK describes a response with status code 200, with default header values.

Success
*/
type GetBifrostTenantVlansOK struct {
	Payload []*models.BifrostConfig
}

// IsSuccess returns true when this get bifrost tenant vlans o k response has a 2xx status code
func (o *GetBifrostTenantVlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bifrost tenant vlans o k response has a 3xx status code
func (o *GetBifrostTenantVlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bifrost tenant vlans o k response has a 4xx status code
func (o *GetBifrostTenantVlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bifrost tenant vlans o k response has a 5xx status code
func (o *GetBifrostTenantVlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bifrost tenant vlans o k response a status code equal to that given
func (o *GetBifrostTenantVlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bifrost tenant vlans o k response
func (o *GetBifrostTenantVlansOK) Code() int {
	return 200
}

func (o *GetBifrostTenantVlansOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/bifrost/vlans][%d] getBifrostTenantVlansOK %s", 200, payload)
}

func (o *GetBifrostTenantVlansOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/bifrost/vlans][%d] getBifrostTenantVlansOK %s", 200, payload)
}

func (o *GetBifrostTenantVlansOK) GetPayload() []*models.BifrostConfig {
	return o.Payload
}

func (o *GetBifrostTenantVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBifrostTenantVlansDefault creates a GetBifrostTenantVlansDefault with default headers values
func NewGetBifrostTenantVlansDefault(code int) *GetBifrostTenantVlansDefault {
	return &GetBifrostTenantVlansDefault{
		_statusCode: code,
	}
}

/*
GetBifrostTenantVlansDefault describes a response with status code -1, with default header values.

Error
*/
type GetBifrostTenantVlansDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get bifrost tenant vlans default response has a 2xx status code
func (o *GetBifrostTenantVlansDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get bifrost tenant vlans default response has a 3xx status code
func (o *GetBifrostTenantVlansDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get bifrost tenant vlans default response has a 4xx status code
func (o *GetBifrostTenantVlansDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get bifrost tenant vlans default response has a 5xx status code
func (o *GetBifrostTenantVlansDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get bifrost tenant vlans default response a status code equal to that given
func (o *GetBifrostTenantVlansDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get bifrost tenant vlans default response
func (o *GetBifrostTenantVlansDefault) Code() int {
	return o._statusCode
}

func (o *GetBifrostTenantVlansDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/bifrost/vlans][%d] GetBifrostTenantVlans default %s", o._statusCode, payload)
}

func (o *GetBifrostTenantVlansDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/bifrost/vlans][%d] GetBifrostTenantVlans default %s", o._statusCode, payload)
}

func (o *GetBifrostTenantVlansDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetBifrostTenantVlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
