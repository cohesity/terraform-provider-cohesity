// Code generated by go-swagger; DO NOT EDIT.

package recovery

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

// GetRecoveryDebugLogsReader is a Reader for the GetRecoveryDebugLogs structure.
type GetRecoveryDebugLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecoveryDebugLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecoveryDebugLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRecoveryDebugLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRecoveryDebugLogsOK creates a GetRecoveryDebugLogsOK with default headers values
func NewGetRecoveryDebugLogsOK() *GetRecoveryDebugLogsOK {
	return &GetRecoveryDebugLogsOK{}
}

/*
GetRecoveryDebugLogsOK describes a response with status code 200, with default header values.

No Content
*/
type GetRecoveryDebugLogsOK struct {
}

// IsSuccess returns true when this get recovery debug logs o k response has a 2xx status code
func (o *GetRecoveryDebugLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get recovery debug logs o k response has a 3xx status code
func (o *GetRecoveryDebugLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recovery debug logs o k response has a 4xx status code
func (o *GetRecoveryDebugLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recovery debug logs o k response has a 5xx status code
func (o *GetRecoveryDebugLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get recovery debug logs o k response a status code equal to that given
func (o *GetRecoveryDebugLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get recovery debug logs o k response
func (o *GetRecoveryDebugLogsOK) Code() int {
	return 200
}

func (o *GetRecoveryDebugLogsOK) Error() string {
	return fmt.Sprintf("[GET /data-protect/recoveries/{id}/debug-logs][%d] getRecoveryDebugLogsOK", 200)
}

func (o *GetRecoveryDebugLogsOK) String() string {
	return fmt.Sprintf("[GET /data-protect/recoveries/{id}/debug-logs][%d] getRecoveryDebugLogsOK", 200)
}

func (o *GetRecoveryDebugLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRecoveryDebugLogsDefault creates a GetRecoveryDebugLogsDefault with default headers values
func NewGetRecoveryDebugLogsDefault(code int) *GetRecoveryDebugLogsDefault {
	return &GetRecoveryDebugLogsDefault{
		_statusCode: code,
	}
}

/*
GetRecoveryDebugLogsDefault describes a response with status code -1, with default header values.

Error
*/
type GetRecoveryDebugLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get recovery debug logs default response has a 2xx status code
func (o *GetRecoveryDebugLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get recovery debug logs default response has a 3xx status code
func (o *GetRecoveryDebugLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get recovery debug logs default response has a 4xx status code
func (o *GetRecoveryDebugLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get recovery debug logs default response has a 5xx status code
func (o *GetRecoveryDebugLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get recovery debug logs default response a status code equal to that given
func (o *GetRecoveryDebugLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get recovery debug logs default response
func (o *GetRecoveryDebugLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetRecoveryDebugLogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/recoveries/{id}/debug-logs][%d] GetRecoveryDebugLogs default %s", o._statusCode, payload)
}

func (o *GetRecoveryDebugLogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/recoveries/{id}/debug-logs][%d] GetRecoveryDebugLogs default %s", o._statusCode, payload)
}

func (o *GetRecoveryDebugLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRecoveryDebugLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
