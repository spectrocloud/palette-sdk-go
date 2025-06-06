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

// NewV1CustomCloudTypeControlPlaneDeleteParams creates a new V1CustomCloudTypeControlPlaneDeleteParams object
// with the default values initialized.
func NewV1CustomCloudTypeControlPlaneDeleteParams() *V1CustomCloudTypeControlPlaneDeleteParams {
	var ()
	return &V1CustomCloudTypeControlPlaneDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeControlPlaneDeleteParamsWithTimeout creates a new V1CustomCloudTypeControlPlaneDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeControlPlaneDeleteParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlaneDeleteParams {
	var ()
	return &V1CustomCloudTypeControlPlaneDeleteParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeControlPlaneDeleteParamsWithContext creates a new V1CustomCloudTypeControlPlaneDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeControlPlaneDeleteParamsWithContext(ctx context.Context) *V1CustomCloudTypeControlPlaneDeleteParams {
	var ()
	return &V1CustomCloudTypeControlPlaneDeleteParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeControlPlaneDeleteParamsWithHTTPClient creates a new V1CustomCloudTypeControlPlaneDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeControlPlaneDeleteParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlaneDeleteParams {
	var ()
	return &V1CustomCloudTypeControlPlaneDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeControlPlaneDeleteParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type control plane delete operation typically these are written to a http.Request
*/
type V1CustomCloudTypeControlPlaneDeleteParams struct {

	/*CloudType
	  Unique cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlaneDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) WithContext(ctx context.Context) *V1CustomCloudTypeControlPlaneDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlaneDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) WithCloudType(cloudType string) *V1CustomCloudTypeControlPlaneDeleteParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type control plane delete params
func (o *V1CustomCloudTypeControlPlaneDeleteParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeControlPlaneDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
