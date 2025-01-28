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

// NewV1ClusterFeatureBackupOnDemandCreateParams creates a new V1ClusterFeatureBackupOnDemandCreateParams object
// with the default values initialized.
func NewV1ClusterFeatureBackupOnDemandCreateParams() *V1ClusterFeatureBackupOnDemandCreateParams {
	var ()
	return &V1ClusterFeatureBackupOnDemandCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureBackupOnDemandCreateParamsWithTimeout creates a new V1ClusterFeatureBackupOnDemandCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureBackupOnDemandCreateParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureBackupOnDemandCreateParams {
	var ()
	return &V1ClusterFeatureBackupOnDemandCreateParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureBackupOnDemandCreateParamsWithContext creates a new V1ClusterFeatureBackupOnDemandCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureBackupOnDemandCreateParamsWithContext(ctx context.Context) *V1ClusterFeatureBackupOnDemandCreateParams {
	var ()
	return &V1ClusterFeatureBackupOnDemandCreateParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureBackupOnDemandCreateParamsWithHTTPClient creates a new V1ClusterFeatureBackupOnDemandCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureBackupOnDemandCreateParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureBackupOnDemandCreateParams {
	var ()
	return &V1ClusterFeatureBackupOnDemandCreateParams{
		HTTPClient: client,
	}
}

/*V1ClusterFeatureBackupOnDemandCreateParams contains all the parameters to send to the API endpoint
for the v1 cluster feature backup on demand create operation typically these are written to a http.Request
*/
type V1ClusterFeatureBackupOnDemandCreateParams struct {

	/*Body*/
	Body *models.V1ClusterBackupConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureBackupOnDemandCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) WithContext(ctx context.Context) *V1ClusterFeatureBackupOnDemandCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureBackupOnDemandCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) WithBody(body *models.V1ClusterBackupConfig) *V1ClusterFeatureBackupOnDemandCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) SetBody(body *models.V1ClusterBackupConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) WithUID(uid string) *V1ClusterFeatureBackupOnDemandCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature backup on demand create params
func (o *V1ClusterFeatureBackupOnDemandCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureBackupOnDemandCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
