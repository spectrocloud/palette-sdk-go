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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1CloudConfigsAzureUIDClusterConfigParams creates a new V1CloudConfigsAzureUIDClusterConfigParams object
// with the default values initialized.
func NewV1CloudConfigsAzureUIDClusterConfigParams() *V1CloudConfigsAzureUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAzureUIDClusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAzureUIDClusterConfigParamsWithTimeout creates a new V1CloudConfigsAzureUIDClusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAzureUIDClusterConfigParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAzureUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAzureUIDClusterConfigParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAzureUIDClusterConfigParamsWithContext creates a new V1CloudConfigsAzureUIDClusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAzureUIDClusterConfigParamsWithContext(ctx context.Context) *V1CloudConfigsAzureUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAzureUIDClusterConfigParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAzureUIDClusterConfigParamsWithHTTPClient creates a new V1CloudConfigsAzureUIDClusterConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAzureUIDClusterConfigParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAzureUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAzureUIDClusterConfigParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAzureUIDClusterConfigParams contains all the parameters to send to the API endpoint
for the v1 cloud configs azure Uid cluster config operation typically these are written to a http.Request
*/
type V1CloudConfigsAzureUIDClusterConfigParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAzureUIDClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) WithContext(ctx context.Context) *V1CloudConfigsAzureUIDClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAzureUIDClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) WithBody(body *models.V1AzureCloudClusterConfigEntity) *V1CloudConfigsAzureUIDClusterConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) SetBody(body *models.V1AzureCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) WithConfigUID(configUID string) *V1CloudConfigsAzureUIDClusterConfigParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs azure Uid cluster config params
func (o *V1CloudConfigsAzureUIDClusterConfigParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAzureUIDClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
