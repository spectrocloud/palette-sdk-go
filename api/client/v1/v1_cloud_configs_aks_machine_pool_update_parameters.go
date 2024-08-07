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

// NewV1CloudConfigsAksMachinePoolUpdateParams creates a new V1CloudConfigsAksMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsAksMachinePoolUpdateParams() *V1CloudConfigsAksMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAksMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAksMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsAksMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAksMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAksMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAksMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAksMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsAksMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAksMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsAksMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAksMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAksMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsAksMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAksMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAksMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAksMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAksMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aks machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsAksMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1AzureMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAksMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsAksMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAksMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) WithBody(body *models.V1AzureMachinePoolConfigEntity) *V1CloudConfigsAksMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) SetBody(body *models.V1AzureMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsAksMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsAksMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs aks machine pool update params
func (o *V1CloudConfigsAksMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAksMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
