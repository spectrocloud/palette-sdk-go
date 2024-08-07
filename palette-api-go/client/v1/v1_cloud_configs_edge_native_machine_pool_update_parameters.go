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

	"github.com/spectrocloud/palette-api-go/models"
)

// NewV1CloudConfigsEdgeNativeMachinePoolUpdateParams creates a new V1CloudConfigsEdgeNativeMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeNativeMachinePoolUpdateParams() *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsEdgeNativeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsEdgeNativeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsEdgeNativeMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsEdgeNativeMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge native machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeNativeMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1EdgeNativeMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) WithBody(body *models.V1EdgeNativeMachinePoolConfigEntity) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) SetBody(body *models.V1EdgeNativeMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEdgeNativeMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs edge native machine pool update params
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeNativeMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
