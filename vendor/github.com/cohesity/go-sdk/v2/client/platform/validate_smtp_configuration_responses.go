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

// ValidateSMTPConfigurationReader is a Reader for the ValidateSMTPConfiguration structure.
type ValidateSMTPConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateSMTPConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateSMTPConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateSMTPConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateSMTPConfigurationNoContent creates a ValidateSMTPConfigurationNoContent with default headers values
func NewValidateSMTPConfigurationNoContent() *ValidateSMTPConfigurationNoContent {
	return &ValidateSMTPConfigurationNoContent{}
}

/*
ValidateSMTPConfigurationNoContent describes a response with status code 204, with default header values.

No Content
*/
type ValidateSMTPConfigurationNoContent struct {
}

// IsSuccess returns true when this validate Smtp configuration no content response has a 2xx status code
func (o *ValidateSMTPConfigurationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate Smtp configuration no content response has a 3xx status code
func (o *ValidateSMTPConfigurationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate Smtp configuration no content response has a 4xx status code
func (o *ValidateSMTPConfigurationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate Smtp configuration no content response has a 5xx status code
func (o *ValidateSMTPConfigurationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this validate Smtp configuration no content response a status code equal to that given
func (o *ValidateSMTPConfigurationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the validate Smtp configuration no content response
func (o *ValidateSMTPConfigurationNoContent) Code() int {
	return 204
}

func (o *ValidateSMTPConfigurationNoContent) Error() string {
	return fmt.Sprintf("[POST /clusters/smtp/validate][%d] validateSmtpConfigurationNoContent", 204)
}

func (o *ValidateSMTPConfigurationNoContent) String() string {
	return fmt.Sprintf("[POST /clusters/smtp/validate][%d] validateSmtpConfigurationNoContent", 204)
}

func (o *ValidateSMTPConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateSMTPConfigurationDefault creates a ValidateSMTPConfigurationDefault with default headers values
func NewValidateSMTPConfigurationDefault(code int) *ValidateSMTPConfigurationDefault {
	return &ValidateSMTPConfigurationDefault{
		_statusCode: code,
	}
}

/*
ValidateSMTPConfigurationDefault describes a response with status code -1, with default header values.

Error
*/
type ValidateSMTPConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this validate SMTP configuration default response has a 2xx status code
func (o *ValidateSMTPConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate SMTP configuration default response has a 3xx status code
func (o *ValidateSMTPConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate SMTP configuration default response has a 4xx status code
func (o *ValidateSMTPConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate SMTP configuration default response has a 5xx status code
func (o *ValidateSMTPConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate SMTP configuration default response a status code equal to that given
func (o *ValidateSMTPConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the validate SMTP configuration default response
func (o *ValidateSMTPConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *ValidateSMTPConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters/smtp/validate][%d] ValidateSMTPConfiguration default %s", o._statusCode, payload)
}

func (o *ValidateSMTPConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters/smtp/validate][%d] ValidateSMTPConfiguration default %s", o._statusCode, payload)
}

func (o *ValidateSMTPConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateSMTPConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
