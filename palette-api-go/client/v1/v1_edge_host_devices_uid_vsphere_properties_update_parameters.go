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

// NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParams creates a new V1EdgeHostDevicesUIDVspherePropertiesUpdateParams object
// with the default values initialized.
func NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParams() *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDVspherePropertiesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParamsWithTimeout creates a new V1EdgeHostDevicesUIDVspherePropertiesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParamsWithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDVspherePropertiesUpdateParams{

		timeout: timeout,
	}
}

// NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParamsWithContext creates a new V1EdgeHostDevicesUIDVspherePropertiesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParamsWithContext(ctx context.Context) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDVspherePropertiesUpdateParams{

		Context: ctx,
	}
}

// NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParamsWithHTTPClient creates a new V1EdgeHostDevicesUIDVspherePropertiesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeHostDevicesUIDVspherePropertiesUpdateParamsWithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	var ()
	return &V1EdgeHostDevicesUIDVspherePropertiesUpdateParams{
		HTTPClient: client,
	}
}

/*V1EdgeHostDevicesUIDVspherePropertiesUpdateParams contains all the parameters to send to the API endpoint
for the v1 edge host devices Uid vsphere properties update operation typically these are written to a http.Request
*/
type V1EdgeHostDevicesUIDVspherePropertiesUpdateParams struct {

	/*Body*/
	Body *models.V1EdgeHostVsphereCloudProperties
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) WithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) WithContext(ctx context.Context) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) WithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) WithBody(body *models.V1EdgeHostVsphereCloudProperties) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) SetBody(body *models.V1EdgeHostVsphereCloudProperties) {
	o.Body = body
}

// WithUID adds the uid to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) WithUID(uid string) *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge host devices Uid vsphere properties update params
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeHostDevicesUIDVspherePropertiesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
