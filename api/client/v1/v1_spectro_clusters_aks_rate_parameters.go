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

// NewV1SpectroClustersAksRateParams creates a new V1SpectroClustersAksRateParams object
// with the default values initialized.
func NewV1SpectroClustersAksRateParams() *V1SpectroClustersAksRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersAksRateParams{
		PeriodType: &periodTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersAksRateParamsWithTimeout creates a new V1SpectroClustersAksRateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersAksRateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersAksRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersAksRateParams{
		PeriodType: &periodTypeDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersAksRateParamsWithContext creates a new V1SpectroClustersAksRateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersAksRateParamsWithContext(ctx context.Context) *V1SpectroClustersAksRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersAksRateParams{
		PeriodType: &periodTypeDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersAksRateParamsWithHTTPClient creates a new V1SpectroClustersAksRateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersAksRateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersAksRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersAksRateParams{
		PeriodType: &periodTypeDefault,
		HTTPClient: client,
	}
}

/*V1SpectroClustersAksRateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters aks rate operation typically these are written to a http.Request
*/
type V1SpectroClustersAksRateParams struct {

	/*Body*/
	Body *models.V1SpectroAzureClusterRateEntity
	/*PeriodType*/
	PeriodType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersAksRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) WithContext(ctx context.Context) *V1SpectroClustersAksRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersAksRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) WithBody(body *models.V1SpectroAzureClusterRateEntity) *V1SpectroClustersAksRateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) SetBody(body *models.V1SpectroAzureClusterRateEntity) {
	o.Body = body
}

// WithPeriodType adds the periodType to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) WithPeriodType(periodType *string) *V1SpectroClustersAksRateParams {
	o.SetPeriodType(periodType)
	return o
}

// SetPeriodType adds the periodType to the v1 spectro clusters aks rate params
func (o *V1SpectroClustersAksRateParams) SetPeriodType(periodType *string) {
	o.PeriodType = periodType
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersAksRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
