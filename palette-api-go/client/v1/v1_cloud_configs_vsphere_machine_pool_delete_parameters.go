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

// NewV1CloudConfigsVsphereMachinePoolDeleteParams creates a new V1CloudConfigsVsphereMachinePoolDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsVsphereMachinePoolDeleteParams() *V1CloudConfigsVsphereMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsVsphereMachinePoolDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithTimeout creates a new V1CloudConfigsVsphereMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsVsphereMachinePoolDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithContext creates a new V1CloudConfigsVsphereMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsVsphereMachinePoolDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithHTTPClient creates a new V1CloudConfigsVsphereMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsVsphereMachinePoolDeleteParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsVsphereMachinePoolDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs vsphere machine pool delete operation typically these are written to a http.Request
*/
type V1CloudConfigsVsphereMachinePoolDeleteParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string
	/*MachinePoolName
	  Machine pool name

	*/
	MachinePoolName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsVsphereMachinePoolDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere machine pool delete params
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsVsphereMachinePoolDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configUid
	if err := r.SetPathParam("configUid", o.ConfigUID); err != nil {
		return err
	}

	// path param machinePoolName
	if err := r.SetPathParam("machinePoolName", o.MachinePoolName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
