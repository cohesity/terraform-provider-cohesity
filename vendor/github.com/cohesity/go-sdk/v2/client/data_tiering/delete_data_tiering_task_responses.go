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

// DeleteDataTieringTaskReader is a Reader for the DeleteDataTieringTask structure.
type DeleteDataTieringTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDataTieringTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDataTieringTaskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDataTieringTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDataTieringTaskNoContent creates a DeleteDataTieringTaskNoContent with default headers values
func NewDeleteDataTieringTaskNoContent() *DeleteDataTieringTaskNoContent {
	return &DeleteDataTieringTaskNoContent{}
}

/*
DeleteDataTieringTaskNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteDataTieringTaskNoContent struct {
}

// IsSuccess returns true when this delete data tiering task no content response has a 2xx status code
func (o *DeleteDataTieringTaskNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete data tiering task no content response has a 3xx status code
func (o *DeleteDataTieringTaskNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete data tiering task no content response has a 4xx status code
func (o *DeleteDataTieringTaskNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete data tiering task no content response has a 5xx status code
func (o *DeleteDataTieringTaskNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete data tiering task no content response a status code equal to that given
func (o *DeleteDataTieringTaskNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete data tiering task no content response
func (o *DeleteDataTieringTaskNoContent) Code() int {
	return 204
}

func (o *DeleteDataTieringTaskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /data-tiering/tasks/{id}][%d] deleteDataTieringTaskNoContent", 204)
}

func (o *DeleteDataTieringTaskNoContent) String() string {
	return fmt.Sprintf("[DELETE /data-tiering/tasks/{id}][%d] deleteDataTieringTaskNoContent", 204)
}

func (o *DeleteDataTieringTaskNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDataTieringTaskDefault creates a DeleteDataTieringTaskDefault with default headers values
func NewDeleteDataTieringTaskDefault(code int) *DeleteDataTieringTaskDefault {
	return &DeleteDataTieringTaskDefault{
		_statusCode: code,
	}
}

/*
DeleteDataTieringTaskDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteDataTieringTaskDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete data tiering task default response has a 2xx status code
func (o *DeleteDataTieringTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete data tiering task default response has a 3xx status code
func (o *DeleteDataTieringTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete data tiering task default response has a 4xx status code
func (o *DeleteDataTieringTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete data tiering task default response has a 5xx status code
func (o *DeleteDataTieringTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete data tiering task default response a status code equal to that given
func (o *DeleteDataTieringTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete data tiering task default response
func (o *DeleteDataTieringTaskDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDataTieringTaskDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /data-tiering/tasks/{id}][%d] DeleteDataTieringTask default %s", o._statusCode, payload)
}

func (o *DeleteDataTieringTaskDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /data-tiering/tasks/{id}][%d] DeleteDataTieringTask default %s", o._statusCode, payload)
}

func (o *DeleteDataTieringTaskDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDataTieringTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
