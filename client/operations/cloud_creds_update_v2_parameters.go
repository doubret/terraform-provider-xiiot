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

// NewCloudCredsUpdateV2Params creates a new CloudCredsUpdateV2Params object
// with the default values initialized.
func NewCloudCredsUpdateV2Params() *CloudCredsUpdateV2Params {
	var ()
	return &CloudCredsUpdateV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudCredsUpdateV2ParamsWithTimeout creates a new CloudCredsUpdateV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudCredsUpdateV2ParamsWithTimeout(timeout time.Duration) *CloudCredsUpdateV2Params {
	var ()
	return &CloudCredsUpdateV2Params{

		timeout: timeout,
	}
}

// NewCloudCredsUpdateV2ParamsWithContext creates a new CloudCredsUpdateV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCloudCredsUpdateV2ParamsWithContext(ctx context.Context) *CloudCredsUpdateV2Params {
	var ()
	return &CloudCredsUpdateV2Params{

		Context: ctx,
	}
}

// NewCloudCredsUpdateV2ParamsWithHTTPClient creates a new CloudCredsUpdateV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudCredsUpdateV2ParamsWithHTTPClient(client *http.Client) *CloudCredsUpdateV2Params {
	var ()
	return &CloudCredsUpdateV2Params{
		HTTPClient: client,
	}
}

/*CloudCredsUpdateV2Params contains all the parameters to send to the API endpoint
for the cloud creds update v2 operation typically these are written to a http.Request
*/
type CloudCredsUpdateV2Params struct {

	/*Authorization*/
	Authorization string
	/*Body*/
	Body *models.CloudCreds
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) WithTimeout(timeout time.Duration) *CloudCredsUpdateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) WithContext(ctx context.Context) *CloudCredsUpdateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) WithHTTPClient(client *http.Client) *CloudCredsUpdateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) WithAuthorization(authorization string) *CloudCredsUpdateV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) WithBody(body *models.CloudCreds) *CloudCredsUpdateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) SetBody(body *models.CloudCreds) {
	o.Body = body
}

// WithID adds the id to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) WithID(id string) *CloudCredsUpdateV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the cloud creds update v2 params
func (o *CloudCredsUpdateV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CloudCredsUpdateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
