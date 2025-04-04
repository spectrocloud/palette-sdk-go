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

// NewV1CloudConfigsCustomPoolMachinesAddParams creates a new V1CloudConfigsCustomPoolMachinesAddParams object
// with the default values initialized.
func NewV1CloudConfigsCustomPoolMachinesAddParams() *V1CloudConfigsCustomPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsCustomPoolMachinesAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsCustomPoolMachinesAddParamsWithTimeout creates a new V1CloudConfigsCustomPoolMachinesAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsCustomPoolMachinesAddParamsWithTimeout(timeout time.Duration) *V1CloudConfigsCustomPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsCustomPoolMachinesAddParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsCustomPoolMachinesAddParamsWithContext creates a new V1CloudConfigsCustomPoolMachinesAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsCustomPoolMachinesAddParamsWithContext(ctx context.Context) *V1CloudConfigsCustomPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsCustomPoolMachinesAddParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsCustomPoolMachinesAddParamsWithHTTPClient creates a new V1CloudConfigsCustomPoolMachinesAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsCustomPoolMachinesAddParamsWithHTTPClient(client *http.Client) *V1CloudConfigsCustomPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsCustomPoolMachinesAddParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsCustomPoolMachinesAddParams contains all the parameters to send to the API endpoint
for the v1 cloud configs custom pool machines add operation typically these are written to a http.Request
*/
type V1CloudConfigsCustomPoolMachinesAddParams struct {

	/*Body*/
	Body *models.V1CustomMachine
	/*CloudType
	  Cluster's cloud type

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WithTimeout(timeout time.Duration) *V1CloudConfigsCustomPoolMachinesAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WithContext(ctx context.Context) *V1CloudConfigsCustomPoolMachinesAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WithHTTPClient(client *http.Client) *V1CloudConfigsCustomPoolMachinesAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WithBody(body *models.V1CustomMachine) *V1CloudConfigsCustomPoolMachinesAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) SetBody(body *models.V1CustomMachine) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WithCloudType(cloudType string) *V1CloudConfigsCustomPoolMachinesAddParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithConfigUID adds the configUID to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WithConfigUID(configUID string) *V1CloudConfigsCustomPoolMachinesAddParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsCustomPoolMachinesAddParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs custom pool machines add params
func (o *V1CloudConfigsCustomPoolMachinesAddParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsCustomPoolMachinesAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
