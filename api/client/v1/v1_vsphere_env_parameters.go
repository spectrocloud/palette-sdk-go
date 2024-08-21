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

// NewV1VsphereEnvParams creates a new V1VsphereEnvParams object
// with the default values initialized.
func NewV1VsphereEnvParams() *V1VsphereEnvParams {
	var ()
	return &V1VsphereEnvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VsphereEnvParamsWithTimeout creates a new V1VsphereEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VsphereEnvParamsWithTimeout(timeout time.Duration) *V1VsphereEnvParams {
	var ()
	return &V1VsphereEnvParams{

		timeout: timeout,
	}
}

// NewV1VsphereEnvParamsWithContext creates a new V1VsphereEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VsphereEnvParamsWithContext(ctx context.Context) *V1VsphereEnvParams {
	var ()
	return &V1VsphereEnvParams{

		Context: ctx,
	}
}

// NewV1VsphereEnvParamsWithHTTPClient creates a new V1VsphereEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VsphereEnvParamsWithHTTPClient(client *http.Client) *V1VsphereEnvParams {
	var ()
	return &V1VsphereEnvParams{
		HTTPClient: client,
	}
}

/*
V1VsphereEnvParams contains all the parameters to send to the API endpoint
for the v1 vsphere env operation typically these are written to a http.Request
*/
type V1VsphereEnvParams struct {

	/*VsphereCloudAccount
	  Request payload for VSphere cloud account

	*/
	VsphereCloudAccount *models.V1VsphereCloudAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 vsphere env params
func (o *V1VsphereEnvParams) WithTimeout(timeout time.Duration) *V1VsphereEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 vsphere env params
func (o *V1VsphereEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 vsphere env params
func (o *V1VsphereEnvParams) WithContext(ctx context.Context) *V1VsphereEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 vsphere env params
func (o *V1VsphereEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 vsphere env params
func (o *V1VsphereEnvParams) WithHTTPClient(client *http.Client) *V1VsphereEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 vsphere env params
func (o *V1VsphereEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVsphereCloudAccount adds the vsphereCloudAccount to the v1 vsphere env params
func (o *V1VsphereEnvParams) WithVsphereCloudAccount(vsphereCloudAccount *models.V1VsphereCloudAccount) *V1VsphereEnvParams {
	o.SetVsphereCloudAccount(vsphereCloudAccount)
	return o
}

// SetVsphereCloudAccount adds the vsphereCloudAccount to the v1 vsphere env params
func (o *V1VsphereEnvParams) SetVsphereCloudAccount(vsphereCloudAccount *models.V1VsphereCloudAccount) {
	o.VsphereCloudAccount = vsphereCloudAccount
}

// WriteToRequest writes these params to a swagger request
func (o *V1VsphereEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VsphereCloudAccount != nil {
		if err := r.SetBodyParam(o.VsphereCloudAccount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}