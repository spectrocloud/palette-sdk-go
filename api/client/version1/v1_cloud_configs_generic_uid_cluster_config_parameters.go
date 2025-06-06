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

// NewV1CloudConfigsGenericUIDClusterConfigParams creates a new V1CloudConfigsGenericUIDClusterConfigParams object
// with the default values initialized.
func NewV1CloudConfigsGenericUIDClusterConfigParams() *V1CloudConfigsGenericUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGenericUIDClusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGenericUIDClusterConfigParamsWithTimeout creates a new V1CloudConfigsGenericUIDClusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGenericUIDClusterConfigParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGenericUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGenericUIDClusterConfigParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsGenericUIDClusterConfigParamsWithContext creates a new V1CloudConfigsGenericUIDClusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGenericUIDClusterConfigParamsWithContext(ctx context.Context) *V1CloudConfigsGenericUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGenericUIDClusterConfigParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsGenericUIDClusterConfigParamsWithHTTPClient creates a new V1CloudConfigsGenericUIDClusterConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGenericUIDClusterConfigParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGenericUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsGenericUIDClusterConfigParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsGenericUIDClusterConfigParams contains all the parameters to send to the API endpoint
for the v1 cloud configs generic Uid cluster config operation typically these are written to a http.Request
*/
type V1CloudConfigsGenericUIDClusterConfigParams struct {

	/*Body*/
	Body *models.V1GenericCloudClusterConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGenericUIDClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) WithContext(ctx context.Context) *V1CloudConfigsGenericUIDClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGenericUIDClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) WithBody(body *models.V1GenericCloudClusterConfigEntity) *V1CloudConfigsGenericUIDClusterConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) SetBody(body *models.V1GenericCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) WithConfigUID(configUID string) *V1CloudConfigsGenericUIDClusterConfigParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs generic Uid cluster config params
func (o *V1CloudConfigsGenericUIDClusterConfigParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGenericUIDClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
