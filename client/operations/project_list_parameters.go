// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewProjectListParams creates a new ProjectListParams object
// with the default values initialized.
func NewProjectListParams() *ProjectListParams {
	var ()
	return &ProjectListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectListParamsWithTimeout creates a new ProjectListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectListParamsWithTimeout(timeout time.Duration) *ProjectListParams {
	var ()
	return &ProjectListParams{

		timeout: timeout,
	}
}

// NewProjectListParamsWithContext creates a new ProjectListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectListParamsWithContext(ctx context.Context) *ProjectListParams {
	var ()
	return &ProjectListParams{

		Context: ctx,
	}
}

// NewProjectListParamsWithHTTPClient creates a new ProjectListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectListParamsWithHTTPClient(client *http.Client) *ProjectListParams {
	var ()
	return &ProjectListParams{
		HTTPClient: client,
	}
}

/*ProjectListParams contains all the parameters to send to the API endpoint
for the project list operation typically these are written to a http.Request
*/
type ProjectListParams struct {

	/*Authorization*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the project list params
func (o *ProjectListParams) WithTimeout(timeout time.Duration) *ProjectListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project list params
func (o *ProjectListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project list params
func (o *ProjectListParams) WithContext(ctx context.Context) *ProjectListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project list params
func (o *ProjectListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project list params
func (o *ProjectListParams) WithHTTPClient(client *http.Client) *ProjectListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project list params
func (o *ProjectListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the project list params
func (o *ProjectListParams) WithAuthorization(authorization string) *ProjectListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the project list params
func (o *ProjectListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
