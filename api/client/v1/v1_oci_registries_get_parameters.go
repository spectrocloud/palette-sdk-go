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
)

// NewV1OciRegistriesGetParams creates a new V1OciRegistriesGetParams object
// with the default values initialized.
func NewV1OciRegistriesGetParams() *V1OciRegistriesGetParams {
	var ()
	return &V1OciRegistriesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OciRegistriesGetParamsWithTimeout creates a new V1OciRegistriesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OciRegistriesGetParamsWithTimeout(timeout time.Duration) *V1OciRegistriesGetParams {
	var ()
	return &V1OciRegistriesGetParams{

		timeout: timeout,
	}
}

// NewV1OciRegistriesGetParamsWithContext creates a new V1OciRegistriesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OciRegistriesGetParamsWithContext(ctx context.Context) *V1OciRegistriesGetParams {
	var ()
	return &V1OciRegistriesGetParams{

		Context: ctx,
	}
}

// NewV1OciRegistriesGetParamsWithHTTPClient creates a new V1OciRegistriesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OciRegistriesGetParamsWithHTTPClient(client *http.Client) *V1OciRegistriesGetParams {
	var ()
	return &V1OciRegistriesGetParams{
		HTTPClient: client,
	}
}

/*
V1OciRegistriesGetParams contains all the parameters to send to the API endpoint
for the v1 oci registries get operation typically these are written to a http.Request
*/
type V1OciRegistriesGetParams struct {

	/*ClusterUID*/
	ClusterUID *string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) WithTimeout(timeout time.Duration) *V1OciRegistriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) WithContext(ctx context.Context) *V1OciRegistriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) WithHTTPClient(client *http.Client) *V1OciRegistriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterUID adds the clusterUID to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) WithClusterUID(clusterUID *string) *V1OciRegistriesGetParams {
	o.SetClusterUID(clusterUID)
	return o
}

// SetClusterUID adds the clusterUid to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) SetClusterUID(clusterUID *string) {
	o.ClusterUID = clusterUID
}

// WithUID adds the uid to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) WithUID(uid string) *V1OciRegistriesGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 oci registries get params
func (o *V1OciRegistriesGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OciRegistriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterUID != nil {

		// query param clusterUid
		var qrClusterUID string
		if o.ClusterUID != nil {
			qrClusterUID = *o.ClusterUID
		}
		qClusterUID := qrClusterUID
		if qClusterUID != "" {
			if err := r.SetQueryParam("clusterUid", qClusterUID); err != nil {
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