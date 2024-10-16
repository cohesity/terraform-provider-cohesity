// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

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

// SearchProductionAdObjectsReader is a Reader for the SearchProductionAdObjects structure.
type SearchProductionAdObjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchProductionAdObjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchProductionAdObjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchProductionAdObjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchProductionAdObjectsOK creates a SearchProductionAdObjectsOK with default headers values
func NewSearchProductionAdObjectsOK() *SearchProductionAdObjectsOK {
	return &SearchProductionAdObjectsOK{}
}

/*
SearchProductionAdObjectsOK describes a response with status code 200, with default header values.

Success
*/
type SearchProductionAdObjectsOK struct {
	Payload []*models.ADObject
}

// IsSuccess returns true when this search production ad objects o k response has a 2xx status code
func (o *SearchProductionAdObjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search production ad objects o k response has a 3xx status code
func (o *SearchProductionAdObjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search production ad objects o k response has a 4xx status code
func (o *SearchProductionAdObjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search production ad objects o k response has a 5xx status code
func (o *SearchProductionAdObjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search production ad objects o k response a status code equal to that given
func (o *SearchProductionAdObjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search production ad objects o k response
func (o *SearchProductionAdObjectsOK) Code() int {
	return 200
}

func (o *SearchProductionAdObjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/restore/adObjects][%d] searchProductionAdObjectsOK %s", 200, payload)
}

func (o *SearchProductionAdObjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/restore/adObjects][%d] searchProductionAdObjectsOK %s", 200, payload)
}

func (o *SearchProductionAdObjectsOK) GetPayload() []*models.ADObject {
	return o.Payload
}

func (o *SearchProductionAdObjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchProductionAdObjectsDefault creates a SearchProductionAdObjectsDefault with default headers values
func NewSearchProductionAdObjectsDefault(code int) *SearchProductionAdObjectsDefault {
	return &SearchProductionAdObjectsDefault{
		_statusCode: code,
	}
}

/*
SearchProductionAdObjectsDefault describes a response with status code -1, with default header values.

Error
*/
type SearchProductionAdObjectsDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this search production ad objects default response has a 2xx status code
func (o *SearchProductionAdObjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search production ad objects default response has a 3xx status code
func (o *SearchProductionAdObjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search production ad objects default response has a 4xx status code
func (o *SearchProductionAdObjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search production ad objects default response has a 5xx status code
func (o *SearchProductionAdObjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search production ad objects default response a status code equal to that given
func (o *SearchProductionAdObjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search production ad objects default response
func (o *SearchProductionAdObjectsDefault) Code() int {
	return o._statusCode
}

func (o *SearchProductionAdObjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/restore/adObjects][%d] SearchProductionAdObjects default %s", o._statusCode, payload)
}

func (o *SearchProductionAdObjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/restore/adObjects][%d] SearchProductionAdObjects default %s", o._statusCode, payload)
}

func (o *SearchProductionAdObjectsDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *SearchProductionAdObjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
