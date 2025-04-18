// Code generated by go-swagger; DO NOT EDIT.

package view

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

// GetViewDirectoryQuotasReader is a Reader for the GetViewDirectoryQuotas structure.
type GetViewDirectoryQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetViewDirectoryQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetViewDirectoryQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetViewDirectoryQuotasDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetViewDirectoryQuotasOK creates a GetViewDirectoryQuotasOK with default headers values
func NewGetViewDirectoryQuotasOK() *GetViewDirectoryQuotasOK {
	return &GetViewDirectoryQuotasOK{}
}

/*
GetViewDirectoryQuotasOK describes a response with status code 200, with default header values.

Success
*/
type GetViewDirectoryQuotasOK struct {
	Payload *models.ViewDirectoryQuotas
}

// IsSuccess returns true when this get view directory quotas o k response has a 2xx status code
func (o *GetViewDirectoryQuotasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get view directory quotas o k response has a 3xx status code
func (o *GetViewDirectoryQuotasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get view directory quotas o k response has a 4xx status code
func (o *GetViewDirectoryQuotasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get view directory quotas o k response has a 5xx status code
func (o *GetViewDirectoryQuotasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get view directory quotas o k response a status code equal to that given
func (o *GetViewDirectoryQuotasOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get view directory quotas o k response
func (o *GetViewDirectoryQuotasOK) Code() int {
	return 200
}

func (o *GetViewDirectoryQuotasOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views/{id}/directory-quotas][%d] getViewDirectoryQuotasOK %s", 200, payload)
}

func (o *GetViewDirectoryQuotasOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views/{id}/directory-quotas][%d] getViewDirectoryQuotasOK %s", 200, payload)
}

func (o *GetViewDirectoryQuotasOK) GetPayload() *models.ViewDirectoryQuotas {
	return o.Payload
}

func (o *GetViewDirectoryQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ViewDirectoryQuotas)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetViewDirectoryQuotasDefault creates a GetViewDirectoryQuotasDefault with default headers values
func NewGetViewDirectoryQuotasDefault(code int) *GetViewDirectoryQuotasDefault {
	return &GetViewDirectoryQuotasDefault{
		_statusCode: code,
	}
}

/*
GetViewDirectoryQuotasDefault describes a response with status code -1, with default header values.

Error
*/
type GetViewDirectoryQuotasDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get view directory quotas default response has a 2xx status code
func (o *GetViewDirectoryQuotasDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get view directory quotas default response has a 3xx status code
func (o *GetViewDirectoryQuotasDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get view directory quotas default response has a 4xx status code
func (o *GetViewDirectoryQuotasDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get view directory quotas default response has a 5xx status code
func (o *GetViewDirectoryQuotasDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get view directory quotas default response a status code equal to that given
func (o *GetViewDirectoryQuotasDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get view directory quotas default response
func (o *GetViewDirectoryQuotasDefault) Code() int {
	return o._statusCode
}

func (o *GetViewDirectoryQuotasDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views/{id}/directory-quotas][%d] GetViewDirectoryQuotas default %s", o._statusCode, payload)
}

func (o *GetViewDirectoryQuotasDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views/{id}/directory-quotas][%d] GetViewDirectoryQuotas default %s", o._statusCode, payload)
}

func (o *GetViewDirectoryQuotasDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetViewDirectoryQuotasDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
