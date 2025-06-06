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

// NewV1ClusterProfilesUIDExportTerraformParams creates a new V1ClusterProfilesUIDExportTerraformParams object
// with the default values initialized.
func NewV1ClusterProfilesUIDExportTerraformParams() *V1ClusterProfilesUIDExportTerraformParams {
	var (
		formatDefault = string("yaml")
	)
	return &V1ClusterProfilesUIDExportTerraformParams{
		Format: &formatDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesUIDExportTerraformParamsWithTimeout creates a new V1ClusterProfilesUIDExportTerraformParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesUIDExportTerraformParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesUIDExportTerraformParams {
	var (
		formatDefault = string("yaml")
	)
	return &V1ClusterProfilesUIDExportTerraformParams{
		Format: &formatDefault,

		timeout: timeout,
	}
}

// NewV1ClusterProfilesUIDExportTerraformParamsWithContext creates a new V1ClusterProfilesUIDExportTerraformParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesUIDExportTerraformParamsWithContext(ctx context.Context) *V1ClusterProfilesUIDExportTerraformParams {
	var (
		formatDefault = string("yaml")
	)
	return &V1ClusterProfilesUIDExportTerraformParams{
		Format: &formatDefault,

		Context: ctx,
	}
}

// NewV1ClusterProfilesUIDExportTerraformParamsWithHTTPClient creates a new V1ClusterProfilesUIDExportTerraformParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesUIDExportTerraformParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesUIDExportTerraformParams {
	var (
		formatDefault = string("yaml")
	)
	return &V1ClusterProfilesUIDExportTerraformParams{
		Format:     &formatDefault,
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesUIDExportTerraformParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles Uid export terraform operation typically these are written to a http.Request
*/
type V1ClusterProfilesUIDExportTerraformParams struct {

	/*Format
	  Cluster profile export file format [ "yaml", "json" ]

	*/
	Format *string
	/*UID
	  Cluster profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesUIDExportTerraformParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) WithContext(ctx context.Context) *V1ClusterProfilesUIDExportTerraformParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesUIDExportTerraformParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) WithFormat(format *string) *V1ClusterProfilesUIDExportTerraformParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) SetFormat(format *string) {
	o.Format = format
}

// WithUID adds the uid to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) WithUID(uid string) *V1ClusterProfilesUIDExportTerraformParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cluster profiles Uid export terraform params
func (o *V1ClusterProfilesUIDExportTerraformParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesUIDExportTerraformParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
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
