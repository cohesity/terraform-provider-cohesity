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

// UpdateAMQPTargetConfigReader is a Reader for the UpdateAMQPTargetConfig structure.
type UpdateAMQPTargetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAMQPTargetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAMQPTargetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAMQPTargetConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAMQPTargetConfigOK creates a UpdateAMQPTargetConfigOK with default headers values
func NewUpdateAMQPTargetConfigOK() *UpdateAMQPTargetConfigOK {
	return &UpdateAMQPTargetConfigOK{}
}

/*
UpdateAMQPTargetConfigOK describes a response with status code 200, with default header values.

Success
*/
type UpdateAMQPTargetConfigOK struct {
	Payload *models.ClusterAMQPTargetConfig
}

// IsSuccess returns true when this update a m q p target config o k response has a 2xx status code
func (o *UpdateAMQPTargetConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update a m q p target config o k response has a 3xx status code
func (o *UpdateAMQPTargetConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update a m q p target config o k response has a 4xx status code
func (o *UpdateAMQPTargetConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update a m q p target config o k response has a 5xx status code
func (o *UpdateAMQPTargetConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update a m q p target config o k response a status code equal to that given
func (o *UpdateAMQPTargetConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update a m q p target config o k response
func (o *UpdateAMQPTargetConfigOK) Code() int {
	return 200
}

func (o *UpdateAMQPTargetConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/amqp-target-config][%d] updateAMQPTargetConfigOK %s", 200, payload)
}

func (o *UpdateAMQPTargetConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/amqp-target-config][%d] updateAMQPTargetConfigOK %s", 200, payload)
}

func (o *UpdateAMQPTargetConfigOK) GetPayload() *models.ClusterAMQPTargetConfig {
	return o.Payload
}

func (o *UpdateAMQPTargetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterAMQPTargetConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAMQPTargetConfigDefault creates a UpdateAMQPTargetConfigDefault with default headers values
func NewUpdateAMQPTargetConfigDefault(code int) *UpdateAMQPTargetConfigDefault {
	return &UpdateAMQPTargetConfigDefault{
		_statusCode: code,
	}
}

/*
UpdateAMQPTargetConfigDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateAMQPTargetConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update a m q p target config default response has a 2xx status code
func (o *UpdateAMQPTargetConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update a m q p target config default response has a 3xx status code
func (o *UpdateAMQPTargetConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update a m q p target config default response has a 4xx status code
func (o *UpdateAMQPTargetConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update a m q p target config default response has a 5xx status code
func (o *UpdateAMQPTargetConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update a m q p target config default response a status code equal to that given
func (o *UpdateAMQPTargetConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update a m q p target config default response
func (o *UpdateAMQPTargetConfigDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAMQPTargetConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/amqp-target-config][%d] UpdateAMQPTargetConfig default %s", o._statusCode, payload)
}

func (o *UpdateAMQPTargetConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /clusters/amqp-target-config][%d] UpdateAMQPTargetConfig default %s", o._statusCode, payload)
}

func (o *UpdateAMQPTargetConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAMQPTargetConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
