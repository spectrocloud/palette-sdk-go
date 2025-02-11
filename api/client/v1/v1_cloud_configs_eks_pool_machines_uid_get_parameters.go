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

// NewV1CloudConfigsEksPoolMachinesUIDGetParams creates a new V1CloudConfigsEksPoolMachinesUIDGetParams object
// with the default values initialized.
func NewV1CloudConfigsEksPoolMachinesUIDGetParams() *V1CloudConfigsEksPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithTimeout creates a new V1CloudConfigsEksPoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithContext creates a new V1CloudConfigsEksPoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithContext(ctx context.Context) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithHTTPClient creates a new V1CloudConfigsEksPoolMachinesUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesUIDGetParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsEksPoolMachinesUIDGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs eks pool machines Uid get operation typically these are written to a http.Request
*/
type V1CloudConfigsEksPoolMachinesUIDGetParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) WithContext(ctx context.Context) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) WithConfigUID(configUID string) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) WithMachineUID(machineUID string) *V1CloudConfigsEksPoolMachinesUIDGetParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs eks pool machines Uid get params
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEksPoolMachinesUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
