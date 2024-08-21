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

// NewV1ClusterFeatureComplianceScanCreateParams creates a new V1ClusterFeatureComplianceScanCreateParams object
// with the default values initialized.
func NewV1ClusterFeatureComplianceScanCreateParams() *V1ClusterFeatureComplianceScanCreateParams {
	var ()
	return &V1ClusterFeatureComplianceScanCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureComplianceScanCreateParamsWithTimeout creates a new V1ClusterFeatureComplianceScanCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureComplianceScanCreateParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureComplianceScanCreateParams {
	var ()
	return &V1ClusterFeatureComplianceScanCreateParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureComplianceScanCreateParamsWithContext creates a new V1ClusterFeatureComplianceScanCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureComplianceScanCreateParamsWithContext(ctx context.Context) *V1ClusterFeatureComplianceScanCreateParams {
	var ()
	return &V1ClusterFeatureComplianceScanCreateParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureComplianceScanCreateParamsWithHTTPClient creates a new V1ClusterFeatureComplianceScanCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureComplianceScanCreateParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureComplianceScanCreateParams {
	var ()
	return &V1ClusterFeatureComplianceScanCreateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterFeatureComplianceScanCreateParams contains all the parameters to send to the API endpoint
for the v1 cluster feature compliance scan create operation typically these are written to a http.Request
*/
type V1ClusterFeatureComplianceScanCreateParams struct {

	/*Body*/
	Body *models.V1ClusterComplianceScheduleConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureComplianceScanCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) WithContext(ctx context.Context) *V1ClusterFeatureComplianceScanCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureComplianceScanCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) WithBody(body *models.V1ClusterComplianceScheduleConfig) *V1ClusterFeatureComplianceScanCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) SetBody(body *models.V1ClusterComplianceScheduleConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) WithUID(uid string) *V1ClusterFeatureComplianceScanCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature compliance scan create params
func (o *V1ClusterFeatureComplianceScanCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureComplianceScanCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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