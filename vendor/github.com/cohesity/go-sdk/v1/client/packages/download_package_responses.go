// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// DownloadPackageReader is a Reader for the DownloadPackage structure.
type DownloadPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDownloadPackageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDownloadPackageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDownloadPackageAccepted creates a DownloadPackageAccepted with default headers values
func NewDownloadPackageAccepted() *DownloadPackageAccepted {
	return &DownloadPackageAccepted{}
}

/*
DownloadPackageAccepted describes a response with status code 202, with default header values.

Success
*/
type DownloadPackageAccepted struct {
	Payload *models.DownloadPackageResult
}

// IsSuccess returns true when this download package accepted response has a 2xx status code
func (o *DownloadPackageAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download package accepted response has a 3xx status code
func (o *DownloadPackageAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download package accepted response has a 4xx status code
func (o *DownloadPackageAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this download package accepted response has a 5xx status code
func (o *DownloadPackageAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this download package accepted response a status code equal to that given
func (o *DownloadPackageAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the download package accepted response
func (o *DownloadPackageAccepted) Code() int {
	return 202
}

func (o *DownloadPackageAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/packages/url][%d] downloadPackageAccepted %s", 202, payload)
}

func (o *DownloadPackageAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/packages/url][%d] downloadPackageAccepted %s", 202, payload)
}

func (o *DownloadPackageAccepted) GetPayload() *models.DownloadPackageResult {
	return o.Payload
}

func (o *DownloadPackageAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DownloadPackageResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadPackageDefault creates a DownloadPackageDefault with default headers values
func NewDownloadPackageDefault(code int) *DownloadPackageDefault {
	return &DownloadPackageDefault{
		_statusCode: code,
	}
}

/*
DownloadPackageDefault describes a response with status code -1, with default header values.

Error
*/
type DownloadPackageDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this download package default response has a 2xx status code
func (o *DownloadPackageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this download package default response has a 3xx status code
func (o *DownloadPackageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this download package default response has a 4xx status code
func (o *DownloadPackageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this download package default response has a 5xx status code
func (o *DownloadPackageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this download package default response a status code equal to that given
func (o *DownloadPackageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the download package default response
func (o *DownloadPackageDefault) Code() int {
	return o._statusCode
}

func (o *DownloadPackageDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/packages/url][%d] DownloadPackage default %s", o._statusCode, payload)
}

func (o *DownloadPackageDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/packages/url][%d] DownloadPackage default %s", o._statusCode, payload)
}

func (o *DownloadPackageDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *DownloadPackageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
