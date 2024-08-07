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

	"github.com/spectrocloud/palette-api-go/models"
)

// NewV1CloudConfigsVirtualUIDUpdateParams creates a new V1CloudConfigsVirtualUIDUpdateParams object
// with the default values initialized.
func NewV1CloudConfigsVirtualUIDUpdateParams() *V1CloudConfigsVirtualUIDUpdateParams {
	var ()
	return &V1CloudConfigsVirtualUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsVirtualUIDUpdateParamsWithTimeout creates a new V1CloudConfigsVirtualUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsVirtualUIDUpdateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsVirtualUIDUpdateParams {
	var ()
	return &V1CloudConfigsVirtualUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsVirtualUIDUpdateParamsWithContext creates a new V1CloudConfigsVirtualUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsVirtualUIDUpdateParamsWithContext(ctx context.Context) *V1CloudConfigsVirtualUIDUpdateParams {
	var ()
	return &V1CloudConfigsVirtualUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsVirtualUIDUpdateParamsWithHTTPClient creates a new V1CloudConfigsVirtualUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsVirtualUIDUpdateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsVirtualUIDUpdateParams {
	var ()
	return &V1CloudConfigsVirtualUIDUpdateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsVirtualUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs virtual Uid update operation typically these are written to a http.Request
*/
type V1CloudConfigsVirtualUIDUpdateParams struct {

	/*Body*/
	Body *models.V1VirtualClusterResize
	/*ConfigUID
	  Specify virtual cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsVirtualUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) WithContext(ctx context.Context) *V1CloudConfigsVirtualUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsVirtualUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) WithBody(body *models.V1VirtualClusterResize) *V1CloudConfigsVirtualUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) SetBody(body *models.V1VirtualClusterResize) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) WithConfigUID(configUID string) *V1CloudConfigsVirtualUIDUpdateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs virtual Uid update params
func (o *V1CloudConfigsVirtualUIDUpdateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsVirtualUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param configUid
	if err := r.SetPathParam("configUid", o.ConfigUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
