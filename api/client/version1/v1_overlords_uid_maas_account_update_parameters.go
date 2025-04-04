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

// NewV1OverlordsUIDMaasAccountUpdateParams creates a new V1OverlordsUIDMaasAccountUpdateParams object
// with the default values initialized.
func NewV1OverlordsUIDMaasAccountUpdateParams() *V1OverlordsUIDMaasAccountUpdateParams {
	var ()
	return &V1OverlordsUIDMaasAccountUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDMaasAccountUpdateParamsWithTimeout creates a new V1OverlordsUIDMaasAccountUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDMaasAccountUpdateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDMaasAccountUpdateParams {
	var ()
	return &V1OverlordsUIDMaasAccountUpdateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDMaasAccountUpdateParamsWithContext creates a new V1OverlordsUIDMaasAccountUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDMaasAccountUpdateParamsWithContext(ctx context.Context) *V1OverlordsUIDMaasAccountUpdateParams {
	var ()
	return &V1OverlordsUIDMaasAccountUpdateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDMaasAccountUpdateParamsWithHTTPClient creates a new V1OverlordsUIDMaasAccountUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDMaasAccountUpdateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDMaasAccountUpdateParams {
	var ()
	return &V1OverlordsUIDMaasAccountUpdateParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDMaasAccountUpdateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid maas account update operation typically these are written to a http.Request
*/
type V1OverlordsUIDMaasAccountUpdateParams struct {

	/*Body*/
	Body *models.V1OverlordMaasAccountEntity
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDMaasAccountUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) WithContext(ctx context.Context) *V1OverlordsUIDMaasAccountUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDMaasAccountUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) WithBody(body *models.V1OverlordMaasAccountEntity) *V1OverlordsUIDMaasAccountUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) SetBody(body *models.V1OverlordMaasAccountEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) WithUID(uid string) *V1OverlordsUIDMaasAccountUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid maas account update params
func (o *V1OverlordsUIDMaasAccountUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDMaasAccountUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
