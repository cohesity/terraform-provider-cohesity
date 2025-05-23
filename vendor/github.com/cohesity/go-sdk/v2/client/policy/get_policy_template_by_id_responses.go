// Code generated by go-swagger; DO NOT EDIT.

package policy

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

// GetPolicyTemplateByIDReader is a Reader for the GetPolicyTemplateByID structure.
type GetPolicyTemplateByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyTemplateByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyTemplateByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPolicyTemplateByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPolicyTemplateByIDOK creates a GetPolicyTemplateByIDOK with default headers values
func NewGetPolicyTemplateByIDOK() *GetPolicyTemplateByIDOK {
	return &GetPolicyTemplateByIDOK{}
}

/*
GetPolicyTemplateByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetPolicyTemplateByIDOK struct {
	Payload *models.PolicyTemplateResponse
}

// IsSuccess returns true when this get policy template by Id o k response has a 2xx status code
func (o *GetPolicyTemplateByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policy template by Id o k response has a 3xx status code
func (o *GetPolicyTemplateByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy template by Id o k response has a 4xx status code
func (o *GetPolicyTemplateByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policy template by Id o k response has a 5xx status code
func (o *GetPolicyTemplateByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy template by Id o k response a status code equal to that given
func (o *GetPolicyTemplateByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get policy template by Id o k response
func (o *GetPolicyTemplateByIDOK) Code() int {
	return 200
}

func (o *GetPolicyTemplateByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/policy-templates/{id}][%d] getPolicyTemplateByIdOK %s", 200, payload)
}

func (o *GetPolicyTemplateByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/policy-templates/{id}][%d] getPolicyTemplateByIdOK %s", 200, payload)
}

func (o *GetPolicyTemplateByIDOK) GetPayload() *models.PolicyTemplateResponse {
	return o.Payload
}

func (o *GetPolicyTemplateByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyTemplateByIDDefault creates a GetPolicyTemplateByIDDefault with default headers values
func NewGetPolicyTemplateByIDDefault(code int) *GetPolicyTemplateByIDDefault {
	return &GetPolicyTemplateByIDDefault{
		_statusCode: code,
	}
}

/*
GetPolicyTemplateByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetPolicyTemplateByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get policy template by Id default response has a 2xx status code
func (o *GetPolicyTemplateByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get policy template by Id default response has a 3xx status code
func (o *GetPolicyTemplateByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get policy template by Id default response has a 4xx status code
func (o *GetPolicyTemplateByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get policy template by Id default response has a 5xx status code
func (o *GetPolicyTemplateByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get policy template by Id default response a status code equal to that given
func (o *GetPolicyTemplateByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get policy template by Id default response
func (o *GetPolicyTemplateByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPolicyTemplateByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/policy-templates/{id}][%d] GetPolicyTemplateById default %s", o._statusCode, payload)
}

func (o *GetPolicyTemplateByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/policy-templates/{id}][%d] GetPolicyTemplateById default %s", o._statusCode, payload)
}

func (o *GetPolicyTemplateByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPolicyTemplateByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
