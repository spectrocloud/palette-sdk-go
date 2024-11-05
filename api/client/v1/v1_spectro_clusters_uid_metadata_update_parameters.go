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

// NewV1SpectroClustersUIDMetadataUpdateParams creates a new V1SpectroClustersUIDMetadataUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersUIDMetadataUpdateParams() *V1SpectroClustersUIDMetadataUpdateParams {
	var ()
	return &V1SpectroClustersUIDMetadataUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDMetadataUpdateParamsWithTimeout creates a new V1SpectroClustersUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDMetadataUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDMetadataUpdateParams {
	var ()
	return &V1SpectroClustersUIDMetadataUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDMetadataUpdateParamsWithContext creates a new V1SpectroClustersUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDMetadataUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersUIDMetadataUpdateParams {
	var ()
	return &V1SpectroClustersUIDMetadataUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDMetadataUpdateParamsWithHTTPClient creates a new V1SpectroClustersUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDMetadataUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDMetadataUpdateParams {
	var ()
	return &V1SpectroClustersUIDMetadataUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDMetadataUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid metadata update operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDMetadataUpdateParams struct {

	/*Body*/
	Body *models.V1ObjectMetaInputEntitySchema
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDMetadataUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersUIDMetadataUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDMetadataUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) WithBody(body *models.V1ObjectMetaInputEntitySchema) *V1SpectroClustersUIDMetadataUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) SetBody(body *models.V1ObjectMetaInputEntitySchema) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) WithUID(uid string) *V1SpectroClustersUIDMetadataUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid metadata update params
func (o *V1SpectroClustersUIDMetadataUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDMetadataUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
