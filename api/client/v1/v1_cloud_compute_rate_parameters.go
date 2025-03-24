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
)

// NewV1CloudComputeRateParams creates a new V1CloudComputeRateParams object
// with the default values initialized.
func NewV1CloudComputeRateParams() *V1CloudComputeRateParams {
	var ()
	return &V1CloudComputeRateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudComputeRateParamsWithTimeout creates a new V1CloudComputeRateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudComputeRateParamsWithTimeout(timeout time.Duration) *V1CloudComputeRateParams {
	var ()
	return &V1CloudComputeRateParams{

		timeout: timeout,
	}
}

// NewV1CloudComputeRateParamsWithContext creates a new V1CloudComputeRateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudComputeRateParamsWithContext(ctx context.Context) *V1CloudComputeRateParams {
	var ()
	return &V1CloudComputeRateParams{

		Context: ctx,
	}
}

// NewV1CloudComputeRateParamsWithHTTPClient creates a new V1CloudComputeRateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudComputeRateParamsWithHTTPClient(client *http.Client) *V1CloudComputeRateParams {
	var ()
	return &V1CloudComputeRateParams{
		HTTPClient: client,
	}
}

/*V1CloudComputeRateParams contains all the parameters to send to the API endpoint
for the v1 cloud compute rate operation typically these are written to a http.Request
*/
type V1CloudComputeRateParams struct {

	/*Cloud
	  cloud for which compute rate is requested

	*/
	Cloud string
	/*Region
	  region for which compute rate is requested

	*/
	Region string
	/*Type
	  instance type for which compute rate is requested

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) WithTimeout(timeout time.Duration) *V1CloudComputeRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) WithContext(ctx context.Context) *V1CloudComputeRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) WithHTTPClient(client *http.Client) *V1CloudComputeRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloud adds the cloud to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) WithCloud(cloud string) *V1CloudComputeRateParams {
	o.SetCloud(cloud)
	return o
}

// SetCloud adds the cloud to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) SetCloud(cloud string) {
	o.Cloud = cloud
}

// WithRegion adds the region to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) WithRegion(region string) *V1CloudComputeRateParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) SetRegion(region string) {
	o.Region = region
}

// WithType adds the typeVar to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) WithType(typeVar string) *V1CloudComputeRateParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the v1 cloud compute rate params
func (o *V1CloudComputeRateParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudComputeRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud
	if err := r.SetPathParam("cloud", o.Cloud); err != nil {
		return err
	}

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {
		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
