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

// NewV1SpectroClustersVMCloneParams creates a new V1SpectroClustersVMCloneParams object
// with the default values initialized.
func NewV1SpectroClustersVMCloneParams() *V1SpectroClustersVMCloneParams {
	var ()
	return &V1SpectroClustersVMCloneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVMCloneParamsWithTimeout creates a new V1SpectroClustersVMCloneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVMCloneParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVMCloneParams {
	var ()
	return &V1SpectroClustersVMCloneParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersVMCloneParamsWithContext creates a new V1SpectroClustersVMCloneParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVMCloneParamsWithContext(ctx context.Context) *V1SpectroClustersVMCloneParams {
	var ()
	return &V1SpectroClustersVMCloneParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersVMCloneParamsWithHTTPClient creates a new V1SpectroClustersVMCloneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVMCloneParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVMCloneParams {
	var ()
	return &V1SpectroClustersVMCloneParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersVMCloneParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters VM clone operation typically these are written to a http.Request
*/
type V1SpectroClustersVMCloneParams struct {

	/*Body*/
	Body *models.V1SpectroClusterVMCloneEntity
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

// WithTimeout adds the timeout to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVMCloneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) WithContext(ctx context.Context) *V1SpectroClustersVMCloneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVMCloneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) WithBody(body *models.V1SpectroClusterVMCloneEntity) *V1SpectroClustersVMCloneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) SetBody(body *models.V1SpectroClusterVMCloneEntity) {
	o.Body = body
}

// WithNamespace adds the namespace to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) WithNamespace(namespace string) *V1SpectroClustersVMCloneParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUID adds the uid to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) WithUID(uid string) *V1SpectroClustersVMCloneParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) SetUID(uid string) {
	o.UID = uid
}

// WithVMName adds the vMName to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) WithVMName(vMName string) *V1SpectroClustersVMCloneParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the v1 spectro clusters VM clone params
func (o *V1SpectroClustersVMCloneParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVMCloneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
