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

// NewV1EdgeHostDevicesUIDTunnelConfigUpdateParams creates a new V1EdgeHostDevicesUIDTunnelConfigUpdateParams object
// with the default values initialized.
func NewV1EdgeHostDevicesUIDTunnelConfigUpdateParams() *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDTunnelConfigUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeHostDevicesUIDTunnelConfigUpdateParamsWithTimeout creates a new V1EdgeHostDevicesUIDTunnelConfigUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeHostDevicesUIDTunnelConfigUpdateParamsWithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDTunnelConfigUpdateParams{

		timeout: timeout,
	}
}

// NewV1EdgeHostDevicesUIDTunnelConfigUpdateParamsWithContext creates a new V1EdgeHostDevicesUIDTunnelConfigUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeHostDevicesUIDTunnelConfigUpdateParamsWithContext(ctx context.Context) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDTunnelConfigUpdateParams{

		Context: ctx,
	}
}

// NewV1EdgeHostDevicesUIDTunnelConfigUpdateParamsWithHTTPClient creates a new V1EdgeHostDevicesUIDTunnelConfigUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeHostDevicesUIDTunnelConfigUpdateParamsWithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDTunnelConfigUpdateParams{
		HTTPClient: client,
	}
}

/*
V1EdgeHostDevicesUIDTunnelConfigUpdateParams contains all the parameters to send to the API endpoint
for the v1 edge host devices Uid tunnel config update operation typically these are written to a http.Request
*/
type V1EdgeHostDevicesUIDTunnelConfigUpdateParams struct {

	/*Body*/
	Body *models.V1SpectroTunnelConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) WithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) WithContext(ctx context.Context) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) WithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) WithBody(body *models.V1SpectroTunnelConfig) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) SetBody(body *models.V1SpectroTunnelConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) WithUID(uid string) *V1EdgeHostDevicesUIDTunnelConfigUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge host devices Uid tunnel config update params
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeHostDevicesUIDTunnelConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
