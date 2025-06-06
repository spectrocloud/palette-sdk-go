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

// NewV1CloudConfigsAksUIDClusterConfigParams creates a new V1CloudConfigsAksUIDClusterConfigParams object
// with the default values initialized.
func NewV1CloudConfigsAksUIDClusterConfigParams() *V1CloudConfigsAksUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAksUIDClusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAksUIDClusterConfigParamsWithTimeout creates a new V1CloudConfigsAksUIDClusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAksUIDClusterConfigParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAksUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAksUIDClusterConfigParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAksUIDClusterConfigParamsWithContext creates a new V1CloudConfigsAksUIDClusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAksUIDClusterConfigParamsWithContext(ctx context.Context) *V1CloudConfigsAksUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAksUIDClusterConfigParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAksUIDClusterConfigParamsWithHTTPClient creates a new V1CloudConfigsAksUIDClusterConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAksUIDClusterConfigParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAksUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAksUIDClusterConfigParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAksUIDClusterConfigParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aks Uid cluster config operation typically these are written to a http.Request
*/
type V1CloudConfigsAksUIDClusterConfigParams struct {

	/*Body*/
	Body *models.V1AzureCloudClusterConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAksUIDClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) WithContext(ctx context.Context) *V1CloudConfigsAksUIDClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAksUIDClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) WithBody(body *models.V1AzureCloudClusterConfigEntity) *V1CloudConfigsAksUIDClusterConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) SetBody(body *models.V1AzureCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) WithConfigUID(configUID string) *V1CloudConfigsAksUIDClusterConfigParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aks Uid cluster config params
func (o *V1CloudConfigsAksUIDClusterConfigParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAksUIDClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
