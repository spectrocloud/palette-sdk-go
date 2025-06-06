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

// NewV1EksLaunchTemplateParams creates a new V1EksLaunchTemplateParams object
// with the default values initialized.
func NewV1EksLaunchTemplateParams() *V1EksLaunchTemplateParams {
	var ()
	return &V1EksLaunchTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EksLaunchTemplateParamsWithTimeout creates a new V1EksLaunchTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EksLaunchTemplateParamsWithTimeout(timeout time.Duration) *V1EksLaunchTemplateParams {
	var ()
	return &V1EksLaunchTemplateParams{

		timeout: timeout,
	}
}

// NewV1EksLaunchTemplateParamsWithContext creates a new V1EksLaunchTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EksLaunchTemplateParamsWithContext(ctx context.Context) *V1EksLaunchTemplateParams {
	var ()
	return &V1EksLaunchTemplateParams{

		Context: ctx,
	}
}

// NewV1EksLaunchTemplateParamsWithHTTPClient creates a new V1EksLaunchTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EksLaunchTemplateParamsWithHTTPClient(client *http.Client) *V1EksLaunchTemplateParams {
	var ()
	return &V1EksLaunchTemplateParams{
		HTTPClient: client,
	}
}

/*
V1EksLaunchTemplateParams contains all the parameters to send to the API endpoint
for the v1 eks launch template operation typically these are written to a http.Request
*/
type V1EksLaunchTemplateParams struct {

	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID string
	/*RegionID
	  AWS region ID for which launch templates are requested

	*/
	RegionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) WithTimeout(timeout time.Duration) *V1EksLaunchTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) WithContext(ctx context.Context) *V1EksLaunchTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) WithHTTPClient(client *http.Client) *V1EksLaunchTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) WithCloudAccountUID(cloudAccountUID string) *V1EksLaunchTemplateParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithRegionID adds the regionID to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) WithRegionID(regionID string) *V1EksLaunchTemplateParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the v1 eks launch template params
func (o *V1EksLaunchTemplateParams) SetRegionID(regionID string) {
	o.RegionID = regionID
}

// WriteToRequest writes these params to a swagger request
func (o *V1EksLaunchTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cloudAccountUid
	qrCloudAccountUID := o.CloudAccountUID
	qCloudAccountUID := qrCloudAccountUID
	if qCloudAccountUID != "" {
		if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
			return err
		}
	}

	// path param regionId
	if err := r.SetPathParam("regionId", o.RegionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
