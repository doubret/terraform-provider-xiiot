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

	models "github.com/doubret/terraform-provider-xiiot/models"
)

// NewCategoryCreateParams creates a new CategoryCreateParams object
// with the default values initialized.
func NewCategoryCreateParams() *CategoryCreateParams {
	var ()
	return &CategoryCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCategoryCreateParamsWithTimeout creates a new CategoryCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCategoryCreateParamsWithTimeout(timeout time.Duration) *CategoryCreateParams {
	var ()
	return &CategoryCreateParams{

		timeout: timeout,
	}
}

// NewCategoryCreateParamsWithContext creates a new CategoryCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCategoryCreateParamsWithContext(ctx context.Context) *CategoryCreateParams {
	var ()
	return &CategoryCreateParams{

		Context: ctx,
	}
}

// NewCategoryCreateParamsWithHTTPClient creates a new CategoryCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCategoryCreateParamsWithHTTPClient(client *http.Client) *CategoryCreateParams {
	var ()
	return &CategoryCreateParams{
		HTTPClient: client,
	}
}

/*CategoryCreateParams contains all the parameters to send to the API endpoint
for the category create operation typically these are written to a http.Request
*/
type CategoryCreateParams struct {

	/*Authorization*/
	Authorization string
	/*Body
	  This is a category creation request description

	*/
	Body *models.Category

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the category create params
func (o *CategoryCreateParams) WithTimeout(timeout time.Duration) *CategoryCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the category create params
func (o *CategoryCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the category create params
func (o *CategoryCreateParams) WithContext(ctx context.Context) *CategoryCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the category create params
func (o *CategoryCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the category create params
func (o *CategoryCreateParams) WithHTTPClient(client *http.Client) *CategoryCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the category create params
func (o *CategoryCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the category create params
func (o *CategoryCreateParams) WithAuthorization(authorization string) *CategoryCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the category create params
func (o *CategoryCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the category create params
func (o *CategoryCreateParams) WithBody(body *models.Category) *CategoryCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the category create params
func (o *CategoryCreateParams) SetBody(body *models.Category) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CategoryCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
