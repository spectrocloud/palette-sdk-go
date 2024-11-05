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

// NewV1CloudConfigsAwsMachinePoolUpdateParams creates a new V1CloudConfigsAwsMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsAwsMachinePoolUpdateParams() *V1CloudConfigsAwsMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAwsMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsAwsMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAwsMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAwsMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAwsMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsAwsMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAwsMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsAwsMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAwsMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsAwsMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAwsMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAwsMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAwsMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aws machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsAwsMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1AwsMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAwsMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsAwsMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAwsMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) WithBody(body *models.V1AwsMachinePoolConfigEntity) *V1CloudConfigsAwsMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) SetBody(body *models.V1AwsMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsAwsMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsAwsMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs aws machine pool update params
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAwsMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
