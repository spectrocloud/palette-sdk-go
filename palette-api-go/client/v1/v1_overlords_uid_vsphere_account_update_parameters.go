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

// NewV1OverlordsUIDVsphereAccountUpdateParams creates a new V1OverlordsUIDVsphereAccountUpdateParams object
// with the default values initialized.
func NewV1OverlordsUIDVsphereAccountUpdateParams() *V1OverlordsUIDVsphereAccountUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDVsphereAccountUpdateParamsWithTimeout creates a new V1OverlordsUIDVsphereAccountUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDVsphereAccountUpdateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereAccountUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountUpdateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDVsphereAccountUpdateParamsWithContext creates a new V1OverlordsUIDVsphereAccountUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDVsphereAccountUpdateParamsWithContext(ctx context.Context) *V1OverlordsUIDVsphereAccountUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountUpdateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDVsphereAccountUpdateParamsWithHTTPClient creates a new V1OverlordsUIDVsphereAccountUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDVsphereAccountUpdateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereAccountUpdateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountUpdateParams{
		HTTPClient: client,
	}
}

/*V1OverlordsUIDVsphereAccountUpdateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid vsphere account update operation typically these are written to a http.Request
*/
type V1OverlordsUIDVsphereAccountUpdateParams struct {

	/*Body*/
	Body *models.V1OverlordVsphereAccountEntity
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereAccountUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) WithContext(ctx context.Context) *V1OverlordsUIDVsphereAccountUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereAccountUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) WithBody(body *models.V1OverlordVsphereAccountEntity) *V1OverlordsUIDVsphereAccountUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) SetBody(body *models.V1OverlordVsphereAccountEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) WithUID(uid string) *V1OverlordsUIDVsphereAccountUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid vsphere account update params
func (o *V1OverlordsUIDVsphereAccountUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDVsphereAccountUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
