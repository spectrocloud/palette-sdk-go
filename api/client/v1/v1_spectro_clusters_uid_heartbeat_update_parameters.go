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

// NewV1SpectroClustersUIDHeartbeatUpdateParams creates a new V1SpectroClustersUIDHeartbeatUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersUIDHeartbeatUpdateParams() *V1SpectroClustersUIDHeartbeatUpdateParams {
	var ()
	return &V1SpectroClustersUIDHeartbeatUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDHeartbeatUpdateParamsWithTimeout creates a new V1SpectroClustersUIDHeartbeatUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDHeartbeatUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDHeartbeatUpdateParams {
	var ()
	return &V1SpectroClustersUIDHeartbeatUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDHeartbeatUpdateParamsWithContext creates a new V1SpectroClustersUIDHeartbeatUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDHeartbeatUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersUIDHeartbeatUpdateParams {
	var ()
	return &V1SpectroClustersUIDHeartbeatUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDHeartbeatUpdateParamsWithHTTPClient creates a new V1SpectroClustersUIDHeartbeatUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDHeartbeatUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDHeartbeatUpdateParams {
	var ()
	return &V1SpectroClustersUIDHeartbeatUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDHeartbeatUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid heartbeat update operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDHeartbeatUpdateParams struct {

	/*Body*/
	Body *models.V1SpectroClusterHeartbeat
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDHeartbeatUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersUIDHeartbeatUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDHeartbeatUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) WithBody(body *models.V1SpectroClusterHeartbeat) *V1SpectroClustersUIDHeartbeatUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) SetBody(body *models.V1SpectroClusterHeartbeat) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) WithUID(uid string) *V1SpectroClustersUIDHeartbeatUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid heartbeat update params
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDHeartbeatUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
