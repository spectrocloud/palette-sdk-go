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

// NewV1ClusterGroupsUIDPacksResolvedValuesGetParams creates a new V1ClusterGroupsUIDPacksResolvedValuesGetParams object
// with the default values initialized.
func NewV1ClusterGroupsUIDPacksResolvedValuesGetParams() *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterGroupsUIDPacksResolvedValuesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterGroupsUIDPacksResolvedValuesGetParamsWithTimeout creates a new V1ClusterGroupsUIDPacksResolvedValuesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterGroupsUIDPacksResolvedValuesGetParamsWithTimeout(timeout time.Duration) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterGroupsUIDPacksResolvedValuesGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterGroupsUIDPacksResolvedValuesGetParamsWithContext creates a new V1ClusterGroupsUIDPacksResolvedValuesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterGroupsUIDPacksResolvedValuesGetParamsWithContext(ctx context.Context) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterGroupsUIDPacksResolvedValuesGetParams{

		Context: ctx,
	}
}

// NewV1ClusterGroupsUIDPacksResolvedValuesGetParamsWithHTTPClient creates a new V1ClusterGroupsUIDPacksResolvedValuesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterGroupsUIDPacksResolvedValuesGetParamsWithHTTPClient(client *http.Client) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterGroupsUIDPacksResolvedValuesGetParams{
		HTTPClient: client,
	}
}

/*V1ClusterGroupsUIDPacksResolvedValuesGetParams contains all the parameters to send to the API endpoint
for the v1 cluster groups Uid packs resolved values get operation typically these are written to a http.Request
*/
type V1ClusterGroupsUIDPacksResolvedValuesGetParams struct {

	/*Body*/
	Body *models.V1SpectroClusterProfilesParamReferenceEntity
	/*UID
	  Cluster group uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) WithTimeout(timeout time.Duration) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) WithContext(ctx context.Context) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) WithHTTPClient(client *http.Client) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) WithBody(body *models.V1SpectroClusterProfilesParamReferenceEntity) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) SetBody(body *models.V1SpectroClusterProfilesParamReferenceEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) WithUID(uid string) *V1ClusterGroupsUIDPacksResolvedValuesGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster groups Uid packs resolved values get params
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
