// Code generated by go-swagger; DO NOT EDIT.

package remote_storage

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

// DeleteRemoteStorageRegistrationReader is a Reader for the DeleteRemoteStorageRegistration structure.
type DeleteRemoteStorageRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRemoteStorageRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRemoteStorageRegistrationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteRemoteStorageRegistrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRemoteStorageRegistrationNoContent creates a DeleteRemoteStorageRegistrationNoContent with default headers values
func NewDeleteRemoteStorageRegistrationNoContent() *DeleteRemoteStorageRegistrationNoContent {
	return &DeleteRemoteStorageRegistrationNoContent{}
}

/*
DeleteRemoteStorageRegistrationNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteRemoteStorageRegistrationNoContent struct {
}

// IsSuccess returns true when this delete remote storage registration no content response has a 2xx status code
func (o *DeleteRemoteStorageRegistrationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete remote storage registration no content response has a 3xx status code
func (o *DeleteRemoteStorageRegistrationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete remote storage registration no content response has a 4xx status code
func (o *DeleteRemoteStorageRegistrationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete remote storage registration no content response has a 5xx status code
func (o *DeleteRemoteStorageRegistrationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete remote storage registration no content response a status code equal to that given
func (o *DeleteRemoteStorageRegistrationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete remote storage registration no content response
func (o *DeleteRemoteStorageRegistrationNoContent) Code() int {
	return 204
}

func (o *DeleteRemoteStorageRegistrationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /remote-storage/{id}][%d] deleteRemoteStorageRegistrationNoContent", 204)
}

func (o *DeleteRemoteStorageRegistrationNoContent) String() string {
	return fmt.Sprintf("[DELETE /remote-storage/{id}][%d] deleteRemoteStorageRegistrationNoContent", 204)
}

func (o *DeleteRemoteStorageRegistrationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRemoteStorageRegistrationDefault creates a DeleteRemoteStorageRegistrationDefault with default headers values
func NewDeleteRemoteStorageRegistrationDefault(code int) *DeleteRemoteStorageRegistrationDefault {
	return &DeleteRemoteStorageRegistrationDefault{
		_statusCode: code,
	}
}

/*
DeleteRemoteStorageRegistrationDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteRemoteStorageRegistrationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete remote storage registration default response has a 2xx status code
func (o *DeleteRemoteStorageRegistrationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete remote storage registration default response has a 3xx status code
func (o *DeleteRemoteStorageRegistrationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete remote storage registration default response has a 4xx status code
func (o *DeleteRemoteStorageRegistrationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete remote storage registration default response has a 5xx status code
func (o *DeleteRemoteStorageRegistrationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete remote storage registration default response a status code equal to that given
func (o *DeleteRemoteStorageRegistrationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete remote storage registration default response
func (o *DeleteRemoteStorageRegistrationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRemoteStorageRegistrationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /remote-storage/{id}][%d] DeleteRemoteStorageRegistration default %s", o._statusCode, payload)
}

func (o *DeleteRemoteStorageRegistrationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /remote-storage/{id}][%d] DeleteRemoteStorageRegistration default %s", o._statusCode, payload)
}

func (o *DeleteRemoteStorageRegistrationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRemoteStorageRegistrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
