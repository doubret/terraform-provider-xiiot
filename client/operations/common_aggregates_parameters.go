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

// NewCommonAggregatesParams creates a new CommonAggregatesParams object
// with the default values initialized.
func NewCommonAggregatesParams() *CommonAggregatesParams {
	var ()
	return &CommonAggregatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCommonAggregatesParamsWithTimeout creates a new CommonAggregatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCommonAggregatesParamsWithTimeout(timeout time.Duration) *CommonAggregatesParams {
	var ()
	return &CommonAggregatesParams{

		timeout: timeout,
	}
}

// NewCommonAggregatesParamsWithContext creates a new CommonAggregatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewCommonAggregatesParamsWithContext(ctx context.Context) *CommonAggregatesParams {
	var ()
	return &CommonAggregatesParams{

		Context: ctx,
	}
}

// NewCommonAggregatesParamsWithHTTPClient creates a new CommonAggregatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCommonAggregatesParamsWithHTTPClient(client *http.Client) *CommonAggregatesParams {
	var ()
	return &CommonAggregatesParams{
		HTTPClient: client,
	}
}

/*CommonAggregatesParams contains all the parameters to send to the API endpoint
for the common aggregates operation typically these are written to a http.Request
*/
type CommonAggregatesParams struct {

	/*AggregateSpec*/
	AggregateSpec *models.AggregateSpec
	/*Authorization*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the common aggregates params
func (o *CommonAggregatesParams) WithTimeout(timeout time.Duration) *CommonAggregatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the common aggregates params
func (o *CommonAggregatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the common aggregates params
func (o *CommonAggregatesParams) WithContext(ctx context.Context) *CommonAggregatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the common aggregates params
func (o *CommonAggregatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the common aggregates params
func (o *CommonAggregatesParams) WithHTTPClient(client *http.Client) *CommonAggregatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the common aggregates params
func (o *CommonAggregatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAggregateSpec adds the aggregateSpec to the common aggregates params
func (o *CommonAggregatesParams) WithAggregateSpec(aggregateSpec *models.AggregateSpec) *CommonAggregatesParams {
	o.SetAggregateSpec(aggregateSpec)
	return o
}

// SetAggregateSpec adds the aggregateSpec to the common aggregates params
func (o *CommonAggregatesParams) SetAggregateSpec(aggregateSpec *models.AggregateSpec) {
	o.AggregateSpec = aggregateSpec
}

// WithAuthorization adds the authorization to the common aggregates params
func (o *CommonAggregatesParams) WithAuthorization(authorization string) *CommonAggregatesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the common aggregates params
func (o *CommonAggregatesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *CommonAggregatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AggregateSpec != nil {
		if err := r.SetBodyParam(o.AggregateSpec); err != nil {
			return err
		}
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
