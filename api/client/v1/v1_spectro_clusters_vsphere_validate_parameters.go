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

// NewV1SpectroClustersVsphereValidateParams creates a new V1SpectroClustersVsphereValidateParams object
// with the default values initialized.
func NewV1SpectroClustersVsphereValidateParams() *V1SpectroClustersVsphereValidateParams {
	var ()
	return &V1SpectroClustersVsphereValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVsphereValidateParamsWithTimeout creates a new V1SpectroClustersVsphereValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVsphereValidateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVsphereValidateParams {
	var ()
	return &V1SpectroClustersVsphereValidateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersVsphereValidateParamsWithContext creates a new V1SpectroClustersVsphereValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVsphereValidateParamsWithContext(ctx context.Context) *V1SpectroClustersVsphereValidateParams {
	var ()
	return &V1SpectroClustersVsphereValidateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersVsphereValidateParamsWithHTTPClient creates a new V1SpectroClustersVsphereValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVsphereValidateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVsphereValidateParams {
	var ()
	return &V1SpectroClustersVsphereValidateParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersVsphereValidateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters vsphere validate operation typically these are written to a http.Request
*/
type V1SpectroClustersVsphereValidateParams struct {

	/*Body*/
	Body *models.V1SpectroVsphereClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVsphereValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) WithContext(ctx context.Context) *V1SpectroClustersVsphereValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVsphereValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) WithBody(body *models.V1SpectroVsphereClusterEntity) *V1SpectroClustersVsphereValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters vsphere validate params
func (o *V1SpectroClustersVsphereValidateParams) SetBody(body *models.V1SpectroVsphereClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVsphereValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
