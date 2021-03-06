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

// NewEdgeGetDatasourcesParams creates a new EdgeGetDatasourcesParams object
// with the default values initialized.
func NewEdgeGetDatasourcesParams() *EdgeGetDatasourcesParams {
	var ()
	return &EdgeGetDatasourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeGetDatasourcesParamsWithTimeout creates a new EdgeGetDatasourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEdgeGetDatasourcesParamsWithTimeout(timeout time.Duration) *EdgeGetDatasourcesParams {
	var ()
	return &EdgeGetDatasourcesParams{

		timeout: timeout,
	}
}

// NewEdgeGetDatasourcesParamsWithContext creates a new EdgeGetDatasourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewEdgeGetDatasourcesParamsWithContext(ctx context.Context) *EdgeGetDatasourcesParams {
	var ()
	return &EdgeGetDatasourcesParams{

		Context: ctx,
	}
}

// NewEdgeGetDatasourcesParamsWithHTTPClient creates a new EdgeGetDatasourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEdgeGetDatasourcesParamsWithHTTPClient(client *http.Client) *EdgeGetDatasourcesParams {
	var ()
	return &EdgeGetDatasourcesParams{
		HTTPClient: client,
	}
}

/*EdgeGetDatasourcesParams contains all the parameters to send to the API endpoint
for the edge get datasources operation typically these are written to a http.Request
*/
type EdgeGetDatasourcesParams struct {

	/*Authorization*/
	Authorization string
	/*EdgeID
	  ID for the edge

	*/
	EdgeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edge get datasources params
func (o *EdgeGetDatasourcesParams) WithTimeout(timeout time.Duration) *EdgeGetDatasourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge get datasources params
func (o *EdgeGetDatasourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge get datasources params
func (o *EdgeGetDatasourcesParams) WithContext(ctx context.Context) *EdgeGetDatasourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge get datasources params
func (o *EdgeGetDatasourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge get datasources params
func (o *EdgeGetDatasourcesParams) WithHTTPClient(client *http.Client) *EdgeGetDatasourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge get datasources params
func (o *EdgeGetDatasourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the edge get datasources params
func (o *EdgeGetDatasourcesParams) WithAuthorization(authorization string) *EdgeGetDatasourcesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the edge get datasources params
func (o *EdgeGetDatasourcesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithEdgeID adds the edgeID to the edge get datasources params
func (o *EdgeGetDatasourcesParams) WithEdgeID(edgeID string) *EdgeGetDatasourcesParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the edge get datasources params
func (o *EdgeGetDatasourcesParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeGetDatasourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param edgeId
	if err := r.SetPathParam("edgeId", o.EdgeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
