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

// NewV1ClusterProfilesUIDPacksNameGetParams creates a new V1ClusterProfilesUIDPacksNameGetParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDPacksNameGetParams() *V1ClusterProfilesUIDPacksNameGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDPacksNameGetParamsWithTimeout creates a new V1ClusterProfilesUIDPacksNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDPacksNameGetParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDPacksNameGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksNameGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDPacksNameGetParamsWithContext creates a new V1ClusterProfilesUIDPacksNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDPacksNameGetParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDPacksNameGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksNameGetParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDPacksNameGetParamsWithHTTPClient creates a new V1ClusterProfilesUIDPacksNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDPacksNameGetParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDPacksNameGetParams {
	var ()
	return &V1ClusterProfilesUIDPacksNameGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDPacksNameGetParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid packs name get operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDPacksNameGetParams struct {

	/*PackName
	  Cluster profile pack name

	*/
	PackName string
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDPacksNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDPacksNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDPacksNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackName adds the packName to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) WithPackName(packName string) *V1ClusterProfilesUIDPacksNameGetParams {
	o.SetPackName(packName)
	return o
}

// SetPackName adds the packName to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) SetPackName(packName string) {
	o.PackName = packName
}

// WithUID adds the uid to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) WithUID(uid string) *V1ClusterProfilesUIDPacksNameGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid packs name get params
func (o *V1ClusterProfilesUIDPacksNameGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDPacksNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packName
	if err := r.SetPathParam("packName", o.PackName); err != nil {
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
