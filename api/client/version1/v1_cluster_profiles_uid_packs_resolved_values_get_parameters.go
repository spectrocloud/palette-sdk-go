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

// NewV1ClusterProfilesUIDPacksResolvedValuesGetParams creates a new V1ClusterProfilesUIDPacksResolvedValuesGetParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDPacksResolvedValuesGetParams() *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksResolvedValuesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDPacksResolvedValuesGetParamsWithTimeout creates a new V1ClusterProfilesUIDPacksResolvedValuesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDPacksResolvedValuesGetParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksResolvedValuesGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDPacksResolvedValuesGetParamsWithContext creates a new V1ClusterProfilesUIDPacksResolvedValuesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDPacksResolvedValuesGetParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksResolvedValuesGetParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDPacksResolvedValuesGetParamsWithHTTPClient creates a new V1ClusterProfilesUIDPacksResolvedValuesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDPacksResolvedValuesGetParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksResolvedValuesGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDPacksResolvedValuesGetParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid packs resolved values get operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDPacksResolvedValuesGetParams struct {

	/*Body*/
	Body *models.V1PackParamsEntity
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) WithBody(body *models.V1PackParamsEntity) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) SetBody(body *models.V1PackParamsEntity) {
	o.Body = body
}

// WithUID adds the uid to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) WithUID(uid string) *V1ClusterProfilesUIDPacksResolvedValuesGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid packs resolved values get params
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDPacksResolvedValuesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
