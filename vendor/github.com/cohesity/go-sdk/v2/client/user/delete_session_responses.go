// Code generated by go-swagger; DO NOT EDIT.

package user

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

// DeleteSessionReader is a Reader for the DeleteSession structure.
type DeleteSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSessionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteSessionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSessionNoContent creates a DeleteSessionNoContent with default headers values
func NewDeleteSessionNoContent() *DeleteSessionNoContent {
	return &DeleteSessionNoContent{}
}

/*
DeleteSessionNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteSessionNoContent struct {
}

// IsSuccess returns true when this delete session no content response has a 2xx status code
func (o *DeleteSessionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete session no content response has a 3xx status code
func (o *DeleteSessionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete session no content response has a 4xx status code
func (o *DeleteSessionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete session no content response has a 5xx status code
func (o *DeleteSessionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete session no content response a status code equal to that given
func (o *DeleteSessionNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete session no content response
func (o *DeleteSessionNoContent) Code() int {
	return 204
}

func (o *DeleteSessionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/sessions][%d] deleteSessionNoContent", 204)
}

func (o *DeleteSessionNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/sessions][%d] deleteSessionNoContent", 204)
}

func (o *DeleteSessionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSessionDefault creates a DeleteSessionDefault with default headers values
func NewDeleteSessionDefault(code int) *DeleteSessionDefault {
	return &DeleteSessionDefault{
		_statusCode: code,
	}
}

/*
DeleteSessionDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteSessionDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete session default response has a 2xx status code
func (o *DeleteSessionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete session default response has a 3xx status code
func (o *DeleteSessionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete session default response has a 4xx status code
func (o *DeleteSessionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete session default response has a 5xx status code
func (o *DeleteSessionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete session default response a status code equal to that given
func (o *DeleteSessionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete session default response
func (o *DeleteSessionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSessionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/sessions][%d] DeleteSession default %s", o._statusCode, payload)
}

func (o *DeleteSessionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/sessions][%d] DeleteSession default %s", o._statusCode, payload)
}

func (o *DeleteSessionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSessionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
