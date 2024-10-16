// Code generated by go-swagger; DO NOT EDIT.

package baseos_patch_management

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

// ApplyBaseosPatchReader is a Reader for the ApplyBaseosPatch structure.
type ApplyBaseosPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplyBaseosPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplyBaseosPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplyBaseosPatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplyBaseosPatchOK creates a ApplyBaseosPatchOK with default headers values
func NewApplyBaseosPatchOK() *ApplyBaseosPatchOK {
	return &ApplyBaseosPatchOK{}
}

/*
ApplyBaseosPatchOK describes a response with status code 200, with default header values.

No Content
*/
type ApplyBaseosPatchOK struct {
}

// IsSuccess returns true when this apply baseos patch o k response has a 2xx status code
func (o *ApplyBaseosPatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this apply baseos patch o k response has a 3xx status code
func (o *ApplyBaseosPatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this apply baseos patch o k response has a 4xx status code
func (o *ApplyBaseosPatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this apply baseos patch o k response has a 5xx status code
func (o *ApplyBaseosPatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this apply baseos patch o k response a status code equal to that given
func (o *ApplyBaseosPatchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the apply baseos patch o k response
func (o *ApplyBaseosPatchOK) Code() int {
	return 200
}

func (o *ApplyBaseosPatchOK) Error() string {
	return fmt.Sprintf("[POST /patch-management/baseos-patch-apply][%d] applyBaseosPatchOK", 200)
}

func (o *ApplyBaseosPatchOK) String() string {
	return fmt.Sprintf("[POST /patch-management/baseos-patch-apply][%d] applyBaseosPatchOK", 200)
}

func (o *ApplyBaseosPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplyBaseosPatchDefault creates a ApplyBaseosPatchDefault with default headers values
func NewApplyBaseosPatchDefault(code int) *ApplyBaseosPatchDefault {
	return &ApplyBaseosPatchDefault{
		_statusCode: code,
	}
}

/*
ApplyBaseosPatchDefault describes a response with status code -1, with default header values.

Error
*/
type ApplyBaseosPatchDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this apply baseos patch default response has a 2xx status code
func (o *ApplyBaseosPatchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this apply baseos patch default response has a 3xx status code
func (o *ApplyBaseosPatchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this apply baseos patch default response has a 4xx status code
func (o *ApplyBaseosPatchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this apply baseos patch default response has a 5xx status code
func (o *ApplyBaseosPatchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this apply baseos patch default response a status code equal to that given
func (o *ApplyBaseosPatchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the apply baseos patch default response
func (o *ApplyBaseosPatchDefault) Code() int {
	return o._statusCode
}

func (o *ApplyBaseosPatchDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /patch-management/baseos-patch-apply][%d] ApplyBaseosPatch default %s", o._statusCode, payload)
}

func (o *ApplyBaseosPatchDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /patch-management/baseos-patch-apply][%d] ApplyBaseosPatch default %s", o._statusCode, payload)
}

func (o *ApplyBaseosPatchDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ApplyBaseosPatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
