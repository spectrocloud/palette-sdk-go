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

// NewV1SpectroClustersVMResumeParams creates a new V1SpectroClustersVMResumeParams object
// with the default values initialized.
func NewV1SpectroClustersVMResumeParams() *V1SpectroClustersVMResumeParams {
	var ()
	return &V1SpectroClustersVMResumeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVMResumeParamsWithTimeout creates a new V1SpectroClustersVMResumeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVMResumeParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVMResumeParams {
	var ()
	return &V1SpectroClustersVMResumeParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersVMResumeParamsWithContext creates a new V1SpectroClustersVMResumeParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVMResumeParamsWithContext(ctx context.Context) *V1SpectroClustersVMResumeParams {
	var ()
	return &V1SpectroClustersVMResumeParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersVMResumeParamsWithHTTPClient creates a new V1SpectroClustersVMResumeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVMResumeParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVMResumeParams {
	var ()
	return &V1SpectroClustersVMResumeParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersVMResumeParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters VM resume operation typically these are written to a http.Request
*/
type V1SpectroClustersVMResumeParams struct {

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

// WithTimeout adds the timeout to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVMResumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) WithContext(ctx context.Context) *V1SpectroClustersVMResumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVMResumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) WithNamespace(namespace string) *V1SpectroClustersVMResumeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUID adds the uid to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) WithUID(uid string) *V1SpectroClustersVMResumeParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) SetUID(uid string) {
	o.UID = uid
}

// WithVMName adds the vMName to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) WithVMName(vMName string) *V1SpectroClustersVMResumeParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the v1 spectro clusters VM resume params
func (o *V1SpectroClustersVMResumeParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVMResumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
