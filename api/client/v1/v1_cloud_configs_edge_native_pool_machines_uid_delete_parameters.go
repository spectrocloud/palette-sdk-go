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

// NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams creates a new V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams() *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithTimeout creates a new V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithContext creates a new V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithHTTPClient creates a new V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge native pool machines Uid delete operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) WithMachineUID(machineUID string) *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs edge native pool machines Uid delete params
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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