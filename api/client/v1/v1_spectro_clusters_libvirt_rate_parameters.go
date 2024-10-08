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

// NewV1SpectroClustersLibvirtRateParams creates a new V1SpectroClustersLibvirtRateParams object
// with the default values initialized.
func NewV1SpectroClustersLibvirtRateParams() *V1SpectroClustersLibvirtRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersLibvirtRateParams{
		PeriodType: &periodTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersLibvirtRateParamsWithTimeout creates a new V1SpectroClustersLibvirtRateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersLibvirtRateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersLibvirtRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersLibvirtRateParams{
		PeriodType: &periodTypeDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersLibvirtRateParamsWithContext creates a new V1SpectroClustersLibvirtRateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersLibvirtRateParamsWithContext(ctx context.Context) *V1SpectroClustersLibvirtRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersLibvirtRateParams{
		PeriodType: &periodTypeDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersLibvirtRateParamsWithHTTPClient creates a new V1SpectroClustersLibvirtRateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersLibvirtRateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersLibvirtRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersLibvirtRateParams{
		PeriodType: &periodTypeDefault,
		HTTPClient: client,
	}
}

/*
V1SpectroClustersLibvirtRateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters libvirt rate operation typically these are written to a http.Request
*/
type V1SpectroClustersLibvirtRateParams struct {

	/*Body*/
	Body *models.V1SpectroLibvirtClusterRateEntity
	/*PeriodType*/
	PeriodType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersLibvirtRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) WithContext(ctx context.Context) *V1SpectroClustersLibvirtRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersLibvirtRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) WithBody(body *models.V1SpectroLibvirtClusterRateEntity) *V1SpectroClustersLibvirtRateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) SetBody(body *models.V1SpectroLibvirtClusterRateEntity) {
	o.Body = body
}

// WithPeriodType adds the periodType to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) WithPeriodType(periodType *string) *V1SpectroClustersLibvirtRateParams {
	o.SetPeriodType(periodType)
	return o
}

// SetPeriodType adds the periodType to the v1 spectro clusters libvirt rate params
func (o *V1SpectroClustersLibvirtRateParams) SetPeriodType(periodType *string) {
	o.PeriodType = periodType
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersLibvirtRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
