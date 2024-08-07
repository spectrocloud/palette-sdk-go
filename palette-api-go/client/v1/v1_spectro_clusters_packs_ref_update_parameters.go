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

// NewV1SpectroClustersPacksRefUpdateParams creates a new V1SpectroClustersPacksRefUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersPacksRefUpdateParams() *V1SpectroClustersPacksRefUpdateParams {
	var ()
	return &V1SpectroClustersPacksRefUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersPacksRefUpdateParamsWithTimeout creates a new V1SpectroClustersPacksRefUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersPacksRefUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersPacksRefUpdateParams {
	var ()
	return &V1SpectroClustersPacksRefUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersPacksRefUpdateParamsWithContext creates a new V1SpectroClustersPacksRefUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersPacksRefUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersPacksRefUpdateParams {
	var ()
	return &V1SpectroClustersPacksRefUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersPacksRefUpdateParamsWithHTTPClient creates a new V1SpectroClustersPacksRefUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersPacksRefUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersPacksRefUpdateParams {
	var ()
	return &V1SpectroClustersPacksRefUpdateParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersPacksRefUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters packs ref update operation typically these are written to a http.Request
*/
type V1SpectroClustersPacksRefUpdateParams struct {

	/*Body*/
	Body *models.V1ClusterNotificationUpdateEntity
	/*Notify*/
	Notify *string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersPacksRefUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersPacksRefUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersPacksRefUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) WithBody(body *models.V1ClusterNotificationUpdateEntity) *V1SpectroClustersPacksRefUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) SetBody(body *models.V1ClusterNotificationUpdateEntity) {
	o.Body = body
}

// WithNotify adds the notify to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) WithNotify(notify *string) *V1SpectroClustersPacksRefUpdateParams {
	o.SetNotify(notify)
	return o
}

// SetNotify adds the notify to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) SetNotify(notify *string) {
	o.Notify = notify
}

// WithUID adds the uid to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) WithUID(uid string) *V1SpectroClustersPacksRefUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters packs ref update params
func (o *V1SpectroClustersPacksRefUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersPacksRefUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Notify != nil {

		// query param notify
		var qrNotify string
		if o.Notify != nil {
			qrNotify = *o.Notify
		}
		qNotify := qrNotify
		if qNotify != "" {
			if err := r.SetQueryParam("notify", qNotify); err != nil {
				return err
			}
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
