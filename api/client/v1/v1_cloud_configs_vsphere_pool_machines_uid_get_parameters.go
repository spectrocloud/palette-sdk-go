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

// NewV1CloudConfigsVspherePoolMachinesUIDGetParams creates a new V1CloudConfigsVspherePoolMachinesUIDGetParams object
// with the default values initialized.
func NewV1CloudConfigsVspherePoolMachinesUIDGetParams() *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithTimeout creates a new V1CloudConfigsVspherePoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext creates a new V1CloudConfigsVspherePoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(ctx context.Context) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithHTTPClient creates a new V1CloudConfigsVspherePoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDGetParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsVspherePoolMachinesUIDGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs vsphere pool machines Uid get operation typically these are written to a http.Request
*/
type V1CloudConfigsVspherePoolMachinesUIDGetParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) WithContext(ctx context.Context) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) WithConfigUID(configUID string) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) WithMachineUID(machineUID string) *V1CloudConfigsVspherePoolMachinesUIDGetParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs vsphere pool machines Uid get params
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsVspherePoolMachinesUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
