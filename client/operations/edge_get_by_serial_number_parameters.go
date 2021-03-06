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

// NewEdgeGetBySerialNumberParams creates a new EdgeGetBySerialNumberParams object
// with the default values initialized.
func NewEdgeGetBySerialNumberParams() *EdgeGetBySerialNumberParams {
	var ()
	return &EdgeGetBySerialNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeGetBySerialNumberParamsWithTimeout creates a new EdgeGetBySerialNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEdgeGetBySerialNumberParamsWithTimeout(timeout time.Duration) *EdgeGetBySerialNumberParams {
	var ()
	return &EdgeGetBySerialNumberParams{

		timeout: timeout,
	}
}

// NewEdgeGetBySerialNumberParamsWithContext creates a new EdgeGetBySerialNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewEdgeGetBySerialNumberParamsWithContext(ctx context.Context) *EdgeGetBySerialNumberParams {
	var ()
	return &EdgeGetBySerialNumberParams{

		Context: ctx,
	}
}

// NewEdgeGetBySerialNumberParamsWithHTTPClient creates a new EdgeGetBySerialNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEdgeGetBySerialNumberParamsWithHTTPClient(client *http.Client) *EdgeGetBySerialNumberParams {
	var ()
	return &EdgeGetBySerialNumberParams{
		HTTPClient: client,
	}
}

/*EdgeGetBySerialNumberParams contains all the parameters to send to the API endpoint
for the edge get by serial number operation typically these are written to a http.Request
*/
type EdgeGetBySerialNumberParams struct {

	/*Body
	  JSON { serialNumber: string }

	*/
	Body *models.SerialNumberPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) WithTimeout(timeout time.Duration) *EdgeGetBySerialNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) WithContext(ctx context.Context) *EdgeGetBySerialNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) WithHTTPClient(client *http.Client) *EdgeGetBySerialNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) WithBody(body *models.SerialNumberPayload) *EdgeGetBySerialNumberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge get by serial number params
func (o *EdgeGetBySerialNumberParams) SetBody(body *models.SerialNumberPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeGetBySerialNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
