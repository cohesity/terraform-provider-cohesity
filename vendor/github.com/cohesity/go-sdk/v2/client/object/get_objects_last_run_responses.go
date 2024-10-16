// Code generated by go-swagger; DO NOT EDIT.

package object

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

// GetObjectsLastRunReader is a Reader for the GetObjectsLastRun structure.
type GetObjectsLastRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetObjectsLastRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetObjectsLastRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetObjectsLastRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetObjectsLastRunOK creates a GetObjectsLastRunOK with default headers values
func NewGetObjectsLastRunOK() *GetObjectsLastRunOK {
	return &GetObjectsLastRunOK{}
}

/*
GetObjectsLastRunOK describes a response with status code 200, with default header values.

Success
*/
type GetObjectsLastRunOK struct {
	Payload *models.ObjectsLastRun
}

// IsSuccess returns true when this get objects last run o k response has a 2xx status code
func (o *GetObjectsLastRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get objects last run o k response has a 3xx status code
func (o *GetObjectsLastRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get objects last run o k response has a 4xx status code
func (o *GetObjectsLastRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get objects last run o k response has a 5xx status code
func (o *GetObjectsLastRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get objects last run o k response a status code equal to that given
func (o *GetObjectsLastRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get objects last run o k response
func (o *GetObjectsLastRunOK) Code() int {
	return 200
}

func (o *GetObjectsLastRunOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/objects/last-run][%d] getObjectsLastRunOK %s", 200, payload)
}

func (o *GetObjectsLastRunOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/objects/last-run][%d] getObjectsLastRunOK %s", 200, payload)
}

func (o *GetObjectsLastRunOK) GetPayload() *models.ObjectsLastRun {
	return o.Payload
}

func (o *GetObjectsLastRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectsLastRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObjectsLastRunDefault creates a GetObjectsLastRunDefault with default headers values
func NewGetObjectsLastRunDefault(code int) *GetObjectsLastRunDefault {
	return &GetObjectsLastRunDefault{
		_statusCode: code,
	}
}

/*
GetObjectsLastRunDefault describes a response with status code -1, with default header values.

Error
*/
type GetObjectsLastRunDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get objects last run default response has a 2xx status code
func (o *GetObjectsLastRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get objects last run default response has a 3xx status code
func (o *GetObjectsLastRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get objects last run default response has a 4xx status code
func (o *GetObjectsLastRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get objects last run default response has a 5xx status code
func (o *GetObjectsLastRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get objects last run default response a status code equal to that given
func (o *GetObjectsLastRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get objects last run default response
func (o *GetObjectsLastRunDefault) Code() int {
	return o._statusCode
}

func (o *GetObjectsLastRunDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/objects/last-run][%d] GetObjectsLastRun default %s", o._statusCode, payload)
}

func (o *GetObjectsLastRunDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/objects/last-run][%d] GetObjectsLastRun default %s", o._statusCode, payload)
}

func (o *GetObjectsLastRunDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetObjectsLastRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
