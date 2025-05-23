// Code generated by go-swagger; DO NOT EDIT.

package audit_log

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

// GetAuditLogsReader is a Reader for the GetAuditLogs structure.
type GetAuditLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditLogsOK creates a GetAuditLogsOK with default headers values
func NewGetAuditLogsOK() *GetAuditLogsOK {
	return &GetAuditLogsOK{}
}

/*
GetAuditLogsOK describes a response with status code 200, with default header values.

Success
*/
type GetAuditLogsOK struct {
	Payload *models.AuditLogs
}

// IsSuccess returns true when this get audit logs o k response has a 2xx status code
func (o *GetAuditLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audit logs o k response has a 3xx status code
func (o *GetAuditLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audit logs o k response has a 4xx status code
func (o *GetAuditLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audit logs o k response has a 5xx status code
func (o *GetAuditLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audit logs o k response a status code equal to that given
func (o *GetAuditLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get audit logs o k response
func (o *GetAuditLogsOK) Code() int {
	return 200
}

func (o *GetAuditLogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /audit-logs][%d] getAuditLogsOK %s", 200, payload)
}

func (o *GetAuditLogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /audit-logs][%d] getAuditLogsOK %s", 200, payload)
}

func (o *GetAuditLogsOK) GetPayload() *models.AuditLogs {
	return o.Payload
}

func (o *GetAuditLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditLogs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogsDefault creates a GetAuditLogsDefault with default headers values
func NewGetAuditLogsDefault(code int) *GetAuditLogsDefault {
	return &GetAuditLogsDefault{
		_statusCode: code,
	}
}

/*
GetAuditLogsDefault describes a response with status code -1, with default header values.

Error
*/
type GetAuditLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get audit logs default response has a 2xx status code
func (o *GetAuditLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get audit logs default response has a 3xx status code
func (o *GetAuditLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get audit logs default response has a 4xx status code
func (o *GetAuditLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get audit logs default response has a 5xx status code
func (o *GetAuditLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get audit logs default response a status code equal to that given
func (o *GetAuditLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get audit logs default response
func (o *GetAuditLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditLogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /audit-logs][%d] GetAuditLogs default %s", o._statusCode, payload)
}

func (o *GetAuditLogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /audit-logs][%d] GetAuditLogs default %s", o._statusCode, payload)
}

func (o *GetAuditLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuditLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
