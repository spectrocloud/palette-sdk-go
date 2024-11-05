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

// NewV1APIKeysUIDUpdateParams creates a new V1APIKeysUIDUpdateParams object
// with the default values initialized.
func NewV1APIKeysUIDUpdateParams() *V1APIKeysUIDUpdateParams {
	var ()
	return &V1APIKeysUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1APIKeysUIDUpdateParamsWithTimeout creates a new V1APIKeysUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1APIKeysUIDUpdateParamsWithTimeout(timeout time.Duration) *V1APIKeysUIDUpdateParams {
	var ()
	return &V1APIKeysUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1APIKeysUIDUpdateParamsWithContext creates a new V1APIKeysUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1APIKeysUIDUpdateParamsWithContext(ctx context.Context) *V1APIKeysUIDUpdateParams {
	var ()
	return &V1APIKeysUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1APIKeysUIDUpdateParamsWithHTTPClient creates a new V1APIKeysUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1APIKeysUIDUpdateParamsWithHTTPClient(client *http.Client) *V1APIKeysUIDUpdateParams {
	var ()
	return &V1APIKeysUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1APIKeysUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 Api keys Uid update operation typically these are written to a http.Request
*/
type V1APIKeysUIDUpdateParams struct {

	/*Body*/
	Body *models.V1APIKeyUpdate
	/*UID
	  Specify API key uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) WithTimeout(timeout time.Duration) *V1APIKeysUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) WithContext(ctx context.Context) *V1APIKeysUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) WithHTTPClient(client *http.Client) *V1APIKeysUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) WithBody(body *models.V1APIKeyUpdate) *V1APIKeysUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) SetBody(body *models.V1APIKeyUpdate) {
	o.Body = body
}

// WithUID adds the uid to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) WithUID(uid string) *V1APIKeysUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 Api keys Uid update params
func (o *V1APIKeysUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1APIKeysUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
