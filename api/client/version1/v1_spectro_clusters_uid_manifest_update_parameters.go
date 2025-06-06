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

// NewV1SpectroClustersUIDManifestUpdateParams creates a new V1SpectroClustersUIDManifestUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersUIDManifestUpdateParams() *V1SpectroClustersUIDManifestUpdateParams {
	var ()
	return &V1SpectroClustersUIDManifestUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDManifestUpdateParamsWithTimeout creates a new V1SpectroClustersUIDManifestUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDManifestUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDManifestUpdateParams {
	var ()
	return &V1SpectroClustersUIDManifestUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDManifestUpdateParamsWithContext creates a new V1SpectroClustersUIDManifestUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDManifestUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersUIDManifestUpdateParams {
	var ()
	return &V1SpectroClustersUIDManifestUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDManifestUpdateParamsWithHTTPClient creates a new V1SpectroClustersUIDManifestUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDManifestUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDManifestUpdateParams {
	var ()
	return &V1SpectroClustersUIDManifestUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDManifestUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid manifest update operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDManifestUpdateParams struct {

	/*Body*/
	Body *models.V1SpectroClusterAssetManifest
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDManifestUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersUIDManifestUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDManifestUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) WithBody(body *models.V1SpectroClusterAssetManifest) *V1SpectroClustersUIDManifestUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) SetBody(body *models.V1SpectroClusterAssetManifest) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) WithUID(uid string) *V1SpectroClustersUIDManifestUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid manifest update params
func (o *V1SpectroClustersUIDManifestUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDManifestUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
