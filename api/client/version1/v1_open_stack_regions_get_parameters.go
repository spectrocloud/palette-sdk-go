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

// NewV1OpenStackRegionsGetParams creates a new V1OpenStackRegionsGetParams object
// with the default values initialized.
func NewV1OpenStackRegionsGetParams() *V1OpenStackRegionsGetParams {
	var ()
	return &V1OpenStackRegionsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OpenStackRegionsGetParamsWithTimeout creates a new V1OpenStackRegionsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OpenStackRegionsGetParamsWithTimeout(timeout time.Duration) *V1OpenStackRegionsGetParams {
	var ()
	return &V1OpenStackRegionsGetParams{

		timeout: timeout,
	}
}

// NewV1OpenStackRegionsGetParamsWithContext creates a new V1OpenStackRegionsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OpenStackRegionsGetParamsWithContext(ctx context.Context) *V1OpenStackRegionsGetParams {
	var ()
	return &V1OpenStackRegionsGetParams{

		Context: ctx,
	}
}

// NewV1OpenStackRegionsGetParamsWithHTTPClient creates a new V1OpenStackRegionsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OpenStackRegionsGetParamsWithHTTPClient(client *http.Client) *V1OpenStackRegionsGetParams {
	var ()
	return &V1OpenStackRegionsGetParams{
		HTTPClient: client,
	}
}

/*
V1OpenStackRegionsGetParams contains all the parameters to send to the API endpoint
for the v1 open stack regions get operation typically these are written to a http.Request
*/
type V1OpenStackRegionsGetParams struct {

	/*CloudAccountUID
	  Uid for the specific OpenStack cloud account

	*/
	CloudAccountUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) WithTimeout(timeout time.Duration) *V1OpenStackRegionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) WithContext(ctx context.Context) *V1OpenStackRegionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) WithHTTPClient(client *http.Client) *V1OpenStackRegionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) WithCloudAccountUID(cloudAccountUID *string) *V1OpenStackRegionsGetParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 open stack regions get params
func (o *V1OpenStackRegionsGetParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1OpenStackRegionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
