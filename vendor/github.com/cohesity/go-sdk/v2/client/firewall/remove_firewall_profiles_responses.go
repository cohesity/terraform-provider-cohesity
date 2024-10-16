// Code generated by go-swagger; DO NOT EDIT.

package firewall

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

// RemoveFirewallProfilesReader is a Reader for the RemoveFirewallProfiles structure.
type RemoveFirewallProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveFirewallProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveFirewallProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveFirewallProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveFirewallProfilesOK creates a RemoveFirewallProfilesOK with default headers values
func NewRemoveFirewallProfilesOK() *RemoveFirewallProfilesOK {
	return &RemoveFirewallProfilesOK{}
}

/*
RemoveFirewallProfilesOK describes a response with status code 200, with default header values.

Success
*/
type RemoveFirewallProfilesOK struct {
	Payload *models.SuccessResp
}

// IsSuccess returns true when this remove firewall profiles o k response has a 2xx status code
func (o *RemoveFirewallProfilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove firewall profiles o k response has a 3xx status code
func (o *RemoveFirewallProfilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove firewall profiles o k response has a 4xx status code
func (o *RemoveFirewallProfilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove firewall profiles o k response has a 5xx status code
func (o *RemoveFirewallProfilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove firewall profiles o k response a status code equal to that given
func (o *RemoveFirewallProfilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove firewall profiles o k response
func (o *RemoveFirewallProfilesOK) Code() int {
	return 200
}

func (o *RemoveFirewallProfilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/firewall/profile/remove][%d] removeFirewallProfilesOK %s", 200, payload)
}

func (o *RemoveFirewallProfilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/firewall/profile/remove][%d] removeFirewallProfilesOK %s", 200, payload)
}

func (o *RemoveFirewallProfilesOK) GetPayload() *models.SuccessResp {
	return o.Payload
}

func (o *RemoveFirewallProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveFirewallProfilesDefault creates a RemoveFirewallProfilesDefault with default headers values
func NewRemoveFirewallProfilesDefault(code int) *RemoveFirewallProfilesDefault {
	return &RemoveFirewallProfilesDefault{
		_statusCode: code,
	}
}

/*
RemoveFirewallProfilesDefault describes a response with status code -1, with default header values.

Error
*/
type RemoveFirewallProfilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this remove firewall profiles default response has a 2xx status code
func (o *RemoveFirewallProfilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove firewall profiles default response has a 3xx status code
func (o *RemoveFirewallProfilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove firewall profiles default response has a 4xx status code
func (o *RemoveFirewallProfilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove firewall profiles default response has a 5xx status code
func (o *RemoveFirewallProfilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove firewall profiles default response a status code equal to that given
func (o *RemoveFirewallProfilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the remove firewall profiles default response
func (o *RemoveFirewallProfilesDefault) Code() int {
	return o._statusCode
}

func (o *RemoveFirewallProfilesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/firewall/profile/remove][%d] RemoveFirewallProfiles default %s", o._statusCode, payload)
}

func (o *RemoveFirewallProfilesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/firewall/profile/remove][%d] RemoveFirewallProfiles default %s", o._statusCode, payload)
}

func (o *RemoveFirewallProfilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveFirewallProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
