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

// NewV1TencentKeypairsParams creates a new V1TencentKeypairsParams object
// with the default values initialized.
func NewV1TencentKeypairsParams() *V1TencentKeypairsParams {
	var ()
	return &V1TencentKeypairsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TencentKeypairsParamsWithTimeout creates a new V1TencentKeypairsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TencentKeypairsParamsWithTimeout(timeout time.Duration) *V1TencentKeypairsParams {
	var ()
	return &V1TencentKeypairsParams{

		timeout: timeout,
	}
}

// NewV1TencentKeypairsParamsWithContext creates a new V1TencentKeypairsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TencentKeypairsParamsWithContext(ctx context.Context) *V1TencentKeypairsParams {
	var ()
	return &V1TencentKeypairsParams{

		Context: ctx,
	}
}

// NewV1TencentKeypairsParamsWithHTTPClient creates a new V1TencentKeypairsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TencentKeypairsParamsWithHTTPClient(client *http.Client) *V1TencentKeypairsParams {
	var ()
	return &V1TencentKeypairsParams{
		HTTPClient: client,
	}
}

/*
V1TencentKeypairsParams contains all the parameters to send to the API endpoint
for the v1 tencent keypairs operation typically these are written to a http.Request
*/
type V1TencentKeypairsParams struct {

	/*CloudAccountUID
	  Uid for the specific Tencent cloud account

	*/
	CloudAccountUID string
	/*Region
	  Region for which keypairs are requested

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) WithTimeout(timeout time.Duration) *V1TencentKeypairsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) WithContext(ctx context.Context) *V1TencentKeypairsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) WithHTTPClient(client *http.Client) *V1TencentKeypairsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) WithCloudAccountUID(cloudAccountUID string) *V1TencentKeypairsParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithRegion adds the region to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) WithRegion(region string) *V1TencentKeypairsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 tencent keypairs params
func (o *V1TencentKeypairsParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1TencentKeypairsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
