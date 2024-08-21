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

// NewV1CloudConfigsVspherePoolMachinesUIDDeleteParams creates a new V1CloudConfigsVspherePoolMachinesUIDDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsVspherePoolMachinesUIDDeleteParams() *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDDeleteParamsWithTimeout creates a new V1CloudConfigsVspherePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsVspherePoolMachinesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDDeleteParamsWithContext creates a new V1CloudConfigsVspherePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsVspherePoolMachinesUIDDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDDeleteParamsWithHTTPClient creates a new V1CloudConfigsVspherePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsVspherePoolMachinesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsVspherePoolMachinesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs vsphere pool machines Uid delete operation typically these are written to a http.Request
*/
type V1CloudConfigsVspherePoolMachinesUIDDeleteParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string
	/*MachinePoolName
	  Machine pool name

	*/
	MachinePoolName string
	/*MachineUID
	  Machine uid

	*/
	MachineUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) WithMachineUID(machineUID string) *V1CloudConfigsVspherePoolMachinesUIDDeleteParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs vsphere pool machines Uid delete params
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsVspherePoolMachinesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param machineUid
	if err := r.SetPathParam("machineUid", o.MachineUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}