// Code generated by go-swagger; DO NOT EDIT.

package families

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

// NewDeleteFamilyParams creates a new DeleteFamilyParams object
// with the default values initialized.
func NewDeleteFamilyParams() *DeleteFamilyParams {
	var ()
	return &DeleteFamilyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFamilyParamsWithTimeout creates a new DeleteFamilyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFamilyParamsWithTimeout(timeout time.Duration) *DeleteFamilyParams {
	var ()
	return &DeleteFamilyParams{

		timeout: timeout,
	}
}

// NewDeleteFamilyParamsWithContext creates a new DeleteFamilyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFamilyParamsWithContext(ctx context.Context) *DeleteFamilyParams {
	var ()
	return &DeleteFamilyParams{

		Context: ctx,
	}
}

// NewDeleteFamilyParamsWithHTTPClient creates a new DeleteFamilyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFamilyParamsWithHTTPClient(client *http.Client) *DeleteFamilyParams {
	var ()
	return &DeleteFamilyParams{
		HTTPClient: client,
	}
}

/*DeleteFamilyParams contains all the parameters to send to the API endpoint
for the delete family operation typically these are written to a http.Request
*/
type DeleteFamilyParams struct {

	/*FamilyID
	  The id of the Family to delete.

	*/
	FamilyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete family params
func (o *DeleteFamilyParams) WithTimeout(timeout time.Duration) *DeleteFamilyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete family params
func (o *DeleteFamilyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete family params
func (o *DeleteFamilyParams) WithContext(ctx context.Context) *DeleteFamilyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete family params
func (o *DeleteFamilyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete family params
func (o *DeleteFamilyParams) WithHTTPClient(client *http.Client) *DeleteFamilyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete family params
func (o *DeleteFamilyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFamilyID adds the familyID to the delete family params
func (o *DeleteFamilyParams) WithFamilyID(familyID string) *DeleteFamilyParams {
	o.SetFamilyID(familyID)
	return o
}

// SetFamilyID adds the familyId to the delete family params
func (o *DeleteFamilyParams) SetFamilyID(familyID string) {
	o.FamilyID = familyID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFamilyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param family_id
	if err := r.SetPathParam("family_id", o.FamilyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}