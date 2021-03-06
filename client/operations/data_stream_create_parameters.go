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

// NewDataStreamCreateParams creates a new DataStreamCreateParams object
// with the default values initialized.
func NewDataStreamCreateParams() *DataStreamCreateParams {
	var ()
	return &DataStreamCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataStreamCreateParamsWithTimeout creates a new DataStreamCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataStreamCreateParamsWithTimeout(timeout time.Duration) *DataStreamCreateParams {
	var ()
	return &DataStreamCreateParams{

		timeout: timeout,
	}
}

// NewDataStreamCreateParamsWithContext creates a new DataStreamCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataStreamCreateParamsWithContext(ctx context.Context) *DataStreamCreateParams {
	var ()
	return &DataStreamCreateParams{

		Context: ctx,
	}
}

// NewDataStreamCreateParamsWithHTTPClient creates a new DataStreamCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataStreamCreateParamsWithHTTPClient(client *http.Client) *DataStreamCreateParams {
	var ()
	return &DataStreamCreateParams{
		HTTPClient: client,
	}
}

/*DataStreamCreateParams contains all the parameters to send to the API endpoint
for the data stream create operation typically these are written to a http.Request
*/
type DataStreamCreateParams struct {

	/*Authorization*/
	Authorization string
	/*Body
	  This is a datastream creation request description

	*/
	Body *models.DataStream

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data stream create params
func (o *DataStreamCreateParams) WithTimeout(timeout time.Duration) *DataStreamCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data stream create params
func (o *DataStreamCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data stream create params
func (o *DataStreamCreateParams) WithContext(ctx context.Context) *DataStreamCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data stream create params
func (o *DataStreamCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data stream create params
func (o *DataStreamCreateParams) WithHTTPClient(client *http.Client) *DataStreamCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data stream create params
func (o *DataStreamCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data stream create params
func (o *DataStreamCreateParams) WithAuthorization(authorization string) *DataStreamCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data stream create params
func (o *DataStreamCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the data stream create params
func (o *DataStreamCreateParams) WithBody(body *models.DataStream) *DataStreamCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data stream create params
func (o *DataStreamCreateParams) SetBody(body *models.DataStream) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DataStreamCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
