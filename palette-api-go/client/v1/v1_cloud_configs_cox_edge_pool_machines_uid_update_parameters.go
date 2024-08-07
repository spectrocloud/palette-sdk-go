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

// NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams creates a new V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams() *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParamsWithTimeout creates a new V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParamsWithContext creates a new V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParamsWithHTTPClient creates a new V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsCoxEdgePoolMachinesUIDUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	var ()
	return &V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs cox edge pool machines Uid update operation typically these are written to a http.Request
*/
type V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams struct {

	/*Body*/
	Body *models.V1CoxEdgeMachine
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string
	/*MachinePoolName
	  Machine pool name

	*/
	MachinePoolName string
	/*MachineUID
	  Machine uid

	*/
	MachineUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WithBody(body *models.V1CoxEdgeMachine) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) SetBody(body *models.V1CoxEdgeMachine) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WithMachineUID adds the machineUID to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WithMachineUID(machineUID string) *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams {
	o.SetMachineUID(machineUID)
	return o
}

// SetMachineUID adds the machineUid to the v1 cloud configs cox edge pool machines Uid update params
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) SetMachineUID(machineUID string) {
	o.MachineUID = machineUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsCoxEdgePoolMachinesUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param machineUid
	if err := r.SetPathParam("machineUid", o.MachineUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
