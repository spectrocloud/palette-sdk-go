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

// NewV1CloudConfigsAwsMachinePoolCreateParams creates a new V1CloudConfigsAwsMachinePoolCreateParams object
// with the default values initialized.
func NewV1CloudConfigsAwsMachinePoolCreateParams() *V1CloudConfigsAwsMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAwsMachinePoolCreateParamsWithTimeout creates a new V1CloudConfigsAwsMachinePoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAwsMachinePoolCreateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAwsMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAwsMachinePoolCreateParamsWithContext creates a new V1CloudConfigsAwsMachinePoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAwsMachinePoolCreateParamsWithContext(ctx context.Context) *V1CloudConfigsAwsMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolCreateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAwsMachinePoolCreateParamsWithHTTPClient creates a new V1CloudConfigsAwsMachinePoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAwsMachinePoolCreateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAwsMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsAwsMachinePoolCreateParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAwsMachinePoolCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aws machine pool create operation typically these are written to a http.Request
*/
type V1CloudConfigsAwsMachinePoolCreateParams struct {

	/*Body*/
	Body *models.V1AwsMachinePoolConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAwsMachinePoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) WithContext(ctx context.Context) *V1CloudConfigsAwsMachinePoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAwsMachinePoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) WithBody(body *models.V1AwsMachinePoolConfigEntity) *V1CloudConfigsAwsMachinePoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) SetBody(body *models.V1AwsMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) WithConfigUID(configUID string) *V1CloudConfigsAwsMachinePoolCreateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aws machine pool create params
func (o *V1CloudConfigsAwsMachinePoolCreateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAwsMachinePoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
