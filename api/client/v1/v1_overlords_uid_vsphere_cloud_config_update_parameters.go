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

// NewV1OverlordsUIDVsphereCloudConfigUpdateParams creates a new V1OverlordsUIDVsphereCloudConfigUpdateParams object
// with the default values initialized.
func NewV1OverlordsUIDVsphereCloudConfigUpdateParams() *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereCloudConfigUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDVsphereCloudConfigUpdateParamsWithTimeout creates a new V1OverlordsUIDVsphereCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDVsphereCloudConfigUpdateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereCloudConfigUpdateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDVsphereCloudConfigUpdateParamsWithContext creates a new V1OverlordsUIDVsphereCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDVsphereCloudConfigUpdateParamsWithContext(ctx context.Context) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereCloudConfigUpdateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDVsphereCloudConfigUpdateParamsWithHTTPClient creates a new V1OverlordsUIDVsphereCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDVsphereCloudConfigUpdateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereCloudConfigUpdateParams{
		HTTPClient: client,
	}
}

/*V1OverlordsUIDVsphereCloudConfigUpdateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid vsphere cloud config update operation typically these are written to a http.Request
*/
type V1OverlordsUIDVsphereCloudConfigUpdateParams struct {

	/*Body*/
	Body *models.V1OverlordVsphereCloudConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) WithContext(ctx context.Context) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) WithBody(body *models.V1OverlordVsphereCloudConfig) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) SetBody(body *models.V1OverlordVsphereCloudConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) WithUID(uid string) *V1OverlordsUIDVsphereCloudConfigUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid vsphere cloud config update params
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDVsphereCloudConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
