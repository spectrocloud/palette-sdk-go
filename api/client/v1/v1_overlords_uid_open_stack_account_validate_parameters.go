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
)

// NewV1OverlordsUIDOpenStackAccountValidateParams creates a new V1OverlordsUIDOpenStackAccountValidateParams object
// with the default values initialized.
func NewV1OverlordsUIDOpenStackAccountValidateParams() *V1OverlordsUIDOpenStackAccountValidateParams {
	var ()
	return &V1OverlordsUIDOpenStackAccountValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDOpenStackAccountValidateParamsWithTimeout creates a new V1OverlordsUIDOpenStackAccountValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDOpenStackAccountValidateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackAccountValidateParams {
	var ()
	return &V1OverlordsUIDOpenStackAccountValidateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDOpenStackAccountValidateParamsWithContext creates a new V1OverlordsUIDOpenStackAccountValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDOpenStackAccountValidateParamsWithContext(ctx context.Context) *V1OverlordsUIDOpenStackAccountValidateParams {
	var ()
	return &V1OverlordsUIDOpenStackAccountValidateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDOpenStackAccountValidateParamsWithHTTPClient creates a new V1OverlordsUIDOpenStackAccountValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDOpenStackAccountValidateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackAccountValidateParams {
	var ()
	return &V1OverlordsUIDOpenStackAccountValidateParams{
		HTTPClient: client,
	}
}

/*V1OverlordsUIDOpenStackAccountValidateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid open stack account validate operation typically these are written to a http.Request
*/
type V1OverlordsUIDOpenStackAccountValidateParams struct {

	/*Body*/
	Body V1OverlordsUIDOpenStackAccountValidateBody
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackAccountValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) WithContext(ctx context.Context) *V1OverlordsUIDOpenStackAccountValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackAccountValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) WithBody(body V1OverlordsUIDOpenStackAccountValidateBody) *V1OverlordsUIDOpenStackAccountValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) SetBody(body V1OverlordsUIDOpenStackAccountValidateBody) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) WithUID(uid string) *V1OverlordsUIDOpenStackAccountValidateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid open stack account validate params
func (o *V1OverlordsUIDOpenStackAccountValidateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDOpenStackAccountValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
