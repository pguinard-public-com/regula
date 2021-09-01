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

// CreateEnvironmentReader is a Reader for the CreateEnvironment structure.
type CreateEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEnvironmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEnvironmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEnvironmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateEnvironmentCreated creates a CreateEnvironmentCreated with default headers values
func NewCreateEnvironmentCreated() *CreateEnvironmentCreated {
	return &CreateEnvironmentCreated{}
}

/*CreateEnvironmentCreated handles this case with default header values.

New environment details.
*/
type CreateEnvironmentCreated struct {
	Payload *models.Environment
}

func (o *CreateEnvironmentCreated) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentCreated  %+v", 201, o.Payload)
}

func (o *CreateEnvironmentCreated) GetPayload() *models.Environment {
	return o.Payload
}

func (o *CreateEnvironmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Environment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentBadRequest creates a CreateEnvironmentBadRequest with default headers values
func NewCreateEnvironmentBadRequest() *CreateEnvironmentBadRequest {
	return &CreateEnvironmentBadRequest{}
}

/*CreateEnvironmentBadRequest handles this case with default header values.

BadRequestError
*/
type CreateEnvironmentBadRequest struct {
	Payload *models.BadRequestError
}

func (o *CreateEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateEnvironmentBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *CreateEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentUnauthorized creates a CreateEnvironmentUnauthorized with default headers values
func NewCreateEnvironmentUnauthorized() *CreateEnvironmentUnauthorized {
	return &CreateEnvironmentUnauthorized{}
}

/*CreateEnvironmentUnauthorized handles this case with default header values.

AuthenticationError
*/
type CreateEnvironmentUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *CreateEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateEnvironmentUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *CreateEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentForbidden creates a CreateEnvironmentForbidden with default headers values
func NewCreateEnvironmentForbidden() *CreateEnvironmentForbidden {
	return &CreateEnvironmentForbidden{}
}

/*CreateEnvironmentForbidden handles this case with default header values.

AuthorizationError
*/
type CreateEnvironmentForbidden struct {
	Payload *models.AuthorizationError
}

func (o *CreateEnvironmentForbidden) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentForbidden  %+v", 403, o.Payload)
}

func (o *CreateEnvironmentForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *CreateEnvironmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentNotFound creates a CreateEnvironmentNotFound with default headers values
func NewCreateEnvironmentNotFound() *CreateEnvironmentNotFound {
	return &CreateEnvironmentNotFound{}
}

/*CreateEnvironmentNotFound handles this case with default header values.

NotFoundError
*/
type CreateEnvironmentNotFound struct {
	Payload *models.NotFoundError
}

func (o *CreateEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentNotFound  %+v", 404, o.Payload)
}

func (o *CreateEnvironmentNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *CreateEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentInternalServerError creates a CreateEnvironmentInternalServerError with default headers values
func NewCreateEnvironmentInternalServerError() *CreateEnvironmentInternalServerError {
	return &CreateEnvironmentInternalServerError{}
}

/*CreateEnvironmentInternalServerError handles this case with default header values.

InternalServerError
*/
type CreateEnvironmentInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *CreateEnvironmentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateEnvironmentInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *CreateEnvironmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
