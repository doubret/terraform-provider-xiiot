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

// NewScriptRuntimeDeleteParams creates a new ScriptRuntimeDeleteParams object
// with the default values initialized.
func NewScriptRuntimeDeleteParams() *ScriptRuntimeDeleteParams {
	var ()
	return &ScriptRuntimeDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScriptRuntimeDeleteParamsWithTimeout creates a new ScriptRuntimeDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScriptRuntimeDeleteParamsWithTimeout(timeout time.Duration) *ScriptRuntimeDeleteParams {
	var ()
	return &ScriptRuntimeDeleteParams{

		timeout: timeout,
	}
}

// NewScriptRuntimeDeleteParamsWithContext creates a new ScriptRuntimeDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewScriptRuntimeDeleteParamsWithContext(ctx context.Context) *ScriptRuntimeDeleteParams {
	var ()
	return &ScriptRuntimeDeleteParams{

		Context: ctx,
	}
}

// NewScriptRuntimeDeleteParamsWithHTTPClient creates a new ScriptRuntimeDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScriptRuntimeDeleteParamsWithHTTPClient(client *http.Client) *ScriptRuntimeDeleteParams {
	var ()
	return &ScriptRuntimeDeleteParams{
		HTTPClient: client,
	}
}

/*ScriptRuntimeDeleteParams contains all the parameters to send to the API endpoint
for the script runtime delete operation typically these are written to a http.Request
*/
type ScriptRuntimeDeleteParams struct {

	/*Authorization*/
	Authorization string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) WithTimeout(timeout time.Duration) *ScriptRuntimeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) WithContext(ctx context.Context) *ScriptRuntimeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) WithHTTPClient(client *http.Client) *ScriptRuntimeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) WithAuthorization(authorization string) *ScriptRuntimeDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) WithID(id string) *ScriptRuntimeDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the script runtime delete params
func (o *ScriptRuntimeDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ScriptRuntimeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
