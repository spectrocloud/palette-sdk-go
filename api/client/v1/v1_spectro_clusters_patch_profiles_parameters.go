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
	"github.com/go-openapi/swag"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1SpectroClustersPatchProfilesParams creates a new V1SpectroClustersPatchProfilesParams object
// with the default values initialized.
func NewV1SpectroClustersPatchProfilesParams() *V1SpectroClustersPatchProfilesParams {
	var (
		resolveNotificationDefault = bool(false)
	)
	return &V1SpectroClustersPatchProfilesParams{
		ResolveNotification: &resolveNotificationDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersPatchProfilesParamsWithTimeout creates a new V1SpectroClustersPatchProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersPatchProfilesParamsWithTimeout(timeout time.Duration) *V1SpectroClustersPatchProfilesParams {
	var (
		resolveNotificationDefault = bool(false)
	)
	return &V1SpectroClustersPatchProfilesParams{
		ResolveNotification: &resolveNotificationDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersPatchProfilesParamsWithContext creates a new V1SpectroClustersPatchProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersPatchProfilesParamsWithContext(ctx context.Context) *V1SpectroClustersPatchProfilesParams {
	var (
		resolveNotificationDefault = bool(false)
	)
	return &V1SpectroClustersPatchProfilesParams{
		ResolveNotification: &resolveNotificationDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersPatchProfilesParamsWithHTTPClient creates a new V1SpectroClustersPatchProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersPatchProfilesParamsWithHTTPClient(client *http.Client) *V1SpectroClustersPatchProfilesParams {
	var (
		resolveNotificationDefault = bool(false)
	)
	return &V1SpectroClustersPatchProfilesParams{
		ResolveNotification: &resolveNotificationDefault,
		HTTPClient:          client,
	}
}

/*
V1SpectroClustersPatchProfilesParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters patch profiles operation typically these are written to a http.Request
*/
type V1SpectroClustersPatchProfilesParams struct {

	/*Body*/
	Body *models.V1SpectroClusterProfiles
	/*ResolveNotification
	  Resolve pending cluster notification if set to true

	*/
	ResolveNotification *bool
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) WithTimeout(timeout time.Duration) *V1SpectroClustersPatchProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) WithContext(ctx context.Context) *V1SpectroClustersPatchProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) WithHTTPClient(client *http.Client) *V1SpectroClustersPatchProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) WithBody(body *models.V1SpectroClusterProfiles) *V1SpectroClustersPatchProfilesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) SetBody(body *models.V1SpectroClusterProfiles) {
	o.Body = body
}

// WithResolveNotification adds the resolveNotification to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) WithResolveNotification(resolveNotification *bool) *V1SpectroClustersPatchProfilesParams {
	o.SetResolveNotification(resolveNotification)
	return o
}

// SetResolveNotification adds the resolveNotification to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) SetResolveNotification(resolveNotification *bool) {
	o.ResolveNotification = resolveNotification
}

// WithUID adds the uid to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) WithUID(uid string) *V1SpectroClustersPatchProfilesParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters patch profiles params
func (o *V1SpectroClustersPatchProfilesParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersPatchProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ResolveNotification != nil {

		// query param resolveNotification
		var qrResolveNotification bool
		if o.ResolveNotification != nil {
			qrResolveNotification = *o.ResolveNotification
		}
		qResolveNotification := swag.FormatBool(qrResolveNotification)
		if qResolveNotification != "" {
			if err := r.SetQueryParam("resolveNotification", qResolveNotification); err != nil {
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