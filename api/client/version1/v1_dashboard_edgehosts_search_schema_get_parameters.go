// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewV1DashboardEdgehostsSearchSchemaGetParams creates a new V1DashboardEdgehostsSearchSchemaGetParams object
// with the default values initialized.
func NewV1DashboardEdgehostsSearchSchemaGetParams() *V1DashboardEdgehostsSearchSchemaGetParams {

	return &V1DashboardEdgehostsSearchSchemaGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardEdgehostsSearchSchemaGetParamsWithTimeout creates a new V1DashboardEdgehostsSearchSchemaGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardEdgehostsSearchSchemaGetParamsWithTimeout(timeout time.Duration) *V1DashboardEdgehostsSearchSchemaGetParams {

	return &V1DashboardEdgehostsSearchSchemaGetParams{

		timeout: timeout,
	}
}

// NewV1DashboardEdgehostsSearchSchemaGetParamsWithContext creates a new V1DashboardEdgehostsSearchSchemaGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardEdgehostsSearchSchemaGetParamsWithContext(ctx context.Context) *V1DashboardEdgehostsSearchSchemaGetParams {

	return &V1DashboardEdgehostsSearchSchemaGetParams{

		Context: ctx,
	}
}

// NewV1DashboardEdgehostsSearchSchemaGetParamsWithHTTPClient creates a new V1DashboardEdgehostsSearchSchemaGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardEdgehostsSearchSchemaGetParamsWithHTTPClient(client *http.Client) *V1DashboardEdgehostsSearchSchemaGetParams {

	return &V1DashboardEdgehostsSearchSchemaGetParams{
		HTTPClient: client,
	}
}

/*
V1DashboardEdgehostsSearchSchemaGetParams contains all the parameters to send to the API endpoint
for the v1 dashboard edgehosts search schema get operation typically these are written to a http.Request
*/
type V1DashboardEdgehostsSearchSchemaGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard edgehosts search schema get params
func (o *V1DashboardEdgehostsSearchSchemaGetParams) WithTimeout(timeout time.Duration) *V1DashboardEdgehostsSearchSchemaGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard edgehosts search schema get params
func (o *V1DashboardEdgehostsSearchSchemaGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard edgehosts search schema get params
func (o *V1DashboardEdgehostsSearchSchemaGetParams) WithContext(ctx context.Context) *V1DashboardEdgehostsSearchSchemaGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard edgehosts search schema get params
func (o *V1DashboardEdgehostsSearchSchemaGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard edgehosts search schema get params
func (o *V1DashboardEdgehostsSearchSchemaGetParams) WithHTTPClient(client *http.Client) *V1DashboardEdgehostsSearchSchemaGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard edgehosts search schema get params
func (o *V1DashboardEdgehostsSearchSchemaGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardEdgehostsSearchSchemaGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
