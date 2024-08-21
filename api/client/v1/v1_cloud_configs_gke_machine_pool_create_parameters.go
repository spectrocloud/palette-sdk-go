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

// NewV1CloudConfigsGkeMachinePoolCreateParams creates a new V1CloudConfigsGkeMachinePoolCreateParams object
// with the default values initialized.
func NewV1CloudConfigsGkeMachinePoolCreateParams() *V1CloudConfigsGkeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsGkeMachinePoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGkeMachinePoolCreateParamsWithTimeout creates a new V1CloudConfigsGkeMachinePoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGkeMachinePoolCreateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGkeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsGkeMachinePoolCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsGkeMachinePoolCreateParamsWithContext creates a new V1CloudConfigsGkeMachinePoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGkeMachinePoolCreateParamsWithContext(ctx context.Context) *V1CloudConfigsGkeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsGkeMachinePoolCreateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsGkeMachinePoolCreateParamsWithHTTPClient creates a new V1CloudConfigsGkeMachinePoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGkeMachinePoolCreateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGkeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsGkeMachinePoolCreateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsGkeMachinePoolCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs gke machine pool create operation typically these are written to a http.Request
*/
type V1CloudConfigsGkeMachinePoolCreateParams struct {

	/*Body*/
	Body *models.V1GcpMachinePoolConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGkeMachinePoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) WithContext(ctx context.Context) *V1CloudConfigsGkeMachinePoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGkeMachinePoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) WithBody(body *models.V1GcpMachinePoolConfigEntity) *V1CloudConfigsGkeMachinePoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) SetBody(body *models.V1GcpMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) WithConfigUID(configUID string) *V1CloudConfigsGkeMachinePoolCreateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs gke machine pool create params
func (o *V1CloudConfigsGkeMachinePoolCreateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGkeMachinePoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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