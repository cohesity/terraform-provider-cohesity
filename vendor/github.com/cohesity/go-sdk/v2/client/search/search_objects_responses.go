// Code generated by go-swagger; DO NOT EDIT.

package search

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

// SearchObjectsReader is a Reader for the SearchObjects structure.
type SearchObjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchObjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchObjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchObjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchObjectsOK creates a SearchObjectsOK with default headers values
func NewSearchObjectsOK() *SearchObjectsOK {
	return &SearchObjectsOK{}
}

/*
SearchObjectsOK describes a response with status code 200, with default header values.

Success
*/
type SearchObjectsOK struct {
	Payload *models.ObjectsSearchResponseBody
}

// IsSuccess returns true when this search objects o k response has a 2xx status code
func (o *SearchObjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search objects o k response has a 3xx status code
func (o *SearchObjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objects o k response has a 4xx status code
func (o *SearchObjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search objects o k response has a 5xx status code
func (o *SearchObjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search objects o k response a status code equal to that given
func (o *SearchObjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search objects o k response
func (o *SearchObjectsOK) Code() int {
	return 200
}

func (o *SearchObjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/search/objects][%d] searchObjectsOK %s", 200, payload)
}

func (o *SearchObjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/search/objects][%d] searchObjectsOK %s", 200, payload)
}

func (o *SearchObjectsOK) GetPayload() *models.ObjectsSearchResponseBody {
	return o.Payload
}

func (o *SearchObjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectsSearchResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchObjectsDefault creates a SearchObjectsDefault with default headers values
func NewSearchObjectsDefault(code int) *SearchObjectsDefault {
	return &SearchObjectsDefault{
		_statusCode: code,
	}
}

/*
SearchObjectsDefault describes a response with status code -1, with default header values.

Error
*/
type SearchObjectsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this search objects default response has a 2xx status code
func (o *SearchObjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search objects default response has a 3xx status code
func (o *SearchObjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search objects default response has a 4xx status code
func (o *SearchObjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search objects default response has a 5xx status code
func (o *SearchObjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search objects default response a status code equal to that given
func (o *SearchObjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search objects default response
func (o *SearchObjectsDefault) Code() int {
	return o._statusCode
}

func (o *SearchObjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/search/objects][%d] SearchObjects default %s", o._statusCode, payload)
}

func (o *SearchObjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/search/objects][%d] SearchObjects default %s", o._statusCode, payload)
}

func (o *SearchObjectsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchObjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
