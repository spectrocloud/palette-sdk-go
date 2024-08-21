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

// NewV1CloudConfigsCoxEdgeMachinePoolUpdateParams creates a new V1CloudConfigsCoxEdgeMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsCoxEdgeMachinePoolUpdateParams() *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsCoxEdgeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsCoxEdgeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsCoxEdgeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgeMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsCoxEdgeMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs cox edge machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsCoxEdgeMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1CoxEdgeMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) WithBody(body *models.V1CoxEdgeMachinePoolConfigEntity) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) SetBody(body *models.V1CoxEdgeMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsCoxEdgeMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs cox edge machine pool update params
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsCoxEdgeMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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