// Code generated by go-swagger; DO NOT EDIT.

package backup_jobs

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

// GetBackupJobByIDReader is a Reader for the GetBackupJobByID structure.
type GetBackupJobByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupJobByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupJobByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackupJobByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackupJobByIDOK creates a GetBackupJobByIDOK with default headers values
func NewGetBackupJobByIDOK() *GetBackupJobByIDOK {
	return &GetBackupJobByIDOK{}
}

/*
GetBackupJobByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetBackupJobByIDOK struct {
	Payload []*models.BackupJobWrapper
}

// IsSuccess returns true when this get backup job by Id o k response has a 2xx status code
func (o *GetBackupJobByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup job by Id o k response has a 3xx status code
func (o *GetBackupJobByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup job by Id o k response has a 4xx status code
func (o *GetBackupJobByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup job by Id o k response has a 5xx status code
func (o *GetBackupJobByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup job by Id o k response a status code equal to that given
func (o *GetBackupJobByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get backup job by Id o k response
func (o *GetBackupJobByIDOK) Code() int {
	return 200
}

func (o *GetBackupJobByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupjobs/{id}][%d] getBackupJobByIdOK %s", 200, payload)
}

func (o *GetBackupJobByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupjobs/{id}][%d] getBackupJobByIdOK %s", 200, payload)
}

func (o *GetBackupJobByIDOK) GetPayload() []*models.BackupJobWrapper {
	return o.Payload
}

func (o *GetBackupJobByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupJobByIDDefault creates a GetBackupJobByIDDefault with default headers values
func NewGetBackupJobByIDDefault(code int) *GetBackupJobByIDDefault {
	return &GetBackupJobByIDDefault{
		_statusCode: code,
	}
}

/*
GetBackupJobByIDDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetBackupJobByIDDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get backup job by Id default response has a 2xx status code
func (o *GetBackupJobByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get backup job by Id default response has a 3xx status code
func (o *GetBackupJobByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get backup job by Id default response has a 4xx status code
func (o *GetBackupJobByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get backup job by Id default response has a 5xx status code
func (o *GetBackupJobByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get backup job by Id default response a status code equal to that given
func (o *GetBackupJobByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get backup job by Id default response
func (o *GetBackupJobByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetBackupJobByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupjobs/{id}][%d] GetBackupJobById default %s", o._statusCode, payload)
}

func (o *GetBackupJobByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupjobs/{id}][%d] GetBackupJobById default %s", o._statusCode, payload)
}

func (o *GetBackupJobByIDDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetBackupJobByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
