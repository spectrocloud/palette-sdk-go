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
)

// NewV1CloudConfigsGkePoolMachinesUIDDeleteParams creates a new V1CloudConfigsGkePoolMachinesUIDDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsGkePoolMachinesUIDDeleteParams() *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsGkePoolMachinesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGkePoolMachinesUIDDeleteParamsWithTimeout creates a new V1CloudConfigsGkePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGkePoolMachinesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsGkePoolMachinesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsGkePoolMachinesUIDDeleteParamsWithContext creates a new V1CloudConfigsGkePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGkePoolMachinesUIDDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsGkePoolMachinesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsGkePoolMachinesUIDDeleteParamsWithHTTPClient creates a new V1CloudConfigsGkePoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGkePoolMachinesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsGkePoolMachinesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsGkePoolMachinesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs gke pool machines Uid delete operation typically these are written to a http.Request
*/
type V1CloudConfigsGkePoolMachinesUIDDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) WithMachineUID(machineUID string) *V1CloudConfigsGkePoolMachinesUIDDeleteParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs gke pool machines Uid delete params
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGkePoolMachinesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
