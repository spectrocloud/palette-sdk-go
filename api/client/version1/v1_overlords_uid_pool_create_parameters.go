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

// NewV1OverlordsUIDPoolCreateParams creates a new V1OverlordsUIDPoolCreateParams object
// with the default values initialized.
func NewV1OverlordsUIDPoolCreateParams() *V1OverlordsUIDPoolCreateParams {
	var ()
	return &V1OverlordsUIDPoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDPoolCreateParamsWithTimeout creates a new V1OverlordsUIDPoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDPoolCreateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDPoolCreateParams {
	var ()
	return &V1OverlordsUIDPoolCreateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDPoolCreateParamsWithContext creates a new V1OverlordsUIDPoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDPoolCreateParamsWithContext(ctx context.Context) *V1OverlordsUIDPoolCreateParams {
	var ()
	return &V1OverlordsUIDPoolCreateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDPoolCreateParamsWithHTTPClient creates a new V1OverlordsUIDPoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDPoolCreateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDPoolCreateParams {
	var ()
	return &V1OverlordsUIDPoolCreateParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDPoolCreateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid pool create operation typically these are written to a http.Request
*/
type V1OverlordsUIDPoolCreateParams struct {

	/*Body*/
	Body *models.V1IPPoolInputEntity
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDPoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) WithContext(ctx context.Context) *V1OverlordsUIDPoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDPoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) WithBody(body *models.V1IPPoolInputEntity) *V1OverlordsUIDPoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) SetBody(body *models.V1IPPoolInputEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) WithUID(uid string) *V1OverlordsUIDPoolCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid pool create params
func (o *V1OverlordsUIDPoolCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDPoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
