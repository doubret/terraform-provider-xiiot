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

// NewLogRequestDownloadParams creates a new LogRequestDownloadParams object
// with the default values initialized.
func NewLogRequestDownloadParams() *LogRequestDownloadParams {
	var ()
	return &LogRequestDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogRequestDownloadParamsWithTimeout creates a new LogRequestDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogRequestDownloadParamsWithTimeout(timeout time.Duration) *LogRequestDownloadParams {
	var ()
	return &LogRequestDownloadParams{

		timeout: timeout,
	}
}

// NewLogRequestDownloadParamsWithContext creates a new LogRequestDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogRequestDownloadParamsWithContext(ctx context.Context) *LogRequestDownloadParams {
	var ()
	return &LogRequestDownloadParams{

		Context: ctx,
	}
}

// NewLogRequestDownloadParamsWithHTTPClient creates a new LogRequestDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogRequestDownloadParamsWithHTTPClient(client *http.Client) *LogRequestDownloadParams {
	var ()
	return &LogRequestDownloadParams{
		HTTPClient: client,
	}
}

/*LogRequestDownloadParams contains all the parameters to send to the API endpoint
for the log request download operation typically these are written to a http.Request
*/
type LogRequestDownloadParams struct {

	/*Authorization*/
	Authorization string
	/*Request*/
	Request *models.RequestLogDownloadPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the log request download params
func (o *LogRequestDownloadParams) WithTimeout(timeout time.Duration) *LogRequestDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log request download params
func (o *LogRequestDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log request download params
func (o *LogRequestDownloadParams) WithContext(ctx context.Context) *LogRequestDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log request download params
func (o *LogRequestDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log request download params
func (o *LogRequestDownloadParams) WithHTTPClient(client *http.Client) *LogRequestDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log request download params
func (o *LogRequestDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the log request download params
func (o *LogRequestDownloadParams) WithAuthorization(authorization string) *LogRequestDownloadParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the log request download params
func (o *LogRequestDownloadParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRequest adds the request to the log request download params
func (o *LogRequestDownloadParams) WithRequest(request *models.RequestLogDownloadPayload) *LogRequestDownloadParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the log request download params
func (o *LogRequestDownloadParams) SetRequest(request *models.RequestLogDownloadPayload) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *LogRequestDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
