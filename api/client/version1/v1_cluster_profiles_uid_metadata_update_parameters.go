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

// NewV1ClusterProfilesUIDMetadataUpdateParams creates a new V1ClusterProfilesUIDMetadataUpdateParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDMetadataUpdateParams() *V1ClusterProfilesUIDMetadataUpdateParams {
	var ()
	return &V1ClusterProfilesUIDMetadataUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDMetadataUpdateParamsWithTimeout creates a new V1ClusterProfilesUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDMetadataUpdateParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDMetadataUpdateParams {
	var ()
	return &V1ClusterProfilesUIDMetadataUpdateParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDMetadataUpdateParamsWithContext creates a new V1ClusterProfilesUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDMetadataUpdateParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDMetadataUpdateParams {
	var ()
	return &V1ClusterProfilesUIDMetadataUpdateParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDMetadataUpdateParamsWithHTTPClient creates a new V1ClusterProfilesUIDMetadataUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDMetadataUpdateParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDMetadataUpdateParams {
	var ()
	return &V1ClusterProfilesUIDMetadataUpdateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDMetadataUpdateParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid metadata update operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDMetadataUpdateParams struct {

	/*Body*/
	Body *models.V1ProfileMetaEntity
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDMetadataUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDMetadataUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDMetadataUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) WithBody(body *models.V1ProfileMetaEntity) *V1ClusterProfilesUIDMetadataUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) SetBody(body *models.V1ProfileMetaEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) WithUID(uid string) *V1ClusterProfilesUIDMetadataUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid metadata update params
func (o *V1ClusterProfilesUIDMetadataUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDMetadataUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
