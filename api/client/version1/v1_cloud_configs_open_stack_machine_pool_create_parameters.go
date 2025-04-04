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

// NewV1CloudConfigsOpenStackMachinePoolCreateParams creates a new V1CloudConfigsOpenStackMachinePoolCreateParams object
// with the default values initialized.
func NewV1CloudConfigsOpenStackMachinePoolCreateParams() *V1CloudConfigsOpenStackMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithTimeout creates a new V1CloudConfigsOpenStackMachinePoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithContext creates a new V1CloudConfigsOpenStackMachinePoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithContext(ctx context.Context) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolCreateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithHTTPClient creates a new V1CloudConfigsOpenStackMachinePoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsOpenStackMachinePoolCreateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsOpenStackMachinePoolCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs open stack machine pool create operation typically these are written to a http.Request
*/
type V1CloudConfigsOpenStackMachinePoolCreateParams struct {

	/*Body*/
	Body *models.V1OpenStackMachinePoolConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) WithContext(ctx context.Context) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) WithBody(body *models.V1OpenStackMachinePoolConfigEntity) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) SetBody(body *models.V1OpenStackMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) WithConfigUID(configUID string) *V1CloudConfigsOpenStackMachinePoolCreateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs open stack machine pool create params
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsOpenStackMachinePoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
