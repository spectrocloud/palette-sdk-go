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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1CloudConfigsMaasPoolMachinesUIDUpdateParams creates a new V1CloudConfigsMaasPoolMachinesUIDUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsMaasPoolMachinesUIDUpdateParams() *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsMaasPoolMachinesUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsMaasPoolMachinesUIDUpdateParamsWithTimeout creates a new V1CloudConfigsMaasPoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsMaasPoolMachinesUIDUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsMaasPoolMachinesUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsMaasPoolMachinesUIDUpdateParamsWithContext creates a new V1CloudConfigsMaasPoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsMaasPoolMachinesUIDUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsMaasPoolMachinesUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsMaasPoolMachinesUIDUpdateParamsWithHTTPClient creates a new V1CloudConfigsMaasPoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsMaasPoolMachinesUIDUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsMaasPoolMachinesUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsMaasPoolMachinesUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs maas pool machines Uid update operation typically these are written to a http.Request
*/
type V1CloudConfigsMaasPoolMachinesUIDUpdateParams struct {

	/*Body*/
	Body *models.V1MaasMachine
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

// WithTimeout adds the timeout to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WithBody(body *models.V1MaasMachine) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) SetBody(body *models.V1MaasMachine) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WithMachineUID(machineUID string) *V1CloudConfigsMaasPoolMachinesUIDUpdateParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs maas pool machines Uid update params
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsMaasPoolMachinesUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
