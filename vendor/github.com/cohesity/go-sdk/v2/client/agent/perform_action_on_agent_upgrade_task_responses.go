// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// PerformActionOnAgentUpgradeTaskReader is a Reader for the PerformActionOnAgentUpgradeTask structure.
type PerformActionOnAgentUpgradeTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformActionOnAgentUpgradeTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPerformActionOnAgentUpgradeTaskCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformActionOnAgentUpgradeTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformActionOnAgentUpgradeTaskCreated creates a PerformActionOnAgentUpgradeTaskCreated with default headers values
func NewPerformActionOnAgentUpgradeTaskCreated() *PerformActionOnAgentUpgradeTaskCreated {
	return &PerformActionOnAgentUpgradeTaskCreated{}
}

/*
PerformActionOnAgentUpgradeTaskCreated describes a response with status code 201, with default header values.

Success
*/
type PerformActionOnAgentUpgradeTaskCreated struct {
	Payload *models.AgentUpgradeTaskActionObject
}

// IsSuccess returns true when this perform action on agent upgrade task created response has a 2xx status code
func (o *PerformActionOnAgentUpgradeTaskCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this perform action on agent upgrade task created response has a 3xx status code
func (o *PerformActionOnAgentUpgradeTaskCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform action on agent upgrade task created response has a 4xx status code
func (o *PerformActionOnAgentUpgradeTaskCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform action on agent upgrade task created response has a 5xx status code
func (o *PerformActionOnAgentUpgradeTaskCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this perform action on agent upgrade task created response a status code equal to that given
func (o *PerformActionOnAgentUpgradeTaskCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the perform action on agent upgrade task created response
func (o *PerformActionOnAgentUpgradeTaskCreated) Code() int {
	return 201
}

func (o *PerformActionOnAgentUpgradeTaskCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/agents/upgrade-tasks/actions][%d] performActionOnAgentUpgradeTaskCreated %s", 201, payload)
}

func (o *PerformActionOnAgentUpgradeTaskCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/agents/upgrade-tasks/actions][%d] performActionOnAgentUpgradeTaskCreated %s", 201, payload)
}

func (o *PerformActionOnAgentUpgradeTaskCreated) GetPayload() *models.AgentUpgradeTaskActionObject {
	return o.Payload
}

func (o *PerformActionOnAgentUpgradeTaskCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentUpgradeTaskActionObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformActionOnAgentUpgradeTaskDefault creates a PerformActionOnAgentUpgradeTaskDefault with default headers values
func NewPerformActionOnAgentUpgradeTaskDefault(code int) *PerformActionOnAgentUpgradeTaskDefault {
	return &PerformActionOnAgentUpgradeTaskDefault{
		_statusCode: code,
	}
}

/*
PerformActionOnAgentUpgradeTaskDefault describes a response with status code -1, with default header values.

Error
*/
type PerformActionOnAgentUpgradeTaskDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this perform action on agent upgrade task default response has a 2xx status code
func (o *PerformActionOnAgentUpgradeTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this perform action on agent upgrade task default response has a 3xx status code
func (o *PerformActionOnAgentUpgradeTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this perform action on agent upgrade task default response has a 4xx status code
func (o *PerformActionOnAgentUpgradeTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this perform action on agent upgrade task default response has a 5xx status code
func (o *PerformActionOnAgentUpgradeTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this perform action on agent upgrade task default response a status code equal to that given
func (o *PerformActionOnAgentUpgradeTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the perform action on agent upgrade task default response
func (o *PerformActionOnAgentUpgradeTaskDefault) Code() int {
	return o._statusCode
}

func (o *PerformActionOnAgentUpgradeTaskDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/agents/upgrade-tasks/actions][%d] PerformActionOnAgentUpgradeTask default %s", o._statusCode, payload)
}

func (o *PerformActionOnAgentUpgradeTaskDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/agents/upgrade-tasks/actions][%d] PerformActionOnAgentUpgradeTask default %s", o._statusCode, payload)
}

func (o *PerformActionOnAgentUpgradeTaskDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PerformActionOnAgentUpgradeTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
