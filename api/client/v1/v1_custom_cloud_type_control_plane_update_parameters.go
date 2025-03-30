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

// NewV1CustomCloudTypeControlPlaneUpdateParams creates a new V1CustomCloudTypeControlPlaneUpdateParams object
// with the default values initialized.
func NewV1CustomCloudTypeControlPlaneUpdateParams() *V1CustomCloudTypeControlPlaneUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlaneUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeControlPlaneUpdateParamsWithTimeout creates a new V1CustomCloudTypeControlPlaneUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeControlPlaneUpdateParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlaneUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlaneUpdateParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeControlPlaneUpdateParamsWithContext creates a new V1CustomCloudTypeControlPlaneUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeControlPlaneUpdateParamsWithContext(ctx context.Context) *V1CustomCloudTypeControlPlaneUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlaneUpdateParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeControlPlaneUpdateParamsWithHTTPClient creates a new V1CustomCloudTypeControlPlaneUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeControlPlaneUpdateParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlaneUpdateParams {
	var ()
	return &V1CustomCloudTypeControlPlaneUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeControlPlaneUpdateParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type control plane update operation typically these are written to a http.Request
*/
type V1CustomCloudTypeControlPlaneUpdateParams struct {

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

// WithTimeout adds the timeout to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeControlPlaneUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) WithContext(ctx context.Context) *V1CustomCloudTypeControlPlaneUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeControlPlaneUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) WithCloudType(cloudType string) *V1CustomCloudTypeControlPlaneUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithFileName adds the fileName to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) WithFileName(fileName runtime.NamedReadCloser) *V1CustomCloudTypeControlPlaneUpdateParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the v1 custom cloud type control plane update params
func (o *V1CustomCloudTypeControlPlaneUpdateParams) SetFileName(fileName runtime.NamedReadCloser) {
	o.FileName = fileName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeControlPlaneUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
