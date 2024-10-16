// Code generated by go-swagger; DO NOT EDIT.

package network

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

// DeleteIPConfigReader is a Reader for the DeleteIPConfig structure.
type DeleteIPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIPConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteIPConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteIPConfigNoContent creates a DeleteIPConfigNoContent with default headers values
func NewDeleteIPConfigNoContent() *DeleteIPConfigNoContent {
	return &DeleteIPConfigNoContent{}
}

/*
DeleteIPConfigNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteIPConfigNoContent struct {
}

// IsSuccess returns true when this delete Ip config no content response has a 2xx status code
func (o *DeleteIPConfigNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Ip config no content response has a 3xx status code
func (o *DeleteIPConfigNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Ip config no content response has a 4xx status code
func (o *DeleteIPConfigNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Ip config no content response has a 5xx status code
func (o *DeleteIPConfigNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Ip config no content response a status code equal to that given
func (o *DeleteIPConfigNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Ip config no content response
func (o *DeleteIPConfigNoContent) Code() int {
	return 204
}

func (o *DeleteIPConfigNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/network/ipConfig][%d] deleteIpConfigNoContent", 204)
}

func (o *DeleteIPConfigNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/network/ipConfig][%d] deleteIpConfigNoContent", 204)
}

func (o *DeleteIPConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPConfigDefault creates a DeleteIPConfigDefault with default headers values
func NewDeleteIPConfigDefault(code int) *DeleteIPConfigDefault {
	return &DeleteIPConfigDefault{
		_statusCode: code,
	}
}

/*
DeleteIPConfigDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteIPConfigDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this delete Ip config default response has a 2xx status code
func (o *DeleteIPConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete Ip config default response has a 3xx status code
func (o *DeleteIPConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete Ip config default response has a 4xx status code
func (o *DeleteIPConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete Ip config default response has a 5xx status code
func (o *DeleteIPConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete Ip config default response a status code equal to that given
func (o *DeleteIPConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete Ip config default response
func (o *DeleteIPConfigDefault) Code() int {
	return o._statusCode
}

func (o *DeleteIPConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/network/ipConfig][%d] DeleteIpConfig default %s", o._statusCode, payload)
}

func (o *DeleteIPConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/network/ipConfig][%d] DeleteIpConfig default %s", o._statusCode, payload)
}

func (o *DeleteIPConfigDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *DeleteIPConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
