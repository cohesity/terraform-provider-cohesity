// Code generated by go-swagger; DO NOT EDIT.

package recovery

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

// CreateDownloadFilesAndFoldersRecoveryReader is a Reader for the CreateDownloadFilesAndFoldersRecovery structure.
type CreateDownloadFilesAndFoldersRecoveryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDownloadFilesAndFoldersRecoveryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDownloadFilesAndFoldersRecoveryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDownloadFilesAndFoldersRecoveryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDownloadFilesAndFoldersRecoveryCreated creates a CreateDownloadFilesAndFoldersRecoveryCreated with default headers values
func NewCreateDownloadFilesAndFoldersRecoveryCreated() *CreateDownloadFilesAndFoldersRecoveryCreated {
	return &CreateDownloadFilesAndFoldersRecoveryCreated{}
}

/*
CreateDownloadFilesAndFoldersRecoveryCreated describes a response with status code 201, with default header values.

Success
*/
type CreateDownloadFilesAndFoldersRecoveryCreated struct {
	Payload *models.Recovery
}

// IsSuccess returns true when this create download files and folders recovery created response has a 2xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create download files and folders recovery created response has a 3xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create download files and folders recovery created response has a 4xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create download files and folders recovery created response has a 5xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create download files and folders recovery created response a status code equal to that given
func (o *CreateDownloadFilesAndFoldersRecoveryCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create download files and folders recovery created response
func (o *CreateDownloadFilesAndFoldersRecoveryCreated) Code() int {
	return 201
}

func (o *CreateDownloadFilesAndFoldersRecoveryCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries/download-files-folders][%d] createDownloadFilesAndFoldersRecoveryCreated %s", 201, payload)
}

func (o *CreateDownloadFilesAndFoldersRecoveryCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries/download-files-folders][%d] createDownloadFilesAndFoldersRecoveryCreated %s", 201, payload)
}

func (o *CreateDownloadFilesAndFoldersRecoveryCreated) GetPayload() *models.Recovery {
	return o.Payload
}

func (o *CreateDownloadFilesAndFoldersRecoveryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recovery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDownloadFilesAndFoldersRecoveryDefault creates a CreateDownloadFilesAndFoldersRecoveryDefault with default headers values
func NewCreateDownloadFilesAndFoldersRecoveryDefault(code int) *CreateDownloadFilesAndFoldersRecoveryDefault {
	return &CreateDownloadFilesAndFoldersRecoveryDefault{
		_statusCode: code,
	}
}

/*
CreateDownloadFilesAndFoldersRecoveryDefault describes a response with status code -1, with default header values.

Error
*/
type CreateDownloadFilesAndFoldersRecoveryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create download files and folders recovery default response has a 2xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create download files and folders recovery default response has a 3xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create download files and folders recovery default response has a 4xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create download files and folders recovery default response has a 5xx status code
func (o *CreateDownloadFilesAndFoldersRecoveryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create download files and folders recovery default response a status code equal to that given
func (o *CreateDownloadFilesAndFoldersRecoveryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create download files and folders recovery default response
func (o *CreateDownloadFilesAndFoldersRecoveryDefault) Code() int {
	return o._statusCode
}

func (o *CreateDownloadFilesAndFoldersRecoveryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries/download-files-folders][%d] CreateDownloadFilesAndFoldersRecovery default %s", o._statusCode, payload)
}

func (o *CreateDownloadFilesAndFoldersRecoveryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries/download-files-folders][%d] CreateDownloadFilesAndFoldersRecovery default %s", o._statusCode, payload)
}

func (o *CreateDownloadFilesAndFoldersRecoveryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDownloadFilesAndFoldersRecoveryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
