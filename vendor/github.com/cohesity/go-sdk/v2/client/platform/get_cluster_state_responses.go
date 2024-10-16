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

// GetClusterStateReader is a Reader for the GetClusterState structure.
type GetClusterStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterStateOK creates a GetClusterStateOK with default headers values
func NewGetClusterStateOK() *GetClusterStateOK {
	return &GetClusterStateOK{}
}

/*
GetClusterStateOK describes a response with status code 200, with default header values.

Success
*/
type GetClusterStateOK struct {
	Payload *models.ClusterStateParams
}

// IsSuccess returns true when this get cluster state o k response has a 2xx status code
func (o *GetClusterStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster state o k response has a 3xx status code
func (o *GetClusterStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster state o k response has a 4xx status code
func (o *GetClusterStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster state o k response has a 5xx status code
func (o *GetClusterStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster state o k response a status code equal to that given
func (o *GetClusterStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster state o k response
func (o *GetClusterStateOK) Code() int {
	return 200
}

func (o *GetClusterStateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/state][%d] getClusterStateOK %s", 200, payload)
}

func (o *GetClusterStateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/state][%d] getClusterStateOK %s", 200, payload)
}

func (o *GetClusterStateOK) GetPayload() *models.ClusterStateParams {
	return o.Payload
}

func (o *GetClusterStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterStateParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterStateDefault creates a GetClusterStateDefault with default headers values
func NewGetClusterStateDefault(code int) *GetClusterStateDefault {
	return &GetClusterStateDefault{
		_statusCode: code,
	}
}

/*
GetClusterStateDefault describes a response with status code -1, with default header values.

Error
*/
type GetClusterStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get cluster state default response has a 2xx status code
func (o *GetClusterStateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster state default response has a 3xx status code
func (o *GetClusterStateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster state default response has a 4xx status code
func (o *GetClusterStateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster state default response has a 5xx status code
func (o *GetClusterStateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster state default response a status code equal to that given
func (o *GetClusterStateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cluster state default response
func (o *GetClusterStateDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterStateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/state][%d] GetClusterState default %s", o._statusCode, payload)
}

func (o *GetClusterStateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/state][%d] GetClusterState default %s", o._statusCode, payload)
}

func (o *GetClusterStateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
