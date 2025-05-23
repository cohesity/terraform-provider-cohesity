// Code generated by go-swagger; DO NOT EDIT.

package source

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

// UpdateM365BackupControllerReader is a Reader for the UpdateM365BackupController structure.
type UpdateM365BackupControllerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateM365BackupControllerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateM365BackupControllerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateM365BackupControllerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateM365BackupControllerOK creates a UpdateM365BackupControllerOK with default headers values
func NewUpdateM365BackupControllerOK() *UpdateM365BackupControllerOK {
	return &UpdateM365BackupControllerOK{}
}

/*
UpdateM365BackupControllerOK describes a response with status code 200, with default header values.

Success
*/
type UpdateM365BackupControllerOK struct {
	Payload *models.GetM365BackupControllerResponseParams
}

// IsSuccess returns true when this update m365 backup controller o k response has a 2xx status code
func (o *UpdateM365BackupControllerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update m365 backup controller o k response has a 3xx status code
func (o *UpdateM365BackupControllerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m365 backup controller o k response has a 4xx status code
func (o *UpdateM365BackupControllerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update m365 backup controller o k response has a 5xx status code
func (o *UpdateM365BackupControllerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update m365 backup controller o k response a status code equal to that given
func (o *UpdateM365BackupControllerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update m365 backup controller o k response
func (o *UpdateM365BackupControllerOK) Code() int {
	return 200
}

func (o *UpdateM365BackupControllerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /data-protect/sources/microsoft365/backup-controllers/{id}][%d] updateM365BackupControllerOK %s", 200, payload)
}

func (o *UpdateM365BackupControllerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /data-protect/sources/microsoft365/backup-controllers/{id}][%d] updateM365BackupControllerOK %s", 200, payload)
}

func (o *UpdateM365BackupControllerOK) GetPayload() *models.GetM365BackupControllerResponseParams {
	return o.Payload
}

func (o *UpdateM365BackupControllerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetM365BackupControllerResponseParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateM365BackupControllerDefault creates a UpdateM365BackupControllerDefault with default headers values
func NewUpdateM365BackupControllerDefault(code int) *UpdateM365BackupControllerDefault {
	return &UpdateM365BackupControllerDefault{
		_statusCode: code,
	}
}

/*
UpdateM365BackupControllerDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateM365BackupControllerDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update m365 backup controller default response has a 2xx status code
func (o *UpdateM365BackupControllerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update m365 backup controller default response has a 3xx status code
func (o *UpdateM365BackupControllerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update m365 backup controller default response has a 4xx status code
func (o *UpdateM365BackupControllerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update m365 backup controller default response has a 5xx status code
func (o *UpdateM365BackupControllerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update m365 backup controller default response a status code equal to that given
func (o *UpdateM365BackupControllerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update m365 backup controller default response
func (o *UpdateM365BackupControllerDefault) Code() int {
	return o._statusCode
}

func (o *UpdateM365BackupControllerDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /data-protect/sources/microsoft365/backup-controllers/{id}][%d] UpdateM365BackupController default %s", o._statusCode, payload)
}

func (o *UpdateM365BackupControllerDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /data-protect/sources/microsoft365/backup-controllers/{id}][%d] UpdateM365BackupController default %s", o._statusCode, payload)
}

func (o *UpdateM365BackupControllerDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateM365BackupControllerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
