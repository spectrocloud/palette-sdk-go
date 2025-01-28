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

// NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParams creates a new V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParams() *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParamsWithTimeout creates a new V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParamsWithContext creates a new V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParamsWithHTTPClient creates a new V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDUpdateParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams{
		HTTPClient: client,
	}
}

/*V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers Uid manifests Uid update operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams struct {

	/*Body*/
	Body *models.V1ManifestRefUpdateEntity
	/*ManifestUID
	  Application profile tier manifest uid

	*/
	ManifestUID string
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

// WithTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WithBody(body *models.V1ManifestRefUpdateEntity) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) SetBody(body *models.V1ManifestRefUpdateEntity) {
	o.Body = body
}

// WithManifestUID adds the manifestUID to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WithManifestUID(manifestUID string) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	o.SetManifestUID(manifestUID)
	return o
}

// SetManifestUID adds the manifestUid to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) SetManifestUID(manifestUID string) {
	o.ManifestUID = manifestUID
}

// WithTierUID adds the tierUID to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WithTierUID(tierUID string) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WithUID(uid string) *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers Uid manifests Uid update params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersUIDManifestsUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
