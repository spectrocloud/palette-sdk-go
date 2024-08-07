// Code generated by go-swagger; DO NOT EDIT.

package v1

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
	"github.com/go-openapi/swag"
)

// NewV1CoxEdgeInstanceTypesParams creates a new V1CoxEdgeInstanceTypesParams object
// with the default values initialized.
func NewV1CoxEdgeInstanceTypesParams() *V1CoxEdgeInstanceTypesParams {
	var ()
	return &V1CoxEdgeInstanceTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CoxEdgeInstanceTypesParamsWithTimeout creates a new V1CoxEdgeInstanceTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CoxEdgeInstanceTypesParamsWithTimeout(timeout time.Duration) *V1CoxEdgeInstanceTypesParams {
	var ()
	return &V1CoxEdgeInstanceTypesParams{

		timeout: timeout,
	}
}

// NewV1CoxEdgeInstanceTypesParamsWithContext creates a new V1CoxEdgeInstanceTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CoxEdgeInstanceTypesParamsWithContext(ctx context.Context) *V1CoxEdgeInstanceTypesParams {
	var ()
	return &V1CoxEdgeInstanceTypesParams{

		Context: ctx,
	}
}

// NewV1CoxEdgeInstanceTypesParamsWithHTTPClient creates a new V1CoxEdgeInstanceTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CoxEdgeInstanceTypesParamsWithHTTPClient(client *http.Client) *V1CoxEdgeInstanceTypesParams {
	var ()
	return &V1CoxEdgeInstanceTypesParams{
		HTTPClient: client,
	}
}

/*V1CoxEdgeInstanceTypesParams contains all the parameters to send to the API endpoint
for the v1 cox edge instance types operation typically these are written to a http.Request
*/
type V1CoxEdgeInstanceTypesParams struct {

	/*CPUGtEq
	  Filter for instances having cpu greater than or equal

	*/
	CPUGtEq *float64
	/*MemoryGtEq
	  Filter for instances having memory greater than or equal

	*/
	MemoryGtEq *float64
	/*Region
	  Region for which CoxEdge instances are listed

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) WithTimeout(timeout time.Duration) *V1CoxEdgeInstanceTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) WithContext(ctx context.Context) *V1CoxEdgeInstanceTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) WithHTTPClient(client *http.Client) *V1CoxEdgeInstanceTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCPUGtEq adds the cPUGtEq to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) WithCPUGtEq(cPUGtEq *float64) *V1CoxEdgeInstanceTypesParams {
	o.SetCPUGtEq(cPUGtEq)
	return o
}

// SetCPUGtEq adds the cpuGtEq to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) SetCPUGtEq(cPUGtEq *float64) {
	o.CPUGtEq = cPUGtEq
}

// WithMemoryGtEq adds the memoryGtEq to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) WithMemoryGtEq(memoryGtEq *float64) *V1CoxEdgeInstanceTypesParams {
	o.SetMemoryGtEq(memoryGtEq)
	return o
}

// SetMemoryGtEq adds the memoryGtEq to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) SetMemoryGtEq(memoryGtEq *float64) {
	o.MemoryGtEq = memoryGtEq
}

// WithRegion adds the region to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) WithRegion(region string) *V1CoxEdgeInstanceTypesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 cox edge instance types params
func (o *V1CoxEdgeInstanceTypesParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1CoxEdgeInstanceTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CPUGtEq != nil {

		// query param cpuGtEq
		var qrCPUGtEq float64
		if o.CPUGtEq != nil {
			qrCPUGtEq = *o.CPUGtEq
		}
		qCPUGtEq := swag.FormatFloat64(qrCPUGtEq)
		if qCPUGtEq != "" {
			if err := r.SetQueryParam("cpuGtEq", qCPUGtEq); err != nil {
				return err
			}
		}

	}

	if o.MemoryGtEq != nil {

		// query param memoryGtEq
		var qrMemoryGtEq float64
		if o.MemoryGtEq != nil {
			qrMemoryGtEq = *o.MemoryGtEq
		}
		qMemoryGtEq := swag.FormatFloat64(qrMemoryGtEq)
		if qMemoryGtEq != "" {
			if err := r.SetQueryParam("memoryGtEq", qMemoryGtEq); err != nil {
				return err
			}
		}

	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
