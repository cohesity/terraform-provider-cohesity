// Code generated by go-swagger; DO NOT EDIT.

package athena_app_ops

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

// RelaunchAppInstanceReader is a Reader for the RelaunchAppInstance structure.
type RelaunchAppInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RelaunchAppInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRelaunchAppInstanceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRelaunchAppInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRelaunchAppInstanceAccepted creates a RelaunchAppInstanceAccepted with default headers values
func NewRelaunchAppInstanceAccepted() *RelaunchAppInstanceAccepted {
	return &RelaunchAppInstanceAccepted{}
}

/*
RelaunchAppInstanceAccepted describes a response with status code 202, with default header values.

Request Accepted
*/
type RelaunchAppInstanceAccepted struct {
}

// IsSuccess returns true when this relaunch app instance accepted response has a 2xx status code
func (o *RelaunchAppInstanceAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this relaunch app instance accepted response has a 3xx status code
func (o *RelaunchAppInstanceAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this relaunch app instance accepted response has a 4xx status code
func (o *RelaunchAppInstanceAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this relaunch app instance accepted response has a 5xx status code
func (o *RelaunchAppInstanceAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this relaunch app instance accepted response a status code equal to that given
func (o *RelaunchAppInstanceAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the relaunch app instance accepted response
func (o *RelaunchAppInstanceAccepted) Code() int {
	return 202
}

func (o *RelaunchAppInstanceAccepted) Error() string {
	return fmt.Sprintf("[PUT /apps/instances/{id}/relaunch][%d] relaunchAppInstanceAccepted", 202)
}

func (o *RelaunchAppInstanceAccepted) String() string {
	return fmt.Sprintf("[PUT /apps/instances/{id}/relaunch][%d] relaunchAppInstanceAccepted", 202)
}

func (o *RelaunchAppInstanceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRelaunchAppInstanceDefault creates a RelaunchAppInstanceDefault with default headers values
func NewRelaunchAppInstanceDefault(code int) *RelaunchAppInstanceDefault {
	return &RelaunchAppInstanceDefault{
		_statusCode: code,
	}
}

/*
RelaunchAppInstanceDefault describes a response with status code -1, with default header values.

Error
*/
type RelaunchAppInstanceDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this relaunch app instance default response has a 2xx status code
func (o *RelaunchAppInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this relaunch app instance default response has a 3xx status code
func (o *RelaunchAppInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this relaunch app instance default response has a 4xx status code
func (o *RelaunchAppInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this relaunch app instance default response has a 5xx status code
func (o *RelaunchAppInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this relaunch app instance default response a status code equal to that given
func (o *RelaunchAppInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the relaunch app instance default response
func (o *RelaunchAppInstanceDefault) Code() int {
	return o._statusCode
}

func (o *RelaunchAppInstanceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /apps/instances/{id}/relaunch][%d] RelaunchAppInstance default %s", o._statusCode, payload)
}

func (o *RelaunchAppInstanceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /apps/instances/{id}/relaunch][%d] RelaunchAppInstance default %s", o._statusCode, payload)
}

func (o *RelaunchAppInstanceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RelaunchAppInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
