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

// NewV1TencentZonesParams creates a new V1TencentZonesParams object
// with the default values initialized.
func NewV1TencentZonesParams() *V1TencentZonesParams {
	var ()
	return &V1TencentZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TencentZonesParamsWithTimeout creates a new V1TencentZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TencentZonesParamsWithTimeout(timeout time.Duration) *V1TencentZonesParams {
	var ()
	return &V1TencentZonesParams{

		timeout: timeout,
	}
}

// NewV1TencentZonesParamsWithContext creates a new V1TencentZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TencentZonesParamsWithContext(ctx context.Context) *V1TencentZonesParams {
	var ()
	return &V1TencentZonesParams{

		Context: ctx,
	}
}

// NewV1TencentZonesParamsWithHTTPClient creates a new V1TencentZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TencentZonesParamsWithHTTPClient(client *http.Client) *V1TencentZonesParams {
	var ()
	return &V1TencentZonesParams{
		HTTPClient: client,
	}
}

/*V1TencentZonesParams contains all the parameters to send to the API endpoint
for the v1 tencent zones operation typically these are written to a http.Request
*/
type V1TencentZonesParams struct {

	/*CloudAccountUID
	  Uid for the specific Tencent cloud account

	*/
	CloudAccountUID string
	/*Region
	  Region for which zones are requested

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tencent zones params
func (o *V1TencentZonesParams) WithTimeout(timeout time.Duration) *V1TencentZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tencent zones params
func (o *V1TencentZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tencent zones params
func (o *V1TencentZonesParams) WithContext(ctx context.Context) *V1TencentZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tencent zones params
func (o *V1TencentZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tencent zones params
func (o *V1TencentZonesParams) WithHTTPClient(client *http.Client) *V1TencentZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tencent zones params
func (o *V1TencentZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 tencent zones params
func (o *V1TencentZonesParams) WithCloudAccountUID(cloudAccountUID string) *V1TencentZonesParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 tencent zones params
func (o *V1TencentZonesParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithRegion adds the region to the v1 tencent zones params
func (o *V1TencentZonesParams) WithRegion(region string) *V1TencentZonesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 tencent zones params
func (o *V1TencentZonesParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1TencentZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
