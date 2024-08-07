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

// NewV1CloudConfigsCustomMachinePoolUpdateParams creates a new V1CloudConfigsCustomMachinePoolUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsCustomMachinePoolUpdateParams() *V1CloudConfigsCustomMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCustomMachinePoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsCustomMachinePoolUpdateParamsWithTimeout creates a new V1CloudConfigsCustomMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsCustomMachinePoolUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsCustomMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCustomMachinePoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsCustomMachinePoolUpdateParamsWithContext creates a new V1CloudConfigsCustomMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsCustomMachinePoolUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsCustomMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCustomMachinePoolUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsCustomMachinePoolUpdateParamsWithHTTPClient creates a new V1CloudConfigsCustomMachinePoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsCustomMachinePoolUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsCustomMachinePoolUpdateParams {
	var ()
	return &V1CloudConfigsCustomMachinePoolUpdateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsCustomMachinePoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs custom machine pool update operation typically these are written to a http.Request
*/
type V1CloudConfigsCustomMachinePoolUpdateParams struct {

	/*Body*/
	Body *models.V1CustomMachinePoolConfigEntity
	/*CloudType
	  Cluster's cloud type

	*/
	CloudType string
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

// WithTimeout adds the timeout to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsCustomMachinePoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsCustomMachinePoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsCustomMachinePoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WithBody(body *models.V1CustomMachinePoolConfigEntity) *V1CloudConfigsCustomMachinePoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) SetBody(body *models.V1CustomMachinePoolConfigEntity) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WithCloudType(cloudType string) *V1CloudConfigsCustomMachinePoolUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithConfigUID adds the configUID to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsCustomMachinePoolUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WithMachinePoolName adds the machinePoolName to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WithMachinePoolName(machinePoolName string) *V1CloudConfigsCustomMachinePoolUpdateParams {
	o.SetMachinePoolName(machinePoolName)
	return o
}

// SetMachinePoolName adds the machinePoolName to the v1 cloud configs custom machine pool update params
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) SetMachinePoolName(machinePoolName string) {
	o.MachinePoolName = machinePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsCustomMachinePoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
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
