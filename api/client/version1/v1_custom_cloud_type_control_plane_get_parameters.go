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

// NewV1CustomCloudTypeControlPlaneGetParams creates a new V1CustomCloudTypeControlPlaneGetParams object
// with the default values initialized.
func NewV1CustomCloudTypeControlPlaneGetParams() *V1CustomCloudTypeControlPlaneGetParams {
	var ()
	return &V1CustomCloudTypeControlPlaneGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeControlPlaneGetParamsWithTimeout creates a new V1CustomCloudTypeControlPlaneGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeControlPlaneGetParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlaneGetParams {
	var ()
	return &V1CustomCloudTypeControlPlaneGetParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeControlPlaneGetParamsWithContext creates a new V1CustomCloudTypeControlPlaneGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeControlPlaneGetParamsWithContext(ctx context.Context) *V1CustomCloudTypeControlPlaneGetParams {
	var ()
	return &V1CustomCloudTypeControlPlaneGetParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeControlPlaneGetParamsWithHTTPClient creates a new V1CustomCloudTypeControlPlaneGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeControlPlaneGetParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlaneGetParams {
	var ()
	return &V1CustomCloudTypeControlPlaneGetParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeControlPlaneGetParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type control plane get operation typically these are written to a http.Request
*/
type V1CustomCloudTypeControlPlaneGetParams struct {

	/*CloudType
	  Unique cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlaneGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) WithContext(ctx context.Context) *V1CustomCloudTypeControlPlaneGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlaneGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) WithCloudType(cloudType string) *V1CustomCloudTypeControlPlaneGetParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type control plane get params
func (o *V1CustomCloudTypeControlPlaneGetParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeControlPlaneGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
