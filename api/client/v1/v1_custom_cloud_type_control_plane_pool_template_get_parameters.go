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

// NewV1CustomCloudTypeControlPlanePoolTemplateGetParams creates a new V1CustomCloudTypeControlPlanePoolTemplateGetParams object
// with the default values initialized.
func NewV1CustomCloudTypeControlPlanePoolTemplateGetParams() *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateGetParamsWithTimeout creates a new V1CustomCloudTypeControlPlanePoolTemplateGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeControlPlanePoolTemplateGetParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateGetParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateGetParamsWithContext creates a new V1CustomCloudTypeControlPlanePoolTemplateGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeControlPlanePoolTemplateGetParamsWithContext(ctx context.Context) *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateGetParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateGetParamsWithHTTPClient creates a new V1CustomCloudTypeControlPlanePoolTemplateGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeControlPlanePoolTemplateGetParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateGetParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeControlPlanePoolTemplateGetParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type control plane pool template get operation typically these are written to a http.Request
*/
type V1CustomCloudTypeControlPlanePoolTemplateGetParams struct {

	/*CloudType
	  Unique cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) WithContext(ctx context.Context) *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) WithCloudType(cloudType string) *V1CustomCloudTypeControlPlanePoolTemplateGetParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type control plane pool template get params
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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