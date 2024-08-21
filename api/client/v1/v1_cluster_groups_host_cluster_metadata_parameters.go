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

// NewV1ClusterGroupsHostClusterMetadataParams creates a new V1ClusterGroupsHostClusterMetadataParams object
// with the default values initialized.
func NewV1ClusterGroupsHostClusterMetadataParams() *V1ClusterGroupsHostClusterMetadataParams {

	return &V1ClusterGroupsHostClusterMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterGroupsHostClusterMetadataParamsWithTimeout creates a new V1ClusterGroupsHostClusterMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterGroupsHostClusterMetadataParamsWithTimeout(timeout time.Duration) *V1ClusterGroupsHostClusterMetadataParams {

	return &V1ClusterGroupsHostClusterMetadataParams{

		timeout: timeout,
	}
}

// NewV1ClusterGroupsHostClusterMetadataParamsWithContext creates a new V1ClusterGroupsHostClusterMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterGroupsHostClusterMetadataParamsWithContext(ctx context.Context) *V1ClusterGroupsHostClusterMetadataParams {

	return &V1ClusterGroupsHostClusterMetadataParams{

		Context: ctx,
	}
}

// NewV1ClusterGroupsHostClusterMetadataParamsWithHTTPClient creates a new V1ClusterGroupsHostClusterMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterGroupsHostClusterMetadataParamsWithHTTPClient(client *http.Client) *V1ClusterGroupsHostClusterMetadataParams {

	return &V1ClusterGroupsHostClusterMetadataParams{
		HTTPClient: client,
	}
}

/*
V1ClusterGroupsHostClusterMetadataParams contains all the parameters to send to the API endpoint
for the v1 cluster groups host cluster metadata operation typically these are written to a http.Request
*/
type V1ClusterGroupsHostClusterMetadataParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster groups host cluster metadata params
func (o *V1ClusterGroupsHostClusterMetadataParams) WithTimeout(timeout time.Duration) *V1ClusterGroupsHostClusterMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster groups host cluster metadata params
func (o *V1ClusterGroupsHostClusterMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster groups host cluster metadata params
func (o *V1ClusterGroupsHostClusterMetadataParams) WithContext(ctx context.Context) *V1ClusterGroupsHostClusterMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster groups host cluster metadata params
func (o *V1ClusterGroupsHostClusterMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster groups host cluster metadata params
func (o *V1ClusterGroupsHostClusterMetadataParams) WithHTTPClient(client *http.Client) *V1ClusterGroupsHostClusterMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster groups host cluster metadata params
func (o *V1ClusterGroupsHostClusterMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterGroupsHostClusterMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}