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

// CreateProtectionJobReader is a Reader for the CreateProtectionJob structure.
type CreateProtectionJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProtectionJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProtectionJobCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateProtectionJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProtectionJobCreated creates a CreateProtectionJobCreated with default headers values
func NewCreateProtectionJobCreated() *CreateProtectionJobCreated {
	return &CreateProtectionJobCreated{}
}

/*
CreateProtectionJobCreated describes a response with status code 201, with default header values.

Success
*/
type CreateProtectionJobCreated struct {
	Payload *models.ProtectionJob
}

// IsSuccess returns true when this create protection job created response has a 2xx status code
func (o *CreateProtectionJobCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create protection job created response has a 3xx status code
func (o *CreateProtectionJobCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create protection job created response has a 4xx status code
func (o *CreateProtectionJobCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create protection job created response has a 5xx status code
func (o *CreateProtectionJobCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create protection job created response a status code equal to that given
func (o *CreateProtectionJobCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create protection job created response
func (o *CreateProtectionJobCreated) Code() int {
	return 201
}

func (o *CreateProtectionJobCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionJobs][%d] createProtectionJobCreated %s", 201, payload)
}

func (o *CreateProtectionJobCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionJobs][%d] createProtectionJobCreated %s", 201, payload)
}

func (o *CreateProtectionJobCreated) GetPayload() *models.ProtectionJob {
	return o.Payload
}

func (o *CreateProtectionJobCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProtectionJobDefault creates a CreateProtectionJobDefault with default headers values
func NewCreateProtectionJobDefault(code int) *CreateProtectionJobDefault {
	return &CreateProtectionJobDefault{
		_statusCode: code,
	}
}

/*
CreateProtectionJobDefault describes a response with status code -1, with default header values.

Error
*/
type CreateProtectionJobDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create protection job default response has a 2xx status code
func (o *CreateProtectionJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create protection job default response has a 3xx status code
func (o *CreateProtectionJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create protection job default response has a 4xx status code
func (o *CreateProtectionJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create protection job default response has a 5xx status code
func (o *CreateProtectionJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create protection job default response a status code equal to that given
func (o *CreateProtectionJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create protection job default response
func (o *CreateProtectionJobDefault) Code() int {
	return o._statusCode
}

func (o *CreateProtectionJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionJobs][%d] CreateProtectionJob default %s", o._statusCode, payload)
}

func (o *CreateProtectionJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionJobs][%d] CreateProtectionJob default %s", o._statusCode, payload)
}

func (o *CreateProtectionJobDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateProtectionJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
