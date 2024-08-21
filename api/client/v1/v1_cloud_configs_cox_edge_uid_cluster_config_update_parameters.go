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

// NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams creates a new V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams() *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParamsWithTimeout creates a new V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParamsWithContext creates a new V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParamsWithHTTPClient creates a new V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsCoxEdgeUIDClusterConfigUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs cox edge Uid cluster config update operation typically these are written to a http.Request
*/
type V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams struct {

	/*Body*/
	Body *models.V1CoxEdgeCloudClusterConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) WithBody(body *models.V1CoxEdgeCloudClusterConfigEntity) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) SetBody(body *models.V1CoxEdgeCloudClusterConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs cox edge Uid cluster config update params
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsCoxEdgeUIDClusterConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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