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

// NewV1CloudConfigsVirtualPoolMachinesUIDGetParams creates a new V1CloudConfigsVirtualPoolMachinesUIDGetParams object
// with the default values initialized.
func NewV1CloudConfigsVirtualPoolMachinesUIDGetParams() *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVirtualPoolMachinesUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithTimeout creates a new V1CloudConfigsVirtualPoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVirtualPoolMachinesUIDGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithContext creates a new V1CloudConfigsVirtualPoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithContext(ctx context.Context) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVirtualPoolMachinesUIDGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithHTTPClient creates a new V1CloudConfigsVirtualPoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsVirtualPoolMachinesUIDGetParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsVirtualPoolMachinesUIDGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs virtual pool machines Uid get operation typically these are written to a http.Request
*/
type V1CloudConfigsVirtualPoolMachinesUIDGetParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) WithContext(ctx context.Context) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) WithConfigUID(configUID string) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) WithMachineUID(machineUID string) *V1CloudConfigsVirtualPoolMachinesUIDGetParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs virtual pool machines Uid get params
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsVirtualPoolMachinesUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
