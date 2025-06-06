// Code generated by go-swagger; DO NOT EDIT.

package version1

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

// NewV1OverlordsUIDMetadataUpdateParams creates a new V1OverlordsUIDMetadataUpdateParams object
// with the default values initialized.
func NewV1OverlordsUIDMetadataUpdateParams() *V1OverlordsUIDMetadataUpdateParams {
	var ()
	return &V1OverlordsUIDMetadataUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDMetadataUpdateParamsWithTimeout creates a new V1OverlordsUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDMetadataUpdateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDMetadataUpdateParams {
	var ()
	return &V1OverlordsUIDMetadataUpdateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDMetadataUpdateParamsWithContext creates a new V1OverlordsUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDMetadataUpdateParamsWithContext(ctx context.Context) *V1OverlordsUIDMetadataUpdateParams {
	var ()
	return &V1OverlordsUIDMetadataUpdateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDMetadataUpdateParamsWithHTTPClient creates a new V1OverlordsUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDMetadataUpdateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDMetadataUpdateParams {
	var ()
	return &V1OverlordsUIDMetadataUpdateParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDMetadataUpdateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid metadata update operation typically these are written to a http.Request
*/
type V1OverlordsUIDMetadataUpdateParams struct {

	/*Body*/
	Body *models.V1ObjectMetaInputEntitySchema
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDMetadataUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) WithContext(ctx context.Context) *V1OverlordsUIDMetadataUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDMetadataUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) WithBody(body *models.V1ObjectMetaInputEntitySchema) *V1OverlordsUIDMetadataUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) SetBody(body *models.V1ObjectMetaInputEntitySchema) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) WithUID(uid string) *V1OverlordsUIDMetadataUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid metadata update params
func (o *V1OverlordsUIDMetadataUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDMetadataUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
