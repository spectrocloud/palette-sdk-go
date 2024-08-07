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

// NewV1ClusterFeatureScanKubeHunterLogUpdateParams creates a new V1ClusterFeatureScanKubeHunterLogUpdateParams object
// with the default values initialized.
func NewV1ClusterFeatureScanKubeHunterLogUpdateParams() *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeHunterLogUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureScanKubeHunterLogUpdateParamsWithTimeout creates a new V1ClusterFeatureScanKubeHunterLogUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureScanKubeHunterLogUpdateParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeHunterLogUpdateParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureScanKubeHunterLogUpdateParamsWithContext creates a new V1ClusterFeatureScanKubeHunterLogUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureScanKubeHunterLogUpdateParamsWithContext(ctx context.Context) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeHunterLogUpdateParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureScanKubeHunterLogUpdateParamsWithHTTPClient creates a new V1ClusterFeatureScanKubeHunterLogUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureScanKubeHunterLogUpdateParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanKubeHunterLogUpdateParams{
		HTTPClient: client,
	}
}

/*V1ClusterFeatureScanKubeHunterLogUpdateParams contains all the parameters to send to the API endpoint
for the v1 cluster feature scan kube hunter log update operation typically these are written to a http.Request
*/
type V1ClusterFeatureScanKubeHunterLogUpdateParams struct {

	/*Body*/
	Body *models.V1KubeHunterEntity
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) WithContext(ctx context.Context) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) WithBody(body *models.V1KubeHunterEntity) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) SetBody(body *models.V1KubeHunterEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) WithUID(uid string) *V1ClusterFeatureScanKubeHunterLogUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature scan kube hunter log update params
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureScanKubeHunterLogUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
