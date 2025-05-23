// Code generated by go-swagger; DO NOT EDIT.

package data_tiering

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

// GetDataTieringTaskByIDReader is a Reader for the GetDataTieringTaskByID structure.
type GetDataTieringTaskByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataTieringTaskByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataTieringTaskByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDataTieringTaskByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDataTieringTaskByIDOK creates a GetDataTieringTaskByIDOK with default headers values
func NewGetDataTieringTaskByIDOK() *GetDataTieringTaskByIDOK {
	return &GetDataTieringTaskByIDOK{}
}

/*
GetDataTieringTaskByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetDataTieringTaskByIDOK struct {
	Payload *models.DataTieringTask
}

// IsSuccess returns true when this get data tiering task by Id o k response has a 2xx status code
func (o *GetDataTieringTaskByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get data tiering task by Id o k response has a 3xx status code
func (o *GetDataTieringTaskByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data tiering task by Id o k response has a 4xx status code
func (o *GetDataTieringTaskByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data tiering task by Id o k response has a 5xx status code
func (o *GetDataTieringTaskByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get data tiering task by Id o k response a status code equal to that given
func (o *GetDataTieringTaskByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get data tiering task by Id o k response
func (o *GetDataTieringTaskByIDOK) Code() int {
	return 200
}

func (o *GetDataTieringTaskByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-tiering/tasks/{id}][%d] getDataTieringTaskByIdOK %s", 200, payload)
}

func (o *GetDataTieringTaskByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-tiering/tasks/{id}][%d] getDataTieringTaskByIdOK %s", 200, payload)
}

func (o *GetDataTieringTaskByIDOK) GetPayload() *models.DataTieringTask {
	return o.Payload
}

func (o *GetDataTieringTaskByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTieringTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataTieringTaskByIDDefault creates a GetDataTieringTaskByIDDefault with default headers values
func NewGetDataTieringTaskByIDDefault(code int) *GetDataTieringTaskByIDDefault {
	return &GetDataTieringTaskByIDDefault{
		_statusCode: code,
	}
}

/*
GetDataTieringTaskByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetDataTieringTaskByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get data tiering task by Id default response has a 2xx status code
func (o *GetDataTieringTaskByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get data tiering task by Id default response has a 3xx status code
func (o *GetDataTieringTaskByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get data tiering task by Id default response has a 4xx status code
func (o *GetDataTieringTaskByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get data tiering task by Id default response has a 5xx status code
func (o *GetDataTieringTaskByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get data tiering task by Id default response a status code equal to that given
func (o *GetDataTieringTaskByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get data tiering task by Id default response
func (o *GetDataTieringTaskByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDataTieringTaskByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-tiering/tasks/{id}][%d] GetDataTieringTaskById default %s", o._statusCode, payload)
}

func (o *GetDataTieringTaskByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-tiering/tasks/{id}][%d] GetDataTieringTaskById default %s", o._statusCode, payload)
}

func (o *GetDataTieringTaskByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDataTieringTaskByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
