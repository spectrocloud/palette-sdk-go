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

// NewV1SpectroClustersGcpRateParams creates a new V1SpectroClustersGcpRateParams object
// with the default values initialized.
func NewV1SpectroClustersGcpRateParams() *V1SpectroClustersGcpRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersGcpRateParams{
		PeriodType: &periodTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersGcpRateParamsWithTimeout creates a new V1SpectroClustersGcpRateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersGcpRateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersGcpRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersGcpRateParams{
		PeriodType: &periodTypeDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersGcpRateParamsWithContext creates a new V1SpectroClustersGcpRateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersGcpRateParamsWithContext(ctx context.Context) *V1SpectroClustersGcpRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersGcpRateParams{
		PeriodType: &periodTypeDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersGcpRateParamsWithHTTPClient creates a new V1SpectroClustersGcpRateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersGcpRateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersGcpRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersGcpRateParams{
		PeriodType: &periodTypeDefault,
		HTTPClient: client,
	}
}

/*
V1SpectroClustersGcpRateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters gcp rate operation typically these are written to a http.Request
*/
type V1SpectroClustersGcpRateParams struct {

	/*Body*/
	Body *models.V1SpectroGcpClusterRateEntity
	/*PeriodType*/
	PeriodType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersGcpRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) WithContext(ctx context.Context) *V1SpectroClustersGcpRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersGcpRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) WithBody(body *models.V1SpectroGcpClusterRateEntity) *V1SpectroClustersGcpRateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) SetBody(body *models.V1SpectroGcpClusterRateEntity) {
	o.Body = body
}

// WithPeriodType adds the periodType to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) WithPeriodType(periodType *string) *V1SpectroClustersGcpRateParams {
	o.SetPeriodType(periodType)
	return o
}

// SetPeriodType adds the periodType to the v1 spectro clusters gcp rate params
func (o *V1SpectroClustersGcpRateParams) SetPeriodType(periodType *string) {
	o.PeriodType = periodType
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersGcpRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.PeriodType != nil {

		// query param periodType
		var qrPeriodType string
		if o.PeriodType != nil {
			qrPeriodType = *o.PeriodType
		}
		qPeriodType := qrPeriodType
		if qPeriodType != "" {
			if err := r.SetQueryParam("periodType", qPeriodType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
