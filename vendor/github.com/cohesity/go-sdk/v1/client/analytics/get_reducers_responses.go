// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// GetReducersReader is a Reader for the GetReducers structure.
type GetReducersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReducersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReducersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetReducersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReducersOK creates a GetReducersOK with default headers values
func NewGetReducersOK() *GetReducersOK {
	return &GetReducersOK{}
}

/*
GetReducersOK describes a response with status code 200, with default header values.

Success
*/
type GetReducersOK struct {
	Payload *models.ReducersWrapper
}

// IsSuccess returns true when this get reducers o k response has a 2xx status code
func (o *GetReducersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get reducers o k response has a 3xx status code
func (o *GetReducersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reducers o k response has a 4xx status code
func (o *GetReducersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get reducers o k response has a 5xx status code
func (o *GetReducersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get reducers o k response a status code equal to that given
func (o *GetReducersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get reducers o k response
func (o *GetReducersOK) Code() int {
	return 200
}

func (o *GetReducersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/reducers][%d] getReducersOK %s", 200, payload)
}

func (o *GetReducersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/reducers][%d] getReducersOK %s", 200, payload)
}

func (o *GetReducersOK) GetPayload() *models.ReducersWrapper {
	return o.Payload
}

func (o *GetReducersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReducersWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReducersDefault creates a GetReducersDefault with default headers values
func NewGetReducersDefault(code int) *GetReducersDefault {
	return &GetReducersDefault{
		_statusCode: code,
	}
}

/*
GetReducersDefault describes a response with status code -1, with default header values.

Error
*/
type GetReducersDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get reducers default response has a 2xx status code
func (o *GetReducersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get reducers default response has a 3xx status code
func (o *GetReducersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get reducers default response has a 4xx status code
func (o *GetReducersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get reducers default response has a 5xx status code
func (o *GetReducersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get reducers default response a status code equal to that given
func (o *GetReducersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get reducers default response
func (o *GetReducersDefault) Code() int {
	return o._statusCode
}

func (o *GetReducersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/reducers][%d] GetReducers default %s", o._statusCode, payload)
}

func (o *GetReducersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/reducers][%d] GetReducers default %s", o._statusCode, payload)
}

func (o *GetReducersDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetReducersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
