// Code generated by go-swagger; DO NOT EDIT.

package m_f_a

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

// VerifySupportUserTotpReader is a Reader for the VerifySupportUserTotp structure.
type VerifySupportUserTotpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifySupportUserTotpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerifySupportUserTotpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVerifySupportUserTotpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVerifySupportUserTotpOK creates a VerifySupportUserTotpOK with default headers values
func NewVerifySupportUserTotpOK() *VerifySupportUserTotpOK {
	return &VerifySupportUserTotpOK{}
}

/*
VerifySupportUserTotpOK describes a response with status code 200, with default header values.

Success
*/
type VerifySupportUserTotpOK struct {
	Payload *models.VerifyTotpResult
}

// IsSuccess returns true when this verify support user totp o k response has a 2xx status code
func (o *VerifySupportUserTotpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this verify support user totp o k response has a 3xx status code
func (o *VerifySupportUserTotpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this verify support user totp o k response has a 4xx status code
func (o *VerifySupportUserTotpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this verify support user totp o k response has a 5xx status code
func (o *VerifySupportUserTotpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this verify support user totp o k response a status code equal to that given
func (o *VerifySupportUserTotpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the verify support user totp o k response
func (o *VerifySupportUserTotpOK) Code() int {
	return 200
}

func (o *VerifySupportUserTotpOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support-user/verify-totp][%d] verifySupportUserTotpOK %s", 200, payload)
}

func (o *VerifySupportUserTotpOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support-user/verify-totp][%d] verifySupportUserTotpOK %s", 200, payload)
}

func (o *VerifySupportUserTotpOK) GetPayload() *models.VerifyTotpResult {
	return o.Payload
}

func (o *VerifySupportUserTotpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VerifyTotpResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifySupportUserTotpDefault creates a VerifySupportUserTotpDefault with default headers values
func NewVerifySupportUserTotpDefault(code int) *VerifySupportUserTotpDefault {
	return &VerifySupportUserTotpDefault{
		_statusCode: code,
	}
}

/*
VerifySupportUserTotpDefault describes a response with status code -1, with default header values.

Error
*/
type VerifySupportUserTotpDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this verify support user totp default response has a 2xx status code
func (o *VerifySupportUserTotpDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this verify support user totp default response has a 3xx status code
func (o *VerifySupportUserTotpDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this verify support user totp default response has a 4xx status code
func (o *VerifySupportUserTotpDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this verify support user totp default response has a 5xx status code
func (o *VerifySupportUserTotpDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this verify support user totp default response a status code equal to that given
func (o *VerifySupportUserTotpDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the verify support user totp default response
func (o *VerifySupportUserTotpDefault) Code() int {
	return o._statusCode
}

func (o *VerifySupportUserTotpDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support-user/verify-totp][%d] VerifySupportUserTotp default %s", o._statusCode, payload)
}

func (o *VerifySupportUserTotpDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /support-user/verify-totp][%d] VerifySupportUserTotp default %s", o._statusCode, payload)
}

func (o *VerifySupportUserTotpDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *VerifySupportUserTotpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
