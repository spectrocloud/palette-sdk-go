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

// NewV1SpectroClustersVirtualValidateParams creates a new V1SpectroClustersVirtualValidateParams object
// with the default values initialized.
func NewV1SpectroClustersVirtualValidateParams() *V1SpectroClustersVirtualValidateParams {
	var ()
	return &V1SpectroClustersVirtualValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersVirtualValidateParamsWithTimeout creates a new V1SpectroClustersVirtualValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersVirtualValidateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersVirtualValidateParams {
	var ()
	return &V1SpectroClustersVirtualValidateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersVirtualValidateParamsWithContext creates a new V1SpectroClustersVirtualValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersVirtualValidateParamsWithContext(ctx context.Context) *V1SpectroClustersVirtualValidateParams {
	var ()
	return &V1SpectroClustersVirtualValidateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersVirtualValidateParamsWithHTTPClient creates a new V1SpectroClustersVirtualValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersVirtualValidateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersVirtualValidateParams {
	var ()
	return &V1SpectroClustersVirtualValidateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersVirtualValidateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters virtual validate operation typically these are written to a http.Request
*/
type V1SpectroClustersVirtualValidateParams struct {

	/*Body*/
	Body *models.V1SpectroVirtualClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersVirtualValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) WithContext(ctx context.Context) *V1SpectroClustersVirtualValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersVirtualValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) WithBody(body *models.V1SpectroVirtualClusterEntity) *V1SpectroClustersVirtualValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters virtual validate params
func (o *V1SpectroClustersVirtualValidateParams) SetBody(body *models.V1SpectroVirtualClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersVirtualValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
