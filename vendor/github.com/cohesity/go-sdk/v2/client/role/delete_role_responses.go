// Code generated by go-swagger; DO NOT EDIT.

package role

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

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRoleNoContent creates a DeleteRoleNoContent with default headers values
func NewDeleteRoleNoContent() *DeleteRoleNoContent {
	return &DeleteRoleNoContent{}
}

/*
DeleteRoleNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteRoleNoContent struct {
}

// IsSuccess returns true when this delete role no content response has a 2xx status code
func (o *DeleteRoleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete role no content response has a 3xx status code
func (o *DeleteRoleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role no content response has a 4xx status code
func (o *DeleteRoleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete role no content response has a 5xx status code
func (o *DeleteRoleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role no content response a status code equal to that given
func (o *DeleteRoleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete role no content response
func (o *DeleteRoleNoContent) Code() int {
	return 204
}

func (o *DeleteRoleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /roles/{name}][%d] deleteRoleNoContent", 204)
}

func (o *DeleteRoleNoContent) String() string {
	return fmt.Sprintf("[DELETE /roles/{name}][%d] deleteRoleNoContent", 204)
}

func (o *DeleteRoleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoleDefault creates a DeleteRoleDefault with default headers values
func NewDeleteRoleDefault(code int) *DeleteRoleDefault {
	return &DeleteRoleDefault{
		_statusCode: code,
	}
}

/*
DeleteRoleDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete role default response has a 2xx status code
func (o *DeleteRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete role default response has a 3xx status code
func (o *DeleteRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete role default response has a 4xx status code
func (o *DeleteRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete role default response has a 5xx status code
func (o *DeleteRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete role default response a status code equal to that given
func (o *DeleteRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete role default response
func (o *DeleteRoleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /roles/{name}][%d] DeleteRole default %s", o._statusCode, payload)
}

func (o *DeleteRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /roles/{name}][%d] DeleteRole default %s", o._statusCode, payload)
}

func (o *DeleteRoleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
