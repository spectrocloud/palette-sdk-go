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
)

// NewV1VsphereComputeClusterResourcesParams creates a new V1VsphereComputeClusterResourcesParams object
// with the default values initialized.
func NewV1VsphereComputeClusterResourcesParams() *V1VsphereComputeClusterResourcesParams {
	var ()
	return &V1VsphereComputeClusterResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VsphereComputeClusterResourcesParamsWithTimeout creates a new V1VsphereComputeClusterResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VsphereComputeClusterResourcesParamsWithTimeout(timeout time.Duration) *V1VsphereComputeClusterResourcesParams {
	var ()
	return &V1VsphereComputeClusterResourcesParams{

		timeout: timeout,
	}
}

// NewV1VsphereComputeClusterResourcesParamsWithContext creates a new V1VsphereComputeClusterResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VsphereComputeClusterResourcesParamsWithContext(ctx context.Context) *V1VsphereComputeClusterResourcesParams {
	var ()
	return &V1VsphereComputeClusterResourcesParams{

		Context: ctx,
	}
}

// NewV1VsphereComputeClusterResourcesParamsWithHTTPClient creates a new V1VsphereComputeClusterResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VsphereComputeClusterResourcesParamsWithHTTPClient(client *http.Client) *V1VsphereComputeClusterResourcesParams {
	var ()
	return &V1VsphereComputeClusterResourcesParams{
		HTTPClient: client,
	}
}

/*
V1VsphereComputeClusterResourcesParams contains all the parameters to send to the API endpoint
for the v1 vsphere compute cluster resources operation typically these are written to a http.Request
*/
type V1VsphereComputeClusterResourcesParams struct {

	/*CloudAccountUID
	  Uid for the specific VSphere cloud account

	*/
	CloudAccountUID string
	/*Computecluster
	  computecluster for which resources is requested

	*/
	Computecluster string
	/*UID
	  VSphere datacenter uid for which resources is requested

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) WithTimeout(timeout time.Duration) *V1VsphereComputeClusterResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) WithContext(ctx context.Context) *V1VsphereComputeClusterResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) WithHTTPClient(client *http.Client) *V1VsphereComputeClusterResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) WithCloudAccountUID(cloudAccountUID string) *V1VsphereComputeClusterResourcesParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithComputecluster adds the computecluster to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) WithComputecluster(computecluster string) *V1VsphereComputeClusterResourcesParams {
	o.SetComputecluster(computecluster)
	return o
}

// SetComputecluster adds the computecluster to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) SetComputecluster(computecluster string) {
	o.Computecluster = computecluster
}

// WithUID adds the uid to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) WithUID(uid string) *V1VsphereComputeClusterResourcesParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 vsphere compute cluster resources params
func (o *V1VsphereComputeClusterResourcesParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1VsphereComputeClusterResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cloudAccountUid
	qrCloudAccountUID := o.CloudAccountUID
	qCloudAccountUID := qrCloudAccountUID
	if qCloudAccountUID != "" {
		if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
			return err
		}
	}

	// path param computecluster
	if err := r.SetPathParam("computecluster", o.Computecluster); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
