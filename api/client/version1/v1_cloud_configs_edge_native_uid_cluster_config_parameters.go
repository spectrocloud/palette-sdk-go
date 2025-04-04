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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1CloudConfigsEdgeNativeUIDClusterConfigParams creates a new V1CloudConfigsEdgeNativeUIDClusterConfigParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeNativeUIDClusterConfigParams() *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeNativeUIDClusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeNativeUIDClusterConfigParamsWithTimeout creates a new V1CloudConfigsEdgeNativeUIDClusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeNativeUIDClusterConfigParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeNativeUIDClusterConfigParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeNativeUIDClusterConfigParamsWithContext creates a new V1CloudConfigsEdgeNativeUIDClusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeNativeUIDClusterConfigParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeNativeUIDClusterConfigParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeNativeUIDClusterConfigParamsWithHTTPClient creates a new V1CloudConfigsEdgeNativeUIDClusterConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeNativeUIDClusterConfigParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeNativeUIDClusterConfigParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEdgeNativeUIDClusterConfigParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge native Uid cluster config operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeNativeUIDClusterConfigParams struct {

	/*Body*/
	Body *models.V1EdgeNativeCloudClusterConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) WithBody(body *models.V1EdgeNativeCloudClusterConfigEntity) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) SetBody(body *models.V1EdgeNativeCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeNativeUIDClusterConfigParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge native Uid cluster config params
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeNativeUIDClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param configUid
	if err := r.SetPathParam("configUid", o.ConfigUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
