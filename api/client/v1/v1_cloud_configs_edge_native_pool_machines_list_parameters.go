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

// NewV1CloudConfigsEdgeNativePoolMachinesListParams creates a new V1CloudConfigsEdgeNativePoolMachinesListParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeNativePoolMachinesListParams() *V1CloudConfigsEdgeNativePoolMachinesListParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithTimeout creates a new V1CloudConfigsEdgeNativePoolMachinesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesListParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext creates a new V1CloudConfigsEdgeNativePoolMachinesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesListParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithHTTPClient creates a new V1CloudConfigsEdgeNativePoolMachinesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	var ()
	return &V1CloudConfigsEdgeNativePoolMachinesListParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEdgeNativePoolMachinesListParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge native pool machines list operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeNativePoolMachinesListParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEdgeNativePoolMachinesListParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs edge native pool machines list params
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeNativePoolMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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