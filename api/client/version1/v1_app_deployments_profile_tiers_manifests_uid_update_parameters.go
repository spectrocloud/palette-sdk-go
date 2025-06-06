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

// NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParams creates a new V1AppDeploymentsProfileTiersManifestsUIDUpdateParams object
// with the default values initialized.
func NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParams() *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParamsWithTimeout creates a new V1AppDeploymentsProfileTiersManifestsUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParamsWithTimeout(timeout time.Duration) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParamsWithContext creates a new V1AppDeploymentsProfileTiersManifestsUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParamsWithContext(ctx context.Context) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParamsWithHTTPClient creates a new V1AppDeploymentsProfileTiersManifestsUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppDeploymentsProfileTiersManifestsUIDUpdateParamsWithHTTPClient(client *http.Client) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1AppDeploymentsProfileTiersManifestsUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 app deployments profile tiers manifests Uid update operation typically these are written to a http.Request
*/
type V1AppDeploymentsProfileTiersManifestsUIDUpdateParams struct {

	/*Body*/
	Body *models.V1ManifestRefUpdateEntity
	/*ManifestUID
	  Application deployment tier manifest uid

	*/
	ManifestUID string
	/*TierUID
	  Application deployment tier uid

	*/
	TierUID string
	/*UID
	  Application deployment uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WithTimeout(timeout time.Duration) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WithContext(ctx context.Context) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WithHTTPClient(client *http.Client) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WithBody(body *models.V1ManifestRefUpdateEntity) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) SetBody(body *models.V1ManifestRefUpdateEntity) {
	o.Body = body
}

// WithManifestUID adds the manifestUID to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WithManifestUID(manifestUID string) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	o.SetManifestUID(manifestUID)
	return o
}

// SetManifestUID adds the manifestUid to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) SetManifestUID(manifestUID string) {
	o.ManifestUID = manifestUID
}

// WithTierUID adds the tierUID to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WithTierUID(tierUID string) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WithUID(uid string) *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app deployments profile tiers manifests Uid update params
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppDeploymentsProfileTiersManifestsUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param manifestUid
	if err := r.SetPathParam("manifestUid", o.ManifestUID); err != nil {
		return err
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
