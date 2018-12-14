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

// NewEdgeListParams creates a new EdgeListParams object
// with the default values initialized.
func NewEdgeListParams() *EdgeListParams {
	var ()
	return &EdgeListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeListParamsWithTimeout creates a new EdgeListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEdgeListParamsWithTimeout(timeout time.Duration) *EdgeListParams {
	var ()
	return &EdgeListParams{

		timeout: timeout,
	}
}

// NewEdgeListParamsWithContext creates a new EdgeListParams object
// with the default values initialized, and the ability to set a context for a request
func NewEdgeListParamsWithContext(ctx context.Context) *EdgeListParams {
	var ()
	return &EdgeListParams{

		Context: ctx,
	}
}

// NewEdgeListParamsWithHTTPClient creates a new EdgeListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEdgeListParamsWithHTTPClient(client *http.Client) *EdgeListParams {
	var ()
	return &EdgeListParams{
		HTTPClient: client,
	}
}

/*EdgeListParams contains all the parameters to send to the API endpoint
for the edge list operation typically these are written to a http.Request
*/
type EdgeListParams struct {

	/*Authorization*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edge list params
func (o *EdgeListParams) WithTimeout(timeout time.Duration) *EdgeListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge list params
func (o *EdgeListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge list params
func (o *EdgeListParams) WithContext(ctx context.Context) *EdgeListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge list params
func (o *EdgeListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge list params
func (o *EdgeListParams) WithHTTPClient(client *http.Client) *EdgeListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge list params
func (o *EdgeListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the edge list params
func (o *EdgeListParams) WithAuthorization(authorization string) *EdgeListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the edge list params
func (o *EdgeListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
