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

// NewV1CustomCloudTypeClusterTemplateUpdateParams creates a new V1CustomCloudTypeClusterTemplateUpdateParams object
// with the default values initialized.
func NewV1CustomCloudTypeClusterTemplateUpdateParams() *V1CustomCloudTypeClusterTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeClusterTemplateUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeClusterTemplateUpdateParamsWithTimeout creates a new V1CustomCloudTypeClusterTemplateUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeClusterTemplateUpdateParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeClusterTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeClusterTemplateUpdateParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeClusterTemplateUpdateParamsWithContext creates a new V1CustomCloudTypeClusterTemplateUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeClusterTemplateUpdateParamsWithContext(ctx context.Context) *V1CustomCloudTypeClusterTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeClusterTemplateUpdateParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeClusterTemplateUpdateParamsWithHTTPClient creates a new V1CustomCloudTypeClusterTemplateUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeClusterTemplateUpdateParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeClusterTemplateUpdateParams {
	var ()
	return &V1CustomCloudTypeClusterTemplateUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeClusterTemplateUpdateParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type cluster template update operation typically these are written to a http.Request
*/
type V1CustomCloudTypeClusterTemplateUpdateParams struct {

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

// WithTimeout adds the timeout to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeClusterTemplateUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) WithContext(ctx context.Context) *V1CustomCloudTypeClusterTemplateUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeClusterTemplateUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) WithCloudType(cloudType string) *V1CustomCloudTypeClusterTemplateUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithFileName adds the fileName to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) WithFileName(fileName runtime.NamedReadCloser) *V1CustomCloudTypeClusterTemplateUpdateParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the v1 custom cloud type cluster template update params
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) SetFileName(fileName runtime.NamedReadCloser) {
	o.FileName = fileName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeClusterTemplateUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
