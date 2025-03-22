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

// NewV1CloudConfigsGcpUIDClusterConfigParams creates a new V1CloudConfigsGcpUIDClusterConfigParams object
// with the default values initialized.
func NewV1CloudConfigsGcpUIDClusterConfigParams() *V1CloudConfigsGcpUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGcpUIDClusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGcpUIDClusterConfigParamsWithTimeout creates a new V1CloudConfigsGcpUIDClusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGcpUIDClusterConfigParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGcpUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGcpUIDClusterConfigParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsGcpUIDClusterConfigParamsWithContext creates a new V1CloudConfigsGcpUIDClusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGcpUIDClusterConfigParamsWithContext(ctx context.Context) *V1CloudConfigsGcpUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGcpUIDClusterConfigParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsGcpUIDClusterConfigParamsWithHTTPClient creates a new V1CloudConfigsGcpUIDClusterConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGcpUIDClusterConfigParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGcpUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGcpUIDClusterConfigParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsGcpUIDClusterConfigParams contains all the parameters to send to the API endpoint
for the v1 cloud configs gcp Uid cluster config operation typically these are written to a http.Request
*/
type V1CloudConfigsGcpUIDClusterConfigParams struct {

	/*Body*/
	Body *models.V1GcpCloudClusterConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGcpUIDClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) WithContext(ctx context.Context) *V1CloudConfigsGcpUIDClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGcpUIDClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) WithBody(body *models.V1GcpCloudClusterConfigEntity) *V1CloudConfigsGcpUIDClusterConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) SetBody(body *models.V1GcpCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) WithConfigUID(configUID string) *V1CloudConfigsGcpUIDClusterConfigParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs gcp Uid cluster config params
func (o *V1CloudConfigsGcpUIDClusterConfigParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGcpUIDClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
