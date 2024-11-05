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

// NewV1AzureInstanceTypesParams creates a new V1AzureInstanceTypesParams object
// with the default values initialized.
func NewV1AzureInstanceTypesParams() *V1AzureInstanceTypesParams {
	var ()
	return &V1AzureInstanceTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzureInstanceTypesParamsWithTimeout creates a new V1AzureInstanceTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzureInstanceTypesParamsWithTimeout(timeout time.Duration) *V1AzureInstanceTypesParams {
	var ()
	return &V1AzureInstanceTypesParams{

		timeout: timeout,
	}
}

// NewV1AzureInstanceTypesParamsWithContext creates a new V1AzureInstanceTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzureInstanceTypesParamsWithContext(ctx context.Context) *V1AzureInstanceTypesParams {
	var ()
	return &V1AzureInstanceTypesParams{

		Context: ctx,
	}
}

// NewV1AzureInstanceTypesParamsWithHTTPClient creates a new V1AzureInstanceTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzureInstanceTypesParamsWithHTTPClient(client *http.Client) *V1AzureInstanceTypesParams {
	var ()
	return &V1AzureInstanceTypesParams{
		HTTPClient: client,
	}
}

/*V1AzureInstanceTypesParams contains all the parameters to send to the API endpoint
for the v1 azure instance types operation typically these are written to a http.Request
*/
type V1AzureInstanceTypesParams struct {

	/*CPUGtEq
	  Filter for instances having cpu greater than or equal

	*/
	CPUGtEq *float64
	/*GpuGtEq
	  Filter for instances having gpu greater than or equal

	*/
	GpuGtEq *float64
	/*MemoryGtEq
	  Filter for instances having memory greater than or equal

	*/
	MemoryGtEq *float64
	/*Region
	  Region for which Azure instance types are requested

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) WithTimeout(timeout time.Duration) *V1AzureInstanceTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) WithContext(ctx context.Context) *V1AzureInstanceTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) WithHTTPClient(client *http.Client) *V1AzureInstanceTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCPUGtEq adds the cPUGtEq to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) WithCPUGtEq(cPUGtEq *float64) *V1AzureInstanceTypesParams {
	o.SetCPUGtEq(cPUGtEq)
	return o
}

// SetCPUGtEq adds the cpuGtEq to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) SetCPUGtEq(cPUGtEq *float64) {
	o.CPUGtEq = cPUGtEq
}

// WithGpuGtEq adds the gpuGtEq to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) WithGpuGtEq(gpuGtEq *float64) *V1AzureInstanceTypesParams {
	o.SetGpuGtEq(gpuGtEq)
	return o
}

// SetGpuGtEq adds the gpuGtEq to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) SetGpuGtEq(gpuGtEq *float64) {
	o.GpuGtEq = gpuGtEq
}

// WithMemoryGtEq adds the memoryGtEq to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) WithMemoryGtEq(memoryGtEq *float64) *V1AzureInstanceTypesParams {
	o.SetMemoryGtEq(memoryGtEq)
	return o
}

// SetMemoryGtEq adds the memoryGtEq to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) SetMemoryGtEq(memoryGtEq *float64) {
	o.MemoryGtEq = memoryGtEq
}

// WithRegion adds the region to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) WithRegion(region string) *V1AzureInstanceTypesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 azure instance types params
func (o *V1AzureInstanceTypesParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzureInstanceTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.GpuGtEq != nil {

		// query param gpuGtEq
		var qrGpuGtEq float64
		if o.GpuGtEq != nil {
			qrGpuGtEq = *o.GpuGtEq
		}
		qGpuGtEq := swag.FormatFloat64(qrGpuGtEq)
		if qGpuGtEq != "" {
			if err := r.SetQueryParam("gpuGtEq", qGpuGtEq); err != nil {
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
