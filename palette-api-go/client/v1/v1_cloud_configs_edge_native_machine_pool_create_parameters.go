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

// NewV1CloudConfigsEdgeNativeMachinePoolCreateParams creates a new V1CloudConfigsEdgeNativeMachinePoolCreateParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeNativeMachinePoolCreateParams() *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithTimeout creates a new V1CloudConfigsEdgeNativeMachinePoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext creates a new V1CloudConfigsEdgeNativeMachinePoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolCreateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithHTTPClient creates a new V1CloudConfigsEdgeNativeMachinePoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEdgeNativeMachinePoolCreateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsEdgeNativeMachinePoolCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge native machine pool create operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeNativeMachinePoolCreateParams struct {

	/*Body*/
	Body *models.V1EdgeNativeMachinePoolConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) WithBody(body *models.V1EdgeNativeMachinePoolConfigEntity) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) SetBody(body *models.V1EdgeNativeMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeNativeMachinePoolCreateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge native machine pool create params
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeNativeMachinePoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
