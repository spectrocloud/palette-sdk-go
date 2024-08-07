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

// NewV1SpectroClustersTkeRateParams creates a new V1SpectroClustersTkeRateParams object
// with the default values initialized.
func NewV1SpectroClustersTkeRateParams() *V1SpectroClustersTkeRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersTkeRateParams{
		PeriodType: &periodTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersTkeRateParamsWithTimeout creates a new V1SpectroClustersTkeRateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersTkeRateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersTkeRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersTkeRateParams{
		PeriodType: &periodTypeDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersTkeRateParamsWithContext creates a new V1SpectroClustersTkeRateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersTkeRateParamsWithContext(ctx context.Context) *V1SpectroClustersTkeRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersTkeRateParams{
		PeriodType: &periodTypeDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersTkeRateParamsWithHTTPClient creates a new V1SpectroClustersTkeRateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersTkeRateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersTkeRateParams {
	var (
		periodTypeDefault = string("hourly")
	)
	return &V1SpectroClustersTkeRateParams{
		PeriodType: &periodTypeDefault,
		HTTPClient: client,
	}
}

/*V1SpectroClustersTkeRateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters tke rate operation typically these are written to a http.Request
*/
type V1SpectroClustersTkeRateParams struct {

	/*Body*/
	Body *models.V1SpectroTencentClusterRateEntity
	/*PeriodType*/
	PeriodType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersTkeRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) WithContext(ctx context.Context) *V1SpectroClustersTkeRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersTkeRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) WithBody(body *models.V1SpectroTencentClusterRateEntity) *V1SpectroClustersTkeRateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) SetBody(body *models.V1SpectroTencentClusterRateEntity) {
	o.Body = body
}

// WithPeriodType adds the periodType to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) WithPeriodType(periodType *string) *V1SpectroClustersTkeRateParams {
	o.SetPeriodType(periodType)
	return o
}

// SetPeriodType adds the periodType to the v1 spectro clusters tke rate params
func (o *V1SpectroClustersTkeRateParams) SetPeriodType(periodType *string) {
	o.PeriodType = periodType
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersTkeRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
