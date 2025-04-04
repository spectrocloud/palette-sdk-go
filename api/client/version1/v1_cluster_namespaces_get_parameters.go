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
	"github.com/go-openapi/swag"
)

// NewV1ClusterNamespacesGetParams creates a new V1ClusterNamespacesGetParams object
// with the default values initialized.
func NewV1ClusterNamespacesGetParams() *V1ClusterNamespacesGetParams {
	var (
		skipEmptyNamespacesDefault = bool(false)
	)
	return &V1ClusterNamespacesGetParams{
		SkipEmptyNamespaces: &skipEmptyNamespacesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterNamespacesGetParamsWithTimeout creates a new V1ClusterNamespacesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterNamespacesGetParamsWithTimeout(timeout time.Duration) *V1ClusterNamespacesGetParams {
	var (
		skipEmptyNamespacesDefault = bool(false)
	)
	return &V1ClusterNamespacesGetParams{
		SkipEmptyNamespaces: &skipEmptyNamespacesDefault,

		timeout: timeout,
	}
}

// NewV1ClusterNamespacesGetParamsWithContext creates a new V1ClusterNamespacesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterNamespacesGetParamsWithContext(ctx context.Context) *V1ClusterNamespacesGetParams {
	var (
		skipEmptyNamespacesDefault = bool(false)
	)
	return &V1ClusterNamespacesGetParams{
		SkipEmptyNamespaces: &skipEmptyNamespacesDefault,

		Context: ctx,
	}
}

// NewV1ClusterNamespacesGetParamsWithHTTPClient creates a new V1ClusterNamespacesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterNamespacesGetParamsWithHTTPClient(client *http.Client) *V1ClusterNamespacesGetParams {
	var (
		skipEmptyNamespacesDefault = bool(false)
	)
	return &V1ClusterNamespacesGetParams{
		SkipEmptyNamespaces: &skipEmptyNamespacesDefault,
		HTTPClient:          client,
	}
}

/*
V1ClusterNamespacesGetParams contains all the parameters to send to the API endpoint
for the v1 cluster namespaces get operation typically these are written to a http.Request
*/
type V1ClusterNamespacesGetParams struct {

	/*SkipEmptyNamespaces*/
	SkipEmptyNamespaces *bool
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) WithTimeout(timeout time.Duration) *V1ClusterNamespacesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) WithContext(ctx context.Context) *V1ClusterNamespacesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) WithHTTPClient(client *http.Client) *V1ClusterNamespacesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkipEmptyNamespaces adds the skipEmptyNamespaces to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) WithSkipEmptyNamespaces(skipEmptyNamespaces *bool) *V1ClusterNamespacesGetParams {
	o.SetSkipEmptyNamespaces(skipEmptyNamespaces)
	return o
}

// SetSkipEmptyNamespaces adds the skipEmptyNamespaces to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) SetSkipEmptyNamespaces(skipEmptyNamespaces *bool) {
	o.SkipEmptyNamespaces = skipEmptyNamespaces
}

// WithUID adds the uid to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) WithUID(uid string) *V1ClusterNamespacesGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster namespaces get params
func (o *V1ClusterNamespacesGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterNamespacesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SkipEmptyNamespaces != nil {

		// query param skipEmptyNamespaces
		var qrSkipEmptyNamespaces bool
		if o.SkipEmptyNamespaces != nil {
			qrSkipEmptyNamespaces = *o.SkipEmptyNamespaces
		}
		qSkipEmptyNamespaces := swag.FormatBool(qrSkipEmptyNamespaces)
		if qSkipEmptyNamespaces != "" {
			if err := r.SetQueryParam("skipEmptyNamespaces", qSkipEmptyNamespaces); err != nil {
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
