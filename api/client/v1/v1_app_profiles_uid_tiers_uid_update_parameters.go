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

// NewV1AppProfilesUIDTiersUIDUpdateParams creates a new V1AppProfilesUIDTiersUIDUpdateParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersUIDUpdateParams() *V1AppProfilesUIDTiersUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersUIDUpdateParamsWithTimeout creates a new V1AppProfilesUIDTiersUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersUIDUpdateParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersUIDUpdateParamsWithContext creates a new V1AppProfilesUIDTiersUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersUIDUpdateParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersUIDUpdateParamsWithHTTPClient creates a new V1AppProfilesUIDTiersUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersUIDUpdateParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDUpdateParams {
	var ()
	return &V1AppProfilesUIDTiersUIDUpdateParams{
		HTTPClient: client,
	}
}

/*V1AppProfilesUIDTiersUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers Uid update operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersUIDUpdateParams struct {

	/*Body*/
	Body *models.V1AppTierUpdateEntity
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

// WithTimeout adds the timeout to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) WithBody(body *models.V1AppTierUpdateEntity) *V1AppProfilesUIDTiersUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) SetBody(body *models.V1AppTierUpdateEntity) {
	o.Body = body
}

// WithTierUID adds the tierUID to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) WithTierUID(tierUID string) *V1AppProfilesUIDTiersUIDUpdateParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) WithUID(uid string) *V1AppProfilesUIDTiersUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers Uid update params
func (o *V1AppProfilesUIDTiersUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
