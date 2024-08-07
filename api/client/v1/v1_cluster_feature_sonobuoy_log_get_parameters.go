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
)

// NewV1ClusterFeatureSonobuoyLogGetParams creates a new V1ClusterFeatureSonobuoyLogGetParams object
// with the default values initialized.
func NewV1ClusterFeatureSonobuoyLogGetParams() *V1ClusterFeatureSonobuoyLogGetParams {
	var ()
	return &V1ClusterFeatureSonobuoyLogGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureSonobuoyLogGetParamsWithTimeout creates a new V1ClusterFeatureSonobuoyLogGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureSonobuoyLogGetParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureSonobuoyLogGetParams {
	var ()
	return &V1ClusterFeatureSonobuoyLogGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureSonobuoyLogGetParamsWithContext creates a new V1ClusterFeatureSonobuoyLogGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureSonobuoyLogGetParamsWithContext(ctx context.Context) *V1ClusterFeatureSonobuoyLogGetParams {
	var ()
	return &V1ClusterFeatureSonobuoyLogGetParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureSonobuoyLogGetParamsWithHTTPClient creates a new V1ClusterFeatureSonobuoyLogGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureSonobuoyLogGetParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureSonobuoyLogGetParams {
	var ()
	return &V1ClusterFeatureSonobuoyLogGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterFeatureSonobuoyLogGetParams contains all the parameters to send to the API endpoint
for the v1 cluster feature sonobuoy log get operation typically these are written to a http.Request
*/
type V1ClusterFeatureSonobuoyLogGetParams struct {

	/*LogUID*/
	LogUID string
	/*ReportID*/
	ReportID *string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureSonobuoyLogGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) WithContext(ctx context.Context) *V1ClusterFeatureSonobuoyLogGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureSonobuoyLogGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogUID adds the logUID to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) WithLogUID(logUID string) *V1ClusterFeatureSonobuoyLogGetParams {
	o.SetLogUID(logUID)
	return o
}

// SetLogUID adds the logUid to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) SetLogUID(logUID string) {
	o.LogUID = logUID
}

// WithReportID adds the reportID to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) WithReportID(reportID *string) *V1ClusterFeatureSonobuoyLogGetParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) SetReportID(reportID *string) {
	o.ReportID = reportID
}

// WithUID adds the uid to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) WithUID(uid string) *V1ClusterFeatureSonobuoyLogGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature sonobuoy log get params
func (o *V1ClusterFeatureSonobuoyLogGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureSonobuoyLogGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param logUid
	if err := r.SetPathParam("logUid", o.LogUID); err != nil {
		return err
	}

	if o.ReportID != nil {

		// query param reportId
		var qrReportID string
		if o.ReportID != nil {
			qrReportID = *o.ReportID
		}
		qReportID := qrReportID
		if qReportID != "" {
			if err := r.SetQueryParam("reportId", qReportID); err != nil {
				return err
			}
		}

	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
