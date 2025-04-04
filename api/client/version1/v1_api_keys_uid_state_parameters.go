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

// NewV1APIKeysUIDStateParams creates a new V1APIKeysUIDStateParams object
// with the default values initialized.
func NewV1APIKeysUIDStateParams() *V1APIKeysUIDStateParams {
	var ()
	return &V1APIKeysUIDStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1APIKeysUIDStateParamsWithTimeout creates a new V1APIKeysUIDStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1APIKeysUIDStateParamsWithTimeout(timeout time.Duration) *V1APIKeysUIDStateParams {
	var ()
	return &V1APIKeysUIDStateParams{

		timeout: timeout,
	}
}

// NewV1APIKeysUIDStateParamsWithContext creates a new V1APIKeysUIDStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1APIKeysUIDStateParamsWithContext(ctx context.Context) *V1APIKeysUIDStateParams {
	var ()
	return &V1APIKeysUIDStateParams{

		Context: ctx,
	}
}

// NewV1APIKeysUIDStateParamsWithHTTPClient creates a new V1APIKeysUIDStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1APIKeysUIDStateParamsWithHTTPClient(client *http.Client) *V1APIKeysUIDStateParams {
	var ()
	return &V1APIKeysUIDStateParams{
		HTTPClient: client,
	}
}

/*
V1APIKeysUIDStateParams contains all the parameters to send to the API endpoint
for the v1 Api keys Uid state operation typically these are written to a http.Request
*/
type V1APIKeysUIDStateParams struct {

	/*Body*/
	Body *models.V1APIKeyActiveState
	/*UID
	  Specify API key uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) WithTimeout(timeout time.Duration) *V1APIKeysUIDStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) WithContext(ctx context.Context) *V1APIKeysUIDStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) WithHTTPClient(client *http.Client) *V1APIKeysUIDStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) WithBody(body *models.V1APIKeyActiveState) *V1APIKeysUIDStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) SetBody(body *models.V1APIKeyActiveState) {
	o.Body = body
}

// WithUID adds the uid to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) WithUID(uid string) *V1APIKeysUIDStateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 Api keys Uid state params
func (o *V1APIKeysUIDStateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1APIKeysUIDStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
