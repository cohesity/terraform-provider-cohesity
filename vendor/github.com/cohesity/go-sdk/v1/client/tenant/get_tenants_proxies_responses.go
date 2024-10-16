// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// GetTenantsProxiesReader is a Reader for the GetTenantsProxies structure.
type GetTenantsProxiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsProxiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsProxiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsProxiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsProxiesOK creates a GetTenantsProxiesOK with default headers values
func NewGetTenantsProxiesOK() *GetTenantsProxiesOK {
	return &GetTenantsProxiesOK{}
}

/*
GetTenantsProxiesOK describes a response with status code 200, with default header values.

Get Tenants Proxies Response.
*/
type GetTenantsProxiesOK struct {
	Payload []*models.TenantProxy
}

// IsSuccess returns true when this get tenants proxies o k response has a 2xx status code
func (o *GetTenantsProxiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tenants proxies o k response has a 3xx status code
func (o *GetTenantsProxiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenants proxies o k response has a 4xx status code
func (o *GetTenantsProxiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tenants proxies o k response has a 5xx status code
func (o *GetTenantsProxiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenants proxies o k response a status code equal to that given
func (o *GetTenantsProxiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tenants proxies o k response
func (o *GetTenantsProxiesOK) Code() int {
	return 200
}

func (o *GetTenantsProxiesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/tenants/proxies][%d] getTenantsProxiesOK %s", 200, payload)
}

func (o *GetTenantsProxiesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/tenants/proxies][%d] getTenantsProxiesOK %s", 200, payload)
}

func (o *GetTenantsProxiesOK) GetPayload() []*models.TenantProxy {
	return o.Payload
}

func (o *GetTenantsProxiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsProxiesDefault creates a GetTenantsProxiesDefault with default headers values
func NewGetTenantsProxiesDefault(code int) *GetTenantsProxiesDefault {
	return &GetTenantsProxiesDefault{
		_statusCode: code,
	}
}

/*
GetTenantsProxiesDefault describes a response with status code -1, with default header values.

Error
*/
type GetTenantsProxiesDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get tenants proxies default response has a 2xx status code
func (o *GetTenantsProxiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get tenants proxies default response has a 3xx status code
func (o *GetTenantsProxiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get tenants proxies default response has a 4xx status code
func (o *GetTenantsProxiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get tenants proxies default response has a 5xx status code
func (o *GetTenantsProxiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get tenants proxies default response a status code equal to that given
func (o *GetTenantsProxiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get tenants proxies default response
func (o *GetTenantsProxiesDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsProxiesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/tenants/proxies][%d] GetTenantsProxies default %s", o._statusCode, payload)
}

func (o *GetTenantsProxiesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/tenants/proxies][%d] GetTenantsProxies default %s", o._statusCode, payload)
}

func (o *GetTenantsProxiesDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetTenantsProxiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
