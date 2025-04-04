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

// NewV1CloudConfigsAwsPoolMachinesUIDDeleteParams creates a new V1CloudConfigsAwsPoolMachinesUIDDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsAwsPoolMachinesUIDDeleteParams() *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsAwsPoolMachinesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAwsPoolMachinesUIDDeleteParamsWithTimeout creates a new V1CloudConfigsAwsPoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAwsPoolMachinesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsAwsPoolMachinesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAwsPoolMachinesUIDDeleteParamsWithContext creates a new V1CloudConfigsAwsPoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAwsPoolMachinesUIDDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsAwsPoolMachinesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAwsPoolMachinesUIDDeleteParamsWithHTTPClient creates a new V1CloudConfigsAwsPoolMachinesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAwsPoolMachinesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	var ()
	return &V1CloudConfigsAwsPoolMachinesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAwsPoolMachinesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aws pool machines Uid delete operation typically these are written to a http.Request
*/
type V1CloudConfigsAwsPoolMachinesUIDDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) WithMachineUID(machineUID string) *V1CloudConfigsAwsPoolMachinesUIDDeleteParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs aws pool machines Uid delete params
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAwsPoolMachinesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
