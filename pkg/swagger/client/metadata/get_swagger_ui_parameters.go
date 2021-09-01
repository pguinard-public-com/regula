// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetSwaggerUIParams creates a new GetSwaggerUIParams object
// with the default values initialized.
func NewGetSwaggerUIParams() *GetSwaggerUIParams {

	return &GetSwaggerUIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSwaggerUIParamsWithTimeout creates a new GetSwaggerUIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSwaggerUIParamsWithTimeout(timeout time.Duration) *GetSwaggerUIParams {

	return &GetSwaggerUIParams{

		timeout: timeout,
	}
}

// NewGetSwaggerUIParamsWithContext creates a new GetSwaggerUIParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSwaggerUIParamsWithContext(ctx context.Context) *GetSwaggerUIParams {

	return &GetSwaggerUIParams{

		Context: ctx,
	}
}

// NewGetSwaggerUIParamsWithHTTPClient creates a new GetSwaggerUIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSwaggerUIParamsWithHTTPClient(client *http.Client) *GetSwaggerUIParams {

	return &GetSwaggerUIParams{
		HTTPClient: client,
	}
}

/*GetSwaggerUIParams contains all the parameters to send to the API endpoint
for the get swagger UI operation typically these are written to a http.Request
*/
type GetSwaggerUIParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get swagger UI params
func (o *GetSwaggerUIParams) WithTimeout(timeout time.Duration) *GetSwaggerUIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get swagger UI params
func (o *GetSwaggerUIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get swagger UI params
func (o *GetSwaggerUIParams) WithContext(ctx context.Context) *GetSwaggerUIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get swagger UI params
func (o *GetSwaggerUIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get swagger UI params
func (o *GetSwaggerUIParams) WithHTTPClient(client *http.Client) *GetSwaggerUIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get swagger UI params
func (o *GetSwaggerUIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSwaggerUIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
