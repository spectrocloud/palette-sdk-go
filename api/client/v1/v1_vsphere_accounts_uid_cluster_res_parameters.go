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
	"github.com/go-openapi/swag"
)

// NewV1VsphereAccountsUIDClusterResParams creates a new V1VsphereAccountsUIDClusterResParams object
// with the default values initialized.
func NewV1VsphereAccountsUIDClusterResParams() *V1VsphereAccountsUIDClusterResParams {
	var ()
	return &V1VsphereAccountsUIDClusterResParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VsphereAccountsUIDClusterResParamsWithTimeout creates a new V1VsphereAccountsUIDClusterResParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VsphereAccountsUIDClusterResParamsWithTimeout(timeout time.Duration) *V1VsphereAccountsUIDClusterResParams {
	var ()
	return &V1VsphereAccountsUIDClusterResParams{

		timeout: timeout,
	}
}

// NewV1VsphereAccountsUIDClusterResParamsWithContext creates a new V1VsphereAccountsUIDClusterResParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VsphereAccountsUIDClusterResParamsWithContext(ctx context.Context) *V1VsphereAccountsUIDClusterResParams {
	var ()
	return &V1VsphereAccountsUIDClusterResParams{

		Context: ctx,
	}
}

// NewV1VsphereAccountsUIDClusterResParamsWithHTTPClient creates a new V1VsphereAccountsUIDClusterResParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VsphereAccountsUIDClusterResParamsWithHTTPClient(client *http.Client) *V1VsphereAccountsUIDClusterResParams {
	var ()
	return &V1VsphereAccountsUIDClusterResParams{
		HTTPClient: client,
	}
}

/*
V1VsphereAccountsUIDClusterResParams contains all the parameters to send to the API endpoint
for the v1 vsphere accounts Uid cluster res operation typically these are written to a http.Request
*/
type V1VsphereAccountsUIDClusterResParams struct {

	/*Computecluster*/
	Computecluster string
	/*Datacenter*/
	Datacenter string
	/*UID*/
	UID string
	/*UseQualifiedNetworkName*/
	UseQualifiedNetworkName *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) WithTimeout(timeout time.Duration) *V1VsphereAccountsUIDClusterResParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) WithContext(ctx context.Context) *V1VsphereAccountsUIDClusterResParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) WithHTTPClient(client *http.Client) *V1VsphereAccountsUIDClusterResParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComputecluster adds the computecluster to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) WithComputecluster(computecluster string) *V1VsphereAccountsUIDClusterResParams {
	o.SetComputecluster(computecluster)
	return o
}

// SetComputecluster adds the computecluster to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) SetComputecluster(computecluster string) {
	o.Computecluster = computecluster
}

// WithDatacenter adds the datacenter to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) WithDatacenter(datacenter string) *V1VsphereAccountsUIDClusterResParams {
	o.SetDatacenter(datacenter)
	return o
}

// SetDatacenter adds the datacenter to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) SetDatacenter(datacenter string) {
	o.Datacenter = datacenter
}

// WithUID adds the uid to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) WithUID(uid string) *V1VsphereAccountsUIDClusterResParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) SetUID(uid string) {
	o.UID = uid
}

// WithUseQualifiedNetworkName adds the useQualifiedNetworkName to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) WithUseQualifiedNetworkName(useQualifiedNetworkName *bool) *V1VsphereAccountsUIDClusterResParams {
	o.SetUseQualifiedNetworkName(useQualifiedNetworkName)
	return o
}

// SetUseQualifiedNetworkName adds the useQualifiedNetworkName to the v1 vsphere accounts Uid cluster res params
func (o *V1VsphereAccountsUIDClusterResParams) SetUseQualifiedNetworkName(useQualifiedNetworkName *bool) {
	o.UseQualifiedNetworkName = useQualifiedNetworkName
}

// WriteToRequest writes these params to a swagger request
func (o *V1VsphereAccountsUIDClusterResParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param computecluster
	qrComputecluster := o.Computecluster
	qComputecluster := qrComputecluster
	if qComputecluster != "" {
		if err := r.SetQueryParam("computecluster", qComputecluster); err != nil {
			return err
		}
	}

	// query param datacenter
	qrDatacenter := o.Datacenter
	qDatacenter := qrDatacenter
	if qDatacenter != "" {
		if err := r.SetQueryParam("datacenter", qDatacenter); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if o.UseQualifiedNetworkName != nil {

		// query param useQualifiedNetworkName
		var qrUseQualifiedNetworkName bool
		if o.UseQualifiedNetworkName != nil {
			qrUseQualifiedNetworkName = *o.UseQualifiedNetworkName
		}
		qUseQualifiedNetworkName := swag.FormatBool(qrUseQualifiedNetworkName)
		if qUseQualifiedNetworkName != "" {
			if err := r.SetQueryParam("useQualifiedNetworkName", qUseQualifiedNetworkName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}