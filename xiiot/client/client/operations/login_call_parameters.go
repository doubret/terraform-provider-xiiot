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

// NewLoginCallParams creates a new LoginCallParams object
// with the default values initialized.
func NewLoginCallParams() *LoginCallParams {
	var ()
	return &LoginCallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoginCallParamsWithTimeout creates a new LoginCallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoginCallParamsWithTimeout(timeout time.Duration) *LoginCallParams {
	var ()
	return &LoginCallParams{

		timeout: timeout,
	}
}

// NewLoginCallParamsWithContext creates a new LoginCallParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoginCallParamsWithContext(ctx context.Context) *LoginCallParams {
	var ()
	return &LoginCallParams{

		Context: ctx,
	}
}

// NewLoginCallParamsWithHTTPClient creates a new LoginCallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoginCallParamsWithHTTPClient(client *http.Client) *LoginCallParams {
	var ()
	return &LoginCallParams{
		HTTPClient: client,
	}
}

/*LoginCallParams contains all the parameters to send to the API endpoint
for the login call operation typically these are written to a http.Request
*/
type LoginCallParams struct {

	/*Request
	  This is a login credential

	*/
	Request *models.Credential

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the login call params
func (o *LoginCallParams) WithTimeout(timeout time.Duration) *LoginCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login call params
func (o *LoginCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login call params
func (o *LoginCallParams) WithContext(ctx context.Context) *LoginCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login call params
func (o *LoginCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login call params
func (o *LoginCallParams) WithHTTPClient(client *http.Client) *LoginCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login call params
func (o *LoginCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the login call params
func (o *LoginCallParams) WithRequest(request *models.Credential) *LoginCallParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the login call params
func (o *LoginCallParams) SetRequest(request *models.Credential) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *LoginCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
