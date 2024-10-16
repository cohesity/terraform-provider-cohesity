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

// UpdateIsDMaaSClusterReader is a Reader for the UpdateIsDMaaSCluster structure.
type UpdateIsDMaaSClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIsDMaaSClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIsDMaaSClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateIsDMaaSClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIsDMaaSClusterOK creates a UpdateIsDMaaSClusterOK with default headers values
func NewUpdateIsDMaaSClusterOK() *UpdateIsDMaaSClusterOK {
	return &UpdateIsDMaaSClusterOK{}
}

/*
UpdateIsDMaaSClusterOK describes a response with status code 200, with default header values.

Success
*/
type UpdateIsDMaaSClusterOK struct {
	Payload *models.DMaaSInfo
}

// IsSuccess returns true when this update is d maa s cluster o k response has a 2xx status code
func (o *UpdateIsDMaaSClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update is d maa s cluster o k response has a 3xx status code
func (o *UpdateIsDMaaSClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update is d maa s cluster o k response has a 4xx status code
func (o *UpdateIsDMaaSClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update is d maa s cluster o k response has a 5xx status code
func (o *UpdateIsDMaaSClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update is d maa s cluster o k response a status code equal to that given
func (o *UpdateIsDMaaSClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update is d maa s cluster o k response
func (o *UpdateIsDMaaSClusterOK) Code() int {
	return 200
}

func (o *UpdateIsDMaaSClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/is-dmaas][%d] updateIsDMaaSClusterOK %s", 200, payload)
}

func (o *UpdateIsDMaaSClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/is-dmaas][%d] updateIsDMaaSClusterOK %s", 200, payload)
}

func (o *UpdateIsDMaaSClusterOK) GetPayload() *models.DMaaSInfo {
	return o.Payload
}

func (o *UpdateIsDMaaSClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DMaaSInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIsDMaaSClusterDefault creates a UpdateIsDMaaSClusterDefault with default headers values
func NewUpdateIsDMaaSClusterDefault(code int) *UpdateIsDMaaSClusterDefault {
	return &UpdateIsDMaaSClusterDefault{
		_statusCode: code,
	}
}

/*
UpdateIsDMaaSClusterDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateIsDMaaSClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update is d maa s cluster default response has a 2xx status code
func (o *UpdateIsDMaaSClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update is d maa s cluster default response has a 3xx status code
func (o *UpdateIsDMaaSClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update is d maa s cluster default response has a 4xx status code
func (o *UpdateIsDMaaSClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update is d maa s cluster default response has a 5xx status code
func (o *UpdateIsDMaaSClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update is d maa s cluster default response a status code equal to that given
func (o *UpdateIsDMaaSClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update is d maa s cluster default response
func (o *UpdateIsDMaaSClusterDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIsDMaaSClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/is-dmaas][%d] UpdateIsDMaaSCluster default %s", o._statusCode, payload)
}

func (o *UpdateIsDMaaSClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/is-dmaas][%d] UpdateIsDMaaSCluster default %s", o._statusCode, payload)
}

func (o *UpdateIsDMaaSClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateIsDMaaSClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
