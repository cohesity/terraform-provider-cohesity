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

// GetVMVolumesInformationReader is a Reader for the GetVMVolumesInformation structure.
type GetVMVolumesInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMVolumesInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMVolumesInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVMVolumesInformationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVMVolumesInformationOK creates a GetVMVolumesInformationOK with default headers values
func NewGetVMVolumesInformationOK() *GetVMVolumesInformationOK {
	return &GetVMVolumesInformationOK{}
}

/*
GetVMVolumesInformationOK describes a response with status code 200, with default header values.

Success
*/
type GetVMVolumesInformationOK struct {
	Payload *models.VMVolumesInformation
}

// IsSuccess returns true when this get Vm volumes information o k response has a 2xx status code
func (o *GetVMVolumesInformationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Vm volumes information o k response has a 3xx status code
func (o *GetVMVolumesInformationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vm volumes information o k response has a 4xx status code
func (o *GetVMVolumesInformationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Vm volumes information o k response has a 5xx status code
func (o *GetVMVolumesInformationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vm volumes information o k response a status code equal to that given
func (o *GetVMVolumesInformationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Vm volumes information o k response
func (o *GetVMVolumesInformationOK) Code() int {
	return 200
}

func (o *GetVMVolumesInformationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/vms/volumesInformation][%d] getVmVolumesInformationOK %s", 200, payload)
}

func (o *GetVMVolumesInformationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/vms/volumesInformation][%d] getVmVolumesInformationOK %s", 200, payload)
}

func (o *GetVMVolumesInformationOK) GetPayload() *models.VMVolumesInformation {
	return o.Payload
}

func (o *GetVMVolumesInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMVolumesInformation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMVolumesInformationDefault creates a GetVMVolumesInformationDefault with default headers values
func NewGetVMVolumesInformationDefault(code int) *GetVMVolumesInformationDefault {
	return &GetVMVolumesInformationDefault{
		_statusCode: code,
	}
}

/*
GetVMVolumesInformationDefault describes a response with status code -1, with default header values.

Error
*/
type GetVMVolumesInformationDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get Vm volumes information default response has a 2xx status code
func (o *GetVMVolumesInformationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get Vm volumes information default response has a 3xx status code
func (o *GetVMVolumesInformationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get Vm volumes information default response has a 4xx status code
func (o *GetVMVolumesInformationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get Vm volumes information default response has a 5xx status code
func (o *GetVMVolumesInformationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get Vm volumes information default response a status code equal to that given
func (o *GetVMVolumesInformationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get Vm volumes information default response
func (o *GetVMVolumesInformationDefault) Code() int {
	return o._statusCode
}

func (o *GetVMVolumesInformationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/vms/volumesInformation][%d] GetVmVolumesInformation default %s", o._statusCode, payload)
}

func (o *GetVMVolumesInformationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/vms/volumesInformation][%d] GetVmVolumesInformation default %s", o._statusCode, payload)
}

func (o *GetVMVolumesInformationDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetVMVolumesInformationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
