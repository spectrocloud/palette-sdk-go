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

// NewV1DashboardClustersSearchSummaryExportParams creates a new V1DashboardClustersSearchSummaryExportParams object
// with the default values initialized.
func NewV1DashboardClustersSearchSummaryExportParams() *V1DashboardClustersSearchSummaryExportParams {
	var (
		formatDefault = string("csv")
	)
	return &V1DashboardClustersSearchSummaryExportParams{
		Format: &formatDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1DashboardClustersSearchSummaryExportParamsWithTimeout creates a new V1DashboardClustersSearchSummaryExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1DashboardClustersSearchSummaryExportParamsWithTimeout(timeout time.Duration) *V1DashboardClustersSearchSummaryExportParams {
	var (
		formatDefault = string("csv")
	)
	return &V1DashboardClustersSearchSummaryExportParams{
		Format: &formatDefault,

		timeout: timeout,
	}
}

// NewV1DashboardClustersSearchSummaryExportParamsWithContext creates a new V1DashboardClustersSearchSummaryExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1DashboardClustersSearchSummaryExportParamsWithContext(ctx context.Context) *V1DashboardClustersSearchSummaryExportParams {
	var (
		formatDefault = string("csv")
	)
	return &V1DashboardClustersSearchSummaryExportParams{
		Format: &formatDefault,

		Context: ctx,
	}
}

// NewV1DashboardClustersSearchSummaryExportParamsWithHTTPClient creates a new V1DashboardClustersSearchSummaryExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1DashboardClustersSearchSummaryExportParamsWithHTTPClient(client *http.Client) *V1DashboardClustersSearchSummaryExportParams {
	var (
		formatDefault = string("csv")
	)
	return &V1DashboardClustersSearchSummaryExportParams{
		Format:     &formatDefault,
		HTTPClient: client,
	}
}

/*V1DashboardClustersSearchSummaryExportParams contains all the parameters to send to the API endpoint
for the v1 dashboard clusters search summary export operation typically these are written to a http.Request
*/
type V1DashboardClustersSearchSummaryExportParams struct {

	/*Body*/
	Body *models.V1SearchFilterSummarySpec
	/*Format*/
	Format *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) WithTimeout(timeout time.Duration) *V1DashboardClustersSearchSummaryExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) WithContext(ctx context.Context) *V1DashboardClustersSearchSummaryExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) WithHTTPClient(client *http.Client) *V1DashboardClustersSearchSummaryExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) WithBody(body *models.V1SearchFilterSummarySpec) *V1DashboardClustersSearchSummaryExportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) SetBody(body *models.V1SearchFilterSummarySpec) {
	o.Body = body
}

// WithFormat adds the format to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) WithFormat(format *string) *V1DashboardClustersSearchSummaryExportParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the v1 dashboard clusters search summary export params
func (o *V1DashboardClustersSearchSummaryExportParams) SetFormat(format *string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *V1DashboardClustersSearchSummaryExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
