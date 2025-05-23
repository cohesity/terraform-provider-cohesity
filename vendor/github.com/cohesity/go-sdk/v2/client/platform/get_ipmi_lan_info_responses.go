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

// GetIpmiLanInfoReader is a Reader for the GetIpmiLanInfo structure.
type GetIpmiLanInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIpmiLanInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIpmiLanInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIpmiLanInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIpmiLanInfoOK creates a GetIpmiLanInfoOK with default headers values
func NewGetIpmiLanInfoOK() *GetIpmiLanInfoOK {
	return &GetIpmiLanInfoOK{}
}

/*
GetIpmiLanInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetIpmiLanInfoOK struct {
	Payload *models.IpmiLanInfo
}

// IsSuccess returns true when this get ipmi lan info o k response has a 2xx status code
func (o *GetIpmiLanInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ipmi lan info o k response has a 3xx status code
func (o *GetIpmiLanInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ipmi lan info o k response has a 4xx status code
func (o *GetIpmiLanInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ipmi lan info o k response has a 5xx status code
func (o *GetIpmiLanInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get ipmi lan info o k response a status code equal to that given
func (o *GetIpmiLanInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get ipmi lan info o k response
func (o *GetIpmiLanInfoOK) Code() int {
	return 200
}

func (o *GetIpmiLanInfoOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipmi/get-lan-info][%d] getIpmiLanInfoOK %s", 200, payload)
}

func (o *GetIpmiLanInfoOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipmi/get-lan-info][%d] getIpmiLanInfoOK %s", 200, payload)
}

func (o *GetIpmiLanInfoOK) GetPayload() *models.IpmiLanInfo {
	return o.Payload
}

func (o *GetIpmiLanInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpmiLanInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIpmiLanInfoDefault creates a GetIpmiLanInfoDefault with default headers values
func NewGetIpmiLanInfoDefault(code int) *GetIpmiLanInfoDefault {
	return &GetIpmiLanInfoDefault{
		_statusCode: code,
	}
}

/*
GetIpmiLanInfoDefault describes a response with status code -1, with default header values.

Error
*/
type GetIpmiLanInfoDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get ipmi lan info default response has a 2xx status code
func (o *GetIpmiLanInfoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get ipmi lan info default response has a 3xx status code
func (o *GetIpmiLanInfoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get ipmi lan info default response has a 4xx status code
func (o *GetIpmiLanInfoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get ipmi lan info default response has a 5xx status code
func (o *GetIpmiLanInfoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get ipmi lan info default response a status code equal to that given
func (o *GetIpmiLanInfoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get ipmi lan info default response
func (o *GetIpmiLanInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetIpmiLanInfoDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipmi/get-lan-info][%d] GetIpmiLanInfo default %s", o._statusCode, payload)
}

func (o *GetIpmiLanInfoDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipmi/get-lan-info][%d] GetIpmiLanInfo default %s", o._statusCode, payload)
}

func (o *GetIpmiLanInfoDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIpmiLanInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
