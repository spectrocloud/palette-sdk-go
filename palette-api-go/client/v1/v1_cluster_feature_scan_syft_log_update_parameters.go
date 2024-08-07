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

// NewV1ClusterFeatureScanSyftLogUpdateParams creates a new V1ClusterFeatureScanSyftLogUpdateParams object
// with the default values initialized.
func NewV1ClusterFeatureScanSyftLogUpdateParams() *V1ClusterFeatureScanSyftLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanSyftLogUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureScanSyftLogUpdateParamsWithTimeout creates a new V1ClusterFeatureScanSyftLogUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureScanSyftLogUpdateParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureScanSyftLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanSyftLogUpdateParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureScanSyftLogUpdateParamsWithContext creates a new V1ClusterFeatureScanSyftLogUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureScanSyftLogUpdateParamsWithContext(ctx context.Context) *V1ClusterFeatureScanSyftLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanSyftLogUpdateParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureScanSyftLogUpdateParamsWithHTTPClient creates a new V1ClusterFeatureScanSyftLogUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureScanSyftLogUpdateParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureScanSyftLogUpdateParams {
	var ()
	return &V1ClusterFeatureScanSyftLogUpdateParams{
		HTTPClient: client,
	}
}

/*V1ClusterFeatureScanSyftLogUpdateParams contains all the parameters to send to the API endpoint
for the v1 cluster feature scan syft log update operation typically these are written to a http.Request
*/
type V1ClusterFeatureScanSyftLogUpdateParams struct {

	/*Body*/
	Body *models.V1SyftEntity
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureScanSyftLogUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) WithContext(ctx context.Context) *V1ClusterFeatureScanSyftLogUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureScanSyftLogUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) WithBody(body *models.V1SyftEntity) *V1ClusterFeatureScanSyftLogUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) SetBody(body *models.V1SyftEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) WithUID(uid string) *V1ClusterFeatureScanSyftLogUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature scan syft log update params
func (o *V1ClusterFeatureScanSyftLogUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureScanSyftLogUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
