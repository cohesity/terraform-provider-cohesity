// Code generated by go-swagger; DO NOT EDIT.

package protection_jobs

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

// UpdateProtectionJobReader is a Reader for the UpdateProtectionJob structure.
type UpdateProtectionJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProtectionJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProtectionJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateProtectionJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProtectionJobOK creates a UpdateProtectionJobOK with default headers values
func NewUpdateProtectionJobOK() *UpdateProtectionJobOK {
	return &UpdateProtectionJobOK{}
}

/*
UpdateProtectionJobOK describes a response with status code 200, with default header values.

Success
*/
type UpdateProtectionJobOK struct {
	Payload *models.ProtectionJob
}

// IsSuccess returns true when this update protection job o k response has a 2xx status code
func (o *UpdateProtectionJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update protection job o k response has a 3xx status code
func (o *UpdateProtectionJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update protection job o k response has a 4xx status code
func (o *UpdateProtectionJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update protection job o k response has a 5xx status code
func (o *UpdateProtectionJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update protection job o k response a status code equal to that given
func (o *UpdateProtectionJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update protection job o k response
func (o *UpdateProtectionJobOK) Code() int {
	return 200
}

func (o *UpdateProtectionJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/protectionJobs/{id}][%d] updateProtectionJobOK %s", 200, payload)
}

func (o *UpdateProtectionJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/protectionJobs/{id}][%d] updateProtectionJobOK %s", 200, payload)
}

func (o *UpdateProtectionJobOK) GetPayload() *models.ProtectionJob {
	return o.Payload
}

func (o *UpdateProtectionJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProtectionJobDefault creates a UpdateProtectionJobDefault with default headers values
func NewUpdateProtectionJobDefault(code int) *UpdateProtectionJobDefault {
	return &UpdateProtectionJobDefault{
		_statusCode: code,
	}
}

/*
UpdateProtectionJobDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateProtectionJobDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update protection job default response has a 2xx status code
func (o *UpdateProtectionJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update protection job default response has a 3xx status code
func (o *UpdateProtectionJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update protection job default response has a 4xx status code
func (o *UpdateProtectionJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update protection job default response has a 5xx status code
func (o *UpdateProtectionJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update protection job default response a status code equal to that given
func (o *UpdateProtectionJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update protection job default response
func (o *UpdateProtectionJobDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProtectionJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/protectionJobs/{id}][%d] UpdateProtectionJob default %s", o._statusCode, payload)
}

func (o *UpdateProtectionJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/protectionJobs/{id}][%d] UpdateProtectionJob default %s", o._statusCode, payload)
}

func (o *UpdateProtectionJobDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateProtectionJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
