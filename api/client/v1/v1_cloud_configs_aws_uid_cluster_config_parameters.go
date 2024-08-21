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

// NewV1CloudConfigsAwsUIDClusterConfigParams creates a new V1CloudConfigsAwsUIDClusterConfigParams object
// with the default values initialized.
func NewV1CloudConfigsAwsUIDClusterConfigParams() *V1CloudConfigsAwsUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAwsUIDClusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAwsUIDClusterConfigParamsWithTimeout creates a new V1CloudConfigsAwsUIDClusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAwsUIDClusterConfigParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAwsUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAwsUIDClusterConfigParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAwsUIDClusterConfigParamsWithContext creates a new V1CloudConfigsAwsUIDClusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAwsUIDClusterConfigParamsWithContext(ctx context.Context) *V1CloudConfigsAwsUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAwsUIDClusterConfigParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAwsUIDClusterConfigParamsWithHTTPClient creates a new V1CloudConfigsAwsUIDClusterConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAwsUIDClusterConfigParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAwsUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsAwsUIDClusterConfigParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAwsUIDClusterConfigParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aws Uid cluster config operation typically these are written to a http.Request
*/
type V1CloudConfigsAwsUIDClusterConfigParams struct {

	/*Body*/
	Body *models.V1AwsCloudClusterConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAwsUIDClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) WithContext(ctx context.Context) *V1CloudConfigsAwsUIDClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAwsUIDClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) WithBody(body *models.V1AwsCloudClusterConfigEntity) *V1CloudConfigsAwsUIDClusterConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) SetBody(body *models.V1AwsCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) WithConfigUID(configUID string) *V1CloudConfigsAwsUIDClusterConfigParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aws Uid cluster config params
func (o *V1CloudConfigsAwsUIDClusterConfigParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAwsUIDClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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