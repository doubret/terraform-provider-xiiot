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

// NewSensorUpdateV2Params creates a new SensorUpdateV2Params object
// with the default values initialized.
func NewSensorUpdateV2Params() *SensorUpdateV2Params {
	var ()
	return &SensorUpdateV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewSensorUpdateV2ParamsWithTimeout creates a new SensorUpdateV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewSensorUpdateV2ParamsWithTimeout(timeout time.Duration) *SensorUpdateV2Params {
	var ()
	return &SensorUpdateV2Params{

		timeout: timeout,
	}
}

// NewSensorUpdateV2ParamsWithContext creates a new SensorUpdateV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewSensorUpdateV2ParamsWithContext(ctx context.Context) *SensorUpdateV2Params {
	var ()
	return &SensorUpdateV2Params{

		Context: ctx,
	}
}

// NewSensorUpdateV2ParamsWithHTTPClient creates a new SensorUpdateV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSensorUpdateV2ParamsWithHTTPClient(client *http.Client) *SensorUpdateV2Params {
	var ()
	return &SensorUpdateV2Params{
		HTTPClient: client,
	}
}

/*SensorUpdateV2Params contains all the parameters to send to the API endpoint
for the sensor update v2 operation typically these are written to a http.Request
*/
type SensorUpdateV2Params struct {

	/*Authorization*/
	Authorization string
	/*Body*/
	Body *models.Sensor
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sensor update v2 params
func (o *SensorUpdateV2Params) WithTimeout(timeout time.Duration) *SensorUpdateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sensor update v2 params
func (o *SensorUpdateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sensor update v2 params
func (o *SensorUpdateV2Params) WithContext(ctx context.Context) *SensorUpdateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sensor update v2 params
func (o *SensorUpdateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sensor update v2 params
func (o *SensorUpdateV2Params) WithHTTPClient(client *http.Client) *SensorUpdateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sensor update v2 params
func (o *SensorUpdateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the sensor update v2 params
func (o *SensorUpdateV2Params) WithAuthorization(authorization string) *SensorUpdateV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the sensor update v2 params
func (o *SensorUpdateV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the sensor update v2 params
func (o *SensorUpdateV2Params) WithBody(body *models.Sensor) *SensorUpdateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the sensor update v2 params
func (o *SensorUpdateV2Params) SetBody(body *models.Sensor) {
	o.Body = body
}

// WithID adds the id to the sensor update v2 params
func (o *SensorUpdateV2Params) WithID(id string) *SensorUpdateV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the sensor update v2 params
func (o *SensorUpdateV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SensorUpdateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
