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

// NewV1CloudConfigsGcpMachinePoolDeleteParams creates a new V1CloudConfigsGcpMachinePoolDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsGcpMachinePoolDeleteParams() *V1CloudConfigsGcpMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGcpMachinePoolDeleteParamsWithTimeout creates a new V1CloudConfigsGcpMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGcpMachinePoolDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGcpMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsGcpMachinePoolDeleteParamsWithContext creates a new V1CloudConfigsGcpMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGcpMachinePoolDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsGcpMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsGcpMachinePoolDeleteParamsWithHTTPClient creates a new V1CloudConfigsGcpMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGcpMachinePoolDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGcpMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsGcpMachinePoolDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs gcp machine pool delete operation typically these are written to a http.Request
*/
type V1CloudConfigsGcpMachinePoolDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGcpMachinePoolDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsGcpMachinePoolDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGcpMachinePoolDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsGcpMachinePoolDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsGcpMachinePoolDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs gcp machine pool delete params
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGcpMachinePoolDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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