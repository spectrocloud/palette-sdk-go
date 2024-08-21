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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1AwsCopyImageFromDefaultRegionParams creates a new V1AwsCopyImageFromDefaultRegionParams object
// with the default values initialized.
func NewV1AwsCopyImageFromDefaultRegionParams() *V1AwsCopyImageFromDefaultRegionParams {
	var ()
	return &V1AwsCopyImageFromDefaultRegionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsCopyImageFromDefaultRegionParamsWithTimeout creates a new V1AwsCopyImageFromDefaultRegionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsCopyImageFromDefaultRegionParamsWithTimeout(timeout time.Duration) *V1AwsCopyImageFromDefaultRegionParams {
	var ()
	return &V1AwsCopyImageFromDefaultRegionParams{

		timeout: timeout,
	}
}

// NewV1AwsCopyImageFromDefaultRegionParamsWithContext creates a new V1AwsCopyImageFromDefaultRegionParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsCopyImageFromDefaultRegionParamsWithContext(ctx context.Context) *V1AwsCopyImageFromDefaultRegionParams {
	var ()
	return &V1AwsCopyImageFromDefaultRegionParams{

		Context: ctx,
	}
}

// NewV1AwsCopyImageFromDefaultRegionParamsWithHTTPClient creates a new V1AwsCopyImageFromDefaultRegionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsCopyImageFromDefaultRegionParamsWithHTTPClient(client *http.Client) *V1AwsCopyImageFromDefaultRegionParams {
	var ()
	return &V1AwsCopyImageFromDefaultRegionParams{
		HTTPClient: client,
	}
}

/*
V1AwsCopyImageFromDefaultRegionParams contains all the parameters to send to the API endpoint
for the v1 aws copy image from default region operation typically these are written to a http.Request
*/
type V1AwsCopyImageFromDefaultRegionParams struct {

	/*Region
	  Region to copy AWS image from

	*/
	Region string
	/*SpectroClusterAwsImageTag
	  Request payload to copy the AWS image

	*/
	SpectroClusterAwsImageTag *models.V1AwsFindImageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) WithTimeout(timeout time.Duration) *V1AwsCopyImageFromDefaultRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) WithContext(ctx context.Context) *V1AwsCopyImageFromDefaultRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) WithHTTPClient(client *http.Client) *V1AwsCopyImageFromDefaultRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) WithRegion(region string) *V1AwsCopyImageFromDefaultRegionParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) SetRegion(region string) {
	o.Region = region
}

// WithSpectroClusterAwsImageTag adds the spectroClusterAwsImageTag to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) WithSpectroClusterAwsImageTag(spectroClusterAwsImageTag *models.V1AwsFindImageRequest) *V1AwsCopyImageFromDefaultRegionParams {
	o.SetSpectroClusterAwsImageTag(spectroClusterAwsImageTag)
	return o
}

// SetSpectroClusterAwsImageTag adds the spectroClusterAwsImageTag to the v1 aws copy image from default region params
func (o *V1AwsCopyImageFromDefaultRegionParams) SetSpectroClusterAwsImageTag(spectroClusterAwsImageTag *models.V1AwsFindImageRequest) {
	o.SpectroClusterAwsImageTag = spectroClusterAwsImageTag
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsCopyImageFromDefaultRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if o.SpectroClusterAwsImageTag != nil {
		if err := r.SetBodyParam(o.SpectroClusterAwsImageTag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}