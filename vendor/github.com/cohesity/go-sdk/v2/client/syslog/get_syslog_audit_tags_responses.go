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

// GetSyslogAuditTagsReader is a Reader for the GetSyslogAuditTags structure.
type GetSyslogAuditTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyslogAuditTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyslogAuditTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSyslogAuditTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSyslogAuditTagsOK creates a GetSyslogAuditTagsOK with default headers values
func NewGetSyslogAuditTagsOK() *GetSyslogAuditTagsOK {
	return &GetSyslogAuditTagsOK{}
}

/*
GetSyslogAuditTagsOK describes a response with status code 200, with default header values.

Success
*/
type GetSyslogAuditTagsOK struct {
	Payload *models.SyslogAuditTag
}

// IsSuccess returns true when this get syslog audit tags o k response has a 2xx status code
func (o *GetSyslogAuditTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get syslog audit tags o k response has a 3xx status code
func (o *GetSyslogAuditTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get syslog audit tags o k response has a 4xx status code
func (o *GetSyslogAuditTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get syslog audit tags o k response has a 5xx status code
func (o *GetSyslogAuditTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get syslog audit tags o k response a status code equal to that given
func (o *GetSyslogAuditTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get syslog audit tags o k response
func (o *GetSyslogAuditTagsOK) Code() int {
	return 200
}

func (o *GetSyslogAuditTagsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog/audit-tags][%d] getSyslogAuditTagsOK %s", 200, payload)
}

func (o *GetSyslogAuditTagsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog/audit-tags][%d] getSyslogAuditTagsOK %s", 200, payload)
}

func (o *GetSyslogAuditTagsOK) GetPayload() *models.SyslogAuditTag {
	return o.Payload
}

func (o *GetSyslogAuditTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyslogAuditTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyslogAuditTagsDefault creates a GetSyslogAuditTagsDefault with default headers values
func NewGetSyslogAuditTagsDefault(code int) *GetSyslogAuditTagsDefault {
	return &GetSyslogAuditTagsDefault{
		_statusCode: code,
	}
}

/*
GetSyslogAuditTagsDefault describes a response with status code -1, with default header values.

Error
*/
type GetSyslogAuditTagsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get syslog audit tags default response has a 2xx status code
func (o *GetSyslogAuditTagsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get syslog audit tags default response has a 3xx status code
func (o *GetSyslogAuditTagsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get syslog audit tags default response has a 4xx status code
func (o *GetSyslogAuditTagsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get syslog audit tags default response has a 5xx status code
func (o *GetSyslogAuditTagsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get syslog audit tags default response a status code equal to that given
func (o *GetSyslogAuditTagsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get syslog audit tags default response
func (o *GetSyslogAuditTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetSyslogAuditTagsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog/audit-tags][%d] GetSyslogAuditTags default %s", o._statusCode, payload)
}

func (o *GetSyslogAuditTagsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /syslog/audit-tags][%d] GetSyslogAuditTags default %s", o._statusCode, payload)
}

func (o *GetSyslogAuditTagsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSyslogAuditTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
