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

// NewV1SpectroClustersVMUpdateParams creates a new V1SpectroClustersVMUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersVMUpdateParams() *V1SpectroClustersVMUpdateParams {
	var ()
	return &V1SpectroClustersVMUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVMUpdateParamsWithTimeout creates a new V1SpectroClustersVMUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVMUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVMUpdateParams {
	var ()
	return &V1SpectroClustersVMUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersVMUpdateParamsWithContext creates a new V1SpectroClustersVMUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVMUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersVMUpdateParams {
	var ()
	return &V1SpectroClustersVMUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersVMUpdateParamsWithHTTPClient creates a new V1SpectroClustersVMUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVMUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVMUpdateParams {
	var ()
	return &V1SpectroClustersVMUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersVMUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters VM update operation typically these are written to a http.Request
*/
type V1SpectroClustersVMUpdateParams struct {

	/*Body*/
	Body *models.V1ClusterVirtualMachine
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

// WithTimeout adds the timeout to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVMUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersVMUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVMUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) WithBody(body *models.V1ClusterVirtualMachine) *V1SpectroClustersVMUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) SetBody(body *models.V1ClusterVirtualMachine) {
	o.Body = body
}

// WithNamespace adds the namespace to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) WithNamespace(namespace string) *V1SpectroClustersVMUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUID adds the uid to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) WithUID(uid string) *V1SpectroClustersVMUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WithVMName adds the vMName to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) WithVMName(vMName string) *V1SpectroClustersVMUpdateParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the v1 spectro clusters VM update params
func (o *V1SpectroClustersVMUpdateParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVMUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
