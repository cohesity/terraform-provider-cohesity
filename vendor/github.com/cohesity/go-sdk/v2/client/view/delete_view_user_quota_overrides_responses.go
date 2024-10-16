// Code generated by go-swagger; DO NOT EDIT.

package view

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

// DeleteViewUserQuotaOverridesReader is a Reader for the DeleteViewUserQuotaOverrides structure.
type DeleteViewUserQuotaOverridesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteViewUserQuotaOverridesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteViewUserQuotaOverridesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteViewUserQuotaOverridesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteViewUserQuotaOverridesNoContent creates a DeleteViewUserQuotaOverridesNoContent with default headers values
func NewDeleteViewUserQuotaOverridesNoContent() *DeleteViewUserQuotaOverridesNoContent {
	return &DeleteViewUserQuotaOverridesNoContent{}
}

/*
DeleteViewUserQuotaOverridesNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteViewUserQuotaOverridesNoContent struct {
}

// IsSuccess returns true when this delete view user quota overrides no content response has a 2xx status code
func (o *DeleteViewUserQuotaOverridesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete view user quota overrides no content response has a 3xx status code
func (o *DeleteViewUserQuotaOverridesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete view user quota overrides no content response has a 4xx status code
func (o *DeleteViewUserQuotaOverridesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete view user quota overrides no content response has a 5xx status code
func (o *DeleteViewUserQuotaOverridesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete view user quota overrides no content response a status code equal to that given
func (o *DeleteViewUserQuotaOverridesNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete view user quota overrides no content response
func (o *DeleteViewUserQuotaOverridesNoContent) Code() int {
	return 204
}

func (o *DeleteViewUserQuotaOverridesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /file-services/views/{viewId}/user-quotas][%d] deleteViewUserQuotaOverridesNoContent", 204)
}

func (o *DeleteViewUserQuotaOverridesNoContent) String() string {
	return fmt.Sprintf("[DELETE /file-services/views/{viewId}/user-quotas][%d] deleteViewUserQuotaOverridesNoContent", 204)
}

func (o *DeleteViewUserQuotaOverridesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteViewUserQuotaOverridesDefault creates a DeleteViewUserQuotaOverridesDefault with default headers values
func NewDeleteViewUserQuotaOverridesDefault(code int) *DeleteViewUserQuotaOverridesDefault {
	return &DeleteViewUserQuotaOverridesDefault{
		_statusCode: code,
	}
}

/*
DeleteViewUserQuotaOverridesDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteViewUserQuotaOverridesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete view user quota overrides default response has a 2xx status code
func (o *DeleteViewUserQuotaOverridesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete view user quota overrides default response has a 3xx status code
func (o *DeleteViewUserQuotaOverridesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete view user quota overrides default response has a 4xx status code
func (o *DeleteViewUserQuotaOverridesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete view user quota overrides default response has a 5xx status code
func (o *DeleteViewUserQuotaOverridesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete view user quota overrides default response a status code equal to that given
func (o *DeleteViewUserQuotaOverridesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete view user quota overrides default response
func (o *DeleteViewUserQuotaOverridesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteViewUserQuotaOverridesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /file-services/views/{viewId}/user-quotas][%d] DeleteViewUserQuotaOverrides default %s", o._statusCode, payload)
}

func (o *DeleteViewUserQuotaOverridesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /file-services/views/{viewId}/user-quotas][%d] DeleteViewUserQuotaOverrides default %s", o._statusCode, payload)
}

func (o *DeleteViewUserQuotaOverridesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteViewUserQuotaOverridesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
