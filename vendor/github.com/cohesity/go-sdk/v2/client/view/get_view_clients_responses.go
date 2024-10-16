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

// GetViewClientsReader is a Reader for the GetViewClients structure.
type GetViewClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetViewClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetViewClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetViewClientsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetViewClientsOK creates a GetViewClientsOK with default headers values
func NewGetViewClientsOK() *GetViewClientsOK {
	return &GetViewClientsOK{}
}

/*
GetViewClientsOK describes a response with status code 200, with default header values.

Success
*/
type GetViewClientsOK struct {
	Payload *models.ViewClients
}

// IsSuccess returns true when this get view clients o k response has a 2xx status code
func (o *GetViewClientsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get view clients o k response has a 3xx status code
func (o *GetViewClientsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get view clients o k response has a 4xx status code
func (o *GetViewClientsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get view clients o k response has a 5xx status code
func (o *GetViewClientsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get view clients o k response a status code equal to that given
func (o *GetViewClientsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get view clients o k response
func (o *GetViewClientsOK) Code() int {
	return 200
}

func (o *GetViewClientsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/view-clients][%d] getViewClientsOK %s", 200, payload)
}

func (o *GetViewClientsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/view-clients][%d] getViewClientsOK %s", 200, payload)
}

func (o *GetViewClientsOK) GetPayload() *models.ViewClients {
	return o.Payload
}

func (o *GetViewClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ViewClients)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetViewClientsDefault creates a GetViewClientsDefault with default headers values
func NewGetViewClientsDefault(code int) *GetViewClientsDefault {
	return &GetViewClientsDefault{
		_statusCode: code,
	}
}

/*
GetViewClientsDefault describes a response with status code -1, with default header values.

Error
*/
type GetViewClientsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get view clients default response has a 2xx status code
func (o *GetViewClientsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get view clients default response has a 3xx status code
func (o *GetViewClientsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get view clients default response has a 4xx status code
func (o *GetViewClientsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get view clients default response has a 5xx status code
func (o *GetViewClientsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get view clients default response a status code equal to that given
func (o *GetViewClientsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get view clients default response
func (o *GetViewClientsDefault) Code() int {
	return o._statusCode
}

func (o *GetViewClientsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/view-clients][%d] GetViewClients default %s", o._statusCode, payload)
}

func (o *GetViewClientsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/view-clients][%d] GetViewClients default %s", o._statusCode, payload)
}

func (o *GetViewClientsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetViewClientsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
