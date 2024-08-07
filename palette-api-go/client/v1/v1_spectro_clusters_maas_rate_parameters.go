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

// NewV1SpectroClustersMaasRateParams creates a new V1SpectroClustersMaasRateParams object
// with the default values initialized.
func NewV1SpectroClustersMaasRateParams() *V1SpectroClustersMaasRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersMaasRateParams{
		PeriodType: &periodTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersMaasRateParamsWithTimeout creates a new V1SpectroClustersMaasRateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersMaasRateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersMaasRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersMaasRateParams{
		PeriodType: &periodTypeDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersMaasRateParamsWithContext creates a new V1SpectroClustersMaasRateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersMaasRateParamsWithContext(ctx context.Context) *V1SpectroClustersMaasRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersMaasRateParams{
		PeriodType: &periodTypeDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersMaasRateParamsWithHTTPClient creates a new V1SpectroClustersMaasRateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersMaasRateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersMaasRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersMaasRateParams{
		PeriodType: &periodTypeDefault,
		HTTPClient: client,
	}
}

/*V1SpectroClustersMaasRateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters maas rate operation typically these are written to a http.Request
*/
type V1SpectroClustersMaasRateParams struct {

	/*Body*/
	Body *models.V1SpectroMaasClusterRateEntity
	/*PeriodType*/
	PeriodType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersMaasRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) WithContext(ctx context.Context) *V1SpectroClustersMaasRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersMaasRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) WithBody(body *models.V1SpectroMaasClusterRateEntity) *V1SpectroClustersMaasRateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) SetBody(body *models.V1SpectroMaasClusterRateEntity) {
	o.Body = body
}

// WithPeriodType adds the periodType to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) WithPeriodType(periodType *string) *V1SpectroClustersMaasRateParams {
	o.SetPeriodType(periodType)
	return o
}

// SetPeriodType adds the periodType to the v1 spectro clusters maas rate params
func (o *V1SpectroClustersMaasRateParams) SetPeriodType(periodType *string) {
	o.PeriodType = periodType
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersMaasRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
