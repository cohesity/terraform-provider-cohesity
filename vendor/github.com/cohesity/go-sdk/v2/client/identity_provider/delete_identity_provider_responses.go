// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

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

// DeleteIdentityProviderReader is a Reader for the DeleteIdentityProvider structure.
type DeleteIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIdentityProviderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIdentityProviderNoContent creates a DeleteIdentityProviderNoContent with default headers values
func NewDeleteIdentityProviderNoContent() *DeleteIdentityProviderNoContent {
	return &DeleteIdentityProviderNoContent{}
}

/*
DeleteIdentityProviderNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteIdentityProviderNoContent struct {
}

// IsSuccess returns true when this delete identity provider no content response has a 2xx status code
func (o *DeleteIdentityProviderNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete identity provider no content response has a 3xx status code
func (o *DeleteIdentityProviderNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identity provider no content response has a 4xx status code
func (o *DeleteIdentityProviderNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identity provider no content response has a 5xx status code
func (o *DeleteIdentityProviderNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identity provider no content response a status code equal to that given
func (o *DeleteIdentityProviderNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete identity provider no content response
func (o *DeleteIdentityProviderNoContent) Code() int {
	return 204
}

func (o *DeleteIdentityProviderNoContent) Error() string {
	return fmt.Sprintf("[DELETE /idps/{id}][%d] deleteIdentityProviderNoContent", 204)
}

func (o *DeleteIdentityProviderNoContent) String() string {
	return fmt.Sprintf("[DELETE /idps/{id}][%d] deleteIdentityProviderNoContent", 204)
}

func (o *DeleteIdentityProviderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIdentityProviderDefault creates a DeleteIdentityProviderDefault with default headers values
func NewDeleteIdentityProviderDefault(code int) *DeleteIdentityProviderDefault {
	return &DeleteIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
DeleteIdentityProviderDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteIdentityProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete identity provider default response has a 2xx status code
func (o *DeleteIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete identity provider default response has a 3xx status code
func (o *DeleteIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete identity provider default response has a 4xx status code
func (o *DeleteIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete identity provider default response has a 5xx status code
func (o *DeleteIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete identity provider default response a status code equal to that given
func (o *DeleteIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete identity provider default response
func (o *DeleteIdentityProviderDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIdentityProviderDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /idps/{id}][%d] DeleteIdentityProvider default %s", o._statusCode, payload)
}

func (o *DeleteIdentityProviderDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /idps/{id}][%d] DeleteIdentityProvider default %s", o._statusCode, payload)
}

func (o *DeleteIdentityProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
