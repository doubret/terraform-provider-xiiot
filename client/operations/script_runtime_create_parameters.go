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

// NewScriptRuntimeCreateParams creates a new ScriptRuntimeCreateParams object
// with the default values initialized.
func NewScriptRuntimeCreateParams() *ScriptRuntimeCreateParams {
	var ()
	return &ScriptRuntimeCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScriptRuntimeCreateParamsWithTimeout creates a new ScriptRuntimeCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScriptRuntimeCreateParamsWithTimeout(timeout time.Duration) *ScriptRuntimeCreateParams {
	var ()
	return &ScriptRuntimeCreateParams{

		timeout: timeout,
	}
}

// NewScriptRuntimeCreateParamsWithContext creates a new ScriptRuntimeCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewScriptRuntimeCreateParamsWithContext(ctx context.Context) *ScriptRuntimeCreateParams {
	var ()
	return &ScriptRuntimeCreateParams{

		Context: ctx,
	}
}

// NewScriptRuntimeCreateParamsWithHTTPClient creates a new ScriptRuntimeCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScriptRuntimeCreateParamsWithHTTPClient(client *http.Client) *ScriptRuntimeCreateParams {
	var ()
	return &ScriptRuntimeCreateParams{
		HTTPClient: client,
	}
}

/*ScriptRuntimeCreateParams contains all the parameters to send to the API endpoint
for the script runtime create operation typically these are written to a http.Request
*/
type ScriptRuntimeCreateParams struct {

	/*Authorization*/
	Authorization string
	/*Body
	  This is a script runtime creation request description

	*/
	Body *models.ScriptRuntime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the script runtime create params
func (o *ScriptRuntimeCreateParams) WithTimeout(timeout time.Duration) *ScriptRuntimeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the script runtime create params
func (o *ScriptRuntimeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the script runtime create params
func (o *ScriptRuntimeCreateParams) WithContext(ctx context.Context) *ScriptRuntimeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the script runtime create params
func (o *ScriptRuntimeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the script runtime create params
func (o *ScriptRuntimeCreateParams) WithHTTPClient(client *http.Client) *ScriptRuntimeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the script runtime create params
func (o *ScriptRuntimeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the script runtime create params
func (o *ScriptRuntimeCreateParams) WithAuthorization(authorization string) *ScriptRuntimeCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the script runtime create params
func (o *ScriptRuntimeCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the script runtime create params
func (o *ScriptRuntimeCreateParams) WithBody(body *models.ScriptRuntime) *ScriptRuntimeCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the script runtime create params
func (o *ScriptRuntimeCreateParams) SetBody(body *models.ScriptRuntime) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ScriptRuntimeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
