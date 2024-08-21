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

// NewV1CustomCloudTypeCloudProviderDeleteParams creates a new V1CustomCloudTypeCloudProviderDeleteParams object
// with the default values initialized.
func NewV1CustomCloudTypeCloudProviderDeleteParams() *V1CustomCloudTypeCloudProviderDeleteParams {
	var ()
	return &V1CustomCloudTypeCloudProviderDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeCloudProviderDeleteParamsWithTimeout creates a new V1CustomCloudTypeCloudProviderDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeCloudProviderDeleteParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeCloudProviderDeleteParams {
	var ()
	return &V1CustomCloudTypeCloudProviderDeleteParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeCloudProviderDeleteParamsWithContext creates a new V1CustomCloudTypeCloudProviderDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeCloudProviderDeleteParamsWithContext(ctx context.Context) *V1CustomCloudTypeCloudProviderDeleteParams {
	var ()
	return &V1CustomCloudTypeCloudProviderDeleteParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeCloudProviderDeleteParamsWithHTTPClient creates a new V1CustomCloudTypeCloudProviderDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeCloudProviderDeleteParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeCloudProviderDeleteParams {
	var ()
	return &V1CustomCloudTypeCloudProviderDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeCloudProviderDeleteParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type cloud provider delete operation typically these are written to a http.Request
*/
type V1CustomCloudTypeCloudProviderDeleteParams struct {

	/*CloudType
	  Unique cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeCloudProviderDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) WithContext(ctx context.Context) *V1CustomCloudTypeCloudProviderDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeCloudProviderDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) WithCloudType(cloudType string) *V1CustomCloudTypeCloudProviderDeleteParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type cloud provider delete params
func (o *V1CustomCloudTypeCloudProviderDeleteParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeCloudProviderDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}