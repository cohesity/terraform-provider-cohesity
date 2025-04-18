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

// UpgradeCheckGetResultsReader is a Reader for the UpgradeCheckGetResults structure.
type UpgradeCheckGetResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeCheckGetResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpgradeCheckGetResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpgradeCheckGetResultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpgradeCheckGetResultsOK creates a UpgradeCheckGetResultsOK with default headers values
func NewUpgradeCheckGetResultsOK() *UpgradeCheckGetResultsOK {
	return &UpgradeCheckGetResultsOK{}
}

/*
UpgradeCheckGetResultsOK describes a response with status code 200, with default header values.

Success
*/
type UpgradeCheckGetResultsOK struct {
	Payload *models.UpgradeChecksResults
}

// IsSuccess returns true when this upgrade check get results o k response has a 2xx status code
func (o *UpgradeCheckGetResultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upgrade check get results o k response has a 3xx status code
func (o *UpgradeCheckGetResultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upgrade check get results o k response has a 4xx status code
func (o *UpgradeCheckGetResultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upgrade check get results o k response has a 5xx status code
func (o *UpgradeCheckGetResultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upgrade check get results o k response a status code equal to that given
func (o *UpgradeCheckGetResultsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upgrade check get results o k response
func (o *UpgradeCheckGetResultsOK) Code() int {
	return 200
}

func (o *UpgradeCheckGetResultsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/upgrade-checks/{testRunInstanceId}][%d] upgradeCheckGetResultsOK %s", 200, payload)
}

func (o *UpgradeCheckGetResultsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/upgrade-checks/{testRunInstanceId}][%d] upgradeCheckGetResultsOK %s", 200, payload)
}

func (o *UpgradeCheckGetResultsOK) GetPayload() *models.UpgradeChecksResults {
	return o.Payload
}

func (o *UpgradeCheckGetResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpgradeChecksResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpgradeCheckGetResultsDefault creates a UpgradeCheckGetResultsDefault with default headers values
func NewUpgradeCheckGetResultsDefault(code int) *UpgradeCheckGetResultsDefault {
	return &UpgradeCheckGetResultsDefault{
		_statusCode: code,
	}
}

/*
UpgradeCheckGetResultsDefault describes a response with status code -1, with default header values.

Error
*/
type UpgradeCheckGetResultsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this upgrade check get results default response has a 2xx status code
func (o *UpgradeCheckGetResultsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this upgrade check get results default response has a 3xx status code
func (o *UpgradeCheckGetResultsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this upgrade check get results default response has a 4xx status code
func (o *UpgradeCheckGetResultsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this upgrade check get results default response has a 5xx status code
func (o *UpgradeCheckGetResultsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this upgrade check get results default response a status code equal to that given
func (o *UpgradeCheckGetResultsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the upgrade check get results default response
func (o *UpgradeCheckGetResultsDefault) Code() int {
	return o._statusCode
}

func (o *UpgradeCheckGetResultsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/upgrade-checks/{testRunInstanceId}][%d] UpgradeCheckGetResults default %s", o._statusCode, payload)
}

func (o *UpgradeCheckGetResultsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cluster/upgrade-checks/{testRunInstanceId}][%d] UpgradeCheckGetResults default %s", o._statusCode, payload)
}

func (o *UpgradeCheckGetResultsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpgradeCheckGetResultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
