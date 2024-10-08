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

// NewV1TencentStorageTypesParams creates a new V1TencentStorageTypesParams object
// with the default values initialized.
func NewV1TencentStorageTypesParams() *V1TencentStorageTypesParams {
	var ()
	return &V1TencentStorageTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TencentStorageTypesParamsWithTimeout creates a new V1TencentStorageTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TencentStorageTypesParamsWithTimeout(timeout time.Duration) *V1TencentStorageTypesParams {
	var ()
	return &V1TencentStorageTypesParams{

		timeout: timeout,
	}
}

// NewV1TencentStorageTypesParamsWithContext creates a new V1TencentStorageTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TencentStorageTypesParamsWithContext(ctx context.Context) *V1TencentStorageTypesParams {
	var ()
	return &V1TencentStorageTypesParams{

		Context: ctx,
	}
}

// NewV1TencentStorageTypesParamsWithHTTPClient creates a new V1TencentStorageTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TencentStorageTypesParamsWithHTTPClient(client *http.Client) *V1TencentStorageTypesParams {
	var ()
	return &V1TencentStorageTypesParams{
		HTTPClient: client,
	}
}

/*
V1TencentStorageTypesParams contains all the parameters to send to the API endpoint
for the v1 tencent storage types operation typically these are written to a http.Request
*/
type V1TencentStorageTypesParams struct {

	/*CloudAccountUID
	  Uid for the specific tencent cloud account

	*/
	CloudAccountUID string
	/*Region
	  Region for which tencent storages are listed

	*/
	Region string
	/*Zone
	  Zone for which tencent storages are listed

	*/
	Zone string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) WithTimeout(timeout time.Duration) *V1TencentStorageTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) WithContext(ctx context.Context) *V1TencentStorageTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) WithHTTPClient(client *http.Client) *V1TencentStorageTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) WithCloudAccountUID(cloudAccountUID string) *V1TencentStorageTypesParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithRegion adds the region to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) WithRegion(region string) *V1TencentStorageTypesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) SetRegion(region string) {
	o.Region = region
}

// WithZone adds the zone to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) WithZone(zone string) *V1TencentStorageTypesParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the v1 tencent storage types params
func (o *V1TencentStorageTypesParams) SetZone(zone string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *V1TencentStorageTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	// query param zone
	qrZone := o.Zone
	qZone := qrZone
	if qZone != "" {
		if err := r.SetQueryParam("zone", qZone); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
