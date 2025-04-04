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

// NewV1ClusterFeatureComplianceScanUpdateParams creates a new V1ClusterFeatureComplianceScanUpdateParams object
// with the default values initialized.
func NewV1ClusterFeatureComplianceScanUpdateParams() *V1ClusterFeatureComplianceScanUpdateParams {
	var ()
	return &V1ClusterFeatureComplianceScanUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureComplianceScanUpdateParamsWithTimeout creates a new V1ClusterFeatureComplianceScanUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureComplianceScanUpdateParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureComplianceScanUpdateParams {
	var ()
	return &V1ClusterFeatureComplianceScanUpdateParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureComplianceScanUpdateParamsWithContext creates a new V1ClusterFeatureComplianceScanUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureComplianceScanUpdateParamsWithContext(ctx context.Context) *V1ClusterFeatureComplianceScanUpdateParams {
	var ()
	return &V1ClusterFeatureComplianceScanUpdateParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureComplianceScanUpdateParamsWithHTTPClient creates a new V1ClusterFeatureComplianceScanUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureComplianceScanUpdateParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureComplianceScanUpdateParams {
	var ()
	return &V1ClusterFeatureComplianceScanUpdateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterFeatureComplianceScanUpdateParams contains all the parameters to send to the API endpoint
for the v1 cluster feature compliance scan update operation typically these are written to a http.Request
*/
type V1ClusterFeatureComplianceScanUpdateParams struct {

	/*Body*/
	Body *models.V1ClusterComplianceScheduleConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureComplianceScanUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) WithContext(ctx context.Context) *V1ClusterFeatureComplianceScanUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureComplianceScanUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) WithBody(body *models.V1ClusterComplianceScheduleConfig) *V1ClusterFeatureComplianceScanUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) SetBody(body *models.V1ClusterComplianceScheduleConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) WithUID(uid string) *V1ClusterFeatureComplianceScanUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature compliance scan update params
func (o *V1ClusterFeatureComplianceScanUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureComplianceScanUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
