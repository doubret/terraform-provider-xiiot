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

// NewSensorDeleteParams creates a new SensorDeleteParams object
// with the default values initialized.
func NewSensorDeleteParams() *SensorDeleteParams {
	var ()
	return &SensorDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSensorDeleteParamsWithTimeout creates a new SensorDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSensorDeleteParamsWithTimeout(timeout time.Duration) *SensorDeleteParams {
	var ()
	return &SensorDeleteParams{

		timeout: timeout,
	}
}

// NewSensorDeleteParamsWithContext creates a new SensorDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewSensorDeleteParamsWithContext(ctx context.Context) *SensorDeleteParams {
	var ()
	return &SensorDeleteParams{

		Context: ctx,
	}
}

// NewSensorDeleteParamsWithHTTPClient creates a new SensorDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSensorDeleteParamsWithHTTPClient(client *http.Client) *SensorDeleteParams {
	var ()
	return &SensorDeleteParams{
		HTTPClient: client,
	}
}

/*SensorDeleteParams contains all the parameters to send to the API endpoint
for the sensor delete operation typically these are written to a http.Request
*/
type SensorDeleteParams struct {

	/*Authorization*/
	Authorization string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sensor delete params
func (o *SensorDeleteParams) WithTimeout(timeout time.Duration) *SensorDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sensor delete params
func (o *SensorDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sensor delete params
func (o *SensorDeleteParams) WithContext(ctx context.Context) *SensorDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sensor delete params
func (o *SensorDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sensor delete params
func (o *SensorDeleteParams) WithHTTPClient(client *http.Client) *SensorDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sensor delete params
func (o *SensorDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the sensor delete params
func (o *SensorDeleteParams) WithAuthorization(authorization string) *SensorDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the sensor delete params
func (o *SensorDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the sensor delete params
func (o *SensorDeleteParams) WithID(id string) *SensorDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the sensor delete params
func (o *SensorDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SensorDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
