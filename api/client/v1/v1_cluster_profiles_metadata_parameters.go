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

// NewV1ClusterProfilesMetadataParams creates a new V1ClusterProfilesMetadataParams object
// with the default values initialized.
func NewV1ClusterProfilesMetadataParams() *V1ClusterProfilesMetadataParams {

	return &V1ClusterProfilesMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesMetadataParamsWithTimeout creates a new V1ClusterProfilesMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesMetadataParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesMetadataParams {

	return &V1ClusterProfilesMetadataParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesMetadataParamsWithContext creates a new V1ClusterProfilesMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesMetadataParamsWithContext(ctx context.Context) *V1ClusterProfilesMetadataParams {

	return &V1ClusterProfilesMetadataParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesMetadataParamsWithHTTPClient creates a new V1ClusterProfilesMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesMetadataParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesMetadataParams {

	return &V1ClusterProfilesMetadataParams{
		HTTPClient: client,
	}
}

/*V1ClusterProfilesMetadataParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles metadata operation typically these are written to a http.Request
*/
type V1ClusterProfilesMetadataParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles metadata params
func (o *V1ClusterProfilesMetadataParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles metadata params
func (o *V1ClusterProfilesMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles metadata params
func (o *V1ClusterProfilesMetadataParams) WithContext(ctx context.Context) *V1ClusterProfilesMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles metadata params
func (o *V1ClusterProfilesMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles metadata params
func (o *V1ClusterProfilesMetadataParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles metadata params
func (o *V1ClusterProfilesMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
