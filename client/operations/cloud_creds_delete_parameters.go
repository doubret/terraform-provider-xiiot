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

// NewCloudCredsDeleteParams creates a new CloudCredsDeleteParams object
// with the default values initialized.
func NewCloudCredsDeleteParams() *CloudCredsDeleteParams {
	var ()
	return &CloudCredsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudCredsDeleteParamsWithTimeout creates a new CloudCredsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudCredsDeleteParamsWithTimeout(timeout time.Duration) *CloudCredsDeleteParams {
	var ()
	return &CloudCredsDeleteParams{

		timeout: timeout,
	}
}

// NewCloudCredsDeleteParamsWithContext creates a new CloudCredsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloudCredsDeleteParamsWithContext(ctx context.Context) *CloudCredsDeleteParams {
	var ()
	return &CloudCredsDeleteParams{

		Context: ctx,
	}
}

// NewCloudCredsDeleteParamsWithHTTPClient creates a new CloudCredsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudCredsDeleteParamsWithHTTPClient(client *http.Client) *CloudCredsDeleteParams {
	var ()
	return &CloudCredsDeleteParams{
		HTTPClient: client,
	}
}

/*CloudCredsDeleteParams contains all the parameters to send to the API endpoint
for the cloud creds delete operation typically these are written to a http.Request
*/
type CloudCredsDeleteParams struct {

	/*Authorization*/
	Authorization string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cloud creds delete params
func (o *CloudCredsDeleteParams) WithTimeout(timeout time.Duration) *CloudCredsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud creds delete params
func (o *CloudCredsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud creds delete params
func (o *CloudCredsDeleteParams) WithContext(ctx context.Context) *CloudCredsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud creds delete params
func (o *CloudCredsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud creds delete params
func (o *CloudCredsDeleteParams) WithHTTPClient(client *http.Client) *CloudCredsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud creds delete params
func (o *CloudCredsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the cloud creds delete params
func (o *CloudCredsDeleteParams) WithAuthorization(authorization string) *CloudCredsDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the cloud creds delete params
func (o *CloudCredsDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the cloud creds delete params
func (o *CloudCredsDeleteParams) WithID(id string) *CloudCredsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cloud creds delete params
func (o *CloudCredsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CloudCredsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
