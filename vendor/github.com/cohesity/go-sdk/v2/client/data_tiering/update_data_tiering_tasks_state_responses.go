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

// UpdateDataTieringTasksStateReader is a Reader for the UpdateDataTieringTasksState structure.
type UpdateDataTieringTasksStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDataTieringTasksStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDataTieringTasksStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateDataTieringTasksStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDataTieringTasksStateOK creates a UpdateDataTieringTasksStateOK with default headers values
func NewUpdateDataTieringTasksStateOK() *UpdateDataTieringTasksStateOK {
	return &UpdateDataTieringTasksStateOK{}
}

/*
UpdateDataTieringTasksStateOK describes a response with status code 200, with default header values.

Success
*/
type UpdateDataTieringTasksStateOK struct {
	Payload *models.UpdateDataTieringState
}

// IsSuccess returns true when this update data tiering tasks state o k response has a 2xx status code
func (o *UpdateDataTieringTasksStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update data tiering tasks state o k response has a 3xx status code
func (o *UpdateDataTieringTasksStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update data tiering tasks state o k response has a 4xx status code
func (o *UpdateDataTieringTasksStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update data tiering tasks state o k response has a 5xx status code
func (o *UpdateDataTieringTasksStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update data tiering tasks state o k response a status code equal to that given
func (o *UpdateDataTieringTasksStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update data tiering tasks state o k response
func (o *UpdateDataTieringTasksStateOK) Code() int {
	return 200
}

func (o *UpdateDataTieringTasksStateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/tasks/states][%d] updateDataTieringTasksStateOK %s", 200, payload)
}

func (o *UpdateDataTieringTasksStateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/tasks/states][%d] updateDataTieringTasksStateOK %s", 200, payload)
}

func (o *UpdateDataTieringTasksStateOK) GetPayload() *models.UpdateDataTieringState {
	return o.Payload
}

func (o *UpdateDataTieringTasksStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDataTieringState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataTieringTasksStateDefault creates a UpdateDataTieringTasksStateDefault with default headers values
func NewUpdateDataTieringTasksStateDefault(code int) *UpdateDataTieringTasksStateDefault {
	return &UpdateDataTieringTasksStateDefault{
		_statusCode: code,
	}
}

/*
UpdateDataTieringTasksStateDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDataTieringTasksStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update data tiering tasks state default response has a 2xx status code
func (o *UpdateDataTieringTasksStateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update data tiering tasks state default response has a 3xx status code
func (o *UpdateDataTieringTasksStateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update data tiering tasks state default response has a 4xx status code
func (o *UpdateDataTieringTasksStateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update data tiering tasks state default response has a 5xx status code
func (o *UpdateDataTieringTasksStateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update data tiering tasks state default response a status code equal to that given
func (o *UpdateDataTieringTasksStateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update data tiering tasks state default response
func (o *UpdateDataTieringTasksStateDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDataTieringTasksStateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/tasks/states][%d] UpdateDataTieringTasksState default %s", o._statusCode, payload)
}

func (o *UpdateDataTieringTasksStateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/tasks/states][%d] UpdateDataTieringTasksState default %s", o._statusCode, payload)
}

func (o *UpdateDataTieringTasksStateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDataTieringTasksStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
