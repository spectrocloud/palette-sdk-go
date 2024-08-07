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

// NewV1CloudConfigsEdgeMachinePoolUpdateParams creates a new V1CloudConfigsEdgeMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeMachinePoolUpdateParams() *V1CloudConfigsEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsEdgeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsEdgeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsEdgeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEdgeMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1EdgeMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) WithBody(body *models.V1EdgeMachinePoolConfigEntity) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) SetBody(body *models.V1EdgeMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEdgeMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs edge machine pool update params
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
