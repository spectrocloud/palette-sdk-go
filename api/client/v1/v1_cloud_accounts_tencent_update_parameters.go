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

// NewV1CloudAccountsTencentUpdateParams creates a new V1CloudAccountsTencentUpdateParams object
// with the default values initialized.
func NewV1CloudAccountsTencentUpdateParams() *V1CloudAccountsTencentUpdateParams {
	var ()
	return &V1CloudAccountsTencentUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsTencentUpdateParamsWithTimeout creates a new V1CloudAccountsTencentUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsTencentUpdateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsTencentUpdateParams {
	var ()
	return &V1CloudAccountsTencentUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsTencentUpdateParamsWithContext creates a new V1CloudAccountsTencentUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsTencentUpdateParamsWithContext(ctx context.Context) *V1CloudAccountsTencentUpdateParams {
	var ()
	return &V1CloudAccountsTencentUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsTencentUpdateParamsWithHTTPClient creates a new V1CloudAccountsTencentUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsTencentUpdateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsTencentUpdateParams {
	var ()
	return &V1CloudAccountsTencentUpdateParams{
		HTTPClient: client,
	}
}

/*V1CloudAccountsTencentUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts tencent update operation typically these are written to a http.Request
*/
type V1CloudAccountsTencentUpdateParams struct {

	/*Body*/
	Body *models.V1TencentAccount
	/*UID
	  Tencent cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsTencentUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) WithContext(ctx context.Context) *V1CloudAccountsTencentUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsTencentUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) WithBody(body *models.V1TencentAccount) *V1CloudAccountsTencentUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) SetBody(body *models.V1TencentAccount) {
	o.Body = body
}

// WithUID adds the uid to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) WithUID(uid string) *V1CloudAccountsTencentUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts tencent update params
func (o *V1CloudAccountsTencentUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsTencentUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
