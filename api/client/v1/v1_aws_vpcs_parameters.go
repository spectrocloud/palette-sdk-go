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

// NewV1AwsVpcsParams creates a new V1AwsVpcsParams object
// with the default values initialized.
func NewV1AwsVpcsParams() *V1AwsVpcsParams {
	var ()
	return &V1AwsVpcsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsVpcsParamsWithTimeout creates a new V1AwsVpcsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsVpcsParamsWithTimeout(timeout time.Duration) *V1AwsVpcsParams {
	var ()
	return &V1AwsVpcsParams{

		timeout: timeout,
	}
}

// NewV1AwsVpcsParamsWithContext creates a new V1AwsVpcsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsVpcsParamsWithContext(ctx context.Context) *V1AwsVpcsParams {
	var ()
	return &V1AwsVpcsParams{

		Context: ctx,
	}
}

// NewV1AwsVpcsParamsWithHTTPClient creates a new V1AwsVpcsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsVpcsParamsWithHTTPClient(client *http.Client) *V1AwsVpcsParams {
	var ()
	return &V1AwsVpcsParams{
		HTTPClient: client,
	}
}

/*
V1AwsVpcsParams contains all the parameters to send to the API endpoint
for the v1 aws vpcs operation typically these are written to a http.Request
*/
type V1AwsVpcsParams struct {

	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID string
	/*Region
	  Region for which VPCs are requested

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws vpcs params
func (o *V1AwsVpcsParams) WithTimeout(timeout time.Duration) *V1AwsVpcsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws vpcs params
func (o *V1AwsVpcsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws vpcs params
func (o *V1AwsVpcsParams) WithContext(ctx context.Context) *V1AwsVpcsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws vpcs params
func (o *V1AwsVpcsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws vpcs params
func (o *V1AwsVpcsParams) WithHTTPClient(client *http.Client) *V1AwsVpcsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws vpcs params
func (o *V1AwsVpcsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 aws vpcs params
func (o *V1AwsVpcsParams) WithCloudAccountUID(cloudAccountUID string) *V1AwsVpcsParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 aws vpcs params
func (o *V1AwsVpcsParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithRegion adds the region to the v1 aws vpcs params
func (o *V1AwsVpcsParams) WithRegion(region string) *V1AwsVpcsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 aws vpcs params
func (o *V1AwsVpcsParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsVpcsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}