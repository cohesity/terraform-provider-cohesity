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

// DeleteUsersReader is a Reader for the DeleteUsers structure.
type DeleteUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUsersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUsersNoContent creates a DeleteUsersNoContent with default headers values
func NewDeleteUsersNoContent() *DeleteUsersNoContent {
	return &DeleteUsersNoContent{}
}

/*
DeleteUsersNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteUsersNoContent struct {
}

// IsSuccess returns true when this delete users no content response has a 2xx status code
func (o *DeleteUsersNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete users no content response has a 3xx status code
func (o *DeleteUsersNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete users no content response has a 4xx status code
func (o *DeleteUsersNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete users no content response has a 5xx status code
func (o *DeleteUsersNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete users no content response a status code equal to that given
func (o *DeleteUsersNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete users no content response
func (o *DeleteUsersNoContent) Code() int {
	return 204
}

func (o *DeleteUsersNoContent) Error() string {
	return fmt.Sprintf("[POST /users/delete][%d] deleteUsersNoContent", 204)
}

func (o *DeleteUsersNoContent) String() string {
	return fmt.Sprintf("[POST /users/delete][%d] deleteUsersNoContent", 204)
}

func (o *DeleteUsersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsersDefault creates a DeleteUsersDefault with default headers values
func NewDeleteUsersDefault(code int) *DeleteUsersDefault {
	return &DeleteUsersDefault{
		_statusCode: code,
	}
}

/*
DeleteUsersDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete users default response has a 2xx status code
func (o *DeleteUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete users default response has a 3xx status code
func (o *DeleteUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete users default response has a 4xx status code
func (o *DeleteUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete users default response has a 5xx status code
func (o *DeleteUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete users default response a status code equal to that given
func (o *DeleteUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete users default response
func (o *DeleteUsersDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUsersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/delete][%d] DeleteUsers default %s", o._statusCode, payload)
}

func (o *DeleteUsersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/delete][%d] DeleteUsers default %s", o._statusCode, payload)
}

func (o *DeleteUsersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
