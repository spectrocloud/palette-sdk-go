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

// NewV1AppProfilesUIDTiersUIDManifestsCreateParams creates a new V1AppProfilesUIDTiersUIDManifestsCreateParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersUIDManifestsCreateParams() *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsCreateParamsWithTimeout creates a new V1AppProfilesUIDTiersUIDManifestsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersUIDManifestsCreateParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsCreateParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsCreateParamsWithContext creates a new V1AppProfilesUIDTiersUIDManifestsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersUIDManifestsCreateParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsCreateParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsCreateParamsWithHTTPClient creates a new V1AppProfilesUIDTiersUIDManifestsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersUIDManifestsCreateParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsCreateParams{
		HTTPClient: client,
	}
}

/*
V1AppProfilesUIDTiersUIDManifestsCreateParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers Uid manifests create operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersUIDManifestsCreateParams struct {

	/*Body*/
	Body *models.V1ManifestInputEntity
	/*TierUID
	  Application profile tier uid

	*/
	TierUID string
	/*UID
	  Application profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) WithBody(body *models.V1ManifestInputEntity) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) SetBody(body *models.V1ManifestInputEntity) {
	o.Body = body
}

// WithTierUID adds the tierUID to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) WithTierUID(tierUID string) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) WithUID(uid string) *V1AppProfilesUIDTiersUIDManifestsCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers Uid manifests create params
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersUIDManifestsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param tierUid
	if err := r.SetPathParam("tierUid", o.TierUID); err != nil {
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