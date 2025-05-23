// Code generated by go-swagger; DO NOT EDIT.

package object

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

// GetSourceHierarchyObjectsReader is a Reader for the GetSourceHierarchyObjects structure.
type GetSourceHierarchyObjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceHierarchyObjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceHierarchyObjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSourceHierarchyObjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSourceHierarchyObjectsOK creates a GetSourceHierarchyObjectsOK with default headers values
func NewGetSourceHierarchyObjectsOK() *GetSourceHierarchyObjectsOK {
	return &GetSourceHierarchyObjectsOK{}
}

/*
GetSourceHierarchyObjectsOK describes a response with status code 200, with default header values.

Success
*/
type GetSourceHierarchyObjectsOK struct {
	Payload *models.SourceHierarchyObjectSummaries
}

// IsSuccess returns true when this get source hierarchy objects o k response has a 2xx status code
func (o *GetSourceHierarchyObjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get source hierarchy objects o k response has a 3xx status code
func (o *GetSourceHierarchyObjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source hierarchy objects o k response has a 4xx status code
func (o *GetSourceHierarchyObjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source hierarchy objects o k response has a 5xx status code
func (o *GetSourceHierarchyObjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get source hierarchy objects o k response a status code equal to that given
func (o *GetSourceHierarchyObjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get source hierarchy objects o k response
func (o *GetSourceHierarchyObjectsOK) Code() int {
	return 200
}

func (o *GetSourceHierarchyObjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/sources/{sourceId}/objects][%d] getSourceHierarchyObjectsOK %s", 200, payload)
}

func (o *GetSourceHierarchyObjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/sources/{sourceId}/objects][%d] getSourceHierarchyObjectsOK %s", 200, payload)
}

func (o *GetSourceHierarchyObjectsOK) GetPayload() *models.SourceHierarchyObjectSummaries {
	return o.Payload
}

func (o *GetSourceHierarchyObjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceHierarchyObjectSummaries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceHierarchyObjectsDefault creates a GetSourceHierarchyObjectsDefault with default headers values
func NewGetSourceHierarchyObjectsDefault(code int) *GetSourceHierarchyObjectsDefault {
	return &GetSourceHierarchyObjectsDefault{
		_statusCode: code,
	}
}

/*
GetSourceHierarchyObjectsDefault describes a response with status code -1, with default header values.

Error
*/
type GetSourceHierarchyObjectsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get source hierarchy objects default response has a 2xx status code
func (o *GetSourceHierarchyObjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get source hierarchy objects default response has a 3xx status code
func (o *GetSourceHierarchyObjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get source hierarchy objects default response has a 4xx status code
func (o *GetSourceHierarchyObjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get source hierarchy objects default response has a 5xx status code
func (o *GetSourceHierarchyObjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get source hierarchy objects default response a status code equal to that given
func (o *GetSourceHierarchyObjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get source hierarchy objects default response
func (o *GetSourceHierarchyObjectsDefault) Code() int {
	return o._statusCode
}

func (o *GetSourceHierarchyObjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/sources/{sourceId}/objects][%d] GetSourceHierarchyObjects default %s", o._statusCode, payload)
}

func (o *GetSourceHierarchyObjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/sources/{sourceId}/objects][%d] GetSourceHierarchyObjects default %s", o._statusCode, payload)
}

func (o *GetSourceHierarchyObjectsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSourceHierarchyObjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
