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

// DownloadMROutputFilesReader is a Reader for the DownloadMROutputFiles structure.
type DownloadMROutputFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadMROutputFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadMROutputFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDownloadMROutputFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDownloadMROutputFilesOK creates a DownloadMROutputFilesOK with default headers values
func NewDownloadMROutputFilesOK() *DownloadMROutputFilesOK {
	return &DownloadMROutputFilesOK{}
}

/*
DownloadMROutputFilesOK describes a response with status code 200, with default header values.

Success
*/
type DownloadMROutputFilesOK struct {
	Payload *models.ExtractFileRangeResult
}

// IsSuccess returns true when this download m r output files o k response has a 2xx status code
func (o *DownloadMROutputFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download m r output files o k response has a 3xx status code
func (o *DownloadMROutputFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download m r output files o k response has a 4xx status code
func (o *DownloadMROutputFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download m r output files o k response has a 5xx status code
func (o *DownloadMROutputFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this download m r output files o k response a status code equal to that given
func (o *DownloadMROutputFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download m r output files o k response
func (o *DownloadMROutputFilesOK) Code() int {
	return 200
}

func (o *DownloadMROutputFilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/mrOutputfiles][%d] downloadMROutputFilesOK %s", 200, payload)
}

func (o *DownloadMROutputFilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/mrOutputfiles][%d] downloadMROutputFilesOK %s", 200, payload)
}

func (o *DownloadMROutputFilesOK) GetPayload() *models.ExtractFileRangeResult {
	return o.Payload
}

func (o *DownloadMROutputFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtractFileRangeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadMROutputFilesDefault creates a DownloadMROutputFilesDefault with default headers values
func NewDownloadMROutputFilesDefault(code int) *DownloadMROutputFilesDefault {
	return &DownloadMROutputFilesDefault{
		_statusCode: code,
	}
}

/*
DownloadMROutputFilesDefault describes a response with status code -1, with default header values.

Error
*/
type DownloadMROutputFilesDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this download m r output files default response has a 2xx status code
func (o *DownloadMROutputFilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this download m r output files default response has a 3xx status code
func (o *DownloadMROutputFilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this download m r output files default response has a 4xx status code
func (o *DownloadMROutputFilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this download m r output files default response has a 5xx status code
func (o *DownloadMROutputFilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this download m r output files default response a status code equal to that given
func (o *DownloadMROutputFilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the download m r output files default response
func (o *DownloadMROutputFilesDefault) Code() int {
	return o._statusCode
}

func (o *DownloadMROutputFilesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/mrOutputfiles][%d] DownloadMROutputFiles default %s", o._statusCode, payload)
}

func (o *DownloadMROutputFilesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/analytics/mrOutputfiles][%d] DownloadMROutputFiles default %s", o._statusCode, payload)
}

func (o *DownloadMROutputFilesDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *DownloadMROutputFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
