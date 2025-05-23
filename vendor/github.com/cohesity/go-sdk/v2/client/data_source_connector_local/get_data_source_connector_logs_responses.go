// Code generated by go-swagger; DO NOT EDIT.

package data_source_connector_local

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

// GetDataSourceConnectorLogsReader is a Reader for the GetDataSourceConnectorLogs structure.
type GetDataSourceConnectorLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataSourceConnectorLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataSourceConnectorLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDataSourceConnectorLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDataSourceConnectorLogsOK creates a GetDataSourceConnectorLogsOK with default headers values
func NewGetDataSourceConnectorLogsOK() *GetDataSourceConnectorLogsOK {
	return &GetDataSourceConnectorLogsOK{}
}

/*
GetDataSourceConnectorLogsOK describes a response with status code 200, with default header values.

Success
*/
type GetDataSourceConnectorLogsOK struct {
	Payload *models.DataSourceConnectorLogs
}

// IsSuccess returns true when this get data source connector logs o k response has a 2xx status code
func (o *GetDataSourceConnectorLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get data source connector logs o k response has a 3xx status code
func (o *GetDataSourceConnectorLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data source connector logs o k response has a 4xx status code
func (o *GetDataSourceConnectorLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data source connector logs o k response has a 5xx status code
func (o *GetDataSourceConnectorLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get data source connector logs o k response a status code equal to that given
func (o *GetDataSourceConnectorLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get data source connector logs o k response
func (o *GetDataSourceConnectorLogsOK) Code() int {
	return 200
}

func (o *GetDataSourceConnectorLogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-source-connector/logs][%d] getDataSourceConnectorLogsOK %s", 200, payload)
}

func (o *GetDataSourceConnectorLogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-source-connector/logs][%d] getDataSourceConnectorLogsOK %s", 200, payload)
}

func (o *GetDataSourceConnectorLogsOK) GetPayload() *models.DataSourceConnectorLogs {
	return o.Payload
}

func (o *GetDataSourceConnectorLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSourceConnectorLogs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourceConnectorLogsDefault creates a GetDataSourceConnectorLogsDefault with default headers values
func NewGetDataSourceConnectorLogsDefault(code int) *GetDataSourceConnectorLogsDefault {
	return &GetDataSourceConnectorLogsDefault{
		_statusCode: code,
	}
}

/*
GetDataSourceConnectorLogsDefault describes a response with status code -1, with default header values.

Error
*/
type GetDataSourceConnectorLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get data source connector logs default response has a 2xx status code
func (o *GetDataSourceConnectorLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get data source connector logs default response has a 3xx status code
func (o *GetDataSourceConnectorLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get data source connector logs default response has a 4xx status code
func (o *GetDataSourceConnectorLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get data source connector logs default response has a 5xx status code
func (o *GetDataSourceConnectorLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get data source connector logs default response a status code equal to that given
func (o *GetDataSourceConnectorLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get data source connector logs default response
func (o *GetDataSourceConnectorLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetDataSourceConnectorLogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-source-connector/logs][%d] GetDataSourceConnectorLogs default %s", o._statusCode, payload)
}

func (o *GetDataSourceConnectorLogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-source-connector/logs][%d] GetDataSourceConnectorLogs default %s", o._statusCode, payload)
}

func (o *GetDataSourceConnectorLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDataSourceConnectorLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
