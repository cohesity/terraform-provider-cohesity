// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// GetBackgroundActivityScheduleReader is a Reader for the GetBackgroundActivitySchedule structure.
type GetBackgroundActivityScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackgroundActivityScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackgroundActivityScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackgroundActivityScheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackgroundActivityScheduleOK creates a GetBackgroundActivityScheduleOK with default headers values
func NewGetBackgroundActivityScheduleOK() *GetBackgroundActivityScheduleOK {
	return &GetBackgroundActivityScheduleOK{}
}

/*
GetBackgroundActivityScheduleOK describes a response with status code 200, with default header values.

Success
*/
type GetBackgroundActivityScheduleOK struct {
	Payload *models.BandwidthLimit
}

// IsSuccess returns true when this get background activity schedule o k response has a 2xx status code
func (o *GetBackgroundActivityScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get background activity schedule o k response has a 3xx status code
func (o *GetBackgroundActivityScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get background activity schedule o k response has a 4xx status code
func (o *GetBackgroundActivityScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get background activity schedule o k response has a 5xx status code
func (o *GetBackgroundActivityScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get background activity schedule o k response a status code equal to that given
func (o *GetBackgroundActivityScheduleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get background activity schedule o k response
func (o *GetBackgroundActivityScheduleOK) Code() int {
	return 200
}

func (o *GetBackgroundActivityScheduleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/cluster/backgroundActivitySchedule][%d] getBackgroundActivityScheduleOK %s", 200, payload)
}

func (o *GetBackgroundActivityScheduleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/cluster/backgroundActivitySchedule][%d] getBackgroundActivityScheduleOK %s", 200, payload)
}

func (o *GetBackgroundActivityScheduleOK) GetPayload() *models.BandwidthLimit {
	return o.Payload
}

func (o *GetBackgroundActivityScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BandwidthLimit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackgroundActivityScheduleDefault creates a GetBackgroundActivityScheduleDefault with default headers values
func NewGetBackgroundActivityScheduleDefault(code int) *GetBackgroundActivityScheduleDefault {
	return &GetBackgroundActivityScheduleDefault{
		_statusCode: code,
	}
}

/*
GetBackgroundActivityScheduleDefault describes a response with status code -1, with default header values.

Error
*/
type GetBackgroundActivityScheduleDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get background activity schedule default response has a 2xx status code
func (o *GetBackgroundActivityScheduleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get background activity schedule default response has a 3xx status code
func (o *GetBackgroundActivityScheduleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get background activity schedule default response has a 4xx status code
func (o *GetBackgroundActivityScheduleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get background activity schedule default response has a 5xx status code
func (o *GetBackgroundActivityScheduleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get background activity schedule default response a status code equal to that given
func (o *GetBackgroundActivityScheduleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get background activity schedule default response
func (o *GetBackgroundActivityScheduleDefault) Code() int {
	return o._statusCode
}

func (o *GetBackgroundActivityScheduleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/cluster/backgroundActivitySchedule][%d] GetBackgroundActivitySchedule default %s", o._statusCode, payload)
}

func (o *GetBackgroundActivityScheduleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/cluster/backgroundActivitySchedule][%d] GetBackgroundActivitySchedule default %s", o._statusCode, payload)
}

func (o *GetBackgroundActivityScheduleDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetBackgroundActivityScheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
