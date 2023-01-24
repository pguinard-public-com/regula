// Code generated by go-swagger; DO NOT EDIT.

package rule_bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/regula/v3/pkg/swagger/models"
)

// GetLatestRuleBundleReader is a Reader for the GetLatestRuleBundle structure.
type GetLatestRuleBundleReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestRuleBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestRuleBundleOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetLatestRuleBundleNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetLatestRuleBundleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLatestRuleBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestRuleBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLatestRuleBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLatestRuleBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLatestRuleBundleOK creates a GetLatestRuleBundleOK with default headers values
func NewGetLatestRuleBundleOK(writer io.Writer) *GetLatestRuleBundleOK {
	return &GetLatestRuleBundleOK{
		Payload: writer,
	}
}

/*GetLatestRuleBundleOK handles this case with default header values.

Rule bundle.
*/
type GetLatestRuleBundleOK struct {
	Payload io.Writer
}

func (o *GetLatestRuleBundleOK) Error() string {
	return fmt.Sprintf("[GET /rule_bundles/latest][%d] getLatestRuleBundleOK  %+v", 200, o.Payload)
}

func (o *GetLatestRuleBundleOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetLatestRuleBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestRuleBundleNotModified creates a GetLatestRuleBundleNotModified with default headers values
func NewGetLatestRuleBundleNotModified() *GetLatestRuleBundleNotModified {
	return &GetLatestRuleBundleNotModified{}
}

/*GetLatestRuleBundleNotModified handles this case with default header values.

NotModified
*/
type GetLatestRuleBundleNotModified struct {
	Payload models.NotModifiedResponse
}

func (o *GetLatestRuleBundleNotModified) Error() string {
	return fmt.Sprintf("[GET /rule_bundles/latest][%d] getLatestRuleBundleNotModified  %+v", 304, o.Payload)
}

func (o *GetLatestRuleBundleNotModified) GetPayload() models.NotModifiedResponse {
	return o.Payload
}

func (o *GetLatestRuleBundleNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestRuleBundleBadRequest creates a GetLatestRuleBundleBadRequest with default headers values
func NewGetLatestRuleBundleBadRequest() *GetLatestRuleBundleBadRequest {
	return &GetLatestRuleBundleBadRequest{}
}

/*GetLatestRuleBundleBadRequest handles this case with default header values.

BadRequestError
*/
type GetLatestRuleBundleBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetLatestRuleBundleBadRequest) Error() string {
	return fmt.Sprintf("[GET /rule_bundles/latest][%d] getLatestRuleBundleBadRequest  %+v", 400, o.Payload)
}

func (o *GetLatestRuleBundleBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetLatestRuleBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestRuleBundleUnauthorized creates a GetLatestRuleBundleUnauthorized with default headers values
func NewGetLatestRuleBundleUnauthorized() *GetLatestRuleBundleUnauthorized {
	return &GetLatestRuleBundleUnauthorized{}
}

/*GetLatestRuleBundleUnauthorized handles this case with default header values.

AuthenticationError
*/
type GetLatestRuleBundleUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *GetLatestRuleBundleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /rule_bundles/latest][%d] getLatestRuleBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestRuleBundleUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *GetLatestRuleBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestRuleBundleForbidden creates a GetLatestRuleBundleForbidden with default headers values
func NewGetLatestRuleBundleForbidden() *GetLatestRuleBundleForbidden {
	return &GetLatestRuleBundleForbidden{}
}

/*GetLatestRuleBundleForbidden handles this case with default header values.

AuthorizationError
*/
type GetLatestRuleBundleForbidden struct {
	Payload *models.AuthorizationError
}

func (o *GetLatestRuleBundleForbidden) Error() string {
	return fmt.Sprintf("[GET /rule_bundles/latest][%d] getLatestRuleBundleForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestRuleBundleForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *GetLatestRuleBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestRuleBundleNotFound creates a GetLatestRuleBundleNotFound with default headers values
func NewGetLatestRuleBundleNotFound() *GetLatestRuleBundleNotFound {
	return &GetLatestRuleBundleNotFound{}
}

/*GetLatestRuleBundleNotFound handles this case with default header values.

NotFoundError
*/
type GetLatestRuleBundleNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetLatestRuleBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /rule_bundles/latest][%d] getLatestRuleBundleNotFound  %+v", 404, o.Payload)
}

func (o *GetLatestRuleBundleNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetLatestRuleBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestRuleBundleInternalServerError creates a GetLatestRuleBundleInternalServerError with default headers values
func NewGetLatestRuleBundleInternalServerError() *GetLatestRuleBundleInternalServerError {
	return &GetLatestRuleBundleInternalServerError{}
}

/*GetLatestRuleBundleInternalServerError handles this case with default header values.

InternalServerError
*/
type GetLatestRuleBundleInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetLatestRuleBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /rule_bundles/latest][%d] getLatestRuleBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestRuleBundleInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetLatestRuleBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
