// Code generated by go-swagger; DO NOT EDIT.

package protection_runs

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

// UpdateProtectionRunsReader is a Reader for the UpdateProtectionRuns structure.
type UpdateProtectionRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProtectionRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateProtectionRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateProtectionRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProtectionRunsNoContent creates a UpdateProtectionRunsNoContent with default headers values
func NewUpdateProtectionRunsNoContent() *UpdateProtectionRunsNoContent {
	return &UpdateProtectionRunsNoContent{}
}

/*
UpdateProtectionRunsNoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateProtectionRunsNoContent struct {
}

// IsSuccess returns true when this update protection runs no content response has a 2xx status code
func (o *UpdateProtectionRunsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update protection runs no content response has a 3xx status code
func (o *UpdateProtectionRunsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update protection runs no content response has a 4xx status code
func (o *UpdateProtectionRunsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update protection runs no content response has a 5xx status code
func (o *UpdateProtectionRunsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update protection runs no content response a status code equal to that given
func (o *UpdateProtectionRunsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update protection runs no content response
func (o *UpdateProtectionRunsNoContent) Code() int {
	return 204
}

func (o *UpdateProtectionRunsNoContent) Error() string {
	return fmt.Sprintf("[PUT /public/protectionRuns][%d] updateProtectionRunsNoContent", 204)
}

func (o *UpdateProtectionRunsNoContent) String() string {
	return fmt.Sprintf("[PUT /public/protectionRuns][%d] updateProtectionRunsNoContent", 204)
}

func (o *UpdateProtectionRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProtectionRunsDefault creates a UpdateProtectionRunsDefault with default headers values
func NewUpdateProtectionRunsDefault(code int) *UpdateProtectionRunsDefault {
	return &UpdateProtectionRunsDefault{
		_statusCode: code,
	}
}

/*
UpdateProtectionRunsDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateProtectionRunsDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update protection runs default response has a 2xx status code
func (o *UpdateProtectionRunsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update protection runs default response has a 3xx status code
func (o *UpdateProtectionRunsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update protection runs default response has a 4xx status code
func (o *UpdateProtectionRunsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update protection runs default response has a 5xx status code
func (o *UpdateProtectionRunsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update protection runs default response a status code equal to that given
func (o *UpdateProtectionRunsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update protection runs default response
func (o *UpdateProtectionRunsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProtectionRunsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/protectionRuns][%d] UpdateProtectionRuns default %s", o._statusCode, payload)
}

func (o *UpdateProtectionRunsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/protectionRuns][%d] UpdateProtectionRuns default %s", o._statusCode, payload)
}

func (o *UpdateProtectionRunsDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateProtectionRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
