// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// ImportCrlFileReader is a Reader for the ImportCrlFile structure.
type ImportCrlFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportCrlFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewImportCrlFileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewImportCrlFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImportCrlFileNoContent creates a ImportCrlFileNoContent with default headers values
func NewImportCrlFileNoContent() *ImportCrlFileNoContent {
	return &ImportCrlFileNoContent{}
}

/*
ImportCrlFileNoContent describes a response with status code 204, with default header values.

No Content
*/
type ImportCrlFileNoContent struct {
}

// IsSuccess returns true when this import crl file no content response has a 2xx status code
func (o *ImportCrlFileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import crl file no content response has a 3xx status code
func (o *ImportCrlFileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import crl file no content response has a 4xx status code
func (o *ImportCrlFileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this import crl file no content response has a 5xx status code
func (o *ImportCrlFileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this import crl file no content response a status code equal to that given
func (o *ImportCrlFileNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the import crl file no content response
func (o *ImportCrlFileNoContent) Code() int {
	return 204
}

func (o *ImportCrlFileNoContent) Error() string {
	return fmt.Sprintf("[PUT /clusters/import-crl-file][%d] importCrlFileNoContent", 204)
}

func (o *ImportCrlFileNoContent) String() string {
	return fmt.Sprintf("[PUT /clusters/import-crl-file][%d] importCrlFileNoContent", 204)
}

func (o *ImportCrlFileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportCrlFileDefault creates a ImportCrlFileDefault with default headers values
func NewImportCrlFileDefault(code int) *ImportCrlFileDefault {
	return &ImportCrlFileDefault{
		_statusCode: code,
	}
}

/*
ImportCrlFileDefault describes a response with status code -1, with default header values.

Error
*/
type ImportCrlFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this import crl file default response has a 2xx status code
func (o *ImportCrlFileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this import crl file default response has a 3xx status code
func (o *ImportCrlFileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this import crl file default response has a 4xx status code
func (o *ImportCrlFileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this import crl file default response has a 5xx status code
func (o *ImportCrlFileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this import crl file default response a status code equal to that given
func (o *ImportCrlFileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the import crl file default response
func (o *ImportCrlFileDefault) Code() int {
	return o._statusCode
}

func (o *ImportCrlFileDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/import-crl-file][%d] ImportCrlFile default %s", o._statusCode, payload)
}

func (o *ImportCrlFileDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/import-crl-file][%d] ImportCrlFile default %s", o._statusCode, payload)
}

func (o *ImportCrlFileDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportCrlFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
