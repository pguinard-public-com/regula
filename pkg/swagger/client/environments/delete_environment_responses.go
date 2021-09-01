// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/regula/pkg/swagger/models"
)

// DeleteEnvironmentReader is a Reader for the DeleteEnvironment structure.
type DeleteEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEnvironmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteEnvironmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteEnvironmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEnvironmentNoContent creates a DeleteEnvironmentNoContent with default headers values
func NewDeleteEnvironmentNoContent() *DeleteEnvironmentNoContent {
	return &DeleteEnvironmentNoContent{}
}

/*DeleteEnvironmentNoContent handles this case with default header values.

Environment deleted.
*/
type DeleteEnvironmentNoContent struct {
}

func (o *DeleteEnvironmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /environments/{environment_id}][%d] deleteEnvironmentNoContent ", 204)
}

func (o *DeleteEnvironmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvironmentBadRequest creates a DeleteEnvironmentBadRequest with default headers values
func NewDeleteEnvironmentBadRequest() *DeleteEnvironmentBadRequest {
	return &DeleteEnvironmentBadRequest{}
}

/*DeleteEnvironmentBadRequest handles this case with default header values.

BadRequestError
*/
type DeleteEnvironmentBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /environments/{environment_id}][%d] deleteEnvironmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteEnvironmentBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnvironmentUnauthorized creates a DeleteEnvironmentUnauthorized with default headers values
func NewDeleteEnvironmentUnauthorized() *DeleteEnvironmentUnauthorized {
	return &DeleteEnvironmentUnauthorized{}
}

/*DeleteEnvironmentUnauthorized handles this case with default header values.

AuthenticationError
*/
type DeleteEnvironmentUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *DeleteEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /environments/{environment_id}][%d] deleteEnvironmentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteEnvironmentUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *DeleteEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnvironmentForbidden creates a DeleteEnvironmentForbidden with default headers values
func NewDeleteEnvironmentForbidden() *DeleteEnvironmentForbidden {
	return &DeleteEnvironmentForbidden{}
}

/*DeleteEnvironmentForbidden handles this case with default header values.

AuthorizationError
*/
type DeleteEnvironmentForbidden struct {
	Payload *models.AuthorizationError
}

func (o *DeleteEnvironmentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /environments/{environment_id}][%d] deleteEnvironmentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteEnvironmentForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *DeleteEnvironmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnvironmentNotFound creates a DeleteEnvironmentNotFound with default headers values
func NewDeleteEnvironmentNotFound() *DeleteEnvironmentNotFound {
	return &DeleteEnvironmentNotFound{}
}

/*DeleteEnvironmentNotFound handles this case with default header values.

NotFoundError
*/
type DeleteEnvironmentNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /environments/{environment_id}][%d] deleteEnvironmentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteEnvironmentNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnvironmentInternalServerError creates a DeleteEnvironmentInternalServerError with default headers values
func NewDeleteEnvironmentInternalServerError() *DeleteEnvironmentInternalServerError {
	return &DeleteEnvironmentInternalServerError{}
}

/*DeleteEnvironmentInternalServerError handles this case with default header values.

InternalServerError
*/
type DeleteEnvironmentInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteEnvironmentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /environments/{environment_id}][%d] deleteEnvironmentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteEnvironmentInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteEnvironmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
