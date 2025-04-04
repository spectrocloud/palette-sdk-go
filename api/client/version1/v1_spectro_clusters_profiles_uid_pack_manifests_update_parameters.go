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

// NewV1SpectroClustersProfilesUIDPackManifestsUpdateParams creates a new V1SpectroClustersProfilesUIDPackManifestsUpdateParams object
// with the default values initialized.
func NewV1SpectroClustersProfilesUIDPackManifestsUpdateParams() *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	var ()
	return &V1SpectroClustersProfilesUIDPackManifestsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersProfilesUIDPackManifestsUpdateParamsWithTimeout creates a new V1SpectroClustersProfilesUIDPackManifestsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersProfilesUIDPackManifestsUpdateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	var ()
	return &V1SpectroClustersProfilesUIDPackManifestsUpdateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersProfilesUIDPackManifestsUpdateParamsWithContext creates a new V1SpectroClustersProfilesUIDPackManifestsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersProfilesUIDPackManifestsUpdateParamsWithContext(ctx context.Context) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	var ()
	return &V1SpectroClustersProfilesUIDPackManifestsUpdateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersProfilesUIDPackManifestsUpdateParamsWithHTTPClient creates a new V1SpectroClustersProfilesUIDPackManifestsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersProfilesUIDPackManifestsUpdateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	var ()
	return &V1SpectroClustersProfilesUIDPackManifestsUpdateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersProfilesUIDPackManifestsUpdateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters profiles Uid pack manifests update operation typically these are written to a http.Request
*/
type V1SpectroClustersProfilesUIDPackManifestsUpdateParams struct {

	/*Body*/
	Body *models.V1ManifestRefInputEntities
	/*PackName
	  Name of the pack

	*/
	PackName string
	/*ProfileUID
	  Cluster profile uid

	*/
	ProfileUID string
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WithContext(ctx context.Context) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WithBody(body *models.V1ManifestRefInputEntities) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) SetBody(body *models.V1ManifestRefInputEntities) {
	o.Body = body
}

// WithPackName adds the packName to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WithPackName(packName string) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	o.SetPackName(packName)
	return o
}

// SetPackName adds the packName to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) SetPackName(packName string) {
	o.PackName = packName
}

// WithProfileUID adds the profileUID to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WithProfileUID(profileUID string) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	o.SetProfileUID(profileUID)
	return o
}

// SetProfileUID adds the profileUid to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) SetProfileUID(profileUID string) {
	o.ProfileUID = profileUID
}

// WithUID adds the uid to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WithUID(uid string) *V1SpectroClustersProfilesUIDPackManifestsUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters profiles Uid pack manifests update params
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersProfilesUIDPackManifestsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param packName
	if err := r.SetPathParam("packName", o.PackName); err != nil {
		return err
	}

	// path param profileUid
	if err := r.SetPathParam("profileUid", o.ProfileUID); err != nil {
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
