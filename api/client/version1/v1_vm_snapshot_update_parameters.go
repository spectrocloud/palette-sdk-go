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

// NewV1VMSnapshotUpdateParams creates a new V1VMSnapshotUpdateParams object
// with the default values initialized.
func NewV1VMSnapshotUpdateParams() *V1VMSnapshotUpdateParams {
	var ()
	return &V1VMSnapshotUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VMSnapshotUpdateParamsWithTimeout creates a new V1VMSnapshotUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VMSnapshotUpdateParamsWithTimeout(timeout time.Duration) *V1VMSnapshotUpdateParams {
	var ()
	return &V1VMSnapshotUpdateParams{

		timeout: timeout,
	}
}

// NewV1VMSnapshotUpdateParamsWithContext creates a new V1VMSnapshotUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VMSnapshotUpdateParamsWithContext(ctx context.Context) *V1VMSnapshotUpdateParams {
	var ()
	return &V1VMSnapshotUpdateParams{

		Context: ctx,
	}
}

// NewV1VMSnapshotUpdateParamsWithHTTPClient creates a new V1VMSnapshotUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VMSnapshotUpdateParamsWithHTTPClient(client *http.Client) *V1VMSnapshotUpdateParams {
	var ()
	return &V1VMSnapshotUpdateParams{
		HTTPClient: client,
	}
}

/*
V1VMSnapshotUpdateParams contains all the parameters to send to the API endpoint
for the v1 VM snapshot update operation typically these are written to a http.Request
*/
type V1VMSnapshotUpdateParams struct {

	/*Body*/
	Body *models.V1VirtualMachineSnapshot
	/*Namespace
	  Namespace name

	*/
	Namespace string
	/*SnapshotName
	  Snapshot name

	*/
	SnapshotName string
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

// WithTimeout adds the timeout to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithTimeout(timeout time.Duration) *V1VMSnapshotUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithContext(ctx context.Context) *V1VMSnapshotUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithHTTPClient(client *http.Client) *V1VMSnapshotUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithBody(body *models.V1VirtualMachineSnapshot) *V1VMSnapshotUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetBody(body *models.V1VirtualMachineSnapshot) {
	o.Body = body
}

// WithNamespace adds the namespace to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithNamespace(namespace string) *V1VMSnapshotUpdateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSnapshotName adds the snapshotName to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithSnapshotName(snapshotName string) *V1VMSnapshotUpdateParams {
	o.SetSnapshotName(snapshotName)
	return o
}

// SetSnapshotName adds the snapshotName to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetSnapshotName(snapshotName string) {
	o.SnapshotName = snapshotName
}

// WithUID adds the uid to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithUID(uid string) *V1VMSnapshotUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WithVMName adds the vMName to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) WithVMName(vMName string) *V1VMSnapshotUpdateParams {
	o.SetVMName(vMName)
	return o
}

// SetVMName adds the vmName to the v1 VM snapshot update params
func (o *V1VMSnapshotUpdateParams) SetVMName(vMName string) {
	o.VMName = vMName
}

// WriteToRequest writes these params to a swagger request
func (o *V1VMSnapshotUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param snapshotName
	if err := r.SetPathParam("snapshotName", o.SnapshotName); err != nil {
		return err
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
