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

// NewV1OverlordsUIDOpenStackCloudConfigCreateParams creates a new V1OverlordsUIDOpenStackCloudConfigCreateParams object
// with the default values initialized.
func NewV1OverlordsUIDOpenStackCloudConfigCreateParams() *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithTimeout creates a new V1OverlordsUIDOpenStackCloudConfigCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigCreateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithContext creates a new V1OverlordsUIDOpenStackCloudConfigCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithContext(ctx context.Context) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigCreateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithHTTPClient creates a new V1OverlordsUIDOpenStackCloudConfigCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigCreateParams{
		HTTPClient: client,
	}
}

/*V1OverlordsUIDOpenStackCloudConfigCreateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid open stack cloud config create operation typically these are written to a http.Request
*/
type V1OverlordsUIDOpenStackCloudConfigCreateParams struct {

	/*Body*/
	Body *models.V1OverlordOpenStackCloudConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) WithContext(ctx context.Context) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) WithBody(body *models.V1OverlordOpenStackCloudConfig) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) SetBody(body *models.V1OverlordOpenStackCloudConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) WithUID(uid string) *V1OverlordsUIDOpenStackCloudConfigCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid open stack cloud config create params
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDOpenStackCloudConfigCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
