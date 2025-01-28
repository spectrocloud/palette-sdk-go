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

// NewV1SpectroClustersVMRestartParams creates a new V1SpectroClustersVMRestartParams object
// with the default values initialized.
func NewV1SpectroClustersVMRestartParams() *V1SpectroClustersVMRestartParams {
	var ()
	return &V1SpectroClustersVMRestartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVMRestartParamsWithTimeout creates a new V1SpectroClustersVMRestartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVMRestartParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVMRestartParams {
	var ()
	return &V1SpectroClustersVMRestartParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersVMRestartParamsWithContext creates a new V1SpectroClustersVMRestartParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVMRestartParamsWithContext(ctx context.Context) *V1SpectroClustersVMRestartParams {
	var ()
	return &V1SpectroClustersVMRestartParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersVMRestartParamsWithHTTPClient creates a new V1SpectroClustersVMRestartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVMRestartParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVMRestartParams {
	var ()
	return &V1SpectroClustersVMRestartParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersVMRestartParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters VM restart operation typically these are written to a http.Request
*/
type V1SpectroClustersVMRestartParams struct {

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

// WithTimeout adds the timeout to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVMRestartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) WithContext(ctx context.Context) *V1SpectroClustersVMRestartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVMRestartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) WithNamespace(namespace string) *V1SpectroClustersVMRestartParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUID adds the uid to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) WithUID(uid string) *V1SpectroClustersVMRestartParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) SetUID(uid string) {
	o.UID = uid
}

// WithVMName adds the vMName to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) WithVMName(vMName string) *V1SpectroClustersVMRestartParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the v1 spectro clusters VM restart params
func (o *V1SpectroClustersVMRestartParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVMRestartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
