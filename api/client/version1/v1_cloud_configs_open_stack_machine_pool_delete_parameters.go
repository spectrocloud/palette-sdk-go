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

// NewV1CloudConfigsOpenStackMachinePoolDeleteParams creates a new V1CloudConfigsOpenStackMachinePoolDeleteParams object
// with the default values initialized.
func NewV1CloudConfigsOpenStackMachinePoolDeleteParams() *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithTimeout creates a new V1CloudConfigsOpenStackMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithTimeout(timeout time.Duration) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithContext creates a new V1CloudConfigsOpenStackMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithContext(ctx context.Context) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithHTTPClient creates a new V1CloudConfigsOpenStackMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithHTTPClient(client *http.Client) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsOpenStackMachinePoolDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud configs open stack machine pool delete operation typically these are written to a http.Request
*/
type V1CloudConfigsOpenStackMachinePoolDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) WithTimeout(timeout time.Duration) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) WithContext(ctx context.Context) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) WithHTTPClient(client *http.Client) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) WithConfigUID(configUID string) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsOpenStackMachinePoolDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs open stack machine pool delete params
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsOpenStackMachinePoolDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
