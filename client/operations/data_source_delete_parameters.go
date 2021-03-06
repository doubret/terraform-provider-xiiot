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

// NewDataSourceDeleteParams creates a new DataSourceDeleteParams object
// with the default values initialized.
func NewDataSourceDeleteParams() *DataSourceDeleteParams {
	var ()
	return &DataSourceDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataSourceDeleteParamsWithTimeout creates a new DataSourceDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataSourceDeleteParamsWithTimeout(timeout time.Duration) *DataSourceDeleteParams {
	var ()
	return &DataSourceDeleteParams{

		timeout: timeout,
	}
}

// NewDataSourceDeleteParamsWithContext creates a new DataSourceDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataSourceDeleteParamsWithContext(ctx context.Context) *DataSourceDeleteParams {
	var ()
	return &DataSourceDeleteParams{

		Context: ctx,
	}
}

// NewDataSourceDeleteParamsWithHTTPClient creates a new DataSourceDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataSourceDeleteParamsWithHTTPClient(client *http.Client) *DataSourceDeleteParams {
	var ()
	return &DataSourceDeleteParams{
		HTTPClient: client,
	}
}

/*DataSourceDeleteParams contains all the parameters to send to the API endpoint
for the data source delete operation typically these are written to a http.Request
*/
type DataSourceDeleteParams struct {

	/*Authorization*/
	Authorization string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data source delete params
func (o *DataSourceDeleteParams) WithTimeout(timeout time.Duration) *DataSourceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data source delete params
func (o *DataSourceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data source delete params
func (o *DataSourceDeleteParams) WithContext(ctx context.Context) *DataSourceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data source delete params
func (o *DataSourceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data source delete params
func (o *DataSourceDeleteParams) WithHTTPClient(client *http.Client) *DataSourceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data source delete params
func (o *DataSourceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data source delete params
func (o *DataSourceDeleteParams) WithAuthorization(authorization string) *DataSourceDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data source delete params
func (o *DataSourceDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the data source delete params
func (o *DataSourceDeleteParams) WithID(id string) *DataSourceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the data source delete params
func (o *DataSourceDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DataSourceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
