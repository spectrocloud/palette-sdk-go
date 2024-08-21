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

// NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParams creates a new V1CustomCloudTypeControlPlanePoolTemplateUpdateParams object
// with the default values initialized.
func NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParams() *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParamsWithTimeout creates a new V1CustomCloudTypeControlPlanePoolTemplateUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateUpdateParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParamsWithContext creates a new V1CustomCloudTypeControlPlanePoolTemplateUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParamsWithContext(ctx context.Context) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateUpdateParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParamsWithHTTPClient creates a new V1CustomCloudTypeControlPlanePoolTemplateUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeControlPlanePoolTemplateUpdateParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlanePoolTemplateUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeControlPlanePoolTemplateUpdateParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type control plane pool template update operation typically these are written to a http.Request
*/
type V1CustomCloudTypeControlPlanePoolTemplateUpdateParams struct {

	/*CloudType
	  Unique cloud type

	*/
	CloudType string
	/*FileName*/
	FileName runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) WithContext(ctx context.Context) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) WithCloudType(cloudType string) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithFileName adds the fileName to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) WithFileName(fileName runtime.NamedReadCloser) *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the v1 custom cloud type control plane pool template update params
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) SetFileName(fileName runtime.NamedReadCloser) {
	o.FileName = fileName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
	}

	if o.FileName != nil {

		if o.FileName != nil {

			// form file param fileName
			if err := r.SetFileParam("fileName", o.FileName); err != nil {
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}