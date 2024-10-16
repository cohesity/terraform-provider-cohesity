// Code generated by go-swagger; DO NOT EDIT.

package view_boxes

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

// DeleteViewBoxReader is a Reader for the DeleteViewBox structure.
type DeleteViewBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteViewBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteViewBoxNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteViewBoxDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteViewBoxNoContent creates a DeleteViewBoxNoContent with default headers values
func NewDeleteViewBoxNoContent() *DeleteViewBoxNoContent {
	return &DeleteViewBoxNoContent{}
}

/*
DeleteViewBoxNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteViewBoxNoContent struct {
}

// IsSuccess returns true when this delete view box no content response has a 2xx status code
func (o *DeleteViewBoxNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete view box no content response has a 3xx status code
func (o *DeleteViewBoxNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete view box no content response has a 4xx status code
func (o *DeleteViewBoxNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete view box no content response has a 5xx status code
func (o *DeleteViewBoxNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete view box no content response a status code equal to that given
func (o *DeleteViewBoxNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete view box no content response
func (o *DeleteViewBoxNoContent) Code() int {
	return 204
}

func (o *DeleteViewBoxNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/viewBoxes/{id}][%d] deleteViewBoxNoContent", 204)
}

func (o *DeleteViewBoxNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/viewBoxes/{id}][%d] deleteViewBoxNoContent", 204)
}

func (o *DeleteViewBoxNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteViewBoxDefault creates a DeleteViewBoxDefault with default headers values
func NewDeleteViewBoxDefault(code int) *DeleteViewBoxDefault {
	return &DeleteViewBoxDefault{
		_statusCode: code,
	}
}

/*
DeleteViewBoxDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteViewBoxDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this delete view box default response has a 2xx status code
func (o *DeleteViewBoxDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete view box default response has a 3xx status code
func (o *DeleteViewBoxDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete view box default response has a 4xx status code
func (o *DeleteViewBoxDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete view box default response has a 5xx status code
func (o *DeleteViewBoxDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete view box default response a status code equal to that given
func (o *DeleteViewBoxDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete view box default response
func (o *DeleteViewBoxDefault) Code() int {
	return o._statusCode
}

func (o *DeleteViewBoxDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/viewBoxes/{id}][%d] DeleteViewBox default %s", o._statusCode, payload)
}

func (o *DeleteViewBoxDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/viewBoxes/{id}][%d] DeleteViewBox default %s", o._statusCode, payload)
}

func (o *DeleteViewBoxDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *DeleteViewBoxDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
