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

	models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
)

// NewDockerProfileCreateParams creates a new DockerProfileCreateParams object
// with the default values initialized.
func NewDockerProfileCreateParams() *DockerProfileCreateParams {
	var ()
	return &DockerProfileCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDockerProfileCreateParamsWithTimeout creates a new DockerProfileCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDockerProfileCreateParamsWithTimeout(timeout time.Duration) *DockerProfileCreateParams {
	var ()
	return &DockerProfileCreateParams{

		timeout: timeout,
	}
}

// NewDockerProfileCreateParamsWithContext creates a new DockerProfileCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDockerProfileCreateParamsWithContext(ctx context.Context) *DockerProfileCreateParams {
	var ()
	return &DockerProfileCreateParams{

		Context: ctx,
	}
}

// NewDockerProfileCreateParamsWithHTTPClient creates a new DockerProfileCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDockerProfileCreateParamsWithHTTPClient(client *http.Client) *DockerProfileCreateParams {
	var ()
	return &DockerProfileCreateParams{
		HTTPClient: client,
	}
}

/*DockerProfileCreateParams contains all the parameters to send to the API endpoint
for the docker profile create operation typically these are written to a http.Request
*/
type DockerProfileCreateParams struct {

	/*Authorization*/
	Authorization string
	/*Body
	  This is a DockerProfile creation request description

	*/
	Body *models.DockerProfile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the docker profile create params
func (o *DockerProfileCreateParams) WithTimeout(timeout time.Duration) *DockerProfileCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the docker profile create params
func (o *DockerProfileCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the docker profile create params
func (o *DockerProfileCreateParams) WithContext(ctx context.Context) *DockerProfileCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the docker profile create params
func (o *DockerProfileCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the docker profile create params
func (o *DockerProfileCreateParams) WithHTTPClient(client *http.Client) *DockerProfileCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the docker profile create params
func (o *DockerProfileCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the docker profile create params
func (o *DockerProfileCreateParams) WithAuthorization(authorization string) *DockerProfileCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the docker profile create params
func (o *DockerProfileCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the docker profile create params
func (o *DockerProfileCreateParams) WithBody(body *models.DockerProfile) *DockerProfileCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the docker profile create params
func (o *DockerProfileCreateParams) SetBody(body *models.DockerProfile) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DockerProfileCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
