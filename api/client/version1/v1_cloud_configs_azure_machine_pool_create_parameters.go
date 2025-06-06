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

// NewV1CloudConfigsAzureMachinePoolCreateParams creates a new V1CloudConfigsAzureMachinePoolCreateParams object
// with the default values initialized.
func NewV1CloudConfigsAzureMachinePoolCreateParams() *V1CloudConfigsAzureMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAzureMachinePoolCreateParamsWithTimeout creates a new V1CloudConfigsAzureMachinePoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAzureMachinePoolCreateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAzureMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAzureMachinePoolCreateParamsWithContext creates a new V1CloudConfigsAzureMachinePoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAzureMachinePoolCreateParamsWithContext(ctx context.Context) *V1CloudConfigsAzureMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolCreateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAzureMachinePoolCreateParamsWithHTTPClient creates a new V1CloudConfigsAzureMachinePoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAzureMachinePoolCreateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAzureMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAzureMachinePoolCreateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAzureMachinePoolCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs azure machine pool create operation typically these are written to a http.Request
*/
type V1CloudConfigsAzureMachinePoolCreateParams struct {

	/*Body*/
	Body *models.V1AzureMachinePoolConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAzureMachinePoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) WithContext(ctx context.Context) *V1CloudConfigsAzureMachinePoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAzureMachinePoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) WithBody(body *models.V1AzureMachinePoolConfigEntity) *V1CloudConfigsAzureMachinePoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) SetBody(body *models.V1AzureMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) WithConfigUID(configUID string) *V1CloudConfigsAzureMachinePoolCreateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs azure machine pool create params
func (o *V1CloudConfigsAzureMachinePoolCreateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAzureMachinePoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
