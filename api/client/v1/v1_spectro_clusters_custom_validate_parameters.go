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

// NewV1SpectroClustersCustomValidateParams creates a new V1SpectroClustersCustomValidateParams object
// with the default values initialized.
func NewV1SpectroClustersCustomValidateParams() *V1SpectroClustersCustomValidateParams {
	var ()
	return &V1SpectroClustersCustomValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersCustomValidateParamsWithTimeout creates a new V1SpectroClustersCustomValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersCustomValidateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersCustomValidateParams {
	var ()
	return &V1SpectroClustersCustomValidateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersCustomValidateParamsWithContext creates a new V1SpectroClustersCustomValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersCustomValidateParamsWithContext(ctx context.Context) *V1SpectroClustersCustomValidateParams {
	var ()
	return &V1SpectroClustersCustomValidateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersCustomValidateParamsWithHTTPClient creates a new V1SpectroClustersCustomValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersCustomValidateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersCustomValidateParams {
	var ()
	return &V1SpectroClustersCustomValidateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersCustomValidateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters custom validate operation typically these are written to a http.Request
*/
type V1SpectroClustersCustomValidateParams struct {

	/*Body*/
	Body *models.V1SpectroCustomClusterEntity
	/*CloudType
	  Cluster's cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersCustomValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) WithContext(ctx context.Context) *V1SpectroClustersCustomValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersCustomValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) WithBody(body *models.V1SpectroCustomClusterEntity) *V1SpectroClustersCustomValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) SetBody(body *models.V1SpectroCustomClusterEntity) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) WithCloudType(cloudType string) *V1SpectroClustersCustomValidateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 spectro clusters custom validate params
func (o *V1SpectroClustersCustomValidateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersCustomValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
