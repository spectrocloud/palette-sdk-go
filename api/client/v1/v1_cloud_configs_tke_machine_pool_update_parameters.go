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

// NewV1CloudConfigsTkeMachinePoolUpdateParams creates a new V1CloudConfigsTkeMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsTkeMachinePoolUpdateParams() *V1CloudConfigsTkeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsTkeMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsTkeMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsTkeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsTkeMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsTkeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsTkeMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsTkeMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsTkeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsTkeMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsTkeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsTkeMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsTkeMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsTkeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsTkeMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsTkeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsTkeMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsTkeMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs tke machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsTkeMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1TencentMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsTkeMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsTkeMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsTkeMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) WithBody(body *models.V1TencentMachinePoolConfigEntity) *V1CloudConfigsTkeMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) SetBody(body *models.V1TencentMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsTkeMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsTkeMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs tke machine pool update params
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsTkeMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
