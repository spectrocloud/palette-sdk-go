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

// NewV1SpectroClustersUIDTokenKubeConfigUpdateParams creates a new V1SpectroClustersUIDTokenKubeConfigUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersUIDTokenKubeConfigUpdateParams() *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDTokenKubeConfigUpdateParamsWithTimeout creates a new V1SpectroClustersUIDTokenKubeConfigUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDTokenKubeConfigUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDTokenKubeConfigUpdateParamsWithContext creates a new V1SpectroClustersUIDTokenKubeConfigUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDTokenKubeConfigUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDTokenKubeConfigUpdateParamsWithHTTPClient creates a new V1SpectroClustersUIDTokenKubeConfigUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDTokenKubeConfigUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDTokenKubeConfigUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid token kube config update operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDTokenKubeConfigUpdateParams struct {

	/*Body*/
	Body *models.V1SpectroClusterAssetTokenKubeConfig
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) WithBody(body *models.V1SpectroClusterAssetTokenKubeConfig) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) SetBody(body *models.V1SpectroClusterAssetTokenKubeConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) WithUID(uid string) *V1SpectroClustersUIDTokenKubeConfigUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid token kube config update params
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDTokenKubeConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
