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

// NewScriptCreateParams creates a new ScriptCreateParams object
// with the default values initialized.
func NewScriptCreateParams() *ScriptCreateParams {
	var ()
	return &ScriptCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScriptCreateParamsWithTimeout creates a new ScriptCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScriptCreateParamsWithTimeout(timeout time.Duration) *ScriptCreateParams {
	var ()
	return &ScriptCreateParams{

		timeout: timeout,
	}
}

// NewScriptCreateParamsWithContext creates a new ScriptCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewScriptCreateParamsWithContext(ctx context.Context) *ScriptCreateParams {
	var ()
	return &ScriptCreateParams{

		Context: ctx,
	}
}

// NewScriptCreateParamsWithHTTPClient creates a new ScriptCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScriptCreateParamsWithHTTPClient(client *http.Client) *ScriptCreateParams {
	var ()
	return &ScriptCreateParams{
		HTTPClient: client,
	}
}

/*ScriptCreateParams contains all the parameters to send to the API endpoint
for the script create operation typically these are written to a http.Request
*/
type ScriptCreateParams struct {

	/*Authorization*/
	Authorization string
	/*Body
	  This is a script creation request description

	*/
	Body *models.Script

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the script create params
func (o *ScriptCreateParams) WithTimeout(timeout time.Duration) *ScriptCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the script create params
func (o *ScriptCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the script create params
func (o *ScriptCreateParams) WithContext(ctx context.Context) *ScriptCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the script create params
func (o *ScriptCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the script create params
func (o *ScriptCreateParams) WithHTTPClient(client *http.Client) *ScriptCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the script create params
func (o *ScriptCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the script create params
func (o *ScriptCreateParams) WithAuthorization(authorization string) *ScriptCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the script create params
func (o *ScriptCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the script create params
func (o *ScriptCreateParams) WithBody(body *models.Script) *ScriptCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the script create params
func (o *ScriptCreateParams) SetBody(body *models.Script) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ScriptCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
