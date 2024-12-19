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

// NewV1OverlordsUIDOpenStackCloudConfigUpdateParams creates a new V1OverlordsUIDOpenStackCloudConfigUpdateParams object
// with the default values initialized.
func NewV1OverlordsUIDOpenStackCloudConfigUpdateParams() *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDOpenStackCloudConfigUpdateParamsWithTimeout creates a new V1OverlordsUIDOpenStackCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDOpenStackCloudConfigUpdateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigUpdateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDOpenStackCloudConfigUpdateParamsWithContext creates a new V1OverlordsUIDOpenStackCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDOpenStackCloudConfigUpdateParamsWithContext(ctx context.Context) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigUpdateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDOpenStackCloudConfigUpdateParamsWithHTTPClient creates a new V1OverlordsUIDOpenStackCloudConfigUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDOpenStackCloudConfigUpdateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	var ()
	return &V1OverlordsUIDOpenStackCloudConfigUpdateParams{
		HTTPClient: client,
	}
}

/*V1OverlordsUIDOpenStackCloudConfigUpdateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid open stack cloud config update operation typically these are written to a http.Request
*/
type V1OverlordsUIDOpenStackCloudConfigUpdateParams struct {

	/*Body*/
	Body *models.V1OverlordOpenStackCloudConfig
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) WithContext(ctx context.Context) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) WithBody(body *models.V1OverlordOpenStackCloudConfig) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) SetBody(body *models.V1OverlordOpenStackCloudConfig) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) WithUID(uid string) *V1OverlordsUIDOpenStackCloudConfigUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid open stack cloud config update params
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDOpenStackCloudConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
