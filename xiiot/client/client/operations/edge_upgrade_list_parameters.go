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

// NewEdgeUpgradeListParams creates a new EdgeUpgradeListParams object
// with the default values initialized.
func NewEdgeUpgradeListParams() *EdgeUpgradeListParams {
	var ()
	return &EdgeUpgradeListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeUpgradeListParamsWithTimeout creates a new EdgeUpgradeListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEdgeUpgradeListParamsWithTimeout(timeout time.Duration) *EdgeUpgradeListParams {
	var ()
	return &EdgeUpgradeListParams{

		timeout: timeout,
	}
}

// NewEdgeUpgradeListParamsWithContext creates a new EdgeUpgradeListParams object
// with the default values initialized, and the ability to set a context for a request
func NewEdgeUpgradeListParamsWithContext(ctx context.Context) *EdgeUpgradeListParams {
	var ()
	return &EdgeUpgradeListParams{

		Context: ctx,
	}
}

// NewEdgeUpgradeListParamsWithHTTPClient creates a new EdgeUpgradeListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEdgeUpgradeListParamsWithHTTPClient(client *http.Client) *EdgeUpgradeListParams {
	var ()
	return &EdgeUpgradeListParams{
		HTTPClient: client,
	}
}

/*EdgeUpgradeListParams contains all the parameters to send to the API endpoint
for the edge upgrade list operation typically these are written to a http.Request
*/
type EdgeUpgradeListParams struct {

	/*Authorization*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edge upgrade list params
func (o *EdgeUpgradeListParams) WithTimeout(timeout time.Duration) *EdgeUpgradeListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge upgrade list params
func (o *EdgeUpgradeListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge upgrade list params
func (o *EdgeUpgradeListParams) WithContext(ctx context.Context) *EdgeUpgradeListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge upgrade list params
func (o *EdgeUpgradeListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge upgrade list params
func (o *EdgeUpgradeListParams) WithHTTPClient(client *http.Client) *EdgeUpgradeListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge upgrade list params
func (o *EdgeUpgradeListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the edge upgrade list params
func (o *EdgeUpgradeListParams) WithAuthorization(authorization string) *EdgeUpgradeListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the edge upgrade list params
func (o *EdgeUpgradeListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeUpgradeListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
