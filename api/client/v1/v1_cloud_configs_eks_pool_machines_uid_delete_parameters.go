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

// NewV1CloudConfigsEksPoolMachinesUIDDeleteParams creates a new V1CloudConfigsEksPoolMachinesUIDDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsEksPoolMachinesUIDDeleteParams() *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDDeleteParamsWithTimeout creates a new V1CloudConfigsEksPoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEksPoolMachinesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDDeleteParamsWithContext creates a new V1CloudConfigsEksPoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEksPoolMachinesUIDDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDDeleteParamsWithHTTPClient creates a new V1CloudConfigsEksPoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEksPoolMachinesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEksPoolMachinesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs eks pool machines Uid delete operation typically these are written to a http.Request
*/
type V1CloudConfigsEksPoolMachinesUIDDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) WithMachineUID(machineUID string) *V1CloudConfigsEksPoolMachinesUIDDeleteParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs eks pool machines Uid delete params
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEksPoolMachinesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
