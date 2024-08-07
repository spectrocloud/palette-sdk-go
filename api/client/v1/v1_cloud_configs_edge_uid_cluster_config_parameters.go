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

// NewV1CloudConfigsEdgeUIDClusterConfigParams creates a new V1CloudConfigsEdgeUIDClusterConfigParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeUIDClusterConfigParams() *V1CloudConfigsEdgeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeUIDClusterConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeUIDClusterConfigParamsWithTimeout creates a new V1CloudConfigsEdgeUIDClusterConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeUIDClusterConfigParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeUIDClusterConfigParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeUIDClusterConfigParamsWithContext creates a new V1CloudConfigsEdgeUIDClusterConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeUIDClusterConfigParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeUIDClusterConfigParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeUIDClusterConfigParamsWithHTTPClient creates a new V1CloudConfigsEdgeUIDClusterConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeUIDClusterConfigParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeUIDClusterConfigParams {
	var ()
	return &V1CloudConfigsEdgeUIDClusterConfigParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEdgeUIDClusterConfigParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge Uid cluster config operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeUIDClusterConfigParams struct {

	/*Body*/
	Body *models.V1EdgeCloudClusterConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeUIDClusterConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeUIDClusterConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeUIDClusterConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) WithBody(body *models.V1EdgeCloudClusterConfigEntity) *V1CloudConfigsEdgeUIDClusterConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) SetBody(body *models.V1EdgeCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeUIDClusterConfigParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge Uid cluster config params
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeUIDClusterConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
