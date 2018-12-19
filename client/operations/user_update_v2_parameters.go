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

// NewUserUpdateV2Params creates a new UserUpdateV2Params object
// with the default values initialized.
func NewUserUpdateV2Params() *UserUpdateV2Params {
	var ()
	return &UserUpdateV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserUpdateV2ParamsWithTimeout creates a new UserUpdateV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserUpdateV2ParamsWithTimeout(timeout time.Duration) *UserUpdateV2Params {
	var ()
	return &UserUpdateV2Params{

		timeout: timeout,
	}
}

// NewUserUpdateV2ParamsWithContext creates a new UserUpdateV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewUserUpdateV2ParamsWithContext(ctx context.Context) *UserUpdateV2Params {
	var ()
	return &UserUpdateV2Params{

		Context: ctx,
	}
}

// NewUserUpdateV2ParamsWithHTTPClient creates a new UserUpdateV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserUpdateV2ParamsWithHTTPClient(client *http.Client) *UserUpdateV2Params {
	var ()
	return &UserUpdateV2Params{
		HTTPClient: client,
	}
}

/*UserUpdateV2Params contains all the parameters to send to the API endpoint
for the user update v2 operation typically these are written to a http.Request
*/
type UserUpdateV2Params struct {

	/*Authorization*/
	Authorization string
	/*Body*/
	Body *models.User
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user update v2 params
func (o *UserUpdateV2Params) WithTimeout(timeout time.Duration) *UserUpdateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user update v2 params
func (o *UserUpdateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user update v2 params
func (o *UserUpdateV2Params) WithContext(ctx context.Context) *UserUpdateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user update v2 params
func (o *UserUpdateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user update v2 params
func (o *UserUpdateV2Params) WithHTTPClient(client *http.Client) *UserUpdateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user update v2 params
func (o *UserUpdateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user update v2 params
func (o *UserUpdateV2Params) WithAuthorization(authorization string) *UserUpdateV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user update v2 params
func (o *UserUpdateV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the user update v2 params
func (o *UserUpdateV2Params) WithBody(body *models.User) *UserUpdateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user update v2 params
func (o *UserUpdateV2Params) SetBody(body *models.User) {
	o.Body = body
}

// WithID adds the id to the user update v2 params
func (o *UserUpdateV2Params) WithID(id string) *UserUpdateV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the user update v2 params
func (o *UserUpdateV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserUpdateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
