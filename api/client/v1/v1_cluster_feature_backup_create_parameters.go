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

// NewV1ClusterFeatureBackupCreateParams creates a new V1ClusterFeatureBackupCreateParams object
// with the default values initialized.
func NewV1ClusterFeatureBackupCreateParams() *V1ClusterFeatureBackupCreateParams {
	var ()
	return &V1ClusterFeatureBackupCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterFeatureBackupCreateParamsWithTimeout creates a new V1ClusterFeatureBackupCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterFeatureBackupCreateParamsWithTimeout(timeout time.Duration) *V1ClusterFeatureBackupCreateParams {
	var ()
	return &V1ClusterFeatureBackupCreateParams{

		timeout: timeout,
	}
}

// NewV1ClusterFeatureBackupCreateParamsWithContext creates a new V1ClusterFeatureBackupCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterFeatureBackupCreateParamsWithContext(ctx context.Context) *V1ClusterFeatureBackupCreateParams {
	var ()
	return &V1ClusterFeatureBackupCreateParams{

		Context: ctx,
	}
}

// NewV1ClusterFeatureBackupCreateParamsWithHTTPClient creates a new V1ClusterFeatureBackupCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterFeatureBackupCreateParamsWithHTTPClient(client *http.Client) *V1ClusterFeatureBackupCreateParams {
	var ()
	return &V1ClusterFeatureBackupCreateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterFeatureBackupCreateParams contains all the parameters to send to the API endpoint
for the v1 cluster feature backup create operation typically these are written to a http.Request
*/
type V1ClusterFeatureBackupCreateParams struct {

	/*Body*/
	Body *models.V1ClusterBackupConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) WithTimeout(timeout time.Duration) *V1ClusterFeatureBackupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) WithContext(ctx context.Context) *V1ClusterFeatureBackupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) WithHTTPClient(client *http.Client) *V1ClusterFeatureBackupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) WithBody(body *models.V1ClusterBackupConfig) *V1ClusterFeatureBackupCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) SetBody(body *models.V1ClusterBackupConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) WithUID(uid string) *V1ClusterFeatureBackupCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster feature backup create params
func (o *V1ClusterFeatureBackupCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterFeatureBackupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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