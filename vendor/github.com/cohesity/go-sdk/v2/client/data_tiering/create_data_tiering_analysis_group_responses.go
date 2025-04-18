// Code generated by go-swagger; DO NOT EDIT.

package data_tiering

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

// CreateDataTieringAnalysisGroupReader is a Reader for the CreateDataTieringAnalysisGroup structure.
type CreateDataTieringAnalysisGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDataTieringAnalysisGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDataTieringAnalysisGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDataTieringAnalysisGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDataTieringAnalysisGroupCreated creates a CreateDataTieringAnalysisGroupCreated with default headers values
func NewCreateDataTieringAnalysisGroupCreated() *CreateDataTieringAnalysisGroupCreated {
	return &CreateDataTieringAnalysisGroupCreated{}
}

/*
CreateDataTieringAnalysisGroupCreated describes a response with status code 201, with default header values.

Success
*/
type CreateDataTieringAnalysisGroupCreated struct {
	Payload *models.DataTieringAnalysisGroup
}

// IsSuccess returns true when this create data tiering analysis group created response has a 2xx status code
func (o *CreateDataTieringAnalysisGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create data tiering analysis group created response has a 3xx status code
func (o *CreateDataTieringAnalysisGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create data tiering analysis group created response has a 4xx status code
func (o *CreateDataTieringAnalysisGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create data tiering analysis group created response has a 5xx status code
func (o *CreateDataTieringAnalysisGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create data tiering analysis group created response a status code equal to that given
func (o *CreateDataTieringAnalysisGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create data tiering analysis group created response
func (o *CreateDataTieringAnalysisGroupCreated) Code() int {
	return 201
}

func (o *CreateDataTieringAnalysisGroupCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/analysis-groups][%d] createDataTieringAnalysisGroupCreated %s", 201, payload)
}

func (o *CreateDataTieringAnalysisGroupCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/analysis-groups][%d] createDataTieringAnalysisGroupCreated %s", 201, payload)
}

func (o *CreateDataTieringAnalysisGroupCreated) GetPayload() *models.DataTieringAnalysisGroup {
	return o.Payload
}

func (o *CreateDataTieringAnalysisGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTieringAnalysisGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataTieringAnalysisGroupDefault creates a CreateDataTieringAnalysisGroupDefault with default headers values
func NewCreateDataTieringAnalysisGroupDefault(code int) *CreateDataTieringAnalysisGroupDefault {
	return &CreateDataTieringAnalysisGroupDefault{
		_statusCode: code,
	}
}

/*
CreateDataTieringAnalysisGroupDefault describes a response with status code -1, with default header values.

Error
*/
type CreateDataTieringAnalysisGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create data tiering analysis group default response has a 2xx status code
func (o *CreateDataTieringAnalysisGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create data tiering analysis group default response has a 3xx status code
func (o *CreateDataTieringAnalysisGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create data tiering analysis group default response has a 4xx status code
func (o *CreateDataTieringAnalysisGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create data tiering analysis group default response has a 5xx status code
func (o *CreateDataTieringAnalysisGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create data tiering analysis group default response a status code equal to that given
func (o *CreateDataTieringAnalysisGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create data tiering analysis group default response
func (o *CreateDataTieringAnalysisGroupDefault) Code() int {
	return o._statusCode
}

func (o *CreateDataTieringAnalysisGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/analysis-groups][%d] CreateDataTieringAnalysisGroup default %s", o._statusCode, payload)
}

func (o *CreateDataTieringAnalysisGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-tiering/analysis-groups][%d] CreateDataTieringAnalysisGroup default %s", o._statusCode, payload)
}

func (o *CreateDataTieringAnalysisGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDataTieringAnalysisGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
