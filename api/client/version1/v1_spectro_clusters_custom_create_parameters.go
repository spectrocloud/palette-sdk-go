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

// NewV1SpectroClustersCustomCreateParams creates a new V1SpectroClustersCustomCreateParams object
// with the default values initialized.
func NewV1SpectroClustersCustomCreateParams() *V1SpectroClustersCustomCreateParams {
	var ()
	return &V1SpectroClustersCustomCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersCustomCreateParamsWithTimeout creates a new V1SpectroClustersCustomCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersCustomCreateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersCustomCreateParams {
	var ()
	return &V1SpectroClustersCustomCreateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersCustomCreateParamsWithContext creates a new V1SpectroClustersCustomCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersCustomCreateParamsWithContext(ctx context.Context) *V1SpectroClustersCustomCreateParams {
	var ()
	return &V1SpectroClustersCustomCreateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersCustomCreateParamsWithHTTPClient creates a new V1SpectroClustersCustomCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersCustomCreateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersCustomCreateParams {
	var ()
	return &V1SpectroClustersCustomCreateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersCustomCreateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters custom create operation typically these are written to a http.Request
*/
type V1SpectroClustersCustomCreateParams struct {

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

// WithTimeout adds the timeout to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersCustomCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) WithContext(ctx context.Context) *V1SpectroClustersCustomCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersCustomCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) WithBody(body *models.V1SpectroCustomClusterEntity) *V1SpectroClustersCustomCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) SetBody(body *models.V1SpectroCustomClusterEntity) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) WithCloudType(cloudType string) *V1SpectroClustersCustomCreateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 spectro clusters custom create params
func (o *V1SpectroClustersCustomCreateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersCustomCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
