// Code generated by go-swagger; DO NOT EDIT.

package external_target

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

// GetExternalTargetSettingsReader is a Reader for the GetExternalTargetSettings structure.
type GetExternalTargetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalTargetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalTargetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetExternalTargetSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalTargetSettingsOK creates a GetExternalTargetSettingsOK with default headers values
func NewGetExternalTargetSettingsOK() *GetExternalTargetSettingsOK {
	return &GetExternalTargetSettingsOK{}
}

/*
GetExternalTargetSettingsOK describes a response with status code 200, with default header values.

Success
*/
type GetExternalTargetSettingsOK struct {
	Payload *models.ExternalTarget
}

// IsSuccess returns true when this get external target settings o k response has a 2xx status code
func (o *GetExternalTargetSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get external target settings o k response has a 3xx status code
func (o *GetExternalTargetSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external target settings o k response has a 4xx status code
func (o *GetExternalTargetSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get external target settings o k response has a 5xx status code
func (o *GetExternalTargetSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get external target settings o k response a status code equal to that given
func (o *GetExternalTargetSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get external target settings o k response
func (o *GetExternalTargetSettingsOK) Code() int {
	return 200
}

func (o *GetExternalTargetSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets/settings][%d] getExternalTargetSettingsOK %s", 200, payload)
}

func (o *GetExternalTargetSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets/settings][%d] getExternalTargetSettingsOK %s", 200, payload)
}

func (o *GetExternalTargetSettingsOK) GetPayload() *models.ExternalTarget {
	return o.Payload
}

func (o *GetExternalTargetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalTargetSettingsDefault creates a GetExternalTargetSettingsDefault with default headers values
func NewGetExternalTargetSettingsDefault(code int) *GetExternalTargetSettingsDefault {
	return &GetExternalTargetSettingsDefault{
		_statusCode: code,
	}
}

/*
GetExternalTargetSettingsDefault describes a response with status code -1, with default header values.

Error
*/
type GetExternalTargetSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get external target settings default response has a 2xx status code
func (o *GetExternalTargetSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get external target settings default response has a 3xx status code
func (o *GetExternalTargetSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get external target settings default response has a 4xx status code
func (o *GetExternalTargetSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get external target settings default response has a 5xx status code
func (o *GetExternalTargetSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get external target settings default response a status code equal to that given
func (o *GetExternalTargetSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get external target settings default response
func (o *GetExternalTargetSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetExternalTargetSettingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets/settings][%d] GetExternalTargetSettings default %s", o._statusCode, payload)
}

func (o *GetExternalTargetSettingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/external-targets/settings][%d] GetExternalTargetSettings default %s", o._statusCode, payload)
}

func (o *GetExternalTargetSettingsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExternalTargetSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
