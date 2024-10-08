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

// NewV1CloudConfigsEdgePoolMachinesUIDGetParams creates a new V1CloudConfigsEdgePoolMachinesUIDGetParams object
// with the default values initialized.
func NewV1CloudConfigsEdgePoolMachinesUIDGetParams() *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEdgePoolMachinesUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithTimeout creates a new V1CloudConfigsEdgePoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEdgePoolMachinesUIDGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithContext creates a new V1CloudConfigsEdgePoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithContext(ctx context.Context) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEdgePoolMachinesUIDGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithHTTPClient creates a new V1CloudConfigsEdgePoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEdgePoolMachinesUIDGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEdgePoolMachinesUIDGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge pool machines Uid get operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgePoolMachinesUIDGetParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) WithContext(ctx context.Context) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) WithConfigUID(configUID string) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) WithMachineUID(machineUID string) *V1CloudConfigsEdgePoolMachinesUIDGetParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs edge pool machines Uid get params
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgePoolMachinesUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
