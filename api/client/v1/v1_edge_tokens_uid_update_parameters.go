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

// NewV1EdgeTokensUIDUpdateParams creates a new V1EdgeTokensUIDUpdateParams object
// with the default values initialized.
func NewV1EdgeTokensUIDUpdateParams() *V1EdgeTokensUIDUpdateParams {
	var ()
	return &V1EdgeTokensUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeTokensUIDUpdateParamsWithTimeout creates a new V1EdgeTokensUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeTokensUIDUpdateParamsWithTimeout(timeout time.Duration) *V1EdgeTokensUIDUpdateParams {
	var ()
	return &V1EdgeTokensUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1EdgeTokensUIDUpdateParamsWithContext creates a new V1EdgeTokensUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeTokensUIDUpdateParamsWithContext(ctx context.Context) *V1EdgeTokensUIDUpdateParams {
	var ()
	return &V1EdgeTokensUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1EdgeTokensUIDUpdateParamsWithHTTPClient creates a new V1EdgeTokensUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeTokensUIDUpdateParamsWithHTTPClient(client *http.Client) *V1EdgeTokensUIDUpdateParams {
	var ()
	return &V1EdgeTokensUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1EdgeTokensUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 edge tokens Uid update operation typically these are written to a http.Request
*/
type V1EdgeTokensUIDUpdateParams struct {

	/*Body*/
	Body *models.V1EdgeTokenUpdate
	/*UID
	  Edge token uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) WithTimeout(timeout time.Duration) *V1EdgeTokensUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) WithContext(ctx context.Context) *V1EdgeTokensUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) WithHTTPClient(client *http.Client) *V1EdgeTokensUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) WithBody(body *models.V1EdgeTokenUpdate) *V1EdgeTokensUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) SetBody(body *models.V1EdgeTokenUpdate) {
	o.Body = body
}

// WithUID adds the uid to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) WithUID(uid string) *V1EdgeTokensUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge tokens Uid update params
func (o *V1EdgeTokensUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeTokensUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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