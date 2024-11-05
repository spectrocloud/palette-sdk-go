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

// NewV1EdgeTokensUIDStateParams creates a new V1EdgeTokensUIDStateParams object
// with the default values initialized.
func NewV1EdgeTokensUIDStateParams() *V1EdgeTokensUIDStateParams {
	var ()
	return &V1EdgeTokensUIDStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeTokensUIDStateParamsWithTimeout creates a new V1EdgeTokensUIDStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeTokensUIDStateParamsWithTimeout(timeout time.Duration) *V1EdgeTokensUIDStateParams {
	var ()
	return &V1EdgeTokensUIDStateParams{

		timeout: timeout,
	}
}

// NewV1EdgeTokensUIDStateParamsWithContext creates a new V1EdgeTokensUIDStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeTokensUIDStateParamsWithContext(ctx context.Context) *V1EdgeTokensUIDStateParams {
	var ()
	return &V1EdgeTokensUIDStateParams{

		Context: ctx,
	}
}

// NewV1EdgeTokensUIDStateParamsWithHTTPClient creates a new V1EdgeTokensUIDStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeTokensUIDStateParamsWithHTTPClient(client *http.Client) *V1EdgeTokensUIDStateParams {
	var ()
	return &V1EdgeTokensUIDStateParams{
		HTTPClient: client,
	}
}

/*
V1EdgeTokensUIDStateParams contains all the parameters to send to the API endpoint
for the v1 edge tokens Uid state operation typically these are written to a http.Request
*/
type V1EdgeTokensUIDStateParams struct {

	/*Body*/
	Body *models.V1EdgeTokenActiveState
	/*UID
	  Edge token uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) WithTimeout(timeout time.Duration) *V1EdgeTokensUIDStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) WithContext(ctx context.Context) *V1EdgeTokensUIDStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) WithHTTPClient(client *http.Client) *V1EdgeTokensUIDStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) WithBody(body *models.V1EdgeTokenActiveState) *V1EdgeTokensUIDStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) SetBody(body *models.V1EdgeTokenActiveState) {
	o.Body = body
}

// WithUID adds the uid to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) WithUID(uid string) *V1EdgeTokensUIDStateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge tokens Uid state params
func (o *V1EdgeTokensUIDStateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeTokensUIDStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
