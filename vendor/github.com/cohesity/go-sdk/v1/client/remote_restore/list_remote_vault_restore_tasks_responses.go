// Code generated by go-swagger; DO NOT EDIT.

package remote_restore

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

// ListRemoteVaultRestoreTasksReader is a Reader for the ListRemoteVaultRestoreTasks structure.
type ListRemoteVaultRestoreTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRemoteVaultRestoreTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRemoteVaultRestoreTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRemoteVaultRestoreTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRemoteVaultRestoreTasksOK creates a ListRemoteVaultRestoreTasksOK with default headers values
func NewListRemoteVaultRestoreTasksOK() *ListRemoteVaultRestoreTasksOK {
	return &ListRemoteVaultRestoreTasksOK{}
}

/*
ListRemoteVaultRestoreTasksOK describes a response with status code 200, with default header values.

Success
*/
type ListRemoteVaultRestoreTasksOK struct {
	Payload []*models.RemoteVaultRestoreTaskStatus
}

// IsSuccess returns true when this list remote vault restore tasks o k response has a 2xx status code
func (o *ListRemoteVaultRestoreTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list remote vault restore tasks o k response has a 3xx status code
func (o *ListRemoteVaultRestoreTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list remote vault restore tasks o k response has a 4xx status code
func (o *ListRemoteVaultRestoreTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list remote vault restore tasks o k response has a 5xx status code
func (o *ListRemoteVaultRestoreTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list remote vault restore tasks o k response a status code equal to that given
func (o *ListRemoteVaultRestoreTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list remote vault restore tasks o k response
func (o *ListRemoteVaultRestoreTasksOK) Code() int {
	return 200
}

func (o *ListRemoteVaultRestoreTasksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/restoreTasks][%d] listRemoteVaultRestoreTasksOK %s", 200, payload)
}

func (o *ListRemoteVaultRestoreTasksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/restoreTasks][%d] listRemoteVaultRestoreTasksOK %s", 200, payload)
}

func (o *ListRemoteVaultRestoreTasksOK) GetPayload() []*models.RemoteVaultRestoreTaskStatus {
	return o.Payload
}

func (o *ListRemoteVaultRestoreTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRemoteVaultRestoreTasksDefault creates a ListRemoteVaultRestoreTasksDefault with default headers values
func NewListRemoteVaultRestoreTasksDefault(code int) *ListRemoteVaultRestoreTasksDefault {
	return &ListRemoteVaultRestoreTasksDefault{
		_statusCode: code,
	}
}

/*
ListRemoteVaultRestoreTasksDefault describes a response with status code -1, with default header values.

Error
*/
type ListRemoteVaultRestoreTasksDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this list remote vault restore tasks default response has a 2xx status code
func (o *ListRemoteVaultRestoreTasksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list remote vault restore tasks default response has a 3xx status code
func (o *ListRemoteVaultRestoreTasksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list remote vault restore tasks default response has a 4xx status code
func (o *ListRemoteVaultRestoreTasksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list remote vault restore tasks default response has a 5xx status code
func (o *ListRemoteVaultRestoreTasksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list remote vault restore tasks default response a status code equal to that given
func (o *ListRemoteVaultRestoreTasksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list remote vault restore tasks default response
func (o *ListRemoteVaultRestoreTasksDefault) Code() int {
	return o._statusCode
}

func (o *ListRemoteVaultRestoreTasksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/restoreTasks][%d] ListRemoteVaultRestoreTasks default %s", o._statusCode, payload)
}

func (o *ListRemoteVaultRestoreTasksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/remoteVaults/restoreTasks][%d] ListRemoteVaultRestoreTasks default %s", o._statusCode, payload)
}

func (o *ListRemoteVaultRestoreTasksDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *ListRemoteVaultRestoreTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
