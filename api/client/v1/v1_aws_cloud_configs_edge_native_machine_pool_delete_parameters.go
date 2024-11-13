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

// NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams creates a new V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams object
// with the default values initialized.
func NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams() *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	var ()
	return &V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParamsWithTimeout creates a new V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParamsWithTimeout(timeout time.Duration) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	var ()
	return &V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams{

		timeout: timeout,
	}
}

// NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext creates a new V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(ctx context.Context) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	var ()
	return &V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams{

		Context: ctx,
	}
}

// NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParamsWithHTTPClient creates a new V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParamsWithHTTPClient(client *http.Client) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	var ()
	return &V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams{
		HTTPClient: client,
	}
}

/*
V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams contains all the parameters to send to the API endpoint
for the v1 aws cloud configs edge native machine pool delete operation typically these are written to a http.Request
*/
type V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams struct {

	/*ConfigUID
	  AWS Cluster's cloud config uid

	*/
	ConfigUID string
	/*MachinePoolName
	  Edge-native machine pool name

	*/
	MachinePoolName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) WithTimeout(timeout time.Duration) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) WithContext(ctx context.Context) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) WithHTTPClient(client *http.Client) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) WithConfigUID(configUID string) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) WithMachinePoolName(machinePoolName string) *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 aws cloud configs edge native machine pool delete params
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsCloudConfigsEdgeNativeMachinePoolDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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