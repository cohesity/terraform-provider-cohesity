// Code generated by go-swagger; DO NOT EDIT.

package statistics

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

// GetProgressMonitorByPathReader is a Reader for the GetProgressMonitorByPath structure.
type GetProgressMonitorByPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProgressMonitorByPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProgressMonitorByPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProgressMonitorByPathDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProgressMonitorByPathOK creates a GetProgressMonitorByPathOK with default headers values
func NewGetProgressMonitorByPathOK() *GetProgressMonitorByPathOK {
	return &GetProgressMonitorByPathOK{}
}

/*
GetProgressMonitorByPathOK describes a response with status code 200, with default header values.

Success
*/
type GetProgressMonitorByPathOK struct {
	Payload *models.GetTasksResult
}

// IsSuccess returns true when this get progress monitor by path o k response has a 2xx status code
func (o *GetProgressMonitorByPathOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get progress monitor by path o k response has a 3xx status code
func (o *GetProgressMonitorByPathOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get progress monitor by path o k response has a 4xx status code
func (o *GetProgressMonitorByPathOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get progress monitor by path o k response has a 5xx status code
func (o *GetProgressMonitorByPathOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get progress monitor by path o k response a status code equal to that given
func (o *GetProgressMonitorByPathOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get progress monitor by path o k response
func (o *GetProgressMonitorByPathOK) Code() int {
	return 200
}

func (o *GetProgressMonitorByPathOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /progressMonitors/{path}][%d] getProgressMonitorByPathOK %s", 200, payload)
}

func (o *GetProgressMonitorByPathOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /progressMonitors/{path}][%d] getProgressMonitorByPathOK %s", 200, payload)
}

func (o *GetProgressMonitorByPathOK) GetPayload() *models.GetTasksResult {
	return o.Payload
}

func (o *GetProgressMonitorByPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTasksResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProgressMonitorByPathDefault creates a GetProgressMonitorByPathDefault with default headers values
func NewGetProgressMonitorByPathDefault(code int) *GetProgressMonitorByPathDefault {
	return &GetProgressMonitorByPathDefault{
		_statusCode: code,
	}
}

/*
GetProgressMonitorByPathDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetProgressMonitorByPathDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get progress monitor by path default response has a 2xx status code
func (o *GetProgressMonitorByPathDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get progress monitor by path default response has a 3xx status code
func (o *GetProgressMonitorByPathDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get progress monitor by path default response has a 4xx status code
func (o *GetProgressMonitorByPathDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get progress monitor by path default response has a 5xx status code
func (o *GetProgressMonitorByPathDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get progress monitor by path default response a status code equal to that given
func (o *GetProgressMonitorByPathDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get progress monitor by path default response
func (o *GetProgressMonitorByPathDefault) Code() int {
	return o._statusCode
}

func (o *GetProgressMonitorByPathDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /progressMonitors/{path}][%d] GetProgressMonitorByPath default %s", o._statusCode, payload)
}

func (o *GetProgressMonitorByPathDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /progressMonitors/{path}][%d] GetProgressMonitorByPath default %s", o._statusCode, payload)
}

func (o *GetProgressMonitorByPathDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetProgressMonitorByPathDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
