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
)

// NewV1ClusterProfilesUIDPacksManifestsGetParams creates a new V1ClusterProfilesUIDPacksManifestsGetParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDPacksManifestsGetParams() *V1ClusterProfilesUIDPacksManifestsGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksManifestsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDPacksManifestsGetParamsWithTimeout creates a new V1ClusterProfilesUIDPacksManifestsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDPacksManifestsGetParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDPacksManifestsGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksManifestsGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDPacksManifestsGetParamsWithContext creates a new V1ClusterProfilesUIDPacksManifestsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDPacksManifestsGetParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDPacksManifestsGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksManifestsGetParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDPacksManifestsGetParamsWithHTTPClient creates a new V1ClusterProfilesUIDPacksManifestsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDPacksManifestsGetParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDPacksManifestsGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksManifestsGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDPacksManifestsGetParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid packs manifests get operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDPacksManifestsGetParams struct {

	/*IncludePackMeta
	  Comma seperated pack meta such as schema, presets

	*/
	IncludePackMeta *string
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDPacksManifestsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDPacksManifestsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDPacksManifestsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludePackMeta adds the includePackMeta to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) WithIncludePackMeta(includePackMeta *string) *V1ClusterProfilesUIDPacksManifestsGetParams {
	o.SetIncludePackMeta(includePackMeta)
	return o
}

// SetIncludePackMeta adds the includePackMeta to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) SetIncludePackMeta(includePackMeta *string) {
	o.IncludePackMeta = includePackMeta
}

// WithUID adds the uid to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) WithUID(uid string) *V1ClusterProfilesUIDPacksManifestsGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid packs manifests get params
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDPacksManifestsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludePackMeta != nil {

		// query param includePackMeta
		var qrIncludePackMeta string
		if o.IncludePackMeta != nil {
			qrIncludePackMeta = *o.IncludePackMeta
		}
		qIncludePackMeta := qrIncludePackMeta
		if qIncludePackMeta != "" {
			if err := r.SetQueryParam("includePackMeta", qIncludePackMeta); err != nil {
				return err
			}
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
