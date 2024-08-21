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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1CloudConfigsVspherePoolMachinesUIDUpdateParams creates a new V1CloudConfigsVspherePoolMachinesUIDUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsVspherePoolMachinesUIDUpdateParams() *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDUpdateParamsWithTimeout creates a new V1CloudConfigsVspherePoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsVspherePoolMachinesUIDUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDUpdateParamsWithContext creates a new V1CloudConfigsVspherePoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsVspherePoolMachinesUIDUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsVspherePoolMachinesUIDUpdateParamsWithHTTPClient creates a new V1CloudConfigsVspherePoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsVspherePoolMachinesUIDUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsVspherePoolMachinesUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsVspherePoolMachinesUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs vsphere pool machines Uid update operation typically these are written to a http.Request
*/
type V1CloudConfigsVspherePoolMachinesUIDUpdateParams struct {

	/*Body*/
	Body *models.V1VsphereMachine
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

// WithTimeout adds the timeout to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WithBody(body *models.V1VsphereMachine) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) SetBody(body *models.V1VsphereMachine) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WithMachineUID(machineUID string) *V1CloudConfigsVspherePoolMachinesUIDUpdateParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs vsphere pool machines Uid update params
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsVspherePoolMachinesUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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