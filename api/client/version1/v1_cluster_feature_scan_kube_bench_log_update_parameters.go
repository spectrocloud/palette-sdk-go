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

// NewV1ClusterFeatureScanKubeBenchLogUpdateParams creates a new V1ClusterFeatureScanKubeBenchLogUpdateParams object
// with the default values initialized.
func NewV1ClusterFeatureScanKubeBenchLogUpdateParams() *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeBenchLogUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureScanKubeBenchLogUpdateParamsWithTimeout creates a new V1ClusterFeatureScanKubeBenchLogUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureScanKubeBenchLogUpdateParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeBenchLogUpdateParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureScanKubeBenchLogUpdateParamsWithContext creates a new V1ClusterFeatureScanKubeBenchLogUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureScanKubeBenchLogUpdateParamsWithContext(ctx context.Context) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeBenchLogUpdateParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureScanKubeBenchLogUpdateParamsWithHTTPClient creates a new V1ClusterFeatureScanKubeBenchLogUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureScanKubeBenchLogUpdateParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeBenchLogUpdateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterFeatureScanKubeBenchLogUpdateParams contains all the parameters to send to the API endpoint
for the v1 cluster feature scan kube bench log update operation typically these are written to a http.Request
*/
type V1ClusterFeatureScanKubeBenchLogUpdateParams struct {

	/*Body*/
	Body *models.V1KubeBenchEntity
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) WithContext(ctx context.Context) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) WithBody(body *models.V1KubeBenchEntity) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) SetBody(body *models.V1KubeBenchEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) WithUID(uid string) *V1ClusterFeatureScanKubeBenchLogUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature scan kube bench log update params
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureScanKubeBenchLogUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
