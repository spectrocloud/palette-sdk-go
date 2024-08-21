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

// NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams creates a new V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams() *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithTimeout creates a new V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithContext creates a new V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithHTTPClient creates a new V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs machine pools machine Uid maintenance update operation typically these are written to a http.Request
*/
type V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams struct {

	/*Body*/
	Body *models.V1MachineMaintenance
	/*CloudType
	  Cloud type

	*/
	CloudType string
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

// WithTimeout adds the timeout to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithBody(body *models.V1MachineMaintenance) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetBody(body *models.V1MachineMaintenance) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithCloudType(cloudType string) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithConfigUID adds the configUID to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WithMachineUID(machineUID string) *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs machine pools machine Uid maintenance update params
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
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