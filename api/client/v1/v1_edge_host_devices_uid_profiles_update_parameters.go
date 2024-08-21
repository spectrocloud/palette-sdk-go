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

// NewV1EdgeHostDevicesUIDProfilesUpdateParams creates a new V1EdgeHostDevicesUIDProfilesUpdateParams object
// with the default values initialized.
func NewV1EdgeHostDevicesUIDProfilesUpdateParams() *V1EdgeHostDevicesUIDProfilesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDProfilesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeHostDevicesUIDProfilesUpdateParamsWithTimeout creates a new V1EdgeHostDevicesUIDProfilesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeHostDevicesUIDProfilesUpdateParamsWithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDProfilesUpdateParams{

		timeout: timeout,
	}
}

// NewV1EdgeHostDevicesUIDProfilesUpdateParamsWithContext creates a new V1EdgeHostDevicesUIDProfilesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeHostDevicesUIDProfilesUpdateParamsWithContext(ctx context.Context) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDProfilesUpdateParams{

		Context: ctx,
	}
}

// NewV1EdgeHostDevicesUIDProfilesUpdateParamsWithHTTPClient creates a new V1EdgeHostDevicesUIDProfilesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeHostDevicesUIDProfilesUpdateParamsWithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDProfilesUpdateParams{
		HTTPClient: client,
	}
}

/*
V1EdgeHostDevicesUIDProfilesUpdateParams contains all the parameters to send to the API endpoint
for the v1 edge host devices Uid profiles update operation typically these are written to a http.Request
*/
type V1EdgeHostDevicesUIDProfilesUpdateParams struct {

	/*Body*/
	Body *models.V1SpectroClusterProfiles
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) WithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) WithContext(ctx context.Context) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) WithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) WithBody(body *models.V1SpectroClusterProfiles) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) SetBody(body *models.V1SpectroClusterProfiles) {
	o.Body = body
}

// WithUID adds the uid to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) WithUID(uid string) *V1EdgeHostDevicesUIDProfilesUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge host devices Uid profiles update params
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeHostDevicesUIDProfilesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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