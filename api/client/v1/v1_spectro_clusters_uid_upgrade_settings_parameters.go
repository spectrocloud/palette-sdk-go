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

// NewV1SpectroClustersUIDUpgradeSettingsParams creates a new V1SpectroClustersUIDUpgradeSettingsParams object
// with the default values initialized.
func NewV1SpectroClustersUIDUpgradeSettingsParams() *V1SpectroClustersUIDUpgradeSettingsParams {
	var ()
	return &V1SpectroClustersUIDUpgradeSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDUpgradeSettingsParamsWithTimeout creates a new V1SpectroClustersUIDUpgradeSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDUpgradeSettingsParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDUpgradeSettingsParams {
	var ()
	return &V1SpectroClustersUIDUpgradeSettingsParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDUpgradeSettingsParamsWithContext creates a new V1SpectroClustersUIDUpgradeSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDUpgradeSettingsParamsWithContext(ctx context.Context) *V1SpectroClustersUIDUpgradeSettingsParams {
	var ()
	return &V1SpectroClustersUIDUpgradeSettingsParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDUpgradeSettingsParamsWithHTTPClient creates a new V1SpectroClustersUIDUpgradeSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDUpgradeSettingsParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDUpgradeSettingsParams {
	var ()
	return &V1SpectroClustersUIDUpgradeSettingsParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDUpgradeSettingsParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid upgrade settings operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDUpgradeSettingsParams struct {

	/*Body*/
	Body *models.V1ClusterUpgradeSettingsEntity
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDUpgradeSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) WithContext(ctx context.Context) *V1SpectroClustersUIDUpgradeSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDUpgradeSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) WithBody(body *models.V1ClusterUpgradeSettingsEntity) *V1SpectroClustersUIDUpgradeSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) SetBody(body *models.V1ClusterUpgradeSettingsEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) WithUID(uid string) *V1SpectroClustersUIDUpgradeSettingsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid upgrade settings params
func (o *V1SpectroClustersUIDUpgradeSettingsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDUpgradeSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
