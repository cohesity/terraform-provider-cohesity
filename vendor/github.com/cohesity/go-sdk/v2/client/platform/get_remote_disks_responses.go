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

// GetRemoteDisksReader is a Reader for the GetRemoteDisks structure.
type GetRemoteDisksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemoteDisksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRemoteDisksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRemoteDisksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRemoteDisksOK creates a GetRemoteDisksOK with default headers values
func NewGetRemoteDisksOK() *GetRemoteDisksOK {
	return &GetRemoteDisksOK{}
}

/*
GetRemoteDisksOK describes a response with status code 200, with default header values.

Success
*/
type GetRemoteDisksOK struct {
	Payload *models.RemoteDisks
}

// IsSuccess returns true when this get remote disks o k response has a 2xx status code
func (o *GetRemoteDisksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get remote disks o k response has a 3xx status code
func (o *GetRemoteDisksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote disks o k response has a 4xx status code
func (o *GetRemoteDisksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get remote disks o k response has a 5xx status code
func (o *GetRemoteDisksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote disks o k response a status code equal to that given
func (o *GetRemoteDisksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get remote disks o k response
func (o *GetRemoteDisksOK) Code() int {
	return 200
}

func (o *GetRemoteDisksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /disks/remote][%d] getRemoteDisksOK %s", 200, payload)
}

func (o *GetRemoteDisksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /disks/remote][%d] getRemoteDisksOK %s", 200, payload)
}

func (o *GetRemoteDisksOK) GetPayload() *models.RemoteDisks {
	return o.Payload
}

func (o *GetRemoteDisksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteDisks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteDisksDefault creates a GetRemoteDisksDefault with default headers values
func NewGetRemoteDisksDefault(code int) *GetRemoteDisksDefault {
	return &GetRemoteDisksDefault{
		_statusCode: code,
	}
}

/*
GetRemoteDisksDefault describes a response with status code -1, with default header values.

Error
*/
type GetRemoteDisksDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get remote disks default response has a 2xx status code
func (o *GetRemoteDisksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get remote disks default response has a 3xx status code
func (o *GetRemoteDisksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get remote disks default response has a 4xx status code
func (o *GetRemoteDisksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get remote disks default response has a 5xx status code
func (o *GetRemoteDisksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get remote disks default response a status code equal to that given
func (o *GetRemoteDisksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get remote disks default response
func (o *GetRemoteDisksDefault) Code() int {
	return o._statusCode
}

func (o *GetRemoteDisksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /disks/remote][%d] GetRemoteDisks default %s", o._statusCode, payload)
}

func (o *GetRemoteDisksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /disks/remote][%d] GetRemoteDisks default %s", o._statusCode, payload)
}

func (o *GetRemoteDisksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRemoteDisksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
