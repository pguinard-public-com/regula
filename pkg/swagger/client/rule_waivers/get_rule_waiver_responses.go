// Code generated by go-swagger; DO NOT EDIT.

package rule_waivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/regula/v3/pkg/swagger/models"
)

// GetRuleWaiverReader is a Reader for the GetRuleWaiver structure.
type GetRuleWaiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleWaiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuleWaiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRuleWaiverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRuleWaiverForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRuleWaiverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRuleWaiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRuleWaiverOK creates a GetRuleWaiverOK with default headers values
func NewGetRuleWaiverOK() *GetRuleWaiverOK {
	return &GetRuleWaiverOK{}
}

/*GetRuleWaiverOK handles this case with default header values.

Rule waiver details.
*/
type GetRuleWaiverOK struct {
	Payload *models.RuleWaiver
}

func (o *GetRuleWaiverOK) Error() string {
	return fmt.Sprintf("[GET /rule_waivers/{rule_waiver_id}][%d] getRuleWaiverOK  %+v", 200, o.Payload)
}

func (o *GetRuleWaiverOK) GetPayload() *models.RuleWaiver {
	return o.Payload
}

func (o *GetRuleWaiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleWaiver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleWaiverUnauthorized creates a GetRuleWaiverUnauthorized with default headers values
func NewGetRuleWaiverUnauthorized() *GetRuleWaiverUnauthorized {
	return &GetRuleWaiverUnauthorized{}
}

/*GetRuleWaiverUnauthorized handles this case with default header values.

AuthenticationError
*/
type GetRuleWaiverUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *GetRuleWaiverUnauthorized) Error() string {
	return fmt.Sprintf("[GET /rule_waivers/{rule_waiver_id}][%d] getRuleWaiverUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRuleWaiverUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *GetRuleWaiverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleWaiverForbidden creates a GetRuleWaiverForbidden with default headers values
func NewGetRuleWaiverForbidden() *GetRuleWaiverForbidden {
	return &GetRuleWaiverForbidden{}
}

/*GetRuleWaiverForbidden handles this case with default header values.

AuthorizationError
*/
type GetRuleWaiverForbidden struct {
	Payload *models.AuthorizationError
}

func (o *GetRuleWaiverForbidden) Error() string {
	return fmt.Sprintf("[GET /rule_waivers/{rule_waiver_id}][%d] getRuleWaiverForbidden  %+v", 403, o.Payload)
}

func (o *GetRuleWaiverForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *GetRuleWaiverForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleWaiverNotFound creates a GetRuleWaiverNotFound with default headers values
func NewGetRuleWaiverNotFound() *GetRuleWaiverNotFound {
	return &GetRuleWaiverNotFound{}
}

/*GetRuleWaiverNotFound handles this case with default header values.

NotFoundError
*/
type GetRuleWaiverNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetRuleWaiverNotFound) Error() string {
	return fmt.Sprintf("[GET /rule_waivers/{rule_waiver_id}][%d] getRuleWaiverNotFound  %+v", 404, o.Payload)
}

func (o *GetRuleWaiverNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetRuleWaiverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleWaiverInternalServerError creates a GetRuleWaiverInternalServerError with default headers values
func NewGetRuleWaiverInternalServerError() *GetRuleWaiverInternalServerError {
	return &GetRuleWaiverInternalServerError{}
}

/*GetRuleWaiverInternalServerError handles this case with default header values.

InternalServerError
*/
type GetRuleWaiverInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetRuleWaiverInternalServerError) Error() string {
	return fmt.Sprintf("[GET /rule_waivers/{rule_waiver_id}][%d] getRuleWaiverInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRuleWaiverInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetRuleWaiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
