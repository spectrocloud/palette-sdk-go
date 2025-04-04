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

// NewV1OverlordsUIDMaasCloudConfigUpdateParams creates a new V1OverlordsUIDMaasCloudConfigUpdateParams object
// with the default values initialized.
func NewV1OverlordsUIDMaasCloudConfigUpdateParams() *V1OverlordsUIDMaasCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDMaasCloudConfigUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDMaasCloudConfigUpdateParamsWithTimeout creates a new V1OverlordsUIDMaasCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDMaasCloudConfigUpdateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDMaasCloudConfigUpdateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDMaasCloudConfigUpdateParamsWithContext creates a new V1OverlordsUIDMaasCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDMaasCloudConfigUpdateParamsWithContext(ctx context.Context) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDMaasCloudConfigUpdateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDMaasCloudConfigUpdateParamsWithHTTPClient creates a new V1OverlordsUIDMaasCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDMaasCloudConfigUpdateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDMaasCloudConfigUpdateParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDMaasCloudConfigUpdateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid maas cloud config update operation typically these are written to a http.Request
*/
type V1OverlordsUIDMaasCloudConfigUpdateParams struct {

	/*Body*/
	Body *models.V1OverlordMaasCloudConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) WithContext(ctx context.Context) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) WithBody(body *models.V1OverlordMaasCloudConfig) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) SetBody(body *models.V1OverlordMaasCloudConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) WithUID(uid string) *V1OverlordsUIDMaasCloudConfigUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid maas cloud config update params
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDMaasCloudConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
