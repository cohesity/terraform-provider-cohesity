// Code generated by go-swagger; DO NOT EDIT.

package protection_group

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

// GetRunDebugLogsReader is a Reader for the GetRunDebugLogs structure.
type GetRunDebugLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunDebugLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunDebugLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRunDebugLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunDebugLogsOK creates a GetRunDebugLogsOK with default headers values
func NewGetRunDebugLogsOK() *GetRunDebugLogsOK {
	return &GetRunDebugLogsOK{}
}

/*
GetRunDebugLogsOK describes a response with status code 200, with default header values.

No Content
*/
type GetRunDebugLogsOK struct {
}

// IsSuccess returns true when this get run debug logs o k response has a 2xx status code
func (o *GetRunDebugLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run debug logs o k response has a 3xx status code
func (o *GetRunDebugLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run debug logs o k response has a 4xx status code
func (o *GetRunDebugLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run debug logs o k response has a 5xx status code
func (o *GetRunDebugLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get run debug logs o k response a status code equal to that given
func (o *GetRunDebugLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get run debug logs o k response
func (o *GetRunDebugLogsOK) Code() int {
	return 200
}

func (o *GetRunDebugLogsOK) Error() string {
	return fmt.Sprintf("[GET /data-protect/protection-groups/{id}/runs/{runId}/debug-logs][%d] getRunDebugLogsOK", 200)
}

func (o *GetRunDebugLogsOK) String() string {
	return fmt.Sprintf("[GET /data-protect/protection-groups/{id}/runs/{runId}/debug-logs][%d] getRunDebugLogsOK", 200)
}

func (o *GetRunDebugLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRunDebugLogsDefault creates a GetRunDebugLogsDefault with default headers values
func NewGetRunDebugLogsDefault(code int) *GetRunDebugLogsDefault {
	return &GetRunDebugLogsDefault{
		_statusCode: code,
	}
}

/*
GetRunDebugLogsDefault describes a response with status code -1, with default header values.

Error
*/
type GetRunDebugLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get run debug logs default response has a 2xx status code
func (o *GetRunDebugLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get run debug logs default response has a 3xx status code
func (o *GetRunDebugLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get run debug logs default response has a 4xx status code
func (o *GetRunDebugLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get run debug logs default response has a 5xx status code
func (o *GetRunDebugLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get run debug logs default response a status code equal to that given
func (o *GetRunDebugLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get run debug logs default response
func (o *GetRunDebugLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetRunDebugLogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/protection-groups/{id}/runs/{runId}/debug-logs][%d] GetRunDebugLogs default %s", o._statusCode, payload)
}

func (o *GetRunDebugLogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/protection-groups/{id}/runs/{runId}/debug-logs][%d] GetRunDebugLogs default %s", o._statusCode, payload)
}

func (o *GetRunDebugLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRunDebugLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
