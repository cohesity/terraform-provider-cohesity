// Code generated by go-swagger; DO NOT EDIT.

package syslog

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

// GetSyslogServersReader is a Reader for the GetSyslogServers structure.
type GetSyslogServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyslogServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyslogServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSyslogServersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSyslogServersOK creates a GetSyslogServersOK with default headers values
func NewGetSyslogServersOK() *GetSyslogServersOK {
	return &GetSyslogServersOK{}
}

/*
GetSyslogServersOK describes a response with status code 200, with default header values.

Success
*/
type GetSyslogServersOK struct {
	Payload *models.SyslogServers
}

// IsSuccess returns true when this get syslog servers o k response has a 2xx status code
func (o *GetSyslogServersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get syslog servers o k response has a 3xx status code
func (o *GetSyslogServersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get syslog servers o k response has a 4xx status code
func (o *GetSyslogServersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get syslog servers o k response has a 5xx status code
func (o *GetSyslogServersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get syslog servers o k response a status code equal to that given
func (o *GetSyslogServersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get syslog servers o k response
func (o *GetSyslogServersOK) Code() int {
	return 200
}

func (o *GetSyslogServersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog][%d] getSyslogServersOK %s", 200, payload)
}

func (o *GetSyslogServersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog][%d] getSyslogServersOK %s", 200, payload)
}

func (o *GetSyslogServersOK) GetPayload() *models.SyslogServers {
	return o.Payload
}

func (o *GetSyslogServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyslogServers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyslogServersDefault creates a GetSyslogServersDefault with default headers values
func NewGetSyslogServersDefault(code int) *GetSyslogServersDefault {
	return &GetSyslogServersDefault{
		_statusCode: code,
	}
}

/*
GetSyslogServersDefault describes a response with status code -1, with default header values.

Error
*/
type GetSyslogServersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get syslog servers default response has a 2xx status code
func (o *GetSyslogServersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get syslog servers default response has a 3xx status code
func (o *GetSyslogServersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get syslog servers default response has a 4xx status code
func (o *GetSyslogServersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get syslog servers default response has a 5xx status code
func (o *GetSyslogServersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get syslog servers default response a status code equal to that given
func (o *GetSyslogServersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get syslog servers default response
func (o *GetSyslogServersDefault) Code() int {
	return o._statusCode
}

func (o *GetSyslogServersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog][%d] GetSyslogServers default %s", o._statusCode, payload)
}

func (o *GetSyslogServersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog][%d] GetSyslogServers default %s", o._statusCode, payload)
}

func (o *GetSyslogServersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSyslogServersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
