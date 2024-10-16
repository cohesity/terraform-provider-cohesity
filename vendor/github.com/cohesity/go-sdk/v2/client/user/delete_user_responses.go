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

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/*
DeleteUserNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteUserNoContent struct {
}

// IsSuccess returns true when this delete user no content response has a 2xx status code
func (o *DeleteUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user no content response has a 3xx status code
func (o *DeleteUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user no content response has a 4xx status code
func (o *DeleteUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user no content response has a 5xx status code
func (o *DeleteUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user no content response a status code equal to that given
func (o *DeleteUserNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete user no content response
func (o *DeleteUserNoContent) Code() int {
	return 204
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{sid}][%d] deleteUserNoContent", 204)
}

func (o *DeleteUserNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/{sid}][%d] deleteUserNoContent", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserDefault creates a DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	return &DeleteUserDefault{
		_statusCode: code,
	}
}

/*
DeleteUserDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete user default response has a 2xx status code
func (o *DeleteUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete user default response has a 3xx status code
func (o *DeleteUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete user default response has a 4xx status code
func (o *DeleteUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete user default response has a 5xx status code
func (o *DeleteUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete user default response a status code equal to that given
func (o *DeleteUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete user default response
func (o *DeleteUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/{sid}][%d] DeleteUser default %s", o._statusCode, payload)
}

func (o *DeleteUserDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /users/{sid}][%d] DeleteUser default %s", o._statusCode, payload)
}

func (o *DeleteUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
