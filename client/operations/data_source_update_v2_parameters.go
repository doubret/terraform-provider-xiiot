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

// NewDataSourceUpdateV2Params creates a new DataSourceUpdateV2Params object
// with the default values initialized.
func NewDataSourceUpdateV2Params() *DataSourceUpdateV2Params {
	var ()
	return &DataSourceUpdateV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataSourceUpdateV2ParamsWithTimeout creates a new DataSourceUpdateV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataSourceUpdateV2ParamsWithTimeout(timeout time.Duration) *DataSourceUpdateV2Params {
	var ()
	return &DataSourceUpdateV2Params{

		timeout: timeout,
	}
}

// NewDataSourceUpdateV2ParamsWithContext creates a new DataSourceUpdateV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDataSourceUpdateV2ParamsWithContext(ctx context.Context) *DataSourceUpdateV2Params {
	var ()
	return &DataSourceUpdateV2Params{

		Context: ctx,
	}
}

// NewDataSourceUpdateV2ParamsWithHTTPClient creates a new DataSourceUpdateV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataSourceUpdateV2ParamsWithHTTPClient(client *http.Client) *DataSourceUpdateV2Params {
	var ()
	return &DataSourceUpdateV2Params{
		HTTPClient: client,
	}
}

/*DataSourceUpdateV2Params contains all the parameters to send to the API endpoint
for the data source update v2 operation typically these are written to a http.Request
*/
type DataSourceUpdateV2Params struct {

	/*Authorization*/
	Authorization string
	/*Body*/
	Body *models.DataSource
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data source update v2 params
func (o *DataSourceUpdateV2Params) WithTimeout(timeout time.Duration) *DataSourceUpdateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data source update v2 params
func (o *DataSourceUpdateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data source update v2 params
func (o *DataSourceUpdateV2Params) WithContext(ctx context.Context) *DataSourceUpdateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data source update v2 params
func (o *DataSourceUpdateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data source update v2 params
func (o *DataSourceUpdateV2Params) WithHTTPClient(client *http.Client) *DataSourceUpdateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data source update v2 params
func (o *DataSourceUpdateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data source update v2 params
func (o *DataSourceUpdateV2Params) WithAuthorization(authorization string) *DataSourceUpdateV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data source update v2 params
func (o *DataSourceUpdateV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the data source update v2 params
func (o *DataSourceUpdateV2Params) WithBody(body *models.DataSource) *DataSourceUpdateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data source update v2 params
func (o *DataSourceUpdateV2Params) SetBody(body *models.DataSource) {
	o.Body = body
}

// WithID adds the id to the data source update v2 params
func (o *DataSourceUpdateV2Params) WithID(id string) *DataSourceUpdateV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the data source update v2 params
func (o *DataSourceUpdateV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DataSourceUpdateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
