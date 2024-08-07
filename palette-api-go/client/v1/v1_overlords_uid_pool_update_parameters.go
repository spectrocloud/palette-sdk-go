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

// NewV1OverlordsUIDPoolUpdateParams creates a new V1OverlordsUIDPoolUpdateParams object
// with the default values initialized.
func NewV1OverlordsUIDPoolUpdateParams() *V1OverlordsUIDPoolUpdateParams {
	var ()
	return &V1OverlordsUIDPoolUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDPoolUpdateParamsWithTimeout creates a new V1OverlordsUIDPoolUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDPoolUpdateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDPoolUpdateParams {
	var ()
	return &V1OverlordsUIDPoolUpdateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDPoolUpdateParamsWithContext creates a new V1OverlordsUIDPoolUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDPoolUpdateParamsWithContext(ctx context.Context) *V1OverlordsUIDPoolUpdateParams {
	var ()
	return &V1OverlordsUIDPoolUpdateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDPoolUpdateParamsWithHTTPClient creates a new V1OverlordsUIDPoolUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDPoolUpdateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDPoolUpdateParams {
	var ()
	return &V1OverlordsUIDPoolUpdateParams{
		HTTPClient: client,
	}
}

/*V1OverlordsUIDPoolUpdateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid pool update operation typically these are written to a http.Request
*/
type V1OverlordsUIDPoolUpdateParams struct {

	/*Body*/
	Body *models.V1IPPoolInputEntity
	/*PoolUID*/
	PoolUID string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDPoolUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) WithContext(ctx context.Context) *V1OverlordsUIDPoolUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDPoolUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) WithBody(body *models.V1IPPoolInputEntity) *V1OverlordsUIDPoolUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) SetBody(body *models.V1IPPoolInputEntity) {
	o.Body = body
}

// WithPoolUID adds the poolUID to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) WithPoolUID(poolUID string) *V1OverlordsUIDPoolUpdateParams {
	o.SetPoolUID(poolUID)
	return o
}

// SetPoolUID adds the poolUid to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) SetPoolUID(poolUID string) {
	o.PoolUID = poolUID
}

// WithUID adds the uid to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) WithUID(uid string) *V1OverlordsUIDPoolUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid pool update params
func (o *V1OverlordsUIDPoolUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDPoolUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param poolUid
	if err := r.SetPathParam("poolUid", o.PoolUID); err != nil {
		return err
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
