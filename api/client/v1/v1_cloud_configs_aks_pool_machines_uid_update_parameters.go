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

// NewV1CloudConfigsAksPoolMachinesUIDUpdateParams creates a new V1CloudConfigsAksPoolMachinesUIDUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsAksPoolMachinesUIDUpdateParams() *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsAksPoolMachinesUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAksPoolMachinesUIDUpdateParamsWithTimeout creates a new V1CloudConfigsAksPoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAksPoolMachinesUIDUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsAksPoolMachinesUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAksPoolMachinesUIDUpdateParamsWithContext creates a new V1CloudConfigsAksPoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAksPoolMachinesUIDUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsAksPoolMachinesUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAksPoolMachinesUIDUpdateParamsWithHTTPClient creates a new V1CloudConfigsAksPoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAksPoolMachinesUIDUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsAksPoolMachinesUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAksPoolMachinesUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aks pool machines Uid update operation typically these are written to a http.Request
*/
type V1CloudConfigsAksPoolMachinesUIDUpdateParams struct {

	/*Body*/
	Body *models.V1AzureMachine
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

// WithTimeout adds the timeout to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WithBody(body *models.V1AzureMachine) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) SetBody(body *models.V1AzureMachine) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WithMachineUID(machineUID string) *V1CloudConfigsAksPoolMachinesUIDUpdateParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs aks pool machines Uid update params
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAksPoolMachinesUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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