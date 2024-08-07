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

// NewV1CloudConfigsEdgeMachinePoolDeleteParams creates a new V1CloudConfigsEdgeMachinePoolDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeMachinePoolDeleteParams() *V1CloudConfigsEdgeMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithTimeout creates a new V1CloudConfigsEdgeMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithContext creates a new V1CloudConfigsEdgeMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithHTTPClient creates a new V1CloudConfigsEdgeMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolDeleteParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsEdgeMachinePoolDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge machine pool delete operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeMachinePoolDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEdgeMachinePoolDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs edge machine pool delete params
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeMachinePoolDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
