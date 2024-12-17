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

// NewV1CloudConfigsAzureMachinePoolUpdateParams creates a new V1CloudConfigsAzureMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsAzureMachinePoolUpdateParams() *V1CloudConfigsAzureMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAzureMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsAzureMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAzureMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAzureMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAzureMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsAzureMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAzureMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsAzureMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAzureMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsAzureMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAzureMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAzureMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsAzureMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs azure machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsAzureMachinePoolUpdateParams struct {

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

// WithTimeout adds the timeout to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAzureMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsAzureMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAzureMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) WithBody(body *models.V1AzureMachinePoolConfigEntity) *V1CloudConfigsAzureMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) SetBody(body *models.V1AzureMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsAzureMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsAzureMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs azure machine pool update params
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAzureMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
