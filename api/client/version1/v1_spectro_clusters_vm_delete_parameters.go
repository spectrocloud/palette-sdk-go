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
)

// NewV1SpectroClustersVMDeleteParams creates a new V1SpectroClustersVMDeleteParams object
// with the default values initialized.
func NewV1SpectroClustersVMDeleteParams() *V1SpectroClustersVMDeleteParams {
	var ()
	return &V1SpectroClustersVMDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVMDeleteParamsWithTimeout creates a new V1SpectroClustersVMDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVMDeleteParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVMDeleteParams {
	var ()
	return &V1SpectroClustersVMDeleteParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersVMDeleteParamsWithContext creates a new V1SpectroClustersVMDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVMDeleteParamsWithContext(ctx context.Context) *V1SpectroClustersVMDeleteParams {
	var ()
	return &V1SpectroClustersVMDeleteParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersVMDeleteParamsWithHTTPClient creates a new V1SpectroClustersVMDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVMDeleteParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVMDeleteParams {
	var ()
	return &V1SpectroClustersVMDeleteParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersVMDeleteParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters VM delete operation typically these are written to a http.Request
*/
type V1SpectroClustersVMDeleteParams struct {

	/*Namespace
	  Namespace name

	*/
	Namespace string
	/*UID
	  Cluster uid

	*/
	UID string
	/*VMName
	  Virtual Machine name

	*/
	VMName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVMDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) WithContext(ctx context.Context) *V1SpectroClustersVMDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVMDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) WithNamespace(namespace string) *V1SpectroClustersVMDeleteParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUID adds the uid to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) WithUID(uid string) *V1SpectroClustersVMDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WithVMName adds the vMName to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) WithVMName(vMName string) *V1SpectroClustersVMDeleteParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the v1 spectro clusters VM delete params
func (o *V1SpectroClustersVMDeleteParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVMDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {
		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	// path param vmName
	if err := r.SetPathParam("vmName", o.VMName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
