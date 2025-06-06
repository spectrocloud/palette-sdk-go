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

// NewV1TagFilterUIDUpdateParams creates a new V1TagFilterUIDUpdateParams object
// with the default values initialized.
func NewV1TagFilterUIDUpdateParams() *V1TagFilterUIDUpdateParams {
	var ()
	return &V1TagFilterUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TagFilterUIDUpdateParamsWithTimeout creates a new V1TagFilterUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TagFilterUIDUpdateParamsWithTimeout(timeout time.Duration) *V1TagFilterUIDUpdateParams {
	var ()
	return &V1TagFilterUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1TagFilterUIDUpdateParamsWithContext creates a new V1TagFilterUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TagFilterUIDUpdateParamsWithContext(ctx context.Context) *V1TagFilterUIDUpdateParams {
	var ()
	return &V1TagFilterUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1TagFilterUIDUpdateParamsWithHTTPClient creates a new V1TagFilterUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TagFilterUIDUpdateParamsWithHTTPClient(client *http.Client) *V1TagFilterUIDUpdateParams {
	var ()
	return &V1TagFilterUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1TagFilterUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 tag filter Uid update operation typically these are written to a http.Request
*/
type V1TagFilterUIDUpdateParams struct {

	/*Body*/
	Body *models.V1TagFilter
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) WithTimeout(timeout time.Duration) *V1TagFilterUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) WithContext(ctx context.Context) *V1TagFilterUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) WithHTTPClient(client *http.Client) *V1TagFilterUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) WithBody(body *models.V1TagFilter) *V1TagFilterUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) SetBody(body *models.V1TagFilter) {
	o.Body = body
}

// WithUID adds the uid to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) WithUID(uid string) *V1TagFilterUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 tag filter Uid update params
func (o *V1TagFilterUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1TagFilterUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
