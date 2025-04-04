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

// NewV1CloudConfigsMaasMachinePoolUpdateParams creates a new V1CloudConfigsMaasMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsMaasMachinePoolUpdateParams() *V1CloudConfigsMaasMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsMaasMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsMaasMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsMaasMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsMaasMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsMaasMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsMaasMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsMaasMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsMaasMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsMaasMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsMaasMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsMaasMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsMaasMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsMaasMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs maas machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsMaasMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1MaasMachinePoolConfigEntity
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

// WithTimeout adds the timeout to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsMaasMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsMaasMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsMaasMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) WithBody(body *models.V1MaasMachinePoolConfigEntity) *V1CloudConfigsMaasMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) SetBody(body *models.V1MaasMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsMaasMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsMaasMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs maas machine pool update params
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsMaasMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
