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

// NewV1MaasZonesGetParams creates a new V1MaasZonesGetParams object
// with the default values initialized.
func NewV1MaasZonesGetParams() *V1MaasZonesGetParams {
	var ()
	return &V1MaasZonesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1MaasZonesGetParamsWithTimeout creates a new V1MaasZonesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1MaasZonesGetParamsWithTimeout(timeout time.Duration) *V1MaasZonesGetParams {
	var ()
	return &V1MaasZonesGetParams{

		timeout: timeout,
	}
}

// NewV1MaasZonesGetParamsWithContext creates a new V1MaasZonesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1MaasZonesGetParamsWithContext(ctx context.Context) *V1MaasZonesGetParams {
	var ()
	return &V1MaasZonesGetParams{

		Context: ctx,
	}
}

// NewV1MaasZonesGetParamsWithHTTPClient creates a new V1MaasZonesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1MaasZonesGetParamsWithHTTPClient(client *http.Client) *V1MaasZonesGetParams {
	var ()
	return &V1MaasZonesGetParams{
		HTTPClient: client,
	}
}

/*
V1MaasZonesGetParams contains all the parameters to send to the API endpoint
for the v1 maas zones get operation typically these are written to a http.Request
*/
type V1MaasZonesGetParams struct {

	/*CloudAccountUID
	  Uid for the specific Maas cloud account

	*/
	CloudAccountUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 maas zones get params
func (o *V1MaasZonesGetParams) WithTimeout(timeout time.Duration) *V1MaasZonesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 maas zones get params
func (o *V1MaasZonesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 maas zones get params
func (o *V1MaasZonesGetParams) WithContext(ctx context.Context) *V1MaasZonesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 maas zones get params
func (o *V1MaasZonesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 maas zones get params
func (o *V1MaasZonesGetParams) WithHTTPClient(client *http.Client) *V1MaasZonesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 maas zones get params
func (o *V1MaasZonesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 maas zones get params
func (o *V1MaasZonesGetParams) WithCloudAccountUID(cloudAccountUID *string) *V1MaasZonesGetParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 maas zones get params
func (o *V1MaasZonesGetParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1MaasZonesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
