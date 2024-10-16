// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

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

// GetOneDriveDocumentsReader is a Reader for the GetOneDriveDocuments structure.
type GetOneDriveDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOneDriveDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOneDriveDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOneDriveDocumentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOneDriveDocumentsOK creates a GetOneDriveDocumentsOK with default headers values
func NewGetOneDriveDocumentsOK() *GetOneDriveDocumentsOK {
	return &GetOneDriveDocumentsOK{}
}

/*
GetOneDriveDocumentsOK describes a response with status code 200, with default header values.

Success
*/
type GetOneDriveDocumentsOK struct {
	Payload *models.FileSearchResults
}

// IsSuccess returns true when this get one drive documents o k response has a 2xx status code
func (o *GetOneDriveDocumentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get one drive documents o k response has a 3xx status code
func (o *GetOneDriveDocumentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get one drive documents o k response has a 4xx status code
func (o *GetOneDriveDocumentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get one drive documents o k response has a 5xx status code
func (o *GetOneDriveDocumentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get one drive documents o k response a status code equal to that given
func (o *GetOneDriveDocumentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get one drive documents o k response
func (o *GetOneDriveDocumentsOK) Code() int {
	return 200
}

func (o *GetOneDriveDocumentsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/office365/onedrive/documents][%d] getOneDriveDocumentsOK %s", 200, payload)
}

func (o *GetOneDriveDocumentsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/office365/onedrive/documents][%d] getOneDriveDocumentsOK %s", 200, payload)
}

func (o *GetOneDriveDocumentsOK) GetPayload() *models.FileSearchResults {
	return o.Payload
}

func (o *GetOneDriveDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileSearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneDriveDocumentsDefault creates a GetOneDriveDocumentsDefault with default headers values
func NewGetOneDriveDocumentsDefault(code int) *GetOneDriveDocumentsDefault {
	return &GetOneDriveDocumentsDefault{
		_statusCode: code,
	}
}

/*
GetOneDriveDocumentsDefault describes a response with status code -1, with default header values.

Error
*/
type GetOneDriveDocumentsDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get one drive documents default response has a 2xx status code
func (o *GetOneDriveDocumentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get one drive documents default response has a 3xx status code
func (o *GetOneDriveDocumentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get one drive documents default response has a 4xx status code
func (o *GetOneDriveDocumentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get one drive documents default response has a 5xx status code
func (o *GetOneDriveDocumentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get one drive documents default response a status code equal to that given
func (o *GetOneDriveDocumentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get one drive documents default response
func (o *GetOneDriveDocumentsDefault) Code() int {
	return o._statusCode
}

func (o *GetOneDriveDocumentsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/office365/onedrive/documents][%d] GetOneDriveDocuments default %s", o._statusCode, payload)
}

func (o *GetOneDriveDocumentsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/office365/onedrive/documents][%d] GetOneDriveDocuments default %s", o._statusCode, payload)
}

func (o *GetOneDriveDocumentsDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetOneDriveDocumentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
