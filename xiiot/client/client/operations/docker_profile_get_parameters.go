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

// NewDockerProfileGetParams creates a new DockerProfileGetParams object
// with the default values initialized.
func NewDockerProfileGetParams() *DockerProfileGetParams {
	var ()
	return &DockerProfileGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDockerProfileGetParamsWithTimeout creates a new DockerProfileGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDockerProfileGetParamsWithTimeout(timeout time.Duration) *DockerProfileGetParams {
	var ()
	return &DockerProfileGetParams{

		timeout: timeout,
	}
}

// NewDockerProfileGetParamsWithContext creates a new DockerProfileGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDockerProfileGetParamsWithContext(ctx context.Context) *DockerProfileGetParams {
	var ()
	return &DockerProfileGetParams{

		Context: ctx,
	}
}

// NewDockerProfileGetParamsWithHTTPClient creates a new DockerProfileGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDockerProfileGetParamsWithHTTPClient(client *http.Client) *DockerProfileGetParams {
	var ()
	return &DockerProfileGetParams{
		HTTPClient: client,
	}
}

/*DockerProfileGetParams contains all the parameters to send to the API endpoint
for the docker profile get operation typically these are written to a http.Request
*/
type DockerProfileGetParams struct {

	/*Authorization*/
	Authorization string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the docker profile get params
func (o *DockerProfileGetParams) WithTimeout(timeout time.Duration) *DockerProfileGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the docker profile get params
func (o *DockerProfileGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the docker profile get params
func (o *DockerProfileGetParams) WithContext(ctx context.Context) *DockerProfileGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the docker profile get params
func (o *DockerProfileGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the docker profile get params
func (o *DockerProfileGetParams) WithHTTPClient(client *http.Client) *DockerProfileGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the docker profile get params
func (o *DockerProfileGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the docker profile get params
func (o *DockerProfileGetParams) WithAuthorization(authorization string) *DockerProfileGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the docker profile get params
func (o *DockerProfileGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the docker profile get params
func (o *DockerProfileGetParams) WithID(id string) *DockerProfileGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the docker profile get params
func (o *DockerProfileGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DockerProfileGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
