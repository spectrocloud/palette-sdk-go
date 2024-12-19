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

// NewV1CloudConfigsEksPoolMachinesAddParams creates a new V1CloudConfigsEksPoolMachinesAddParams object
// with the default values initialized.
func NewV1CloudConfigsEksPoolMachinesAddParams() *V1CloudConfigsEksPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEksPoolMachinesAddParamsWithTimeout creates a new V1CloudConfigsEksPoolMachinesAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEksPoolMachinesAddParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEksPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesAddParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEksPoolMachinesAddParamsWithContext creates a new V1CloudConfigsEksPoolMachinesAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEksPoolMachinesAddParamsWithContext(ctx context.Context) *V1CloudConfigsEksPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesAddParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEksPoolMachinesAddParamsWithHTTPClient creates a new V1CloudConfigsEksPoolMachinesAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEksPoolMachinesAddParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEksPoolMachinesAddParams {
	var ()
	return &V1CloudConfigsEksPoolMachinesAddParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsEksPoolMachinesAddParams contains all the parameters to send to the API endpoint
for the v1 cloud configs eks pool machines add operation typically these are written to a http.Request
*/
type V1CloudConfigsEksPoolMachinesAddParams struct {

	/*Body*/
	Body *models.V1AwsMachine
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

// WithTimeout adds the timeout to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEksPoolMachinesAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) WithContext(ctx context.Context) *V1CloudConfigsEksPoolMachinesAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEksPoolMachinesAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) WithBody(body *models.V1AwsMachine) *V1CloudConfigsEksPoolMachinesAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) SetBody(body *models.V1AwsMachine) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) WithConfigUID(configUID string) *V1CloudConfigsEksPoolMachinesAddParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsEksPoolMachinesAddParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs eks pool machines add params
func (o *V1CloudConfigsEksPoolMachinesAddParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEksPoolMachinesAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
