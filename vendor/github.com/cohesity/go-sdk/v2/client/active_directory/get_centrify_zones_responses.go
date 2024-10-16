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

// GetCentrifyZonesReader is a Reader for the GetCentrifyZones structure.
type GetCentrifyZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCentrifyZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCentrifyZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCentrifyZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCentrifyZonesOK creates a GetCentrifyZonesOK with default headers values
func NewGetCentrifyZonesOK() *GetCentrifyZonesOK {
	return &GetCentrifyZonesOK{}
}

/*
GetCentrifyZonesOK describes a response with status code 200, with default header values.

Success
*/
type GetCentrifyZonesOK struct {
	Payload *models.CentrifyZones
}

// IsSuccess returns true when this get centrify zones o k response has a 2xx status code
func (o *GetCentrifyZonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get centrify zones o k response has a 3xx status code
func (o *GetCentrifyZonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get centrify zones o k response has a 4xx status code
func (o *GetCentrifyZonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get centrify zones o k response has a 5xx status code
func (o *GetCentrifyZonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get centrify zones o k response a status code equal to that given
func (o *GetCentrifyZonesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get centrify zones o k response
func (o *GetCentrifyZonesOK) Code() int {
	return 200
}

func (o *GetCentrifyZonesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /centrify-zones][%d] getCentrifyZonesOK %s", 200, payload)
}

func (o *GetCentrifyZonesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /centrify-zones][%d] getCentrifyZonesOK %s", 200, payload)
}

func (o *GetCentrifyZonesOK) GetPayload() *models.CentrifyZones {
	return o.Payload
}

func (o *GetCentrifyZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CentrifyZones)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCentrifyZonesDefault creates a GetCentrifyZonesDefault with default headers values
func NewGetCentrifyZonesDefault(code int) *GetCentrifyZonesDefault {
	return &GetCentrifyZonesDefault{
		_statusCode: code,
	}
}

/*
GetCentrifyZonesDefault describes a response with status code -1, with default header values.

Error
*/
type GetCentrifyZonesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get centrify zones default response has a 2xx status code
func (o *GetCentrifyZonesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get centrify zones default response has a 3xx status code
func (o *GetCentrifyZonesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get centrify zones default response has a 4xx status code
func (o *GetCentrifyZonesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get centrify zones default response has a 5xx status code
func (o *GetCentrifyZonesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get centrify zones default response a status code equal to that given
func (o *GetCentrifyZonesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get centrify zones default response
func (o *GetCentrifyZonesDefault) Code() int {
	return o._statusCode
}

func (o *GetCentrifyZonesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /centrify-zones][%d] GetCentrifyZones default %s", o._statusCode, payload)
}

func (o *GetCentrifyZonesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /centrify-zones][%d] GetCentrifyZones default %s", o._statusCode, payload)
}

func (o *GetCentrifyZonesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCentrifyZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
