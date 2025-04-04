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

// NewV1CloudConfigsGcpMachinePoolUpdateParams creates a new V1CloudConfigsGcpMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsGcpMachinePoolUpdateParams() *V1CloudConfigsGcpMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGcpMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsGcpMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGcpMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGcpMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsGcpMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsGcpMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGcpMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsGcpMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsGcpMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsGcpMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGcpMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGcpMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsGcpMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsGcpMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs gcp machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsGcpMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1GcpMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGcpMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsGcpMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGcpMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) WithBody(body *models.V1GcpMachinePoolConfigEntity) *V1CloudConfigsGcpMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) SetBody(body *models.V1GcpMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsGcpMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsGcpMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs gcp machine pool update params
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGcpMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
