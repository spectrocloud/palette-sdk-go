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

// NewV1CloudConfigsOpenStackPoolMachinesListParams creates a new V1CloudConfigsOpenStackPoolMachinesListParams object
// with the default values initialized.
func NewV1CloudConfigsOpenStackPoolMachinesListParams() *V1CloudConfigsOpenStackPoolMachinesListParams {
	var ()
	return &V1CloudConfigsOpenStackPoolMachinesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsOpenStackPoolMachinesListParamsWithTimeout creates a new V1CloudConfigsOpenStackPoolMachinesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsOpenStackPoolMachinesListParamsWithTimeout(timeout time.Duration) *V1CloudConfigsOpenStackPoolMachinesListParams {
	var ()
	return &V1CloudConfigsOpenStackPoolMachinesListParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsOpenStackPoolMachinesListParamsWithContext creates a new V1CloudConfigsOpenStackPoolMachinesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsOpenStackPoolMachinesListParamsWithContext(ctx context.Context) *V1CloudConfigsOpenStackPoolMachinesListParams {
	var ()
	return &V1CloudConfigsOpenStackPoolMachinesListParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsOpenStackPoolMachinesListParamsWithHTTPClient creates a new V1CloudConfigsOpenStackPoolMachinesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsOpenStackPoolMachinesListParamsWithHTTPClient(client *http.Client) *V1CloudConfigsOpenStackPoolMachinesListParams {
	var ()
	return &V1CloudConfigsOpenStackPoolMachinesListParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsOpenStackPoolMachinesListParams contains all the parameters to send to the API endpoint
for the v1 cloud configs open stack pool machines list operation typically these are written to a http.Request
*/
type V1CloudConfigsOpenStackPoolMachinesListParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) WithTimeout(timeout time.Duration) *V1CloudConfigsOpenStackPoolMachinesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) WithContext(ctx context.Context) *V1CloudConfigsOpenStackPoolMachinesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) WithHTTPClient(client *http.Client) *V1CloudConfigsOpenStackPoolMachinesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) WithConfigUID(configUID string) *V1CloudConfigsOpenStackPoolMachinesListParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsOpenStackPoolMachinesListParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs open stack pool machines list params
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsOpenStackPoolMachinesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
