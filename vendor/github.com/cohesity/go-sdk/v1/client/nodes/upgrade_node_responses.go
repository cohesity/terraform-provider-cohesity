// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// UpgradeNodeReader is a Reader for the UpgradeNode structure.
type UpgradeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpgradeNodeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpgradeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpgradeNodeAccepted creates a UpgradeNodeAccepted with default headers values
func NewUpgradeNodeAccepted() *UpgradeNodeAccepted {
	return &UpgradeNodeAccepted{}
}

/*
UpgradeNodeAccepted describes a response with status code 202, with default header values.

Success
*/
type UpgradeNodeAccepted struct {
	Payload *models.UpgradeNodeResult
}

// IsSuccess returns true when this upgrade node accepted response has a 2xx status code
func (o *UpgradeNodeAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upgrade node accepted response has a 3xx status code
func (o *UpgradeNodeAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upgrade node accepted response has a 4xx status code
func (o *UpgradeNodeAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this upgrade node accepted response has a 5xx status code
func (o *UpgradeNodeAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this upgrade node accepted response a status code equal to that given
func (o *UpgradeNodeAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the upgrade node accepted response
func (o *UpgradeNodeAccepted) Code() int {
	return 202
}

func (o *UpgradeNodeAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/nodes/software][%d] upgradeNodeAccepted %s", 202, payload)
}

func (o *UpgradeNodeAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/nodes/software][%d] upgradeNodeAccepted %s", 202, payload)
}

func (o *UpgradeNodeAccepted) GetPayload() *models.UpgradeNodeResult {
	return o.Payload
}

func (o *UpgradeNodeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpgradeNodeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpgradeNodeDefault creates a UpgradeNodeDefault with default headers values
func NewUpgradeNodeDefault(code int) *UpgradeNodeDefault {
	return &UpgradeNodeDefault{
		_statusCode: code,
	}
}

/*
UpgradeNodeDefault describes a response with status code -1, with default header values.

Error
*/
type UpgradeNodeDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this upgrade node default response has a 2xx status code
func (o *UpgradeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upgrade node default response has a 3xx status code
func (o *UpgradeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upgrade node default response has a 4xx status code
func (o *UpgradeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upgrade node default response has a 5xx status code
func (o *UpgradeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upgrade node default response a status code equal to that given
func (o *UpgradeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the upgrade node default response
func (o *UpgradeNodeDefault) Code() int {
	return o._statusCode
}

func (o *UpgradeNodeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/nodes/software][%d] UpgradeNode default %s", o._statusCode, payload)
}

func (o *UpgradeNodeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/nodes/software][%d] UpgradeNode default %s", o._statusCode, payload)
}

func (o *UpgradeNodeDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpgradeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
