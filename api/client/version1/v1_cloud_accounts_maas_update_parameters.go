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

// NewV1CloudAccountsMaasUpdateParams creates a new V1CloudAccountsMaasUpdateParams object
// with the default values initialized.
func NewV1CloudAccountsMaasUpdateParams() *V1CloudAccountsMaasUpdateParams {
	var ()
	return &V1CloudAccountsMaasUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsMaasUpdateParamsWithTimeout creates a new V1CloudAccountsMaasUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsMaasUpdateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsMaasUpdateParams {
	var ()
	return &V1CloudAccountsMaasUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsMaasUpdateParamsWithContext creates a new V1CloudAccountsMaasUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsMaasUpdateParamsWithContext(ctx context.Context) *V1CloudAccountsMaasUpdateParams {
	var ()
	return &V1CloudAccountsMaasUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsMaasUpdateParamsWithHTTPClient creates a new V1CloudAccountsMaasUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsMaasUpdateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsMaasUpdateParams {
	var ()
	return &V1CloudAccountsMaasUpdateParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsMaasUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts maas update operation typically these are written to a http.Request
*/
type V1CloudAccountsMaasUpdateParams struct {

	/*Body*/
	Body *models.V1MaasAccount
	/*UID
	  Maas cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsMaasUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) WithContext(ctx context.Context) *V1CloudAccountsMaasUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsMaasUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) WithBody(body *models.V1MaasAccount) *V1CloudAccountsMaasUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) SetBody(body *models.V1MaasAccount) {
	o.Body = body
}

// WithUID adds the uid to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) WithUID(uid string) *V1CloudAccountsMaasUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts maas update params
func (o *V1CloudAccountsMaasUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsMaasUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
