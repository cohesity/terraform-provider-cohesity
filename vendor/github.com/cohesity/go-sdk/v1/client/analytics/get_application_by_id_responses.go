// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// GetApplicationByIDReader is a Reader for the GetApplicationByID structure.
type GetApplicationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetApplicationByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplicationByIDOK creates a GetApplicationByIDOK with default headers values
func NewGetApplicationByIDOK() *GetApplicationByIDOK {
	return &GetApplicationByIDOK{}
}

/*
GetApplicationByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetApplicationByIDOK struct {
	Payload *models.MapReduceInfo
}

// IsSuccess returns true when this get application by Id o k response has a 2xx status code
func (o *GetApplicationByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application by Id o k response has a 3xx status code
func (o *GetApplicationByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application by Id o k response has a 4xx status code
func (o *GetApplicationByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application by Id o k response has a 5xx status code
func (o *GetApplicationByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application by Id o k response a status code equal to that given
func (o *GetApplicationByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application by Id o k response
func (o *GetApplicationByIDOK) Code() int {
	return 200
}

func (o *GetApplicationByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/apps/{id}][%d] getApplicationByIdOK %s", 200, payload)
}

func (o *GetApplicationByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/apps/{id}][%d] getApplicationByIdOK %s", 200, payload)
}

func (o *GetApplicationByIDOK) GetPayload() *models.MapReduceInfo {
	return o.Payload
}

func (o *GetApplicationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MapReduceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationByIDDefault creates a GetApplicationByIDDefault with default headers values
func NewGetApplicationByIDDefault(code int) *GetApplicationByIDDefault {
	return &GetApplicationByIDDefault{
		_statusCode: code,
	}
}

/*
GetApplicationByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetApplicationByIDDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get application by Id default response has a 2xx status code
func (o *GetApplicationByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get application by Id default response has a 3xx status code
func (o *GetApplicationByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get application by Id default response has a 4xx status code
func (o *GetApplicationByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get application by Id default response has a 5xx status code
func (o *GetApplicationByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get application by Id default response a status code equal to that given
func (o *GetApplicationByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get application by Id default response
func (o *GetApplicationByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetApplicationByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/apps/{id}][%d] GetApplicationById default %s", o._statusCode, payload)
}

func (o *GetApplicationByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/apps/{id}][%d] GetApplicationById default %s", o._statusCode, payload)
}

func (o *GetApplicationByIDDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetApplicationByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
