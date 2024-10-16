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

// GetProxyServersReader is a Reader for the GetProxyServers structure.
type GetProxyServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProxyServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProxyServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProxyServersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProxyServersOK creates a GetProxyServersOK with default headers values
func NewGetProxyServersOK() *GetProxyServersOK {
	return &GetProxyServersOK{}
}

/*
GetProxyServersOK describes a response with status code 200, with default header values.

Success
*/
type GetProxyServersOK struct {
	Payload []*models.ProxyServer
}

// IsSuccess returns true when this get proxy servers o k response has a 2xx status code
func (o *GetProxyServersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get proxy servers o k response has a 3xx status code
func (o *GetProxyServersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get proxy servers o k response has a 4xx status code
func (o *GetProxyServersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get proxy servers o k response has a 5xx status code
func (o *GetProxyServersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get proxy servers o k response a status code equal to that given
func (o *GetProxyServersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get proxy servers o k response
func (o *GetProxyServersOK) Code() int {
	return 200
}

func (o *GetProxyServersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /proxyServers][%d] getProxyServersOK %s", 200, payload)
}

func (o *GetProxyServersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /proxyServers][%d] getProxyServersOK %s", 200, payload)
}

func (o *GetProxyServersOK) GetPayload() []*models.ProxyServer {
	return o.Payload
}

func (o *GetProxyServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProxyServersDefault creates a GetProxyServersDefault with default headers values
func NewGetProxyServersDefault(code int) *GetProxyServersDefault {
	return &GetProxyServersDefault{
		_statusCode: code,
	}
}

/*
GetProxyServersDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetProxyServersDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get proxy servers default response has a 2xx status code
func (o *GetProxyServersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get proxy servers default response has a 3xx status code
func (o *GetProxyServersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get proxy servers default response has a 4xx status code
func (o *GetProxyServersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get proxy servers default response has a 5xx status code
func (o *GetProxyServersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get proxy servers default response a status code equal to that given
func (o *GetProxyServersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get proxy servers default response
func (o *GetProxyServersDefault) Code() int {
	return o._statusCode
}

func (o *GetProxyServersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /proxyServers][%d] GetProxyServers default %s", o._statusCode, payload)
}

func (o *GetProxyServersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /proxyServers][%d] GetProxyServers default %s", o._statusCode, payload)
}

func (o *GetProxyServersDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetProxyServersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
