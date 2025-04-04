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

// NewV1MaasSubnetsGetParams creates a new V1MaasSubnetsGetParams object
// with the default values initialized.
func NewV1MaasSubnetsGetParams() *V1MaasSubnetsGetParams {
	var ()
	return &V1MaasSubnetsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MaasSubnetsGetParamsWithTimeout creates a new V1MaasSubnetsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MaasSubnetsGetParamsWithTimeout(timeout time.Duration) *V1MaasSubnetsGetParams {
	var ()
	return &V1MaasSubnetsGetParams{

		timeout: timeout,
	}
}

// NewV1MaasSubnetsGetParamsWithContext creates a new V1MaasSubnetsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MaasSubnetsGetParamsWithContext(ctx context.Context) *V1MaasSubnetsGetParams {
	var ()
	return &V1MaasSubnetsGetParams{

		Context: ctx,
	}
}

// NewV1MaasSubnetsGetParamsWithHTTPClient creates a new V1MaasSubnetsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MaasSubnetsGetParamsWithHTTPClient(client *http.Client) *V1MaasSubnetsGetParams {
	var ()
	return &V1MaasSubnetsGetParams{
		HTTPClient: client,
	}
}

/*
V1MaasSubnetsGetParams contains all the parameters to send to the API endpoint
for the v1 maas subnets get operation typically these are written to a http.Request
*/
type V1MaasSubnetsGetParams struct {

	/*CloudAccountUID
	  Uid for the specific Maas cloud account

	*/
	CloudAccountUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) WithTimeout(timeout time.Duration) *V1MaasSubnetsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) WithContext(ctx context.Context) *V1MaasSubnetsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) WithHTTPClient(client *http.Client) *V1MaasSubnetsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) WithCloudAccountUID(cloudAccountUID *string) *V1MaasSubnetsGetParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 maas subnets get params
func (o *V1MaasSubnetsGetParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1MaasSubnetsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudAccountUID != nil {

		// query param cloudAccountUid
		var qrCloudAccountUID string
		if o.CloudAccountUID != nil {
			qrCloudAccountUID = *o.CloudAccountUID
		}
		qCloudAccountUID := qrCloudAccountUID
		if qCloudAccountUID != "" {
			if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
